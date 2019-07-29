package accesscontrol

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

//
// The Decision enumeration represents a choice of two available options for an access request:
//

// Decision type
type Decision string

const (

	// The decision is to grant access.
	DecisionGranted Decision = "Granted"

	// The decision is to deny access.
	DecisionDenied Decision = "Denied"
)

//
// Non-normative enum that describes the various reasons for denying access.
// The following strings shall be used for the reason field:
//

// DenyReason type
type DenyReason string

const (

	// The device shall provide the following event, whenever a valid credential
	// is not enabled or has been disabled (e.g., due to credential being lost etc.) to prevent
	// unauthorized entry.
	//
	DenyReasonCredentialNotEnabled DenyReason = "CredentialNotEnabled"

	// The device shall provide the following event, whenever a valid credential
	// is presented though it is not active yet;: e.g, the credential was presented before the
	// start date.
	//
	DenyReasonCredentialNotActive DenyReason = "CredentialNotActive"

	// The device shall provide the following event, whenever a valid credential
	// was presented after its expiry date.
	//
	DenyReasonCredentialExpired DenyReason = "CredentialExpired"

	// The device shall provide the following event, whenever an entered PIN code
	// does not match the credential.
	//
	DenyReasonInvalidPIN DenyReason = "InvalidPIN"

	// The device shall provide the following event, whenever a valid credential
	// is denied access to the requested AccessPoint because the credential is not permitted at
	// the moment.
	//
	DenyReasonNotPermittedAtThisTime DenyReason = "NotPermittedAtThisTime"

	// The device shall provide the following event, whenever the presented
	// credential is not authorized.
	//
	DenyReasonUnauthorized DenyReason = "Unauthorized"

	// The device shall provide the following event, whenever the request is
	// denied and no other specific event matches it or is supported by the service.
	//
	DenyReasonOther DenyReason = "Other"
)

// Capabilities type
type Capabilities ServiceCapabilities

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capability response message contains the requested Access Control
	// service capabilities using a hierarchical XML capability structure.
	//
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty"`
}

// GetAccessPointInfoList type
type GetAccessPointInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAccessPointInfoList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl StartReference,omitempty"`
}

// GetAccessPointInfoListResponse type
type GetAccessPointInfoListResponse struct {
	XMLName xml.Name `xml:"GetAccessPointInfoListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of AccessPointInfo items.
	AccessPointInfo []AccessPointInfo `xml:"AccessPointInfo,omitempty"`
}

// GetAccessPointInfo type
type GetAccessPointInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAccessPointInfo"`

	// Tokens of AccessPointInfo items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// GetAccessPointInfoResponse type
type GetAccessPointInfoResponse struct {
	XMLName xml.Name `xml:"GetAccessPointInfoResponse"`

	// List of AccessPointInfo items.
	AccessPointInfo []AccessPointInfo `xml:"AccessPointInfo,omitempty"`
}

// GetAccessPointList type
type GetAccessPointList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAccessPointList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl StartReference,omitempty"`
}

// GetAccessPointListResponse type
type GetAccessPointListResponse struct {
	XMLName xml.Name `xml:"GetAccessPointListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of AccessPoint items.
	AccessPoint []AccessPoint `xml:"AccessPoint,omitempty"`
}

// GetAccessPoints type
type GetAccessPoints struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAccessPoints"`

	// Tokens of AccessPoint items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// GetAccessPointsResponse type
type GetAccessPointsResponse struct {
	XMLName xml.Name `xml:"GetAccessPointsResponse"`

	// List of AccessPoint items.
	AccessPoint []AccessPoint `xml:"AccessPoint,omitempty"`
}

// CreateAccessPoint type
type CreateAccessPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl CreateAccessPoint"`

	// AccessPoint item to create
	AccessPoint AccessPoint `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AccessPoint,omitempty"`
}

// CreateAccessPointResponse type
type CreateAccessPointResponse struct {
	XMLName xml.Name `xml:"CreateAccessPointResponse"`

	// Token of created AccessPoint item
	Token ReferenceToken `xml:"Token,omitempty"`
}

