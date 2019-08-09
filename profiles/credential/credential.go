package credential

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capability response message contains the requested credential
	// service capabilities using a hierarchical XML capability structure.
	//
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty"`
}

// GetSupportedFormatTypes type
type GetSupportedFormatTypes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetSupportedFormatTypes"`

	// Name of the credential identifier type
	CredentialIdentifierTypeName string `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialIdentifierTypeName,omitempty"`
}

// GetSupportedFormatTypesResponse type
type GetSupportedFormatTypesResponse struct {
	XMLName xml.Name `xml:"GetSupportedFormatTypesResponse"`

	// Identifier format type
	FormatTypeInfo []CredentialIdentifierFormatTypeInfo `xml:"FormatTypeInfo,omitempty"`
}

// GetCredentialInfo type
type GetCredentialInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentialInfo"`

	// Tokens of CredentialInfo items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl Token,omitempty"`
}

// GetCredentialInfoResponse type
type GetCredentialInfoResponse struct {
	XMLName xml.Name `xml:"GetCredentialInfoResponse"`

	// List of CredentialInfo items.
	CredentialInfo []CredentialInfo `xml:"CredentialInfo,omitempty"`
}

// GetCredentialInfoList type
type GetCredentialInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentialInfoList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/credential/wsdl StartReference,omitempty"`
}

// GetCredentialInfoListResponse type
type GetCredentialInfoListResponse struct {
	XMLName xml.Name `xml:"GetCredentialInfoListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of CredentialInfo items.
	CredentialInfo []CredentialInfo `xml:"CredentialInfo,omitempty"`
}

// GetCredentials type
type GetCredentials struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentials"`

	// Token of Credentials to get
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl Token,omitempty"`
}

// GetCredentialsResponse type
type GetCredentialsResponse struct {
	XMLName xml.Name `xml:"GetCredentialsResponse"`

	// List of Credential items.
	Credential []Credential `xml:"Credential,omitempty"`
}

// GetCredentialList type
type GetCredentialList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentialList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/credential/wsdl StartReference,omitempty"`
}

// GetCredentialListResponse type
type GetCredentialListResponse struct {
	XMLName xml.Name `xml:"GetCredentialListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of Credential items.
	Credential []Credential `xml:"Credential,omitempty"`
}

// CreateCredential type
type CreateCredential struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl CreateCredential"`

	// The credential to create.
	Credential Credential `xml:"http://www.onvif.org/ver10/credential/wsdl Credential,omitempty"`

	// The state of the credential.
	State CredentialState `xml:"http://www.onvif.org/ver10/credential/wsdl State,omitempty"`
}

// CreateCredentialResponse type
type CreateCredentialResponse struct {
	XMLName xml.Name `xml:"CreateCredentialResponse"`

	// The token of the created credential
	Token ReferenceToken `xml:"Token,omitempty"`
}

// ModifyCredential type
type ModifyCredential struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl ModifyCredential"`

	// Details of the credential.
	Credential Credential `xml:"http://www.onvif.org/ver10/credential/wsdl Credential,omitempty"`
}

// ModifyCredentialResponse type
type ModifyCredentialResponse struct {
	XMLName xml.Name `xml:"ModifyCredentialResponse"`
}

// SetCredential type
type SetCredential struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl SetCredential"`

	// Details of the credential.
	CredentialData CredentialData `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialData,omitempty"`
}

// SetCredentialResponse type
type SetCredentialResponse struct {
	XMLName xml.Name `xml:"SetCredentialResponse"`
}

// DeleteCredential type
type DeleteCredential struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl DeleteCredential"`

	// The token of the credential to delete.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl Token,omitempty"`
}

// DeleteCredentialResponse type
type DeleteCredentialResponse struct {
	XMLName xml.Name `xml:"DeleteCredentialResponse"`
}

// GetCredentialState type
type GetCredentialState struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentialState"`

	// Token of Credential
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl Token,omitempty"`
}

// GetCredentialStateResponse type
type GetCredentialStateResponse struct {
	XMLName xml.Name `xml:"GetCredentialStateResponse"`

	// State of the credential.
	State CredentialState `xml:"State,omitempty"`
}

// EnableCredential type
type EnableCredential struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl EnableCredential"`

	// The token of the credential
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl Token,omitempty"`

	// Reason for enabling the credential.
	Reason Name `xml:"http://www.onvif.org/ver10/credential/wsdl Reason,omitempty"`
}

