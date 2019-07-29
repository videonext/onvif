package accessrules

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// Capabilities type
type Capabilities ServiceCapabilities

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capability response message contains the requested access rules
	// service capabilities using a hierarchical XML capability structure.
	//
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty"`
}

// GetAccessProfileInfo type
type GetAccessProfileInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl GetAccessProfileInfo"`

	// Tokens of AccessProfileInfo items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/accessrules/wsdl Token,omitempty"`
}

// GetAccessProfileInfoResponse type
type GetAccessProfileInfoResponse struct {
	XMLName xml.Name `xml:"GetAccessProfileInfoResponse"`

	// List of AccessProfileInfo items.
	AccessProfileInfo []AccessProfileInfo `xml:"AccessProfileInfo,omitempty"`
}

// GetAccessProfileInfoList type
type GetAccessProfileInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl GetAccessProfileInfoList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/accessrules/wsdl StartReference,omitempty"`
}

// GetAccessProfileInfoListResponse type
type GetAccessProfileInfoListResponse struct {
	XMLName xml.Name `xml:"GetAccessProfileInfoListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of AccessProfileInfo items.
	AccessProfileInfo []AccessProfileInfo `xml:"AccessProfileInfo,omitempty"`
}

// GetAccessProfiles type
type GetAccessProfiles struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl GetAccessProfiles"`

	// Tokens of AccessProfile items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/accessrules/wsdl Token,omitempty"`
}

// GetAccessProfilesResponse type
type GetAccessProfilesResponse struct {
	XMLName xml.Name `xml:"GetAccessProfilesResponse"`

	// List of Access Profile items.
	AccessProfile []AccessProfile `xml:"AccessProfile,omitempty"`
}

// GetAccessProfileList type
type GetAccessProfileList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl GetAccessProfileList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/accessrules/wsdl StartReference,omitempty"`
}

// GetAccessProfileListResponse type
type GetAccessProfileListResponse struct {
	XMLName xml.Name `xml:"GetAccessProfileListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of Access Profile items.
	AccessProfile []AccessProfile `xml:"AccessProfile,omitempty"`
}

// CreateAccessProfile type
type CreateAccessProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl CreateAccessProfile"`

	// The AccessProfile to create.
	AccessProfile AccessProfile `xml:"http://www.onvif.org/ver10/accessrules/wsdl AccessProfile,omitempty"`
}

// CreateAccessProfileResponse type
type CreateAccessProfileResponse struct {
	XMLName xml.Name `xml:"CreateAccessProfileResponse"`

	// The Token of created AccessProfile.
	Token ReferenceToken `xml:"Token,omitempty"`
}

// ModifyAccessProfile type
type ModifyAccessProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl ModifyAccessProfile"`

	// The details of Access Profile
	AccessProfile AccessProfile `xml:"http://www.onvif.org/ver10/accessrules/wsdl AccessProfile,omitempty"`
}

// ModifyAccessProfileResponse type
type ModifyAccessProfileResponse struct {
	XMLName xml.Name `xml:"ModifyAccessProfileResponse"`
}

// SetAccessProfile type
type SetAccessProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl SetAccessProfile"`

	// The AccessProfile item to create or modify
	AccessProfile AccessProfile `xml:"http://www.onvif.org/ver10/accessrules/wsdl AccessProfile,omitempty"`
}

// SetAccessProfileResponse type
type SetAccessProfileResponse struct {
	XMLName xml.Name `xml:"SetAccessProfileResponse"`
}

// DeleteAccessProfile type
type DeleteAccessProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl DeleteAccessProfile"`

	// The token of the access profile to delete.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/accessrules/wsdl Token,omitempty"`
}

// DeleteAccessProfileResponse type
type DeleteAccessProfileResponse struct {
	XMLName xml.Name `xml:"DeleteAccessProfileResponse"`
}

// ServiceCapabilities type
type ServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl Capabilities"`

	//
	// The maximum number of entries returned by a single Get<Entity>List or Get<Entity>
	// request. The device shall never return more than this number of entities in a single
	// response.
	//

	MaxLimit uint32 `xml:"http://www.onvif.org/ver10/schema MaxLimit,attr,omitempty"`

	//
	// Indicates the maximum number of access profiles supported by the device.
	//

	MaxAccessProfiles uint32 `xml:"http://www.onvif.org/ver10/schema MaxAccessProfiles,attr,omitempty"`

	//
	// Indicates the maximum number of access policies per access profile supported by the device.
	//

	MaxAccessPoliciesPerAccessProfile uint32 `xml:"http://www.onvif.org/ver10/schema MaxAccessPoliciesPerAccessProfile,attr,omitempty"`

	//
	// Indicates whether or not several access policies can refer to the same access point in an
	// access profile.
	//

	MultipleSchedulesPerAccessPointSupported bool `xml:"http://www.onvif.org/ver10/accessrules/wsdl MultipleSchedulesPerAccessPointSupported,attr,omitempty"`

	//
	// Indicates that the client is allowed to supply the token when creating access profiles. To
	// enable the use of the command SetAccessProfile, the value must be set to true.
	//

	ClientSuppliedTokenSupported bool `xml:"http://www.onvif.org/ver10/accessrules/wsdl ClientSuppliedTokenSupported,attr,omitempty"`
}