// SetAccessPoint type
type SetAccessPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl SetAccessPoint"`

	// AccessPoint item to create or modify
	AccessPoint AccessPoint `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AccessPoint,omitempty"`
}

// SetAccessPointResponse type
type SetAccessPointResponse struct {
	XMLName xml.Name `xml:"SetAccessPointResponse"`
}

// ModifyAccessPoint type
type ModifyAccessPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl ModifyAccessPoint"`

	// AccessPoint item to modify
	AccessPoint AccessPoint `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AccessPoint,omitempty"`
}

// ModifyAccessPointResponse type
type ModifyAccessPointResponse struct {
	XMLName xml.Name `xml:"ModifyAccessPointResponse"`
}

// DeleteAccessPoint type
type DeleteAccessPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl DeleteAccessPoint"`

	// Token of AccessPoint item to delete.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// DeleteAccessPointResponse type
type DeleteAccessPointResponse struct {
	XMLName xml.Name `xml:"DeleteAccessPointResponse"`
}

// SetAccessPointAuthenticationProfile type
type SetAccessPointAuthenticationProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl SetAccessPointAuthenticationProfile"`

	// Token of the AccessPoint.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`

	// Token of the AuthenticationProfile.
	AuthenticationProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AuthenticationProfileToken,omitempty"`
}

// SetAccessPointAuthenticationProfileResponse type
type SetAccessPointAuthenticationProfileResponse struct {
	XMLName xml.Name `xml:"SetAccessPointAuthenticationProfileResponse"`
}

// DeleteAccessPointAuthenticationProfile type
type DeleteAccessPointAuthenticationProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl DeleteAccessPointAuthenticationProfile"`

	// Token of the AccessPoint.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// DeleteAccessPointAuthenticationProfileResponse type
type DeleteAccessPointAuthenticationProfileResponse struct {
	XMLName xml.Name `xml:"DeleteAccessPointAuthenticationProfileResponse"`
}

// GetAreaInfoList type
type GetAreaInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAreaInfoList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl StartReference,omitempty"`
}

// GetAreaInfoListResponse type
type GetAreaInfoListResponse struct {
	XMLName xml.Name `xml:"GetAreaInfoListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of AreaInfo items.
	AreaInfo []AreaInfo `xml:"AreaInfo,omitempty"`
}

// GetAreaInfo type
type GetAreaInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAreaInfo"`

	// Tokens of AreaInfo items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// GetAreaInfoResponse type
type GetAreaInfoResponse struct {
	XMLName xml.Name `xml:"GetAreaInfoResponse"`

	// List of AreaInfo items.
	AreaInfo []AreaInfo `xml:"AreaInfo,omitempty"`
}

// GetAreaList type
type GetAreaList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAreaList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl StartReference,omitempty"`
}

// GetAreaListResponse type
type GetAreaListResponse struct {
	XMLName xml.Name `xml:"GetAreaListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of Area items.
	Area []Area `xml:"Area,omitempty"`
}

// GetAreas type
type GetAreas struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAreas"`

	// Tokens of Area items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// GetAreasResponse type
type GetAreasResponse struct {
	XMLName xml.Name `xml:"GetAreasResponse"`

	// List of Area items.
	Area []Area `xml:"Area,omitempty"`
}

// CreateArea type
type CreateArea struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl CreateArea"`

	// Area item to create
	Area Area `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Area,omitempty"`
}

// CreateAreaResponse type
type CreateAreaResponse struct {
	XMLName xml.Name `xml:"CreateAreaResponse"`

	// Token of created Area item
	Token ReferenceToken `xml:"Token,omitempty"`
}

// SetArea type
type SetArea struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl SetArea"`

	// Area item to create or modify
	Area Area `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Area,omitempty"`
}

// SetAreaResponse type
type SetAreaResponse struct {
	XMLName xml.Name `xml:"SetAreaResponse"`
}

// ModifyArea type
type ModifyArea struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl ModifyArea"`

	// Area item to modify
	Area Area `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Area,omitempty"`
}

