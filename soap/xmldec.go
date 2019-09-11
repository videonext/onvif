// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xml implements a simple XML 1.0 parser that
// understands XML name spaces.
package soap

// References:
//    Annotated XML spec: https://www.xml.com/axml/testaxml.htm
//    XML name spaces: https://www.w3.org/TR/REC-xml-names/

// TODO(rsc):
//	Test error handling.

import (
	"bufio"
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

// A SyntaxError represents a syntax error in the XML input stream.
type SyntaxError struct {
	Msg  string
	Line int
}

func (e *SyntaxError) Error() string {
	return "XML syntax error on line " + strconv.Itoa(e.Line) + ": " + e.Msg
}

// A Name represents an XML name (Local) annotated
// with a name space identifier (Space).
// In tokens returned by Decoder.Token, the Space identifier
// is given as a canonical URL, not the short prefix used
// in the document being parsed.
type Name struct {
	Space, Local string
}

// An Attr represents an attribute in an XML element (Name=Value).
type Attr struct {
	Name  Name
	Value string
}

// A Token is an interface holding one of the token types:
// StartElement, EndElement, CharData, Comment, ProcInst, or Directive.
type Token interface{}

// A StartElement represents an XML start element.
type StartElement struct {
	Name Name
	Attr []Attr
}

// Copy creates a new copy of StartElement.
func (e StartElement) Copy() StartElement {
	attrs := make([]Attr, len(e.Attr))
	copy(attrs, e.Attr)
	e.Attr = attrs
	return e
}

// End returns the corresponding XML end element.
func (e StartElement) End() EndElement {
	return EndElement{e.Name}
}

// An EndElement represents an XML end element.
type EndElement struct {
	Name Name
}

// A CharData represents XML character data (raw text),
// in which XML escape sequences have been replaced by
// the characters they represent.
type CharData []byte

func makeCopy(b []byte) []byte {
	b1 := make([]byte, len(b))
	copy(b1, b)
	return b1
}

// Copy creates a new copy of CharData.
func (c CharData) Copy() CharData { return CharData(makeCopy(c)) }

// A Comment represents an XML comment of the form <!--comment-->.
// The bytes do not include the <!-- and --> comment markers.
type Comment []byte

// Copy creates a new copy of Comment.
func (c Comment) Copy() Comment { return Comment(makeCopy(c)) }

// A ProcInst represents an XML processing instruction of the form <?target inst?>
type ProcInst struct {
	Target string
	Inst   []byte
}

// Copy creates a new copy of ProcInst.
func (p ProcInst) Copy() ProcInst {
	p.Inst = makeCopy(p.Inst)
	return p
}

// A Directive represents an XML directive of the form <!text>.
// The bytes do not include the <! and > markers.
type Directive []byte

// Copy creates a new copy of Directive.
func (d Directive) Copy() Directive { return Directive(makeCopy(d)) }

// CopyToken returns a copy of a Token.
func CopyToken(t Token) Token {
	switch v := t.(type) {
	case CharData:
		return v.Copy()
	case Comment:
		return v.Copy()
	case Directive:
		return v.Copy()
	case ProcInst:
		return v.Copy()
	case StartElement:
		return v.Copy()
	}
	return t
}

// A TokenReader is anything that can decode a stream of XML tokens, including a
// Decoder.
//
// When Token encounters an error or end-of-file condition after successfully
// reading a token, it returns the token. It may return the (non-nil) error from
// the same call or return the error (and a nil token) from a subsequent call.
// An instance of this general case is that a TokenReader returning a non-nil
// token at the end of the token stream may return either io.EOF or a nil error.
// The next Read should return nil, io.EOF.
//
// Implementations of Token are discouraged from returning a nil token with a
// nil error. Callers should treat a return of nil, nil as indicating that
// nothing happened; in particular it does not indicate EOF.
type TokenReader interface {
	Token() (Token, error)
}

// A Decoder represents an XML parser reading a particular input stream.
// The parser assumes that its input is encoded in UTF-8.
type Decoder struct {
	// Strict defaults to true, enforcing the requirements
	// of the XML specification.
	// If set to false, the parser allows input containing common
	// mistakes:
	//	* If an element is missing an end tag, the parser invents
	//	  end tags as necessary to keep the return values from Token
	//	  properly balanced.
	//	* In attribute values and character data, unknown or malformed
	//	  character entities (sequences beginning with &) are left alone.
	//
	// Setting:
	//
	//	d.Strict = false
	//	d.AutoClose = xml.HTMLAutoClose
	//	d.Entity = xml.HTMLEntity
	//
	// creates a parser that can handle typical HTML.
	//
	// Strict mode does not enforce the requirements of the XML name spaces TR.
	// In particular it does not reject name space tags using undefined prefixes.
	// Such tags are recorded with the unknown prefix as the name space URL.
	Strict bool

	// When Strict == false, AutoClose indicates a set of elements to
	// consider closed immediately after they are opened, regardless
	// of whether an end element is present.
	AutoClose []string

	// Entity can be used to map non-standard entity names to string replacements.
	// The parser behaves as if these standard mappings are present in the map,
	// regardless of the actual map content:
	//
	//	"lt": "<",
	//	"gt": ">",
	//	"amp": "&",
	//	"apos": "'",
	//	"quot": `"`,
	Entity map[string]string

	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// non-UTF-8 charset into UTF-8. If CharsetReader is nil or
	// returns an error, parsing stops with an error. One of the
	// CharsetReader's result values must be non-nil.
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)

	// DefaultSpace sets the default name space used for unadorned tags,
	// as if the entire XML stream were wrapped in an element containing
	// the attribute xmlns="DefaultSpace".
	DefaultSpace string

	r              io.ByteReader
	t              TokenReader
	buf            bytes.Buffer
	saved          *bytes.Buffer
	stk            *stack
	free           *stack
	needClose      bool
	toClose        Name
	nextToken      Token
	nextByte       int
	ns             map[string]string
	err            error
	line           int
	offset         int64
	unmarshalDepth int
}

// NewDecoder creates a new XML parser reading from r.
// If r does not implement io.ByteReader, NewDecoder will
// do its own buffering.
func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		ns:       make(map[string]string),
		nextByte: -1,
		line:     1,
		Strict:   true,
	}
	d.switchToReader(r)
	return d
}

// NewTokenDecoder creates a new XML parser using an underlying token stream.
func NewTokenDecoder(t TokenReader) *Decoder {
	// Is it already a Decoder?
	if d, ok := t.(*Decoder); ok {
		return d
	}
	d := &Decoder{
		ns:       make(map[string]string),
		t:        t,
		nextByte: -1,
		line:     1,
		Strict:   true,
	}
	return d
}

// Token returns the next XML token in the input stream.
// At the end of the input stream, Token returns nil, io.EOF.
//
// Slices of bytes in the returned token data refer to the
// parser's internal buffer and remain valid only until the next
// call to Token. To acquire a copy of the bytes, call CopyToken
// or the token's Copy method.
//
// Token expands self-closing elements such as <br/>
// into separate start and end elements returned by successive calls.
//
// Token guarantees that the StartElement and EndElement
// tokens it returns are properly nested and matched:
// if Token encounters an unexpected end element
// or EOF before all expected end elements,
// it will return an error.
//
// Token implements XML name spaces as described by
// https://www.w3.org/TR/REC-xml-names/.  Each of the
// Name structures contained in the Token has the Space
// set to the URL identifying its name space when known.
// If Token encounters an unrecognized name space prefix,
// it uses the prefix as the Space rather than report an error.
func (d *Decoder) Token() (Token, error) {
	var t Token
	var err error
	if d.stk != nil && d.stk.kind == stkEOF {
		return nil, io.EOF
	}
	if d.nextToken != nil {
		t = d.nextToken
		d.nextToken = nil
	} else if t, err = d.rawToken(); err != nil {
		if err == io.EOF && d.stk != nil && d.stk.kind != stkEOF {
			err = d.syntaxError("unexpected EOF")
		}
		return t, err
	}

	if !d.Strict {
		if t1, ok := d.autoClose(t); ok {
			d.nextToken = t
			t = t1
		}
	}
	switch t1 := t.(type) {
	case StartElement:
		// In XML name spaces, the translations listed in the
		// attributes apply to the element name and
		// to the other attribute names, so process
		// the translations first.
		for _, a := range t1.Attr {
			if a.Name.Space == xmlnsPrefix {
				v, ok := d.ns[a.Name.Local]
				d.pushNs(a.Name.Local, v, ok)
				d.ns[a.Name.Local] = a.Value
			}
			if a.Name.Space == "" && a.Name.Local == xmlnsPrefix {
				// Default space for untagged names
				v, ok := d.ns[""]
				d.pushNs("", v, ok)
				d.ns[""] = a.Value
			}
		}

		d.translate(&t1.Name, true)
		for i := range t1.Attr {
			d.translate(&t1.Attr[i].Name, false)
		}
		d.pushElement(t1.Name)
		t = t1

	case EndElement:
		d.translate(&t1.Name, true)
		if !d.popElement(&t1) {
			return nil, d.err
		}
		t = t1
	}
	return t, err
}

const (
	xmlURL      = "http://www.w3.org/XML/1998/namespace"
	xmlnsPrefix = "xmlns"
	xmlPrefix   = "xml"
)

// Apply name space translation to name n.
// The default name space (for Space=="")
// applies only to element names, not to attribute names.
func (d *Decoder) translate(n *Name, isElementName bool) {
	switch {
	case n.Space == xmlnsPrefix:
		return
	case n.Space == "" && !isElementName:
		return
	case n.Space == xmlPrefix:
		n.Space = xmlURL
	case n.Space == "" && n.Local == xmlnsPrefix:
		return
	}
	if v, ok := d.ns[n.Space]; ok {
		n.Space = v
	} else if n.Space == "" {
		n.Space = d.DefaultSpace
	}
}

func (d *Decoder) switchToReader(r io.Reader) {
	// Get efficient byte at a time reader.
	// Assume that if reader has its own
	// ReadByte, it's efficient enough.
	// Otherwise, use bufio.
	if rb, ok := r.(io.ByteReader); ok {
		d.r = rb
	} else {
		d.r = bufio.NewReader(r)
	}
}

// Parsing state - stack holds old name space translations
// and the current set of open elements. The translations to pop when
// ending a given tag are *below* it on the stack, which is
// more work but forced on us by XML.
type stack struct {
	next *stack
	kind int
	name Name
	ok   bool
}

const (
	stkStart = iota
	stkNs
	stkEOF
)

func (d *Decoder) push(kind int) *stack {
	s := d.free
	if s != nil {
		d.free = s.next
	} else {
		s = new(stack)
	}
	s.next = d.stk
	s.kind = kind
	d.stk = s
	return s
}

func (d *Decoder) pop() *stack {
	s := d.stk
	if s != nil {
		d.stk = s.next
		s.next = d.free
		d.free = s
	}
	return s
}

// Record that after the current element is finished
// (that element is already pushed on the stack)
// Token should return EOF until popEOF is called.
func (d *Decoder) pushEOF() {
	// Walk down stack to find Start.
	// It might not be the top, because there might be stkNs
	// entries above it.
	start := d.stk
	for start.kind != stkStart {
		start = start.next
	}
	// The stkNs entries below a start are associated with that
	// element too; skip over them.
	for start.next != nil && start.next.kind == stkNs {
		start = start.next
	}
	s := d.free
	if s != nil {
		d.free = s.next
	} else {
		s = new(stack)
	}
	s.kind = stkEOF
	s.next = start.next
	start.next = s
}

// Undo a pushEOF.
// The element must have been finished, so the EOF should be at the top of the stack.
func (d *Decoder) popEOF() bool {
	if d.stk == nil || d.stk.kind != stkEOF {
		return false
	}
	d.pop()
	return true
}