// AccessPolicy type
type AccessPolicy struct {

	// Reference to the schedule used by the access policy.
	ScheduleToken ReferenceToken `xml:"http://www.onvif.org/ver10/accessrules/wsdl ScheduleToken,omitempty"`

	//
	// Reference to the entity used by the rule engine, the entity type may be specified by the
	// optional EntityType field explained below but is typically an access point.
	//
	Entity ReferenceToken `xml:"http://www.onvif.org/ver10/accessrules/wsdl Entity,omitempty"`

	//
	// Optional entity type; if missing, an access point type as defined by the ONVIF Access
	// Control Service Specification should be assumed. This can also be represented by the
	// QName value	“tac:AccessPoint” where tac is the namespace of ONVIF Access Control
	// Service Specification. This field is provided for future extensions; it will allow an
	// access policy being	extended to cover entity types other than access points as well.
	//
	EntityType QName `xml:"http://www.onvif.org/ver10/accessrules/wsdl EntityType,omitempty"`

	Extension AccessPolicyExtension `xml:"http://www.onvif.org/ver10/accessrules/wsdl Extension,omitempty"`
}

// AccessPolicyExtension type
type AccessPolicyExtension struct {
}

// AccessProfileInfo type
type AccessProfileInfo struct {
	*DataEntity

	// A user readable name. It shall be up to 64 characters.
	//
	Name Name `xml:"http://www.onvif.org/ver10/accessrules/wsdl Name,omitempty"`

	// User readable description for the access profile. It shall be up
	// to 1024 characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/accessrules/wsdl Description,omitempty"`
}

// AccessProfile type
type AccessProfile struct {
	*AccessProfileInfo

	// A list of access policy structures, where each access policy
	// defines during which schedule an access point can be accessed.
	//
	AccessPolicy []AccessPolicy `xml:"http://www.onvif.org/ver10/accessrules/wsdl AccessPolicy,omitempty"`

	Extension AccessProfileExtension `xml:"http://www.onvif.org/ver10/accessrules/wsdl Extension,omitempty"`
}