// ModifyAreaResponse type
type ModifyAreaResponse struct {
	XMLName xml.Name `xml:"ModifyAreaResponse"`
}

// DeleteArea type
type DeleteArea struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl DeleteArea"`

	// Token of Area item to delete.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// DeleteAreaResponse type
type DeleteAreaResponse struct {
	XMLName xml.Name `xml:"DeleteAreaResponse"`
}

// GetAccessPointState type
type GetAccessPointState struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl GetAccessPointState"`

	// Token of AccessPoint instance to get AccessPointState for.
	//
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// GetAccessPointStateResponse type
type GetAccessPointStateResponse struct {
	XMLName xml.Name `xml:"GetAccessPointStateResponse"`

	// AccessPointState item.
	AccessPointState AccessPointState `xml:"AccessPointState,omitempty"`
}

// EnableAccessPoint type
type EnableAccessPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl EnableAccessPoint"`

	// Token of the AccessPoint instance to enable.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// EnableAccessPointResponse type
type EnableAccessPointResponse struct {
	XMLName xml.Name `xml:"EnableAccessPointResponse"`
}

// DisableAccessPoint type
type DisableAccessPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl DisableAccessPoint"`

	// Token of the AccessPoint instance to disable.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Token,omitempty"`
}

// DisableAccessPointResponse type
type DisableAccessPointResponse struct {
	XMLName xml.Name `xml:"DisableAccessPointResponse"`
}

// ExternalAuthorization type
type ExternalAuthorization struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl ExternalAuthorization"`

	// Token of the Access Point instance.
	AccessPointToken ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AccessPointToken,omitempty"`

	// Optional token of the Credential involved.
	CredentialToken ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl CredentialToken,omitempty"`

	// Optional reason for decision.
	Reason string `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Reason,omitempty"`

	// Decision - Granted or Denied.
	Decision Decision `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Decision,omitempty"`
}

// ExternalAuthorizationResponse type
type ExternalAuthorizationResponse struct {
	XMLName xml.Name `xml:"ExternalAuthorizationResponse"`
}

// ServiceCapabilities type
type ServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Capabilities"`

	// The maximum number of entries returned by a single Get<Entity>List or
	// Get<Entity> request.
	// The device shall never return more than this number of entities in a single response.
	//

	MaxLimit uint32 `xml:"http://www.onvif.org/ver10/schema MaxLimit,attr,omitempty"`

	//
	// Indicates the maximum number of access points supported by the device.
	//

	MaxAccessPoints uint32 `xml:"http://www.onvif.org/ver10/schema MaxAccessPoints,attr,omitempty"`

	//
	// Indicates the maximum number of areas supported by the device.
	//

	MaxAreas uint32 `xml:"http://www.onvif.org/ver10/schema MaxAreas,attr,omitempty"`

	//
	// Indicates that the client is allowed to supply the token when creating access
	// points and areas.
	// To enable the use of the commands SetAccessPoint and SetArea, the value must be set to true.
	//

	ClientSuppliedTokenSupported bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl ClientSuppliedTokenSupported,attr,omitempty"`
}

// AccessPointInfoBase type
type AccessPointInfoBase struct {
	*DataEntity

	// A user readable name. It shall be up to 64 characters.
	//
	Name Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Name,omitempty"`

	// Optional user readable description for the AccessPoint. It shall
	// be up to 1024 characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Description,omitempty"`

	// Optional reference to the Area from which access is requested.
	//
	AreaFrom ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AreaFrom,omitempty"`

	// Optional reference to the Area to which access is requested.
	//
	AreaTo ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AreaTo,omitempty"`

	//
	// Optional entity type; if missing, a Door type as defined by [ONVIF Door Control
	// Service Specification] should be assumed. This can also be represented by the
	// QName value "tdc:Door" – where tdc is the namespace of the door control service:
	// "http://www.onvif.org/ver10/doorcontrol/wsdl". This field is provided for future
	// extensions; it will allow an access point being extended to cover entity types
	// other than doors as well.
	//
	EntityType QName `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl EntityType,omitempty"`

	// Reference to the entity used to control access; the entity type
	// may be specified by the optional EntityType field explained below but is
	// typically a Door.
	//
	Entity ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Entity,omitempty"`
}