// Record that we are starting an element with the given name.
func (d *Decoder) pushElement(name Name) {
	s := d.push(stkStart)
	s.name = name
}

// Record that we are changing the value of ns[local].
// The old value is url, ok.
func (d *Decoder) pushNs(local string, url string, ok bool) {
	s := d.push(stkNs)
	s.name.Local = local
	s.name.Space = url
	s.ok = ok
}

// Creates a SyntaxError with the current line number.
func (d *Decoder) syntaxError(msg string) error {
	return &SyntaxError{Msg: msg, Line: d.line}
}

// Record that we are ending an element with the given name.
// The name must match the record at the top of the stack,
// which must be a pushElement record.
// After popping the element, apply any undo records from
// the stack to restore the name translations that existed
// before we saw this element.
func (d *Decoder) popElement(t *EndElement) bool {
	s := d.pop()
	name := t.Name
	switch {
	case s == nil || s.kind != stkStart:
		d.err = d.syntaxError("unexpected end element </" + name.Local + ">")
		return false
	case s.name.Local != name.Local:
		if !d.Strict {
			d.needClose = true
			d.toClose = t.Name
			t.Name = s.name
			return true
		}
		d.err = d.syntaxError("element <" + s.name.Local + "> closed by </" + name.Local + ">")
		return false
	case s.name.Space != name.Space:
		d.err = d.syntaxError("element <" + s.name.Local + "> in space " + s.name.Space +
			"closed by </" + name.Local + "> in space " + name.Space)
		return false
	}

	// Pop stack until a Start or EOF is on the top, undoing the
	// translations that were associated with the element we just closed.
	for d.stk != nil && d.stk.kind != stkStart && d.stk.kind != stkEOF {
		s := d.pop()
		if s.ok {
			d.ns[s.name.Local] = s.name.Space
		} else {
			delete(d.ns, s.name.Local)
		}
	}

	return true
}

// If the top element on the stack is autoclosing and
// t is not the end tag, invent the end tag.
func (d *Decoder) autoClose(t Token) (Token, bool) {
	if d.stk == nil || d.stk.kind != stkStart {
		return nil, false
	}
	name := strings.ToLower(d.stk.name.Local)
	for _, s := range d.AutoClose {
		if strings.ToLower(s) == name {
			// This one should be auto closed if t doesn't close it.
			et, ok := t.(EndElement)
			if !ok || et.Name.Local != name {
				return EndElement{d.stk.name}, true
			}
			break
		}
	}
	return nil, false
}

var errRawToken = errors.New("xml: cannot use RawToken from UnmarshalXML method")

// RawToken is like Token but does not verify that
// start and end elements match and does not translate
// name space prefixes to their corresponding URLs.
func (d *Decoder) RawToken() (Token, error) {
	if d.unmarshalDepth > 0 {
		return nil, errRawToken
	}
	return d.rawToken()
}

func (d *Decoder) rawToken() (Token, error) {
	if d.t != nil {
		return d.t.Token()
	}
	if d.err != nil {
		return nil, d.err
	}
	if d.needClose {
		// The last element we read was self-closing and
		// we returned just the StartElement half.
		// Return the EndElement half now.
		d.needClose = false
		return EndElement{d.toClose}, nil
	}

	b, ok := d.getc()
	if !ok {
		return nil, d.err
	}

	if b != '<' {
		// Text section.
		d.ungetc(b)
		data := d.text(-1, false)
		if data == nil {
			return nil, d.err
		}
		return CharData(data), nil
	}

	if b, ok = d.mustgetc(); !ok {
		return nil, d.err
	}
	switch b {
	case '/':
		// </: End element
		var name Name
		if name, ok = d.nsname(); !ok {
			if d.err == nil {
				d.err = d.syntaxError("expected element name after </")
			}
			return nil, d.err
		}
		d.space()
		if b, ok = d.mustgetc(); !ok {
			return nil, d.err
		}
		if b != '>' {
			d.err = d.syntaxError("invalid characters between </" + name.Local + " and >")
			return nil, d.err
		}
		return EndElement{name}, nil

	case '?':
		// <?: Processing instruction.
		var target string
		if target, ok = d.name(); !ok {
			if d.err == nil {
				d.err = d.syntaxError("expected target name after <?")
			}
			return nil, d.err
		}
		d.space()
		d.buf.Reset()
		var b0 byte
		for {
			if b, ok = d.mustgetc(); !ok {
				return nil, d.err
			}
			d.buf.WriteByte(b)
			if b0 == '?' && b == '>' {
				break
			}
			b0 = b
		}
		data := d.buf.Bytes()
		data = data[0 : len(data)-2] // chop ?>

		if target == "xml" {
			content := string(data)
			ver := procInst("version", content)
			if ver != "" && ver != "1.0" {
				d.err = fmt.Errorf("xml: unsupported version %q; only version 1.0 is supported", ver)
				return nil, d.err
			}
			enc := procInst("encoding", content)
			if enc != "" && enc != "utf-8" && enc != "UTF-8" && !strings.EqualFold(enc, "utf-8") {
				if d.CharsetReader == nil {
					d.err = fmt.Errorf("xml: encoding %q declared but Decoder.CharsetReader is nil", enc)
					return nil, d.err
				}
				newr, err := d.CharsetReader(enc, d.r.(io.Reader))
				if err != nil {
					d.err = fmt.Errorf("xml: opening charset %q: %v", enc, err)
					return nil, d.err
				}
				if newr == nil {
					panic("CharsetReader returned a nil Reader for charset " + enc)
				}
				d.switchToReader(newr)
			}
		}
		return ProcInst{target, data}, nil

	case '!':
		// <!: Maybe comment, maybe CDATA.
		if b, ok = d.mustgetc(); !ok {
			return nil, d.err
		}
		switch b {
		case '-': // <!-
			// Probably <!-- for a comment.
			if b, ok = d.mustgetc(); !ok {
				return nil, d.err
			}
			if b != '-' {
				d.err = d.syntaxError("invalid sequence <!- not part of <!--")
				return nil, d.err
			}
			// Look for terminator.
			d.buf.Reset()
			var b0, b1 byte
			for {
				if b, ok = d.mustgetc(); !ok {
					return nil, d.err
				}
				d.buf.WriteByte(b)
				if b0 == '-' && b1 == '-' {
					if b != '>' {
						d.err = d.syntaxError(
							`invalid sequence "--" not allowed in comments`)
						return nil, d.err
					}
					break
				}
				b0, b1 = b1, b
			}
			data := d.buf.Bytes()
			data = data[0 : len(data)-3] // chop -->
			return Comment(data), nil

		case '[': // <![
			// Probably <![CDATA[.
			for i := 0; i < 6; i++ {
				if b, ok = d.mustgetc(); !ok {
					return nil, d.err
				}
				if b != "CDATA["[i] {
					d.err = d.syntaxError("invalid <![ sequence")
					return nil, d.err
				}
			}
			// Have <![CDATA[.  Read text until ]]>.
			data := d.text(-1, true)
			if data == nil {
				return nil, d.err
			}
			return CharData(data), nil
		}

		// Probably a directive: <!DOCTYPE ...>, <!ENTITY ...>, etc.
		// We don't care, but accumulate for caller. Quoted angle
		// brackets do not count for nesting.
		d.buf.Reset()
		d.buf.WriteByte(b)
		inquote := uint8(0)
		depth := 0
		for {
			if b, ok = d.mustgetc(); !ok {
				return nil, d.err
			}
			if inquote == 0 && b == '>' && depth == 0 {
				break
			}
		HandleB:
			d.buf.WriteByte(b)
			switch {
			case b == inquote:
				inquote = 0

			case inquote != 0:
				// in quotes, no special action

			case b == '\'' || b == '"':
				inquote = b

			case b == '>' && inquote == 0:
				depth--

			case b == '<' && inquote == 0:
				// Look for <!-- to begin comment.
				s := "!--"
				for i := 0; i < len(s); i++ {
					if b, ok = d.mustgetc(); !ok {
						return nil, d.err
					}
					if b != s[i] {
						for j := 0; j < i; j++ {
							d.buf.WriteByte(s[j])
						}
						depth++
						goto HandleB
					}
				}

				// Remove < that was written above.
				d.buf.Truncate(d.buf.Len() - 1)

				// Look for terminator.
				var b0, b1 byte
				for {
					if b, ok = d.mustgetc(); !ok {
						return nil, d.err
					}
					if b0 == '-' && b1 == '-' && b == '>' {
						break
					}
					b0, b1 = b1, b
				}
			}
		}
		return Directive(d.buf.Bytes()), nil
	}

	// Must be an open element like <a href="foo">
	d.ungetc(b)

	var (
		name  Name
		empty bool
		attr  []Attr
	)
	if name, ok = d.nsname(); !ok {
		if d.err == nil {
			d.err = d.syntaxError("expected element name after <")
		}
		return nil, d.err
	}

	attr = []Attr{}
	for {
		d.space()
		if b, ok = d.mustgetc(); !ok {
			return nil, d.err
		}
		if b == '/' {
			empty = true
			if b, ok = d.mustgetc(); !ok {
				return nil, d.err
			}
			if b != '>' {
				d.err = d.syntaxError("expected /> in element")
				return nil, d.err
			}
			break
		}
		if b == '>' {
			break
		}
		d.ungetc(b)

		a := Attr{}
		if a.Name, ok = d.nsname(); !ok {
			if d.err == nil {
				d.err = d.syntaxError("expected attribute name in element")
			}
			return nil, d.err
		}
		d.space()
		if b, ok = d.mustgetc(); !ok {
			return nil, d.err
		}
		if b != '=' {
			if d.Strict {
				d.err = d.syntaxError("attribute name without = in element")
				return nil, d.err
			}
			d.ungetc(b)
			a.Value = a.Name.Local
		} else {
			d.space()
			data := d.attrval()
			if data == nil {
				return nil, d.err
			}
			a.Value = string(data)
		}
		attr = append(attr, a)
	}
	if empty {
		d.needClose = true
		d.toClose = name
	}
	return StartElement{name, attr}, nil
}

func (d *Decoder) attrval() []byte {
	b, ok := d.mustgetc()
	if !ok {
		return nil
	}
	// Handle quoted attribute values
	if b == '"' || b == '\'' {
		return d.text(int(b), false)
	}
	// Handle unquoted attribute values for strict parsers
	if d.Strict {
		d.err = d.syntaxError("unquoted or missing attribute value in element")
		return nil
	}
	// Handle unquoted attribute values for unstrict parsers
	d.ungetc(b)
	d.buf.Reset()
	for {
		b, ok = d.mustgetc()
		if !ok {
			return nil
		}
		// https://www.w3.org/TR/REC-html40/intro/sgmltut.html#h-3.2.2
		if 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' ||
			'0' <= b && b <= '9' || b == '_' || b == ':' || b == '-' {
			d.buf.WriteByte(b)
		} else {
			d.ungetc(b)
			break
		}
	}
	return d.buf.Bytes()
}

// Skip spaces if any
func (d *Decoder) space() {
	for {
		b, ok := d.getc()
		if !ok {
			return
		}
		switch b {
		case ' ', '\r', '\n', '\t':
		default:
			d.ungetc(b)
			return
		}
	}
}

// Read a single byte.
// If there is no byte to read, return ok==false
// and leave the error in d.err.
// Maintain line number.
func (d *Decoder) getc() (b byte, ok bool) {
	if d.err != nil {
		return 0, false
	}
	if d.nextByte >= 0 {
		b = byte(d.nextByte)
		d.nextByte = -1
	} else {
		b, d.err = d.r.ReadByte()
		if d.err != nil {
			return 0, false
		}
		if d.saved != nil {
			d.saved.WriteByte(b)
		}
	}
	if b == '\n' {
		d.line++
	}
	d.offset++
	return b, true
}