// EnableCredentialResponse type
type EnableCredentialResponse struct {
	XMLName xml.Name `xml:"EnableCredentialResponse"`
}

// DisableCredential type
type DisableCredential struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl DisableCredential"`

	// Token of the Credential
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl Token,omitempty"`

	// Reason for disabling the credential
	Reason Name `xml:"http://www.onvif.org/ver10/credential/wsdl Reason,omitempty"`
}

// DisableCredentialResponse type
type DisableCredentialResponse struct {
	XMLName xml.Name `xml:"DisableCredentialResponse"`
}

// ResetAntipassbackViolation type
type ResetAntipassbackViolation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl ResetAntipassbackViolation"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`
}

// ResetAntipassbackViolationResponse type
type ResetAntipassbackViolationResponse struct {
	XMLName xml.Name `xml:"ResetAntipassbackViolationResponse"`
}

// GetCredentialIdentifiers type
type GetCredentialIdentifiers struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentialIdentifiers"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`
}

// GetCredentialIdentifiersResponse type
type GetCredentialIdentifiersResponse struct {
	XMLName xml.Name `xml:"GetCredentialIdentifiersResponse"`

	// Identifier of the credential
	CredentialIdentifier []CredentialIdentifier `xml:"CredentialIdentifier,omitempty"`
}

// SetCredentialIdentifier type
type SetCredentialIdentifier struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl SetCredentialIdentifier"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`

	// Identifier of the credential
	CredentialIdentifier CredentialIdentifier `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialIdentifier,omitempty"`
}

// SetCredentialIdentifierResponse type
type SetCredentialIdentifierResponse struct {
	XMLName xml.Name `xml:"SetCredentialIdentifierResponse"`
}

// DeleteCredentialIdentifier type
type DeleteCredentialIdentifier struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl DeleteCredentialIdentifier"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`

	// Identifier type name of a credential
	CredentialIdentifierTypeName Name `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialIdentifierTypeName,omitempty"`
}

// DeleteCredentialIdentifierResponse type
type DeleteCredentialIdentifierResponse struct {
	XMLName xml.Name `xml:"DeleteCredentialIdentifierResponse"`
}

// GetCredentialAccessProfiles type
type GetCredentialAccessProfiles struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl GetCredentialAccessProfiles"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`
}

// GetCredentialAccessProfilesResponse type
type GetCredentialAccessProfilesResponse struct {
	XMLName xml.Name `xml:"GetCredentialAccessProfilesResponse"`

	// Access Profiles of the credential
	CredentialAccessProfile []CredentialAccessProfile `xml:"CredentialAccessProfile,omitempty"`
}

// SetCredentialAccessProfiles type
type SetCredentialAccessProfiles struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl SetCredentialAccessProfiles"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`

	// Access Profiles of the credential
	CredentialAccessProfile []CredentialAccessProfile `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialAccessProfile,omitempty"`
}

// SetCredentialAccessProfilesResponse type
type SetCredentialAccessProfilesResponse struct {
	XMLName xml.Name `xml:"SetCredentialAccessProfilesResponse"`
}

// DeleteCredentialAccessProfiles type
type DeleteCredentialAccessProfiles struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl DeleteCredentialAccessProfiles"`

	// Token of the Credential
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialToken,omitempty"`

	// Tokens of Access Profiles
	AccessProfileToken []ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl AccessProfileToken,omitempty"`
}

// DeleteCredentialAccessProfilesResponse type
type DeleteCredentialAccessProfilesResponse struct {
	XMLName xml.Name `xml:"DeleteCredentialAccessProfilesResponse"`
}

