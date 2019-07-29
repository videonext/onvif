package doorcontrol

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
// The physical state of a Door.
//

// DoorPhysicalState type
type DoorPhysicalState string

const (

	// Value is currently unknown (possibly due to initialization or monitors not
	// giving a conclusive result).
	//
	DoorPhysicalStateUnknown DoorPhysicalState = "Unknown"

	// Door is open.
	DoorPhysicalStateOpen DoorPhysicalState = "Open"

	// Door is closed.
	DoorPhysicalStateClosed DoorPhysicalState = "Closed"

	// Door monitor fault is detected.
	DoorPhysicalStateFault DoorPhysicalState = "Fault"
)

//
// The physical state of a Lock (including Double Lock).
//

// LockPhysicalState type
type LockPhysicalState string

const (

	// Value is currently not known.
	LockPhysicalStateUnknown LockPhysicalState = "Unknown"

	// Lock is activated.
	LockPhysicalStateLocked LockPhysicalState = "Locked"

	// Lock is not activated.
	LockPhysicalStateUnlocked LockPhysicalState = "Unlocked"

	// Lock fault is detected.
	LockPhysicalStateFault LockPhysicalState = "Fault"
)

//
// Describes the state of a Door with regard to alarms.
//

// DoorAlarmState type
type DoorAlarmState string

const (

	// No alarm.
	DoorAlarmStateNormal DoorAlarmState = "Normal"

	// Door is forced open.
	DoorAlarmStateDoorForcedOpen DoorAlarmState = "DoorForcedOpen"

	// Door is held open too long.
	DoorAlarmStateDoorOpenTooLong DoorAlarmState = "DoorOpenTooLong"
)

//
// Describes the state of a Tamper detector.
//

// DoorTamperState type
type DoorTamperState string

const (

	// Value is currently not known.
	DoorTamperStateUnknown DoorTamperState = "Unknown"

	// No tampering is detected.
	DoorTamperStateNotInTamper DoorTamperState = "NotInTamper"

	// Tampering is detected.
	DoorTamperStateTamperDetected DoorTamperState = "TamperDetected"
)

//
// Describes the state of a Door fault.
//

// DoorFaultState type
type DoorFaultState string

const (

	// Fault state is unknown.
	DoorFaultStateUnknown DoorFaultState = "Unknown"

	// No fault is detected.
	DoorFaultStateNotInFault DoorFaultState = "NotInFault"

	// Fault is detected.
	DoorFaultStateFaultDetected DoorFaultState = "FaultDetected"
)

//
// The DoorMode describe the mode of operation from a logical perspective.
// Setting a door mode reflects the intent to set a door in a physical state.
//

// DoorMode type
type DoorMode string

const (

	// The mode of operation is unknown.
	DoorModeUnknown DoorMode = "Unknown"

	//
	// The intention is to set the door to a physical locked state.
	// In this mode the device shall provide momentary access using the AccessDoor
	// method if supported by the door instance.
	//
	DoorModeLocked DoorMode = "Locked"

	//
	// The intention is to set the door to a physical unlocked state.
	// Alarms related to door timing operations such as open too long
	// or forced open are masked in this mode.
	//
	DoorModeUnlocked DoorMode = "Unlocked"

	//
	// The intention is to momentary set the door to a physical unlocked state.
	// After a predefined time the device shall revert the door to its previous mode.
	// Alarms related to timing operations such as door forced open are masked in this mode.
	//
	DoorModeAccessed DoorMode = "Accessed"

	//
	// The intention is to set the door to a physical locked state and the
	// device shall not allow AccessDoor requests, i.e. it is not possible
	// for the door to go to the accessed mode.
	// All other requests to change the door mode are allowed.
	//
	DoorModeBlocked DoorMode = "Blocked"

	//
	// The intention is to set the door to a physical locked state and the device
	// shall only allow the LockDownReleaseDoor request.
	// All other requests to change the door mode are not allowed.
	//
	DoorModeLockedDown DoorMode = "LockedDown"

	//
	// The intention is to set the door to a physical unlocked state and the
	// device shall only allow the LockOpenReleaseDoor request.
	// All other requests to change the door mode are not allowed.
	//
	DoorModeLockedOpen DoorMode = "LockedOpen"

	//
	// The intention is to set the door with multiple locks to a physical double locked state.
	// If the door does not support double locking the devices shall
	// treat this as a normal locked mode.
	// When changing to an unlocked mode from the double locked mode, the physical state
	// of the door may first go to locked state before unlocking.
	//
	DoorModeDoubleLocked DoorMode = "DoubleLocked"
)