// InputOffset returns the input stream byte offset of the current decoder position.
// The offset gives the location of the end of the most recently returned token
// and the beginning of the next token.
func (d *Decoder) InputOffset() int64 {
	return d.offset
}

// Return saved offset.
// If we did ungetc (nextByte >= 0), have to back up one.
func (d *Decoder) savedOffset() int {
	n := d.saved.Len()
	if d.nextByte >= 0 {
		n--
	}
	return n
}

// Must read a single byte.
// If there is no byte to read,
// set d.err to SyntaxError("unexpected EOF")
// and return ok==false
func (d *Decoder) mustgetc() (b byte, ok bool) {
	if b, ok = d.getc(); !ok {
		if d.err == io.EOF {
			d.err = d.syntaxError("unexpected EOF")
		}
	}
	return
}

// Unread a single byte.
func (d *Decoder) ungetc(b byte) {
	if b == '\n' {
		d.line--
	}
	d.nextByte = int(b)
	d.offset--
}

var entity = map[string]int{
	"lt":   '<',
	"gt":   '>',
	"amp":  '&',
	"apos": '\'',
	"quot": '"',
}

// Read plain text section (XML calls it character data).
// If quote >= 0, we are in a quoted string and need to find the matching quote.
// If cdata == true, we are in a <![CDATA[ section and need to find ]]>.
// On failure return nil and leave the error in d.err.
func (d *Decoder) text(quote int, cdata bool) []byte {
	var b0, b1 byte
	var trunc int
	d.buf.Reset()
Input:
	for {
		b, ok := d.getc()
		if !ok {
			if cdata {
				if d.err == io.EOF {
					d.err = d.syntaxError("unexpected EOF in CDATA section")
				}
				return nil
			}
			break Input
		}

		// <![CDATA[ section ends with ]]>.
		// It is an error for ]]> to appear in ordinary text.
		if b0 == ']' && b1 == ']' && b == '>' {
			if cdata {
				trunc = 2
				break Input
			}
			d.err = d.syntaxError("unescaped ]]> not in CDATA section")
			return nil
		}

		// Stop reading text if we see a <.
		if b == '<' && !cdata {
			if quote >= 0 {
				d.err = d.syntaxError("unescaped < inside quoted string")
				return nil
			}
			d.ungetc('<')
			break Input
		}
		if quote >= 0 && b == byte(quote) {
			break Input
		}
		if b == '&' && !cdata {
			// Read escaped character expression up to semicolon.
			// XML in all its glory allows a document to define and use
			// its own character names with <!ENTITY ...> directives.
			// Parsers are required to recognize lt, gt, amp, apos, and quot
			// even if they have not been declared.
			before := d.buf.Len()
			d.buf.WriteByte('&')
			var ok bool
			var text string
			var haveText bool
			if b, ok = d.mustgetc(); !ok {
				return nil
			}
			if b == '#' {
				d.buf.WriteByte(b)
				if b, ok = d.mustgetc(); !ok {
					return nil
				}
				base := 10
				if b == 'x' {
					base = 16
					d.buf.WriteByte(b)
					if b, ok = d.mustgetc(); !ok {
						return nil
					}
				}
				start := d.buf.Len()
				for '0' <= b && b <= '9' ||
					base == 16 && 'a' <= b && b <= 'f' ||
					base == 16 && 'A' <= b && b <= 'F' {
					d.buf.WriteByte(b)
					if b, ok = d.mustgetc(); !ok {
						return nil
					}
				}
				if b != ';' {
					d.ungetc(b)
				} else {
					s := string(d.buf.Bytes()[start:])
					d.buf.WriteByte(';')
					n, err := strconv.ParseUint(s, base, 64)
					if err == nil && n <= unicode.MaxRune {
						text = string(n)
						haveText = true
					}
				}
			} else {
				d.ungetc(b)
				if !d.readName() {
					if d.err != nil {
						return nil
					}
				}
				if b, ok = d.mustgetc(); !ok {
					return nil
				}
				if b != ';' {
					d.ungetc(b)
				} else {
					name := d.buf.Bytes()[before+1:]
					d.buf.WriteByte(';')
					if isName(name) {
						s := string(name)
						if r, ok := entity[s]; ok {
							text = string(r)
							haveText = true
						} else if d.Entity != nil {
							text, haveText = d.Entity[s]
						}
					}
				}
			}

			if haveText {
				d.buf.Truncate(before)
				d.buf.Write([]byte(text))
				b0, b1 = 0, 0
				continue Input
			}
			if !d.Strict {
				b0, b1 = 0, 0
				continue Input
			}
			ent := string(d.buf.Bytes()[before:])
			if ent[len(ent)-1] != ';' {
				ent += " (no semicolon)"
			}
			d.err = d.syntaxError("invalid character entity " + ent)
			return nil
		}

		// We must rewrite unescaped \r and \r\n into \n.
		if b == '\r' {
			d.buf.WriteByte('\n')
		} else if b1 == '\r' && b == '\n' {
			// Skip \r\n--we already wrote \n.
		} else {
			d.buf.WriteByte(b)
		}

		b0, b1 = b1, b
	}
	data := d.buf.Bytes()
	data = data[0 : len(data)-trunc]

	// Inspect each rune for being a disallowed character.
	buf := data
	for len(buf) > 0 {
		r, size := utf8.DecodeRune(buf)
		if r == utf8.RuneError && size == 1 {
			d.err = d.syntaxError("invalid UTF-8")
			return nil
		}
		buf = buf[size:]
		if !isInCharacterRange(r) {
			d.err = d.syntaxError(fmt.Sprintf("illegal character code %U", r))
			return nil
		}
	}

	return data
}

// Decide whether the given rune is in the XML Character Range, per
// the Char production of https://www.xml.com/axml/testaxml.htm,
// Section 2.2 Characters.
func isInCharacterRange(r rune) (inrange bool) {
	return r == 0x09 ||
		r == 0x0A ||
		r == 0x0D ||
		r >= 0x20 && r <= 0xD7FF ||
		r >= 0xE000 && r <= 0xFFFD ||
		r >= 0x10000 && r <= 0x10FFFF
}

// Get name space name: name with a : stuck in the middle.
// The part before the : is the name space identifier.
func (d *Decoder) nsname() (name Name, ok bool) {
	s, ok := d.name()
	if !ok {
		return
	}
	i := strings.Index(s, ":")
	if i < 0 {
		name.Local = s
	} else {
		name.Space = s[0:i]
		name.Local = s[i+1:]
	}
	return name, true
}

// Get name: /first(first|second)*/
// Do not set d.err if the name is missing (unless unexpected EOF is received):
// let the caller provide better context.
func (d *Decoder) name() (s string, ok bool) {
	d.buf.Reset()
	if !d.readName() {
		return "", false
	}

	// Now we check the characters.
	b := d.buf.Bytes()
	if !isName(b) {
		d.err = d.syntaxError("invalid XML name: " + string(b))
		return "", false
	}
	return string(b), true
}

// Read a name and append its bytes to d.buf.
// The name is delimited by any single-byte character not valid in names.
// All multi-byte characters are accepted; the caller must check their validity.
func (d *Decoder) readName() (ok bool) {
	var b byte
	if b, ok = d.mustgetc(); !ok {
		return
	}
	if b < utf8.RuneSelf && !isNameByte(b) {
		d.ungetc(b)
		return false
	}
	d.buf.WriteByte(b)

	for {
		if b, ok = d.mustgetc(); !ok {
			return
		}
		if b < utf8.RuneSelf && !isNameByte(b) {
			d.ungetc(b)
			break
		}
		d.buf.WriteByte(b)
	}
	return true
}

func isNameByte(c byte) bool {
	return 'A' <= c && c <= 'Z' ||
		'a' <= c && c <= 'z' ||
		'0' <= c && c <= '9' ||
		c == '_' || c == ':' || c == '.' || c == '-'
}

func isName(s []byte) bool {
	if len(s) == 0 {
		return false
	}
	c, n := utf8.DecodeRune(s)
	if c == utf8.RuneError && n == 1 {
		return false
	}
	if !unicode.Is(first, c) {
		return false
	}
	for n < len(s) {
		s = s[n:]
		c, n = utf8.DecodeRune(s)
		if c == utf8.RuneError && n == 1 {
			return false
		}
		if !unicode.Is(first, c) && !unicode.Is(second, c) {
			return false
		}
	}
	return true
}

func isNameString(s string) bool {
	if len(s) == 0 {
		return false
	}
	c, n := utf8.DecodeRuneInString(s)
	if c == utf8.RuneError && n == 1 {
		return false
	}
	if !unicode.Is(first, c) {
		return false
	}
	for n < len(s) {
		s = s[n:]
		c, n = utf8.DecodeRuneInString(s)
		if c == utf8.RuneError && n == 1 {
			return false
		}
		if !unicode.Is(first, c) && !unicode.Is(second, c) {
			return false
		}
	}
	return true
}

// These tables were generated by cut and paste from Appendix B of
// the XML spec at https://www.xml.com/axml/testaxml.htm
// and then reformatting. First corresponds to (Letter | '_' | ':')
// and second corresponds to NameChar.