// AccessPointInfo type
type AccessPointInfo struct {
	*AccessPointInfoBase

	// The capabilities for the AccessPoint.
	Capabilities AccessPointCapabilities `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Capabilities,omitempty"`
}

// AccessPoint type
type AccessPoint struct {
	*AccessPointInfo

	//
	// A reference to an authentication profile which defines the authentication
	// behavior of the access point.
	//
	AuthenticationProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AuthenticationProfileToken,omitempty"`

	Extension AccessPointExtension `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Extension,omitempty"`
}

// AccessPointExtension type
type AccessPointExtension struct {
}

// AccessPointCapabilities type
type AccessPointCapabilities struct {

	// A list of security level tokens that this access point supports.
	// See [Authentication Behavior Service Specification].
	//
	SupportedSecurityLevels []ReferenceToken `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl SupportedSecurityLevels,omitempty"`

	Extension SupportedSecurityLevelsExtension `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Extension,omitempty"`

	//
	// Indicates whether or not this AccessPoint instance supports EnableAccessPoint
	// and DisableAccessPoint commands.
	//

	DisableAccessPoint bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl DisableAccessPoint,attr,omitempty"`

	//
	// Indicates whether or not this AccessPoint instance supports generation of duress events.
	//

	Duress bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Duress,attr,omitempty"`

	//
	// Indicates whether or not this AccessPoint has a REX switch or other input that
	// allows anonymous access.
	//

	AnonymousAccess bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AnonymousAccess,attr,omitempty"`

	//
	// Indicates whether or not this AccessPoint instance supports generation of
	// AccessTaken and AccessNotTaken events. If AnonymousAccess and AccessTaken are both true, it
	// indicates that the Anonymous versions of AccessTaken and AccessNotTaken are supported.
	//

	AccessTaken bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl AccessTaken,attr,omitempty"`

	//
	// Indicates whether or not this AccessPoint instance supports the
	// ExternalAuthorization operation and the generation of Request events. If AnonymousAccess and
	// ExternalAuthorization are both true, it indicates that the Anonymous version is supported as
	// well.
	//

	ExternalAuthorization bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl ExternalAuthorization,attr,omitempty"`
}

// SupportedSecurityLevelsExtension type
type SupportedSecurityLevelsExtension struct {
}

// AreaInfoBase type
type AreaInfoBase struct {
	*DataEntity

	// User readable name. It shall be up to 64 characters.
	//
	Name Name `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Name,omitempty"`

	// User readable description for the Area. It shall be up to 1024
	// characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Description,omitempty"`
}

// AreaInfo type
type AreaInfo struct {
	*AreaInfoBase
}

// Area type
type Area struct {
	*AreaInfo

	Extension AreaExtension `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Extension,omitempty"`
}

// AreaExtension type
type AreaExtension struct {
}