// ServiceCapabilities type
type ServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/credential/wsdl Capabilities"`

	//
	// A list of identifier types that the device supports. Identifiers types starting with
	// the prefix pt: are reserved to define ONVIF specific types. For custom defined identifier types
	// shall all share the "pt:<Name>" syntax.
	//
	SupportedIdentifierType []Name `xml:"http://www.onvif.org/ver10/credential/wsdl SupportedIdentifierType,omitempty"`

	Extension ServiceCapabilitiesExtension `xml:"http://www.onvif.org/ver10/credential/wsdl Extension,omitempty"`

	//
	// The maximum number of entries returned by a single Get<Entity>List or Get<Entity>
	// request. The device shall never return more than this number of entities in a single response.
	//

	MaxLimit PositiveInteger `xml:"http://www.onvif.org/ver10/credential/wsdl MaxLimit,attr,omitempty"`

	//
	// Indicates that the device supports credential validity.
	//

	CredentialValiditySupported bool `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialValiditySupported,attr,omitempty"`

	//
	// Indicates that the device supports validity on the association between a credential and an
	// access profile.
	//

	CredentialAccessProfileValiditySupported bool `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialAccessProfileValiditySupported,attr,omitempty"`

	//
	// Indicates that the device supports both date and time value for validity. If set to false,
	// then the time value is ignored.
	//

	ValiditySupportsTimeValue bool `xml:"http://www.onvif.org/ver10/credential/wsdl ValiditySupportsTimeValue,attr,omitempty"`

	//
	// The maximum number of credential supported by the device.
	//

	MaxCredentials PositiveInteger `xml:"http://www.onvif.org/ver10/credential/wsdl MaxCredentials,attr,omitempty"`

	//
	// The maximum number of access profiles for a credential.
	//

	MaxAccessProfilesPerCredential PositiveInteger `xml:"http://www.onvif.org/ver10/credential/wsdl MaxAccessProfilesPerCredential,attr,omitempty"`

	//
	// Indicates the device supports resetting of anti-passback violations and notifying on
	// anti-passback violations.
	//

	ResetAntipassbackSupported bool `xml:"http://www.onvif.org/ver10/credential/wsdl ResetAntipassbackSupported,attr,omitempty"`

	//
	// Indicates that the client is allowed to supply the token when creating credentials.
	// To enable the use of the command SetCredential, the value must be set to true.
	//

	ClientSuppliedTokenSupported bool `xml:"http://www.onvif.org/ver10/credential/wsdl ClientSuppliedTokenSupported,attr,omitempty"`

	//
	// The default time period that the credential will temporary be suspended (e.g. by using
	// the wrong PIN a predetermined number of times).
	// The time period is defined as an [ISO 8601] duration string (e.g. “PT5M”).
	//

	DefaultCredentialSuspensionDuration Duration `xml:"http://www.onvif.org/ver10/schema DefaultCredentialSuspensionDuration,attr,omitempty"`
}

// ServiceCapabilitiesExtension type
type ServiceCapabilitiesExtension struct {

	//
	// A list of exemptions that the device supports. Supported exemptions starting with the
	// prefix pt: are reserved to define ONVIF specific exemption types and these reserved
	// exemption types shall all share "pt:<Name>" syntax.
	//
	SupportedExemptionType []Name `xml:"http://www.onvif.org/ver10/credential/wsdl SupportedExemptionType,omitempty"`
}

// CredentialInfo type
type CredentialInfo struct {
	*DataEntity

	//
	// User readable description for the credential. It shall be up to 1024 characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/credential/wsdl Description,omitempty"`

	// An external reference to a person holding this credential. The
	// reference is a username or used ID in an external system, such as a directory
	// service.
	//

	CredentialHolderReference string `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialHolderReference,omitempty"`

	// The start date/time validity of the credential. If the
	// ValiditySupportsTimeValue capability is set to false, then only date is
	// supported (time is ignored).
	//
	ValidFrom string `xml:"http://www.onvif.org/ver10/schema ValidFrom,omitempty"`

	// The expiration date/time validity of the credential. If the
	// ValiditySupportsTimeValue capability is set to false, then only date is
	// supported (time is ignored).
	//
	ValidTo string `xml:"http://www.onvif.org/ver10/schema ValidTo,omitempty"`
}

// Credential type
type Credential struct {
	*CredentialInfo

	// A list of credential identifier structures. At least one
	// credential identifier is required. Maximum one credential identifier structure
	// per type is allowed.
	//
	CredentialIdentifier []CredentialIdentifier `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialIdentifier,omitempty"`

	// A list of credential access profile structures.
	CredentialAccessProfile []CredentialAccessProfile `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialAccessProfile,omitempty"`

	//
	// A boolean indicating that the credential holder needs extra time to get through the door.
	// ExtendedReleaseTime will be added to ReleaseTime, and ExtendedOpenTime will be added to OpenTime
	//
	ExtendedGrantTime bool `xml:"http://www.onvif.org/ver10/credential/wsdl ExtendedGrantTime,omitempty"`

	// A list of credential attributes as name value pairs. Key names
	// starting with the prefix pt: are reserved to define PACS specific attributes
	// following the "pt:<Name>" syntax.
	//
	Attribute []Attribute `xml:"http://www.onvif.org/ver10/credential/wsdl Attribute,omitempty"`

	Extension CredentialExtension `xml:"http://www.onvif.org/ver10/credential/wsdl Extension,omitempty"`
}