var first = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x003A, 0x003A, 1},
		{0x0041, 0x005A, 1},
		{0x005F, 0x005F, 1},
		{0x0061, 0x007A, 1},
		{0x00C0, 0x00D6, 1},
		{0x00D8, 0x00F6, 1},
		{0x00F8, 0x00FF, 1},
		{0x0100, 0x0131, 1},
		{0x0134, 0x013E, 1},
		{0x0141, 0x0148, 1},
		{0x014A, 0x017E, 1},
		{0x0180, 0x01C3, 1},
		{0x01CD, 0x01F0, 1},
		{0x01F4, 0x01F5, 1},
		{0x01FA, 0x0217, 1},
		{0x0250, 0x02A8, 1},
		{0x02BB, 0x02C1, 1},
		{0x0386, 0x0386, 1},
		{0x0388, 0x038A, 1},
		{0x038C, 0x038C, 1},
		{0x038E, 0x03A1, 1},
		{0x03A3, 0x03CE, 1},
		{0x03D0, 0x03D6, 1},
		{0x03DA, 0x03E0, 2},
		{0x03E2, 0x03F3, 1},
		{0x0401, 0x040C, 1},
		{0x040E, 0x044F, 1},
		{0x0451, 0x045C, 1},
		{0x045E, 0x0481, 1},
		{0x0490, 0x04C4, 1},
		{0x04C7, 0x04C8, 1},
		{0x04CB, 0x04CC, 1},
		{0x04D0, 0x04EB, 1},
		{0x04EE, 0x04F5, 1},
		{0x04F8, 0x04F9, 1},
		{0x0531, 0x0556, 1},
		{0x0559, 0x0559, 1},
		{0x0561, 0x0586, 1},
		{0x05D0, 0x05EA, 1},
		{0x05F0, 0x05F2, 1},
		{0x0621, 0x063A, 1},
		{0x0641, 0x064A, 1},
		{0x0671, 0x06B7, 1},
		{0x06BA, 0x06BE, 1},
		{0x06C0, 0x06CE, 1},
		{0x06D0, 0x06D3, 1},
		{0x06D5, 0x06D5, 1},
		{0x06E5, 0x06E6, 1},
		{0x0905, 0x0939, 1},
		{0x093D, 0x093D, 1},
		{0x0958, 0x0961, 1},
		{0x0985, 0x098C, 1},
		{0x098F, 0x0990, 1},
		{0x0993, 0x09A8, 1},
		{0x09AA, 0x09B0, 1},
		{0x09B2, 0x09B2, 1},
		{0x09B6, 0x09B9, 1},
		{0x09DC, 0x09DD, 1},
		{0x09DF, 0x09E1, 1},
		{0x09F0, 0x09F1, 1},
		{0x0A05, 0x0A0A, 1},
		{0x0A0F, 0x0A10, 1},
		{0x0A13, 0x0A28, 1},
		{0x0A2A, 0x0A30, 1},
		{0x0A32, 0x0A33, 1},
		{0x0A35, 0x0A36, 1},
		{0x0A38, 0x0A39, 1},
		{0x0A59, 0x0A5C, 1},
		{0x0A5E, 0x0A5E, 1},
		{0x0A72, 0x0A74, 1},
		{0x0A85, 0x0A8B, 1},
		{0x0A8D, 0x0A8D, 1},
		{0x0A8F, 0x0A91, 1},
		{0x0A93, 0x0AA8, 1},
		{0x0AAA, 0x0AB0, 1},
		{0x0AB2, 0x0AB3, 1},
		{0x0AB5, 0x0AB9, 1},
		{0x0ABD, 0x0AE0, 0x23},
		{0x0B05, 0x0B0C, 1},
		{0x0B0F, 0x0B10, 1},
		{0x0B13, 0x0B28, 1},
		{0x0B2A, 0x0B30, 1},
		{0x0B32, 0x0B33, 1},
		{0x0B36, 0x0B39, 1},
		{0x0B3D, 0x0B3D, 1},
		{0x0B5C, 0x0B5D, 1},
		{0x0B5F, 0x0B61, 1},
		{0x0B85, 0x0B8A, 1},
		{0x0B8E, 0x0B90, 1},
		{0x0B92, 0x0B95, 1},
		{0x0B99, 0x0B9A, 1},
		{0x0B9C, 0x0B9C, 1},
		{0x0B9E, 0x0B9F, 1},
		{0x0BA3, 0x0BA4, 1},
		{0x0BA8, 0x0BAA, 1},
		{0x0BAE, 0x0BB5, 1},
		{0x0BB7, 0x0BB9, 1},
		{0x0C05, 0x0C0C, 1},
		{0x0C0E, 0x0C10, 1},
		{0x0C12, 0x0C28, 1},
		{0x0C2A, 0x0C33, 1},
		{0x0C35, 0x0C39, 1},
		{0x0C60, 0x0C61, 1},
		{0x0C85, 0x0C8C, 1},
		{0x0C8E, 0x0C90, 1},
		{0x0C92, 0x0CA8, 1},
		{0x0CAA, 0x0CB3, 1},
		{0x0CB5, 0x0CB9, 1},
		{0x0CDE, 0x0CDE, 1},
		{0x0CE0, 0x0CE1, 1},
		{0x0D05, 0x0D0C, 1},
		{0x0D0E, 0x0D10, 1},
		{0x0D12, 0x0D28, 1},
		{0x0D2A, 0x0D39, 1},
		{0x0D60, 0x0D61, 1},
		{0x0E01, 0x0E2E, 1},
		{0x0E30, 0x0E30, 1},
		{0x0E32, 0x0E33, 1},
		{0x0E40, 0x0E45, 1},
		{0x0E81, 0x0E82, 1},
		{0x0E84, 0x0E84, 1},
		{0x0E87, 0x0E88, 1},
		{0x0E8A, 0x0E8D, 3},
		{0x0E94, 0x0E97, 1},
		{0x0E99, 0x0E9F, 1},
		{0x0EA1, 0x0EA3, 1},
		{0x0EA5, 0x0EA7, 2},
		{0x0EAA, 0x0EAB, 1},
		{0x0EAD, 0x0EAE, 1},
		{0x0EB0, 0x0EB0, 1},
		{0x0EB2, 0x0EB3, 1},
		{0x0EBD, 0x0EBD, 1},
		{0x0EC0, 0x0EC4, 1},
		{0x0F40, 0x0F47, 1},
		{0x0F49, 0x0F69, 1},
		{0x10A0, 0x10C5, 1},
		{0x10D0, 0x10F6, 1},
		{0x1100, 0x1100, 1},
		{0x1102, 0x1103, 1},
		{0x1105, 0x1107, 1},
		{0x1109, 0x1109, 1},
		{0x110B, 0x110C, 1},
		{0x110E, 0x1112, 1},
		{0x113C, 0x1140, 2},
		{0x114C, 0x1150, 2},
		{0x1154, 0x1155, 1},
		{0x1159, 0x1159, 1},
		{0x115F, 0x1161, 1},
		{0x1163, 0x1169, 2},
		{0x116D, 0x116E, 1},
		{0x1172, 0x1173, 1},
		{0x1175, 0x119E, 0x119E - 0x1175},
		{0x11A8, 0x11AB, 0x11AB - 0x11A8},
		{0x11AE, 0x11AF, 1},
		{0x11B7, 0x11B8, 1},
		{0x11BA, 0x11BA, 1},
		{0x11BC, 0x11C2, 1},
		{0x11EB, 0x11F0, 0x11F0 - 0x11EB},
		{0x11F9, 0x11F9, 1},
		{0x1E00, 0x1E9B, 1},
		{0x1EA0, 0x1EF9, 1},
		{0x1F00, 0x1F15, 1},
		{0x1F18, 0x1F1D, 1},
		{0x1F20, 0x1F45, 1},
		{0x1F48, 0x1F4D, 1},
		{0x1F50, 0x1F57, 1},
		{0x1F59, 0x1F5B, 0x1F5B - 0x1F59},
		{0x1F5D, 0x1F5D, 1},
		{0x1F5F, 0x1F7D, 1},
		{0x1F80, 0x1FB4, 1},
		{0x1FB6, 0x1FBC, 1},
		{0x1FBE, 0x1FBE, 1},
		{0x1FC2, 0x1FC4, 1},
		{0x1FC6, 0x1FCC, 1},
		{0x1FD0, 0x1FD3, 1},
		{0x1FD6, 0x1FDB, 1},
		{0x1FE0, 0x1FEC, 1},
		{0x1FF2, 0x1FF4, 1},
		{0x1FF6, 0x1FFC, 1},
		{0x2126, 0x2126, 1},
		{0x212A, 0x212B, 1},
		{0x212E, 0x212E, 1},
		{0x2180, 0x2182, 1},
		{0x3007, 0x3007, 1},
		{0x3021, 0x3029, 1},
		{0x3041, 0x3094, 1},
		{0x30A1, 0x30FA, 1},
		{0x3105, 0x312C, 1},
		{0x4E00, 0x9FA5, 1},
		{0xAC00, 0xD7A3, 1},
	},
}

var second = &unicode.RangeTable{
	R16: []unicode.Range16{
		{0x002D, 0x002E, 1},
		{0x0030, 0x0039, 1},
		{0x00B7, 0x00B7, 1},
		{0x02D0, 0x02D1, 1},
		{0x0300, 0x0345, 1},
		{0x0360, 0x0361, 1},
		{0x0387, 0x0387, 1},
		{0x0483, 0x0486, 1},
		{0x0591, 0x05A1, 1},
		{0x05A3, 0x05B9, 1},
		{0x05BB, 0x05BD, 1},
		{0x05BF, 0x05BF, 1},
		{0x05C1, 0x05C2, 1},
		{0x05C4, 0x0640, 0x0640 - 0x05C4},
		{0x064B, 0x0652, 1},
		{0x0660, 0x0669, 1},
		{0x0670, 0x0670, 1},
		{0x06D6, 0x06DC, 1},
		{0x06DD, 0x06DF, 1},
		{0x06E0, 0x06E4, 1},
		{0x06E7, 0x06E8, 1},
		{0x06EA, 0x06ED, 1},
		{0x06F0, 0x06F9, 1},
		{0x0901, 0x0903, 1},
		{0x093C, 0x093C, 1},
		{0x093E, 0x094C, 1},
		{0x094D, 0x094D, 1},
		{0x0951, 0x0954, 1},
		{0x0962, 0x0963, 1},
		{0x0966, 0x096F, 1},
		{0x0981, 0x0983, 1},
		{0x09BC, 0x09BC, 1},
		{0x09BE, 0x09BF, 1},
		{0x09C0, 0x09C4, 1},
		{0x09C7, 0x09C8, 1},
		{0x09CB, 0x09CD, 1},
		{0x09D7, 0x09D7, 1},
		{0x09E2, 0x09E3, 1},
		{0x09E6, 0x09EF, 1},
		{0x0A02, 0x0A3C, 0x3A},
		{0x0A3E, 0x0A3F, 1},
		{0x0A40, 0x0A42, 1},
		{0x0A47, 0x0A48, 1},
		{0x0A4B, 0x0A4D, 1},
		{0x0A66, 0x0A6F, 1},
		{0x0A70, 0x0A71, 1},
		{0x0A81, 0x0A83, 1},
		{0x0ABC, 0x0ABC, 1},
		{0x0ABE, 0x0AC5, 1},
		{0x0AC7, 0x0AC9, 1},
		{0x0ACB, 0x0ACD, 1},
		{0x0AE6, 0x0AEF, 1},
		{0x0B01, 0x0B03, 1},
		{0x0B3C, 0x0B3C, 1},
		{0x0B3E, 0x0B43, 1},
		{0x0B47, 0x0B48, 1},
		{0x0B4B, 0x0B4D, 1},
		{0x0B56, 0x0B57, 1},
		{0x0B66, 0x0B6F, 1},
		{0x0B82, 0x0B83, 1},
		{0x0BBE, 0x0BC2, 1},
		{0x0BC6, 0x0BC8, 1},
		{0x0BCA, 0x0BCD, 1},
		{0x0BD7, 0x0BD7, 1},
		{0x0BE7, 0x0BEF, 1},
		{0x0C01, 0x0C03, 1},
		{0x0C3E, 0x0C44, 1},
		{0x0C46, 0x0C48, 1},
		{0x0C4A, 0x0C4D, 1},
		{0x0C55, 0x0C56, 1},
		{0x0C66, 0x0C6F, 1},
		{0x0C82, 0x0C83, 1},
		{0x0CBE, 0x0CC4, 1},
		{0x0CC6, 0x0CC8, 1},
		{0x0CCA, 0x0CCD, 1},
		{0x0CD5, 0x0CD6, 1},
		{0x0CE6, 0x0CEF, 1},
		{0x0D02, 0x0D03, 1},
		{0x0D3E, 0x0D43, 1},
		{0x0D46, 0x0D48, 1},
		{0x0D4A, 0x0D4D, 1},
		{0x0D57, 0x0D57, 1},
		{0x0D66, 0x0D6F, 1},
		{0x0E31, 0x0E31, 1},
		{0x0E34, 0x0E3A, 1},
		{0x0E46, 0x0E46, 1},
		{0x0E47, 0x0E4E, 1},
		{0x0E50, 0x0E59, 1},
		{0x0EB1, 0x0EB1, 1},
		{0x0EB4, 0x0EB9, 1},
		{0x0EBB, 0x0EBC, 1},
		{0x0EC6, 0x0EC6, 1},
		{0x0EC8, 0x0ECD, 1},
		{0x0ED0, 0x0ED9, 1},
		{0x0F18, 0x0F19, 1},
		{0x0F20, 0x0F29, 1},
		{0x0F35, 0x0F39, 2},
		{0x0F3E, 0x0F3F, 1},
		{0x0F71, 0x0F84, 1},
		{0x0F86, 0x0F8B, 1},
		{0x0F90, 0x0F95, 1},
		{0x0F97, 0x0F97, 1},
		{0x0F99, 0x0FAD, 1},
		{0x0FB1, 0x0FB7, 1},
		{0x0FB9, 0x0FB9, 1},
		{0x20D0, 0x20DC, 1},
		{0x20E1, 0x3005, 0x3005 - 0x20E1},
		{0x302A, 0x302F, 1},
		{0x3031, 0x3035, 1},
		{0x3099, 0x309A, 1},
		{0x309D, 0x309E, 1},
		{0x30FC, 0x30FE, 1},
	},
}

