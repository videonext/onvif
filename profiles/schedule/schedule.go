package schedule

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Capabilities ServiceCapabilities

type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetServiceCapabilitiesResponse"`

	//
	// The capability response message contains the requested schedule service
	// capabilities using a hierarchical XML capability structure.
	//
	Capabilities ServiceCapabilities `xml:"http://www.onvif.org/ver10/schedule/wsdl Capabilities,omitempty"`
}

type GetScheduleState struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleState"`

	//
	// Token of schedule instance to get ScheduleState.
	//
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type GetScheduleStateResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleStateResponse"`

	//
	// ScheduleState item.
	//
	ScheduleState ScheduleState `xml:"http://www.onvif.org/ver10/schedule/wsdl ScheduleState,omitempty"`
}

type GetScheduleInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleInfo"`

	//
	// Tokens of ScheduleInfo items to get.
	//
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type GetScheduleInfoResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleInfoResponse"`

	//
	// List of ScheduleInfo items.
	//
	ScheduleInfo ScheduleInfo `xml:"http://www.onvif.org/ver10/schedule/wsdl ScheduleInfo,omitempty"`
}

type GetScheduleInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleInfoList"`

	//
	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is
	// determined by the device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	//
	// Start returning entries from this start reference.
	// If not specified, entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl StartReference,omitempty"`
}

type GetScheduleInfoListResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleInfoListResponse"`

	//
	// StartReference to use in next call to get the following items.
	// If absent, no more items to get.
	//
	NextStartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl NextStartReference,omitempty"`

	//
	// List of ScheduleInfo items.
	//
	ScheduleInfo ScheduleInfo `xml:"http://www.onvif.org/ver10/schedule/wsdl ScheduleInfo,omitempty"`
}

type GetSchedules struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSchedules"`

	// Tokens of Schedule items to get
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type GetSchedulesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSchedulesResponse"`

	// List of schedule items.
	Schedule Schedule `xml:"http://www.onvif.org/ver10/schedule/wsdl Schedule,omitempty"`
}

type GetScheduleList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleList"`

	//
	// Maximum number of entries to return.
	// If not specified, less than one or higher than what the device supports,
	// the number of items is determined by the device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	//
	// Start returning entries from this start reference.
	// If not specified, entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl StartReference,omitempty"`
}

type GetScheduleListResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetScheduleListResponse"`

	//
	// StartReference to use in next call to get the following items.
	// If absent, no more items to get.
	//
	NextStartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl NextStartReference,omitempty"`

	// List of Schedule items.
	Schedule Schedule `xml:"http://www.onvif.org/ver10/schedule/wsdl Schedule,omitempty"`
}

type CreateSchedule struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl CreateSchedule"`

	// The Schedule to create
	Schedule Schedule `xml:"http://www.onvif.org/ver10/schedule/wsdl Schedule,omitempty"`
}

type CreateScheduleResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl CreateScheduleResponse"`

	// The token of created Schedule
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type SetSchedule struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl SetSchedule"`

	// The Schedule to modify/create
	Schedule Schedule `xml:"http://www.onvif.org/ver10/schedule/wsdl Schedule,omitempty"`
}

type SetScheduleResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl SetScheduleResponse"`
}

type ModifySchedule struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl ModifySchedule"`

	// The Schedule to modify/update
	Schedule Schedule `xml:"http://www.onvif.org/ver10/schedule/wsdl Schedule,omitempty"`
}

type ModifyScheduleResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl ModifyScheduleResponse"`
}

type DeleteSchedule struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl DeleteSchedule"`

	// The token of the schedule to delete.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type DeleteScheduleResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl DeleteScheduleResponse"`
}

type GetSpecialDayGroupInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupInfo"`

	// Tokens of SpecialDayGroupInfo items to get.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type GetSpecialDayGroupInfoResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupInfoResponse"`

	// List of SpecialDayGroupInfo items.
	SpecialDayGroupInfo SpecialDayGroupInfo `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroupInfo,omitempty"`
}

type GetSpecialDayGroupInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupInfoList"`

	//
	// Maximum number of entries to return. If not specified, less than
	// one or higher than what the device supports, the number
	// of items is determined by the device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	//
	// Start returning entries from this start reference.
	// If not specified, entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl StartReference,omitempty"`
}

type GetSpecialDayGroupInfoListResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupInfoListResponse"`

	//
	// StartReference to use in next call to get the following items.
	// If absent, no more items to get.
	//
	NextStartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl NextStartReference,omitempty"`

	// List of SpecialDayGroupInfo items.
	SpecialDayGroupInfo SpecialDayGroupInfo `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroupInfo,omitempty"`
}

type GetSpecialDayGroups struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroups"`

	// Tokens of the SpecialDayGroup items to get
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type GetSpecialDayGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupsResponse"`

	// List of SpecialDayGroup items.
	SpecialDayGroup SpecialDayGroup `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroup,omitempty"`
}

type GetSpecialDayGroupList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupList"`

	//
	// Maximum number of entries to return. If not specified, less than
	// one or higher than what the device supports, the number of
	// items is determined by the device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	//
	// Start returning entries from this start reference.
	// If not specified, entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl StartReference,omitempty"`
}

type GetSpecialDayGroupListResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl GetSpecialDayGroupListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"http://www.onvif.org/ver10/schedule/wsdl NextStartReference,omitempty"`

	// List of SpecialDayGroup items.
	SpecialDayGroup SpecialDayGroup `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroup,omitempty"`
}

type CreateSpecialDayGroup struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl CreateSpecialDayGroup"`

	//
	// The special day group to create.
	//
	SpecialDayGroup SpecialDayGroup `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroup,omitempty"`
}

type CreateSpecialDayGroupResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl CreateSpecialDayGroupResponse"`

	// The token of created special day group.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type SetSpecialDayGroup struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl SetSpecialDayGroup"`

	// The SpecialDayGroup to modify/create
	SpecialDayGroup SpecialDayGroup `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroup,omitempty"`
}

type SetSpecialDayGroupResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl SetSpecialDayGroupResponse"`
}

type ModifySpecialDayGroup struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl ModifySpecialDayGroup"`

	//
	// The special day group to modify/update.
	//
	SpecialDayGroup SpecialDayGroup `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDayGroup,omitempty"`
}

type ModifySpecialDayGroupResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl ModifySpecialDayGroupResponse"`
}

type DeleteSpecialDayGroup struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl DeleteSpecialDayGroup"`

	//
	// The token of the special day group item to delete.
	//
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl Token,omitempty"`
}

type DeleteSpecialDayGroupResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl DeleteSpecialDayGroupResponse"`
}

type ServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schedule/wsdl Capabilities"`

	//
	// The maximum number of entries returned by a single Get<Entity>List or
	// Get<Entity> request. The device shall never return more than this number
	// of entities in a single response.
	//

	MaxLimit PositiveInteger `xml:"http://www.onvif.org/ver10/schedule/wsdl MaxLimit,attr,omitempty"`

	//
	// Indicates the maximum number of schedules the device supports.
	// The device shall support at least one schedule.
	//

	MaxSchedules PositiveInteger `xml:"http://www.onvif.org/ver10/schedule/wsdl MaxSchedules,attr,omitempty"`

	//
	// Indicates the maximum number of time periods per day the device supports
	// in a schedule including special days schedule. The device shall support
	// at least one time period per day.
	//

	MaxTimePeriodsPerDay PositiveInteger `xml:"http://www.onvif.org/ver10/schedule/wsdl MaxTimePeriodsPerDay,attr,omitempty"`

	//
	// Indicates the maximum number of special day group entities the device supports.
	// The device shall support at least one ‘SpecialDayGroup’ entity.
	//

	MaxSpecialDayGroups PositiveInteger `xml:"http://www.onvif.org/ver10/schedule/wsdl MaxSpecialDayGroups,attr,omitempty"`

	//
	// Indicates the maximum number of days per ‘SpecialDayGroup’ entity the device
	// supports. The device shall support at least one day per ‘SpecialDayGroup’ entity.
	//

	MaxDaysInSpecialDayGroup PositiveInteger `xml:"http://www.onvif.org/ver10/schedule/wsdl MaxDaysInSpecialDayGroup,attr,omitempty"`

	//
	// Indicates the maximum number of ‘SpecialDaysSchedule’ entities referred by a
	// schedule that the device supports.
	//

	MaxSpecialDaysSchedules PositiveInteger `xml:"http://www.onvif.org/ver10/schedule/wsdl MaxSpecialDaysSchedules,attr,omitempty"`

	//
	// For schedules:
	// If this capability is supported, then all iCalendar recurrence types shall
	// be supported by the device. The device shall also support the start and end dates (or
	// iCalendar occurrence count) in recurring events (see iCalendar examples in section 3).
	// If this capability is not supported, then only the weekly iCalendar recurrence
	// type shall be supported. Non-recurring events and other recurring types are
	// not supported. The device shall only accept a start date with the year ‘1970’
	// (the month and day is needed to reflect the week day of the recurrence)
	// and will not accept an occurrence count (or iCalendar until date) in recurring events.
	// For special days (only applicable if SpecialDaysSupported is set to true):
	// If this capability is supported, then all iCalendar recurrence types shall
	// be supported by the device. The device shall also support the start and
	// end dates (or occurrence count) in recurring events.
	// If this capability is not supported, then only non-recurring special days are supported.
	//

	ExtendedRecurrenceSupported bool `xml:"http://www.onvif.org/ver10/schedule/wsdl ExtendedRecurrenceSupported,attr,omitempty"`

	//
	// If this capability is supported, then the device shall support special days.
	//

	SpecialDaysSupported bool `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDaysSupported,attr,omitempty"`

	//
	// If this capability is set to true, the device shall implement the
	// GetScheduleState command, and shall notify subscribing clients whenever
	// schedules become active or inactive.
	//

	StateReportingSupported bool `xml:"http://www.onvif.org/ver10/schedule/wsdl StateReportingSupported,attr,omitempty"`

	//
	// Indicates that the client is allowed to supply the token when creating schedules and special day groups.
	// To enable the use of the commands SetSchedule and SetSpecialDayGroup, the value must be set to true.
	//

	ClientSuppliedTokenSupported bool `xml:"http://www.onvif.org/ver10/schedule/wsdl ClientSuppliedTokenSupported,attr,omitempty"`
}

type ScheduleInfo struct {
	*DataEntity

	//
	// A user readable name. It shall be up to 64 characters.
	//
	Name Name `xml:"http://www.onvif.org/ver10/schedule/wsdl Name,omitempty"`

	//
	// User readable description for the schedule. It shall be up to 1024 characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/schedule/wsdl Description,omitempty"`
}

type Schedule struct {
	*ScheduleInfo

	//
	// An iCalendar structure that defines a number of events. Events
	// can be recurring or non-recurring. The events can, for instance,
	// be used to control when a camera should record or when a facility
	// is accessible. Some devices might not be able to fully support
	// all the features of iCalendar. Setting the service capability
	// ExtendedRecurrenceSupported to false will enable more devices
	// to be ONVIF compliant. Is of type string (but contains an iCalendar structure).
	//
	Standard string `xml:"http://www.onvif.org/ver10/schedule/wsdl Standard,omitempty"`

	//
	// For devices that are not able to support all the features of iCalendar,
	// supporting special days is essential. Each SpecialDaysSchedule
	// instance defines an alternate set of time periods that overrides
	// the regular schedule for a specified list of special days.
	// Is of type SpecialDaysSchedule.
	//
	SpecialDays SpecialDaysSchedule `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDays,omitempty"`

	Extension ScheduleExtension `xml:"http://www.onvif.org/ver10/schedule/wsdl Extension,omitempty"`
}

type ScheduleExtension struct {
}

type SpecialDaysSchedule struct {

	//
	// Indicates the list of special days in a schedule.
	//
	GroupToken ReferenceToken `xml:"http://www.onvif.org/ver10/schedule/wsdl GroupToken,omitempty"`

	//
	// Indicates the alternate time periods for the list of special days
	// (overrides the regular schedule). For example, the regular schedule indicates
	// that it is active from 8AM to 5PM on Mondays. However, this particular
	// Monday is a special day, and the alternate time periods state that the
	// schedule is active from 9 AM to 11 AM and 1 PM to 4 PM.
	// If no time periods are defined, then no access is allowed.
	// Is of type TimePeriod.
	//
	TimeRange TimePeriod `xml:"http://www.onvif.org/ver10/schedule/wsdl TimeRange,omitempty"`

	Extension SpecialDaysScheduleExtension `xml:"http://www.onvif.org/ver10/schedule/wsdl Extension,omitempty"`
}

