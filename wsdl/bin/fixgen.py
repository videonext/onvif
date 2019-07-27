#!/usr/bin/env python3
import sys
import os
import re

# How to use:
# go get github.com/hooklift/gowsdl/...
# ./fixgen.py [onvif profile file name without .wsdl]

if len(sys.argv) != 2:
    print("Usage:\n./fixgen.py [profile name without .wsdl]")
    exit(1)

go_package = os.path.basename(sys.argv[1])
go_src = os.path.basename(sys.argv[1]) + '.go'
wsdl_file = sys.argv[1] + '.wsdl'

with open(wsdl_file, 'r') as file:
    wsdl = file.read()

os.system('gowsdl -o ' + go_src + ' -p ' + go_package + ' ' + wsdl_file + ' | grep -v expected')

with open(go_package + '/' + go_src, 'r') as file:
    data = file.read()

data = data.replace('"github.com/hooklift/gowsdl/soap"',
                    '"github.com/videonext/onvif/soap"')

# working around some bugs in the gowsdl
data = data.replace('interface{}', '')
data = data.replace('TLS1.0 bool', 'TLS1_0 bool')
data = data.replace('TLS1.1 bool', 'TLS1_1 bool')
data = data.replace('TLS1.2 bool', 'TLS1_2 bool')
data = data.replace('X.509Token bool', 'X_509Token bool')
# remove dupcicated types
data = re.sub(r'type \b(\w+)\s+\1\b', '', data)
data = data.replace('type IntList IntAttrList', '')
data = data.replace('type FloatList FloatAttrList', '')
data = data.replace('type Capabilities DeviceServiceCapabilities', '')
data = data.replace('type FaultcodeEnum *QName', 'type FaultcodeEnum QName')
data = data.replace('type FaultCodesType *QName', 'type FaultCodesType QName')
data = data.replace('type RelationshipType *AnyURI', 'type RelationshipType AnyURI')


# get wsdl namespace
r = re.findall(r'targetNamespace="(http://www.onvif.org/.+)"', wsdl)
if len(r):
    targetNamespace = r[0]
#print("Namespace: " + targetNamespace)

# add wsdl namespace and method name to the soap action
data = re.sub(r'(?s)func \(service (\*\w+\)) (\w+)Context\s*\(ctx context.Context, request (\*\w+\)) \((\*\w+), error\)(.+?)service\.client\.CallContext\(ctx, "(\'\')",',
              r'func (service \1 \2Context(ctx context.Context, request \3 (\4, error)\5service.client.CallContext(ctx, "' + targetNamespace + '/' + r'\2",', data)

# patch struct
data = re.sub(r'(?s)type\s+(\w+?)\s+struct\s+\{\s+?client\s+\*soap\.Client',
              r'type \1 struct {\nclient *soap.Client\nxaddr  string\n', data)


# and initialization
data = re.sub(r'(?s)func New(.+?)\(client \*soap.Client\) (.+?) \{.+?return \&(\w+)\{.+?}',
              r'func New\1(client *soap.Client, xaddr string) \2 {\n return &\3{\nclient: client,\nxaddr: xaddr,\n}', data)


# patch call service.client.CallContext
data = data.replace('service.client.CallContext(ctx,', 'service.client.CallContext(ctx, service.xaddr,')

# add some types
data += "\ntype AnyURI string\n"
data += "type Duration string\n"
data += "type QName string\n"
data += "type NCName string\n"
data += "type NonNegativeInteger int64\n"
data += "type PositiveInteger int64\n"
data += "type NonPositiveInteger int64\n"
data += "type AnySimpleType string\n"


with open(go_package + '/' + go_src, 'w') as file:
    file.write(data)

os.system("gofmt -w " + go_package + '/' + go_src)