// HTMLEntity is an entity map containing translations for the
// standard HTML entity characters.
//
// See the Decoder.Strict and Decoder.Entity fields' documentation.
var HTMLEntity map[string]string = htmlEntity

var htmlEntity = map[string]string{
	/*
		hget http://www.w3.org/TR/html4/sgml/entities.html |
		ssam '
			,y /\&gt;/ x/\&lt;(.|\n)+/ s/\n/ /g
			,x v/^\&lt;!ENTITY/d
			,s/\&lt;!ENTITY ([^ ]+) .*U\+([0-9A-F][0-9A-F][0-9A-F][0-9A-F]) .+/	"\1": "\\u\2",/g
		'
	*/
	"nbsp":     "\u00A0",
	"iexcl":    "\u00A1",
	"cent":     "\u00A2",
	"pound":    "\u00A3",
	"curren":   "\u00A4",
	"yen":      "\u00A5",
	"brvbar":   "\u00A6",
	"sect":     "\u00A7",
	"uml":      "\u00A8",
	"copy":     "\u00A9",
	"ordf":     "\u00AA",
	"laquo":    "\u00AB",
	"not":      "\u00AC",
	"shy":      "\u00AD",
	"reg":      "\u00AE",
	"macr":     "\u00AF",
	"deg":      "\u00B0",
	"plusmn":   "\u00B1",
	"sup2":     "\u00B2",
	"sup3":     "\u00B3",
	"acute":    "\u00B4",
	"micro":    "\u00B5",
	"para":     "\u00B6",
	"middot":   "\u00B7",
	"cedil":    "\u00B8",
	"sup1":     "\u00B9",
	"ordm":     "\u00BA",
	"raquo":    "\u00BB",
	"frac14":   "\u00BC",
	"frac12":   "\u00BD",
	"frac34":   "\u00BE",
	"iquest":   "\u00BF",
	"Agrave":   "\u00C0",
	"Aacute":   "\u00C1",
	"Acirc":    "\u00C2",
	"Atilde":   "\u00C3",
	"Auml":     "\u00C4",
	"Aring":    "\u00C5",
	"AElig":    "\u00C6",
	"Ccedil":   "\u00C7",
	"Egrave":   "\u00C8",
	"Eacute":   "\u00C9",
	"Ecirc":    "\u00CA",
	"Euml":     "\u00CB",
	"Igrave":   "\u00CC",
	"Iacute":   "\u00CD",
	"Icirc":    "\u00CE",
	"Iuml":     "\u00CF",
	"ETH":      "\u00D0",
	"Ntilde":   "\u00D1",
	"Ograve":   "\u00D2",
	"Oacute":   "\u00D3",
	"Ocirc":    "\u00D4",
	"Otilde":   "\u00D5",
	"Ouml":     "\u00D6",
	"times":    "\u00D7",
	"Oslash":   "\u00D8",
	"Ugrave":   "\u00D9",
	"Uacute":   "\u00DA",
	"Ucirc":    "\u00DB",
	"Uuml":     "\u00DC",
	"Yacute":   "\u00DD",
	"THORN":    "\u00DE",
	"szlig":    "\u00DF",
	"agrave":   "\u00E0",
	"aacute":   "\u00E1",
	"acirc":    "\u00E2",
	"atilde":   "\u00E3",
	"auml":     "\u00E4",
	"aring":    "\u00E5",
	"aelig":    "\u00E6",
	"ccedil":   "\u00E7",
	"egrave":   "\u00E8",
	"eacute":   "\u00E9",
	"ecirc":    "\u00EA",
	"euml":     "\u00EB",
	"igrave":   "\u00EC",
	"iacute":   "\u00ED",
	"icirc":    "\u00EE",
	"iuml":     "\u00EF",
	"eth":      "\u00F0",
	"ntilde":   "\u00F1",
	"ograve":   "\u00F2",
	"oacute":   "\u00F3",
	"ocirc":    "\u00F4",
	"otilde":   "\u00F5",
	"ouml":     "\u00F6",
	"divide":   "\u00F7",
	"oslash":   "\u00F8",
	"ugrave":   "\u00F9",
	"uacute":   "\u00FA",
	"ucirc":    "\u00FB",
	"uuml":     "\u00FC",
	"yacute":   "\u00FD",
	"thorn":    "\u00FE",
	"yuml":     "\u00FF",
	"fnof":     "\u0192",
	"Alpha":    "\u0391",
	"Beta":     "\u0392",
	"Gamma":    "\u0393",
	"Delta":    "\u0394",
	"Epsilon":  "\u0395",
	"Zeta":     "\u0396",
	"Eta":      "\u0397",
	"Theta":    "\u0398",
	"Iota":     "\u0399",
	"Kappa":    "\u039A",
	"Lambda":   "\u039B",
	"Mu":       "\u039C",
	"Nu":       "\u039D",
	"Xi":       "\u039E",
	"Omicron":  "\u039F",
	"Pi":       "\u03A0",
	"Rho":      "\u03A1",
	"Sigma":    "\u03A3",
	"Tau":      "\u03A4",
	"Upsilon":  "\u03A5",
	"Phi":      "\u03A6",
	"Chi":      "\u03A7",
	"Psi":      "\u03A8",
	"Omega":    "\u03A9",
	"alpha":    "\u03B1",
	"beta":     "\u03B2",
	"gamma":    "\u03B3",
	"delta":    "\u03B4",
	"epsilon":  "\u03B5",
	"zeta":     "\u03B6",
	"eta":      "\u03B7",
	"theta":    "\u03B8",
	"iota":     "\u03B9",
	"kappa":    "\u03BA",
	"lambda":   "\u03BB",
	"mu":       "\u03BC",
	"nu":       "\u03BD",
	"xi":       "\u03BE",
	"omicron":  "\u03BF",
	"pi":       "\u03C0",
	"rho":      "\u03C1",
	"sigmaf":   "\u03C2",
	"sigma":    "\u03C3",
	"tau":      "\u03C4",
	"upsilon":  "\u03C5",
	"phi":      "\u03C6",
	"chi":      "\u03C7",
	"psi":      "\u03C8",
	"omega":    "\u03C9",
	"thetasym": "\u03D1",
	"upsih":    "\u03D2",
	"piv":      "\u03D6",
	"bull":     "\u2022",
	"hellip":   "\u2026",
	"prime":    "\u2032",
	"Prime":    "\u2033",
	"oline":    "\u203E",
	"frasl":    "\u2044",
	"weierp":   "\u2118",
	"image":    "\u2111",
	"real":     "\u211C",
	"trade":    "\u2122",
	"alefsym":  "\u2135",
	"larr":     "\u2190",
	"uarr":     "\u2191",
	"rarr":     "\u2192",
	"darr":     "\u2193",
	"harr":     "\u2194",
	"crarr":    "\u21B5",
	"lArr":     "\u21D0",
	"uArr":     "\u21D1",
	"rArr":     "\u21D2",
	"dArr":     "\u21D3",
	"hArr":     "\u21D4",
	"forall":   "\u2200",
	"part":     "\u2202",
	"exist":    "\u2203",
	"empty":    "\u2205",
	"nabla":    "\u2207",
	"isin":     "\u2208",
	"notin":    "\u2209",
	"ni":       "\u220B",
	"prod":     "\u220F",
	"sum":      "\u2211",
	"minus":    "\u2212",
	"lowast":   "\u2217",
	"radic":    "\u221A",
	"prop":     "\u221D",
	"infin":    "\u221E",
	"ang":      "\u2220",
	"and":      "\u2227",
	"or":       "\u2228",
	"cap":      "\u2229",
	"cup":      "\u222A",
	"int":      "\u222B",
	"there4":   "\u2234",
	"sim":      "\u223C",
	"cong":     "\u2245",
	"asymp":    "\u2248",
	"ne":       "\u2260",
	"equiv":    "\u2261",
	"le":       "\u2264",
	"ge":       "\u2265",
	"sub":      "\u2282",
	"sup":      "\u2283",
	"nsub":     "\u2284",
	"sube":     "\u2286",
	"supe":     "\u2287",
	"oplus":    "\u2295",
	"otimes":   "\u2297",
	"perp":     "\u22A5",
	"sdot":     "\u22C5",
	"lceil":    "\u2308",
	"rceil":    "\u2309",
	"lfloor":   "\u230A",
	"rfloor":   "\u230B",
	"lang":     "\u2329",
	"rang":     "\u232A",
	"loz":      "\u25CA",
	"spades":   "\u2660",
	"clubs":    "\u2663",
	"hearts":   "\u2665",
	"diams":    "\u2666",
	"quot":     "\u0022",
	"amp":      "\u0026",
	"lt":       "\u003C",
	"gt":       "\u003E",
	"OElig":    "\u0152",
	"oelig":    "\u0153",
	"Scaron":   "\u0160",
	"scaron":   "\u0161",
	"Yuml":     "\u0178",
	"circ":     "\u02C6",
	"tilde":    "\u02DC",
	"ensp":     "\u2002",
	"emsp":     "\u2003",
	"thinsp":   "\u2009",
	"zwnj":     "\u200C",
	"zwj":      "\u200D",
	"lrm":      "\u200E",
	"rlm":      "\u200F",
	"ndash":    "\u2013",
	"mdash":    "\u2014",
	"lsquo":    "\u2018",
	"rsquo":    "\u2019",
	"sbquo":    "\u201A",
	"ldquo":    "\u201C",
	"rdquo":    "\u201D",
	"bdquo":    "\u201E",
	"dagger":   "\u2020",
	"Dagger":   "\u2021",
	"permil":   "\u2030",
	"lsaquo":   "\u2039",
	"rsaquo":   "\u203A",
	"euro":     "\u20AC",
}

// HTMLAutoClose is the set of HTML elements that
// should be considered to close automatically.
//
// See the Decoder.Strict and Decoder.Entity fields' documentation.
var HTMLAutoClose []string = htmlAutoClose

var htmlAutoClose = []string{
	/*
		hget http://www.w3.org/TR/html4/loose.dtd |
		9 sed -n 's/<!ELEMENT ([^ ]*) +- O EMPTY.+/	"\1",/p' | tr A-Z a-z
	*/
	"basefont",
	"br",
	"area",
	"link",
	"img",
	"param",
	"hr",
	"input",
	"col",
	"frame",
	"isindex",
	"base",
	"meta",
}

var (
	escQuot = []byte("&#34;") // shorter than "&quot;"
	escApos = []byte("&#39;") // shorter than "&apos;"
	escAmp  = []byte("&amp;")
	escLT   = []byte("&lt;")
	escGT   = []byte("&gt;")
	escTab  = []byte("&#x9;")
	escNL   = []byte("&#xA;")
	escCR   = []byte("&#xD;")
	escFFFD = []byte("\uFFFD") // Unicode replacement character
)

// EscapeText writes to w the properly escaped XML equivalent
// of the plain text data s.
func EscapeText(w io.Writer, s []byte) error {
	return escapeText(w, s, true)
}

// escapeText writes to w the properly escaped XML equivalent
// of the plain text data s. If escapeNewline is true, newline
// characters will be escaped.
func escapeText(w io.Writer, s []byte, escapeNewline bool) error {
	var esc []byte
	last := 0
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRune(s[i:])
		i += width
		switch r {
		case '"':
			esc = escQuot
		case '\'':
			esc = escApos
		case '&':
			esc = escAmp
		case '<':
			esc = escLT
		case '>':
			esc = escGT
		case '\t':
			esc = escTab
		case '\n':
			if !escapeNewline {
				continue
			}
			esc = escNL
		case '\r':
			esc = escCR
		default:
			if !isInCharacterRange(r) || (r == 0xFFFD && width == 1) {
				esc = escFFFD
				break
			}
			continue
		}
		if _, err := w.Write(s[last : i-width]); err != nil {
			return err
		}
		if _, err := w.Write(esc); err != nil {
			return err
		}
		last = i
	}
	_, err := w.Write(s[last:])
	return err
}