type SpecialDaysScheduleExtension struct {
}

type ScheduleState struct {

	//
	// Indicates that the current time is within the boundaries of the schedule
	// or its special days schedules’ time periods. For example, if this
	// schedule is being used for triggering automatic recording on a video source,
	// the Active flag will be true when the schedule-based recording is supposed to record.
	//
	Active bool `xml:"http://www.onvif.org/ver10/schedule/wsdl Active,omitempty"`

	//
	// Indicates that the current time is within the boundaries of its special
	// days schedules’ time periods. For example, if this schedule is being used
	// for recording at a lower frame rate on a video source during special days,
	// the SpecialDay flag will be true. If special days are not supported by the device,
	// this field may be omitted and interpreted as false by the client.
	//
	SpecialDay bool `xml:"http://www.onvif.org/ver10/schedule/wsdl SpecialDay,omitempty"`

	Extension ScheduleStateExtension `xml:"http://www.onvif.org/ver10/schedule/wsdl Extension,omitempty"`
}

type ScheduleStateExtension struct {
}

type TimePeriod struct {

	//
	// Indicates the start time.
	//
	From string `xml:"http://www.onvif.org/ver10/schema From,omitempty"`

	//
	// Indicates the end time. Is optional, if omitted, the period ends at midnight.
	// The end time is exclusive, meaning that that exact moment in time is not
	// part of the period. To determine if a moment in time (t) is part of a time period,
	// the formula StartTime ≤ t < EndTime is used.
	//
	Until string `xml:"http://www.onvif.org/ver10/schema Until,omitempty"`

	Extension TimePeriodExtension `xml:"http://www.onvif.org/ver10/schedule/wsdl Extension,omitempty"`
}

type TimePeriodExtension struct {
}

type SpecialDayGroupInfo struct {
	*DataEntity

	//
	// User readable name. It shall be up to 64 characters.
	//
	Name Name `xml:"http://www.onvif.org/ver10/schedule/wsdl Name,omitempty"`

	//
	// User readable description for the special days. It shall be up to 1024
	// characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/schedule/wsdl Description,omitempty"`
}

type SpecialDayGroup struct {
	*SpecialDayGroupInfo

	//
	// An iCalendar structure that contains a group of special days.
	// Is of type string (containing an iCalendar structure).
	//
	Days string `xml:"http://www.onvif.org/ver10/schedule/wsdl Days,omitempty"`

	Extension SpecialDayGroupExtension `xml:"http://www.onvif.org/ver10/schedule/wsdl Extension,omitempty"`
}

type SpecialDayGroupExtension struct {
}

// Type used to reference logical and physical entities.

type ReferenceToken string

// Type used for names of logical and physical entities.

type Name string

// Description is optional and the maximum length is device specific.
// If the length is more than maximum length, it is silently chopped to the maximum length
// supported by the device/service (which may be 0).
//

type Description string

// Type used to represent the numbers from 1 ,2 , 3,...

type PositiveInteger uint32