// Capabilities type
type Capabilities ServiceCapabilities

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capability response message contains the requested DoorControl
	// service capabilities using a hierarchical XML capability structure.
	//
	Capabilities ServiceCapabilities `xml:"Capabilities,omitempty"`
}

// GetDoorInfoList type
type GetDoorInfoList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl GetDoorInfoList"`

	// Maximum number of entries to return. If Limit is omitted or if the
	// value of Limit is higher than what the device supports, then the device shall
	// return its maximum amount of entries.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl StartReference,omitempty"`
}

// GetDoorInfoListResponse type
type GetDoorInfoListResponse struct {
	XMLName xml.Name `xml:"GetDoorInfoListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of DoorInfo items.
	DoorInfo []DoorInfo `xml:"DoorInfo,omitempty"`
}

// GetDoorInfo type
type GetDoorInfo struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl GetDoorInfo"`

	// Tokens of DoorInfo items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// GetDoorInfoResponse type
type GetDoorInfoResponse struct {
	XMLName xml.Name `xml:"GetDoorInfoResponse"`

	// List of DoorInfo items.
	DoorInfo []DoorInfo `xml:"DoorInfo,omitempty"`
}

// GetDoorList type
type GetDoorList struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl GetDoorList"`

	// Maximum number of entries to return. If not specified, less than one
	// or higher than what the device supports, the number of items is determined by the
	// device.
	//
	Limit int32 `xml:"http://www.onvif.org/ver10/schema Limit,omitempty"`

	// Start returning entries from this start reference. If not specified,
	// entries shall start from the beginning of the dataset.
	//
	StartReference string `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl StartReference,omitempty"`
}

// GetDoorListResponse type
type GetDoorListResponse struct {
	XMLName xml.Name `xml:"GetDoorListResponse"`

	// StartReference to use in next call to get the following items. If
	// absent, no more items to get.
	//
	NextStartReference string `xml:"NextStartReference,omitempty"`

	// List of Door items.
	Door []Door `xml:"Door,omitempty"`
}

// GetDoors type
type GetDoors struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl GetDoors"`

	// Tokens of Door items to get.
	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// GetDoorsResponse type
type GetDoorsResponse struct {
	XMLName xml.Name `xml:"GetDoorsResponse"`

	// List of Door items.
	Door []Door `xml:"Door,omitempty"`
}

// CreateDoor type
type CreateDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl CreateDoor"`

	// Door item to create
	Door Door `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Door,omitempty"`
}

// CreateDoorResponse type
type CreateDoorResponse struct {
	XMLName xml.Name `xml:"CreateDoorResponse"`

	// Token of created Door item
	Token ReferenceToken `xml:"Token,omitempty"`
}

// SetDoor type
type SetDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl SetDoor"`

	// The Door item to create or modify
	Door Door `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Door,omitempty"`
}

// SetDoorResponse type
type SetDoorResponse struct {
	XMLName xml.Name `xml:"SetDoorResponse"`
}

// ModifyDoor type
type ModifyDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl ModifyDoor"`

	// The details of the door
	Door Door `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Door,omitempty"`
}

// ModifyDoorResponse type
type ModifyDoorResponse struct {
	XMLName xml.Name `xml:"ModifyDoorResponse"`
}

// DeleteDoor type
type DeleteDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DeleteDoor"`

	// The Token of the door to delete.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// DeleteDoorResponse type
type DeleteDoorResponse struct {
	XMLName xml.Name `xml:"DeleteDoorResponse"`
}

// GetDoorState type
type GetDoorState struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl GetDoorState"`

	// Token of the Door instance to get the state for.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// GetDoorStateResponse type
type GetDoorStateResponse struct {
	XMLName xml.Name `xml:"GetDoorStateResponse"`

	// The state of the door.
	DoorState DoorState `xml:"DoorState,omitempty"`
}

// AccessDoor type
type AccessDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl AccessDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`

	// Optional - Indicates that the configured extended time should be
	// used.
	//
	UseExtendedTime bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl UseExtendedTime,omitempty"`

	// Optional - overrides ReleaseTime if specified.
	AccessTime Duration `xml:"http://www.onvif.org/ver10/schema AccessTime,omitempty"`

	// Optional - overrides OpenTime if specified.
	//
	OpenTooLongTime Duration `xml:"http://www.onvif.org/ver10/schema OpenTooLongTime,omitempty"`

	// Optional - overrides PreAlarmTime if specified.
	PreAlarmTime Duration `xml:"http://www.onvif.org/ver10/schema PreAlarmTime,omitempty"`

	// Future extension.
	Extension AccessDoorExtension `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Extension,omitempty"`
}