// EscapeString writes to p the properly escaped XML equivalent
// of the plain text data s.
func (p *printer) EscapeString(s string) {
	var esc []byte
	last := 0
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRuneInString(s[i:])
		i += width
		switch r {
		case '"':
			esc = escQuot
		case '\'':
			esc = escApos
		case '&':
			esc = escAmp
		case '<':
			esc = escLT
		case '>':
			esc = escGT
		case '\t':
			esc = escTab
		case '\n':
			esc = escNL
		case '\r':
			esc = escCR
		default:
			if !isInCharacterRange(r) || (r == 0xFFFD && width == 1) {
				esc = escFFFD
				break
			}
			continue
		}
		p.WriteString(s[last : i-width])
		p.Write(esc)
		last = i
	}
	p.WriteString(s[last:])
}

// Escape is like EscapeText but omits the error return value.
// It is provided for backwards compatibility with Go 1.0.
// Code targeting Go 1.1 or later should use EscapeText.
func Escape(w io.Writer, s []byte) {
	EscapeText(w, s)
}

var (
	cdataStart  = []byte("<![CDATA[")
	cdataEnd    = []byte("]]>")
	cdataEscape = []byte("]]]]><![CDATA[>")
)

// emitCDATA writes to w the CDATA-wrapped plain text data s.
// It escapes CDATA directives nested in s.
func emitCDATA(w io.Writer, s []byte) error {
	if len(s) == 0 {
		return nil
	}
	if _, err := w.Write(cdataStart); err != nil {
		return err
	}
	for {
		i := bytes.Index(s, cdataEnd)
		if i >= 0 && i+len(cdataEnd) <= len(s) {
			// Found a nested CDATA directive end.
			if _, err := w.Write(s[:i]); err != nil {
				return err
			}
			if _, err := w.Write(cdataEscape); err != nil {
				return err
			}
			i += len(cdataEnd)
		} else {
			if _, err := w.Write(s); err != nil {
				return err
			}
			break
		}
		s = s[i:]
	}
	_, err := w.Write(cdataEnd)
	return err
}

// procInst parses the `param="..."` or `param='...'`
// value out of the provided string, returning "" if not found.
func procInst(param, s string) string {
	// TODO: this parsing is somewhat lame and not exact.
	// It works for all actual cases, though.
	param = param + "="
	idx := strings.Index(s, param)
	if idx == -1 {
		return ""
	}
	v := s[idx+len(param):]
	if v == "" {
		return ""
	}
	if v[0] != '\'' && v[0] != '"' {
		return ""
	}
	idx = strings.IndexRune(v[1:], rune(v[0]))
	if idx == -1 {
		return ""
	}
	return v[1 : idx+1]
}

type printer struct {
	*bufio.Writer
	seq        int
	indent     string
	prefix     string
	depth      int
	indentedIn bool
	putNewline bool
	attrNS     map[string]string // map prefix -> name space
	attrPrefix map[string]string // map name space -> prefix
	prefixes   []string
	tags       []Name
}

// read.go
// BUG(rsc): Mapping between XML elements and data structures is inherently flawed:
// an XML element is an order-dependent collection of anonymous
// values, while a data structure is an order-independent collection
// of named values.
// See package json for a textual representation more suitable
// to data structures.

// Unmarshal parses the XML-encoded data and stores the result in
// the value pointed to by v, which must be an arbitrary struct,
// slice, or string. Well-formed data that does not fit into v is
// discarded.
//
// Because Unmarshal uses the reflect package, it can only assign
// to exported (upper case) fields. Unmarshal uses a case-sensitive
// comparison to match XML element names to tag values and struct
// field names.
//
// Unmarshal maps an XML element to a struct using the following rules.
// In the rules, the tag of a field refers to the value associated with the
// key 'xml' in the struct field's tag (see the example above).
//
//   * If the struct has a field of type []byte or string with tag
//      ",innerxml", Unmarshal accumulates the raw XML nested inside the
//      element in that field. The rest of the rules still apply.
//
//   * If the struct has a field named XMLName of type Name,
//      Unmarshal records the element name in that field.
//
//   * If the XMLName field has an associated tag of the form
//      "name" or "namespace-URL name", the XML element must have
//      the given name (and, optionally, name space) or else Unmarshal
//      returns an error.
//
//   * If the XML element has an attribute whose name matches a
//      struct field name with an associated tag containing ",attr" or
//      the explicit name in a struct field tag of the form "name,attr",
//      Unmarshal records the attribute value in that field.
//
//   * If the XML element has an attribute not handled by the previous
//      rule and the struct has a field with an associated tag containing
//      ",any,attr", Unmarshal records the attribute value in the first
//      such field.
//
//   * If the XML element contains character data, that data is
//      accumulated in the first struct field that has tag ",chardata".
//      The struct field may have type []byte or string.
//      If there is no such field, the character data is discarded.
//
//   * If the XML element contains comments, they are accumulated in
//      the first struct field that has tag ",comment".  The struct
//      field may have type []byte or string. If there is no such
//      field, the comments are discarded.
//
//   * If the XML element contains a sub-element whose name matches
//      the prefix of a tag formatted as "a" or "a>b>c", unmarshal
//      will descend into the XML structure looking for elements with the
//      given names, and will map the innermost elements to that struct
//      field. A tag starting with ">" is equivalent to one starting
//      with the field name followed by ">".
//
//   * If the XML element contains a sub-element whose name matches
//      a struct field's XMLName tag and the struct field has no
//      explicit name tag as per the previous rule, unmarshal maps
//      the sub-element to that struct field.
//
//   * If the XML element contains a sub-element whose name matches a
//      field without any mode flags (",attr", ",chardata", etc), Unmarshal
//      maps the sub-element to that struct field.
//
//   * If the XML element contains a sub-element that hasn't matched any
//      of the above rules and the struct has a field with tag ",any",
//      unmarshal maps the sub-element to that struct field.
//
//   * An anonymous struct field is handled as if the fields of its
//      value were part of the outer struct.
//
//   * A struct field with tag "-" is never unmarshaled into.
//
// If Unmarshal encounters a field type that implements the Unmarshaler
// interface, Unmarshal calls its UnmarshalXML method to produce the value from
// the XML element.  Otherwise, if the value implements
// encoding.TextUnmarshaler, Unmarshal calls that value's UnmarshalText method.
//
// Unmarshal maps an XML element to a string or []byte by saving the
// concatenation of that element's character data in the string or
// []byte. The saved []byte is never nil.
//
// Unmarshal maps an attribute value to a string or []byte by saving
// the value in the string or slice.
//
// Unmarshal maps an attribute value to an Attr by saving the attribute,
// including its name, in the Attr.
//
// Unmarshal maps an XML element or attribute value to a slice by
// extending the length of the slice and mapping the element or attribute
// to the newly created value.
//
// Unmarshal maps an XML element or attribute value to a bool by
// setting it to the boolean value represented by the string. Whitespace
// is trimmed and ignored.
//
// Unmarshal maps an XML element or attribute value to an integer or
// floating-point field by setting the field to the result of
// interpreting the string value in decimal. There is no check for
// overflow. Whitespace is trimmed and ignored.
//
// Unmarshal maps an XML element to a Name by recording the element
// name.
//
// Unmarshal maps an XML element to a pointer by setting the pointer
// to a freshly allocated value and then mapping the element to that value.
//
// A missing element or empty attribute value will be unmarshaled as a zero value.
// If the field is a slice, a zero value will be appended to the field. Otherwise, the
// field will be set to its zero value.
func Unmarshal(data []byte, v interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(v)
}

// Decode works like Unmarshal, except it reads the decoder
// stream to find the start element.
func (d *Decoder) Decode(v interface{}) error {
	return d.DecodeElement(v, nil)
}

// DecodeElement works like Unmarshal except that it takes
// a pointer to the start XML element to decode into v.
// It is useful when a client reads some raw XML tokens itself
// but also wants to defer to Unmarshal for some elements.
func (d *Decoder) DecodeElement(v interface{}, start *StartElement) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("non-pointer passed to Unmarshal")
	}
	return d.unmarshal(val.Elem(), start)
}

// An UnmarshalError represents an error in the unmarshaling process.
type UnmarshalError string

func (e UnmarshalError) Error() string { return string(e) }

// Unmarshaler is the interface implemented by objects that can unmarshal
// an XML element description of themselves.
//
// UnmarshalXML decodes a single XML element
// beginning with the given start element.
// If it returns an error, the outer call to Unmarshal stops and
// returns that error.
// UnmarshalXML must consume exactly one XML element.
// One common implementation strategy is to unmarshal into
// a separate value with a layout matching the expected XML
// using d.DecodeElement, and then to copy the data from
// that value into the receiver.
// Another common strategy is to use d.Token to process the
// XML object one token at a time.
// UnmarshalXML may not use d.RawToken.
type Unmarshaler interface {
	UnmarshalXML(d *Decoder, start StartElement) error
}

// UnmarshalerAttr is the interface implemented by objects that can unmarshal
// an XML attribute description of themselves.
//
// UnmarshalXMLAttr decodes a single XML attribute.
// If it returns an error, the outer call to Unmarshal stops and
// returns that error.
// UnmarshalXMLAttr is used only for struct fields with the
// "attr" option in the field tag.
type UnmarshalerAttr interface {
	UnmarshalXMLAttr(attr Attr) error
}

// receiverType returns the receiver type to use in an expression like "%s.MethodName".
func receiverType(val interface{}) string {
	t := reflect.TypeOf(val)
	if t.Name() != "" {
		return t.String()
	}
	return "(" + t.String() + ")"
}

// unmarshalInterface unmarshals a single XML element into val.
// start is the opening tag of the element.
func (d *Decoder) unmarshalInterface(val Unmarshaler, start *StartElement) error {
	// Record that decoder must stop at end tag corresponding to start.
	d.pushEOF()

	d.unmarshalDepth++
	err := val.UnmarshalXML(d, *start)
	d.unmarshalDepth--
	if err != nil {
		d.popEOF()
		return err
	}

	if !d.popEOF() {
		return fmt.Errorf("xml: %s.UnmarshalXML did not consume entire <%s> element", receiverType(val), start.Name.Local)
	}

	return nil
}

// unmarshalTextInterface unmarshals a single XML element into val.
// The chardata contained in the element (but not its children)
// is passed to the text unmarshaler.
func (d *Decoder) unmarshalTextInterface(val encoding.TextUnmarshaler) error {
	var buf []byte
	depth := 1
	for depth > 0 {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch t := t.(type) {
		case CharData:
			if depth == 1 {
				buf = append(buf, t...)
			}
		case StartElement:
			depth++
		case EndElement:
			depth--
		}
	}
	return val.UnmarshalText(buf)
}