// AccessPointState type
type AccessPointState struct {

	// Indicates that the AccessPoint is enabled. By default this field value
	// shall be True, if the DisableAccessPoint capabilities is not supported.
	//
	Enabled bool `xml:"http://www.onvif.org/ver10/accesscontrol/wsdl Enabled,omitempty"`
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

// Removed Attribute by fixgen.py

// PACSPort type
type PACSPort interface {

	/*
		This operation returns the capabilities of the access control service.
		A device which provides the access control service shall implement this method.
	*/
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/*
		This operation requests a list of AccessPointInfo items matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list if
		there are no items matching the specified tokens.
		The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetAccessPointInfo(request *GetAccessPointInfo) (*GetAccessPointInfoResponse, error)

	GetAccessPointInfoContext(ctx context.Context, request *GetAccessPointInfo) (*GetAccessPointInfoResponse, error)

	/*
		This operation requests a list of all AccessPointInfo items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and more
		data is available. The reference shall be valid for retrieving the next set of data.
		Please refer to section 4.8.3 in [ONVIF PACS Architecture and Design Considerations] for more details.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetAccessPointInfoList(request *GetAccessPointInfoList) (*GetAccessPointInfoListResponse, error)

	GetAccessPointInfoListContext(ctx context.Context, request *GetAccessPointInfoList) (*GetAccessPointInfoListResponse, error)

	/*
		This operation requests a list of AccessPoint items matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list if there are
		no items matching the specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetAccessPoints(request *GetAccessPoints) (*GetAccessPointsResponse, error)

	GetAccessPointsContext(ctx context.Context, request *GetAccessPoints) (*GetAccessPointsResponse, error)

	/*
		This operation requests a list of all AccessPoint items provided by the device. A call to
		this method shall return a StartReference when not all data is returned and more data is available.
		The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetAccessPointList(request *GetAccessPointList) (*GetAccessPointListResponse, error)

	GetAccessPointListContext(ctx context.Context, request *GetAccessPointList) (*GetAccessPointListResponse, error)

	/*
		This operation creates the specified access point in the device. The token field of the
		AccessPoint structure shall be empty and the device shall allocate a token for the access point.
		The allocated token shall be returned in the response. If the client sends any value in the
		token field, the device shall return InvalidArgVal as a generic fault code.
	*/
	CreateAccessPoint(request *CreateAccessPoint) (*CreateAccessPointResponse, error)

	CreateAccessPointContext(ctx context.Context, request *CreateAccessPoint) (*CreateAccessPointResponse, error)

	/*
		This method is used to synchronize an access point in a client with the device. If an access
		point with the specified token does not exist in the device, the access point is created.
		If an access point with the specified token exists, then the access point is modified.
		A call to this method takes an AccessPoint structure as input parameter.
		The token field of the AccessPoint structure shall not be empty.
		A device that signals support for the ClientSuppliedTokenSupported capability shall implement this command.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	SetAccessPoint(request *SetAccessPoint) (*SetAccessPointResponse, error)

	SetAccessPointContext(ctx context.Context, request *SetAccessPoint) (*SetAccessPointResponse, error)

	/*
		This operation modifies the specified access point. The token of the access point to modify
		is specified in the token field of the AccessPoint structure and shall not be empty.
		All other fields in the structure shall overwrite the fields in the specified access point.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	ModifyAccessPoint(request *ModifyAccessPoint) (*ModifyAccessPointResponse, error)

	ModifyAccessPointContext(ctx context.Context, request *ModifyAccessPoint) (*ModifyAccessPointResponse, error)

	/*
		This operation deletes the specified access point. If it is associated with one or more
		entities some devices may not be able to delete the access point, and consequently a
		ReferenceInUse fault shall be generated.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	DeleteAccessPoint(request *DeleteAccessPoint) (*DeleteAccessPointResponse, error)

	DeleteAccessPointContext(ctx context.Context, request *DeleteAccessPoint) (*DeleteAccessPointResponse, error)

	/*
		This operation defines the authentication behavior for an access point.
	*/
	SetAccessPointAuthenticationProfile(request *SetAccessPointAuthenticationProfile) (*SetAccessPointAuthenticationProfileResponse, error)

	SetAccessPointAuthenticationProfileContext(ctx context.Context, request *SetAccessPointAuthenticationProfile) (*SetAccessPointAuthenticationProfileResponse, error)

	/*
		This operation reverts the authentication behavior for an access point to its default behavior.
	*/
	DeleteAccessPointAuthenticationProfile(request *DeleteAccessPointAuthenticationProfile) (*DeleteAccessPointAuthenticationProfileResponse, error)

	DeleteAccessPointAuthenticationProfileContext(ctx context.Context, request *DeleteAccessPointAuthenticationProfile) (*DeleteAccessPointAuthenticationProfileResponse, error)

	/*
		This operation requests a list of AreaInfo items matching the given tokens. The device shall
		ignore tokens it cannot resolve and shall return an empty list if there are no items
		matching the specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetAreaInfo(request *GetAreaInfo) (*GetAreaInfoResponse, error)

	GetAreaInfoContext(ctx context.Context, request *GetAreaInfo) (*GetAreaInfoResponse, error)

	/*
		This operation requests a list of all AreaInfo items provided by the device. A call to this
		method shall return a StartReference when not all data is returned and more data is available.
		The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetAreaInfoList(request *GetAreaInfoList) (*GetAreaInfoListResponse, error)

	GetAreaInfoListContext(ctx context.Context, request *GetAreaInfoList) (*GetAreaInfoListResponse, error)

	/*
		This operation requests a list of Area items matching the given tokens. The device shall
		ignore tokens it cannot resolve and shall return an empty list if there are no items
		matching the specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetAreas(request *GetAreas) (*GetAreasResponse, error)

	GetAreasContext(ctx context.Context, request *GetAreas) (*GetAreasResponse, error)

	/*
		This operation requests a list of all Area items provided by the device. A call to this
		method shall return a StartReference when not all data is returned and more data is available.
		The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetAreaList(request *GetAreaList) (*GetAreaListResponse, error)

	GetAreaListContext(ctx context.Context, request *GetAreaList) (*GetAreaListResponse, error)

	/*
		This operation creates the specified area in the device. The token field of the Area
		structure shall be empty and the device shall allocate a token for the area.
		The allocated token shall be returned in the response.
		If the client sends any value in the token field, the device shall return InvalidArgVal as a generic fault code.
	*/
	CreateArea(request *CreateArea) (*CreateAreaResponse, error)

	CreateAreaContext(ctx context.Context, request *CreateArea) (*CreateAreaResponse, error)

	/*
		This method is used to synchronize an area in a client with the device. If an area with the
		specified token does not exist in the device, the area is created. If an area with the
		specified token exists, then the area is modified. A call to this method takes an Area
		structure as input parameter. The token field of the Area structure shall not be empty.
		A device that signals support for the ClientSuppliedTokenSupported capability shall
		implement this command.
		If no token was specified in the request, the device shall return
		InvalidArgs as a generic fault code.
	*/
	SetArea(request *SetArea) (*SetAreaResponse, error)

	SetAreaContext(ctx context.Context, request *SetArea) (*SetAreaResponse, error)

	/*
		This operation modifies the specified area. The token of the area to modify is specified in
		the token field of the Area structure and shall not be empty. All other fields in the
		structure shall overwrite the fields in the specified area.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	ModifyArea(request *ModifyArea) (*ModifyAreaResponse, error)

	ModifyAreaContext(ctx context.Context, request *ModifyArea) (*ModifyAreaResponse, error)

	/*
		This operation deletes the specified area. If it is associated with one or more entities
		some devices may not be able to delete the area, and consequently a ReferenceInUse fault shall be generated.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	DeleteArea(request *DeleteArea) (*DeleteAreaResponse, error)

	DeleteAreaContext(ctx context.Context, request *DeleteArea) (*DeleteAreaResponse, error)

	/*
		This operation requests the AccessPointState for the access point instance specified by the token.
	*/
	GetAccessPointState(request *GetAccessPointState) (*GetAccessPointStateResponse, error)

	GetAccessPointStateContext(ctx context.Context, request *GetAccessPointState) (*GetAccessPointStateResponse, error)

	/*
		This operation allows enabling an access point. A device that signals support for
		DisableAccessPoint capability for a particular access point instance shall implement this command.
	*/
	EnableAccessPoint(request *EnableAccessPoint) (*EnableAccessPointResponse, error)

	EnableAccessPointContext(ctx context.Context, request *EnableAccessPoint) (*EnableAccessPointResponse, error)

	/*
		This operation allows disabling an access point. A device that signals support for the
		DisableAccessPoint capability for a particular access point instance shall implement this command.
	*/
	DisableAccessPoint(request *DisableAccessPoint) (*DisableAccessPointResponse, error)

	DisableAccessPointContext(ctx context.Context, request *DisableAccessPoint) (*DisableAccessPointResponse, error)

	/*
		This operation allows to deny or grant decision at an access point instance. A device that
		signals support for ExternalAuthorization capability for a particular access point instance
		shall implement this method.
	*/
	ExternalAuthorization(request *ExternalAuthorization) (*ExternalAuthorizationResponse, error)

	ExternalAuthorizationContext(ctx context.Context, request *ExternalAuthorization) (*ExternalAuthorizationResponse, error)
}

// pACSPort type
type pACSPort struct {
	client *soap.Client
	xaddr  string
}

func NewPACSPort(client *soap.Client, xaddr string) PACSPort {
	return &pACSPort{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *pACSPort) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAccessPointInfoContext(ctx context.Context, request *GetAccessPointInfo) (*GetAccessPointInfoResponse, error) {
	response := new(GetAccessPointInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAccessPointInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAccessPointInfo(request *GetAccessPointInfo) (*GetAccessPointInfoResponse, error) {
	return service.GetAccessPointInfoContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAccessPointInfoListContext(ctx context.Context, request *GetAccessPointInfoList) (*GetAccessPointInfoListResponse, error) {
	response := new(GetAccessPointInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAccessPointInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAccessPointInfoList(request *GetAccessPointInfoList) (*GetAccessPointInfoListResponse, error) {
	return service.GetAccessPointInfoListContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAccessPointsContext(ctx context.Context, request *GetAccessPoints) (*GetAccessPointsResponse, error) {
	response := new(GetAccessPointsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAccessPoints", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAccessPoints(request *GetAccessPoints) (*GetAccessPointsResponse, error) {
	return service.GetAccessPointsContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAccessPointListContext(ctx context.Context, request *GetAccessPointList) (*GetAccessPointListResponse, error) {
	response := new(GetAccessPointListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAccessPointList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAccessPointList(request *GetAccessPointList) (*GetAccessPointListResponse, error) {
	return service.GetAccessPointListContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) CreateAccessPointContext(ctx context.Context, request *CreateAccessPoint) (*CreateAccessPointResponse, error) {
	response := new(CreateAccessPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/CreateAccessPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) CreateAccessPoint(request *CreateAccessPoint) (*CreateAccessPointResponse, error) {
	return service.CreateAccessPointContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) SetAccessPointContext(ctx context.Context, request *SetAccessPoint) (*SetAccessPointResponse, error) {
	response := new(SetAccessPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/SetAccessPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) SetAccessPoint(request *SetAccessPoint) (*SetAccessPointResponse, error) {
	return service.SetAccessPointContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) ModifyAccessPointContext(ctx context.Context, request *ModifyAccessPoint) (*ModifyAccessPointResponse, error) {
	response := new(ModifyAccessPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/ModifyAccessPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) ModifyAccessPoint(request *ModifyAccessPoint) (*ModifyAccessPointResponse, error) {
	return service.ModifyAccessPointContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) DeleteAccessPointContext(ctx context.Context, request *DeleteAccessPoint) (*DeleteAccessPointResponse, error) {
	response := new(DeleteAccessPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/DeleteAccessPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) DeleteAccessPoint(request *DeleteAccessPoint) (*DeleteAccessPointResponse, error) {
	return service.DeleteAccessPointContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) SetAccessPointAuthenticationProfileContext(ctx context.Context, request *SetAccessPointAuthenticationProfile) (*SetAccessPointAuthenticationProfileResponse, error) {
	response := new(SetAccessPointAuthenticationProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/SetAccessPointAuthenticationProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) SetAccessPointAuthenticationProfile(request *SetAccessPointAuthenticationProfile) (*SetAccessPointAuthenticationProfileResponse, error) {
	return service.SetAccessPointAuthenticationProfileContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) DeleteAccessPointAuthenticationProfileContext(ctx context.Context, request *DeleteAccessPointAuthenticationProfile) (*DeleteAccessPointAuthenticationProfileResponse, error) {
	response := new(DeleteAccessPointAuthenticationProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/DeleteAccessPointAuthenticationProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) DeleteAccessPointAuthenticationProfile(request *DeleteAccessPointAuthenticationProfile) (*DeleteAccessPointAuthenticationProfileResponse, error) {
	return service.DeleteAccessPointAuthenticationProfileContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAreaInfoContext(ctx context.Context, request *GetAreaInfo) (*GetAreaInfoResponse, error) {
	response := new(GetAreaInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAreaInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAreaInfo(request *GetAreaInfo) (*GetAreaInfoResponse, error) {
	return service.GetAreaInfoContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAreaInfoListContext(ctx context.Context, request *GetAreaInfoList) (*GetAreaInfoListResponse, error) {
	response := new(GetAreaInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAreaInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAreaInfoList(request *GetAreaInfoList) (*GetAreaInfoListResponse, error) {
	return service.GetAreaInfoListContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAreasContext(ctx context.Context, request *GetAreas) (*GetAreasResponse, error) {
	response := new(GetAreasResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAreas", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAreas(request *GetAreas) (*GetAreasResponse, error) {
	return service.GetAreasContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAreaListContext(ctx context.Context, request *GetAreaList) (*GetAreaListResponse, error) {
	response := new(GetAreaListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAreaList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAreaList(request *GetAreaList) (*GetAreaListResponse, error) {
	return service.GetAreaListContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) CreateAreaContext(ctx context.Context, request *CreateArea) (*CreateAreaResponse, error) {
	response := new(CreateAreaResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/CreateArea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) CreateArea(request *CreateArea) (*CreateAreaResponse, error) {
	return service.CreateAreaContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) SetAreaContext(ctx context.Context, request *SetArea) (*SetAreaResponse, error) {
	response := new(SetAreaResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/SetArea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) SetArea(request *SetArea) (*SetAreaResponse, error) {
	return service.SetAreaContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) ModifyAreaContext(ctx context.Context, request *ModifyArea) (*ModifyAreaResponse, error) {
	response := new(ModifyAreaResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/ModifyArea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) ModifyArea(request *ModifyArea) (*ModifyAreaResponse, error) {
	return service.ModifyAreaContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) DeleteAreaContext(ctx context.Context, request *DeleteArea) (*DeleteAreaResponse, error) {
	response := new(DeleteAreaResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/DeleteArea", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) DeleteArea(request *DeleteArea) (*DeleteAreaResponse, error) {
	return service.DeleteAreaContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) GetAccessPointStateContext(ctx context.Context, request *GetAccessPointState) (*GetAccessPointStateResponse, error) {
	response := new(GetAccessPointStateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/GetAccessPointState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) GetAccessPointState(request *GetAccessPointState) (*GetAccessPointStateResponse, error) {
	return service.GetAccessPointStateContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) EnableAccessPointContext(ctx context.Context, request *EnableAccessPoint) (*EnableAccessPointResponse, error) {
	response := new(EnableAccessPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/EnableAccessPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) EnableAccessPoint(request *EnableAccessPoint) (*EnableAccessPointResponse, error) {
	return service.EnableAccessPointContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) DisableAccessPointContext(ctx context.Context, request *DisableAccessPoint) (*DisableAccessPointResponse, error) {
	response := new(DisableAccessPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/DisableAccessPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) DisableAccessPoint(request *DisableAccessPoint) (*DisableAccessPointResponse, error) {
	return service.DisableAccessPointContext(
		context.Background(),
		request,
	)
}

func (service *pACSPort) ExternalAuthorizationContext(ctx context.Context, request *ExternalAuthorization) (*ExternalAuthorizationResponse, error) {
	response := new(ExternalAuthorizationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accesscontrol/wsdl/ExternalAuthorization", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pACSPort) ExternalAuthorization(request *ExternalAuthorization) (*ExternalAuthorizationResponse, error) {
	return service.ExternalAuthorizationContext(
		context.Background(),
		request,
	)
}

// AnyURI type
type AnyURI string

// Duration type
type Duration string

// QName type
type QName string

// NCName type
type NCName string

// NonNegativeInteger type
type NonNegativeInteger int64

// NonPositiveInteger type
type NonPositiveInteger int64

// AnySimpleType type
type AnySimpleType string

// String type
type String string