// AccessDoorResponse type
type AccessDoorResponse struct {
	XMLName xml.Name `xml:"AccessDoorResponse"`
}

// LockDoor type
type LockDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// LockDoorResponse type
type LockDoorResponse struct {
	XMLName xml.Name `xml:"LockDoorResponse"`
}

// UnlockDoor type
type UnlockDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl UnlockDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// UnlockDoorResponse type
type UnlockDoorResponse struct {
	XMLName xml.Name `xml:"UnlockDoorResponse"`
}

// BlockDoor type
type BlockDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl BlockDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// BlockDoorResponse type
type BlockDoorResponse struct {
	XMLName xml.Name `xml:"BlockDoorResponse"`
}

// LockDownDoor type
type LockDownDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockDownDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// LockDownDoorResponse type
type LockDownDoorResponse struct {
	XMLName xml.Name `xml:"LockDownDoorResponse"`
}

// LockDownReleaseDoor type
type LockDownReleaseDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockDownReleaseDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// LockDownReleaseDoorResponse type
type LockDownReleaseDoorResponse struct {
	XMLName xml.Name `xml:"LockDownReleaseDoorResponse"`
}

// LockOpenDoor type
type LockOpenDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockOpenDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// LockOpenDoorResponse type
type LockOpenDoorResponse struct {
	XMLName xml.Name `xml:"LockOpenDoorResponse"`
}

// LockOpenReleaseDoor type
type LockOpenReleaseDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockOpenReleaseDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// LockOpenReleaseDoorResponse type
type LockOpenReleaseDoorResponse struct {
	XMLName xml.Name `xml:"LockOpenReleaseDoorResponse"`
}

// DoubleLockDoor type
type DoubleLockDoor struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoubleLockDoor"`

	// Token of the Door instance to control.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Token,omitempty"`
}

// DoubleLockDoorResponse type
type DoubleLockDoorResponse struct {
	XMLName xml.Name `xml:"DoubleLockDoorResponse"`
}

// ServiceCapabilities type
type ServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Capabilities"`

	//
	// The maximum number of entries returned by a single Get<Entity>List or
	// Get<Entity> request. The device shall never return more than this number of entities
	// in a single response.
	//

	MaxLimit uint32 `xml:"http://www.onvif.org/ver10/schema MaxLimit,attr,omitempty"`

	//
	// Indicates the maximum number of doors supported by the device.
	//

	MaxDoors uint32 `xml:"http://www.onvif.org/ver10/schema MaxDoors,attr,omitempty"`

	//
	// Indicates that the client is allowed to supply the token when creating doors.
	// To enable the use of the command SetDoor, the value must be set to true.
	//

	ClientSuppliedTokenSupported bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl ClientSuppliedTokenSupported,attr,omitempty"`
}

// DoorInfoBase type
type DoorInfoBase struct {
	*DataEntity

	// A user readable name. It shall be up to 64 characters.
	//
	Name Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Name,omitempty"`

	// A user readable description. It shall be up to 1024 characters.
	//
	Description Description `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Description,omitempty"`
}

// DoorInfo type
type DoorInfo struct {
	*DoorInfoBase

	// The capabilities of the Door.
	Capabilities DoorCapabilities `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Capabilities,omitempty"`
}

// Door type
type Door struct {
	*DoorInfo

	//
	// The type of door. Is of type text. Can be either one of the following reserved
	// ONVIF types: "pt:Door", "pt:ManTrap", "pt:Turnstile", "pt:RevolvingDoor",
	// "pt:Barrier", or a custom defined type.
	//
	DoorType Name `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoorType,omitempty"`

	//
	// A structure defining times such as how long the door is unlocked when
	// accessed, extended grant time, etc.
	//
	Timings Timings `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Timings,omitempty"`

	Extension DoorExtension `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Extension,omitempty"`
}

// DoorExtension type
type DoorExtension struct {
}