// unmarshalAttr unmarshals a single XML attribute into val.
func (d *Decoder) unmarshalAttr(val reflect.Value, attr Attr) error {
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}
	if val.CanInterface() && val.Type().Implements(unmarshalerAttrType) {
		// This is an unmarshaler with a non-pointer receiver,
		// so it's likely to be incorrect, but we do what we're told.
		return val.Interface().(UnmarshalerAttr).UnmarshalXMLAttr(attr)
	}
	if val.CanAddr() {
		pv := val.Addr()
		if pv.CanInterface() && pv.Type().Implements(unmarshalerAttrType) {
			return pv.Interface().(UnmarshalerAttr).UnmarshalXMLAttr(attr)
		}
	}

	// Not an UnmarshalerAttr; try encoding.TextUnmarshaler.
	if val.CanInterface() && val.Type().Implements(textUnmarshalerType) {
		// This is an unmarshaler with a non-pointer receiver,
		// so it's likely to be incorrect, but we do what we're told.
		return val.Interface().(encoding.TextUnmarshaler).UnmarshalText([]byte(attr.Value))
	}
	if val.CanAddr() {
		pv := val.Addr()
		if pv.CanInterface() && pv.Type().Implements(textUnmarshalerType) {
			return pv.Interface().(encoding.TextUnmarshaler).UnmarshalText([]byte(attr.Value))
		}
	}

	if val.Type().Kind() == reflect.Slice && val.Type().Elem().Kind() != reflect.Uint8 {
		// Slice of element values.
		// Grow slice.
		n := val.Len()
		val.Set(reflect.Append(val, reflect.Zero(val.Type().Elem())))

		// Recur to read element into slice.
		if err := d.unmarshalAttr(val.Index(n), attr); err != nil {
			val.SetLen(n)
			return err
		}
		return nil
	}

	if val.Type() == attrType {
		val.Set(reflect.ValueOf(attr))
		return nil
	}

	return copyValue(val, []byte(attr.Value))
}

var (
	attrType            = reflect.TypeOf(Attr{})
	unmarshalerType     = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
	unmarshalerAttrType = reflect.TypeOf((*UnmarshalerAttr)(nil)).Elem()
	textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

// Unmarshal a single XML element into val.
func (d *Decoder) unmarshal(val reflect.Value, start *StartElement) error {
	// Find start element if we need it.
	if start == nil {
		for {
			tok, err := d.Token()
			if err != nil {
				return err
			}
			if t, ok := tok.(StartElement); ok {
				start = &t
				break
			}
		}
	}

	// Load value from interface, but only if the result will be
	// usefully addressable.
	if val.Kind() == reflect.Interface && !val.IsNil() {
		e := val.Elem()
		if e.Kind() == reflect.Ptr && !e.IsNil() {
			val = e
		}
	}

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		val = val.Elem()
	}

	if val.CanInterface() && val.Type().Implements(unmarshalerType) {
		// This is an unmarshaler with a non-pointer receiver,
		// so it's likely to be incorrect, but we do what we're told.
		return d.unmarshalInterface(val.Interface().(Unmarshaler), start)
	}

	if val.CanAddr() {
		pv := val.Addr()
		if pv.CanInterface() && pv.Type().Implements(unmarshalerType) {
			return d.unmarshalInterface(pv.Interface().(Unmarshaler), start)
		}
	}

	if val.CanInterface() && val.Type().Implements(textUnmarshalerType) {
		return d.unmarshalTextInterface(val.Interface().(encoding.TextUnmarshaler))
	}

	if val.CanAddr() {
		pv := val.Addr()
		if pv.CanInterface() && pv.Type().Implements(textUnmarshalerType) {
			return d.unmarshalTextInterface(pv.Interface().(encoding.TextUnmarshaler))
		}
	}

	var (
		data         []byte
		saveData     reflect.Value
		comment      []byte
		saveComment  reflect.Value
		saveXML      reflect.Value
		saveXMLIndex int
		saveXMLData  []byte
		saveAny      reflect.Value
		sv           reflect.Value
		tinfo        *typeInfo
		err          error
	)

	switch v := val; v.Kind() {
	default:
		return errors.New("unknown type " + v.Type().String())

	case reflect.Interface:
		// TODO: For now, simply ignore the field. In the near
		//       future we may choose to unmarshal the start
		//       element on it, if not nil.
		return d.Skip()

	case reflect.Slice:
		typ := v.Type()
		if typ.Elem().Kind() == reflect.Uint8 {
			// []byte
			saveData = v
			break
		}

		// Slice of element values.
		// Grow slice.
		n := v.Len()
		v.Set(reflect.Append(val, reflect.Zero(v.Type().Elem())))

		// Recur to read element into slice.
		if err := d.unmarshal(v.Index(n), start); err != nil {
			v.SetLen(n)
			return err
		}
		return nil

	case reflect.Bool, reflect.Float32, reflect.Float64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.String:
		saveData = v

	case reflect.Struct:
		typ := v.Type()
		if typ == nameType {
			v.Set(reflect.ValueOf(start.Name))
			break
		}

		sv = v
		tinfo, err = getTypeInfo(typ)
		if err != nil {
			return err
		}

		// Validate and assign element name.
		if tinfo.xmlname != nil {
			finfo := tinfo.xmlname
			if finfo.name != "" && finfo.name != start.Name.Local {
				return UnmarshalError("expected element type <" + finfo.name + "> but have <" + start.Name.Local + ">")
			}
			if finfo.xmlns != "" && finfo.xmlns != start.Name.Space {
				e := "expected element <" + finfo.name + "> in name space " + finfo.xmlns + " but have "
				if start.Name.Space == "" {
					e += "no name space"
				} else {
					e += start.Name.Space
				}
				return UnmarshalError(e)
			}
			fv := finfo.value(sv)
			if _, ok := fv.Interface().(Name); ok {
				fv.Set(reflect.ValueOf(start.Name))
			}
		}

		// Assign attributes.
		for _, a := range start.Attr {
			handled := false
			any := -1
			for i := range tinfo.fields {
				finfo := &tinfo.fields[i]
				switch finfo.flags & fMode {
				case fAttr:
					strv := finfo.value(sv)
					if a.Name.Local == finfo.name /*&& (finfo.xmlns == "" || finfo.xmlns == a.Name.Space)*/ {
						if err := d.unmarshalAttr(strv, a); err != nil {
							return err
						}
						handled = true
					}

				case fAny | fAttr:
					if any == -1 {
						any = i
					}
				}
			}
			if !handled && any >= 0 {
				finfo := &tinfo.fields[any]
				strv := finfo.value(sv)
				if err := d.unmarshalAttr(strv, a); err != nil {
					return err
				}
			}
		}

		// Determine whether we need to save character data or comments.
		for i := range tinfo.fields {
			finfo := &tinfo.fields[i]
			switch finfo.flags & fMode {
			case fCDATA, fCharData:
				if !saveData.IsValid() {
					saveData = finfo.value(sv)
				}

			case fComment:
				if !saveComment.IsValid() {
					saveComment = finfo.value(sv)
				}

			case fAny, fAny | fElement:
				if !saveAny.IsValid() {
					saveAny = finfo.value(sv)
				}

			case fInnerXml:
				if !saveXML.IsValid() {
					saveXML = finfo.value(sv)
					if d.saved == nil {
						saveXMLIndex = 0
						d.saved = new(bytes.Buffer)
					} else {
						saveXMLIndex = d.savedOffset()
					}
				}
			}
		}
	}

	// Find end element.
	// Process sub-elements along the way.
Loop:
	for {
		var savedOffset int
		if saveXML.IsValid() {
			savedOffset = d.savedOffset()
		}
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case StartElement:
			consumed := false
			if sv.IsValid() {
				consumed, err = d.unmarshalPath(tinfo, sv, nil, &t)
				if err != nil {
					return err
				}
				if !consumed && saveAny.IsValid() {
					consumed = true
					if err := d.unmarshal(saveAny, &t); err != nil {
						return err
					}
				}
			}
			if !consumed {
				if err := d.Skip(); err != nil {
					return err
				}
			}

		case EndElement:
			if saveXML.IsValid() {
				saveXMLData = d.saved.Bytes()[saveXMLIndex:savedOffset]
				if saveXMLIndex == 0 {
					d.saved = nil
				}
			}
			break Loop

		case CharData:
			if saveData.IsValid() {
				data = append(data, t...)
			}

		case Comment:
			if saveComment.IsValid() {
				comment = append(comment, t...)
			}
		}
	}

	if saveData.IsValid() && saveData.CanInterface() && saveData.Type().Implements(textUnmarshalerType) {
		if err := saveData.Interface().(encoding.TextUnmarshaler).UnmarshalText(data); err != nil {
			return err
		}
		saveData = reflect.Value{}
	}

	if saveData.IsValid() && saveData.CanAddr() {
		pv := saveData.Addr()
		if pv.CanInterface() && pv.Type().Implements(textUnmarshalerType) {
			if err := pv.Interface().(encoding.TextUnmarshaler).UnmarshalText(data); err != nil {
				return err
			}
			saveData = reflect.Value{}
		}
	}

	if err := copyValue(saveData, data); err != nil {
		return err
	}

	switch t := saveComment; t.Kind() {
	case reflect.String:
		t.SetString(string(comment))
	case reflect.Slice:
		t.Set(reflect.ValueOf(comment))
	}

	switch t := saveXML; t.Kind() {
	case reflect.String:
		t.SetString(string(saveXMLData))
	case reflect.Slice:
		if t.Type().Elem().Kind() == reflect.Uint8 {
			t.Set(reflect.ValueOf(saveXMLData))
		}
	}

	return nil
}

func copyValue(dst reflect.Value, src []byte) (err error) {
	dst0 := dst

	if dst.Kind() == reflect.Ptr {
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		dst = dst.Elem()
	}

	// Save accumulated data.
	switch dst.Kind() {
	case reflect.Invalid:
		// Probably a comment.
	default:
		return errors.New("cannot unmarshal into " + dst0.Type().String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(src) == 0 {
			dst.SetInt(0)
			return nil
		}
		itmp, err := strconv.ParseInt(strings.TrimSpace(string(src)), 10, dst.Type().Bits())
		if err != nil {
			return err
		}
		dst.SetInt(itmp)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if len(src) == 0 {
			dst.SetUint(0)
			return nil
		}
		utmp, err := strconv.ParseUint(strings.TrimSpace(string(src)), 10, dst.Type().Bits())
		if err != nil {
			return err
		}
		dst.SetUint(utmp)
	case reflect.Float32, reflect.Float64:
		if len(src) == 0 {
			dst.SetFloat(0)
			return nil
		}
		ftmp, err := strconv.ParseFloat(strings.TrimSpace(string(src)), dst.Type().Bits())
		if err != nil {
			return err
		}
		dst.SetFloat(ftmp)
	case reflect.Bool:
		if len(src) == 0 {
			dst.SetBool(false)
			return nil
		}
		value, err := strconv.ParseBool(strings.TrimSpace(string(src)))
		if err != nil {
			return err
		}
		dst.SetBool(value)
	case reflect.String:
		dst.SetString(string(src))
	case reflect.Slice:
		if len(src) == 0 {
			// non-nil to flag presence
			src = []byte{}
		}
		dst.SetBytes(src)
	}
	return nil
}

// unmarshalPath walks down an XML structure looking for wanted
// paths, and calls unmarshal on them.
// The consumed result tells whether XML elements have been consumed
// from the Decoder until start's matching end element, or if it's
// still untouched because start is uninteresting for sv's fields.
func (d *Decoder) unmarshalPath(tinfo *typeInfo, sv reflect.Value, parents []string, start *StartElement) (consumed bool, err error) {
	recurse := false
Loop:
	for i := range tinfo.fields {
		finfo := &tinfo.fields[i]
		if finfo.flags&fElement == 0 || len(finfo.parents) < len(parents) /*|| finfo.xmlns != "" && finfo.xmlns != start.Name.Space*/ {
			continue
		}
		for j := range parents {
			if parents[j] != finfo.parents[j] {
				continue Loop
			}
		}
		if len(finfo.parents) == len(parents) && finfo.name == start.Name.Local {
			// It's a perfect match, unmarshal the field.
			return true, d.unmarshal(finfo.value(sv), start)
		}
		if len(finfo.parents) > len(parents) && finfo.parents[len(parents)] == start.Name.Local {
			// It's a prefix for the field. Break and recurse
			// since it's not ok for one field path to be itself
			// the prefix for another field path.
			recurse = true

			// We can reuse the same slice as long as we
			// don't try to append to it.
			parents = finfo.parents[:len(parents)+1]
			break
		}
	}
	if !recurse {
		// We have no business with this element.
		return false, nil
	}
	// The element is not a perfect match for any field, but one
	// or more fields have the path to this element as a parent
	// prefix. Recurse and attempt to match these.
	for {
		var tok Token
		tok, err = d.Token()
		if err != nil {
			return true, err
		}
		switch t := tok.(type) {
		case StartElement:
			consumed2, err := d.unmarshalPath(tinfo, sv, parents, &t)
			if err != nil {
				return true, err
			}
			if !consumed2 {
				if err := d.Skip(); err != nil {
					return true, err
				}
			}
		case EndElement:
			return true, nil
		}
	}
}

// Skip reads tokens until it has consumed the end element
// matching the most recent start element already consumed.
// It recurs if it encounters a start element, so it can be used to
// skip nested structures.
// It returns nil if it finds an end element matching the start
// element; otherwise it returns an error describing the problem.
func (d *Decoder) Skip() error {
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch tok.(type) {
		case StartElement:
			if err := d.Skip(); err != nil {
				return err
			}
		case EndElement:
			return nil
		}
	}
}