// CredentialExtension type
type CredentialExtension struct {
}

// CredentialIdentifier type
type CredentialIdentifier struct {

	// Contains the details of the credential identifier type. Is of type
	// CredentialIdentifierType.
	//
	Type CredentialIdentifierType `xml:"http://www.onvif.org/ver10/credential/wsdl Type,omitempty"`

	// If set to true, this credential identifier is not considered for
	// authentication. For example if the access point requests Card plus PIN, and the credential
	// identifier of type PIN is exempted from authentication, then the access point will not prompt
	// for the PIN.
	//
	ExemptedFromAuthentication bool `xml:"http://www.onvif.org/ver10/credential/wsdl ExemptedFromAuthentication,omitempty"`

	// The value of the identifier in hexadecimal representation.
	//
	Value []byte `xml:"http://www.onvif.org/ver10/schema Value,omitempty"`
}

// CredentialIdentifierType type
type CredentialIdentifierType struct {

	// The name of the credential identifier type, such as pt:Card, pt:PIN,
	// etc.
	//
	Name Name `xml:"http://www.onvif.org/ver10/credential/wsdl Name,omitempty"`

	// Specifies the format of the credential value for the specified identifier
	// type name.
	//
	FormatType string `xml:"http://www.onvif.org/ver10/credential/wsdl FormatType,omitempty"`
}

// CredentialAccessProfile type
type CredentialAccessProfile struct {

	// The reference token of the associated access profile.
	AccessProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/credential/wsdl AccessProfileToken,omitempty"`

	// The start date/time of the validity for the association between the
	// credential and the access profile. If the ValiditySupportsTimeValue capability is set to
	// false, then only date is supported (time is ignored).
	//
	ValidFrom string `xml:"http://www.onvif.org/ver10/schema ValidFrom,omitempty"`

	// The end date/time of the validity for the association between the
	// credential and the access profile. If the ValiditySupportsTimeValue capability is set to
	// false, then only date is supported (time is ignored).
	//
	ValidTo string `xml:"http://www.onvif.org/ver10/schema ValidTo,omitempty"`
}

// CredentialState type
type CredentialState struct {

	// True if the credential is enabled or false if the credential is
	// disabled.
	//
	Enabled bool `xml:"http://www.onvif.org/ver10/credential/wsdl Enabled,omitempty"`

	// Predefined ONVIF reasons as mentioned in the section 5.4.2.7
	// of credential service specification document. For any other reason, free
	// text can be used.
	//
	Reason Name `xml:"http://www.onvif.org/ver10/credential/wsdl Reason,omitempty"`

	// A structure indicating the anti-passback state. This field shall be
	// supported if the ResetAntipassbackSupported capability is set to true.
	//
	AntipassbackState AntipassbackState `xml:"http://www.onvif.org/ver10/credential/wsdl AntipassbackState,omitempty"`

	Extension CredentialStateExtension `xml:"http://www.onvif.org/ver10/credential/wsdl Extension,omitempty"`
}

// CredentialStateExtension type
type CredentialStateExtension struct {
}

// AntipassbackState type
type AntipassbackState struct {

	// Indicates if anti-passback is violated for the credential.
	//
	AntipassbackViolated bool `xml:"http://www.onvif.org/ver10/credential/wsdl AntipassbackViolated,omitempty"`
}

// CredentialIdentifierFormatTypeInfo type
type CredentialIdentifierFormatTypeInfo struct {

	// A format type supported by the device. A list of supported format types is
	// provided in [ISO 16484-5:2014-09 Annex P]. The BACnet type "CUSTOM" is not used in this
	// specification. Instead device manufacturers can define their own format types.
	//
	FormatType string `xml:"http://www.onvif.org/ver10/credential/wsdl FormatType,omitempty"`

	// User readable description of the credential identifier format type. It
	// shall be up to 1024 characters. For custom types, it is recommended to describe how the
	// octet string is encoded (following the structure in column Authentication Factor Value
	// Encoding of [ISO 16484-5:2014-09 Annex P]).
	//
	Description Description `xml:"http://www.onvif.org/ver10/credential/wsdl Description,omitempty"`

	Extension CredentialIdentifierFormatTypeInfoExtension `xml:"http://www.onvif.org/ver10/credential/wsdl Extension,omitempty"`
}