type DataEntity struct {

	// A service-unique identifier of the item.

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

// Removed Attribute by fixgen.py

type SchedulePort interface {

	/*
		This operation returns the capabilities of the schedule service.
	*/
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/*
		This operation requests the ScheduleState for the schedule instance specified by the given token.
	*/
	GetScheduleState(request *GetScheduleState) (*GetScheduleStateResponse, error)

	GetScheduleStateContext(ctx context.Context, request *GetScheduleState) (*GetScheduleStateResponse, error)

	/*
		This method returns a list of schedule info items, specified in the request.
		Only found schedules shall be returned, i.e., the returned numbers of elements can
		differ from the requested element.
		The device shall ignore tokens it cannot resolve and shall return an empty list if
		there are no items matching the specified tokens.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetScheduleInfo(request *GetScheduleInfo) (*GetScheduleInfoResponse, error)

	GetScheduleInfoContext(ctx context.Context, request *GetScheduleInfo) (*GetScheduleInfoResponse, error)

	/*
		This operation requests a list of all of ScheduleInfo items provided by the device.
		A call to this method shall return a StartReference when not all data is returned
		and more data is available. The reference shall be valid for retrieving the next set of data.
		Please refer Access Control Service Specification for more details.
		The number of items returned shall not be greater the Limit parameter.
	*/
	GetScheduleInfoList(request *GetScheduleInfoList) (*GetScheduleInfoListResponse, error)

	GetScheduleInfoListContext(ctx context.Context, request *GetScheduleInfoList) (*GetScheduleInfoListResponse, error)

	/*
		This operation returns the specified schedule item matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list
		if there are no items matching the specified tokens.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned
	*/
	GetSchedules(request *GetSchedules) (*GetSchedulesResponse, error)

	GetSchedulesContext(ctx context.Context, request *GetSchedules) (*GetSchedulesResponse, error)

	/*
		This operation requests a list of all of Schedule items provided by the device.
		A call to this method shall return a StartReference when not all data is returned
		and more data is available. The reference shall be valid for retrieving the next set of data.
		Please refer Access Control Service Specification for more details.
		The number of items returned shall not be greater the Limit parameter.
	*/
	GetScheduleList(request *GetScheduleList) (*GetScheduleListResponse, error)

	GetScheduleListContext(ctx context.Context, request *GetScheduleList) (*GetScheduleListResponse, error)

	/*
		This operation creates the specified schedule. The token field of the schedule structure
		shall be empty, the device shall allocate a token for the schedule. The allocated token
		shall be returned in the response. If the client sends any value in the token field,
		the device shall return InvalidArgVal as generic fault code.
	*/
	CreateSchedule(request *CreateSchedule) (*CreateScheduleResponse, error)

	CreateScheduleContext(ctx context.Context, request *CreateSchedule) (*CreateScheduleResponse, error)

	/*
		This operation modifies or creates the specified schedule.
	*/
	SetSchedule(request *SetSchedule) (*SetScheduleResponse, error)

	SetScheduleContext(ctx context.Context, request *SetSchedule) (*SetScheduleResponse, error)

	/*
		This operation modifies or updates the specified schedule.
	*/
	ModifySchedule(request *ModifySchedule) (*ModifyScheduleResponse, error)

	ModifyScheduleContext(ctx context.Context, request *ModifySchedule) (*ModifyScheduleResponse, error)

	/*
		This operation will delete the specified schedule.
		If it is associated with one or more entities some devices may not be able to delete the schedule,
		and consequently a ReferenceInUse fault shall be generated.
	*/
	DeleteSchedule(request *DeleteSchedule) (*DeleteScheduleResponse, error)

	DeleteScheduleContext(ctx context.Context, request *DeleteSchedule) (*DeleteScheduleResponse, error)

	/*
		This operation requests a list of SpecialDayGroupInfo items matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list if
		there are no items matching specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetSpecialDayGroupInfo(request *GetSpecialDayGroupInfo) (*GetSpecialDayGroupInfoResponse, error)

	GetSpecialDayGroupInfoContext(ctx context.Context, request *GetSpecialDayGroupInfo) (*GetSpecialDayGroupInfoResponse, error)

	/*
		This operation requests a list of all of SpecialDayGroupInfo items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and
		more data is available. The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than Limit parameter.
	*/
	GetSpecialDayGroupInfoList(request *GetSpecialDayGroupInfoList) (*GetSpecialDayGroupInfoListResponse, error)

	GetSpecialDayGroupInfoListContext(ctx context.Context, request *GetSpecialDayGroupInfoList) (*GetSpecialDayGroupInfoListResponse, error)

	/*
		This operation returns the specified special day group item matching the given token.
	*/
	GetSpecialDayGroups(request *GetSpecialDayGroups) (*GetSpecialDayGroupsResponse, error)

	GetSpecialDayGroupsContext(ctx context.Context, request *GetSpecialDayGroups) (*GetSpecialDayGroupsResponse, error)

	/*
		This operation requests a list of all of SpecialDayGroupList items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and
		more data is available. The reference shall be valid for retrieving the next set of data.
		Please refer Access Control Service Specification for more details.
		The number of items returned shall not be greater the Limit parameter.
	*/
	GetSpecialDayGroupList(request *GetSpecialDayGroupList) (*GetSpecialDayGroupListResponse, error)

	GetSpecialDayGroupListContext(ctx context.Context, request *GetSpecialDayGroupList) (*GetSpecialDayGroupListResponse, error)

	/*
		This operation creates the specified special day group. The token field of the
		SpecialDayGroup structure shall be empty, the device shall allocate a token for the
		special day group. The allocated token shall be returned in the response.
		If there is any value in the token field, the device shall return InvalidArgVal as generic fault code.
	*/
	CreateSpecialDayGroup(request *CreateSpecialDayGroup) (*CreateSpecialDayGroupResponse, error)

	CreateSpecialDayGroupContext(ctx context.Context, request *CreateSpecialDayGroup) (*CreateSpecialDayGroupResponse, error)

	/*
		This operation modifies or creates the specified special day group.
	*/
	SetSpecialDayGroup(request *SetSpecialDayGroup) (*SetSpecialDayGroupResponse, error)

	SetSpecialDayGroupContext(ctx context.Context, request *SetSpecialDayGroup) (*SetSpecialDayGroupResponse, error)

	/*
		This operation updates the specified special day group.
	*/
	ModifySpecialDayGroup(request *ModifySpecialDayGroup) (*ModifySpecialDayGroupResponse, error)

	ModifySpecialDayGroupContext(ctx context.Context, request *ModifySpecialDayGroup) (*ModifySpecialDayGroupResponse, error)

	/*
		This method deletes the specified special day group.
		If it is associated with one or more schedules some devices may not be able to delete
		the special day group, and consequently a ReferenceInUse fault must be generated.
	*/
	DeleteSpecialDayGroup(request *DeleteSpecialDayGroup) (*DeleteSpecialDayGroupResponse, error)

	DeleteSpecialDayGroupContext(ctx context.Context, request *DeleteSpecialDayGroup) (*DeleteSpecialDayGroupResponse, error)
}

type schedulePort struct {
	client *soap.Client
	xaddr  string
}

func NewSchedulePort(client *soap.Client, xaddr string) SchedulePort {
	return &schedulePort{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *schedulePort) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetScheduleStateContext(ctx context.Context, request *GetScheduleState) (*GetScheduleStateResponse, error) {
	response := new(GetScheduleStateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetScheduleState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetScheduleState(request *GetScheduleState) (*GetScheduleStateResponse, error) {
	return service.GetScheduleStateContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetScheduleInfoContext(ctx context.Context, request *GetScheduleInfo) (*GetScheduleInfoResponse, error) {
	response := new(GetScheduleInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetScheduleInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetScheduleInfo(request *GetScheduleInfo) (*GetScheduleInfoResponse, error) {
	return service.GetScheduleInfoContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetScheduleInfoListContext(ctx context.Context, request *GetScheduleInfoList) (*GetScheduleInfoListResponse, error) {
	response := new(GetScheduleInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetScheduleInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetScheduleInfoList(request *GetScheduleInfoList) (*GetScheduleInfoListResponse, error) {
	return service.GetScheduleInfoListContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetSchedulesContext(ctx context.Context, request *GetSchedules) (*GetSchedulesResponse, error) {
	response := new(GetSchedulesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetSchedules", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetSchedules(request *GetSchedules) (*GetSchedulesResponse, error) {
	return service.GetSchedulesContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetScheduleListContext(ctx context.Context, request *GetScheduleList) (*GetScheduleListResponse, error) {
	response := new(GetScheduleListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetScheduleList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetScheduleList(request *GetScheduleList) (*GetScheduleListResponse, error) {
	return service.GetScheduleListContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) CreateScheduleContext(ctx context.Context, request *CreateSchedule) (*CreateScheduleResponse, error) {
	response := new(CreateScheduleResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/CreateSchedule", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) CreateSchedule(request *CreateSchedule) (*CreateScheduleResponse, error) {
	return service.CreateScheduleContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) SetScheduleContext(ctx context.Context, request *SetSchedule) (*SetScheduleResponse, error) {
	response := new(SetScheduleResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/SetSchedule", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) SetSchedule(request *SetSchedule) (*SetScheduleResponse, error) {
	return service.SetScheduleContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) ModifyScheduleContext(ctx context.Context, request *ModifySchedule) (*ModifyScheduleResponse, error) {
	response := new(ModifyScheduleResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/ModifySchedule", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) ModifySchedule(request *ModifySchedule) (*ModifyScheduleResponse, error) {
	return service.ModifyScheduleContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) DeleteScheduleContext(ctx context.Context, request *DeleteSchedule) (*DeleteScheduleResponse, error) {
	response := new(DeleteScheduleResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/DeleteSchedule", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) DeleteSchedule(request *DeleteSchedule) (*DeleteScheduleResponse, error) {
	return service.DeleteScheduleContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetSpecialDayGroupInfoContext(ctx context.Context, request *GetSpecialDayGroupInfo) (*GetSpecialDayGroupInfoResponse, error) {
	response := new(GetSpecialDayGroupInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetSpecialDayGroupInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetSpecialDayGroupInfo(request *GetSpecialDayGroupInfo) (*GetSpecialDayGroupInfoResponse, error) {
	return service.GetSpecialDayGroupInfoContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetSpecialDayGroupInfoListContext(ctx context.Context, request *GetSpecialDayGroupInfoList) (*GetSpecialDayGroupInfoListResponse, error) {
	response := new(GetSpecialDayGroupInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetSpecialDayGroupInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetSpecialDayGroupInfoList(request *GetSpecialDayGroupInfoList) (*GetSpecialDayGroupInfoListResponse, error) {
	return service.GetSpecialDayGroupInfoListContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetSpecialDayGroupsContext(ctx context.Context, request *GetSpecialDayGroups) (*GetSpecialDayGroupsResponse, error) {
	response := new(GetSpecialDayGroupsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetSpecialDayGroups", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetSpecialDayGroups(request *GetSpecialDayGroups) (*GetSpecialDayGroupsResponse, error) {
	return service.GetSpecialDayGroupsContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) GetSpecialDayGroupListContext(ctx context.Context, request *GetSpecialDayGroupList) (*GetSpecialDayGroupListResponse, error) {
	response := new(GetSpecialDayGroupListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/GetSpecialDayGroupList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) GetSpecialDayGroupList(request *GetSpecialDayGroupList) (*GetSpecialDayGroupListResponse, error) {
	return service.GetSpecialDayGroupListContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) CreateSpecialDayGroupContext(ctx context.Context, request *CreateSpecialDayGroup) (*CreateSpecialDayGroupResponse, error) {
	response := new(CreateSpecialDayGroupResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/CreateSpecialDayGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) CreateSpecialDayGroup(request *CreateSpecialDayGroup) (*CreateSpecialDayGroupResponse, error) {
	return service.CreateSpecialDayGroupContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) SetSpecialDayGroupContext(ctx context.Context, request *SetSpecialDayGroup) (*SetSpecialDayGroupResponse, error) {
	response := new(SetSpecialDayGroupResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/SetSpecialDayGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) SetSpecialDayGroup(request *SetSpecialDayGroup) (*SetSpecialDayGroupResponse, error) {
	return service.SetSpecialDayGroupContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) ModifySpecialDayGroupContext(ctx context.Context, request *ModifySpecialDayGroup) (*ModifySpecialDayGroupResponse, error) {
	response := new(ModifySpecialDayGroupResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/ModifySpecialDayGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) ModifySpecialDayGroup(request *ModifySpecialDayGroup) (*ModifySpecialDayGroupResponse, error) {
	return service.ModifySpecialDayGroupContext(
		context.Background(),
		request,
	)
}

func (service *schedulePort) DeleteSpecialDayGroupContext(ctx context.Context, request *DeleteSpecialDayGroup) (*DeleteSpecialDayGroupResponse, error) {
	response := new(DeleteSpecialDayGroupResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/schedule/wsdl/DeleteSpecialDayGroup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *schedulePort) DeleteSpecialDayGroup(request *DeleteSpecialDayGroup) (*DeleteSpecialDayGroupResponse, error) {
	return service.DeleteSpecialDayGroupContext(
		context.Background(),
		request,
	)
}

type AnyURI string
type Duration string
type QName string
type NCName string
type NonNegativeInteger int64
type NonPositiveInteger int64
type AnySimpleType string
type String string