// Timings type
type Timings struct {

	//
	// When access is granted (door mode becomes Accessed), the latch is unlocked.
	// ReleaseTime is the time from when the latch is unlocked until it is
	// relocked again (unless the door is physically opened).
	//
	ReleaseTime Duration `xml:"http://www.onvif.org/ver10/schema ReleaseTime,omitempty"`

	//
	// The time from when the door is physically opened until the door is set in the
	// DoorOpenTooLong alarm state.
	//
	OpenTime Duration `xml:"http://www.onvif.org/ver10/schema OpenTime,omitempty"`

	//
	// Some individuals need extra time to open the door before the latch relocks.
	// If supported, ExtendedReleaseTime shall be added to ReleaseTime if UseExtendedTime
	// is set to true in the AccessDoor command.
	//
	ExtendedReleaseTime Duration `xml:"http://www.onvif.org/ver10/schema ExtendedReleaseTime,omitempty"`

	//
	// If the door is physically opened after access is granted,
	// then DelayTimeBeforeRelock is the time from when the door is physically
	// opened until the latch goes back to locked state.
	//
	DelayTimeBeforeRelock Duration `xml:"http://www.onvif.org/ver10/schema DelayTimeBeforeRelock,omitempty"`

	//
	// Some individuals need extra time to pass through the door. If supported,
	// ExtendedOpenTime shall be added to OpenTime if UseExtendedTime is set to true
	// in the AccessDoor command.
	//
	ExtendedOpenTime Duration `xml:"http://www.onvif.org/ver10/schema ExtendedOpenTime,omitempty"`

	//
	// Before a DoorOpenTooLong alarm state is generated, a signal will sound to indicate
	// that the door must be closed. PreAlarmTime defines how long before DoorOpenTooLong
	// the warning signal shall sound.
	//
	PreAlarmTime Duration `xml:"http://www.onvif.org/ver10/schema PreAlarmTime,omitempty"`

	Extension TimingsExtension `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Extension,omitempty"`
}

// TimingsExtension type
type TimingsExtension struct {
}

// DoorCapabilities type
type DoorCapabilities struct {

	// Indicates whether or not this Door instance supports AccessDoor command to
	// perform momentary access.
	//

	Access bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Access,attr,omitempty"`

	// Indicates that this Door instance supports overriding configured timing in the
	// AccessDoor command.
	//

	AccessTimingOverride bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl AccessTimingOverride,attr,omitempty"`

	// Indicates that this Door instance supports LockDoor command to lock the
	// door.
	//

	Lock bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Lock,attr,omitempty"`

	// Indicates that this Door instance supports UnlockDoor command to unlock the
	// door.
	//

	Unlock bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Unlock,attr,omitempty"`

	// Indicates that this Door instance supports BlockDoor command to block the
	// door.
	//

	Block bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Block,attr,omitempty"`

	// Indicates that this Door instance supports DoubleLockDoor command to lock
	// multiple locks on the door.
	//

	DoubleLock bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoubleLock,attr,omitempty"`

	// Indicates that this Door instance supports LockDown (and LockDownRelease)
	// commands to lock the door and put it in LockedDown mode.
	//

	LockDown bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockDown,attr,omitempty"`

	// Indicates that this Door instance supports LockOpen (and LockOpenRelease)
	// commands to unlock the door and put it in LockedOpen mode.
	//

	LockOpen bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockOpen,attr,omitempty"`

	// Indicates that this Door instance has a DoorMonitor and supports the
	// DoorPhysicalState event.
	//

	DoorMonitor bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoorMonitor,attr,omitempty"`

	// Indicates that this Door instance has a LockMonitor and supports the
	// LockPhysicalState event.
	//

	LockMonitor bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockMonitor,attr,omitempty"`

	// Indicates that this Door instance has a DoubleLockMonitor and supports the
	// DoubleLockPhysicalState event.
	//

	DoubleLockMonitor bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoubleLockMonitor,attr,omitempty"`

	// Indicates that this Door instance supports door alarm and the DoorAlarm
	// event.
	//

	Alarm bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Alarm,attr,omitempty"`

	// Indicates that this Door instance has a Tamper detector and supports the
	// DoorTamper event.
	//

	Tamper bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Tamper,attr,omitempty"`

	// Indicates that this Door instance supports door fault and the DoorFault
	// event.
	//

	Fault bool `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Fault,attr,omitempty"`
}

