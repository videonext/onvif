package soap

import (
	"bytes"
	"context"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type SOAPEncoder interface {
	Encode(v interface{}) error
	Flush() error
}

type SOAPDecoder interface {
	Decode(v interface{}) error
}

type SOAPHeader struct {
	XMLName xml.Name      `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
	Headers []interface{} `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  SOAPHeader
	Body    SOAPBody
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// UnmarshalXML unmarshals SOAPBody xml
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://www.w3.org/2003/05/soap-envelope" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

type SOAPFaultSubCode struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Subcode"`

	Value string
}

type SOAPFaultCode struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Code"`

	Value   string
	Subcode SOAPFaultSubCode
}

type SOAPFaultReason struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Reason"`

	Text string
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`

	//	Code   string `xml:"faultcode,omitempty"`
	Code   SOAPFaultCode
	Reason SOAPFaultReason
	//	String string `xml:"faultstring,omitempty"`
	//	Actor  string `xml:"faultactor,omitempty"`
	//	Detail string `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	return f.Reason.Text
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE       string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU        string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType       string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
	mtomContentType string = `multipart/related; start-info="application/soap+xml"; type="application/xop+xml"; boundary="%s"`
)

type WSSPassword struct {
	Type  string `xml:",attr"`
	Value string `xml:",chardata"`
}

type WSSNonce struct {
	EncodingType string `xml:",attr"`
	Value        string `xml:",chardata"`
}

type WSSCreated struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Created"`
	Value   string   `xml:",chardata"`
}

type WSSUsernameToken struct {
	//	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope UsernameToken"`

	Username string

	Password WSSPassword

	Nonce WSSNonce

	Created WSSCreated
}

type WSSSecurityHeader struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`

	MustUnderstand string `xml:"mustUnderstand,attr"`

	UsernameToken WSSUsernameToken
}

// NewWSSSecurityHeader creates WSSSecurityHeader instance
func NewWSSSecurityHeader(user, pass string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{MustUnderstand: "1"}

	// Username
	hdr.UsernameToken.Username = user

	// Created
	hdr.UsernameToken.Created.Value = time.Now().Format("2006-01-02T15:04:05.999") + "Z"

	// Nonce
	b := make([]byte, 16)
	rand.Read(b)
	nonce := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	hdr.UsernameToken.Nonce.EncodingType = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
	hdr.UsernameToken.Nonce.Value = base64.StdEncoding.EncodeToString([]byte(nonce))

	// Password
	h := sha1.New()
	h.Write([]byte(nonce + hdr.UsernameToken.Created.Value + pass))
	hdr.UsernameToken.Password.Type = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest"
	hdr.UsernameToken.Password.Value = base64.StdEncoding.EncodeToString(h.Sum(nil))

	return hdr
}

type basicAuth struct {
	Login    string
	Password string
}

type options struct {
	tlsCfg           *tls.Config
	auth             *basicAuth
	timeout          time.Duration
	contimeout       time.Duration
	tlshshaketimeout time.Duration
	client           HTTPClient
	httpHeaders      map[string]string
	mtom             bool
}

var defaultOptions = options{
	timeout:          time.Duration(30 * time.Second),
	contimeout:       time.Duration(90 * time.Second),
	tlshshaketimeout: time.Duration(15 * time.Second),
}

// A Option sets options such as credentials, tls, etc.
type Option func(*options)

// WithHTTPClient is an Option to set the HTTP client to use
// This cannot be used with WithTLSHandshakeTimeout, WithTLS,
// WithTimeout options
func WithHTTPClient(c HTTPClient) Option {
	return func(o *options) {
		o.client = c
	}
}

// WithTLSHandshakeTimeout is an Option to set default tls handshake timeout
// This option cannot be used with WithHTTPClient
func WithTLSHandshakeTimeout(t time.Duration) Option {
	return func(o *options) {
		o.tlshshaketimeout = t
	}
}

// WithRequestTimeout is an Option to set default end-end connection timeout
// This option cannot be used with WithHTTPClient
func WithRequestTimeout(t time.Duration) Option {
	return func(o *options) {
		o.contimeout = t
	}
}

// WithBasicAuth is an Option to set BasicAuth
func WithBasicAuth(login, password string) Option {
	return func(o *options) {
		o.auth = &basicAuth{Login: login, Password: password}
	}
}

// WithTLS is an Option to set tls config
// This option cannot be used with WithHTTPClient
func WithTLS(tls *tls.Config) Option {
	return func(o *options) {
		o.tlsCfg = tls
	}
}

// WithTimeout is an Option to set default HTTP dial timeout
func WithTimeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// WithHTTPHeaders is an Option to set global HTTP headers for all requests
func WithHTTPHeaders(headers map[string]string) Option {
	return func(o *options) {
		o.httpHeaders = headers
	}
}

// WithMTOM is an Option to set Message Transmission Optimization Mechanism
// MTOM encodes fields of type Binary using XOP.
func WithMTOM() Option {
	return func(o *options) {
		o.mtom = true
	}
}

// Client is soap client
type Client struct {
	url     string
	opts    *options
	headers []interface{}
}

// HTTPClient is a client which can make HTTP requests
// An example implementation is net/http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient creates new SOAP client instance
func NewClient(url string, opt ...Option) *Client {
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	return &Client{
		url:  url,
		opts: &opts,
	}
}

// AddHeader adds envelope header
func (s *Client) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

// CallContext performs HTTP POST request with a context
func (s *Client) CallContext(ctx context.Context, soapAction string, request, response interface{}) error {
	return s.call(ctx, soapAction, request, response)
}

// Call performs HTTP POST request
func (s *Client) Call(soapAction string, request, response interface{}) error {
	return s.call(context.Background(), soapAction, request, response)
}

func (s *Client) call(ctx context.Context, soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		envelope.Header.Headers = s.headers
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)
	var encoder SOAPEncoder
	if s.opts.mtom {
		encoder = newMtomEncoder(buffer)
	} else {
		encoder = xml.NewEncoder(buffer)
	}

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	fmt.Printf("POST: %s\n", s.url)
	if err != nil {
		return err
	}
	if s.opts.auth != nil {
		req.SetBasicAuth(s.opts.auth.Login, s.opts.auth.Password)
	}

	req.WithContext(ctx)

	if s.opts.mtom {
		req.Header.Add("Content-Type", fmt.Sprintf(mtomContentType, encoder.(*mtomEncoder).Boundary()))
	} else {
		req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	}
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Set("User-Agent", "gowsdl/0.1")
	if s.opts.httpHeaders != nil {
		for k, v := range s.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	client := s.opts.client
	if client == nil {
		tr := &http.Transport{
			TLSClientConfig: s.opts.tlsCfg,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := net.Dialer{Timeout: s.opts.timeout}
				return d.DialContext(ctx, network, addr)
			},
			TLSHandshakeTimeout: s.opts.tlshshaketimeout,
		}
		client = &http.Client{Timeout: s.opts.contimeout, Transport: tr}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}

	mtomBoundary, err := getMtomHeader(res.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	var dec SOAPDecoder
	if mtomBoundary != "" {
		dec = newMtomDecoder(res.Body, mtomBoundary)
	} else {
		dec = xml.NewDecoder(res.Body)
	}

	if err := dec.Decode(respEnvelope); err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