// AccessProfileExtension type
type AccessProfileExtension struct {
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

// AccessRulesPort type
type AccessRulesPort interface {

	/* This operation returns the capabilities of the access rules service.
	 */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/*
		This operation requests a list of AccessProfileInfo items matching the given tokens. The device shall
		ignore tokens it cannot resolve and shall return an empty list if there are no items matching the
		specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetAccessProfileInfo(request *GetAccessProfileInfo) (*GetAccessProfileInfoResponse, error)

	GetAccessProfileInfoContext(ctx context.Context, request *GetAccessProfileInfo) (*GetAccessProfileInfoResponse, error)

	/*
		This operation requests a list of all of AccessProfileInfo items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and more data is
		available. The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetAccessProfileInfoList(request *GetAccessProfileInfoList) (*GetAccessProfileInfoListResponse, error)

	GetAccessProfileInfoListContext(ctx context.Context, request *GetAccessProfileInfoList) (*GetAccessProfileInfoListResponse, error)

	/*
		This operation returns the specified access profile item matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list if there are no items
		matching specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetAccessProfiles(request *GetAccessProfiles) (*GetAccessProfilesResponse, error)

	GetAccessProfilesContext(ctx context.Context, request *GetAccessProfiles) (*GetAccessProfilesResponse, error)

	/*
		This operation requests a list of all of access profile items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and more data is
		available. The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetAccessProfileList(request *GetAccessProfileList) (*GetAccessProfileListResponse, error)

	GetAccessProfileListContext(ctx context.Context, request *GetAccessProfileList) (*GetAccessProfileListResponse, error)

	/*
		This operation creates the specified access profile in the device. The token field of the access profile shall be
		empty, the service shall allocate a token for the access profile. The allocated token shall be returned
		in the response. If the client sends any value in the token field, the device shall return InvalidArgVal
		as generic fault code.
		In an access profile, if several access policies specifying different schedules for the same access
		point will result in a union of the schedules.
	*/
	CreateAccessProfile(request *CreateAccessProfile) (*CreateAccessProfileResponse, error)

	CreateAccessProfileContext(ctx context.Context, request *CreateAccessProfile) (*CreateAccessProfileResponse, error)

	/*
		This operation will modify the access profile for the specified access profile token. The token of the
		access profile to modify is specified in the token field of the AccessProile structure and shall not
		be empty. All other fields in the structure shall overwrite the fields in the specified access profile.
		If several access policies specifying different schedules for the same access point will result in a
		union of the schedules.
		If the device could not store the access profile information then a fault will be generated.
	*/
	ModifyAccessProfile(request *ModifyAccessProfile) (*ModifyAccessProfileResponse, error)

	ModifyAccessProfileContext(ctx context.Context, request *ModifyAccessProfile) (*ModifyAccessProfileResponse, error)

	/*
		This operation will synchronize an access profile in a client with the device.
		If an access profile with the specified token does not exist in the device, the access profile is
		created. If an access profile with the specified token exists, then the access profile is modified.
		A call to this method takes an access profile structure as input parameter. The token field of the
		access profile must not be empty.
		A device that signals support for the ClientSuppliedTokenSupported capability shall implement this command.
	*/
	SetAccessProfile(request *SetAccessProfile) (*SetAccessProfileResponse, error)

	SetAccessProfileContext(ctx context.Context, request *SetAccessProfile) (*SetAccessProfileResponse, error)

	/*
		This operation will delete the specified access profile.
		If the access profile is deleted, all access policies associated to the access profile will also be
		deleted.
		If it is associated with one or more entities some devices may not be able to delete the access profile,
		and consequently a ReferenceInUse fault shall be generated.
	*/
	DeleteAccessProfile(request *DeleteAccessProfile) (*DeleteAccessProfileResponse, error)

	DeleteAccessProfileContext(ctx context.Context, request *DeleteAccessProfile) (*DeleteAccessProfileResponse, error)
}

// accessRulesPort type
type accessRulesPort struct {
	client *soap.Client
	xaddr  string
}

func NewAccessRulesPort(client *soap.Client, xaddr string) AccessRulesPort {
	return &accessRulesPort{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *accessRulesPort) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) GetAccessProfileInfoContext(ctx context.Context, request *GetAccessProfileInfo) (*GetAccessProfileInfoResponse, error) {
	response := new(GetAccessProfileInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/GetAccessProfileInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) GetAccessProfileInfo(request *GetAccessProfileInfo) (*GetAccessProfileInfoResponse, error) {
	return service.GetAccessProfileInfoContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) GetAccessProfileInfoListContext(ctx context.Context, request *GetAccessProfileInfoList) (*GetAccessProfileInfoListResponse, error) {
	response := new(GetAccessProfileInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/GetAccessProfileInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) GetAccessProfileInfoList(request *GetAccessProfileInfoList) (*GetAccessProfileInfoListResponse, error) {
	return service.GetAccessProfileInfoListContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) GetAccessProfilesContext(ctx context.Context, request *GetAccessProfiles) (*GetAccessProfilesResponse, error) {
	response := new(GetAccessProfilesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/GetAccessProfiles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) GetAccessProfiles(request *GetAccessProfiles) (*GetAccessProfilesResponse, error) {
	return service.GetAccessProfilesContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) GetAccessProfileListContext(ctx context.Context, request *GetAccessProfileList) (*GetAccessProfileListResponse, error) {
	response := new(GetAccessProfileListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/GetAccessProfileList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) GetAccessProfileList(request *GetAccessProfileList) (*GetAccessProfileListResponse, error) {
	return service.GetAccessProfileListContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) CreateAccessProfileContext(ctx context.Context, request *CreateAccessProfile) (*CreateAccessProfileResponse, error) {
	response := new(CreateAccessProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/CreateAccessProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) CreateAccessProfile(request *CreateAccessProfile) (*CreateAccessProfileResponse, error) {
	return service.CreateAccessProfileContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) ModifyAccessProfileContext(ctx context.Context, request *ModifyAccessProfile) (*ModifyAccessProfileResponse, error) {
	response := new(ModifyAccessProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/ModifyAccessProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) ModifyAccessProfile(request *ModifyAccessProfile) (*ModifyAccessProfileResponse, error) {
	return service.ModifyAccessProfileContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) SetAccessProfileContext(ctx context.Context, request *SetAccessProfile) (*SetAccessProfileResponse, error) {
	response := new(SetAccessProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/SetAccessProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) SetAccessProfile(request *SetAccessProfile) (*SetAccessProfileResponse, error) {
	return service.SetAccessProfileContext(
		context.Background(),
		request,
	)
}

func (service *accessRulesPort) DeleteAccessProfileContext(ctx context.Context, request *DeleteAccessProfile) (*DeleteAccessProfileResponse, error) {
	response := new(DeleteAccessProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/accessrules/wsdl/DeleteAccessProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *accessRulesPort) DeleteAccessProfile(request *DeleteAccessProfile) (*DeleteAccessProfileResponse, error) {
	return service.DeleteAccessProfileContext(
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