// DoorState type
type DoorState struct {

	//
	// Physical state of the Door; it is of type DoorPhysicalState. A device that
	// signals support for DoorMonitor capability for a particular door instance shall provide
	// this field.
	//
	DoorPhysicalState DoorPhysicalState `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoorPhysicalState,omitempty"`

	//
	// Physical state of the Lock; it is of type LockPhysicalState. A device that
	// signals support for LockMonitor capability for a particular door instance shall provide
	// this field.
	//
	LockPhysicalState LockPhysicalState `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl LockPhysicalState,omitempty"`

	//
	// Physical state of the DoubleLock; it is of type LockPhysicalState. A
	// device that signals support for DoubleLockMonitor capability for a particular door
	// instance shall provide this field.
	//
	DoubleLockPhysicalState LockPhysicalState `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoubleLockPhysicalState,omitempty"`

	//
	// Alarm state of the door; it is of type DoorAlarmState. A device that
	// signals support for Alarm capability for a particular door instance shall provide this
	// field.
	//
	Alarm DoorAlarmState `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Alarm,omitempty"`

	//
	// Tampering state of the door; it is of type DoorTamper. A device that
	// signals support for Tamper capability for a particular door instance shall provide this
	// field.
	//
	Tamper DoorTamper `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Tamper,omitempty"`

	//
	// Fault information for door; it is of type DoorFault. A device that signals
	// support for Fault capability for a particular door instance shall provide this field.
	//
	Fault DoorFault `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Fault,omitempty"`

	//
	// The logical operating mode of the door; it is of type DoorMode. An ONVIF
	// compatible device shall report current operating mode in this field.
	//
	DoorMode DoorMode `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl DoorMode,omitempty"`
}

// DoorTamper type
type DoorTamper struct {

	// Optional field; Details describing tampering state change (e.g., reason,
	// place and time).
	// NOTE: All fields (including this one) which are designed to give
	// end-user prompts can be localized to the customer's native language.
	//
	Reason string `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Reason,omitempty"`

	// State of the tamper detector; it is of type DoorTamperState.
	//
	State DoorTamperState `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl State,omitempty"`
}

// DoorFault type
type DoorFault struct {

	// Optional reason for fault.
	Reason string `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl Reason,omitempty"`

	// Overall fault state for the door; it is of type DoorFaultState. If there
	// are any faults, the value shall be: FaultDetected. Details of the detected fault shall
	// be found in the Reason field, and/or the various DoorState fields and/or in extensions
	// to this structure.
	//
	State DoorFaultState `xml:"http://www.onvif.org/ver10/doorcontrol/wsdl State,omitempty"`
}