// typeinfo.go
// typeInfo holds details for the xml representation of a type.
type typeInfo struct {
	xmlname *fieldInfo
	fields  []fieldInfo
}

// fieldInfo holds details for the xml representation of a single field.
type fieldInfo struct {
	idx     []int
	name    string
	xmlns   string
	flags   fieldFlags
	parents []string
}

type fieldFlags int

const (
	fElement fieldFlags = 1 << iota
	fAttr
	fCDATA
	fCharData
	fInnerXml
	fComment
	fAny

	fOmitEmpty

	fMode = fElement | fAttr | fCDATA | fCharData | fInnerXml | fComment | fAny

	xmlName = "XMLName"
)

var tinfoMap sync.Map // map[reflect.Type]*typeInfo

var nameType = reflect.TypeOf(Name{})

// getTypeInfo returns the typeInfo structure with details necessary
// for marshaling and unmarshaling typ.
func getTypeInfo(typ reflect.Type) (*typeInfo, error) {
	if ti, ok := tinfoMap.Load(typ); ok {
		return ti.(*typeInfo), nil
	}

	tinfo := &typeInfo{}
	if typ.Kind() == reflect.Struct && typ != nameType {
		n := typ.NumField()
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if (f.PkgPath != "" && !f.Anonymous) || f.Tag.Get("xml") == "-" {
				continue // Private field
			}

			// For embedded structs, embed its fields.
			if f.Anonymous {
				t := f.Type
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
				}
				if t.Kind() == reflect.Struct {
					inner, err := getTypeInfo(t)
					if err != nil {
						return nil, err
					}
					if tinfo.xmlname == nil {
						tinfo.xmlname = inner.xmlname
					}
					for _, finfo := range inner.fields {
						finfo.idx = append([]int{i}, finfo.idx...)
						if err := addFieldInfo(typ, tinfo, &finfo); err != nil {
							return nil, err
						}
					}
					continue
				}
			}

			finfo, err := structFieldInfo(typ, &f)
			if err != nil {
				return nil, err
			}

			if f.Name == xmlName {
				tinfo.xmlname = finfo
				continue
			}

			// Add the field if it doesn't conflict with other fields.
			if err := addFieldInfo(typ, tinfo, finfo); err != nil {
				return nil, err
			}
		}
	}

	ti, _ := tinfoMap.LoadOrStore(typ, tinfo)
	return ti.(*typeInfo), nil
}

// structFieldInfo builds and returns a fieldInfo for f.
func structFieldInfo(typ reflect.Type, f *reflect.StructField) (*fieldInfo, error) {
	finfo := &fieldInfo{idx: f.Index}

	// Split the tag from the xml namespace if necessary.
	tag := f.Tag.Get("xml")
	if i := strings.Index(tag, " "); i >= 0 {
		finfo.xmlns, tag = tag[:i], tag[i+1:]
	}

	// Parse flags.
	tokens := strings.Split(tag, ",")
	if len(tokens) == 1 {
		finfo.flags = fElement
	} else {
		tag = tokens[0]
		for _, flag := range tokens[1:] {
			switch flag {
			case "attr":
				finfo.flags |= fAttr
			case "cdata":
				finfo.flags |= fCDATA
			case "chardata":
				finfo.flags |= fCharData
			case "innerxml":
				finfo.flags |= fInnerXml
			case "comment":
				finfo.flags |= fComment
			case "any":
				finfo.flags |= fAny
			case "omitempty":
				finfo.flags |= fOmitEmpty
			}
		}

		// Validate the flags used.
		valid := true
		switch mode := finfo.flags & fMode; mode {
		case 0:
			finfo.flags |= fElement
		case fAttr, fCDATA, fCharData, fInnerXml, fComment, fAny, fAny | fAttr:
			if f.Name == xmlName || tag != "" && mode != fAttr {
				valid = false
			}
		default:
			// This will also catch multiple modes in a single field.
			valid = false
		}
		if finfo.flags&fMode == fAny {
			finfo.flags |= fElement
		}
		if finfo.flags&fOmitEmpty != 0 && finfo.flags&(fElement|fAttr) == 0 {
			valid = false
		}
		if !valid {
			return nil, fmt.Errorf("xml: invalid tag in field %s of type %s: %q",
				f.Name, typ, f.Tag.Get("xml"))
		}
	}

	// Use of xmlns without a name is not allowed.
	if finfo.xmlns != "" && tag == "" {
		return nil, fmt.Errorf("xml: namespace without name in field %s of type %s: %q",
			f.Name, typ, f.Tag.Get("xml"))
	}

	if f.Name == xmlName {
		// The XMLName field records the XML element name. Don't
		// process it as usual because its name should default to
		// empty rather than to the field name.
		finfo.name = tag
		return finfo, nil
	}

	if tag == "" {
		// If the name part of the tag is completely empty, get
		// default from XMLName of underlying struct if feasible,
		// or field name otherwise.
		if xmlname := lookupXMLName(f.Type); xmlname != nil {
			finfo.xmlns, finfo.name = xmlname.xmlns, xmlname.name
		} else {
			finfo.name = f.Name
		}
		return finfo, nil
	}

	// Prepare field name and parents.
	parents := strings.Split(tag, ">")
	if parents[0] == "" {
		parents[0] = f.Name
	}
	if parents[len(parents)-1] == "" {
		return nil, fmt.Errorf("xml: trailing '>' in field %s of type %s", f.Name, typ)
	}
	finfo.name = parents[len(parents)-1]
	if len(parents) > 1 {
		if (finfo.flags & fElement) == 0 {
			return nil, fmt.Errorf("xml: %s chain not valid with %s flag", tag, strings.Join(tokens[1:], ","))
		}
		finfo.parents = parents[:len(parents)-1]
	}

	// If the field type has an XMLName field, the names must match
	// so that the behavior of both marshaling and unmarshaling
	// is straightforward and unambiguous.
	if finfo.flags&fElement != 0 {
		ftyp := f.Type
		xmlname := lookupXMLName(ftyp)
		if xmlname != nil && xmlname.name != finfo.name {
			return nil, fmt.Errorf("xml: name %q in tag of %s.%s conflicts with name %q in %s.XMLName",
				finfo.name, typ, f.Name, xmlname.name, ftyp)
		}
	}
	return finfo, nil
}

// lookupXMLName returns the fieldInfo for typ's XMLName field
// in case it exists and has a valid xml field tag, otherwise
// it returns nil.
func lookupXMLName(typ reflect.Type) (xmlname *fieldInfo) {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	for i, n := 0, typ.NumField(); i < n; i++ {
		f := typ.Field(i)
		if f.Name != xmlName {
			continue
		}
		finfo, err := structFieldInfo(typ, &f)
		if err == nil && finfo.name != "" {
			return finfo
		}
		// Also consider errors as a non-existent field tag
		// and let getTypeInfo itself report the error.
		break
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// addFieldInfo adds finfo to tinfo.fields if there are no
// conflicts, or if conflicts arise from previous fields that were
// obtained from deeper embedded structures than finfo. In the latter
// case, the conflicting entries are dropped.
// A conflict occurs when the path (parent + name) to a field is
// itself a prefix of another path, or when two paths match exactly.
// It is okay for field paths to share a common, shorter prefix.
func addFieldInfo(typ reflect.Type, tinfo *typeInfo, newf *fieldInfo) error {
	var conflicts []int
Loop:
	// First, figure all conflicts. Most working code will have none.
	for i := range tinfo.fields {
		oldf := &tinfo.fields[i]
		if oldf.flags&fMode != newf.flags&fMode {
			continue
		}
		if oldf.xmlns != "" && newf.xmlns != "" && oldf.xmlns != newf.xmlns {
			continue
		}
		minl := min(len(newf.parents), len(oldf.parents))
		for p := 0; p < minl; p++ {
			if oldf.parents[p] != newf.parents[p] {
				continue Loop
			}
		}
		if len(oldf.parents) > len(newf.parents) {
			if oldf.parents[len(newf.parents)] == newf.name {
				conflicts = append(conflicts, i)
			}
		} else if len(oldf.parents) < len(newf.parents) {
			if newf.parents[len(oldf.parents)] == oldf.name {
				conflicts = append(conflicts, i)
			}
		} else {
			if newf.name == oldf.name {
				conflicts = append(conflicts, i)
			}
		}
	}
	// Without conflicts, add the new field and return.
	if conflicts == nil {
		tinfo.fields = append(tinfo.fields, *newf)
		return nil
	}

	// If any conflict is shallower, ignore the new field.
	// This matches the Go field resolution on embedding.
	for _, i := range conflicts {
		if len(tinfo.fields[i].idx) < len(newf.idx) {
			return nil
		}
	}

	// Otherwise, if any of them is at the same depth level, it's an error.
	for _, i := range conflicts {
		oldf := &tinfo.fields[i]
		if len(oldf.idx) == len(newf.idx) {
			f1 := typ.FieldByIndex(oldf.idx)
			f2 := typ.FieldByIndex(newf.idx)
			return &TagPathError{typ, f1.Name, f1.Tag.Get("xml"), f2.Name, f2.Tag.Get("xml")}
		}
	}

	// Otherwise, the new field is shallower, and thus takes precedence,
	// so drop the conflicting fields from tinfo and append the new one.
	for c := len(conflicts) - 1; c >= 0; c-- {
		i := conflicts[c]
		copy(tinfo.fields[i:], tinfo.fields[i+1:])
		tinfo.fields = tinfo.fields[:len(tinfo.fields)-1]
	}
	tinfo.fields = append(tinfo.fields, *newf)
	return nil
}

// A TagPathError represents an error in the unmarshaling process
// caused by the use of field tags with conflicting paths.
type TagPathError struct {
	Struct       reflect.Type
	Field1, Tag1 string
	Field2, Tag2 string
}

func (e *TagPathError) Error() string {
	return fmt.Sprintf("%s field %q with tag %q conflicts with field %q with tag %q", e.Struct, e.Field1, e.Tag1, e.Field2, e.Tag2)
}

// value returns v's field value corresponding to finfo.
// It's equivalent to v.FieldByIndex(finfo.idx), but initializes
// and dereferences pointers as necessary.
func (finfo *fieldInfo) value(v reflect.Value) reflect.Value {
	for i, x := range finfo.idx {
		if i > 0 {
			t := v.Type()
			if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
				if v.IsNil() {
					v.Set(reflect.New(v.Type().Elem()))
				}
				v = v.Elem()
			}
		}
		v = v.Field(x)
	}
	return v
}