// CredentialIdentifierFormatTypeInfoExtension type
type CredentialIdentifierFormatTypeInfoExtension struct {
}

// CredentialData type
type CredentialData struct {

	// A format type supported by the device. A list of supported format types is
	// provided in [ISO 16484-5:2014-09 Annex P]. The BACnet type "CUSTOM" is not used in this
	// specification. Instead device manufacturers can define their own format types.
	//
	Credential Credential `xml:"http://www.onvif.org/ver10/credential/wsdl Credential,omitempty"`

	// User readable description of the credential identifier format type. It
	// shall be up to 1024 characters. For custom types, it is recommended to describe how the
	// octet string is encoded (following the structure in column Authentication Factor Value
	// Encoding of [ISO 16484-5:2014-09 Annex P]).
	//
	CredentialState CredentialState `xml:"http://www.onvif.org/ver10/credential/wsdl CredentialState,omitempty"`

	Extension CredentialDataExtension `xml:"http://www.onvif.org/ver10/credential/wsdl Extension,omitempty"`
}

// CredentialDataExtension type
type CredentialDataExtension struct {
}

// FaultResponse type
type FaultResponse struct {

	// A format type supported by the device. A list of supported format types is
	// provided in [ISO 16484-5:2014-09 Annex P]. The BACnet type "CUSTOM" is not used in this
	// specification. Instead device manufacturers can define their own format types.
	//
	Token ReferenceToken `xml:"Token,omitempty"`

	// User readable description of the credential identifier format type. It
	// shall be up to 1024 characters. For custom types, it is recommended to describe how the
	// octet string is encoded (following the structure in column Authentication Factor Value
	// Encoding of [ISO 16484-5:2014-09 Annex P]).
	//
	Fault string `xml:"Fault,omitempty"`

	Extension FaultResponseExtension `xml:"Extension,omitempty"`
}

// FaultResponseExtension type
type FaultResponseExtension struct {
}

// Type used to reference logical and physical entities.

// ReferenceToken type
type ReferenceToken string

// Type used for names of logical and physical entities.

// Name type
type Name string

// Description is optional and the maximum length is device specific.
// If the length is more than maximum length, it is silently chopped to the maximum length
// supported by the device/service (which may be 0).
//

// Description type
type Description string

// Type used to represent the numbers from 1 ,2 , 3,...

// PositiveInteger type
type PositiveInteger uint32