// AccessDoorExtension type
type AccessDoorExtension struct {
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

// DoorControlPort type
type DoorControlPort interface {

	/*
		This operation returns the capabilities of the service.
		An ONVIF compliant device which provides the Door Control service shall implement this method.
	*/
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/*
		This operation requests a list of all DoorInfo items provided by the device.
		An ONVIF compliant device that provides Door Control service shall implement
		this method.
		A call to this method shall return a StartReference when not all data is returned and more data is
		available.
		The reference shall be valid for retrieving the next set of data.
		The number of items returned shall not be greater than Limit parameter.
	*/
	GetDoorInfoList(request *GetDoorInfoList) (*GetDoorInfoListResponse, error)

	GetDoorInfoListContext(ctx context.Context, request *GetDoorInfoList) (*GetDoorInfoListResponse, error)

	/*
		This operation requests a list of DoorInfo items matching the given tokens.
		An ONVIF-compliant device that provides Door Control service shall implement this method.
		The device shall ignore tokens it cannot resolve and shall return an empty list
		if there are no items matching specified tokens.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetDoorInfo(request *GetDoorInfo) (*GetDoorInfoResponse, error)

	GetDoorInfoContext(ctx context.Context, request *GetDoorInfo) (*GetDoorInfoResponse, error)

	/*
		This operation requests a list of all Door items provided by the device.
		A call to this method shall return a StartReference when not all data is returned and more data is
		available. The reference shall be valid for retrieving the next set of data.
		Please refer to section 4.8.3 in [Access Control Service Specification] for more details.
		The number of items returned shall not be greater than the Limit parameter.
	*/
	GetDoorList(request *GetDoorList) (*GetDoorListResponse, error)

	GetDoorListContext(ctx context.Context, request *GetDoorList) (*GetDoorListResponse, error)

	/*
		This operation requests a list of Door items matching the given tokens.
		The device shall ignore tokens it cannot resolve and shall return an empty list if there are no items
		matching specified tokens. The device shall not return a fault in this case.
		If the number of requested items is greater than MaxLimit, a TooManyItems fault shall be returned.
	*/
	GetDoors(request *GetDoors) (*GetDoorsResponse, error)

	GetDoorsContext(ctx context.Context, request *GetDoors) (*GetDoorsResponse, error)

	/*
		This operation creates the specified door in the device.
		The token field of the Door structure shall be empty and the device shall allocate a token for
		the door. The allocated token shall be returned in the response.
		If the client sends any value in the token field, the device shall return
		InvalidArgVal as a generic fault code.
	*/
	CreateDoor(request *CreateDoor) (*CreateDoorResponse, error)

	CreateDoorContext(ctx context.Context, request *CreateDoor) (*CreateDoorResponse, error)

	/*
		This method is used to synchronize a door in a client with the device.
		If a door with the specified token does not exist in the device, the door is created.
		If a door with the specified token exists, then the door is modified.
		A call to this method takes a door structure as input parameter. The token field of the Door
		structure shall not be empty.
		A device that signals support for the ClientSuppliedTokenSupported capability shall
		implement this command.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	SetDoor(request *SetDoor) (*SetDoorResponse, error)

	SetDoorContext(ctx context.Context, request *SetDoor) (*SetDoorResponse, error)

	/*
		This operation modifies the specified door.
		The token of the door to modify is specified in the token field of the Door structure and shall
		not be empty. All other fields in the structure shall overwrite the fields in the specified door.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	ModifyDoor(request *ModifyDoor) (*ModifyDoorResponse, error)

	ModifyDoorContext(ctx context.Context, request *ModifyDoor) (*ModifyDoorResponse, error)

	/*
		This operation deletes the specified door.
		If it is associated with one or more entities some devices may not be able to delete the door,
		and consequently a ReferenceInUse fault shall be generated.
		If no token was specified in the request, the device shall return InvalidArgs as a generic fault code.
	*/
	DeleteDoor(request *DeleteDoor) (*DeleteDoorResponse, error)

	DeleteDoorContext(ctx context.Context, request *DeleteDoor) (*DeleteDoorResponse, error)

	/*
		This operation requests the state of a Door specified by the Token.
		A device implementing the Door Control service shall be capable of reporting
		the status of a door using a DoorState structure available from the
		GetDoorState command.
	*/
	GetDoorState(request *GetDoorState) (*GetDoorStateResponse, error)

	GetDoorStateContext(ctx context.Context, request *GetDoorState) (*GetDoorStateResponse, error)

	/*
		This operation allows momentarily accessing a Door.
		It invokes the functionality typically used when a card holder presents a
		card to a card reader at the door and is granted access.
		The DoorMode shall change to Accessed state. Please refer to Accessed mode in section [DoorMode] for
		more details.
		The Door shall remain accessible for the defined time. When the time span
		elapses, the DoorMode shall change back to its previous state.
		If the request cannot be fulfilled, a Failure fault shall be returned.
		Please refer to section [DoorMode] for details about Door Modes restrictions.
		A device that signals support for Access capability for a particular Door
		instance shall implement this method. A device that signals support for
		AccessTimingOverride capability for a particular Door instance shall also
		provide optional timing parameters (AccessTime, OpenTooLongTime and
		PreAlarmTime) when performing AccessDoor command.
		The device shall take the best effort approach for parameters not supported,
		it must fallback to preconfigured time or limit the time to the closest
		supported time if the specified time is out of range.
	*/
	AccessDoor(request *AccessDoor) (*AccessDoorResponse, error)

	AccessDoorContext(ctx context.Context, request *AccessDoor) (*AccessDoorResponse, error)

	/*
		This operation allows locking a Door.
		The DoorMode shall change to Locked state.
		Please refer to Locked mode in section [DoorMode] for more details.
		A device that signals support for Lock capability for a particular Door
		instance shall implement this method.
		If the request cannot be fulfilled, a Failure fault shall be returned.
		Please refer to section [DoorMode] for more details about Door Modes restrictions.
	*/
	LockDoor(request *LockDoor) (*LockDoorResponse, error)

	LockDoorContext(ctx context.Context, request *LockDoor) (*LockDoorResponse, error)

	/*
		This operation allows unlocking a Door.
		The DoorMode shall change to Unlocked state.
		Please refer to Unlocked mode in section [DoorMode] for more details.
		A device that signals support for Unlock capability for a particular Door
		instance shall implement this method.
		If the request cannot be fulfilled, a Failure fault shall be returned.
		Please refer to section [DoorMode] for more details about Door Modes restrictions.
	*/
	UnlockDoor(request *UnlockDoor) (*UnlockDoorResponse, error)

	UnlockDoorContext(ctx context.Context, request *UnlockDoor) (*UnlockDoorResponse, error)

	/*
		This operation allows blocking a Door and preventing momentary access (AccessDoor command).
		The DoorMode shall change to Blocked state.
		Please refer to Blocked mode in section [DoorMode] for more details.
		A device that signals support for Block capability for a particular Door
		instance shall implement this method.
		If the request cannot be fulfilled, a Failure fault shall be returned.
		Please refer to section [DoorMode] for more details about Door Modes restrictions.
	*/
	BlockDoor(request *BlockDoor) (*BlockDoorResponse, error)

	BlockDoorContext(ctx context.Context, request *BlockDoor) (*BlockDoorResponse, error)

	/*
		This operation allows locking and preventing other actions until a LockDownRelease command is invoked.
		The DoorMode shall change to LockedDown state.
		Please refer to LockedDown mode in section [DoorMode] for more details.
		The device shall ignore other door control commands until a LockDownRelease command is performed.
		A device that signals support for LockDown capability for a particular Door
		instance shall implement this method.
		If a device supports DoubleLock capability for a particular Door instance,
		that operation may be engaged as well.
		If the request cannot be fulfilled, a Failure fault shall be returned.
		Please refer to section [DoorMode] for more details about Door Modes restrictions.
	*/
	LockDownDoor(request *LockDownDoor) (*LockDownDoorResponse, error)

	LockDownDoorContext(ctx context.Context, request *LockDownDoor) (*LockDownDoorResponse, error)

	/*
		This operation allows releasing the LockedDown state of a Door.
		The DoorMode shall change back to its previous/next state.
		It is not defined what the previous/next state shall be, but typically - Locked.
		This method shall only succeed if the current DoorMode is LockedDown.
	*/
	LockDownReleaseDoor(request *LockDownReleaseDoor) (*LockDownReleaseDoorResponse, error)

	LockDownReleaseDoorContext(ctx context.Context, request *LockDownReleaseDoor) (*LockDownReleaseDoorResponse, error)

	/*
		This operation allows unlocking a Door and preventing other actions until LockOpenRelease method is
		invoked.
		The DoorMode shall change to LockedOpen state.
		Please refer to LockedOpen mode in section [DoorMode] for more details.
		The device shall ignore other door control commands until a LockOpenRelease command is performed.
		A device that signals support for LockOpen capability for a particular Door instance shall implement
		this method.
		If the request cannot be fulfilled, a Failure fault shall be returned.
		Please refer to section [DoorMode] for more details about Door Modes restrictions.
	*/
	LockOpenDoor(request *LockOpenDoor) (*LockOpenDoorResponse, error)

	LockOpenDoorContext(ctx context.Context, request *LockOpenDoor) (*LockOpenDoorResponse, error)

	/*
		This operation allows releasing the LockedOpen state of a Door.
		The DoorMode shall change state from the LockedOpen state back to its previous/next state.
		It is not defined what the previous/next state shall be, but typically - Unlocked.
		A device that signals support for LockOpen capability for a particular Door instance shall support
		this command.
		This method shall only succeed if the current DoorMode is LockedOpen.
	*/
	LockOpenReleaseDoor(request *LockOpenReleaseDoor) (*LockOpenReleaseDoorResponse, error)

	LockOpenReleaseDoorContext(ctx context.Context, request *LockOpenReleaseDoor) (*LockOpenReleaseDoorResponse, error)

	/*
		This operation is used for securely locking a Door.
		A call to this method shall change DoorMode state to DoubleLocked.
		Please refer to DoubleLocked mode in section [DoorMode] for more details.
		A device that signals support for DoubleLock capability for a particular
		Door instance shall implement this method. Otherwise this method can be
		performed as a standard Lock operation (see [LockDoor command]).
		If the door has an extra lock that shall be locked as well.
		If the request cannot be fulfilled, a Failure fault shall be returned.
	*/
	DoubleLockDoor(request *DoubleLockDoor) (*DoubleLockDoorResponse, error)

	DoubleLockDoorContext(ctx context.Context, request *DoubleLockDoor) (*DoubleLockDoorResponse, error)
}

// doorControlPort type
type doorControlPort struct {
	client *soap.Client
	xaddr  string
}

func NewDoorControlPort(client *soap.Client, xaddr string) DoorControlPort {
	return &doorControlPort{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *doorControlPort) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) GetDoorInfoListContext(ctx context.Context, request *GetDoorInfoList) (*GetDoorInfoListResponse, error) {
	response := new(GetDoorInfoListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/GetDoorInfoList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) GetDoorInfoList(request *GetDoorInfoList) (*GetDoorInfoListResponse, error) {
	return service.GetDoorInfoListContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) GetDoorInfoContext(ctx context.Context, request *GetDoorInfo) (*GetDoorInfoResponse, error) {
	response := new(GetDoorInfoResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/GetDoorInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) GetDoorInfo(request *GetDoorInfo) (*GetDoorInfoResponse, error) {
	return service.GetDoorInfoContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) GetDoorListContext(ctx context.Context, request *GetDoorList) (*GetDoorListResponse, error) {
	response := new(GetDoorListResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/GetDoorList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) GetDoorList(request *GetDoorList) (*GetDoorListResponse, error) {
	return service.GetDoorListContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) GetDoorsContext(ctx context.Context, request *GetDoors) (*GetDoorsResponse, error) {
	response := new(GetDoorsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/GetDoors", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) GetDoors(request *GetDoors) (*GetDoorsResponse, error) {
	return service.GetDoorsContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) CreateDoorContext(ctx context.Context, request *CreateDoor) (*CreateDoorResponse, error) {
	response := new(CreateDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/CreateDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) CreateDoor(request *CreateDoor) (*CreateDoorResponse, error) {
	return service.CreateDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) SetDoorContext(ctx context.Context, request *SetDoor) (*SetDoorResponse, error) {
	response := new(SetDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/SetDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) SetDoor(request *SetDoor) (*SetDoorResponse, error) {
	return service.SetDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) ModifyDoorContext(ctx context.Context, request *ModifyDoor) (*ModifyDoorResponse, error) {
	response := new(ModifyDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/ModifyDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) ModifyDoor(request *ModifyDoor) (*ModifyDoorResponse, error) {
	return service.ModifyDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) DeleteDoorContext(ctx context.Context, request *DeleteDoor) (*DeleteDoorResponse, error) {
	response := new(DeleteDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/DeleteDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) DeleteDoor(request *DeleteDoor) (*DeleteDoorResponse, error) {
	return service.DeleteDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) GetDoorStateContext(ctx context.Context, request *GetDoorState) (*GetDoorStateResponse, error) {
	response := new(GetDoorStateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/GetDoorState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) GetDoorState(request *GetDoorState) (*GetDoorStateResponse, error) {
	return service.GetDoorStateContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) AccessDoorContext(ctx context.Context, request *AccessDoor) (*AccessDoorResponse, error) {
	response := new(AccessDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/AccessDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) AccessDoor(request *AccessDoor) (*AccessDoorResponse, error) {
	return service.AccessDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) LockDoorContext(ctx context.Context, request *LockDoor) (*LockDoorResponse, error) {
	response := new(LockDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/LockDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) LockDoor(request *LockDoor) (*LockDoorResponse, error) {
	return service.LockDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) UnlockDoorContext(ctx context.Context, request *UnlockDoor) (*UnlockDoorResponse, error) {
	response := new(UnlockDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/UnlockDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) UnlockDoor(request *UnlockDoor) (*UnlockDoorResponse, error) {
	return service.UnlockDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) BlockDoorContext(ctx context.Context, request *BlockDoor) (*BlockDoorResponse, error) {
	response := new(BlockDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/BlockDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) BlockDoor(request *BlockDoor) (*BlockDoorResponse, error) {
	return service.BlockDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) LockDownDoorContext(ctx context.Context, request *LockDownDoor) (*LockDownDoorResponse, error) {
	response := new(LockDownDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/LockDownDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) LockDownDoor(request *LockDownDoor) (*LockDownDoorResponse, error) {
	return service.LockDownDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) LockDownReleaseDoorContext(ctx context.Context, request *LockDownReleaseDoor) (*LockDownReleaseDoorResponse, error) {
	response := new(LockDownReleaseDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/LockDownReleaseDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) LockDownReleaseDoor(request *LockDownReleaseDoor) (*LockDownReleaseDoorResponse, error) {
	return service.LockDownReleaseDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) LockOpenDoorContext(ctx context.Context, request *LockOpenDoor) (*LockOpenDoorResponse, error) {
	response := new(LockOpenDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/LockOpenDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) LockOpenDoor(request *LockOpenDoor) (*LockOpenDoorResponse, error) {
	return service.LockOpenDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) LockOpenReleaseDoorContext(ctx context.Context, request *LockOpenReleaseDoor) (*LockOpenReleaseDoorResponse, error) {
	response := new(LockOpenReleaseDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/LockOpenReleaseDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) LockOpenReleaseDoor(request *LockOpenReleaseDoor) (*LockOpenReleaseDoorResponse, error) {
	return service.LockOpenReleaseDoorContext(
		context.Background(),
		request,
	)
}

func (service *doorControlPort) DoubleLockDoorContext(ctx context.Context, request *DoubleLockDoor) (*DoubleLockDoorResponse, error) {
	response := new(DoubleLockDoorResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/doorcontrol/wsdl/DoubleLockDoor", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *doorControlPort) DoubleLockDoor(request *DoubleLockDoor) (*DoubleLockDoorResponse, error) {
	return service.DoubleLockDoorContext(
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