// DataEntity type
type DataEntity struct {

	// A service-unique identifier of the item.

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

// Attribute type
type Attribute struct {

	// Name of attribute. Key names starting with "ONVIF" (any case) are reserved for ONVIF
	// use.
	//

	Name string `xml:"http://www.onvif.org/ver10/credential/wsdl Name,attr,omitempty"`

	// Value of attribute

	Value string `xml:"http://www.onvif.org/ver10/credential/wsdl Value,attr,omitempty"`
}

// CredentialPort type
type CredentialPort interface {

	/* This operation returns the capabilities of the credential service. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/*
		This method returns all the supported format types of a specified identifier type that is supported by
		the device.
	*/
	GetSupportedFormatTypes(request *GetSupportedFormatTypes) (*GetSupportedFormatTypesResponse, error)

	GetSupportedFormatTypesContext(ctx context.Context, request *GetSupportedFormatTypes) (*GetSupportedFormatTypesResponse, error)

	/*
	   This operation requests a list of CredentialInfo items matching the given tokens.
	   The device shall ignore tokens it cannot resolve and shall return an empty list if there are no
	   items matching the specified tokens. The device shall not return a fault in this case.
	   If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetCredentialInfo(request *GetCredentialInfo) (*GetCredentialInfoResponse, error)

	GetCredentialInfoContext(ctx context.Context, request *GetCredentialInfo) (*GetCredentialInfoResponse, error)

	/*
	   This operation requests a list of all CredentialInfo items provided by the device.
	   A call to this method shall return a StartReference when not all data is returned and more data is available.
	   The reference shall be valid for retrieving the next set of data. Please refer to section 4.8.3 in
	   [ONVIF Access Control Service Specification] for more details.
	   The number of items returned shall not be greater than the Limit parameter.
	*/
	GetCredentialInfoList(request *GetCredentialInfoList) (*GetCredentialInfoListResponse, error)

	GetCredentialInfoListContext(ctx context.Context, request *GetCredentialInfoList) (*GetCredentialInfoListResponse, error)

	/*
		This operation returns the specified credential items matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list if there are no items
		matching specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetCredentials(request *GetCredentials) (*GetCredentialsResponse, error)

	GetCredentialsContext(ctx context.Context, request *GetCredentials) (*GetCredentialsResponse, error)

	/*
		This operation requests a list of all credential items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and more data is
		available. The reference shall be valid for retrieving the next set of data. Please refer section 4.8.3
		in [Access Control Service Specification] for more details. The number of items returned shall not be
		greater the Limit parameter.
	*/
	GetCredentialList(request *GetCredentialList) (*GetCredentialListResponse, error)

	GetCredentialListContext(ctx context.Context, request *GetCredentialList) (*GetCredentialListResponse, error)

	/*
		This operation creates a credential. A call to this method takes a credential structure and a credential
		state structure as input parameters. The credential state can be created in disabled or enabled state.
		The token field of the credential shall be empty, the device shall allocate a token for the credential.
		The allocated token shall be returned in the response. If the client sends any value in the token field,
		the device shall return InvalidArgVal as generic fault code.
	*/
	CreateCredential(request *CreateCredential) (*CreateCredentialResponse, error)

	CreateCredentialContext(ctx context.Context, request *CreateCredential) (*CreateCredentialResponse, error)

	/*
		This method is used to synchronize a credential in a client with the device.
	*/
	SetCredential(request *SetCredential) (*SetCredentialResponse, error)

	SetCredentialContext(ctx context.Context, request *SetCredential) (*SetCredentialResponse, error)

	/*
	   This operation modifies the specified credential.
	   The token of the credential to modify is specified in the token field of the Credential structure and
	   shall not be empty. All other fields in the structure shall overwrite the fields in the specified credential.
	   When an existing credential is modified, the state is not modified explicitly. The only way for a client to
	   change the state of a credential is to explicitly call the EnableCredential, DisableCredential or
	   ResetAntipassback command.
	   All existing credential identifiers and credential access profiles are removed and replaced with the
	   specified entities.
	*/
	ModifyCredential(request *ModifyCredential) (*ModifyCredentialResponse, error)

	ModifyCredentialContext(ctx context.Context, request *ModifyCredential) (*ModifyCredentialResponse, error)

	/*
	   This method deletes the specified credential.
	   If it is associated with one or more entities some devices may not be able to delete the credential,
	   and consequently a ReferenceInUse fault shall be generated.
	*/
	DeleteCredential(request *DeleteCredential) (*DeleteCredentialResponse, error)

	DeleteCredentialContext(ctx context.Context, request *DeleteCredential) (*DeleteCredentialResponse, error)

	/*
		This method returns the state for the specified credential.
		If the capability ResetAntipassbackSupported is set to true, then the device shall supply the
		anti-passback state in the returned credential state structure.
	*/
	GetCredentialState(request *GetCredentialState) (*GetCredentialStateResponse, error)

	GetCredentialStateContext(ctx context.Context, request *GetCredentialState) (*GetCredentialStateResponse, error)

	/*
		This method is used to enable a credential.
	*/
	EnableCredential(request *EnableCredential) (*EnableCredentialResponse, error)

	EnableCredentialContext(ctx context.Context, request *EnableCredential) (*EnableCredentialResponse, error)

	/*
		This method is used to disable a credential.
	*/
	DisableCredential(request *DisableCredential) (*DisableCredentialResponse, error)

	DisableCredentialContext(ctx context.Context, request *DisableCredential) (*DisableCredentialResponse, error)

	/*
		This method is used to reset anti-passback violations for a specified credential.
	*/
	ResetAntipassbackViolation(request *ResetAntipassbackViolation) (*ResetAntipassbackViolationResponse, error)

	ResetAntipassbackViolationContext(ctx context.Context, request *ResetAntipassbackViolation) (*ResetAntipassbackViolationResponse, error)

	/*
		This method returns all the credential identifiers for a credential.
	*/
	GetCredentialIdentifiers(request *GetCredentialIdentifiers) (*GetCredentialIdentifiersResponse, error)

	GetCredentialIdentifiersContext(ctx context.Context, request *GetCredentialIdentifiers) (*GetCredentialIdentifiersResponse, error)

	/*
		This operation creates or updates a credential identifier for a credential.
		If the type of specified credential identifier already exists, the current credential identifier of that
		type is replaced. Otherwise the credential identifier is added.
	*/
	SetCredentialIdentifier(request *SetCredentialIdentifier) (*SetCredentialIdentifierResponse, error)

	SetCredentialIdentifierContext(ctx context.Context, request *SetCredentialIdentifier) (*SetCredentialIdentifierResponse, error)

	/*
		This method deletes all the identifier values for the specified type. However, if the identifier type
		name doesn’t exist in the device, it will be silently ignored without any response.
	*/
	DeleteCredentialIdentifier(request *DeleteCredentialIdentifier) (*DeleteCredentialIdentifierResponse, error)

	DeleteCredentialIdentifierContext(ctx context.Context, request *DeleteCredentialIdentifier) (*DeleteCredentialIdentifierResponse, error)

	/*
		This method returns all the credential access profiles for a credential.
	*/
	GetCredentialAccessProfiles(request *GetCredentialAccessProfiles) (*GetCredentialAccessProfilesResponse, error)

	GetCredentialAccessProfilesContext(ctx context.Context, request *GetCredentialAccessProfiles) (*GetCredentialAccessProfilesResponse, error)

	/*
		This operation add or updates the credential access profiles for a credential.
		The device shall update the credential access profile if the access profile token in the specified
		credential access profile matches. Otherwise the credential access profile is added.
	*/
	SetCredentialAccessProfiles(request *SetCredentialAccessProfiles) (*SetCredentialAccessProfilesResponse, error)

	SetCredentialAccessProfilesContext(ctx context.Context, request *SetCredentialAccessProfiles) (*SetCredentialAccessProfilesResponse, error)

	/*
		This method deletes all the credential access profiles for the specified tokens.
		However, if no matching credential access profiles are found, the corresponding access profile tokens
		are silently ignored without any response.
	*/
	DeleteCredentialAccessProfiles(request *DeleteCredentialAccessProfiles) (*DeleteCredentialAccessProfilesResponse, error)

	DeleteCredentialAccessProfilesContext(ctx context.Context, request *DeleteCredentialAccessProfiles) (*DeleteCredentialAccessProfilesResponse, error)
}

// credentialPort type
type credentialPort struct {
	client *soap.Client
	xaddr  string
}

func NewCredentialPort(client *soap.Client, xaddr string) CredentialPort {
	return &credentialPort{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *credentialPort) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetSupportedFormatTypesContext(ctx context.Context, request *GetSupportedFormatTypes) (*GetSupportedFormatTypesResponse, error) {
	response := new(GetSupportedFormatTypesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetSupportedFormatTypes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetSupportedFormatTypes(request *GetSupportedFormatTypes) (*GetSupportedFormatTypesResponse, error) {
	return service.GetSupportedFormatTypesContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialInfoContext(ctx context.Context, request *GetCredentialInfo) (*GetCredentialInfoResponse, error) {
	response := new(GetCredentialInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentialInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentialInfo(request *GetCredentialInfo) (*GetCredentialInfoResponse, error) {
	return service.GetCredentialInfoContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialInfoListContext(ctx context.Context, request *GetCredentialInfoList) (*GetCredentialInfoListResponse, error) {
	response := new(GetCredentialInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentialInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentialInfoList(request *GetCredentialInfoList) (*GetCredentialInfoListResponse, error) {
	return service.GetCredentialInfoListContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialsContext(ctx context.Context, request *GetCredentials) (*GetCredentialsResponse, error) {
	response := new(GetCredentialsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentials", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentials(request *GetCredentials) (*GetCredentialsResponse, error) {
	return service.GetCredentialsContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialListContext(ctx context.Context, request *GetCredentialList) (*GetCredentialListResponse, error) {
	response := new(GetCredentialListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentialList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentialList(request *GetCredentialList) (*GetCredentialListResponse, error) {
	return service.GetCredentialListContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) CreateCredentialContext(ctx context.Context, request *CreateCredential) (*CreateCredentialResponse, error) {
	response := new(CreateCredentialResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/CreateCredential", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) CreateCredential(request *CreateCredential) (*CreateCredentialResponse, error) {
	return service.CreateCredentialContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) SetCredentialContext(ctx context.Context, request *SetCredential) (*SetCredentialResponse, error) {
	response := new(SetCredentialResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/SetCredential", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) SetCredential(request *SetCredential) (*SetCredentialResponse, error) {
	return service.SetCredentialContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) ModifyCredentialContext(ctx context.Context, request *ModifyCredential) (*ModifyCredentialResponse, error) {
	response := new(ModifyCredentialResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/ModifyCredential", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) ModifyCredential(request *ModifyCredential) (*ModifyCredentialResponse, error) {
	return service.ModifyCredentialContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) DeleteCredentialContext(ctx context.Context, request *DeleteCredential) (*DeleteCredentialResponse, error) {
	response := new(DeleteCredentialResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/DeleteCredential", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) DeleteCredential(request *DeleteCredential) (*DeleteCredentialResponse, error) {
	return service.DeleteCredentialContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialStateContext(ctx context.Context, request *GetCredentialState) (*GetCredentialStateResponse, error) {
	response := new(GetCredentialStateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentialState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentialState(request *GetCredentialState) (*GetCredentialStateResponse, error) {
	return service.GetCredentialStateContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) EnableCredentialContext(ctx context.Context, request *EnableCredential) (*EnableCredentialResponse, error) {
	response := new(EnableCredentialResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/EnableCredential", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) EnableCredential(request *EnableCredential) (*EnableCredentialResponse, error) {
	return service.EnableCredentialContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) DisableCredentialContext(ctx context.Context, request *DisableCredential) (*DisableCredentialResponse, error) {
	response := new(DisableCredentialResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/DisableCredential", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) DisableCredential(request *DisableCredential) (*DisableCredentialResponse, error) {
	return service.DisableCredentialContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) ResetAntipassbackViolationContext(ctx context.Context, request *ResetAntipassbackViolation) (*ResetAntipassbackViolationResponse, error) {
	response := new(ResetAntipassbackViolationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/ResetAntipassbackViolation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) ResetAntipassbackViolation(request *ResetAntipassbackViolation) (*ResetAntipassbackViolationResponse, error) {
	return service.ResetAntipassbackViolationContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialIdentifiersContext(ctx context.Context, request *GetCredentialIdentifiers) (*GetCredentialIdentifiersResponse, error) {
	response := new(GetCredentialIdentifiersResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentialIdentifiers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentialIdentifiers(request *GetCredentialIdentifiers) (*GetCredentialIdentifiersResponse, error) {
	return service.GetCredentialIdentifiersContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) SetCredentialIdentifierContext(ctx context.Context, request *SetCredentialIdentifier) (*SetCredentialIdentifierResponse, error) {
	response := new(SetCredentialIdentifierResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/SetCredentialIdentifier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) SetCredentialIdentifier(request *SetCredentialIdentifier) (*SetCredentialIdentifierResponse, error) {
	return service.SetCredentialIdentifierContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) DeleteCredentialIdentifierContext(ctx context.Context, request *DeleteCredentialIdentifier) (*DeleteCredentialIdentifierResponse, error) {
	response := new(DeleteCredentialIdentifierResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/DeleteCredentialIdentifier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) DeleteCredentialIdentifier(request *DeleteCredentialIdentifier) (*DeleteCredentialIdentifierResponse, error) {
	return service.DeleteCredentialIdentifierContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) GetCredentialAccessProfilesContext(ctx context.Context, request *GetCredentialAccessProfiles) (*GetCredentialAccessProfilesResponse, error) {
	response := new(GetCredentialAccessProfilesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/GetCredentialAccessProfiles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) GetCredentialAccessProfiles(request *GetCredentialAccessProfiles) (*GetCredentialAccessProfilesResponse, error) {
	return service.GetCredentialAccessProfilesContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) SetCredentialAccessProfilesContext(ctx context.Context, request *SetCredentialAccessProfiles) (*SetCredentialAccessProfilesResponse, error) {
	response := new(SetCredentialAccessProfilesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/SetCredentialAccessProfiles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) SetCredentialAccessProfiles(request *SetCredentialAccessProfiles) (*SetCredentialAccessProfilesResponse, error) {
	return service.SetCredentialAccessProfilesContext(
		context.Background(),
		request,
	)
}

func (service *credentialPort) DeleteCredentialAccessProfilesContext(ctx context.Context, request *DeleteCredentialAccessProfiles) (*DeleteCredentialAccessProfilesResponse, error) {
	response := new(DeleteCredentialAccessProfilesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/credential/wsdl/DeleteCredentialAccessProfiles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *credentialPort) DeleteCredentialAccessProfiles(request *DeleteCredentialAccessProfiles) (*DeleteCredentialAccessProfilesResponse, error) {
	return service.DeleteCredentialAccessProfilesContext(
		context.Background(),
		request,
	)
}

// Duration type
type Duration string
