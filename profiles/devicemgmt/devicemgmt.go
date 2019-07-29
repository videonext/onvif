package devicemgmt

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// EAPMethodTypes type
type EAPMethodTypes []int32

// AutoGeoModes type
type AutoGeoModes string

const (

	// Automatic adjustment of the device location.
	AutoGeoModesLocation AutoGeoModes = "Location"

	// Automatic adjustment of the device orientation relative to the compass also called yaw.
	AutoGeoModesHeading AutoGeoModes = "Heading"

	// Automatic adjustment of the deviation from the horizon also called pitch and roll.
	AutoGeoModesLeveling AutoGeoModes = "Leveling"
)

// StorageType type
type StorageType string

const (

	// NFS protocol
	StorageTypeNFS StorageType = "NFS"

	// CIFS protocol
	StorageTypeCIFS StorageType = "CIFS"

	// CDMI protocol
	StorageTypeCDMI StorageType = "CDMI"

	// FTP protocol
	StorageTypeFTP StorageType = "FTP"
)

// GetServices type
type GetServices struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetServices"`

	// Indicates if the service capabilities (untyped) should be included in the response.
	IncludeCapability bool `xml:"http://www.onvif.org/ver10/device/wsdl IncludeCapability,omitempty"`
}

// GetServicesResponse type
type GetServicesResponse struct {
	XMLName xml.Name `xml:"GetServicesResponse"`

	// Each Service element contains information about one service.
	Service []Service `xml:"Service,omitempty"`
}

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capabilities for the device service is returned in the Capabilities element.
	Capabilities DeviceServiceCapabilities `xml:"Capabilities,omitempty"`
}

// GetDeviceInformation type
type GetDeviceInformation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDeviceInformation"`
}

// GetDeviceInformationResponse type
type GetDeviceInformationResponse struct {
	XMLName xml.Name `xml:"GetDeviceInformationResponse"`

	// The manufactor of the device.
	Manufacturer string `xml:"Manufacturer,omitempty"`

	// The device model.
	Model string `xml:"Model,omitempty"`

	// The firmware version in the device.
	FirmwareVersion string `xml:"FirmwareVersion,omitempty"`

	// The serial number of the device.
	SerialNumber string `xml:"SerialNumber,omitempty"`

	// The hardware ID of the device.
	HardwareId string `xml:"HardwareId,omitempty"`
}

// SetSystemDateAndTime type
type SetSystemDateAndTime struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetSystemDateAndTime"`

	// Defines if the date and time is set via NTP or manually.
	DateTimeType SetDateTimeType `xml:"http://www.onvif.org/ver10/device/wsdl DateTimeType,omitempty"`

	// Automatically adjust Daylight savings if defined in TimeZone.
	DaylightSavings bool `xml:"http://www.onvif.org/ver10/device/wsdl DaylightSavings,omitempty"`

	// The time zone in POSIX 1003.1 format
	TimeZone TimeZone `xml:"http://www.onvif.org/ver10/device/wsdl TimeZone,omitempty"`

	// Date and time in UTC. If time is obtained via NTP, UTCDateTime has no meaning
	UTCDateTime string `xml:"http://www.onvif.org/ver10/schema UTCDateTime,omitempty"`
}

// SetSystemDateAndTimeResponse type
type SetSystemDateAndTimeResponse struct {
	XMLName xml.Name `xml:"SetSystemDateAndTimeResponse"`
}

// GetSystemDateAndTime type
type GetSystemDateAndTime struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemDateAndTime"`
}

// GetSystemDateAndTimeResponse type
type GetSystemDateAndTimeResponse struct {
	XMLName xml.Name `xml:"GetSystemDateAndTimeResponse"`

	// Contains information whether system date and time are set manually or by NTP, daylight savings is on or off, time zone in POSIX 1003.1 format and system date and time in UTC and also local system date and time.
	SystemDateAndTime SystemDateTime `xml:"SystemDateAndTime,omitempty"`
}

// SetSystemFactoryDefault type
type SetSystemFactoryDefault struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetSystemFactoryDefault"`

	// Specifies the factory default action type.
	FactoryDefault FactoryDefaultType `xml:"http://www.onvif.org/ver10/device/wsdl FactoryDefault,omitempty"`
}

// SetSystemFactoryDefaultResponse type
type SetSystemFactoryDefaultResponse struct {
	XMLName xml.Name `xml:"SetSystemFactoryDefaultResponse"`
}

// UpgradeSystemFirmware type
type UpgradeSystemFirmware struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl UpgradeSystemFirmware"`

	Firmware AttachmentData `xml:"http://www.onvif.org/ver10/device/wsdl Firmware,omitempty"`
}

// UpgradeSystemFirmwareResponse type
type UpgradeSystemFirmwareResponse struct {
	XMLName xml.Name `xml:"UpgradeSystemFirmwareResponse"`

	Message string `xml:"Message,omitempty"`
}

// SystemReboot type
type SystemReboot struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SystemReboot"`
}

// SystemRebootResponse type
type SystemRebootResponse struct {
	XMLName xml.Name `xml:"SystemRebootResponse"`

	// Contains the reboot message sent by the device.
	Message string `xml:"Message,omitempty"`
}

// RestoreSystem type
type RestoreSystem struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl RestoreSystem"`

	BackupFiles []BackupFile `xml:"http://www.onvif.org/ver10/device/wsdl BackupFiles,omitempty"`
}

// RestoreSystemResponse type
type RestoreSystemResponse struct {
	XMLName xml.Name `xml:"RestoreSystemResponse"`
}

// GetSystemBackup type
type GetSystemBackup struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemBackup"`
}

// GetSystemBackupResponse type
type GetSystemBackupResponse struct {
	XMLName xml.Name `xml:"GetSystemBackupResponse"`

	BackupFiles []BackupFile `xml:"BackupFiles,omitempty"`
}

// GetSystemSupportInformation type
type GetSystemSupportInformation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemSupportInformation"`
}

// GetSystemSupportInformationResponse type
type GetSystemSupportInformationResponse struct {
	XMLName xml.Name `xml:"GetSystemSupportInformationResponse"`

	// Contains the arbitary device diagnostics information.
	SupportInformation SupportInformation `xml:"SupportInformation,omitempty"`
}

// GetSystemLog type
type GetSystemLog struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemLog"`

	// Specifies the type of system log to get.
	LogType SystemLogType `xml:"http://www.onvif.org/ver10/device/wsdl LogType,omitempty"`
}

// GetSystemLogResponse type
type GetSystemLogResponse struct {
	XMLName xml.Name `xml:"GetSystemLogResponse"`

	// Contains the system log information.
	SystemLog SystemLog `xml:"SystemLog,omitempty"`
}

// GetScopes type
type GetScopes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetScopes"`
}

// GetScopesResponse type
type GetScopesResponse struct {
	XMLName xml.Name `xml:"GetScopesResponse"`

	// Contains a list of URI definining the device scopes. Scope parameters can be of two types: fixed and configurable. Fixed parameters can not be altered.
	Scopes []Scope `xml:"Scopes,omitempty"`
}

// SetScopes type
type SetScopes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetScopes"`

	// Contains a list of scope parameters that will replace all existing configurable scope parameters.
	Scopes []AnyURI `xml:"http://www.onvif.org/ver10/schema Scopes,omitempty"`
}

// SetScopesResponse type
type SetScopesResponse struct {
	XMLName xml.Name `xml:"SetScopesResponse"`
}

// AddScopes type
type AddScopes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl AddScopes"`

	// Contains a list of new configurable scope parameters that will be added to the existing configurable scope.
	ScopeItem []AnyURI `xml:"http://www.onvif.org/ver10/schema ScopeItem,omitempty"`
}

// AddScopesResponse type
type AddScopesResponse struct {
	XMLName xml.Name `xml:"AddScopesResponse"`
}

// RemoveScopes type
type RemoveScopes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl RemoveScopes"`

	// Contains a list of URIs that should be removed from the device scope.
	// Note that the response message always will match the request or an error will be returned. The use of the response is for that reason deprecated.
	//
	ScopeItem []AnyURI `xml:"http://www.onvif.org/ver10/schema ScopeItem,omitempty"`
}

// RemoveScopesResponse type
type RemoveScopesResponse struct {
	XMLName xml.Name `xml:"RemoveScopesResponse"`

	// Contains a list of URIs that has been removed from the device scope
	ScopeItem []AnyURI `xml:"ScopeItem,omitempty"`
}

// GetDiscoveryMode type
type GetDiscoveryMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDiscoveryMode"`
}

// GetDiscoveryModeResponse type
type GetDiscoveryModeResponse struct {
	XMLName xml.Name `xml:"GetDiscoveryModeResponse"`

	//
	// Indicator of discovery mode: Discoverable, NonDiscoverable.
	//
	DiscoveryMode DiscoveryMode `xml:"DiscoveryMode,omitempty"`
}

// SetDiscoveryMode type
type SetDiscoveryMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetDiscoveryMode"`

	//
	// Indicator of discovery mode: Discoverable, NonDiscoverable.
	//
	DiscoveryMode DiscoveryMode `xml:"http://www.onvif.org/ver10/device/wsdl DiscoveryMode,omitempty"`
}

// SetDiscoveryModeResponse type
type SetDiscoveryModeResponse struct {
	XMLName xml.Name `xml:"SetDiscoveryModeResponse"`
}

// GetRemoteDiscoveryMode type
type GetRemoteDiscoveryMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetRemoteDiscoveryMode"`
}

// GetRemoteDiscoveryModeResponse type
type GetRemoteDiscoveryModeResponse struct {
	XMLName xml.Name `xml:"GetRemoteDiscoveryModeResponse"`

	//
	// Indicator of discovery mode: Discoverable, NonDiscoverable.
	//
	RemoteDiscoveryMode DiscoveryMode `xml:"RemoteDiscoveryMode,omitempty"`
}

// SetRemoteDiscoveryMode type
type SetRemoteDiscoveryMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetRemoteDiscoveryMode"`

	//
	// Indicator of discovery mode: Discoverable, NonDiscoverable.
	//
	RemoteDiscoveryMode DiscoveryMode `xml:"http://www.onvif.org/ver10/device/wsdl RemoteDiscoveryMode,omitempty"`
}

// SetRemoteDiscoveryModeResponse type
type SetRemoteDiscoveryModeResponse struct {
	XMLName xml.Name `xml:"SetRemoteDiscoveryModeResponse"`
}

// GetDPAddresses type
type GetDPAddresses struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDPAddresses"`
}

// GetDPAddressesResponse type
type GetDPAddressesResponse struct {
	XMLName xml.Name `xml:"GetDPAddressesResponse"`

	DPAddress []NetworkHost `xml:"DPAddress,omitempty"`
}

// SetDPAddresses type
type SetDPAddresses struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetDPAddresses"`

	DPAddress []NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl DPAddress,omitempty"`
}

// SetDPAddressesResponse type
type SetDPAddressesResponse struct {
	XMLName xml.Name `xml:"SetDPAddressesResponse"`
}

// GetEndpointReference type
type GetEndpointReference struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetEndpointReference"`
}

// GetEndpointReferenceResponse type
type GetEndpointReferenceResponse struct {
	XMLName xml.Name `xml:"GetEndpointReferenceResponse"`

	GUID string `xml:"GUID,omitempty"`
}

// GetRemoteUser type
type GetRemoteUser struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetRemoteUser"`
}

// GetRemoteUserResponse type
type GetRemoteUserResponse struct {
	XMLName xml.Name `xml:"GetRemoteUserResponse"`

	RemoteUser RemoteUser `xml:"RemoteUser,omitempty"`
}

// SetRemoteUser type
type SetRemoteUser struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetRemoteUser"`

	RemoteUser RemoteUser `xml:"http://www.onvif.org/ver10/device/wsdl RemoteUser,omitempty"`
}

// SetRemoteUserResponse type
type SetRemoteUserResponse struct {
	XMLName xml.Name `xml:"SetRemoteUserResponse"`
}

// GetUsers type
type GetUsers struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetUsers"`
}

// GetUsersResponse type
type GetUsersResponse struct {
	XMLName xml.Name `xml:"GetUsersResponse"`

	// Contains a list of the onvif users and following information is included in each entry: username and user level.
	User []User `xml:"User,omitempty"`
}

// CreateUsers type
type CreateUsers struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl CreateUsers"`

	// Creates new device users and corresponding credentials. Each user entry includes: username, password and user level. Either all users are created successfully or a fault message MUST be returned without creating any user. If trying to create several users with exactly the same username the request is rejected and no users are created. If password is missing, then fault message Too weak password is returned.
	User []User `xml:"http://www.onvif.org/ver10/device/wsdl User,omitempty"`
}

// CreateUsersResponse type
type CreateUsersResponse struct {
	XMLName xml.Name `xml:"CreateUsersResponse"`
}

// DeleteUsers type
type DeleteUsers struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl DeleteUsers"`

	// Deletes users on an device and there may exist users that cannot be deleted to ensure access to the unit. Either all users are deleted successfully or a fault message MUST be returned and no users be deleted. If a username exists multiple times in the request, then a fault message is returned.
	Username []string `xml:"http://www.onvif.org/ver10/device/wsdl Username,omitempty"`
}

// DeleteUsersResponse type
type DeleteUsersResponse struct {
	XMLName xml.Name `xml:"DeleteUsersResponse"`
}

// SetUser type
type SetUser struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetUser"`

	// Updates the credentials for one or several users on an device. Either all change requests are processed successfully or a fault message MUST be returned. If the request contains the same username multiple times, a fault message is returned.
	User []User `xml:"http://www.onvif.org/ver10/device/wsdl User,omitempty"`
}

// SetUserResponse type
type SetUserResponse struct {
	XMLName xml.Name `xml:"SetUserResponse"`
}

// GetWsdlUrl type
type GetWsdlUrl struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetWsdlUrl"`
}

// GetWsdlUrlResponse type
type GetWsdlUrlResponse struct {
	XMLName xml.Name `xml:"GetWsdlUrlResponse"`

	WsdlUrl AnyURI `xml:"WsdlUrl,omitempty"`
}

// GetCapabilities type
type GetCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetCapabilities"`

	//
	// List of categories to retrieve capability information on.
	//
	Category []CapabilityCategory `xml:"http://www.onvif.org/ver10/device/wsdl Category,omitempty"`
}

// GetCapabilitiesResponse type
type GetCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetCapabilitiesResponse"`

	//
	// Capability information.
	//
	Capabilities Capabilities `xml:"Capabilities,omitempty"`
}

// GetHostname type
type GetHostname struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetHostname"`
}

// GetHostnameResponse type
type GetHostnameResponse struct {
	XMLName xml.Name `xml:"GetHostnameResponse"`

	// Contains the hostname information.
	HostnameInformation HostnameInformation `xml:"HostnameInformation,omitempty"`
}

// SetHostname type
type SetHostname struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetHostname"`

	// The hostname to set.
	Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`
}

// SetHostnameResponse type
type SetHostnameResponse struct {
	XMLName xml.Name `xml:"SetHostnameResponse"`
}

// SetHostnameFromDHCP type
type SetHostnameFromDHCP struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetHostnameFromDHCP"`

	// True if the hostname shall be obtained via DHCP.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP,omitempty"`
}

// SetHostnameFromDHCPResponse type
type SetHostnameFromDHCPResponse struct {
	XMLName xml.Name `xml:"SetHostnameFromDHCPResponse"`

	//
	// Indicates whether or not a reboot is required after configuration updates.
	//
	RebootNeeded bool `xml:"RebootNeeded,omitempty"`
}

// GetDNS type
type GetDNS struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDNS"`
}

// GetDNSResponse type
type GetDNSResponse struct {
	XMLName xml.Name `xml:"GetDNSResponse"`

	//
	// DNS information.
	//
	DNSInformation DNSInformation `xml:"DNSInformation,omitempty"`
}

// SetDNS type
type SetDNS struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetDNS"`

	//
	// Indicate if the DNS address is to be retrieved using DHCP.
	//
	FromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP,omitempty"`

	//
	// DNS search domain.
	//
	SearchDomain []string `xml:"http://www.onvif.org/ver10/device/wsdl SearchDomain,omitempty"`

	//
	// DNS address(es) set manually.
	//
	DNSManual []IPAddress `xml:"http://www.onvif.org/ver10/device/wsdl DNSManual,omitempty"`
}

// SetDNSResponse type
type SetDNSResponse struct {
	XMLName xml.Name `xml:"SetDNSResponse"`
}

// GetNTP type
type GetNTP struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetNTP"`
}

// GetNTPResponse type
type GetNTPResponse struct {
	XMLName xml.Name `xml:"GetNTPResponse"`

	//
	// NTP information.
	//
	NTPInformation NTPInformation `xml:"NTPInformation,omitempty"`
}

// SetNTP type
type SetNTP struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetNTP"`

	//
	// Indicate if NTP address information is to be retrieved using DHCP.
	//
	FromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP,omitempty"`

	//
	// Manual NTP settings.
	//
	NTPManual []NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl NTPManual,omitempty"`
}

// SetNTPResponse type
type SetNTPResponse struct {
	XMLName xml.Name `xml:"SetNTPResponse"`
}

// GetDynamicDNS type
type GetDynamicDNS struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDynamicDNS"`
}

// GetDynamicDNSResponse type
type GetDynamicDNSResponse struct {
	XMLName xml.Name `xml:"GetDynamicDNSResponse"`

	//
	// Dynamic DNS information.
	//
	DynamicDNSInformation DynamicDNSInformation `xml:"DynamicDNSInformation,omitempty"`
}

// SetDynamicDNS type
type SetDynamicDNS struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetDynamicDNS"`

	//
	// Dynamic DNS type.
	//
	Type DynamicDNSType `xml:"http://www.onvif.org/ver10/device/wsdl Type,omitempty"`

	//
	// DNS name.
	//
	Name DNSName `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`

	//
	// DNS record time to live.
	//
	TTL Duration `xml:"http://www.onvif.org/ver10/schema TTL,omitempty"`
}

// SetDynamicDNSResponse type
type SetDynamicDNSResponse struct {
	XMLName xml.Name `xml:"SetDynamicDNSResponse"`
}

// GetNetworkInterfaces type
type GetNetworkInterfaces struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkInterfaces"`
}

// GetNetworkInterfacesResponse type
type GetNetworkInterfacesResponse struct {
	XMLName xml.Name `xml:"GetNetworkInterfacesResponse"`

	//
	// List of network interfaces.
	//
	NetworkInterfaces []NetworkInterface `xml:"NetworkInterfaces,omitempty"`
}

// SetNetworkInterfaces type
type SetNetworkInterfaces struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkInterfaces"`

	//
	// Symbolic network interface name.
	//
	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken,omitempty"`

	//
	// Network interface name.
	//
	NetworkInterface NetworkInterfaceSetConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl NetworkInterface,omitempty"`
}

// SetNetworkInterfacesResponse type
type SetNetworkInterfacesResponse struct {
	XMLName xml.Name `xml:"SetNetworkInterfacesResponse"`

	//
	// Indicates whether or not a reboot is required after configuration updates.
	// If a device responds with RebootNeeded set to false, the device can be reached
	// via the new IP address without further action. A client should be aware that a device
	// may not be responsive for a short period of time until it signals availability at
	// the new address via the discovery Hello messages.
	// If a device responds with RebootNeeded set to true, it will be further available under
	// its previous IP address. The settings will only be activated when the device is
	// rebooted via the SystemReboot command.
	//
	RebootNeeded bool `xml:"RebootNeeded,omitempty"`
}

// GetNetworkProtocols type
type GetNetworkProtocols struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkProtocols"`
}

// GetNetworkProtocolsResponse type
type GetNetworkProtocolsResponse struct {
	XMLName xml.Name `xml:"GetNetworkProtocolsResponse"`

	// Contains an array of defined protocols supported by the device. There are three protocols defined; HTTP, HTTPS and RTSP. The following parameters can be retrieved for each protocol: port and enable/disable.
	NetworkProtocols []NetworkProtocol `xml:"NetworkProtocols,omitempty"`
}

// SetNetworkProtocols type
type SetNetworkProtocols struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkProtocols"`

	// Configures one or more defined network protocols supported by the device. There are currently three protocols defined; HTTP, HTTPS and RTSP. The following parameters can be set for each protocol: port and enable/disable.
	NetworkProtocols []NetworkProtocol `xml:"http://www.onvif.org/ver10/device/wsdl NetworkProtocols,omitempty"`
}

// SetNetworkProtocolsResponse type
type SetNetworkProtocolsResponse struct {
	XMLName xml.Name `xml:"SetNetworkProtocolsResponse"`
}

// GetNetworkDefaultGateway type
type GetNetworkDefaultGateway struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetNetworkDefaultGateway"`
}

// GetNetworkDefaultGatewayResponse type
type GetNetworkDefaultGatewayResponse struct {
	XMLName xml.Name `xml:"GetNetworkDefaultGatewayResponse"`

	// Gets the default IPv4 and IPv6 gateway settings from the device.
	NetworkGateway NetworkGateway `xml:"NetworkGateway,omitempty"`
}

// SetNetworkDefaultGateway type
type SetNetworkDefaultGateway struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetNetworkDefaultGateway"`

	// Sets IPv4 gateway address used as default setting.
	IPv4Address []IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv4Address,omitempty"`

	// Sets IPv6 gateway address used as default setting.
	IPv6Address []IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv6Address,omitempty"`
}

// SetNetworkDefaultGatewayResponse type
type SetNetworkDefaultGatewayResponse struct {
	XMLName xml.Name `xml:"SetNetworkDefaultGatewayResponse"`
}

// GetZeroConfiguration type
type GetZeroConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetZeroConfiguration"`
}

// GetZeroConfigurationResponse type
type GetZeroConfigurationResponse struct {
	XMLName xml.Name `xml:"GetZeroConfigurationResponse"`

	// Contains the zero-configuration.
	ZeroConfiguration NetworkZeroConfiguration `xml:"ZeroConfiguration,omitempty"`
}

// SetZeroConfiguration type
type SetZeroConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetZeroConfiguration"`

	// Unique identifier referencing the physical interface.
	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken,omitempty"`

	// Specifies if the zero-configuration should be enabled or not.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`
}

// SetZeroConfigurationResponse type
type SetZeroConfigurationResponse struct {
	XMLName xml.Name `xml:"SetZeroConfigurationResponse"`
}

// GetIPAddressFilter type
type GetIPAddressFilter struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetIPAddressFilter"`
}

// GetIPAddressFilterResponse type
type GetIPAddressFilterResponse struct {
	XMLName xml.Name `xml:"GetIPAddressFilterResponse"`

	IPAddressFilter IPAddressFilter `xml:"IPAddressFilter,omitempty"`
}

// SetIPAddressFilter type
type SetIPAddressFilter struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetIPAddressFilter"`

	IPAddressFilter IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter,omitempty"`
}

// SetIPAddressFilterResponse type
type SetIPAddressFilterResponse struct {
	XMLName xml.Name `xml:"SetIPAddressFilterResponse"`
}

// AddIPAddressFilter type
type AddIPAddressFilter struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl AddIPAddressFilter"`

	IPAddressFilter IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter,omitempty"`
}

// AddIPAddressFilterResponse type
type AddIPAddressFilterResponse struct {
	XMLName xml.Name `xml:"AddIPAddressFilterResponse"`
}

// RemoveIPAddressFilter type
type RemoveIPAddressFilter struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl RemoveIPAddressFilter"`

	IPAddressFilter IPAddressFilter `xml:"http://www.onvif.org/ver10/device/wsdl IPAddressFilter,omitempty"`
}

// RemoveIPAddressFilterResponse type
type RemoveIPAddressFilterResponse struct {
	XMLName xml.Name `xml:"RemoveIPAddressFilterResponse"`
}

// GetAccessPolicy type
type GetAccessPolicy struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetAccessPolicy"`
}

// GetAccessPolicyResponse type
type GetAccessPolicyResponse struct {
	XMLName xml.Name `xml:"GetAccessPolicyResponse"`

	PolicyFile BinaryData `xml:"PolicyFile,omitempty"`
}

// SetAccessPolicy type
type SetAccessPolicy struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetAccessPolicy"`

	PolicyFile BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl PolicyFile,omitempty"`
}

// SetAccessPolicyResponse type
type SetAccessPolicyResponse struct {
	XMLName xml.Name `xml:"SetAccessPolicyResponse"`
}

// CreateCertificate type
type CreateCertificate struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl CreateCertificate"`

	// Certificate id.
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	// Identification of the entity associated with the public-key.
	Subject string `xml:"http://www.onvif.org/ver10/device/wsdl Subject,omitempty"`

	// Certificate validity start date.
	ValidNotBefore string `xml:"http://www.onvif.org/ver10/schema ValidNotBefore,omitempty"`

	// Certificate expiry start date.
	ValidNotAfter string `xml:"http://www.onvif.org/ver10/schema ValidNotAfter,omitempty"`
}

// CreateCertificateResponse type
type CreateCertificateResponse struct {
	XMLName xml.Name `xml:"CreateCertificateResponse"`

	//
	// base64 encoded DER representation of certificate.
	//
	NvtCertificate Certificate `xml:"NvtCertificate,omitempty"`
}

// GetCertificates type
type GetCertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificates"`
}

// GetCertificatesResponse type
type GetCertificatesResponse struct {
	XMLName xml.Name `xml:"GetCertificatesResponse"`

	//
	// Id and base64 encoded DER representation of all available certificates.
	//
	NvtCertificate []Certificate `xml:"NvtCertificate,omitempty"`
}

// GetCertificatesStatus type
type GetCertificatesStatus struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificatesStatus"`
}

// GetCertificatesStatusResponse type
type GetCertificatesStatusResponse struct {
	XMLName xml.Name `xml:"GetCertificatesStatusResponse"`

	//
	// Indicates if a certificate is used in an optional HTTPS configuration of the device.
	//
	CertificateStatus []CertificateStatus `xml:"CertificateStatus,omitempty"`
}

// SetCertificatesStatus type
type SetCertificatesStatus struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetCertificatesStatus"`

	//
	// Indicates if a certificate is to be used in an optional HTTPS configuration of the device.
	//
	CertificateStatus []CertificateStatus `xml:"http://www.onvif.org/ver10/device/wsdl CertificateStatus,omitempty"`
}

// SetCertificatesStatusResponse type
type SetCertificatesStatusResponse struct {
	XMLName xml.Name `xml:"SetCertificatesStatusResponse"`
}

// DeleteCertificates type
type DeleteCertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl DeleteCertificates"`

	//
	// List of ids of certificates to delete.
	//
	CertificateID []string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`
}

// DeleteCertificatesResponse type
type DeleteCertificatesResponse struct {
	XMLName xml.Name `xml:"DeleteCertificatesResponse"`
}

// GetPkcs10Request type
type GetPkcs10Request struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetPkcs10Request"`

	//
	// List of ids of certificates to delete.
	//
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	//
	// Relative Dinstinguished Name(RDN) CommonName(CN).
	//
	Subject string `xml:"http://www.onvif.org/ver10/device/wsdl Subject,omitempty"`

	//
	// Optional base64 encoded DER attributes.
	//
	Attributes BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl Attributes,omitempty"`
}

// GetPkcs10RequestResponse type
type GetPkcs10RequestResponse struct {
	XMLName xml.Name `xml:"GetPkcs10RequestResponse"`

	//
	// base64 encoded DER representation of certificate.
	//
	Pkcs10Request BinaryData `xml:"Pkcs10Request,omitempty"`
}

// LoadCertificates type
type LoadCertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl LoadCertificates"`

	//
	// Optional id and base64 encoded DER representation of certificate.
	//
	NVTCertificate []Certificate `xml:"http://www.onvif.org/ver10/device/wsdl NVTCertificate,omitempty"`
}

// LoadCertificatesResponse type
type LoadCertificatesResponse struct {
	XMLName xml.Name `xml:"LoadCertificatesResponse"`
}

// GetClientCertificateMode type
type GetClientCertificateMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetClientCertificateMode"`
}

// GetClientCertificateModeResponse type
type GetClientCertificateModeResponse struct {
	XMLName xml.Name `xml:"GetClientCertificateModeResponse"`

	//
	// Indicates whether or not client certificates are required by device.
	//
	Enabled bool `xml:"Enabled,omitempty"`
}

// SetClientCertificateMode type
type SetClientCertificateMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetClientCertificateMode"`

	//
	// Indicates whether or not client certificates are required by device.
	//
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`
}

// SetClientCertificateModeResponse type
type SetClientCertificateModeResponse struct {
	XMLName xml.Name `xml:"SetClientCertificateModeResponse"`
}

// GetCACertificates type
type GetCACertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetCACertificates"`
}

// GetCACertificatesResponse type
type GetCACertificatesResponse struct {
	XMLName xml.Name `xml:"GetCACertificatesResponse"`

	CACertificate []Certificate `xml:"CACertificate,omitempty"`
}

// LoadCertificateWithPrivateKey type
type LoadCertificateWithPrivateKey struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl LoadCertificateWithPrivateKey"`

	CertificateWithPrivateKey []CertificateWithPrivateKey `xml:"http://www.onvif.org/ver10/device/wsdl CertificateWithPrivateKey,omitempty"`
}

// LoadCertificateWithPrivateKeyResponse type
type LoadCertificateWithPrivateKeyResponse struct {
	XMLName xml.Name `xml:"LoadCertificateWithPrivateKeyResponse"`
}

// GetCertificateInformation type
type GetCertificateInformation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetCertificateInformation"`

	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`
}

// GetCertificateInformationResponse type
type GetCertificateInformationResponse struct {
	XMLName xml.Name `xml:"GetCertificateInformationResponse"`

	CertificateInformation CertificateInformation `xml:"CertificateInformation,omitempty"`
}

// LoadCACertificates type
type LoadCACertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl LoadCACertificates"`

	CACertificate []Certificate `xml:"http://www.onvif.org/ver10/device/wsdl CACertificate,omitempty"`
}

// LoadCACertificatesResponse type
type LoadCACertificatesResponse struct {
	XMLName xml.Name `xml:"LoadCACertificatesResponse"`
}

// CreateDot1XConfiguration type
type CreateDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl CreateDot1XConfiguration"`

	Dot1XConfiguration Dot1XConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfiguration,omitempty"`
}

// CreateDot1XConfigurationResponse type
type CreateDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"CreateDot1XConfigurationResponse"`
}

// SetDot1XConfiguration type
type SetDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetDot1XConfiguration"`

	Dot1XConfiguration Dot1XConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfiguration,omitempty"`
}

// SetDot1XConfigurationResponse type
type SetDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"SetDot1XConfigurationResponse"`
}

// GetDot1XConfiguration type
type GetDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDot1XConfiguration"`

	Dot1XConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken,omitempty"`
}

// GetDot1XConfigurationResponse type
type GetDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"GetDot1XConfigurationResponse"`

	Dot1XConfiguration Dot1XConfiguration `xml:"Dot1XConfiguration,omitempty"`
}

// GetDot1XConfigurations type
type GetDot1XConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDot1XConfigurations"`
}

// GetDot1XConfigurationsResponse type
type GetDot1XConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetDot1XConfigurationsResponse"`

	Dot1XConfiguration []Dot1XConfiguration `xml:"Dot1XConfiguration,omitempty"`
}

// DeleteDot1XConfiguration type
type DeleteDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl DeleteDot1XConfiguration"`

	Dot1XConfigurationToken []ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken,omitempty"`
}

// DeleteDot1XConfigurationResponse type
type DeleteDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"DeleteDot1XConfigurationResponse"`
}

// GetRelayOutputs type
type GetRelayOutputs struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetRelayOutputs"`
}

// GetRelayOutputsResponse type
type GetRelayOutputsResponse struct {
	XMLName xml.Name `xml:"GetRelayOutputsResponse"`

	RelayOutputs []RelayOutput `xml:"RelayOutputs,omitempty"`
}

// SetRelayOutputSettings type
type SetRelayOutputSettings struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetRelayOutputSettings"`

	RelayOutputToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputToken,omitempty"`

	Properties RelayOutputSettings `xml:"http://www.onvif.org/ver10/device/wsdl Properties,omitempty"`
}

// SetRelayOutputSettingsResponse type
type SetRelayOutputSettingsResponse struct {
	XMLName xml.Name `xml:"SetRelayOutputSettingsResponse"`
}

// SetRelayOutputState type
type SetRelayOutputState struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetRelayOutputState"`

	RelayOutputToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl RelayOutputToken,omitempty"`

	LogicalState RelayLogicalState `xml:"http://www.onvif.org/ver10/device/wsdl LogicalState,omitempty"`
}

// SetRelayOutputStateResponse type
type SetRelayOutputStateResponse struct {
	XMLName xml.Name `xml:"SetRelayOutputStateResponse"`
}

// SendAuxiliaryCommand type
type SendAuxiliaryCommand struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SendAuxiliaryCommand"`

	AuxiliaryCommand AuxiliaryData `xml:"http://www.onvif.org/ver10/device/wsdl AuxiliaryCommand,omitempty"`
}

// SendAuxiliaryCommandResponse type
type SendAuxiliaryCommandResponse struct {
	XMLName xml.Name `xml:"SendAuxiliaryCommandResponse"`

	AuxiliaryCommandResponse AuxiliaryData `xml:"AuxiliaryCommandResponse,omitempty"`
}

// GetDot11Capabilities type
type GetDot11Capabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDot11Capabilities"`
}

// GetDot11CapabilitiesResponse type
type GetDot11CapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetDot11CapabilitiesResponse"`

	Capabilities Dot11Capabilities `xml:"Capabilities,omitempty"`
}

// GetDot11Status type
type GetDot11Status struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetDot11Status"`

	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken,omitempty"`
}

// GetDot11StatusResponse type
type GetDot11StatusResponse struct {
	XMLName xml.Name `xml:"GetDot11StatusResponse"`

	Status Dot11Status `xml:"Status,omitempty"`
}

// ScanAvailableDot11Networks type
type ScanAvailableDot11Networks struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl ScanAvailableDot11Networks"`

	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken,omitempty"`
}

// ScanAvailableDot11NetworksResponse type
type ScanAvailableDot11NetworksResponse struct {
	XMLName xml.Name `xml:"ScanAvailableDot11NetworksResponse"`

	Networks []Dot11AvailableNetworks `xml:"Networks,omitempty"`
}

// GetSystemUris type
type GetSystemUris struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetSystemUris"`
}

// GetSystemUrisResponse type
type GetSystemUrisResponse struct {
	XMLName xml.Name `xml:"GetSystemUrisResponse"`

	SystemLogUris SystemLogUriList `xml:"SystemLogUris,omitempty"`

	SupportInfoUri AnyURI `xml:"SupportInfoUri,omitempty"`

	SystemBackupUri AnyURI `xml:"SystemBackupUri,omitempty"`

	Extension struct {
	} `xml:"Extension,omitempty"`
}

// StartFirmwareUpgrade type
type StartFirmwareUpgrade struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl StartFirmwareUpgrade"`
}

// StartFirmwareUpgradeResponse type
type StartFirmwareUpgradeResponse struct {
	XMLName xml.Name `xml:"StartFirmwareUpgradeResponse"`

	UploadUri AnyURI `xml:"UploadUri,omitempty"`

	UploadDelay Duration `xml:"UploadDelay,omitempty"`

	ExpectedDownTime Duration `xml:"ExpectedDownTime,omitempty"`
}

// StartSystemRestore type
type StartSystemRestore struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl StartSystemRestore"`
}

// StartSystemRestoreResponse type
type StartSystemRestoreResponse struct {
	XMLName xml.Name `xml:"StartSystemRestoreResponse"`

	UploadUri AnyURI `xml:"UploadUri,omitempty"`

	ExpectedDownTime Duration `xml:"ExpectedDownTime,omitempty"`
}

// GetStorageConfigurations type
type GetStorageConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetStorageConfigurations"`
}

// GetStorageConfigurationsResponse type
type GetStorageConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetStorageConfigurationsResponse"`

	StorageConfigurations []StorageConfiguration `xml:"StorageConfigurations,omitempty"`
}

// CreateStorageConfiguration type
type CreateStorageConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl CreateStorageConfiguration"`

	StorageConfiguration StorageConfigurationData `xml:"http://www.onvif.org/ver10/device/wsdl StorageConfiguration,omitempty"`
}

// CreateStorageConfigurationResponse type
type CreateStorageConfigurationResponse struct {
	XMLName xml.Name `xml:"CreateStorageConfigurationResponse"`

	Token ReferenceToken `xml:"Token,omitempty"`
}

// GetStorageConfiguration type
type GetStorageConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetStorageConfiguration"`

	Token ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token,omitempty"`
}

// GetStorageConfigurationResponse type
type GetStorageConfigurationResponse struct {
	XMLName xml.Name `xml:"GetStorageConfigurationResponse"`

	StorageConfiguration StorageConfiguration `xml:"StorageConfiguration,omitempty"`
}

// SetStorageConfiguration type
type SetStorageConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetStorageConfiguration"`

	StorageConfiguration StorageConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl StorageConfiguration,omitempty"`
}

// SetStorageConfigurationResponse type
type SetStorageConfigurationResponse struct {
	XMLName xml.Name `xml:"SetStorageConfigurationResponse"`
}

// DeleteStorageConfiguration type
type DeleteStorageConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl DeleteStorageConfiguration"`

	Token ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token,omitempty"`
}

// DeleteStorageConfigurationResponse type
type DeleteStorageConfigurationResponse struct {
	XMLName xml.Name `xml:"DeleteStorageConfigurationResponse"`
}

// GetGeoLocation type
type GetGeoLocation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl GetGeoLocation"`
}

// GetGeoLocationResponse type
type GetGeoLocationResponse struct {
	XMLName xml.Name `xml:"GetGeoLocationResponse"`

	Location []LocationEntity `xml:"Location,omitempty"`
}

// SetGeoLocation type
type SetGeoLocation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl SetGeoLocation"`

	Location []LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location,omitempty"`
}

// SetGeoLocationResponse type
type SetGeoLocationResponse struct {
	XMLName xml.Name `xml:"SetGeoLocationResponse"`
}

// DeleteGeoLocation type
type DeleteGeoLocation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl DeleteGeoLocation"`

	Location []LocationEntity `xml:"http://www.onvif.org/ver10/device/wsdl Location,omitempty"`
}

// DeleteGeoLocationResponse type
type DeleteGeoLocationResponse struct {
	XMLName xml.Name `xml:"DeleteGeoLocationResponse"`
}

// Service type
type Service struct {

	// Namespace of the service being described. This parameter allows to match the service capabilities to the service. Note that only one set of capabilities is supported per namespace.
	Namespace AnyURI `xml:"http://www.onvif.org/ver10/schema Namespace,omitempty"`

	// The transport addresses where the service can be reached. The scheme and IP part shall match the one used in the request (i.e. the GetServices request).
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	Capabilities struct {
	} `xml:"Capabilities,omitempty"`

	// The version of the service (not the ONVIF core spec version).
	Version OnvifVersion `xml:"http://www.onvif.org/ver10/device/wsdl Version,omitempty"`
}

// DeviceServiceCapabilities type
type DeviceServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/device/wsdl Capabilities"`

	// Network capabilities.
	Network NetworkCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl Network,omitempty"`

	// Security capabilities.
	Security SecurityCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl Security,omitempty"`

	// System capabilities.
	System SystemCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl System,omitempty"`

	// Capabilities that do not fit in any of the other categories.
	Misc MiscCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl Misc,omitempty"`
}

// NetworkCapabilities type
type NetworkCapabilities struct {

	// Indicates support for IP filtering.

	IPFilter bool `xml:"http://www.onvif.org/ver10/device/wsdl IPFilter,attr,omitempty"`

	// Indicates support for zeroconf.

	ZeroConfiguration bool `xml:"http://www.onvif.org/ver10/device/wsdl ZeroConfiguration,attr,omitempty"`

	// Indicates support for IPv6.

	IPVersion6 bool `xml:"http://www.onvif.org/ver10/device/wsdl IPVersion6,attr,omitempty"`

	// Indicates support for dynamic DNS configuration.

	DynDNS bool `xml:"http://www.onvif.org/ver10/device/wsdl DynDNS,attr,omitempty"`

	// Indicates support for IEEE 802.11 configuration.

	Dot11Configuration bool `xml:"http://www.onvif.org/ver10/device/wsdl Dot11Configuration,attr,omitempty"`

	// Indicates the maximum number of Dot1X configurations supported by the device

	Dot1XConfigurations int32 `xml:"http://www.onvif.org/ver10/schema Dot1XConfigurations,attr,omitempty"`

	// Indicates support for retrieval of hostname from DHCP.

	HostnameFromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl HostnameFromDHCP,attr,omitempty"`

	// Maximum number of NTP servers supported by the devices SetNTP command.

	NTP int32 `xml:"http://www.onvif.org/ver10/schema NTP,attr,omitempty"`

	// Indicates support for Stateful IPv6 DHCP.

	DHCPv6 bool `xml:"http://www.onvif.org/ver10/device/wsdl DHCPv6,attr,omitempty"`
}

// SecurityCapabilities type
type SecurityCapabilities struct {

	// Indicates support for TLS 1.0.

	TLS1_0 bool `xml:"TLS1.0,attr,omitempty"`

	// Indicates support for TLS 1.1.

	TLS1_1 bool `xml:"TLS1.1,attr,omitempty"`

	// Indicates support for TLS 1.2.

	TLS1_2 bool `xml:"TLS1.2,attr,omitempty"`

	// Indicates support for onboard key generation.

	OnboardKeyGeneration bool `xml:"http://www.onvif.org/ver10/device/wsdl OnboardKeyGeneration,attr,omitempty"`

	// Indicates support for access policy configuration.

	AccessPolicyConfig bool `xml:"http://www.onvif.org/ver10/device/wsdl AccessPolicyConfig,attr,omitempty"`

	// Indicates support for the ONVIF default access policy.

	DefaultAccessPolicy bool `xml:"http://www.onvif.org/ver10/device/wsdl DefaultAccessPolicy,attr,omitempty"`

	// Indicates support for IEEE 802.1X configuration.

	Dot1X bool `xml:"http://www.onvif.org/ver10/device/wsdl Dot1X,attr,omitempty"`

	// Indicates support for remote user configuration. Used when accessing another device.

	RemoteUserHandling bool `xml:"http://www.onvif.org/ver10/device/wsdl RemoteUserHandling,attr,omitempty"`

	// Indicates support for WS-Security X.509 token.

	X_509Token bool `xml:"X.509Token,attr,omitempty"`

	// Indicates support for WS-Security SAML token.

	SAMLToken bool `xml:"http://www.onvif.org/ver10/device/wsdl SAMLToken,attr,omitempty"`

	// Indicates support for WS-Security Kerberos token.

	KerberosToken bool `xml:"http://www.onvif.org/ver10/device/wsdl KerberosToken,attr,omitempty"`

	// Indicates support for WS-Security Username token.

	UsernameToken bool `xml:"http://www.onvif.org/ver10/device/wsdl UsernameToken,attr,omitempty"`

	// Indicates support for WS over HTTP digest authenticated communication layer.

	HttpDigest bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpDigest,attr,omitempty"`

	// Indicates support for WS-Security REL token.

	RELToken bool `xml:"http://www.onvif.org/ver10/device/wsdl RELToken,attr,omitempty"`

	// EAP Methods supported by the device. The int values refer to the .

	SupportedEAPMethods EAPMethodTypes `xml:"http://www.onvif.org/ver10/device/wsdl SupportedEAPMethods,attr,omitempty"`

	// The maximum number of users that the device supports.

	MaxUsers int32 `xml:"http://www.onvif.org/ver10/schema MaxUsers,attr,omitempty"`

	// Maximum number of characters supported for the username by CreateUsers.

	MaxUserNameLength int32 `xml:"http://www.onvif.org/ver10/schema MaxUserNameLength,attr,omitempty"`

	// Maximum number of characters supported for the password by CreateUsers and SetUser.

	MaxPasswordLength int32 `xml:"http://www.onvif.org/ver10/schema MaxPasswordLength,attr,omitempty"`
}

// SystemCapabilities type
type SystemCapabilities struct {

	// Indicates support for WS Discovery resolve requests.

	DiscoveryResolve bool `xml:"http://www.onvif.org/ver10/device/wsdl DiscoveryResolve,attr,omitempty"`

	// Indicates support for WS-Discovery Bye.

	DiscoveryBye bool `xml:"http://www.onvif.org/ver10/device/wsdl DiscoveryBye,attr,omitempty"`

	// Indicates support for remote discovery.

	RemoteDiscovery bool `xml:"http://www.onvif.org/ver10/device/wsdl RemoteDiscovery,attr,omitempty"`

	// Indicates support for system backup through MTOM.

	SystemBackup bool `xml:"http://www.onvif.org/ver10/device/wsdl SystemBackup,attr,omitempty"`

	// Indicates support for retrieval of system logging through MTOM.

	SystemLogging bool `xml:"http://www.onvif.org/ver10/device/wsdl SystemLogging,attr,omitempty"`

	// Indicates support for firmware upgrade through MTOM.

	FirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/device/wsdl FirmwareUpgrade,attr,omitempty"`

	// Indicates support for firmware upgrade through HTTP.

	HttpFirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpFirmwareUpgrade,attr,omitempty"`

	// Indicates support for system backup through HTTP.

	HttpSystemBackup bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpSystemBackup,attr,omitempty"`

	// Indicates support for retrieval of system logging through HTTP.

	HttpSystemLogging bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpSystemLogging,attr,omitempty"`

	// Indicates support for retrieving support information through HTTP.

	HttpSupportInformation bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpSupportInformation,attr,omitempty"`

	// Indicates support for storage configuration interfaces.

	StorageConfiguration bool `xml:"http://www.onvif.org/ver10/device/wsdl StorageConfiguration,attr,omitempty"`

	// Indicates maximum number of storage configurations supported.

	MaxStorageConfigurations int32 `xml:"http://www.onvif.org/ver10/schema MaxStorageConfigurations,attr,omitempty"`

	// If present signals support for geo location. The value signals the supported number of entries.

	GeoLocationEntries int32 `xml:"http://www.onvif.org/ver10/schema GeoLocationEntries,attr,omitempty"`

	// List of supported automatic GeoLocation adjustment supported by the device. Valid items are defined by tds:AutoGeoMode.

	AutoGeo StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl AutoGeo,attr,omitempty"`

	// Enumerates the supported StorageTypes, see tds:StorageType.

	StorageTypesSupported StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl StorageTypesSupported,attr,omitempty"`
}

// MiscCapabilities type
type MiscCapabilities struct {

	// Lists of commands supported by SendAuxiliaryCommand.

	AuxiliaryCommands StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl AuxiliaryCommands,attr,omitempty"`
}

// UserCredential type
type UserCredential struct {

	// User name
	UserName string `xml:"http://www.onvif.org/ver10/device/wsdl UserName,omitempty"`

	// optional password
	Password string `xml:"http://www.onvif.org/ver10/device/wsdl Password,omitempty"`

	Extension struct {
	} `xml:"Extension,omitempty"`
}

// StorageConfigurationData type
type StorageConfigurationData struct {

	// local path
	LocalPath AnyURI `xml:"http://www.onvif.org/ver10/schema LocalPath,omitempty"`

	// Storage server address
	StorageUri AnyURI `xml:"http://www.onvif.org/ver10/schema StorageUri,omitempty"`

	// User credential for the storage server
	User UserCredential `xml:"http://www.onvif.org/ver10/device/wsdl User,omitempty"`

	Extension struct {
	} `xml:"Extension,omitempty"`

	// StorageType lists the acceptable values for type attribute

	Type string `xml:"type,attr,omitempty"`
}

// StorageConfiguration type
type StorageConfiguration struct {
	*DeviceEntity

	Data StorageConfigurationData `xml:"http://www.onvif.org/ver10/device/wsdl Data,omitempty"`
}

// Removed Base64Binary by fixgen.py

// HexBinary type
type HexBinary struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/05/xmlmime hexBinary"`

	Value []byte

	ContentType string `xml:"contentType,attr,omitempty"`
}

// FaultcodeEnum type
type FaultcodeEnum QName

const (
	FaultcodeEnumTnsDataEncodingUnknown FaultcodeEnum = "tns:DataEncodingUnknown"

	FaultcodeEnumTnsMustUnderstand FaultcodeEnum = "tns:MustUnderstand"

	FaultcodeEnumTnsReceiver FaultcodeEnum = "tns:Receiver"

	FaultcodeEnumTnsSender FaultcodeEnum = "tns:Sender"

	FaultcodeEnumTnsVersionMismatch FaultcodeEnum = "tns:VersionMismatch"
)

// NotUnderstood type
type NotUnderstood NotUnderstoodType

// Upgrade type
type Upgrade UpgradeType

// Envelope type
type Envelope struct {
	Header Header `xml:"Header,omitempty"`

	Body Body `xml:"Body,omitempty"`
}

// Header type
type Header struct {
}

// Body type
type Body struct {
}

// Fault type
type Fault struct {
	Code Faultcode `xml:"Code,omitempty"`

	Reason Faultreason `xml:"Reason,omitempty"`

	Node AnyURI `xml:"http://www.onvif.org/ver10/schema Node,omitempty"`

	Role AnyURI `xml:"http://www.onvif.org/ver10/schema Role,omitempty"`

	Detail Detail `xml:"Detail,omitempty"`
}

// Faultreason type
type Faultreason struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope faultreason"`

	Text []Reasontext `xml:"Text,omitempty"`
}

// Reasontext type
type Reasontext struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope reasontext"`

	Value string

	string `xml:",attr,omitempty"`
}

// Faultcode type
type Faultcode struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope faultcode"`

	Value FaultcodeEnum `xml:"Value,omitempty"`

	Subcode *Subcode `xml:"Subcode,omitempty"`
}

// Subcode type
type Subcode struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope subcode"`

	Value QName `xml:"http://www.onvif.org/ver10/schema Value,omitempty"`

	Subcode *Subcode `xml:"Subcode,omitempty"`
}

// Detail type
type Detail struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope detail"`
}

// NotUnderstoodType type
type NotUnderstoodType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope NotUnderstood"`

	Qname QName `xml:"qname,attr,omitempty"`
}

// SupportedEnvType type
type SupportedEnvType struct {
	Qname QName `xml:"qname,attr,omitempty"`
}

// UpgradeType type
type UpgradeType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Upgrade"`

	SupportedEnvelope []SupportedEnvType `xml:"SupportedEnvelope,omitempty"`
}

// RelationshipTypeOpenEnum type
type RelationshipTypeOpenEnum string

// RelationshipType type
type RelationshipType AnyURI

const (
	RelationshipTypeHttpwwww3org200508addressingreply RelationshipType = "http://www.w3.org/2005/08/addressing/reply"
)

// FaultCodesOpenEnumType type
type FaultCodesOpenEnumType string

// FaultCodesType type
type FaultCodesType QName

const (
	FaultCodesTypeTnsInvalidAddressingHeader FaultCodesType = "tns:InvalidAddressingHeader"

	FaultCodesTypeTnsInvalidAddress FaultCodesType = "tns:InvalidAddress"

	FaultCodesTypeTnsInvalidEPR FaultCodesType = "tns:InvalidEPR"

	FaultCodesTypeTnsInvalidCardinality FaultCodesType = "tns:InvalidCardinality"

	FaultCodesTypeTnsMissingAddressInEPR FaultCodesType = "tns:MissingAddressInEPR"

	FaultCodesTypeTnsDuplicateMessageID FaultCodesType = "tns:DuplicateMessageID"

	FaultCodesTypeTnsActionMismatch FaultCodesType = "tns:ActionMismatch"

	FaultCodesTypeTnsMessageAddressingHeaderRequired FaultCodesType = "tns:MessageAddressingHeaderRequired"

	FaultCodesTypeTnsDestinationUnreachable FaultCodesType = "tns:DestinationUnreachable"

	FaultCodesTypeTnsActionNotSupported FaultCodesType = "tns:ActionNotSupported"

	FaultCodesTypeTnsEndpointUnavailable FaultCodesType = "tns:EndpointUnavailable"
)

// EndpointReference type
type EndpointReference EndpointReferenceType

// Metadata type
type Metadata MetadataType

// MessageID type
type MessageID AttributedURIType

// RelatesTo type
type RelatesTo RelatesToType

// ReplyTo type
type ReplyTo EndpointReferenceType

// From type
type From EndpointReferenceType

// FaultTo type
type FaultTo EndpointReferenceType

// To type
type To AttributedURIType

// Action type
type Action AttributedURIType

// RetryAfter type
type RetryAfter AttributedUnsignedLongType

// ProblemHeaderQName type
type ProblemHeaderQName AttributedQNameType

// ProblemHeader type
type ProblemHeader AttributedAnyType

// ProblemIRI type
type ProblemIRI AttributedURIType

// ProblemAction type
type ProblemAction ProblemActionType

// EndpointReferenceType type
type EndpointReferenceType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing EndpointReference"`

	Address AttributedURIType `xml:"Address,omitempty"`

	ReferenceParameters ReferenceParametersType `xml:"ReferenceParameters,omitempty"`

	Metadata Metadata `xml:"Metadata,omitempty"`
}

// ReferenceParametersType type
type ReferenceParametersType struct {
}

// MetadataType type
type MetadataType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing Metadata"`
}

// RelatesToType type
type RelatesToType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing RelatesTo"`

	Value AnyURI

	RelationshipType RelationshipTypeOpenEnum `xml:"RelationshipType,attr,omitempty"`
}

// AttributedURIType type
type AttributedURIType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing MessageID"`

	Value AnyURI
}

// AttributedUnsignedLongType type
type AttributedUnsignedLongType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing RetryAfter"`

	Value uint64
}

// AttributedQNameType type
type AttributedQNameType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemHeaderQName"`

	Value QName
}

// AttributedAnyType type
type AttributedAnyType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemHeader"`
}

// ProblemActionType type
type ProblemActionType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemAction"`

	Action Action `xml:"Action,omitempty"`

	SoapAction AnyURI `xml:"http://www.onvif.org/ver10/schema SoapAction,omitempty"`
}

// BaseFault type
type BaseFault BaseFaultType

// BaseFaultType type
type BaseFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsrf/bf-2 BaseFault"`

	Timestamp string `xml:"http://www.onvif.org/ver10/schema Timestamp,omitempty"`

	Originator EndpointReferenceType `xml:"Originator,omitempty"`

	ErrorCode struct {
		Dialect AnyURI `xml:"dialect,attr,omitempty"`
	} `xml:"ErrorCode,omitempty"`

	Description []struct {
		Value string

		string `xml:",attr,omitempty"`
	} `xml:"Description,omitempty"`

	FaultCause struct {
	} `xml:"FaultCause,omitempty"`
}

// FullTopicExpression type
type FullTopicExpression string

// ConcreteTopicExpression type
type ConcreteTopicExpression string

// SimpleTopicExpression type
type SimpleTopicExpression QName

// TopicNamespace type
type TopicNamespace TopicNamespaceType

// TopicSet type
type TopicSet TopicSetType

// Documentation type
type Documentation struct {
}

// ExtensibleDocumented type
type ExtensibleDocumented struct {
	Documentation Documentation `xml:"documentation,omitempty"`
}

// QueryExpressionType type
type QueryExpressionType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 ProducerProperties"`

	Dialect AnyURI `xml:"http://www.onvif.org/ver10/schema Dialect,attr,omitempty"`
}

// TopicNamespaceType type
type TopicNamespaceType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 TopicNamespace"`

	*ExtensibleDocumented

	Topic []struct {
		*TopicType

		Parent ConcreteTopicExpression `xml:"parent,attr,omitempty"`
	} `xml:"Topic,omitempty"`

	Name NCName `xml:"name,attr,omitempty"`

	TargetNamespace AnyURI `xml:"targetNamespace,attr,omitempty"`

	Final bool `xml:"final,attr,omitempty"`
}

// TopicType type
type TopicType struct {
	*ExtensibleDocumented

	MessagePattern QueryExpressionType `xml:"MessagePattern,omitempty"`

	Topic []TopicType `xml:"Topic,omitempty"`

	Name NCName `xml:"name,attr,omitempty"`

	MessageTypes string `xml:"messageTypes,attr,omitempty"`

	Final bool `xml:"final,attr,omitempty"`
}

// TopicSetType type
type TopicSetType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 TopicSet"`

	*ExtensibleDocumented
}

// AbsoluteOrRelativeTimeType type
type AbsoluteOrRelativeTimeType string

// TopicExpression type
type TopicExpression TopicExpressionType

// FixedTopicSet type
type FixedTopicSet bool

// TopicExpressionDialect type
type TopicExpressionDialect AnyURI

// NotificationProducerRP type
type NotificationProducerRP struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationProducerRP"`

	TopicExpression []TopicExpression `xml:"TopicExpression,omitempty"`

	FixedTopicSet FixedTopicSet `xml:"FixedTopicSet,omitempty"`

	TopicExpressionDialect []TopicExpressionDialect `xml:"TopicExpressionDialect,omitempty"`

	TopicSet TopicSet `xml:"TopicSet,omitempty"`
}

// ConsumerReference type
type ConsumerReference EndpointReferenceType

// Filter type
type Filter FilterType

// SubscriptionPolicy type
type SubscriptionPolicy SubscriptionPolicyType

// CreationTime type
type CreationTime time.Time

// SubscriptionManagerRP type
type SubscriptionManagerRP struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscriptionManagerRP"`

	ConsumerReference ConsumerReference `xml:"ConsumerReference,omitempty"`

	Filter Filter `xml:"Filter,omitempty"`

	SubscriptionPolicy SubscriptionPolicy `xml:"SubscriptionPolicy,omitempty"`

	CreationTime CreationTime `xml:"CreationTime,omitempty"`
}

// SubscriptionReference type
type SubscriptionReference EndpointReferenceType

// Topic type
type Topic TopicExpressionType

// ProducerReference type
type ProducerReference EndpointReferenceType

// NotificationMessage type
type NotificationMessage NotificationMessageHolderType

// Notify type
type Notify struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Notify"`

	NotificationMessage []NotificationMessage `xml:"NotificationMessage,omitempty"`
}

// CurrentTime type
type CurrentTime time.Time

// TerminationTime type
type TerminationTime time.Time

// ProducerProperties type
type ProducerProperties QueryExpressionType

// MessageContent type
type MessageContent QueryExpressionType

// UseRaw type
type UseRaw struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UseRaw"`
}

// Subscribe type
type Subscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Subscribe"`

	ConsumerReference EndpointReferenceType `xml:"ConsumerReference,omitempty"`

	Filter FilterType `xml:"Filter,omitempty"`

	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
}

// SubscribeResponse type
type SubscribeResponse struct {
	XMLName xml.Name `xml:"SubscribeResponse"`

	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference,omitempty"`

	CurrentTime CurrentTime `xml:"CurrentTime,omitempty"`

	TerminationTime TerminationTime `xml:"TerminationTime,omitempty"`
}

// GetCurrentMessage type
type GetCurrentMessage struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetCurrentMessage"`

	Topic TopicExpressionType `xml:"Topic,omitempty"`
}

// GetCurrentMessageResponse type
type GetCurrentMessageResponse struct {
	XMLName xml.Name `xml:"GetCurrentMessageResponse"`
}

// SubscribeCreationFailedFault type
type SubscribeCreationFailedFault SubscribeCreationFailedFaultType

// InvalidFilterFault type
type InvalidFilterFault InvalidFilterFaultType

// TopicExpressionDialectUnknownFault type
type TopicExpressionDialectUnknownFault TopicExpressionDialectUnknownFaultType

// InvalidTopicExpressionFault type
type InvalidTopicExpressionFault InvalidTopicExpressionFaultType

// TopicNotSupportedFault type
type TopicNotSupportedFault TopicNotSupportedFaultType

// MultipleTopicsSpecifiedFault type
type MultipleTopicsSpecifiedFault MultipleTopicsSpecifiedFaultType

// InvalidProducerPropertiesExpressionFault type
type InvalidProducerPropertiesExpressionFault InvalidProducerPropertiesExpressionFaultType

// InvalidMessageContentExpressionFault type
type InvalidMessageContentExpressionFault InvalidMessageContentExpressionFaultType

// UnrecognizedPolicyRequestFault type
type UnrecognizedPolicyRequestFault UnrecognizedPolicyRequestFaultType

// UnsupportedPolicyRequestFault type
type UnsupportedPolicyRequestFault UnsupportedPolicyRequestFaultType

// NotifyMessageNotSupportedFault type
type NotifyMessageNotSupportedFault NotifyMessageNotSupportedFaultType

// UnacceptableInitialTerminationTimeFault type
type UnacceptableInitialTerminationTimeFault UnacceptableInitialTerminationTimeFaultType

// NoCurrentMessageOnTopicFault type
type NoCurrentMessageOnTopicFault NoCurrentMessageOnTopicFaultType

// GetMessages type
type GetMessages struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetMessages"`

	MaximumNumber NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema MaximumNumber,omitempty"`
}

// GetMessagesResponse type
type GetMessagesResponse struct {
	XMLName xml.Name `xml:"GetMessagesResponse"`

	NotificationMessage []NotificationMessage `xml:"NotificationMessage,omitempty"`
}

// DestroyPullPoint type
type DestroyPullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 DestroyPullPoint"`
}

// DestroyPullPointResponse type
type DestroyPullPointResponse struct {
	XMLName xml.Name `xml:"DestroyPullPointResponse"`
}

// UnableToGetMessagesFault type
type UnableToGetMessagesFault UnableToGetMessagesFaultType

// UnableToDestroyPullPointFault type
type UnableToDestroyPullPointFault UnableToDestroyPullPointFaultType

// CreatePullPoint type
type CreatePullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 CreatePullPoint"`
}

// CreatePullPointResponse type
type CreatePullPointResponse struct {
	XMLName xml.Name `xml:"CreatePullPointResponse"`

	PullPoint EndpointReferenceType `xml:"PullPoint,omitempty"`
}

// UnableToCreatePullPointFault type
type UnableToCreatePullPointFault UnableToCreatePullPointFaultType

// Renew type
type Renew struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Renew"`

	TerminationTime AbsoluteOrRelativeTimeType `xml:"TerminationTime,omitempty"`
}

// RenewResponse type
type RenewResponse struct {
	XMLName xml.Name `xml:"RenewResponse"`

	TerminationTime TerminationTime `xml:"TerminationTime,omitempty"`

	CurrentTime CurrentTime `xml:"CurrentTime,omitempty"`
}

// UnacceptableTerminationTimeFault type
type UnacceptableTerminationTimeFault UnacceptableTerminationTimeFaultType

// Unsubscribe type
type Unsubscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Unsubscribe"`
}

// UnsubscribeResponse type
type UnsubscribeResponse struct {
	XMLName xml.Name `xml:"UnsubscribeResponse"`
}

// UnableToDestroySubscriptionFault type
type UnableToDestroySubscriptionFault UnableToDestroySubscriptionFaultType

// PauseSubscription type
type PauseSubscription struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseSubscription"`
}

// PauseSubscriptionResponse type
type PauseSubscriptionResponse struct {
	XMLName xml.Name `xml:"PauseSubscriptionResponse"`
}

// ResumeSubscription type
type ResumeSubscription struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeSubscription"`
}

// ResumeSubscriptionResponse type
type ResumeSubscriptionResponse struct {
	XMLName xml.Name `xml:"ResumeSubscriptionResponse"`
}

// PauseFailedFault type
type PauseFailedFault PauseFailedFaultType

// ResumeFailedFault type
type ResumeFailedFault ResumeFailedFaultType

// Removed QueryExpressionType

// TopicExpressionType type
type TopicExpressionType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicExpression"`

	Dialect AnyURI `xml:"http://www.onvif.org/ver10/schema Dialect,attr,omitempty"`
}

// FilterType type
type FilterType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Filter"`
}

// SubscriptionPolicyType type
type SubscriptionPolicyType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscriptionPolicy"`
}

// NotificationMessageHolderType type
type NotificationMessageHolderType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationMessage"`

	SubscriptionReference SubscriptionReference `xml:"SubscriptionReference,omitempty"`

	Topic Topic `xml:"Topic,omitempty"`

	ProducerReference ProducerReference `xml:"ProducerReference,omitempty"`

	Message struct {
	} `xml:"Message,omitempty"`
}

// SubscribeCreationFailedFaultType type
type SubscribeCreationFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscribeCreationFailedFault"`

	*BaseFaultType
}

// InvalidFilterFaultType type
type InvalidFilterFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidFilterFault"`

	*BaseFaultType

	UnknownFilter []QName `xml:"http://www.onvif.org/ver10/schema UnknownFilter,omitempty"`
}

// TopicExpressionDialectUnknownFaultType type
type TopicExpressionDialectUnknownFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicExpressionDialectUnknownFault"`

	*BaseFaultType
}

// InvalidTopicExpressionFaultType type
type InvalidTopicExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidTopicExpressionFault"`

	*BaseFaultType
}

// TopicNotSupportedFaultType type
type TopicNotSupportedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicNotSupportedFault"`

	*BaseFaultType
}

// MultipleTopicsSpecifiedFaultType type
type MultipleTopicsSpecifiedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 MultipleTopicsSpecifiedFault"`

	*BaseFaultType
}

// InvalidProducerPropertiesExpressionFaultType type
type InvalidProducerPropertiesExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidProducerPropertiesExpressionFault"`

	*BaseFaultType
}

// InvalidMessageContentExpressionFaultType type
type InvalidMessageContentExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidMessageContentExpressionFault"`

	*BaseFaultType
}

// UnrecognizedPolicyRequestFaultType type
type UnrecognizedPolicyRequestFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnrecognizedPolicyRequestFault"`

	*BaseFaultType

	UnrecognizedPolicy []QName `xml:"http://www.onvif.org/ver10/schema UnrecognizedPolicy,omitempty"`
}

// UnsupportedPolicyRequestFaultType type
type UnsupportedPolicyRequestFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnsupportedPolicyRequestFault"`

	*BaseFaultType

	UnsupportedPolicy []QName `xml:"http://www.onvif.org/ver10/schema UnsupportedPolicy,omitempty"`
}

// NotifyMessageNotSupportedFaultType type
type NotifyMessageNotSupportedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotifyMessageNotSupportedFault"`

	*BaseFaultType
}

// UnacceptableInitialTerminationTimeFaultType type
type UnacceptableInitialTerminationTimeFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnacceptableInitialTerminationTimeFault"`

	*BaseFaultType

	MinimumTime string `xml:"http://www.onvif.org/ver10/schema MinimumTime,omitempty"`

	MaximumTime string `xml:"http://www.onvif.org/ver10/schema MaximumTime,omitempty"`
}

// NoCurrentMessageOnTopicFaultType type
type NoCurrentMessageOnTopicFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NoCurrentMessageOnTopicFault"`

	*BaseFaultType
}

// UnableToGetMessagesFaultType type
type UnableToGetMessagesFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToGetMessagesFault"`

	*BaseFaultType
}

// UnableToDestroyPullPointFaultType type
type UnableToDestroyPullPointFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToDestroyPullPointFault"`

	*BaseFaultType
}

// UnableToCreatePullPointFaultType type
type UnableToCreatePullPointFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToCreatePullPointFault"`

	*BaseFaultType
}

// UnacceptableTerminationTimeFaultType type
type UnacceptableTerminationTimeFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnacceptableTerminationTimeFault"`

	*BaseFaultType

	MinimumTime string `xml:"http://www.onvif.org/ver10/schema MinimumTime,omitempty"`

	MaximumTime string `xml:"http://www.onvif.org/ver10/schema MaximumTime,omitempty"`
}

// UnableToDestroySubscriptionFaultType type
type UnableToDestroySubscriptionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToDestroySubscriptionFault"`

	*BaseFaultType
}

// PauseFailedFaultType type
type PauseFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseFailedFault"`

	*BaseFaultType
}

// ResumeFailedFaultType type
type ResumeFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeFailedFault"`

	*BaseFaultType
}

// Include type
type Include struct {
	Href AnyURI `xml:"href,attr,omitempty"`
}

// Unique identifier for a physical or logical resource.
// Tokens should be assigned such that they are unique within a device. Tokens must be at least unique within its class.
// Length up to 64 characters.

// ReferenceToken type
type ReferenceToken string

// MoveStatus type
type MoveStatus string

const (
	MoveStatusIDLE MoveStatus = "IDLE"

	MoveStatusMOVING MoveStatus = "MOVING"

	MoveStatusUNKNOWN MoveStatus = "UNKNOWN"
)

// Entity type
type Entity string

const (
	EntityDevice Entity = "Device"

	EntityVideoSource Entity = "VideoSource"

	EntityAudioSource Entity = "AudioSource"
)

// IntRange type
type IntRange struct {
	Min int32 `xml:"http://www.onvif.org/ver10/schema Min,omitempty"`

	Max int32 `xml:"http://www.onvif.org/ver10/schema Max,omitempty"`
}

// Vector2D type
type Vector2D struct {
	X float32 `xml:"x,attr,omitempty"`

	Y float32 `xml:"y,attr,omitempty"`

	//
	// Pan/tilt coordinate space selector. The following options are defined:
	//

	Space AnyURI `xml:"space,attr,omitempty"`
}

// Vector1D type
type Vector1D struct {
	X float32 `xml:"x,attr,omitempty"`

	//
	// Zoom coordinate space selector. The following options are defined:
	//

	Space AnyURI `xml:"space,attr,omitempty"`
}

// PTZVector type
type PTZVector struct {

	// Pan and tilt position. The x component corresponds to pan and the y component to tilt.
	PanTilt Vector2D `xml:"http://www.onvif.org/ver10/schema PanTilt,omitempty"`

	//
	// A zoom position.
	//
	Zoom Vector1D `xml:"http://www.onvif.org/ver10/schema Zoom,omitempty"`
}

// PTZStatus type
type PTZStatus struct {

	//
	// Specifies the absolute position of the PTZ unit together with the Space references. The default absolute spaces of the corresponding PTZ configuration MUST be referenced within the Position element.
	//
	Position PTZVector `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`

	//
	// Indicates if the Pan/Tilt/Zoom device unit is currently moving, idle or in an unknown state.
	//
	MoveStatus PTZMoveStatus `xml:"http://www.onvif.org/ver10/schema MoveStatus,omitempty"`

	//
	// States a current PTZ error.
	//
	Error string `xml:"http://www.onvif.org/ver10/device/wsdl Error,omitempty"`

	//
	// Specifies the UTC time when this status was generated.
	//
	UtcTime string `xml:"http://www.onvif.org/ver10/schema UtcTime,omitempty"`
}

// PTZMoveStatus type
type PTZMoveStatus struct {
	PanTilt MoveStatus `xml:"http://www.onvif.org/ver10/schema PanTilt,omitempty"`

	Zoom MoveStatus `xml:"http://www.onvif.org/ver10/schema Zoom,omitempty"`
}

// Vector type
type Vector struct {
	X float32 `xml:"x,attr,omitempty"`

	Y float32 `xml:"y,attr,omitempty"`
}

// Rectangle type
type Rectangle struct {
	Bottom float32 `xml:"bottom,attr,omitempty"`

	Top float32 `xml:"top,attr,omitempty"`

	Right float32 `xml:"right,attr,omitempty"`

	Left float32 `xml:"left,attr,omitempty"`
}

// Polygon type
type Polygon struct {
	Point []Vector `xml:"http://www.onvif.org/ver10/schema Point,omitempty"`
}

// Color type
type Color struct {
	X float32 `xml:"http://www.onvif.org/ver10/schema X,attr,omitempty"`

	Y float32 `xml:"http://www.onvif.org/ver10/schema Y,attr,omitempty"`

	Z float32 `xml:"http://www.onvif.org/ver10/schema Z,attr,omitempty"`

	//
	// Acceptable values:
	//
	// If the Colorspace attribute is absent, YCbCr is implied.
	//
	// Deprecated values:
	//
	//

	Colorspace AnyURI `xml:"http://www.onvif.org/ver10/schema Colorspace,attr,omitempty"`
}

// Removed ColorCovariance by fixgen.py

// Transformation type
type Transformation struct {
	Translate Vector `xml:"http://www.onvif.org/ver10/schema Translate,omitempty"`

	Scale Vector `xml:"http://www.onvif.org/ver10/schema Scale,omitempty"`

	Extension TransformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// TransformationExtension type
type TransformationExtension struct {
}

// GeoLocation type
type GeoLocation struct {

	// East west location as angle.

	Lon float64 `xml:"lon,attr,omitempty"`

	// North south location as angle.

	Lat float64 `xml:"lat,attr,omitempty"`

	// Hight in meters above sea level.

	Elevation float32 `xml:"elevation,attr,omitempty"`
}

// GeoOrientation type
type GeoOrientation struct {

	// Rotation around the x axis.

	Roll float32 `xml:"roll,attr,omitempty"`

	// Rotation around the y axis.

	Pitch float32 `xml:"pitch,attr,omitempty"`

	// Rotation around the z axis.

	Yaw float32 `xml:"yaw,attr,omitempty"`
}

// LocalLocation type
type LocalLocation struct {

	// East west location as angle.

	X float32 `xml:"x,attr,omitempty"`

	// North south location as angle.

	Y float32 `xml:"y,attr,omitempty"`

	// Offset in meters from the sea level.

	Z float32 `xml:"z,attr,omitempty"`
}

// LocalOrientation type
type LocalOrientation struct {

	// Rotation around the y axis.

	Pan float32 `xml:"pan,attr,omitempty"`

	// Rotation around the z axis.

	Tilt float32 `xml:"http://www.onvif.org/ver10/schema tilt,attr,omitempty"`

	// Rotation around the x axis.

	Roll float32 `xml:"roll,attr,omitempty"`
}

// LocationEntity type
type LocationEntity struct {

	// Location on earth.
	GeoLocation GeoLocation `xml:"http://www.onvif.org/ver10/schema GeoLocation,omitempty"`

	// Orientation relative to earth.
	GeoOrientation GeoOrientation `xml:"http://www.onvif.org/ver10/schema GeoOrientation,omitempty"`

	// Indoor location offset.
	LocalLocation LocalLocation `xml:"http://www.onvif.org/ver10/schema LocalLocation,omitempty"`

	// Indoor orientation offset.
	LocalOrientation LocalOrientation `xml:"http://www.onvif.org/ver10/schema LocalOrientation,omitempty"`

	// Entity type the entry refers to, use a value from the tt:Entity enumeration.

	Entity string `xml:"http://www.onvif.org/ver10/device/wsdl Entity,attr,omitempty"`

	// Optional entity token.

	Token ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token,attr,omitempty"`

	// If this value is true the entity cannot be deleted.

	Fixed bool `xml:"http://www.onvif.org/ver10/device/wsdl Fixed,attr,omitempty"`

	// Optional reference to the XAddr of another devices DeviceManagement service.

	GeoSource AnyURI `xml:"http://www.onvif.org/ver10/schema GeoSource,attr,omitempty"`

	// If set the geo location is obtained internally.

	AutoGeo bool `xml:"http://www.onvif.org/ver10/device/wsdl AutoGeo,attr,omitempty"`
}

// User readable name. Length up to 64 characters.

// Name type
type Name string

// IntAttrList type
type IntAttrList []int32

// FloatAttrList type
type FloatAttrList []float32

// StringAttrList type
type StringAttrList []string

// ReferenceTokenList type
type ReferenceTokenList []ReferenceToken

// RotateMode type
type RotateMode string

const (

	// Enable the Rotate feature. Degree of rotation is specified Degree parameter.
	RotateModeOFF RotateMode = "OFF"

	// Disable the Rotate feature.
	RotateModeON RotateMode = "ON"

	// Rotate feature is automatically activated by the device.
	RotateModeAUTO RotateMode = "AUTO"
)

// SceneOrientationMode type
type SceneOrientationMode string

const (
	SceneOrientationModeMANUAL SceneOrientationMode = "MANUAL"

	SceneOrientationModeAUTO SceneOrientationMode = "AUTO"
)

// SceneOrientationOption type
type SceneOrientationOption string

const (
	SceneOrientationOptionBelow SceneOrientationOption = "Below"

	SceneOrientationOptionHorizon SceneOrientationOption = "Horizon"

	SceneOrientationOptionAbove SceneOrientationOption = "Above"
)

// Source view modes supported by device.

// ViewModes type
type ViewModes string

const (

	// Undewarped viewmode from device supporting fisheye lens.
	ViewModesTtFisheye ViewModes = "tt:Fisheye"

	// 360 degree panoramic view.
	ViewModesTt360Panorama ViewModes = "tt:360Panorama"

	// 180 degree panoramic view.
	ViewModesTt180Panorama ViewModes = "tt:180Panorama"

	// View mode combining four streams in single Quad, eg., applicable for devices supporting four heads.
	ViewModesTtQuad ViewModes = "tt:Quad"

	// Unaltered view from the sensor.
	ViewModesTtOriginal ViewModes = "tt:Original"

	// Viewmode combining the left side sensors, applicable for devices supporting multiple sensors.
	ViewModesTtLeftHalf ViewModes = "tt:LeftHalf"

	// Viewmode combining the right side sensors, applicable for devices supporting multiple sensors.
	ViewModesTtRightHalf ViewModes = "tt:RightHalf"

	// Dewarped view mode for device supporting fisheye lens.
	ViewModesTtDewarp ViewModes = "tt:Dewarp"
)

// VideoEncoding type
type VideoEncoding string

const (
	VideoEncodingJPEG VideoEncoding = "JPEG"

	VideoEncodingMPEG4 VideoEncoding = "MPEG4"

	VideoEncodingH264 VideoEncoding = "H264"
)

// Mpeg4Profile type
type Mpeg4Profile string

const (
	Mpeg4ProfileSP Mpeg4Profile = "SP"

	Mpeg4ProfileASP Mpeg4Profile = "ASP"
)

// H264Profile type
type H264Profile string

const (
	H264ProfileBaseline H264Profile = "Baseline"

	H264ProfileMain H264Profile = "Main"

	H264ProfileExtended H264Profile = "Extended"

	H264ProfileHigh H264Profile = "High"
)

// Video Media Subtypes as referenced by IANA (without the leading "video/" Video Media Type).  See also .

// VideoEncodingMimeNames type
type VideoEncodingMimeNames string

const (
	VideoEncodingMimeNamesJPEG VideoEncodingMimeNames = "JPEG"

	VideoEncodingMimeNamesMPV4ES VideoEncodingMimeNames = "MPV4-ES"

	VideoEncodingMimeNamesH264 VideoEncodingMimeNames = "H264"

	VideoEncodingMimeNamesH265 VideoEncodingMimeNames = "H265"
)

// VideoEncodingProfiles type
type VideoEncodingProfiles string

const (
	VideoEncodingProfilesSimple VideoEncodingProfiles = "Simple"

	VideoEncodingProfilesAdvancedSimple VideoEncodingProfiles = "AdvancedSimple"

	VideoEncodingProfilesBaseline VideoEncodingProfiles = "Baseline"

	VideoEncodingProfilesMain VideoEncodingProfiles = "Main"

	VideoEncodingProfilesMain10 VideoEncodingProfiles = "Main10"

	VideoEncodingProfilesExtended VideoEncodingProfiles = "Extended"

	VideoEncodingProfilesHigh VideoEncodingProfiles = "High"
)

// AudioEncoding type
type AudioEncoding string

const (
	AudioEncodingG711 AudioEncoding = "G711"

	AudioEncodingG726 AudioEncoding = "G726"

	AudioEncodingAAC AudioEncoding = "AAC"
)

// Audio Media Subtypes as referenced by IANA (without the leading "audio/" Audio Media Type).  See also .

// AudioEncodingMimeNames type
type AudioEncodingMimeNames string

const (
	AudioEncodingMimeNamesPCMU AudioEncodingMimeNames = "PCMU"

	AudioEncodingMimeNamesG726 AudioEncodingMimeNames = "G726"

	AudioEncodingMimeNamesMP4ALATM AudioEncodingMimeNames = "MP4A-LATM"

	AudioEncodingMimeNamesMpeg4generic AudioEncodingMimeNames = "mpeg4-generic"
)

// MetadataCompressionType type
type MetadataCompressionType string

const (
	MetadataCompressionTypeNone MetadataCompressionType = "None"

	MetadataCompressionTypeGZIP MetadataCompressionType = "GZIP"

	MetadataCompressionTypeEXI MetadataCompressionType = "EXI"
)

// StreamType type
type StreamType string

const (
	StreamTypeRTPUnicast StreamType = "RTP-Unicast"

	StreamTypeRTPMulticast StreamType = "RTP-Multicast"
)

// TransportProtocol type
type TransportProtocol string

const (
	TransportProtocolUDP TransportProtocol = "UDP"

	// This value is deprecated.
	TransportProtocolTCP TransportProtocol = "TCP"

	TransportProtocolRTSP TransportProtocol = "RTSP"

	TransportProtocolHTTP TransportProtocol = "HTTP"
)

// ScopeDefinition type
type ScopeDefinition string

const (
	ScopeDefinitionFixed ScopeDefinition = "Fixed"

	ScopeDefinitionConfigurable ScopeDefinition = "Configurable"
)

// DiscoveryMode type
type DiscoveryMode string

const (
	DiscoveryModeDiscoverable DiscoveryMode = "Discoverable"

	DiscoveryModeNonDiscoverable DiscoveryMode = "NonDiscoverable"
)

// NetworkInterfaceConfigPriority type
type NetworkInterfaceConfigPriority int32

// Duplex type
type Duplex string

const (
	DuplexFull Duplex = "Full"

	DuplexHalf Duplex = "Half"
)

// IANAIfTypes type
type IANAIfTypes int32

// IPv6DHCPConfiguration type
type IPv6DHCPConfiguration string

const (
	IPv6DHCPConfigurationAuto IPv6DHCPConfiguration = "Auto"

	IPv6DHCPConfigurationStateful IPv6DHCPConfiguration = "Stateful"

	IPv6DHCPConfigurationStateless IPv6DHCPConfiguration = "Stateless"

	IPv6DHCPConfigurationOff IPv6DHCPConfiguration = "Off"
)

// NetworkProtocolType type
type NetworkProtocolType string

const (
	NetworkProtocolTypeHTTP NetworkProtocolType = "HTTP"

	NetworkProtocolTypeHTTPS NetworkProtocolType = "HTTPS"

	NetworkProtocolTypeRTSP NetworkProtocolType = "RTSP"
)

// NetworkHostType type
type NetworkHostType string

const (
	NetworkHostTypeIPv4 NetworkHostType = "IPv4"

	NetworkHostTypeIPv6 NetworkHostType = "IPv6"

	NetworkHostTypeDNS NetworkHostType = "DNS"
)

// IPv4Address type
type IPv4Address string

// IPv6Address type
type IPv6Address string

// HwAddress type
type HwAddress string

// IPType type
type IPType string

const (
	IPTypeIPv4 IPType = "IPv4"

	IPTypeIPv6 IPType = "IPv6"
)

// DNSName type
type DNSName string

// Domain type
type Domain string

// IPAddressFilterType type
type IPAddressFilterType string

const (
	IPAddressFilterTypeAllow IPAddressFilterType = "Allow"

	IPAddressFilterTypeDeny IPAddressFilterType = "Deny"
)

// DynamicDNSType type
type DynamicDNSType string

const (
	DynamicDNSTypeNoUpdate DynamicDNSType = "NoUpdate"

	DynamicDNSTypeClientUpdates DynamicDNSType = "ClientUpdates"

	DynamicDNSTypeServerUpdates DynamicDNSType = "ServerUpdates"
)

// Dot11SSIDType type
type Dot11SSIDType []byte

// Dot11StationMode type
type Dot11StationMode string

const (
	Dot11StationModeAdhoc Dot11StationMode = "Ad-hoc"

	Dot11StationModeInfrastructure Dot11StationMode = "Infrastructure"

	Dot11StationModeExtended Dot11StationMode = "Extended"
)

// Dot11SecurityMode type
type Dot11SecurityMode string

const (
	Dot11SecurityModeNone Dot11SecurityMode = "None"

	Dot11SecurityModeWEP Dot11SecurityMode = "WEP"

	Dot11SecurityModePSK Dot11SecurityMode = "PSK"

	Dot11SecurityModeDot1X Dot11SecurityMode = "Dot1X"

	Dot11SecurityModeExtended Dot11SecurityMode = "Extended"
)

// Dot11Cipher type
type Dot11Cipher string

const (
	Dot11CipherCCMP Dot11Cipher = "CCMP"

	Dot11CipherTKIP Dot11Cipher = "TKIP"

	Dot11CipherAny Dot11Cipher = "Any"

	Dot11CipherExtended Dot11Cipher = "Extended"
)

// Dot11PSK type
type Dot11PSK []byte

// Dot11PSKPassphrase type
type Dot11PSKPassphrase string

// Dot11SignalStrength type
type Dot11SignalStrength string

const (
	Dot11SignalStrengthNone Dot11SignalStrength = "None"

	Dot11SignalStrengthVeryBad Dot11SignalStrength = "Very Bad"

	Dot11SignalStrengthBad Dot11SignalStrength = "Bad"

	Dot11SignalStrengthGood Dot11SignalStrength = "Good"

	Dot11SignalStrengthVeryGood Dot11SignalStrength = "Very Good"

	Dot11SignalStrengthExtended Dot11SignalStrength = "Extended"
)

// Dot11AuthAndMangementSuite type
type Dot11AuthAndMangementSuite string

const (
	Dot11AuthAndMangementSuiteNone Dot11AuthAndMangementSuite = "None"

	Dot11AuthAndMangementSuiteDot1X Dot11AuthAndMangementSuite = "Dot1X"

	Dot11AuthAndMangementSuitePSK Dot11AuthAndMangementSuite = "PSK"

	Dot11AuthAndMangementSuiteExtended Dot11AuthAndMangementSuite = "Extended"
)

// CapabilityCategory type
type CapabilityCategory string

const (
	CapabilityCategoryAll CapabilityCategory = "All"

	CapabilityCategoryAnalytics CapabilityCategory = "Analytics"

	CapabilityCategoryDevice CapabilityCategory = "Device"

	CapabilityCategoryEvents CapabilityCategory = "Events"

	CapabilityCategoryImaging CapabilityCategory = "Imaging"

	CapabilityCategoryMedia CapabilityCategory = "Media"

	CapabilityCategoryPTZ CapabilityCategory = "PTZ"
)

// Enumeration describing the available system log modes.

// SystemLogType type
type SystemLogType string

const (

	// Indicates that a system log is requested.
	SystemLogTypeSystem SystemLogType = "System"

	// Indicates that a access log is requested.
	SystemLogTypeAccess SystemLogType = "Access"
)

// Enumeration describing the available factory default modes.

// FactoryDefaultType type
type FactoryDefaultType string

const (

	// Indicates that a hard factory default is requested.
	FactoryDefaultTypeHard FactoryDefaultType = "Hard"

	// Indicates that a soft factory default is requested.
	FactoryDefaultTypeSoft FactoryDefaultType = "Soft"
)

// SetDateTimeType type
type SetDateTimeType string

const (

	// Indicates that the date and time are set manually.
	SetDateTimeTypeManual SetDateTimeType = "Manual"

	// Indicates that the date and time are set through NTP
	SetDateTimeTypeNTP SetDateTimeType = "NTP"
)

// UserLevel type
type UserLevel string

const (
	UserLevelAdministrator UserLevel = "Administrator"

	UserLevelOperator UserLevel = "Operator"

	UserLevelUser UserLevel = "User"

	UserLevelAnonymous UserLevel = "Anonymous"

	UserLevelExtended UserLevel = "Extended"
)

// RelayLogicalState type
type RelayLogicalState string

const (
	RelayLogicalStateActive RelayLogicalState = "active"

	RelayLogicalStateInactive RelayLogicalState = "inactive"
)

// RelayIdleState type
type RelayIdleState string

const (
	RelayIdleStateClosed RelayIdleState = "closed"

	RelayIdleStateOpen RelayIdleState = "open"
)

// RelayMode type
type RelayMode string

const (
	RelayModeMonostable RelayMode = "Monostable"

	RelayModeBistable RelayMode = "Bistable"
)

// DigitalIdleState type
type DigitalIdleState string

const (
	DigitalIdleStateClosed DigitalIdleState = "closed"

	DigitalIdleStateOpen DigitalIdleState = "open"
)

// EFlipMode type
type EFlipMode string

const (
	EFlipModeOFF EFlipMode = "OFF"

	EFlipModeON EFlipMode = "ON"

	EFlipModeExtended EFlipMode = "Extended"
)

// ReverseMode type
type ReverseMode string

const (
	ReverseModeOFF ReverseMode = "OFF"

	ReverseModeON ReverseMode = "ON"

	ReverseModeAUTO ReverseMode = "AUTO"

	ReverseModeExtended ReverseMode = "Extended"
)

// AuxiliaryData type
type AuxiliaryData string

// PTZPresetTourState type
type PTZPresetTourState string

const (
	PTZPresetTourStateIdle PTZPresetTourState = "Idle"

	PTZPresetTourStateTouring PTZPresetTourState = "Touring"

	PTZPresetTourStatePaused PTZPresetTourState = "Paused"

	PTZPresetTourStateExtended PTZPresetTourState = "Extended"
)

// PTZPresetTourDirection type
type PTZPresetTourDirection string

const (
	PTZPresetTourDirectionForward PTZPresetTourDirection = "Forward"

	PTZPresetTourDirectionBackward PTZPresetTourDirection = "Backward"

	PTZPresetTourDirectionExtended PTZPresetTourDirection = "Extended"
)

// PTZPresetTourOperation type
type PTZPresetTourOperation string

const (
	PTZPresetTourOperationStart PTZPresetTourOperation = "Start"

	PTZPresetTourOperationStop PTZPresetTourOperation = "Stop"

	PTZPresetTourOperationPause PTZPresetTourOperation = "Pause"

	PTZPresetTourOperationExtended PTZPresetTourOperation = "Extended"
)

// AutoFocusMode type
type AutoFocusMode string

const (
	AutoFocusModeAUTO AutoFocusMode = "AUTO"

	AutoFocusModeMANUAL AutoFocusMode = "MANUAL"
)

// AFModes type
type AFModes string

const (

	// Focus of a moving camera is updated only once after stopping a pan, tilt or zoom movement.
	AFModesOnceAfterMove AFModes = "OnceAfterMove"
)

// WideDynamicMode type
type WideDynamicMode string

const (
	WideDynamicModeOFF WideDynamicMode = "OFF"

	WideDynamicModeON WideDynamicMode = "ON"
)

// Enumeration describing the available backlight compenstation modes.

// BacklightCompensationMode type
type BacklightCompensationMode string

const (

	// Backlight compensation is disabled.
	BacklightCompensationModeOFF BacklightCompensationMode = "OFF"

	// Backlight compensation is enabled.
	BacklightCompensationModeON BacklightCompensationMode = "ON"
)

// ExposurePriority type
type ExposurePriority string

const (
	ExposurePriorityLowNoise ExposurePriority = "LowNoise"

	ExposurePriorityFrameRate ExposurePriority = "FrameRate"
)

// ExposureMode type
type ExposureMode string

const (
	ExposureModeAUTO ExposureMode = "AUTO"

	ExposureModeMANUAL ExposureMode = "MANUAL"
)

// Enabled type
type Enabled string

const (
	EnabledENABLED Enabled = "ENABLED"

	EnabledDISABLED Enabled = "DISABLED"
)

// WhiteBalanceMode type
type WhiteBalanceMode string

const (
	WhiteBalanceModeAUTO WhiteBalanceMode = "AUTO"

	WhiteBalanceModeMANUAL WhiteBalanceMode = "MANUAL"
)

// IrCutFilterMode type
type IrCutFilterMode string

const (
	IrCutFilterModeON IrCutFilterMode = "ON"

	IrCutFilterModeOFF IrCutFilterMode = "OFF"

	IrCutFilterModeAUTO IrCutFilterMode = "AUTO"
)

// ImageStabilizationMode type
type ImageStabilizationMode string

const (
	ImageStabilizationModeOFF ImageStabilizationMode = "OFF"

	ImageStabilizationModeON ImageStabilizationMode = "ON"

	ImageStabilizationModeAUTO ImageStabilizationMode = "AUTO"

	ImageStabilizationModeExtended ImageStabilizationMode = "Extended"
)

// IrCutFilterAutoBoundaryType type
type IrCutFilterAutoBoundaryType string

const (
	IrCutFilterAutoBoundaryTypeCommon IrCutFilterAutoBoundaryType = "Common"

	IrCutFilterAutoBoundaryTypeToOn IrCutFilterAutoBoundaryType = "ToOn"

	IrCutFilterAutoBoundaryTypeToOff IrCutFilterAutoBoundaryType = "ToOff"

	IrCutFilterAutoBoundaryTypeExtended IrCutFilterAutoBoundaryType = "Extended"
)

// ToneCompensationMode type
type ToneCompensationMode string

const (
	ToneCompensationModeOFF ToneCompensationMode = "OFF"

	ToneCompensationModeON ToneCompensationMode = "ON"

	ToneCompensationModeAUTO ToneCompensationMode = "AUTO"
)

// DefoggingMode type
type DefoggingMode string

const (
	DefoggingModeOFF DefoggingMode = "OFF"

	DefoggingModeON DefoggingMode = "ON"

	DefoggingModeAUTO DefoggingMode = "AUTO"
)

// TopicNamespaceLocation type
type TopicNamespaceLocation AnyURI

// PropertyOperation type
type PropertyOperation string

const (
	PropertyOperationInitialized PropertyOperation = "Initialized"

	PropertyOperationDeleted PropertyOperation = "Deleted"

	PropertyOperationChanged PropertyOperation = "Changed"
)

// Direction type
type Direction string

const (
	DirectionLeft Direction = "Left"

	DirectionRight Direction = "Right"

	DirectionAny Direction = "Any"
)

//
// Specifies a receiver connection mode.
//

// ReceiverMode type
type ReceiverMode string

const (

	// The receiver connects on demand, as required by consumers of the media streams.
	ReceiverModeAutoConnect ReceiverMode = "AutoConnect"

	// The receiver attempts to maintain a persistent connection to the configured endpoint.
	ReceiverModeAlwaysConnect ReceiverMode = "AlwaysConnect"

	// The receiver does not attempt to connect.
	ReceiverModeNeverConnect ReceiverMode = "NeverConnect"

	// This case should never happen.
	ReceiverModeUnknown ReceiverMode = "Unknown"
)

//
// Specifies the current connection state of the receiver.
//

// ReceiverState type
type ReceiverState string

const (

	// The receiver is not connected.
	ReceiverStateNotConnected ReceiverState = "NotConnected"

	// The receiver is attempting to connect.
	ReceiverStateConnecting ReceiverState = "Connecting"

	// The receiver is connected.
	ReceiverStateConnected ReceiverState = "Connected"

	// This case should never happen.
	ReceiverStateUnknown ReceiverState = "Unknown"
)

// ReceiverReference type
type ReceiverReference ReferenceToken

// RecordingReference type
type RecordingReference ReferenceToken

// TrackReference type
type TrackReference ReferenceToken

// Description type
type Description string

// XPathExpression type
type XPathExpression string

// SearchState type
type SearchState string

const (

	// The search is queued and not yet started.
	SearchStateQueued SearchState = "Queued"

	// The search is underway and not yet completed.
	SearchStateSearching SearchState = "Searching"

	// The search has been completed and no new results will be found.
	SearchStateCompleted SearchState = "Completed"

	// The state of the search is unknown. (This is not a valid response from GetSearchState.)
	SearchStateUnknown SearchState = "Unknown"
)

// JobToken type
type JobToken ReferenceToken

// RecordingStatus type
type RecordingStatus string

const (
	RecordingStatusInitiated RecordingStatus = "Initiated"

	RecordingStatusRecording RecordingStatus = "Recording"

	RecordingStatusStopped RecordingStatus = "Stopped"

	RecordingStatusRemoving RecordingStatus = "Removing"

	RecordingStatusRemoved RecordingStatus = "Removed"

	// This case should never happen.
	RecordingStatusUnknown RecordingStatus = "Unknown"
)

// TrackType type
type TrackType string

const (
	TrackTypeVideo TrackType = "Video"

	TrackTypeAudio TrackType = "Audio"

	TrackTypeMetadata TrackType = "Metadata"

	// Placeholder for future extension.
	TrackTypeExtended TrackType = "Extended"
)

// RecordingJobReference type
type RecordingJobReference ReferenceToken

// RecordingJobMode type
type RecordingJobMode string

// RecordingJobState type
type RecordingJobState string

// ModeOfOperation type
type ModeOfOperation string

const (
	ModeOfOperationIdle ModeOfOperation = "Idle"

	ModeOfOperationActive ModeOfOperation = "Active"

	// This case should never happen.
	ModeOfOperationUnknown ModeOfOperation = "Unknown"
)

//
// AudioClassType acceptable values are;
// gun_shot, scream, glass_breaking, tire_screech
//

// AudioClassType type
type AudioClassType string

// OSDType type
type OSDType string

const (
	OSDTypeText OSDType = "Text"

	OSDTypeImage OSDType = "Image"

	OSDTypeExtended OSDType = "Extended"
)

// StringItems type
type StringItems struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema StringItems"`

	Item []string `xml:"http://www.onvif.org/ver10/device/wsdl Item,omitempty"`
}

// StringList type
type StringList StringAttrList

// Message type
type Message struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Message"`

	// Token value pairs that triggered this message. Typically only one item is present.
	Source ItemList `xml:"http://www.onvif.org/ver10/schema Source,omitempty"`

	Key ItemList `xml:"http://www.onvif.org/ver10/schema Key,omitempty"`

	Data ItemList `xml:"http://www.onvif.org/ver10/schema Data,omitempty"`

	Extension MessageExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	UtcTime string `xml:"http://www.onvif.org/ver10/schema UtcTime,attr,omitempty"`

	PropertyOperation PropertyOperation `xml:"http://www.onvif.org/ver10/schema PropertyOperation,attr,omitempty"`
}

// DeviceEntity type
type DeviceEntity struct {

	// Unique identifier referencing the physical entity.

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

// IntRectangle type
type IntRectangle struct {
	X int32 `xml:"x,attr,omitempty"`

	Y int32 `xml:"y,attr,omitempty"`

	Width int32 `xml:"width,attr,omitempty"`

	Height int32 `xml:"height,attr,omitempty"`
}

// IntRectangleRange type
type IntRectangleRange struct {

	// Range of X-axis.
	XRange IntRange `xml:"http://www.onvif.org/ver10/schema XRange,omitempty"`

	// Range of Y-axis.
	YRange IntRange `xml:"http://www.onvif.org/ver10/schema YRange,omitempty"`

	// Range of width.
	WidthRange IntRange `xml:"http://www.onvif.org/ver10/schema WidthRange,omitempty"`

	// Range of height.
	HeightRange IntRange `xml:"http://www.onvif.org/ver10/schema HeightRange,omitempty"`
}

// FloatRange type
type FloatRange struct {
	Min float32 `xml:"http://www.onvif.org/ver10/schema Min,omitempty"`

	Max float32 `xml:"http://www.onvif.org/ver10/schema Max,omitempty"`
}

// DurationRange type
type DurationRange struct {
	Min Duration `xml:"http://www.onvif.org/ver10/schema Min,omitempty"`

	Max Duration `xml:"http://www.onvif.org/ver10/schema Max,omitempty"`
}

// IntList type
type IntList struct {
	Items []int32 `xml:"http://www.onvif.org/ver10/schema Items,omitempty"`
}

// Removed FloatList by fixgen.py

// Removed AnyHolder by fixgen.py

// VideoSource type
type VideoSource struct {
	*DeviceEntity

	// Frame rate in frames per second.
	Framerate float32 `xml:"http://www.onvif.org/ver10/schema Framerate,omitempty"`

	// Horizontal and vertical resolution
	Resolution VideoResolution `xml:"http://www.onvif.org/ver10/schema Resolution,omitempty"`

	// Optional configuration of the image sensor.
	Imaging ImagingSettings `xml:"http://www.onvif.org/ver10/schema Imaging,omitempty"`

	Extension VideoSourceExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoSourceExtension type
type VideoSourceExtension struct {

	// Optional configuration of the image sensor. To be used if imaging service 2.00 is supported.
	Imaging ImagingSettings20 `xml:"http://www.onvif.org/ver10/schema Imaging,omitempty"`

	Extension VideoSourceExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoSourceExtension2 type
type VideoSourceExtension2 struct {
}

// AudioSource type
type AudioSource struct {
	*DeviceEntity

	// number of available audio channels. (1: mono, 2: stereo)
	Channels int32 `xml:"http://www.onvif.org/ver10/schema Channels,omitempty"`
}

// Profile type
type Profile struct {

	// User readable name of the profile.
	Name Name `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	// Optional configuration of the Video input.
	VideoSourceConfiguration VideoSourceConfiguration `xml:"http://www.onvif.org/ver10/schema VideoSourceConfiguration,omitempty"`

	// Optional configuration of the Audio input.
	AudioSourceConfiguration AudioSourceConfiguration `xml:"http://www.onvif.org/ver10/schema AudioSourceConfiguration,omitempty"`

	// Optional configuration of the Video encoder.
	VideoEncoderConfiguration VideoEncoderConfiguration `xml:"http://www.onvif.org/ver10/schema VideoEncoderConfiguration,omitempty"`

	// Optional configuration of the Audio encoder.
	AudioEncoderConfiguration AudioEncoderConfiguration `xml:"http://www.onvif.org/ver10/schema AudioEncoderConfiguration,omitempty"`

	// Optional configuration of the video analytics module and rule engine.
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `xml:"http://www.onvif.org/ver10/schema VideoAnalyticsConfiguration,omitempty"`

	// Optional configuration of the pan tilt zoom unit.
	PTZConfiguration PTZConfiguration `xml:"http://www.onvif.org/ver10/schema PTZConfiguration,omitempty"`

	// Optional configuration of the metadata stream.
	MetadataConfiguration MetadataConfiguration `xml:"http://www.onvif.org/ver10/schema MetadataConfiguration,omitempty"`

	// Extensions defined in ONVIF 2.0
	Extension ProfileExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Unique identifier of the profile.

	Token ReferenceToken `xml:"token,attr,omitempty"`

	// A value of true signals that the profile cannot be deleted. Default is false.

	Fixed bool `xml:"fixed,attr,omitempty"`
}

// ProfileExtension type
type ProfileExtension struct {

	// Optional configuration of the Audio output.
	AudioOutputConfiguration AudioOutputConfiguration `xml:"http://www.onvif.org/ver10/schema AudioOutputConfiguration,omitempty"`

	// Optional configuration of the Audio decoder.
	AudioDecoderConfiguration AudioDecoderConfiguration `xml:"http://www.onvif.org/ver10/schema AudioDecoderConfiguration,omitempty"`

	Extension ProfileExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ProfileExtension2 type
type ProfileExtension2 struct {
}

// ConfigurationEntity type
type ConfigurationEntity struct {

	// User readable name. Length up to 64 characters.
	Name Name `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	// Number of internal references currently using this configuration.
	UseCount int32 `xml:"http://www.onvif.org/ver10/schema UseCount,omitempty"`

	// Token that uniquely references this configuration. Length up to 64 characters.

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

// VideoSourceConfiguration type
type VideoSourceConfiguration struct {
	*ConfigurationEntity

	// Reference to the physical input.
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl SourceToken,omitempty"`

	// Rectangle specifying the Video capturing area. The capturing area shall not be larger than the whole Video source area.
	Bounds IntRectangle `xml:"http://www.onvif.org/ver10/schema Bounds,omitempty"`

	Extension VideoSourceConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Readonly parameter signalling Source configuration's view mode, for devices supporting different view modes as defined in tt:viewModes.

	ViewMode string `xml:"http://www.onvif.org/ver10/device/wsdl ViewMode,attr,omitempty"`
}

// VideoSourceConfigurationExtension type
type VideoSourceConfigurationExtension struct {

	// Optional element to configure rotation of captured image.
	// What resolutions a device supports shall be unaffected by the Rotate parameters.
	// If a device is configured with Rotate=AUTO, the device shall take control over the Degree parameter and automatically update it so that a client can query current rotation.
	// The device shall automatically apply the same rotation to its pan/tilt control direction depending on the following condition:
	// if Reverse=AUTO in PTControlDirection or if the device doesn’t support Reverse in PTControlDirection
	//
	Rotate Rotate `xml:"http://www.onvif.org/ver10/schema Rotate,omitempty"`

	Extension VideoSourceConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoSourceConfigurationExtension2 type
type VideoSourceConfigurationExtension2 struct {

	// Optional element describing the geometric lens distortion. Multiple instances for future variable lens support.
	LensDescription []LensDescription `xml:"http://www.onvif.org/ver10/schema LensDescription,omitempty"`

	// Optional element describing the scene orientation in the camera’s field of view.
	SceneOrientation SceneOrientation `xml:"http://www.onvif.org/ver10/schema SceneOrientation,omitempty"`
}

// Rotate type
type Rotate struct {

	// Parameter to enable/disable Rotation feature.
	Mode RotateMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Optional parameter to configure how much degree of clockwise rotation of image  for On mode. Omitting this parameter for On mode means 180 degree rotation.
	Degree int32 `xml:"http://www.onvif.org/ver10/schema Degree,omitempty"`

	Extension RotateExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// RotateExtension type
type RotateExtension struct {
}

// LensProjection type
type LensProjection struct {

	// Angle of incidence.
	Angle float32 `xml:"http://www.onvif.org/ver10/schema Angle,omitempty"`

	// Mapping radius as a consequence of the emergent angle.
	Radius float32 `xml:"http://www.onvif.org/ver10/schema Radius,omitempty"`

	// Optional ray absorption at the given angle due to vignetting. A value of one means no absorption.
	Transmittance float32 `xml:"http://www.onvif.org/ver10/schema Transmittance,omitempty"`
}

// LensOffset type
type LensOffset struct {

	// Optional horizontal offset of the lens center in normalized coordinates.

	X float32 `xml:"x,attr,omitempty"`

	// Optional vertical offset of the lens center in normalized coordinates.

	Y float32 `xml:"y,attr,omitempty"`
}

// LensDescription type
type LensDescription struct {

	// Offset of the lens center to the imager center in normalized coordinates.
	Offset LensOffset `xml:"http://www.onvif.org/ver10/schema Offset,omitempty"`

	// Radial description of the projection characteristics. The resulting curve is defined by the B-Spline interpolation
	// over the given elements. The element for Radius zero shall not be provided. The projection points shall be ordered with ascending Radius.
	// Items outside the last projection Radius shall be assumed to be invisible (black).
	Projection []LensProjection `xml:"http://www.onvif.org/ver10/schema Projection,omitempty"`

	// Compensation of the x coordinate needed for the ONVIF normalized coordinate system.
	XFactor float32 `xml:"http://www.onvif.org/ver10/schema XFactor,omitempty"`

	// Optional focal length of the optical system.

	FocalLength float32 `xml:"http://www.onvif.org/ver10/schema FocalLength,attr,omitempty"`
}

// VideoSourceConfigurationOptions type
type VideoSourceConfigurationOptions struct {

	//
	// Supported range for the capturing area.
	// Device that does not support cropped streaming shall express BoundsRange option as mentioned below
	// BoundsRange->XRange and BoundsRange->YRange with same Min/Max values HeightRange and WidthRange Min/Max values same as VideoSource Height and Width Limits.
	//
	BoundsRange IntRectangleRange `xml:"http://www.onvif.org/ver10/schema BoundsRange,omitempty"`

	// List of physical inputs.
	VideoSourceTokensAvailable []ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl VideoSourceTokensAvailable,omitempty"`

	Extension VideoSourceConfigurationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Maximum number of profiles.

	MaximumNumberOfProfiles int32 `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfProfiles,attr,omitempty"`
}

// VideoSourceConfigurationOptionsExtension type
type VideoSourceConfigurationOptionsExtension struct {

	// Options of parameters for Rotation feature.
	Rotate RotateOptions `xml:"http://www.onvif.org/ver10/schema Rotate,omitempty"`

	Extension VideoSourceConfigurationOptionsExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoSourceConfigurationOptionsExtension2 type
type VideoSourceConfigurationOptionsExtension2 struct {

	// Scene orientation modes supported by the device for this configuration.
	SceneOrientationMode []SceneOrientationMode `xml:"http://www.onvif.org/ver10/schema SceneOrientationMode,omitempty"`
}

// RotateOptions type
type RotateOptions struct {

	// Supported options of Rotate mode parameter.
	Mode []RotateMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// List of supported degree value for rotation.
	DegreeList IntList `xml:"http://www.onvif.org/ver10/schema DegreeList,omitempty"`

	Extension RotateOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// After setting the rotation, if a device starts to reboot this value is true.
	// If a device can handle rotation setting without rebooting this value is false.

	Reboot bool `xml:"http://www.onvif.org/ver10/device/wsdl Reboot,attr,omitempty"`
}

// RotateOptionsExtension type
type RotateOptionsExtension struct {
}

// SceneOrientation type
type SceneOrientation struct {

	//
	// Parameter to assign the way the camera determines the scene orientation.
	//
	Mode SceneOrientationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// Assigned or determined scene orientation based on the Mode. When assigning the Mode to AUTO, this field
	// is optional and will be ignored by the device. When assigning the Mode to MANUAL, this field is required
	// and the device will return an InvalidArgs fault if missing.
	//
	Orientation string `xml:"http://www.onvif.org/ver10/device/wsdl Orientation,omitempty"`
}

// VideoEncoderConfiguration type
type VideoEncoderConfiguration struct {
	*ConfigurationEntity

	// Used video codec, either Jpeg, H.264 or Mpeg4
	Encoding VideoEncoding `xml:"http://www.onvif.org/ver10/schema Encoding,omitempty"`

	// Configured video resolution
	Resolution VideoResolution `xml:"http://www.onvif.org/ver10/schema Resolution,omitempty"`

	// Relative value for the video quantizers and the quality of the video. A high value within supported quality range means higher quality
	Quality float32 `xml:"http://www.onvif.org/ver10/schema Quality,omitempty"`

	// Optional element to configure rate control related parameters.
	RateControl VideoRateControl `xml:"http://www.onvif.org/ver10/schema RateControl,omitempty"`

	// Optional element to configure Mpeg4 related parameters.
	MPEG4 Mpeg4Configuration `xml:"http://www.onvif.org/ver10/schema MPEG4,omitempty"`

	// Optional element to configure H.264 related parameters.
	H264 H264Configuration `xml:"http://www.onvif.org/ver10/schema H264,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast,omitempty"`

	// The rtsp session timeout for the related video stream
	SessionTimeout Duration `xml:"http://www.onvif.org/ver10/schema SessionTimeout,omitempty"`

	//
	// A value of true indicates that frame rate is a fixed value rather than an upper limit,
	// and that the video encoder shall prioritize frame rate over all other adaptable
	// configuration values such as bitrate.  Default is false.
	//

	GuaranteedFrameRate bool `xml:"http://www.onvif.org/ver10/device/wsdl GuaranteedFrameRate,attr,omitempty"`
}

// VideoResolution type
type VideoResolution struct {

	// Number of the columns of the Video image.
	Width int32 `xml:"http://www.onvif.org/ver10/schema Width,omitempty"`

	// Number of the lines of the Video image.
	Height int32 `xml:"http://www.onvif.org/ver10/schema Height,omitempty"`
}

// VideoRateControl type
type VideoRateControl struct {

	// Maximum output framerate in fps. If an EncodingInterval is provided the resulting encoded framerate will be reduced by the given factor.
	FrameRateLimit int32 `xml:"http://www.onvif.org/ver10/schema FrameRateLimit,omitempty"`

	// Interval at which images are encoded and transmitted. (A value of 1 means that every frame is encoded, a value of 2 means that every 2nd frame is encoded ...)
	EncodingInterval int32 `xml:"http://www.onvif.org/ver10/schema EncodingInterval,omitempty"`

	// the maximum output bitrate in kbps
	BitrateLimit int32 `xml:"http://www.onvif.org/ver10/schema BitrateLimit,omitempty"`
}

// Mpeg4Configuration type
type Mpeg4Configuration struct {

	// Determines the interval in which the I-Frames will be coded. An entry of 1 indicates I-Frames are continuously generated. An entry of 2 indicates that every 2nd image is an I-Frame, and 3 only every 3rd frame, etc. The frames in between are coded as P or B Frames.
	GovLength int32 `xml:"http://www.onvif.org/ver10/schema GovLength,omitempty"`

	// the Mpeg4 profile, either simple profile (SP) or advanced simple profile (ASP)
	Mpeg4Profile Mpeg4Profile `xml:"http://www.onvif.org/ver10/schema Mpeg4Profile,omitempty"`
}

// H264Configuration type
type H264Configuration struct {

	// Group of Video frames length. Determines typically the interval in which the I-Frames will be coded. An entry of 1 indicates I-Frames are continuously generated. An entry of 2 indicates that every 2nd image is an I-Frame, and 3 only every 3rd frame, etc. The frames in between are coded as P or B Frames.
	GovLength int32 `xml:"http://www.onvif.org/ver10/schema GovLength,omitempty"`

	// the H.264 profile, either baseline, main, extended or high
	H264Profile H264Profile `xml:"http://www.onvif.org/ver10/schema H264Profile,omitempty"`
}

// Removed VideoEncoderConfigurationOptions by fixgen.py

// VideoEncoderOptionsExtension type
type VideoEncoderOptionsExtension struct {

	// Optional JPEG encoder settings ranges.
	JPEG JpegOptions2 `xml:"http://www.onvif.org/ver10/schema JPEG,omitempty"`

	// Optional MPEG-4 encoder settings ranges.
	MPEG4 Mpeg4Options2 `xml:"http://www.onvif.org/ver10/schema MPEG4,omitempty"`

	// Optional H.264 encoder settings ranges.
	H264 H264Options2 `xml:"http://www.onvif.org/ver10/schema H264,omitempty"`

	Extension VideoEncoderOptionsExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoEncoderOptionsExtension2 type
type VideoEncoderOptionsExtension2 struct {
}

// JpegOptions type
type JpegOptions struct {

	// List of supported image sizes.
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/schema ResolutionsAvailable,omitempty"`

	// Supported frame rate in fps (frames per second).
	FrameRateRange IntRange `xml:"http://www.onvif.org/ver10/schema FrameRateRange,omitempty"`

	// Supported encoding interval range. The encoding interval corresponds to the number of frames devided by the encoded frames. An encoding interval value of "1" means that all frames are encoded.
	EncodingIntervalRange IntRange `xml:"http://www.onvif.org/ver10/schema EncodingIntervalRange,omitempty"`
}

// JpegOptions2 type
type JpegOptions2 struct {
	*JpegOptions

	// Supported range of encoded bitrate in kbps.
	BitrateRange IntRange `xml:"http://www.onvif.org/ver10/schema BitrateRange,omitempty"`
}

// Mpeg4Options type
type Mpeg4Options struct {

	// List of supported image sizes.
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/schema ResolutionsAvailable,omitempty"`

	// Supported group of Video frames length. This value typically corresponds to the I-Frame distance.
	GovLengthRange IntRange `xml:"http://www.onvif.org/ver10/schema GovLengthRange,omitempty"`

	// Supported frame rate in fps (frames per second).
	FrameRateRange IntRange `xml:"http://www.onvif.org/ver10/schema FrameRateRange,omitempty"`

	// Supported encoding interval range. The encoding interval corresponds to the number of frames devided by the encoded frames. An encoding interval value of "1" means that all frames are encoded.
	EncodingIntervalRange IntRange `xml:"http://www.onvif.org/ver10/schema EncodingIntervalRange,omitempty"`

	// List of supported MPEG-4 profiles.
	Mpeg4ProfilesSupported []Mpeg4Profile `xml:"http://www.onvif.org/ver10/schema Mpeg4ProfilesSupported,omitempty"`
}

// Mpeg4Options2 type
type Mpeg4Options2 struct {
	*Mpeg4Options

	// Supported range of encoded bitrate in kbps.
	BitrateRange IntRange `xml:"http://www.onvif.org/ver10/schema BitrateRange,omitempty"`
}

// H264Options type
type H264Options struct {

	// List of supported image sizes.
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/schema ResolutionsAvailable,omitempty"`

	// Supported group of Video frames length. This value typically corresponds to the I-Frame distance.
	GovLengthRange IntRange `xml:"http://www.onvif.org/ver10/schema GovLengthRange,omitempty"`

	// Supported frame rate in fps (frames per second).
	FrameRateRange IntRange `xml:"http://www.onvif.org/ver10/schema FrameRateRange,omitempty"`

	// Supported encoding interval range. The encoding interval corresponds to the number of frames devided by the encoded frames. An encoding interval value of "1" means that all frames are encoded.
	EncodingIntervalRange IntRange `xml:"http://www.onvif.org/ver10/schema EncodingIntervalRange,omitempty"`

	// List of supported H.264 profiles.
	H264ProfilesSupported []H264Profile `xml:"http://www.onvif.org/ver10/schema H264ProfilesSupported,omitempty"`
}

// H264Options2 type
type H264Options2 struct {
	*H264Options

	// Supported range of encoded bitrate in kbps.
	BitrateRange IntRange `xml:"http://www.onvif.org/ver10/schema BitrateRange,omitempty"`
}

// Removed VideoEncoder2Configuration by fixgen.py

// VideoResolution2 type
type VideoResolution2 struct {

	// Number of the columns of the Video image.
	Width int32 `xml:"http://www.onvif.org/ver10/schema Width,omitempty"`

	// Number of the lines of the Video image.
	Height int32 `xml:"http://www.onvif.org/ver10/schema Height,omitempty"`
}

// Removed VideoRateControl2 by fixgen.py

// Removed VideoEncoder2ConfigurationOptions by fixgen.py

// AudioSourceConfiguration type
type AudioSourceConfiguration struct {
	*ConfigurationEntity

	// Token of the Audio Source the configuration applies to
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl SourceToken,omitempty"`
}

// Removed AudioSourceConfigurationOptions by fixgen.py

// Removed AudioSourceOptionsExtension by fixgen.py

// AudioEncoderConfiguration type
type AudioEncoderConfiguration struct {
	*ConfigurationEntity

	// Audio codec used for encoding the audio input (either G.711, G.726 or AAC)
	Encoding AudioEncoding `xml:"http://www.onvif.org/ver10/schema Encoding,omitempty"`

	// The output bitrate in kbps.
	Bitrate int32 `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// The output sample rate in kHz.
	SampleRate int32 `xml:"http://www.onvif.org/ver10/schema SampleRate,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast,omitempty"`

	// The rtsp session timeout for the related audio stream
	SessionTimeout Duration `xml:"http://www.onvif.org/ver10/schema SessionTimeout,omitempty"`
}

// AudioEncoderConfigurationOptions type
type AudioEncoderConfigurationOptions struct {

	// list of supported AudioEncoderConfigurations
	Options []AudioEncoderConfigurationOption `xml:"http://www.onvif.org/ver10/schema Options,omitempty"`
}

// AudioEncoderConfigurationOption type
type AudioEncoderConfigurationOption struct {

	// The enoding used for audio data (either G.711, G.726 or AAC)
	Encoding AudioEncoding `xml:"http://www.onvif.org/ver10/schema Encoding,omitempty"`

	// List of supported bitrates in kbps for the specified Encoding
	BitrateList IntList `xml:"http://www.onvif.org/ver10/schema BitrateList,omitempty"`

	// List of supported Sample Rates in kHz for the specified Encoding
	SampleRateList IntList `xml:"http://www.onvif.org/ver10/schema SampleRateList,omitempty"`
}

// Removed AudioEncoder2Configuration by fixgen.py

// Removed AudioEncoder2ConfigurationOptions by fixgen.py

// VideoAnalyticsConfiguration type
type VideoAnalyticsConfiguration struct {
	*ConfigurationEntity

	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration,omitempty"`

	RuleEngineConfiguration RuleEngineConfiguration `xml:"http://www.onvif.org/ver10/schema RuleEngineConfiguration,omitempty"`
}

// MetadataConfiguration type
type MetadataConfiguration struct {
	*ConfigurationEntity

	// optional element to configure which PTZ related data is to include in the metadata stream
	PTZStatus PTZFilter `xml:"http://www.onvif.org/ver10/schema PTZStatus,omitempty"`

	// Optional element to configure the streaming of events. A client might be interested in receiving all,
	// none or some of the events produced by the device:
	//
	Events EventSubscription `xml:"http://www.onvif.org/ver10/schema Events,omitempty"`

	// Defines whether the streamed metadata will include metadata from the analytics engines (video, cell motion, audio etc.)
	Analytics bool `xml:"http://www.onvif.org/ver10/device/wsdl Analytics,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast,omitempty"`

	// The rtsp session timeout for the related audio stream (when using Media2 Service, this value is deprecated and ignored)
	SessionTimeout Duration `xml:"http://www.onvif.org/ver10/schema SessionTimeout,omitempty"`

	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration,omitempty"`

	Extension MetadataConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Optional parameter to configure compression type of Metadata payload. Use values from enumeration MetadataCompressionType.

	CompressionType string `xml:"http://www.onvif.org/ver10/device/wsdl CompressionType,attr,omitempty"`

	// Optional parameter to configure if the metadata stream shall contain the Geo Location coordinates of each target.

	GeoLocation bool `xml:"http://www.onvif.org/ver10/device/wsdl GeoLocation,attr,omitempty"`
}

// MetadataConfigurationExtension type
type MetadataConfigurationExtension struct {
}

// PTZFilter type
type PTZFilter struct {

	// True if the metadata stream shall contain the PTZ status (IDLE, MOVING or UNKNOWN)
	Status bool `xml:"http://www.onvif.org/ver10/device/wsdl Status,omitempty"`

	// True if the metadata stream shall contain the PTZ position
	Position bool `xml:"http://www.onvif.org/ver10/device/wsdl Position,omitempty"`
}

// EventSubscription type
type EventSubscription struct {
	Filter FilterType `xml:"Filter,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
}

// MetadataConfigurationOptions type
type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions `xml:"http://www.onvif.org/ver10/schema PTZStatusFilterOptions,omitempty"`

	Extension MetadataConfigurationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// True if the device is able to stream the Geo Located positions of each target.

	GeoLocation bool `xml:"http://www.onvif.org/ver10/device/wsdl GeoLocation,attr,omitempty"`
}

// MetadataConfigurationOptionsExtension type
type MetadataConfigurationOptionsExtension struct {

	// List of supported metadata compression type. Its options shall be chosen from tt:MetadataCompressionType.
	CompressionType []string `xml:"http://www.onvif.org/ver10/device/wsdl CompressionType,omitempty"`

	Extension MetadataConfigurationOptionsExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// MetadataConfigurationOptionsExtension2 type
type MetadataConfigurationOptionsExtension2 struct {
}

// PTZStatusFilterOptions type
type PTZStatusFilterOptions struct {

	// True if the device is able to stream pan or tilt status information.
	PanTiltStatusSupported bool `xml:"http://www.onvif.org/ver10/device/wsdl PanTiltStatusSupported,omitempty"`

	// True if the device is able to stream zoom status inforamtion.
	ZoomStatusSupported bool `xml:"http://www.onvif.org/ver10/device/wsdl ZoomStatusSupported,omitempty"`

	// True if the device is able to stream the pan or tilt position.
	PanTiltPositionSupported bool `xml:"http://www.onvif.org/ver10/device/wsdl PanTiltPositionSupported,omitempty"`

	// True if the device is able to stream zoom position information.
	ZoomPositionSupported bool `xml:"http://www.onvif.org/ver10/device/wsdl ZoomPositionSupported,omitempty"`

	Extension PTZStatusFilterOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZStatusFilterOptionsExtension type
type PTZStatusFilterOptionsExtension struct {
}

// VideoOutput type
type VideoOutput struct {
	*DeviceEntity

	Layout Layout `xml:"http://www.onvif.org/ver10/schema Layout,omitempty"`

	// Resolution of the display in Pixel.
	Resolution VideoResolution `xml:"http://www.onvif.org/ver10/schema Resolution,omitempty"`

	// Refresh rate of the display in Hertz.
	RefreshRate float32 `xml:"http://www.onvif.org/ver10/schema RefreshRate,omitempty"`

	// Aspect ratio of the display as physical extent of width divided by height.
	AspectRatio float32 `xml:"http://www.onvif.org/ver10/schema AspectRatio,omitempty"`

	Extension VideoOutputExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoOutputExtension type
type VideoOutputExtension struct {
}

// Removed VideoOutputConfiguration by fixgen.py

// Removed VideoOutputConfigurationOptions by fixgen.py

// VideoDecoderConfigurationOptions type
type VideoDecoderConfigurationOptions struct {

	// If the device is able to decode Jpeg streams this element describes the supported codecs and configurations
	JpegDecOptions JpegDecOptions `xml:"http://www.onvif.org/ver10/schema JpegDecOptions,omitempty"`

	// If the device is able to decode H.264 streams this element describes the supported codecs and configurations
	H264DecOptions H264DecOptions `xml:"http://www.onvif.org/ver10/schema H264DecOptions,omitempty"`

	// If the device is able to decode Mpeg4 streams this element describes the supported codecs and configurations
	Mpeg4DecOptions Mpeg4DecOptions `xml:"http://www.onvif.org/ver10/schema Mpeg4DecOptions,omitempty"`

	Extension VideoDecoderConfigurationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// H264DecOptions type
type H264DecOptions struct {

	// List of supported H.264 Video Resolutions
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/schema ResolutionsAvailable,omitempty"`

	// List of supported H264 Profiles (either baseline, main, extended or high)
	SupportedH264Profiles []H264Profile `xml:"http://www.onvif.org/ver10/schema SupportedH264Profiles,omitempty"`

	// Supported H.264 bitrate range in kbps
	SupportedInputBitrate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedInputBitrate,omitempty"`

	// Supported H.264 framerate range in fps
	SupportedFrameRate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedFrameRate,omitempty"`
}

// JpegDecOptions type
type JpegDecOptions struct {

	// List of supported Jpeg Video Resolutions
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/schema ResolutionsAvailable,omitempty"`

	// Supported Jpeg bitrate range in kbps
	SupportedInputBitrate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedInputBitrate,omitempty"`

	// Supported Jpeg framerate range in fps
	SupportedFrameRate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedFrameRate,omitempty"`
}

// Mpeg4DecOptions type
type Mpeg4DecOptions struct {

	// List of supported Mpeg4 Video Resolutions
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/schema ResolutionsAvailable,omitempty"`

	// List of supported Mpeg4 Profiles (either SP or ASP)
	SupportedMpeg4Profiles []Mpeg4Profile `xml:"http://www.onvif.org/ver10/schema SupportedMpeg4Profiles,omitempty"`

	// Supported Mpeg4 bitrate range in kbps
	SupportedInputBitrate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedInputBitrate,omitempty"`

	// Supported Mpeg4 framerate range in fps
	SupportedFrameRate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedFrameRate,omitempty"`
}

// VideoDecoderConfigurationOptionsExtension type
type VideoDecoderConfigurationOptionsExtension struct {
}

// AudioOutput type
type AudioOutput struct {
	*DeviceEntity
}

// AudioOutputConfiguration type
type AudioOutputConfiguration struct {
	*ConfigurationEntity

	// Token of the phsycial Audio output.
	OutputToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl OutputToken,omitempty"`

	//
	// An audio channel MAY support different types of audio transmission. While for full duplex
	// operation no special handling is required, in half duplex operation the transmission direction
	// needs to be switched.
	// The optional SendPrimacy parameter inside the AudioOutputConfiguration indicates which
	// direction is currently active. An NVC can switch between different modes by setting the
	// AudioOutputConfiguration.
	// The following modes for the Send-Primacy are defined:
	// Acoustic echo cancellation is out of ONVIF scope.
	SendPrimacy AnyURI `xml:"http://www.onvif.org/ver10/schema SendPrimacy,omitempty"`

	// Volume setting of the output. The applicable range is defined via the option AudioOutputOptions.OutputLevelRange.
	OutputLevel int32 `xml:"http://www.onvif.org/ver10/schema OutputLevel,omitempty"`
}

// Removed AudioOutputConfigurationOptions by fixgen.py

// AudioDecoderConfiguration type
type AudioDecoderConfiguration struct {
	*ConfigurationEntity
}

// AudioDecoderConfigurationOptions type
type AudioDecoderConfigurationOptions struct {

	// If the device is able to decode AAC encoded audio this section describes the supported configurations
	AACDecOptions AACDecOptions `xml:"http://www.onvif.org/ver10/schema AACDecOptions,omitempty"`

	// If the device is able to decode G711 encoded audio this section describes the supported configurations
	G711DecOptions G711DecOptions `xml:"http://www.onvif.org/ver10/schema G711DecOptions,omitempty"`

	// If the device is able to decode G726 encoded audio this section describes the supported configurations
	G726DecOptions G726DecOptions `xml:"http://www.onvif.org/ver10/schema G726DecOptions,omitempty"`

	Extension AudioDecoderConfigurationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// G711DecOptions type
type G711DecOptions struct {

	// List of supported bitrates in kbps
	Bitrate IntList `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// List of supported sample rates in kHz
	SampleRateRange IntList `xml:"http://www.onvif.org/ver10/schema SampleRateRange,omitempty"`
}

// AACDecOptions type
type AACDecOptions struct {

	// List of supported bitrates in kbps
	Bitrate IntList `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// List of supported sample rates in kHz
	SampleRateRange IntList `xml:"http://www.onvif.org/ver10/schema SampleRateRange,omitempty"`
}

// G726DecOptions type
type G726DecOptions struct {

	// List of supported bitrates in kbps
	Bitrate IntList `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// List of supported sample rates in kHz
	SampleRateRange IntList `xml:"http://www.onvif.org/ver10/schema SampleRateRange,omitempty"`
}

// AudioDecoderConfigurationOptionsExtension type
type AudioDecoderConfigurationOptionsExtension struct {
}

// MulticastConfiguration type
type MulticastConfiguration struct {

	// The multicast address (if this address is set to 0 no multicast streaming is enaled)
	Address IPAddress `xml:"http://www.onvif.org/ver10/device/wsdl Address,omitempty"`

	// The RTP mutlicast destination port. A device may support RTCP. In this case the port value shall be even to allow the corresponding RTCP stream to be mapped to the next higher (odd) destination port number as defined in the RTSP specification.
	Port int32 `xml:"http://www.onvif.org/ver10/schema Port,omitempty"`

	// In case of IPv6 the TTL value is assumed as the hop limit. Note that for IPV6 and administratively scoped IPv4 multicast the primary use for hop limit / TTL is to prevent packets from (endlessly) circulating and not limiting scope. In these cases the address contains the scope.
	TTL int32 `xml:"http://www.onvif.org/ver10/schema TTL,omitempty"`

	// Read only property signalling that streaming is persistant. Use the methods StartMulticastStreaming and StopMulticastStreaming to switch its state.
	AutoStart bool `xml:"http://www.onvif.org/ver10/device/wsdl AutoStart,omitempty"`
}

// StreamSetup type
type StreamSetup struct {

	// Defines if a multicast or unicast stream is requested
	Stream StreamType `xml:"http://www.onvif.org/ver10/schema Stream,omitempty"`

	Transport Transport `xml:"http://www.onvif.org/ver10/schema Transport,omitempty"`
}

// Transport type
type Transport struct {

	// Defines the network protocol for streaming, either UDP=RTP/UDP, RTSP=RTP/RTSP/TCP or HTTP=RTP/RTSP/HTTP/TCP
	Protocol TransportProtocol `xml:"http://www.onvif.org/ver10/schema Protocol,omitempty"`

	// Optional element to describe further tunnel options. This element is normally not needed
	Tunnel *Transport `xml:"http://www.onvif.org/ver10/schema Tunnel,omitempty"`
}

// Removed MediaUri by fixgen.py

// Scope type
type Scope struct {

	// Indicates if the scope is fixed or configurable.
	ScopeDef ScopeDefinition `xml:"http://www.onvif.org/ver10/schema ScopeDef,omitempty"`

	// Scope item URI.
	ScopeItem AnyURI `xml:"http://www.onvif.org/ver10/schema ScopeItem,omitempty"`
}

// NetworkInterface type
type NetworkInterface struct {
	*DeviceEntity

	// Indicates whether or not an interface is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// Network interface information
	Info NetworkInterfaceInfo `xml:"http://www.onvif.org/ver10/schema Info,omitempty"`

	// Link configuration.
	Link NetworkInterfaceLink `xml:"http://www.onvif.org/ver10/schema Link,omitempty"`

	// IPv4 network interface configuration.
	IPv4 IPv4NetworkInterface `xml:"http://www.onvif.org/ver10/schema IPv4,omitempty"`

	// IPv6 network interface configuration.
	IPv6 IPv6NetworkInterface `xml:"http://www.onvif.org/ver10/schema IPv6,omitempty"`

	Extension NetworkInterfaceExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkInterfaceExtension type
type NetworkInterfaceExtension struct {
	InterfaceType IANAIfTypes `xml:"InterfaceType,omitempty"`

	// Extension point prepared for future 802.3 configuration.
	Dot3 []Dot3Configuration `xml:"http://www.onvif.org/ver10/schema Dot3,omitempty"`

	Dot11 []Dot11Configuration `xml:"http://www.onvif.org/ver10/schema Dot11,omitempty"`

	Extension NetworkInterfaceExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// Dot3Configuration type
type Dot3Configuration struct {
}

// NetworkInterfaceExtension2 type
type NetworkInterfaceExtension2 struct {
}

// NetworkInterfaceLink type
type NetworkInterfaceLink struct {

	// Configured link settings.
	AdminSettings NetworkInterfaceConnectionSetting `xml:"http://www.onvif.org/ver10/schema AdminSettings,omitempty"`

	// Current active link settings.
	OperSettings NetworkInterfaceConnectionSetting `xml:"http://www.onvif.org/ver10/schema OperSettings,omitempty"`

	// Integer indicating interface type, for example: 6 is ethernet.
	InterfaceType IANAIfTypes `xml:"InterfaceType,omitempty"`
}

// NetworkInterfaceConnectionSetting type
type NetworkInterfaceConnectionSetting struct {

	// Auto negotiation on/off.
	AutoNegotiation bool `xml:"http://www.onvif.org/ver10/device/wsdl AutoNegotiation,omitempty"`

	// Speed.
	Speed int32 `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`

	// Duplex type, Half or Full.
	Duplex Duplex `xml:"http://www.onvif.org/ver10/schema Duplex,omitempty"`
}

// NetworkInterfaceInfo type
type NetworkInterfaceInfo struct {

	// Network interface name, for example eth0.
	Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`

	// Network interface MAC address.
	HwAddress HwAddress `xml:"http://www.onvif.org/ver10/schema HwAddress,omitempty"`

	// Maximum transmission unit.
	MTU int32 `xml:"http://www.onvif.org/ver10/schema MTU,omitempty"`
}

// IPv6NetworkInterface type
type IPv6NetworkInterface struct {

	// Indicates whether or not IPv6 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// IPv6 configuration.
	Config IPv6Configuration `xml:"http://www.onvif.org/ver10/schema Config,omitempty"`
}

// IPv4NetworkInterface type
type IPv4NetworkInterface struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// IPv4 configuration.
	Config IPv4Configuration `xml:"http://www.onvif.org/ver10/schema Config,omitempty"`
}

// IPv4Configuration type
type IPv4Configuration struct {

	// List of manually added IPv4 addresses.
	Manual []PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// Link local address.
	LinkLocal PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema LinkLocal,omitempty"`

	// IPv4 address configured by using DHCP.
	FromDHCP PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema FromDHCP,omitempty"`

	// Indicates whether or not DHCP is used.
	DHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl DHCP,omitempty"`
}

// IPv6Configuration type
type IPv6Configuration struct {

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"http://www.onvif.org/ver10/device/wsdl AcceptRouterAdvert,omitempty"`

	// DHCP configuration.
	DHCP IPv6DHCPConfiguration `xml:"http://www.onvif.org/ver10/schema DHCP,omitempty"`

	// List of manually entered IPv6 addresses.
	Manual []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// List of link local IPv6 addresses.
	LinkLocal []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema LinkLocal,omitempty"`

	// List of IPv6 addresses configured by using DHCP.
	FromDHCP []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema FromDHCP,omitempty"`

	// List of IPv6 addresses configured by using router advertisment.
	FromRA []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema FromRA,omitempty"`

	Extension IPv6ConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IPv6ConfigurationExtension type
type IPv6ConfigurationExtension struct {
}

// NetworkProtocol type
type NetworkProtocol struct {

	// Network protocol type string.
	Name NetworkProtocolType `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	// Indicates if the protocol is enabled or not.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// The port that is used by the protocol.
	Port []int32 `xml:"http://www.onvif.org/ver10/schema Port,omitempty"`

	Extension NetworkProtocolExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkProtocolExtension type
type NetworkProtocolExtension struct {
}

// NetworkHost type
type NetworkHost struct {

	// Network host type: IPv4, IPv6 or DNS.
	Type NetworkHostType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// IPv4 address.
	IPv4Address IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv4Address,omitempty"`

	// IPv6 address.
	IPv6Address IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv6Address,omitempty"`

	// DNS name.
	DNSname DNSName `xml:"http://www.onvif.org/ver10/device/wsdl DNSname,omitempty"`

	Extension NetworkHostExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkHostExtension type
type NetworkHostExtension struct {
}

// IPAddress type
type IPAddress struct {

	// Indicates if the address is an IPv4 or IPv6 address.
	Type IPType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// IPv4 address.
	IPv4Address IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv4Address,omitempty"`

	// IPv6 address
	IPv6Address IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv6Address,omitempty"`
}

// PrefixedIPv4Address type
type PrefixedIPv4Address struct {

	// IPv4 address
	Address IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl Address,omitempty"`

	// Prefix/submask length
	PrefixLength int32 `xml:"http://www.onvif.org/ver10/schema PrefixLength,omitempty"`
}

// PrefixedIPv6Address type
type PrefixedIPv6Address struct {

	// IPv6 address
	Address IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl Address,omitempty"`

	// Prefix/submask length
	PrefixLength int32 `xml:"http://www.onvif.org/ver10/schema PrefixLength,omitempty"`
}

// HostnameInformation type
type HostnameInformation struct {

	// Indicates whether the hostname is obtained from DHCP or not.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP,omitempty"`

	// Indicates the hostname.
	Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`

	Extension HostnameInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// HostnameInformationExtension type
type HostnameInformationExtension struct {
}

// DNSInformation type
type DNSInformation struct {

	// Indicates whether or not DNS information is retrieved from DHCP.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP,omitempty"`

	// Search domain.
	SearchDomain []string `xml:"http://www.onvif.org/ver10/device/wsdl SearchDomain,omitempty"`

	// List of DNS addresses received from DHCP.
	DNSFromDHCP []IPAddress `xml:"http://www.onvif.org/ver10/device/wsdl DNSFromDHCP,omitempty"`

	// List of manually entered DNS addresses.
	DNSManual []IPAddress `xml:"http://www.onvif.org/ver10/device/wsdl DNSManual,omitempty"`

	Extension DNSInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// DNSInformationExtension type
type DNSInformationExtension struct {
}

// NTPInformation type
type NTPInformation struct {

	// Indicates if NTP information is to be retrieved by using DHCP.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl FromDHCP,omitempty"`

	// List of NTP addresses retrieved by using DHCP.
	NTPFromDHCP []NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl NTPFromDHCP,omitempty"`

	// List of manually entered NTP addresses.
	NTPManual []NetworkHost `xml:"http://www.onvif.org/ver10/device/wsdl NTPManual,omitempty"`

	Extension NTPInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NTPInformationExtension type
type NTPInformationExtension struct {
}

// DynamicDNSInformation type
type DynamicDNSInformation struct {

	// Dynamic DNS type.
	Type DynamicDNSType `xml:"http://www.onvif.org/ver10/device/wsdl Type,omitempty"`

	// DNS name.
	Name DNSName `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`

	// Time to live.
	TTL Duration `xml:"http://www.onvif.org/ver10/schema TTL,omitempty"`

	Extension DynamicDNSInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// DynamicDNSInformationExtension type
type DynamicDNSInformationExtension struct {
}

// NetworkInterfaceSetConfiguration type
type NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not an interface is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// Link configuration.
	Link NetworkInterfaceConnectionSetting `xml:"http://www.onvif.org/ver10/schema Link,omitempty"`

	// Maximum transmission unit.
	MTU int32 `xml:"http://www.onvif.org/ver10/schema MTU,omitempty"`

	// IPv4 network interface configuration.
	IPv4 IPv4NetworkInterfaceSetConfiguration `xml:"http://www.onvif.org/ver10/schema IPv4,omitempty"`

	// IPv6 network interface configuration.
	IPv6 IPv6NetworkInterfaceSetConfiguration `xml:"http://www.onvif.org/ver10/schema IPv6,omitempty"`

	Extension NetworkInterfaceSetConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkInterfaceSetConfigurationExtension type
type NetworkInterfaceSetConfigurationExtension struct {
	Dot3 []Dot3Configuration `xml:"http://www.onvif.org/ver10/schema Dot3,omitempty"`

	Dot11 []Dot11Configuration `xml:"http://www.onvif.org/ver10/schema Dot11,omitempty"`

	Extension NetworkInterfaceSetConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IPv6NetworkInterfaceSetConfiguration type
type IPv6NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not IPv6 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"http://www.onvif.org/ver10/device/wsdl AcceptRouterAdvert,omitempty"`

	// List of manually added IPv6 addresses.
	Manual []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// DHCP configuration.
	DHCP IPv6DHCPConfiguration `xml:"http://www.onvif.org/ver10/schema DHCP,omitempty"`
}

// IPv4NetworkInterfaceSetConfiguration type
type IPv4NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// List of manually added IPv4 addresses.
	Manual []PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// Indicates whether or not DHCP is used.
	DHCP bool `xml:"http://www.onvif.org/ver10/device/wsdl DHCP,omitempty"`
}

// NetworkGateway type
type NetworkGateway struct {

	// IPv4 address string.
	IPv4Address []IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv4Address,omitempty"`

	// IPv6 address string.
	IPv6Address []IPv6Address `xml:"http://www.onvif.org/ver10/device/wsdl IPv6Address,omitempty"`
}

// NetworkZeroConfiguration type
type NetworkZeroConfiguration struct {

	// Unique identifier of network interface.
	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl InterfaceToken,omitempty"`

	// Indicates whether the zero-configuration is enabled or not.
	Enabled bool `xml:"http://www.onvif.org/ver10/device/wsdl Enabled,omitempty"`

	// The zero-configuration IPv4 address(es)
	Addresses []IPv4Address `xml:"http://www.onvif.org/ver10/device/wsdl Addresses,omitempty"`

	Extension *NetworkZeroConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkZeroConfigurationExtension type
type NetworkZeroConfigurationExtension struct {

	// Optional array holding the configuration for the second and possibly further interfaces.
	Additional []NetworkZeroConfiguration `xml:"http://www.onvif.org/ver10/device/wsdl Additional,omitempty"`

	Extension NetworkZeroConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkZeroConfigurationExtension2 type
type NetworkZeroConfigurationExtension2 struct {
}

// IPAddressFilter type
type IPAddressFilter struct {
	Type IPAddressFilterType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	IPv4Address []PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema IPv4Address,omitempty"`

	IPv6Address []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema IPv6Address,omitempty"`

	Extension IPAddressFilterExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IPAddressFilterExtension type
type IPAddressFilterExtension struct {
}

// Dot11Configuration type
type Dot11Configuration struct {
	SSID Dot11SSIDType `xml:"http://www.onvif.org/ver10/schema SSID,omitempty"`

	Mode Dot11StationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Alias Name `xml:"http://www.onvif.org/ver10/schema Alias,omitempty"`

	Priority NetworkInterfaceConfigPriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	Security Dot11SecurityConfiguration `xml:"http://www.onvif.org/ver10/schema Security,omitempty"`
}

// Dot11SecurityConfiguration type
type Dot11SecurityConfiguration struct {
	Mode Dot11SecurityMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Algorithm Dot11Cipher `xml:"http://www.onvif.org/ver10/schema Algorithm,omitempty"`

	PSK Dot11PSKSet `xml:"http://www.onvif.org/ver10/schema PSK,omitempty"`

	Dot1X ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1X,omitempty"`

	Extension Dot11SecurityConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// Dot11SecurityConfigurationExtension type
type Dot11SecurityConfigurationExtension struct {
}

// Dot11PSKSet type
type Dot11PSKSet struct {

	//
	// According to IEEE802.11-2007 H.4.1 the RSNA PSK consists of 256 bits, or 64 octets when represented in hex
	// Either Key or Passphrase shall be given, if both are supplied Key shall be used by the device and Passphrase ignored.
	//
	Key Dot11PSK `xml:"http://www.onvif.org/ver10/schema Key,omitempty"`

	//
	// According to IEEE802.11-2007 H.4.1 a pass-phrase is a sequence of between 8 and 63 ASCII-encoded characters and
	// each character in the pass-phrase must have an encoding in the range of 32 to 126 (decimal),inclusive.
	// If only Passpharse is supplied the Key shall be derived using the algorithm described in IEEE802.11-2007 section H.4
	//
	Passphrase Dot11PSKPassphrase `xml:"http://www.onvif.org/ver10/schema Passphrase,omitempty"`

	Extension Dot11PSKSetExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// Dot11PSKSetExtension type
type Dot11PSKSetExtension struct {
}

// NetworkInterfaceSetConfigurationExtension2 type
type NetworkInterfaceSetConfigurationExtension2 struct {
}

// Dot11Capabilities type
type Dot11Capabilities struct {
	TKIP bool `xml:"http://www.onvif.org/ver10/device/wsdl TKIP,omitempty"`

	ScanAvailableNetworks bool `xml:"http://www.onvif.org/ver10/device/wsdl ScanAvailableNetworks,omitempty"`

	MultipleConfiguration bool `xml:"http://www.onvif.org/ver10/device/wsdl MultipleConfiguration,omitempty"`

	AdHocStationMode bool `xml:"http://www.onvif.org/ver10/device/wsdl AdHocStationMode,omitempty"`

	WEP bool `xml:"http://www.onvif.org/ver10/device/wsdl WEP,omitempty"`
}

// Dot11Status type
type Dot11Status struct {
	SSID Dot11SSIDType `xml:"http://www.onvif.org/ver10/schema SSID,omitempty"`

	BSSID string `xml:"http://www.onvif.org/ver10/device/wsdl BSSID,omitempty"`

	PairCipher Dot11Cipher `xml:"http://www.onvif.org/ver10/schema PairCipher,omitempty"`

	GroupCipher Dot11Cipher `xml:"http://www.onvif.org/ver10/schema GroupCipher,omitempty"`

	SignalStrength Dot11SignalStrength `xml:"http://www.onvif.org/ver10/schema SignalStrength,omitempty"`

	ActiveConfigAlias ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl ActiveConfigAlias,omitempty"`
}

// Dot11AvailableNetworks type
type Dot11AvailableNetworks struct {
	SSID Dot11SSIDType `xml:"http://www.onvif.org/ver10/schema SSID,omitempty"`

	BSSID string `xml:"http://www.onvif.org/ver10/device/wsdl BSSID,omitempty"`

	// See IEEE802.11 7.3.2.25.2 for details.
	AuthAndMangementSuite []Dot11AuthAndMangementSuite `xml:"http://www.onvif.org/ver10/schema AuthAndMangementSuite,omitempty"`

	PairCipher []Dot11Cipher `xml:"http://www.onvif.org/ver10/schema PairCipher,omitempty"`

	GroupCipher []Dot11Cipher `xml:"http://www.onvif.org/ver10/schema GroupCipher,omitempty"`

	SignalStrength Dot11SignalStrength `xml:"http://www.onvif.org/ver10/schema SignalStrength,omitempty"`

	Extension Dot11AvailableNetworksExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// Dot11AvailableNetworksExtension type
type Dot11AvailableNetworksExtension struct {
}

// Capabilities type
type Capabilities struct {

	// Analytics capabilities
	Analytics AnalyticsCapabilities `xml:"http://www.onvif.org/ver10/schema Analytics,omitempty"`

	// Device capabilities
	Device DeviceCapabilities `xml:"http://www.onvif.org/ver10/schema Device,omitempty"`

	// Event capabilities
	Events EventCapabilities `xml:"http://www.onvif.org/ver10/schema Events,omitempty"`

	// Imaging capabilities
	Imaging ImagingCapabilities `xml:"http://www.onvif.org/ver10/schema Imaging,omitempty"`

	// Media capabilities
	Media MediaCapabilities `xml:"http://www.onvif.org/ver10/schema Media,omitempty"`

	// PTZ capabilities
	PTZ PTZCapabilities `xml:"http://www.onvif.org/ver10/schema PTZ,omitempty"`

	Extension CapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// CapabilitiesExtension type
type CapabilitiesExtension struct {
	DeviceIO DeviceIOCapabilities `xml:"http://www.onvif.org/ver10/schema DeviceIO,omitempty"`

	Display DisplayCapabilities `xml:"http://www.onvif.org/ver10/schema Display,omitempty"`

	Recording RecordingCapabilities `xml:"http://www.onvif.org/ver10/schema Recording,omitempty"`

	Search SearchCapabilities `xml:"http://www.onvif.org/ver10/schema Search,omitempty"`

	Replay ReplayCapabilities `xml:"http://www.onvif.org/ver10/schema Replay,omitempty"`

	Receiver ReceiverCapabilities `xml:"http://www.onvif.org/ver10/schema Receiver,omitempty"`

	AnalyticsDevice AnalyticsDeviceCapabilities `xml:"http://www.onvif.org/ver10/schema AnalyticsDevice,omitempty"`

	Extensions CapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extensions,omitempty"`
}

// CapabilitiesExtension2 type
type CapabilitiesExtension2 struct {
}

// AnalyticsCapabilities type
type AnalyticsCapabilities struct {

	// Analytics service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Indicates whether or not rules are supported.
	RuleSupport bool `xml:"http://www.onvif.org/ver10/device/wsdl RuleSupport,omitempty"`

	// Indicates whether or not modules are supported.
	AnalyticsModuleSupport bool `xml:"http://www.onvif.org/ver10/device/wsdl AnalyticsModuleSupport,omitempty"`
}

// DeviceCapabilities type
type DeviceCapabilities struct {

	// Device service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Network capabilities.
	Network NetworkCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl Network,omitempty"`

	// System capabilities.
	System SystemCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl System,omitempty"`

	// I/O capabilities.
	IO IOCapabilities `xml:"http://www.onvif.org/ver10/schema IO,omitempty"`

	// Security capabilities.
	Security SecurityCapabilities `xml:"http://www.onvif.org/ver10/device/wsdl Security,omitempty"`

	Extension DeviceCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// DeviceCapabilitiesExtension type
type DeviceCapabilitiesExtension struct {
}

// EventCapabilities type
type EventCapabilities struct {

	// Event service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Indicates whether or not WS Subscription policy is supported.
	WSSubscriptionPolicySupport bool `xml:"http://www.onvif.org/ver10/device/wsdl WSSubscriptionPolicySupport,omitempty"`

	// Indicates whether or not WS Pull Point is supported.
	WSPullPointSupport bool `xml:"http://www.onvif.org/ver10/device/wsdl WSPullPointSupport,omitempty"`

	// Indicates whether or not WS Pausable Subscription Manager Interface is supported.
	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"http://www.onvif.org/ver10/device/wsdl WSPausableSubscriptionManagerInterfaceSupport,omitempty"`
}

// IOCapabilities type
type IOCapabilities struct {

	// Number of input connectors.
	InputConnectors int32 `xml:"http://www.onvif.org/ver10/schema InputConnectors,omitempty"`

	// Number of relay outputs.
	RelayOutputs int32 `xml:"http://www.onvif.org/ver10/schema RelayOutputs,omitempty"`

	Extension IOCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IOCapabilitiesExtension type
type IOCapabilitiesExtension struct {
	Auxiliary bool `xml:"http://www.onvif.org/ver10/device/wsdl Auxiliary,omitempty"`

	AuxiliaryCommands []AuxiliaryData `xml:"http://www.onvif.org/ver10/device/wsdl AuxiliaryCommands,omitempty"`

	Extension IOCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IOCapabilitiesExtension2 type
type IOCapabilitiesExtension2 struct {
}

// MediaCapabilities type
type MediaCapabilities struct {

	// Media service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Streaming capabilities.
	StreamingCapabilities RealTimeStreamingCapabilities `xml:"http://www.onvif.org/ver10/schema StreamingCapabilities,omitempty"`

	Extension MediaCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// MediaCapabilitiesExtension type
type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities `xml:"http://www.onvif.org/ver10/schema ProfileCapabilities,omitempty"`
}

// RealTimeStreamingCapabilities type
type RealTimeStreamingCapabilities struct {

	// Indicates whether or not RTP multicast is supported.
	RTPMulticast bool `xml:"http://www.onvif.org/ver10/device/wsdl RTPMulticast,omitempty"`

	// Indicates whether or not RTP over TCP is supported.
	RTP_TCP bool `xml:"http://www.onvif.org/ver10/device/wsdl RTP_TCP,omitempty"`

	// Indicates whether or not RTP/RTSP/TCP is supported.
	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/device/wsdl RTP_RTSP_TCP,omitempty"`

	Extension RealTimeStreamingCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// RealTimeStreamingCapabilitiesExtension type
type RealTimeStreamingCapabilitiesExtension struct {
}

// ProfileCapabilities type
type ProfileCapabilities struct {

	// Maximum number of profiles.
	MaximumNumberOfProfiles int32 `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfProfiles,omitempty"`
}

// NetworkCapabilitiesExtension type
type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool `xml:"http://www.onvif.org/ver10/device/wsdl Dot11Configuration,omitempty"`

	Extension NetworkCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkCapabilitiesExtension2 type
type NetworkCapabilitiesExtension2 struct {
}

// SecurityCapabilitiesExtension type
type SecurityCapabilitiesExtension struct {
	TLS10 bool `xml:"TLS1.0,omitempty"`

	Extension SecurityCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SecurityCapabilitiesExtension2 type
type SecurityCapabilitiesExtension2 struct {
	Dot1X bool `xml:"http://www.onvif.org/ver10/device/wsdl Dot1X,omitempty"`

	// EAP Methods supported by the device. The int values refer to the .
	SupportedEAPMethod []int32 `xml:"http://www.onvif.org/ver10/schema SupportedEAPMethod,omitempty"`

	RemoteUserHandling bool `xml:"http://www.onvif.org/ver10/device/wsdl RemoteUserHandling,omitempty"`
}

// SystemCapabilitiesExtension type
type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpFirmwareUpgrade,omitempty"`

	HttpSystemBackup bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpSystemBackup,omitempty"`

	HttpSystemLogging bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpSystemLogging,omitempty"`

	HttpSupportInformation bool `xml:"http://www.onvif.org/ver10/device/wsdl HttpSupportInformation,omitempty"`

	Extension SystemCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SystemCapabilitiesExtension2 type
type SystemCapabilitiesExtension2 struct {
}

// OnvifVersion type
type OnvifVersion struct {

	// Major version number.
	Major int32 `xml:"http://www.onvif.org/ver10/schema Major,omitempty"`

	//
	// Two digit minor version number.
	// If major version number is less than "16", X.0.1 maps to "01" and X.2.1 maps to "21" where X stands for Major version number.
	// Otherwise, minor number is month of release, such as "06" for June.
	//
	Minor int32 `xml:"http://www.onvif.org/ver10/schema Minor,omitempty"`
}

// ImagingCapabilities type
type ImagingCapabilities struct {

	// Imaging service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`
}

// PTZCapabilities type
type PTZCapabilities struct {

	// PTZ service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`
}

// DeviceIOCapabilities type
type DeviceIOCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	VideoSources int32 `xml:"http://www.onvif.org/ver10/schema VideoSources,omitempty"`

	VideoOutputs int32 `xml:"http://www.onvif.org/ver10/schema VideoOutputs,omitempty"`

	AudioSources int32 `xml:"http://www.onvif.org/ver10/schema AudioSources,omitempty"`

	AudioOutputs int32 `xml:"http://www.onvif.org/ver10/schema AudioOutputs,omitempty"`

	RelayOutputs int32 `xml:"http://www.onvif.org/ver10/schema RelayOutputs,omitempty"`
}

// DisplayCapabilities type
type DisplayCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Indication that the SetLayout command supports only predefined layouts.
	FixedLayout bool `xml:"http://www.onvif.org/ver10/device/wsdl FixedLayout,omitempty"`
}

// RecordingCapabilities type
type RecordingCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	ReceiverSource bool `xml:"http://www.onvif.org/ver10/device/wsdl ReceiverSource,omitempty"`

	MediaProfileSource bool `xml:"http://www.onvif.org/ver10/device/wsdl MediaProfileSource,omitempty"`

	DynamicRecordings bool `xml:"http://www.onvif.org/ver10/device/wsdl DynamicRecordings,omitempty"`

	DynamicTracks bool `xml:"http://www.onvif.org/ver10/device/wsdl DynamicTracks,omitempty"`

	MaxStringLength int32 `xml:"http://www.onvif.org/ver10/schema MaxStringLength,omitempty"`
}

// SearchCapabilities type
type SearchCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	MetadataSearch bool `xml:"http://www.onvif.org/ver10/device/wsdl MetadataSearch,omitempty"`
}

// ReplayCapabilities type
type ReplayCapabilities struct {

	// The address of the replay service.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`
}

// ReceiverCapabilities type
type ReceiverCapabilities struct {

	// The address of the receiver service.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Indicates whether the device can receive RTP multicast streams.
	RTP_Multicast bool `xml:"http://www.onvif.org/ver10/device/wsdl RTP_Multicast,omitempty"`

	// Indicates whether the device can receive RTP/TCP streams
	RTP_TCP bool `xml:"http://www.onvif.org/ver10/device/wsdl RTP_TCP,omitempty"`

	// Indicates whether the device can receive RTP/RTSP/TCP streams.
	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/device/wsdl RTP_RTSP_TCP,omitempty"`

	// The maximum number of receivers supported by the device.
	SupportedReceivers int32 `xml:"http://www.onvif.org/ver10/schema SupportedReceivers,omitempty"`

	// The maximum allowed length for RTSP URIs.
	MaximumRTSPURILength int32 `xml:"http://www.onvif.org/ver10/schema MaximumRTSPURILength,omitempty"`
}

// AnalyticsDeviceCapabilities type
type AnalyticsDeviceCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Obsolete property.
	RuleSupport bool `xml:"http://www.onvif.org/ver10/device/wsdl RuleSupport,omitempty"`

	Extension AnalyticsDeviceExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AnalyticsDeviceExtension type
type AnalyticsDeviceExtension struct {
}

// SystemLog type
type SystemLog struct {

	// The log information as attachment data.
	Binary AttachmentData `xml:"http://www.onvif.org/ver10/device/wsdl Binary,omitempty"`

	// The log information as character data.
	String string `xml:"http://www.onvif.org/ver10/device/wsdl String,omitempty"`
}

// SupportInformation type
type SupportInformation struct {

	// The support information as attachment data.
	Binary AttachmentData `xml:"http://www.onvif.org/ver10/device/wsdl Binary,omitempty"`

	// The support information as character data.
	String string `xml:"http://www.onvif.org/ver10/device/wsdl String,omitempty"`
}

// BinaryData type
type BinaryData struct {

	// base64 encoded binary data.
	Data []byte `xml:"http://www.onvif.org/ver10/schema Data,omitempty"`

	ContentType string `xml:"contentType,attr,omitempty"`
}

// AttachmentData type
type AttachmentData struct {
	Include Include `xml:"Include,omitempty"`

	ContentType string `xml:"contentType,attr,omitempty"`
}

// BackupFile type
type BackupFile struct {
	Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`

	Data AttachmentData `xml:"http://www.onvif.org/ver10/device/wsdl Data,omitempty"`
}

// SystemLogUriList type
type SystemLogUriList struct {
	SystemLog []SystemLogUri `xml:"http://www.onvif.org/ver10/device/wsdl SystemLog,omitempty"`
}

// SystemLogUri type
type SystemLogUri struct {
	Type SystemLogType `xml:"http://www.onvif.org/ver10/device/wsdl Type,omitempty"`

	Uri AnyURI `xml:"http://www.onvif.org/ver10/schema Uri,omitempty"`
}

// SystemDateTime type
type SystemDateTime struct {

	// Indicates if the time is set manully or through NTP.
	DateTimeType SetDateTimeType `xml:"http://www.onvif.org/ver10/device/wsdl DateTimeType,omitempty"`

	// Informative indicator whether daylight savings is currently on/off.
	DaylightSavings bool `xml:"http://www.onvif.org/ver10/device/wsdl DaylightSavings,omitempty"`

	// Timezone information in Posix format.
	TimeZone TimeZone `xml:"http://www.onvif.org/ver10/device/wsdl TimeZone,omitempty"`

	// Current system date and time in UTC format. This field is mandatory since version 2.0.
	UTCDateTime string `xml:"http://www.onvif.org/ver10/schema UTCDateTime,omitempty"`

	// Date and time in local format.
	LocalDateTime string `xml:"http://www.onvif.org/ver10/schema LocalDateTime,omitempty"`

	Extension SystemDateTimeExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SystemDateTimeExtension type
type SystemDateTimeExtension struct {
}

// DateTime type
type DateTime struct {
	Time string `xml:"http://www.onvif.org/ver10/schema Time,omitempty"`

	Date string `xml:"http://www.onvif.org/ver10/schema Date,omitempty"`
}

// Date type
type Date struct {
	Year int32 `xml:"http://www.onvif.org/ver10/schema Year,omitempty"`

	// Range is 1 to 12.
	Month int32 `xml:"http://www.onvif.org/ver10/schema Month,omitempty"`

	// Range is 1 to 31.
	Day int32 `xml:"http://www.onvif.org/ver10/schema Day,omitempty"`
}

// Time type
type Time struct {

	// Range is 0 to 23.
	Hour int32 `xml:"http://www.onvif.org/ver10/schema Hour,omitempty"`

	// Range is 0 to 59.
	Minute int32 `xml:"http://www.onvif.org/ver10/schema Minute,omitempty"`

	// Range is 0 to 61 (typically 59).
	Second int32 `xml:"http://www.onvif.org/ver10/schema Second,omitempty"`
}

// TimeZone type
type TimeZone struct {

	// Posix timezone string.
	TZ string `xml:"http://www.onvif.org/ver10/device/wsdl TZ,omitempty"`
}

// RemoteUser type
type RemoteUser struct {
	Username string `xml:"http://www.onvif.org/ver10/device/wsdl Username,omitempty"`

	Password string `xml:"http://www.onvif.org/ver10/device/wsdl Password,omitempty"`

	UseDerivedPassword bool `xml:"http://www.onvif.org/ver10/device/wsdl UseDerivedPassword,omitempty"`
}

// User type
type User struct {

	// Username string.
	Username string `xml:"http://www.onvif.org/ver10/device/wsdl Username,omitempty"`

	// Password string.
	Password string `xml:"http://www.onvif.org/ver10/device/wsdl Password,omitempty"`

	// User level string.
	UserLevel UserLevel `xml:"http://www.onvif.org/ver10/schema UserLevel,omitempty"`

	Extension UserExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// UserExtension type
type UserExtension struct {
}

// CertificateGenerationParameters type
type CertificateGenerationParameters struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	Subject string `xml:"http://www.onvif.org/ver10/device/wsdl Subject,omitempty"`

	ValidNotBefore string `xml:"http://www.onvif.org/ver10/device/wsdl ValidNotBefore,omitempty"`

	ValidNotAfter string `xml:"http://www.onvif.org/ver10/device/wsdl ValidNotAfter,omitempty"`

	Extension CertificateGenerationParametersExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// CertificateGenerationParametersExtension type
type CertificateGenerationParametersExtension struct {
}

// Certificate type
type Certificate struct {

	// Certificate id.
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	// base64 encoded DER representation of certificate.
	Certificate BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl Certificate,omitempty"`
}

// CertificateStatus type
type CertificateStatus struct {

	// Certificate id.
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	// Indicates whether or not a certificate is used in a HTTPS configuration.
	Status bool `xml:"http://www.onvif.org/ver10/device/wsdl Status,omitempty"`
}

// CertificateWithPrivateKey type
type CertificateWithPrivateKey struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	Certificate BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl Certificate,omitempty"`

	PrivateKey BinaryData `xml:"http://www.onvif.org/ver10/device/wsdl PrivateKey,omitempty"`
}

// CertificateInformation type
type CertificateInformation struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`

	IssuerDN string `xml:"http://www.onvif.org/ver10/device/wsdl IssuerDN,omitempty"`

	SubjectDN string `xml:"http://www.onvif.org/ver10/device/wsdl SubjectDN,omitempty"`

	KeyUsage CertificateUsage `xml:"http://www.onvif.org/ver10/schema KeyUsage,omitempty"`

	ExtendedKeyUsage CertificateUsage `xml:"http://www.onvif.org/ver10/schema ExtendedKeyUsage,omitempty"`

	KeyLength int32 `xml:"http://www.onvif.org/ver10/schema KeyLength,omitempty"`

	Version string `xml:"http://www.onvif.org/ver10/device/wsdl Version,omitempty"`

	SerialNum string `xml:"http://www.onvif.org/ver10/device/wsdl SerialNum,omitempty"`

	// Validity Range is from "NotBefore" to "NotAfter"; the corresponding DateTimeRange is from "From" to "Until"
	SignatureAlgorithm string `xml:"http://www.onvif.org/ver10/device/wsdl SignatureAlgorithm,omitempty"`

	Validity DateTimeRange `xml:"http://www.onvif.org/ver10/schema Validity,omitempty"`

	Extension CertificateInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// CertificateUsage type
type CertificateUsage struct {
	Value string

	Critical bool `xml:"http://www.onvif.org/ver10/device/wsdl Critical,attr,omitempty"`
}

// CertificateInformationExtension type
type CertificateInformationExtension struct {
}

// Dot1XConfiguration type
type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Dot1XConfigurationToken,omitempty"`

	Identity string `xml:"http://www.onvif.org/ver10/device/wsdl Identity,omitempty"`

	AnonymousID string `xml:"http://www.onvif.org/ver10/device/wsdl AnonymousID,omitempty"`

	//
	// EAP Method type as defined in .
	//
	EAPMethod int32 `xml:"http://www.onvif.org/ver10/schema EAPMethod,omitempty"`

	CACertificateID []string `xml:"http://www.onvif.org/ver10/device/wsdl CACertificateID,omitempty"`

	EAPMethodConfiguration EAPMethodConfiguration `xml:"http://www.onvif.org/ver10/schema EAPMethodConfiguration,omitempty"`

	Extension Dot1XConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// Dot1XConfigurationExtension type
type Dot1XConfigurationExtension struct {
}

// EAPMethodConfiguration type
type EAPMethodConfiguration struct {

	// Confgiuration information for TLS Method.
	TLSConfiguration TLSConfiguration `xml:"http://www.onvif.org/ver10/schema TLSConfiguration,omitempty"`

	// Password for those EAP Methods that require a password. The password shall never be returned on a get method.
	Password string `xml:"http://www.onvif.org/ver10/device/wsdl Password,omitempty"`

	Extension EapMethodExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// EapMethodExtension type
type EapMethodExtension struct {
}

// TLSConfiguration type
type TLSConfiguration struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/device/wsdl CertificateID,omitempty"`
}

// Removed GenericEapPwdConfigurationExtension by fixgen.py

// RelayOutputSettings type
type RelayOutputSettings struct {

	//
	// 'Bistable' or 'Monostable'
	//
	//
	Mode RelayMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Time after which the relay returns to its idle state if it is in monostable mode. If the Mode field is set to bistable mode the value of the parameter can be ignored.
	DelayTime Duration `xml:"http://www.onvif.org/ver10/schema DelayTime,omitempty"`

	//
	// 'open' or 'closed'
	//
	//
	IdleState RelayIdleState `xml:"http://www.onvif.org/ver10/schema IdleState,omitempty"`
}

// RelayOutput type
type RelayOutput struct {
	*DeviceEntity

	Properties RelayOutputSettings `xml:"http://www.onvif.org/ver10/device/wsdl Properties,omitempty"`
}

// Removed DigitalInput by fixgen.py

// PTZNode type
type PTZNode struct {
	*DeviceEntity

	//
	// A unique identifier that is used to reference PTZ Nodes.
	//
	Name Name `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	//
	// A list of Coordinate Systems available for the PTZ Node. For each Coordinate System, the PTZ Node MUST specify its allowed range.
	//
	SupportedPTZSpaces PTZSpaces `xml:"http://www.onvif.org/ver10/schema SupportedPTZSpaces,omitempty"`

	//
	// All preset operations MUST be available for this PTZ Node if one preset is supported.
	//
	MaximumNumberOfPresets int32 `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfPresets,omitempty"`

	//
	// A boolean operator specifying the availability of a home position. If set to true, the Home Position Operations MUST be available for this PTZ Node.
	//
	HomeSupported bool `xml:"http://www.onvif.org/ver10/device/wsdl HomeSupported,omitempty"`

	//
	// A list of supported Auxiliary commands. If the list is not empty, the Auxiliary Operations MUST be available for this PTZ Node.
	//
	AuxiliaryCommands []AuxiliaryData `xml:"http://www.onvif.org/ver10/device/wsdl AuxiliaryCommands,omitempty"`

	Extension PTZNodeExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	//
	// Indication whether the HomePosition of a Node is fixed or it can be changed via the SetHomePosition command.
	//

	FixedHomePosition bool `xml:"http://www.onvif.org/ver10/device/wsdl FixedHomePosition,attr,omitempty"`

	//
	// Indication whether the Node supports the geo-referenced move command.
	//

	GeoMove bool `xml:"http://www.onvif.org/ver10/device/wsdl GeoMove,attr,omitempty"`
}

// PTZNodeExtension type
type PTZNodeExtension struct {

	//
	// Detail of supported Preset Tour feature.
	//
	SupportedPresetTour PTZPresetTourSupported `xml:"http://www.onvif.org/ver10/schema SupportedPresetTour,omitempty"`

	Extension PTZNodeExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZNodeExtension2 type
type PTZNodeExtension2 struct {
}

// PTZPresetTourSupported type
type PTZPresetTourSupported struct {

	// Indicates number of preset tours that can be created. Required preset tour operations shall be available for this PTZ Node if one or more preset tour is supported.
	MaximumNumberOfPresetTours int32 `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfPresetTours,omitempty"`

	// Indicates which preset tour operations are available for this PTZ Node.
	PTZPresetTourOperation []PTZPresetTourOperation `xml:"http://www.onvif.org/ver10/schema PTZPresetTourOperation,omitempty"`

	Extension PTZPresetTourSupportedExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZPresetTourSupportedExtension type
type PTZPresetTourSupportedExtension struct {
}

// PTZConfiguration type
type PTZConfiguration struct {
	*ConfigurationEntity

	//
	// A mandatory reference to the PTZ Node that the PTZ Configuration belongs to.
	//
	NodeToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl NodeToken,omitempty"`

	//
	// If the PTZ Node supports absolute Pan/Tilt movements, it shall specify one Absolute Pan/Tilt Position Space as default.
	//
	DefaultAbsolutePantTiltPositionSpace AnyURI `xml:"http://www.onvif.org/ver10/schema DefaultAbsolutePantTiltPositionSpace,omitempty"`

	//
	// If the PTZ Node supports absolute zoom movements, it shall specify one Absolute Zoom Position Space as default.
	//
	DefaultAbsoluteZoomPositionSpace AnyURI `xml:"http://www.onvif.org/ver10/schema DefaultAbsoluteZoomPositionSpace,omitempty"`

	//
	// If the PTZ Node supports relative Pan/Tilt movements, it shall specify one RelativePan/Tilt Translation Space as default.
	//
	DefaultRelativePanTiltTranslationSpace AnyURI `xml:"http://www.onvif.org/ver10/schema DefaultRelativePanTiltTranslationSpace,omitempty"`

	//
	// If the PTZ Node supports relative zoom movements, it shall specify one Relative Zoom Translation Space as default.
	//
	DefaultRelativeZoomTranslationSpace AnyURI `xml:"http://www.onvif.org/ver10/schema DefaultRelativeZoomTranslationSpace,omitempty"`

	//
	// If the PTZ Node supports continuous Pan/Tilt movements, it shall specify one Continuous Pan/Tilt Velocity Space as default.
	//
	DefaultContinuousPanTiltVelocitySpace AnyURI `xml:"http://www.onvif.org/ver10/schema DefaultContinuousPanTiltVelocitySpace,omitempty"`

	//
	// If the PTZ Node supports continuous zoom movements, it shall specify one Continuous Zoom Velocity Space as default.
	//
	DefaultContinuousZoomVelocitySpace AnyURI `xml:"http://www.onvif.org/ver10/schema DefaultContinuousZoomVelocitySpace,omitempty"`

	//
	// If the PTZ Node supports absolute or relative PTZ movements, it shall specify corresponding default Pan/Tilt and Zoom speeds.
	//
	DefaultPTZSpeed PTZSpeed `xml:"http://www.onvif.org/ver10/schema DefaultPTZSpeed,omitempty"`

	//
	// If the PTZ Node supports continuous movements, it shall specify a default timeout, after which the movement stops.
	//
	DefaultPTZTimeout Duration `xml:"http://www.onvif.org/ver10/schema DefaultPTZTimeout,omitempty"`

	//
	// The Pan/Tilt limits element should be present for a PTZ Node that supports an absolute Pan/Tilt. If the element is present it signals the support for configurable Pan/Tilt limits. If limits are enabled, the Pan/Tilt movements shall always stay within the specified range. The Pan/Tilt limits are disabled by setting the limits to –INF or +INF.
	//
	PanTiltLimits PanTiltLimits `xml:"http://www.onvif.org/ver10/schema PanTiltLimits,omitempty"`

	//
	// The Zoom limits element should be present for a PTZ Node that supports absolute zoom. If the element is present it signals the supports for configurable Zoom limits. If limits are enabled the zoom movements shall always stay within the specified range. The Zoom limits are disabled by settings the limits to -INF and +INF.
	//
	ZoomLimits ZoomLimits `xml:"http://www.onvif.org/ver10/schema ZoomLimits,omitempty"`

	Extension PTZConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// The optional acceleration ramp used by the device when moving.

	MoveRamp int32 `xml:"http://www.onvif.org/ver10/schema MoveRamp,attr,omitempty"`

	// The optional acceleration ramp used by the device when recalling presets.

	PresetRamp int32 `xml:"http://www.onvif.org/ver10/schema PresetRamp,attr,omitempty"`

	// The optional acceleration ramp used by the device when executing PresetTours.

	PresetTourRamp int32 `xml:"http://www.onvif.org/ver10/schema PresetTourRamp,attr,omitempty"`
}

// PTZConfigurationExtension type
type PTZConfigurationExtension struct {

	// Optional element to configure PT Control Direction related features.
	PTControlDirection PTControlDirection `xml:"http://www.onvif.org/ver10/schema PTControlDirection,omitempty"`

	Extension PTZConfigurationExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZConfigurationExtension2 type
type PTZConfigurationExtension2 struct {
}

// PTControlDirection type
type PTControlDirection struct {

	// Optional element to configure related parameters for E-Flip.
	EFlip EFlip `xml:"http://www.onvif.org/ver10/schema EFlip,omitempty"`

	// Optional element to configure related parameters for reversing of PT Control Direction.
	Reverse Reverse `xml:"http://www.onvif.org/ver10/schema Reverse,omitempty"`

	Extension PTControlDirectionExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTControlDirectionExtension type
type PTControlDirectionExtension struct {
}

// EFlip type
type EFlip struct {

	// Parameter to enable/disable E-Flip feature.
	Mode EFlipMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`
}

// Reverse type
type Reverse struct {

	// Parameter to enable/disable Reverse feature.
	Mode ReverseMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`
}

// PTZConfigurationOptions type
type PTZConfigurationOptions struct {

	//
	// A list of supported coordinate systems including their range limitations.
	//
	Spaces PTZSpaces `xml:"http://www.onvif.org/ver10/schema Spaces,omitempty"`

	//
	// A timeout Range within which Timeouts are accepted by the PTZ Node.
	//
	PTZTimeout DurationRange `xml:"http://www.onvif.org/ver10/schema PTZTimeout,omitempty"`

	// Supported options for PT Direction Control.
	PTControlDirection PTControlDirectionOptions `xml:"http://www.onvif.org/ver10/schema PTControlDirection,omitempty"`

	Extension PTZConfigurationOptions2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	//
	// The list of acceleration ramps supported by the device. The
	// smallest acceleration value corresponds to the minimal index, the
	// highest acceleration corresponds to the maximum index.
	//

	PTZRamps IntAttrList `xml:"http://www.onvif.org/ver10/schema PTZRamps,attr,omitempty"`
}

// PTZConfigurationOptions2 type
type PTZConfigurationOptions2 struct {
}

// PTControlDirectionOptions type
type PTControlDirectionOptions struct {

	// Supported options for EFlip feature.
	EFlip EFlipOptions `xml:"http://www.onvif.org/ver10/schema EFlip,omitempty"`

	// Supported options for Reverse feature.
	Reverse ReverseOptions `xml:"http://www.onvif.org/ver10/schema Reverse,omitempty"`

	Extension PTControlDirectionOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTControlDirectionOptionsExtension type
type PTControlDirectionOptionsExtension struct {
}

// EFlipOptions type
type EFlipOptions struct {

	// Options of EFlip mode parameter.
	Mode []EFlipMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Extension EFlipOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// EFlipOptionsExtension type
type EFlipOptionsExtension struct {
}

// ReverseOptions type
type ReverseOptions struct {

	// Options of Reverse mode parameter.
	Mode []ReverseMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Extension ReverseOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ReverseOptionsExtension type
type ReverseOptionsExtension struct {
}

// PanTiltLimits type
type PanTiltLimits struct {

	//
	// A range of pan tilt limits.
	//
	Range Space2DDescription `xml:"http://www.onvif.org/ver10/schema Range,omitempty"`
}

// ZoomLimits type
type ZoomLimits struct {

	//
	// A range of zoom limit
	//
	Range Space1DDescription `xml:"http://www.onvif.org/ver10/schema Range,omitempty"`
}

// PTZSpaces type
type PTZSpaces struct {

	//
	// The Generic Pan/Tilt Position space is provided by every PTZ node that supports absolute Pan/Tilt, since it does not relate to a specific physical range.
	// Instead, the range should be defined as the full range of the PTZ unit normalized to the range -1 to 1 resulting in the following space description.
	//
	AbsolutePanTiltPositionSpace []Space2DDescription `xml:"http://www.onvif.org/ver10/schema AbsolutePanTiltPositionSpace,omitempty"`

	//
	// The Generic Zoom Position Space is provided by every PTZ node that supports absolute Zoom, since it does not relate to a specific physical range.
	// Instead, the range should be defined as the full range of the Zoom normalized to the range 0 (wide) to 1 (tele).
	// There is no assumption about how the generic zoom range is mapped to magnification, FOV or other physical zoom dimension.
	//
	AbsoluteZoomPositionSpace []Space1DDescription `xml:"http://www.onvif.org/ver10/schema AbsoluteZoomPositionSpace,omitempty"`

	//
	// The Generic Pan/Tilt translation space is provided by every PTZ node that supports relative Pan/Tilt, since it does not relate to a specific physical range.
	// Instead, the range should be defined as the full positive and negative translation range of the PTZ unit normalized to the range -1 to 1,
	// where positive translation would mean clockwise rotation or movement in right/up direction resulting in the following space description.
	//
	RelativePanTiltTranslationSpace []Space2DDescription `xml:"http://www.onvif.org/ver10/schema RelativePanTiltTranslationSpace,omitempty"`

	//
	// The Generic Zoom Translation Space is provided by every PTZ node that supports relative Zoom, since it does not relate to a specific physical range.
	// Instead, the corresponding absolute range should be defined as the full positive and negative translation range of the Zoom normalized to the range -1 to1,
	// where a positive translation maps to a movement in TELE direction. The translation is signed to indicate direction (negative is to wide, positive is to tele).
	// There is no assumption about how the generic zoom range is mapped to magnification, FOV or other physical zoom dimension. This results in the following space description.
	//
	RelativeZoomTranslationSpace []Space1DDescription `xml:"http://www.onvif.org/ver10/schema RelativeZoomTranslationSpace,omitempty"`

	//
	// The generic Pan/Tilt velocity space shall be provided by every PTZ node, since it does not relate to a specific physical range.
	// Instead, the range should be defined as a range of the PTZ unit’s speed normalized to the range -1 to 1, where a positive velocity would map to clockwise
	// rotation or movement in the right/up direction. A signed speed can be independently specified for the pan and tilt component resulting in the following space description.
	//
	ContinuousPanTiltVelocitySpace []Space2DDescription `xml:"http://www.onvif.org/ver10/schema ContinuousPanTiltVelocitySpace,omitempty"`

	//
	// The generic zoom velocity space specifies a zoom factor velocity without knowing the underlying physical model. The range should be normalized from -1 to 1,
	// where a positive velocity would map to TELE direction. A generic zoom velocity space description resembles the following.
	//
	ContinuousZoomVelocitySpace []Space1DDescription `xml:"http://www.onvif.org/ver10/schema ContinuousZoomVelocitySpace,omitempty"`

	//
	// The speed space specifies the speed for a Pan/Tilt movement when moving to an absolute position or to a relative translation.
	// In contrast to the velocity spaces, speed spaces do not contain any directional information. The speed of a combined Pan/Tilt
	// movement is represented by a single non-negative scalar value.
	//
	PanTiltSpeedSpace []Space1DDescription `xml:"http://www.onvif.org/ver10/schema PanTiltSpeedSpace,omitempty"`

	//
	// The speed space specifies the speed for a Zoom movement when moving to an absolute position or to a relative translation.
	// In contrast to the velocity spaces, speed spaces do not contain any directional information.
	//
	ZoomSpeedSpace []Space1DDescription `xml:"http://www.onvif.org/ver10/schema ZoomSpeedSpace,omitempty"`

	Extension PTZSpacesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZSpacesExtension type
type PTZSpacesExtension struct {
}

// Space2DDescription type
type Space2DDescription struct {

	//
	// A URI of coordinate systems.
	//
	URI AnyURI `xml:"http://www.onvif.org/ver10/schema URI,omitempty"`

	//
	// A range of x-axis.
	//
	XRange FloatRange `xml:"http://www.onvif.org/ver10/schema XRange,omitempty"`

	//
	// A range of y-axis.
	//
	YRange FloatRange `xml:"http://www.onvif.org/ver10/schema YRange,omitempty"`
}

// Space1DDescription type
type Space1DDescription struct {

	//
	// A URI of coordinate systems.
	//
	URI AnyURI `xml:"http://www.onvif.org/ver10/schema URI,omitempty"`

	//
	// A range of x-axis.
	//
	XRange FloatRange `xml:"http://www.onvif.org/ver10/schema XRange,omitempty"`
}

// PTZSpeed type
type PTZSpeed struct {

	// Pan and tilt speed. The x component corresponds to pan and the y component to tilt. If omitted in a request, the current (if any) PanTilt movement should not be affected.
	PanTilt Vector2D `xml:"http://www.onvif.org/ver10/schema PanTilt,omitempty"`

	//
	// A zoom speed. If omitted in a request, the current (if any) Zoom movement should not be affected.
	//
	Zoom Vector1D `xml:"http://www.onvif.org/ver10/schema Zoom,omitempty"`
}

// PTZPreset type
type PTZPreset struct {

	//
	// A list of preset position name.
	//
	Name Name `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	//
	// A list of preset position.
	//
	PTZPosition PTZVector `xml:"http://www.onvif.org/ver10/schema PTZPosition,omitempty"`

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

// Removed PresetTour by fixgen.py

// Removed PTZPresetTourExtension by fixgen.py

// PTZPresetTourSpot type
type PTZPresetTourSpot struct {

	// Detail definition of preset position of the tour spot.
	PresetDetail PTZPresetTourPresetDetail `xml:"http://www.onvif.org/ver10/schema PresetDetail,omitempty"`

	// Optional parameter to specify Pan/Tilt and Zoom speed on moving toward this tour spot.
	Speed PTZSpeed `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`

	// Optional parameter to specify time duration of staying on this tour sport.
	StayTime Duration `xml:"http://www.onvif.org/ver10/schema StayTime,omitempty"`

	Extension PTZPresetTourSpotExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZPresetTourSpotExtension type
type PTZPresetTourSpotExtension struct {
}

// PTZPresetTourPresetDetail type
type PTZPresetTourPresetDetail struct {

	// Option to specify the preset position with Preset Token defined in advance.
	PresetToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl PresetToken,omitempty"`

	// Option to specify the preset position with the home position of this PTZ Node. "False" to this parameter shall be treated as an invalid argument.
	Home bool `xml:"http://www.onvif.org/ver10/device/wsdl Home,omitempty"`

	// Option to specify the preset position with vector of PTZ node directly.
	PTZPosition PTZVector `xml:"http://www.onvif.org/ver10/schema PTZPosition,omitempty"`

	TypeExtension PTZPresetTourTypeExtension `xml:"http://www.onvif.org/ver10/schema TypeExtension,omitempty"`
}

// PTZPresetTourTypeExtension type
type PTZPresetTourTypeExtension struct {
}

// PTZPresetTourStatus type
type PTZPresetTourStatus struct {

	// Indicates state of this preset tour by Idle/Touring/Paused.
	State PTZPresetTourState `xml:"http://www.onvif.org/ver10/schema State,omitempty"`

	// Indicates a tour spot currently staying.
	CurrentTourSpot PTZPresetTourSpot `xml:"http://www.onvif.org/ver10/schema CurrentTourSpot,omitempty"`

	Extension PTZPresetTourStatusExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZPresetTourStatusExtension type
type PTZPresetTourStatusExtension struct {
}

// PTZPresetTourStartingCondition type
type PTZPresetTourStartingCondition struct {

	// Optional parameter to specify how many times the preset tour is recurred.
	RecurringTime int32 `xml:"http://www.onvif.org/ver10/schema RecurringTime,omitempty"`

	// Optional parameter to specify how long time duration the preset tour is recurred.
	RecurringDuration Duration `xml:"http://www.onvif.org/ver10/schema RecurringDuration,omitempty"`

	// Optional parameter to choose which direction the preset tour goes. Forward shall be chosen in case it is omitted.
	Direction PTZPresetTourDirection `xml:"http://www.onvif.org/ver10/schema Direction,omitempty"`

	Extension PTZPresetTourStartingConditionExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Execute presets in random order. If set to true and Direction is also present, Direction will be ignored and presets of the Tour will be recalled randomly.

	RandomPresetOrder bool `xml:"http://www.onvif.org/ver10/device/wsdl RandomPresetOrder,attr,omitempty"`
}

// PTZPresetTourStartingConditionExtension type
type PTZPresetTourStartingConditionExtension struct {
}

// Removed PTZPresetTourOptions by fixgen.py

// Removed PTZPresetTourSpotOptions by fixgen.py

// PTZPresetTourPresetDetailOptions type
type PTZPresetTourPresetDetailOptions struct {

	// A list of available Preset Tokens for tour spots.
	PresetToken []ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl PresetToken,omitempty"`

	// An option to indicate Home postion for tour spots.
	Home bool `xml:"http://www.onvif.org/ver10/device/wsdl Home,omitempty"`

	// Supported range of Pan and Tilt for tour spots.
	PanTiltPositionSpace Space2DDescription `xml:"http://www.onvif.org/ver10/schema PanTiltPositionSpace,omitempty"`

	// Supported range of Zoom for a tour spot.
	ZoomPositionSpace Space1DDescription `xml:"http://www.onvif.org/ver10/schema ZoomPositionSpace,omitempty"`

	Extension PTZPresetTourPresetDetailOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZPresetTourPresetDetailOptionsExtension type
type PTZPresetTourPresetDetailOptionsExtension struct {
}

// PTZPresetTourStartingConditionOptions type
type PTZPresetTourStartingConditionOptions struct {

	// Supported range of Recurring Time.
	RecurringTime IntRange `xml:"http://www.onvif.org/ver10/schema RecurringTime,omitempty"`

	// Supported range of Recurring Duration.
	RecurringDuration DurationRange `xml:"http://www.onvif.org/ver10/schema RecurringDuration,omitempty"`

	// Supported options for Direction of Preset Tour.
	Direction []PTZPresetTourDirection `xml:"http://www.onvif.org/ver10/schema Direction,omitempty"`

	Extension PTZPresetTourStartingConditionOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PTZPresetTourStartingConditionOptionsExtension type
type PTZPresetTourStartingConditionOptionsExtension struct {
}

// ImagingStatus type
type ImagingStatus struct {
	FocusStatus FocusStatus `xml:"http://www.onvif.org/ver10/schema FocusStatus,omitempty"`
}

// FocusStatus type
type FocusStatus struct {

	//
	// Status of focus position.
	//
	Position float32 `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`

	//
	// Status of focus MoveStatus.
	//
	MoveStatus MoveStatus `xml:"http://www.onvif.org/ver10/schema MoveStatus,omitempty"`

	//
	// Error status of focus.
	//
	Error string `xml:"http://www.onvif.org/ver10/device/wsdl Error,omitempty"`
}

// FocusConfiguration type
type FocusConfiguration struct {
	AutoFocusMode AutoFocusMode `xml:"http://www.onvif.org/ver10/schema AutoFocusMode,omitempty"`

	DefaultSpeed float32 `xml:"http://www.onvif.org/ver10/schema DefaultSpeed,omitempty"`

	// Parameter to set autofocus near limit (unit: meter).
	NearLimit float32 `xml:"http://www.onvif.org/ver10/schema NearLimit,omitempty"`

	// Parameter to set autofocus far limit (unit: meter).
	// If set to 0.0, infinity will be used.
	FarLimit float32 `xml:"http://www.onvif.org/ver10/schema FarLimit,omitempty"`
}

// ImagingSettings type
type ImagingSettings struct {

	// Enabled/disabled BLC mode (on/off).
	BacklightCompensation BacklightCompensation `xml:"http://www.onvif.org/ver10/schema BacklightCompensation,omitempty"`

	// Image brightness (unit unspecified).
	Brightness float32 `xml:"http://www.onvif.org/ver10/schema Brightness,omitempty"`

	// Color saturation of the image (unit unspecified).
	ColorSaturation float32 `xml:"http://www.onvif.org/ver10/schema ColorSaturation,omitempty"`

	// Contrast of the image (unit unspecified).
	Contrast float32 `xml:"http://www.onvif.org/ver10/schema Contrast,omitempty"`

	// Exposure mode of the device.
	Exposure Exposure `xml:"http://www.onvif.org/ver10/schema Exposure,omitempty"`

	// Focus configuration.
	Focus FocusConfiguration `xml:"http://www.onvif.org/ver10/schema Focus,omitempty"`

	// Infrared Cutoff Filter settings.
	IrCutFilter IrCutFilterMode `xml:"http://www.onvif.org/ver10/schema IrCutFilter,omitempty"`

	// Sharpness of the Video image.
	Sharpness float32 `xml:"http://www.onvif.org/ver10/schema Sharpness,omitempty"`

	// WDR settings.
	WideDynamicRange WideDynamicRange `xml:"http://www.onvif.org/ver10/schema WideDynamicRange,omitempty"`

	// White balance settings.
	WhiteBalance WhiteBalance `xml:"http://www.onvif.org/ver10/schema WhiteBalance,omitempty"`

	Extension ImagingSettingsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingSettingsExtension type
type ImagingSettingsExtension struct {
}

// Exposure type
type Exposure struct {

	//
	// Exposure Mode
	//
	//
	Mode ExposureMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// The exposure priority mode (low noise/framerate).
	//
	Priority ExposurePriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	//
	// Rectangular exposure mask.
	//
	Window Rectangle `xml:"http://www.onvif.org/ver10/schema Window,omitempty"`

	//
	// Minimum value of exposure time range allowed to be used by the algorithm.
	//
	MinExposureTime float32 `xml:"http://www.onvif.org/ver10/schema MinExposureTime,omitempty"`

	//
	// Maximum value of exposure time range allowed to be used by the algorithm.
	//
	MaxExposureTime float32 `xml:"http://www.onvif.org/ver10/schema MaxExposureTime,omitempty"`

	//
	// Minimum value of the sensor gain range that is allowed to be used by the algorithm.
	//
	MinGain float32 `xml:"http://www.onvif.org/ver10/schema MinGain,omitempty"`

	//
	// Maximum value of the sensor gain range that is allowed to be used by the algorithm.
	//
	MaxGain float32 `xml:"http://www.onvif.org/ver10/schema MaxGain,omitempty"`

	//
	// Minimum value of the iris range allowed to be used by the algorithm.
	//
	MinIris float32 `xml:"http://www.onvif.org/ver10/schema MinIris,omitempty"`

	//
	// Maximum value of the iris range allowed to be used by the algorithm.
	//
	MaxIris float32 `xml:"http://www.onvif.org/ver10/schema MaxIris,omitempty"`

	//
	// The fixed exposure time used by the image sensor (μs).
	//
	ExposureTime float32 `xml:"http://www.onvif.org/ver10/schema ExposureTime,omitempty"`

	//
	// The fixed gain used by the image sensor (dB).
	//
	Gain float32 `xml:"http://www.onvif.org/ver10/schema Gain,omitempty"`

	//
	// The fixed attenuation of input light affected by the iris (dB). 0dB maps to a fully opened iris.
	//
	Iris float32 `xml:"http://www.onvif.org/ver10/schema Iris,omitempty"`
}

// WideDynamicRange type
type WideDynamicRange struct {

	//
	// White dynamic range (on/off)
	//
	Mode WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// Optional level parameter (unitless)
	//
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// BacklightCompensation type
type BacklightCompensation struct {

	// Backlight compensation mode (on/off).
	Mode BacklightCompensationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Optional level parameter (unit unspecified).
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// ImagingOptions type
type ImagingOptions struct {
	BacklightCompensation BacklightCompensationOptions `xml:"http://www.onvif.org/ver10/schema BacklightCompensation,omitempty"`

	Brightness FloatRange `xml:"http://www.onvif.org/ver10/schema Brightness,omitempty"`

	ColorSaturation FloatRange `xml:"http://www.onvif.org/ver10/schema ColorSaturation,omitempty"`

	Contrast FloatRange `xml:"http://www.onvif.org/ver10/schema Contrast,omitempty"`

	Exposure ExposureOptions `xml:"http://www.onvif.org/ver10/schema Exposure,omitempty"`

	Focus FocusOptions `xml:"http://www.onvif.org/ver10/schema Focus,omitempty"`

	IrCutFilterModes []IrCutFilterMode `xml:"http://www.onvif.org/ver10/schema IrCutFilterModes,omitempty"`

	Sharpness FloatRange `xml:"http://www.onvif.org/ver10/schema Sharpness,omitempty"`

	WideDynamicRange WideDynamicRangeOptions `xml:"http://www.onvif.org/ver10/schema WideDynamicRange,omitempty"`

	WhiteBalance WhiteBalanceOptions `xml:"http://www.onvif.org/ver10/schema WhiteBalance,omitempty"`
}

// WideDynamicRangeOptions type
type WideDynamicRangeOptions struct {
	Mode []WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Level FloatRange `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// BacklightCompensationOptions type
type BacklightCompensationOptions struct {
	Mode []WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Level FloatRange `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// FocusOptions type
type FocusOptions struct {
	AutoFocusModes []AutoFocusMode `xml:"http://www.onvif.org/ver10/schema AutoFocusModes,omitempty"`

	DefaultSpeed FloatRange `xml:"http://www.onvif.org/ver10/schema DefaultSpeed,omitempty"`

	NearLimit FloatRange `xml:"http://www.onvif.org/ver10/schema NearLimit,omitempty"`

	FarLimit FloatRange `xml:"http://www.onvif.org/ver10/schema FarLimit,omitempty"`
}

// ExposureOptions type
type ExposureOptions struct {
	Mode []ExposureMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Priority []ExposurePriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	MinExposureTime FloatRange `xml:"http://www.onvif.org/ver10/schema MinExposureTime,omitempty"`

	MaxExposureTime FloatRange `xml:"http://www.onvif.org/ver10/schema MaxExposureTime,omitempty"`

	MinGain FloatRange `xml:"http://www.onvif.org/ver10/schema MinGain,omitempty"`

	MaxGain FloatRange `xml:"http://www.onvif.org/ver10/schema MaxGain,omitempty"`

	MinIris FloatRange `xml:"http://www.onvif.org/ver10/schema MinIris,omitempty"`

	MaxIris FloatRange `xml:"http://www.onvif.org/ver10/schema MaxIris,omitempty"`

	ExposureTime FloatRange `xml:"http://www.onvif.org/ver10/schema ExposureTime,omitempty"`

	Gain FloatRange `xml:"http://www.onvif.org/ver10/schema Gain,omitempty"`

	Iris FloatRange `xml:"http://www.onvif.org/ver10/schema Iris,omitempty"`
}

// WhiteBalanceOptions type
type WhiteBalanceOptions struct {
	Mode []WhiteBalanceMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	YrGain FloatRange `xml:"http://www.onvif.org/ver10/schema YrGain,omitempty"`

	YbGain FloatRange `xml:"http://www.onvif.org/ver10/schema YbGain,omitempty"`
}

// Removed FocusMove by fixgen.py

// AbsoluteFocus type
type AbsoluteFocus struct {

	//
	// Position parameter for the absolute focus control.
	//
	Position float32 `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`

	//
	// Speed parameter for the absolute focus control.
	//
	Speed float32 `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`
}

// RelativeFocus type
type RelativeFocus struct {

	//
	// Distance parameter for the relative focus control.
	//
	Distance float32 `xml:"http://www.onvif.org/ver10/schema Distance,omitempty"`

	//
	// Speed parameter for the relative focus control.
	//
	Speed float32 `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`
}

// ContinuousFocus type
type ContinuousFocus struct {

	//
	// Speed parameter for the Continuous focus control.
	//
	Speed float32 `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`
}

// Removed MoveOptions by fixgen.py

// AbsoluteFocusOptions type
type AbsoluteFocusOptions struct {

	//
	// Valid ranges of the position.
	//
	Position FloatRange `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`

	//
	// Valid ranges of the speed.
	//
	Speed FloatRange `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`
}

// RelativeFocusOptions type
type RelativeFocusOptions struct {

	//
	// Valid ranges of the distance.
	//
	Distance FloatRange `xml:"http://www.onvif.org/ver10/schema Distance,omitempty"`

	//
	// Valid ranges of the speed.
	//
	Speed FloatRange `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`
}

// ContinuousFocusOptions type
type ContinuousFocusOptions struct {

	//
	// Valid ranges of the speed.
	//
	Speed FloatRange `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`
}

// WhiteBalance type
type WhiteBalance struct {

	// Auto whitebalancing mode (auto/manual).
	Mode WhiteBalanceMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Rgain (unitless).
	CrGain float32 `xml:"http://www.onvif.org/ver10/schema CrGain,omitempty"`

	// Bgain (unitless).
	CbGain float32 `xml:"http://www.onvif.org/ver10/schema CbGain,omitempty"`
}

// ImagingStatus20 type
type ImagingStatus20 struct {

	//
	// Status of focus.
	//
	FocusStatus20 FocusStatus20 `xml:"http://www.onvif.org/ver10/schema FocusStatus20,omitempty"`

	Extension ImagingStatus20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingStatus20Extension type
type ImagingStatus20Extension struct {
}

// FocusStatus20 type
type FocusStatus20 struct {

	//
	// Status of focus position.
	//
	Position float32 `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`

	//
	// Status of focus MoveStatus.
	//
	MoveStatus MoveStatus `xml:"http://www.onvif.org/ver10/schema MoveStatus,omitempty"`

	//
	// Error status of focus.
	//
	Error string `xml:"http://www.onvif.org/ver10/device/wsdl Error,omitempty"`

	Extension FocusStatus20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// FocusStatus20Extension type
type FocusStatus20Extension struct {
}

// ImagingSettings20 type
type ImagingSettings20 struct {

	// Enabled/disabled BLC mode (on/off).
	BacklightCompensation BacklightCompensation20 `xml:"http://www.onvif.org/ver10/schema BacklightCompensation,omitempty"`

	// Image brightness (unit unspecified).
	Brightness float32 `xml:"http://www.onvif.org/ver10/schema Brightness,omitempty"`

	// Color saturation of the image (unit unspecified).
	ColorSaturation float32 `xml:"http://www.onvif.org/ver10/schema ColorSaturation,omitempty"`

	// Contrast of the image (unit unspecified).
	Contrast float32 `xml:"http://www.onvif.org/ver10/schema Contrast,omitempty"`

	// Exposure mode of the device.
	Exposure Exposure20 `xml:"http://www.onvif.org/ver10/schema Exposure,omitempty"`

	// Focus configuration.
	Focus FocusConfiguration20 `xml:"http://www.onvif.org/ver10/schema Focus,omitempty"`

	// Infrared Cutoff Filter settings.
	IrCutFilter IrCutFilterMode `xml:"http://www.onvif.org/ver10/schema IrCutFilter,omitempty"`

	// Sharpness of the Video image.
	Sharpness float32 `xml:"http://www.onvif.org/ver10/schema Sharpness,omitempty"`

	// WDR settings.
	WideDynamicRange WideDynamicRange20 `xml:"http://www.onvif.org/ver10/schema WideDynamicRange,omitempty"`

	// White balance settings.
	WhiteBalance WhiteBalance20 `xml:"http://www.onvif.org/ver10/schema WhiteBalance,omitempty"`

	Extension ImagingSettingsExtension20 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingSettingsExtension20 type
type ImagingSettingsExtension20 struct {

	// Optional element to configure Image Stabilization feature.
	ImageStabilization ImageStabilization `xml:"http://www.onvif.org/ver10/schema ImageStabilization,omitempty"`

	Extension ImagingSettingsExtension202 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingSettingsExtension202 type
type ImagingSettingsExtension202 struct {

	// An optional parameter applied to only auto mode to adjust timing of toggling Ir cut filter.
	IrCutFilterAutoAdjustment []IrCutFilterAutoAdjustment `xml:"http://www.onvif.org/ver10/schema IrCutFilterAutoAdjustment,omitempty"`

	Extension ImagingSettingsExtension203 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingSettingsExtension203 type
type ImagingSettingsExtension203 struct {

	// Optional element to configure Image Contrast Compensation.
	ToneCompensation ToneCompensation `xml:"http://www.onvif.org/ver10/schema ToneCompensation,omitempty"`

	// Optional element to configure Image Defogging.
	Defogging Defogging `xml:"http://www.onvif.org/ver10/schema Defogging,omitempty"`

	// Optional element to configure Image Noise Reduction.
	NoiseReduction NoiseReduction `xml:"http://www.onvif.org/ver10/schema NoiseReduction,omitempty"`

	Extension ImagingSettingsExtension204 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingSettingsExtension204 type
type ImagingSettingsExtension204 struct {
}

// ImageStabilization type
type ImageStabilization struct {

	// Parameter to enable/disable Image Stabilization feature.
	Mode ImageStabilizationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Optional level parameter (unit unspecified)
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`

	Extension ImageStabilizationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImageStabilizationExtension type
type ImageStabilizationExtension struct {
}

// IrCutFilterAutoAdjustment type
type IrCutFilterAutoAdjustment struct {

	// Specifies which boundaries to automatically toggle Ir cut filter following parameters are applied to. Its options shall be chosen from tt:IrCutFilterAutoBoundaryType.
	BoundaryType string `xml:"http://www.onvif.org/ver10/device/wsdl BoundaryType,omitempty"`

	// Adjusts boundary exposure level for toggling Ir cut filter to on/off specified with unitless normalized value from +1.0 to -1.0. Zero is default and -1.0 is the darkest adjustment (Unitless).
	BoundaryOffset float32 `xml:"http://www.onvif.org/ver10/schema BoundaryOffset,omitempty"`

	// Delay time of toggling Ir cut filter to on/off after crossing of the boundary exposure levels.
	ResponseTime Duration `xml:"http://www.onvif.org/ver10/schema ResponseTime,omitempty"`

	Extension IrCutFilterAutoAdjustmentExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IrCutFilterAutoAdjustmentExtension type
type IrCutFilterAutoAdjustmentExtension struct {
}

// WideDynamicRange20 type
type WideDynamicRange20 struct {

	// Wide dynamic range mode (on/off).
	Mode WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Optional level parameter (unit unspecified).
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// BacklightCompensation20 type
type BacklightCompensation20 struct {

	// Backlight compensation mode (on/off).
	Mode BacklightCompensationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Optional level parameter (unit unspecified).
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// Exposure20 type
type Exposure20 struct {

	//
	// Exposure Mode
	//
	//
	Mode ExposureMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// The exposure priority mode (low noise/framerate).
	//
	Priority ExposurePriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	//
	// Rectangular exposure mask.
	//
	Window Rectangle `xml:"http://www.onvif.org/ver10/schema Window,omitempty"`

	//
	// Minimum value of exposure time range allowed to be used by the algorithm.
	//
	MinExposureTime float32 `xml:"http://www.onvif.org/ver10/schema MinExposureTime,omitempty"`

	//
	// Maximum value of exposure time range allowed to be used by the algorithm.
	//
	MaxExposureTime float32 `xml:"http://www.onvif.org/ver10/schema MaxExposureTime,omitempty"`

	//
	// Minimum value of the sensor gain range that is allowed to be used by the algorithm.
	//
	MinGain float32 `xml:"http://www.onvif.org/ver10/schema MinGain,omitempty"`

	//
	// Maximum value of the sensor gain range that is allowed to be used by the algorithm.
	//
	MaxGain float32 `xml:"http://www.onvif.org/ver10/schema MaxGain,omitempty"`

	//
	// Minimum value of the iris range allowed to be used by the algorithm.  0dB maps to a fully opened iris and positive values map to higher attenuation.
	//
	MinIris float32 `xml:"http://www.onvif.org/ver10/schema MinIris,omitempty"`

	//
	// Maximum value of the iris range allowed to be used by the algorithm. 0dB maps to a fully opened iris and positive values map to higher attenuation.
	//
	MaxIris float32 `xml:"http://www.onvif.org/ver10/schema MaxIris,omitempty"`

	//
	// The fixed exposure time used by the image sensor (μs).
	//
	ExposureTime float32 `xml:"http://www.onvif.org/ver10/schema ExposureTime,omitempty"`

	//
	// The fixed gain used by the image sensor (dB).
	//
	Gain float32 `xml:"http://www.onvif.org/ver10/schema Gain,omitempty"`

	//
	// The fixed attenuation of input light affected by the iris (dB). 0dB maps to a fully opened iris and positive values map to higher attenuation.
	//
	Iris float32 `xml:"http://www.onvif.org/ver10/schema Iris,omitempty"`
}

// ToneCompensation type
type ToneCompensation struct {

	// Parameter to enable/disable or automatic ToneCompensation feature. Its options shall be chosen from tt:ToneCompensationMode Type.
	Mode string `xml:"http://www.onvif.org/ver10/device/wsdl Mode,omitempty"`

	// Optional level parameter specified with unitless normalized value from 0.0 to +1.0.
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`

	Extension ToneCompensationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ToneCompensationExtension type
type ToneCompensationExtension struct {
}

// Defogging type
type Defogging struct {

	// Parameter to enable/disable or automatic Defogging feature. Its options shall be chosen from tt:DefoggingMode Type.
	Mode string `xml:"http://www.onvif.org/ver10/device/wsdl Mode,omitempty"`

	// Optional level parameter specified with unitless normalized value from 0.0 to +1.0.
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`

	Extension DefoggingExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// DefoggingExtension type
type DefoggingExtension struct {
}

// NoiseReduction type
type NoiseReduction struct {

	// Level parameter specified with unitless normalized value from 0.0 to +1.0. Level=0 means no noise reduction or minimal noise reduction.
	Level float32 `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// ImagingOptions20 type
type ImagingOptions20 struct {

	//
	// Valid range of Backlight Compensation.
	//
	BacklightCompensation BacklightCompensationOptions20 `xml:"http://www.onvif.org/ver10/schema BacklightCompensation,omitempty"`

	//
	// Valid range of Brightness.
	//
	Brightness FloatRange `xml:"http://www.onvif.org/ver10/schema Brightness,omitempty"`

	//
	// Valid range of Color Saturation.
	//
	ColorSaturation FloatRange `xml:"http://www.onvif.org/ver10/schema ColorSaturation,omitempty"`

	//
	// Valid range of Contrast.
	//
	Contrast FloatRange `xml:"http://www.onvif.org/ver10/schema Contrast,omitempty"`

	//
	// Valid range of Exposure.
	//
	Exposure ExposureOptions20 `xml:"http://www.onvif.org/ver10/schema Exposure,omitempty"`

	//
	// Valid range of Focus.
	//
	Focus FocusOptions20 `xml:"http://www.onvif.org/ver10/schema Focus,omitempty"`

	//
	// Valid range of IrCutFilterModes.
	//
	IrCutFilterModes []IrCutFilterMode `xml:"http://www.onvif.org/ver10/schema IrCutFilterModes,omitempty"`

	//
	// Valid range of Sharpness.
	//
	Sharpness FloatRange `xml:"http://www.onvif.org/ver10/schema Sharpness,omitempty"`

	//
	// Valid range of WideDynamicRange.
	//
	WideDynamicRange WideDynamicRangeOptions20 `xml:"http://www.onvif.org/ver10/schema WideDynamicRange,omitempty"`

	//
	// Valid range of WhiteBalance.
	//
	WhiteBalance WhiteBalanceOptions20 `xml:"http://www.onvif.org/ver10/schema WhiteBalance,omitempty"`

	Extension ImagingOptions20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingOptions20Extension type
type ImagingOptions20Extension struct {

	// Options of parameters for Image Stabilization feature.
	ImageStabilization ImageStabilizationOptions `xml:"http://www.onvif.org/ver10/schema ImageStabilization,omitempty"`

	Extension ImagingOptions20Extension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingOptions20Extension2 type
type ImagingOptions20Extension2 struct {

	// Options of parameters for adjustment of Ir cut filter auto mode.
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustmentOptions `xml:"http://www.onvif.org/ver10/schema IrCutFilterAutoAdjustment,omitempty"`

	Extension ImagingOptions20Extension3 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingOptions20Extension3 type
type ImagingOptions20Extension3 struct {

	// Options of parameters for Tone Compensation feature.
	ToneCompensationOptions ToneCompensationOptions `xml:"http://www.onvif.org/ver10/schema ToneCompensationOptions,omitempty"`

	// Options of parameters for Defogging feature.
	DefoggingOptions DefoggingOptions `xml:"http://www.onvif.org/ver10/schema DefoggingOptions,omitempty"`

	// Options of parameter for Noise Reduction feature.
	NoiseReductionOptions NoiseReductionOptions `xml:"http://www.onvif.org/ver10/schema NoiseReductionOptions,omitempty"`

	Extension ImagingOptions20Extension4 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImagingOptions20Extension4 type
type ImagingOptions20Extension4 struct {
}

// ImageStabilizationOptions type
type ImageStabilizationOptions struct {

	// Supported options of Image Stabilization mode parameter.
	Mode []ImageStabilizationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Valid range of the Image Stabilization.
	Level FloatRange `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`

	Extension ImageStabilizationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ImageStabilizationOptionsExtension type
type ImageStabilizationOptionsExtension struct {
}

// IrCutFilterAutoAdjustmentOptions type
type IrCutFilterAutoAdjustmentOptions struct {

	// Supported options of boundary types for adjustment of Ir cut filter auto mode. The opptions shall be chosen from tt:IrCutFilterAutoBoundaryType.
	BoundaryType []string `xml:"http://www.onvif.org/ver10/device/wsdl BoundaryType,omitempty"`

	// Indicates whether or not boundary offset for toggling Ir cut filter is supported.
	BoundaryOffset bool `xml:"http://www.onvif.org/ver10/device/wsdl BoundaryOffset,omitempty"`

	// Supported range of delay time for toggling Ir cut filter.
	ResponseTimeRange DurationRange `xml:"http://www.onvif.org/ver10/schema ResponseTimeRange,omitempty"`

	Extension IrCutFilterAutoAdjustmentOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// IrCutFilterAutoAdjustmentOptionsExtension type
type IrCutFilterAutoAdjustmentOptionsExtension struct {
}

// WideDynamicRangeOptions20 type
type WideDynamicRangeOptions20 struct {
	Mode []WideDynamicMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Level FloatRange `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// BacklightCompensationOptions20 type
type BacklightCompensationOptions20 struct {

	//
	// 'ON' or 'OFF'
	//
	Mode []BacklightCompensationMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// Level range of BacklightCompensation.
	//
	Level FloatRange `xml:"http://www.onvif.org/ver10/schema Level,omitempty"`
}

// ExposureOptions20 type
type ExposureOptions20 struct {

	//
	// Exposure Mode
	//
	//
	Mode []ExposureMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// The exposure priority mode (low noise/framerate).
	//
	Priority []ExposurePriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	//
	// Valid range of the Minimum ExposureTime.
	//
	MinExposureTime FloatRange `xml:"http://www.onvif.org/ver10/schema MinExposureTime,omitempty"`

	//
	// Valid range of the Maximum ExposureTime.
	//
	MaxExposureTime FloatRange `xml:"http://www.onvif.org/ver10/schema MaxExposureTime,omitempty"`

	//
	// Valid range of the Minimum Gain.
	//
	MinGain FloatRange `xml:"http://www.onvif.org/ver10/schema MinGain,omitempty"`

	//
	// Valid range of the Maximum Gain.
	//
	MaxGain FloatRange `xml:"http://www.onvif.org/ver10/schema MaxGain,omitempty"`

	//
	// Valid range of the Minimum Iris.
	//
	MinIris FloatRange `xml:"http://www.onvif.org/ver10/schema MinIris,omitempty"`

	//
	// Valid range of the Maximum Iris.
	//
	MaxIris FloatRange `xml:"http://www.onvif.org/ver10/schema MaxIris,omitempty"`

	//
	// Valid range of the ExposureTime.
	//
	ExposureTime FloatRange `xml:"http://www.onvif.org/ver10/schema ExposureTime,omitempty"`

	//
	// Valid range of the Gain.
	//
	Gain FloatRange `xml:"http://www.onvif.org/ver10/schema Gain,omitempty"`

	//
	// Valid range of the Iris.
	//
	Iris FloatRange `xml:"http://www.onvif.org/ver10/schema Iris,omitempty"`
}

// Removed MoveOptions20 by fixgen.py

// Removed RelativeFocusOptions20 by fixgen.py

// WhiteBalance20 type
type WhiteBalance20 struct {

	//
	// 'AUTO' or 'MANUAL'
	//
	Mode WhiteBalanceMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	//
	// Rgain (unitless).
	//
	CrGain float32 `xml:"http://www.onvif.org/ver10/schema CrGain,omitempty"`

	//
	// Bgain (unitless).
	//
	CbGain float32 `xml:"http://www.onvif.org/ver10/schema CbGain,omitempty"`

	Extension WhiteBalance20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// WhiteBalance20Extension type
type WhiteBalance20Extension struct {
}

// FocusConfiguration20 type
type FocusConfiguration20 struct {

	//
	// Mode of auto focus.
	//
	// Note: for devices supporting both manual and auto operation at the same time manual operation may be supported even if the Mode parameter is set to Auto.
	//
	AutoFocusMode AutoFocusMode `xml:"http://www.onvif.org/ver10/schema AutoFocusMode,omitempty"`

	DefaultSpeed float32 `xml:"http://www.onvif.org/ver10/schema DefaultSpeed,omitempty"`

	// Parameter to set autofocus near limit (unit: meter).
	NearLimit float32 `xml:"http://www.onvif.org/ver10/schema NearLimit,omitempty"`

	// Parameter to set autofocus far limit (unit: meter).
	FarLimit float32 `xml:"http://www.onvif.org/ver10/schema FarLimit,omitempty"`

	Extension FocusConfiguration20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Zero or more modes as defined in enumeration tt:AFModes.

	AFMode StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl AFMode,attr,omitempty"`
}

// FocusConfiguration20Extension type
type FocusConfiguration20Extension struct {
}

// WhiteBalanceOptions20 type
type WhiteBalanceOptions20 struct {

	//
	// Mode of WhiteBalance.
	//
	//
	Mode []WhiteBalanceMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	YrGain FloatRange `xml:"http://www.onvif.org/ver10/schema YrGain,omitempty"`

	YbGain FloatRange `xml:"http://www.onvif.org/ver10/schema YbGain,omitempty"`

	Extension WhiteBalanceOptions20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// WhiteBalanceOptions20Extension type
type WhiteBalanceOptions20Extension struct {
}

// FocusOptions20 type
type FocusOptions20 struct {

	//
	// Supported modes for auto focus.
	//
	//
	AutoFocusModes []AutoFocusMode `xml:"http://www.onvif.org/ver10/schema AutoFocusModes,omitempty"`

	//
	// Valid range of DefaultSpeed.
	//
	DefaultSpeed FloatRange `xml:"http://www.onvif.org/ver10/schema DefaultSpeed,omitempty"`

	//
	// Valid range of NearLimit.
	//
	NearLimit FloatRange `xml:"http://www.onvif.org/ver10/schema NearLimit,omitempty"`

	//
	// Valid range of FarLimit.
	//
	FarLimit FloatRange `xml:"http://www.onvif.org/ver10/schema FarLimit,omitempty"`

	Extension FocusOptions20Extension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// FocusOptions20Extension type
type FocusOptions20Extension struct {

	// Supported options for auto focus. Options shall be chosen from tt:AFModes.
	AFModes StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl AFModes,omitempty"`
}

// ToneCompensationOptions type
type ToneCompensationOptions struct {

	// Supported options for Tone Compensation mode. Its options shall be chosen from tt:ToneCompensationMode Type.
	Mode []string `xml:"http://www.onvif.org/ver10/device/wsdl Mode,omitempty"`

	// Indicates whether or not support Level parameter for Tone Compensation.
	Level bool `xml:"http://www.onvif.org/ver10/device/wsdl Level,omitempty"`
}

// DefoggingOptions type
type DefoggingOptions struct {

	// Supported options for Defogging mode. Its options shall be chosen from tt:DefoggingMode Type.
	Mode []string `xml:"http://www.onvif.org/ver10/device/wsdl Mode,omitempty"`

	// Indicates whether or not support Level parameter for Defogging.
	Level bool `xml:"http://www.onvif.org/ver10/device/wsdl Level,omitempty"`
}

// NoiseReductionOptions type
type NoiseReductionOptions struct {

	// Indicates whether or not support Level parameter for NoiseReduction.
	Level bool `xml:"http://www.onvif.org/ver10/device/wsdl Level,omitempty"`
}

// MessageExtension type
type MessageExtension struct {
}

// ItemList type
type ItemList struct {
	SimpleItem []struct {

		// Item name.

		Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,attr,omitempty"`

		// Item value. The type is defined in the corresponding description.

		Value AnySimpleType `xml:"Value,attr,omitempty"`
	} `xml:"SimpleItem,omitempty"`

	ElementItem []struct {

		// Item name.

		Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,attr,omitempty"`
	} `xml:"ElementItem,omitempty"`

	Extension ItemListExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ItemListExtension type
type ItemListExtension struct {
}

// MessageDescription type
type MessageDescription struct {

	// Set of tokens producing this message. The list may only contain SimpleItemDescription items.
	// The set of tokens identify the component within the WS-Endpoint, which is responsible for the producing the message.
	// For analytics events the token set shall include the VideoSourceConfigurationToken, the VideoAnalyticsConfigurationToken
	// and the name of the analytics module or rule.
	//
	Source ItemListDescription `xml:"http://www.onvif.org/ver10/schema Source,omitempty"`

	// Describes optional message payload parameters that may be used as key. E.g. object IDs of tracked objects are conveyed as key.
	Key ItemListDescription `xml:"http://www.onvif.org/ver10/schema Key,omitempty"`

	// Describes the payload of the message.
	Data ItemListDescription `xml:"http://www.onvif.org/ver10/schema Data,omitempty"`

	Extension MessageDescriptionExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Must be set to true when the described Message relates to a property. An alternative term of "property" is a "state" in contrast to a pure event, which contains relevant information for only a single point in time.Default is false.

	IsProperty bool `xml:"http://www.onvif.org/ver10/device/wsdl IsProperty,attr,omitempty"`
}

// MessageDescriptionExtension type
type MessageDescriptionExtension struct {
}

// ItemListDescription type
type ItemListDescription struct {
	SimpleItemDescription []struct {

		// Item name. Must be unique within a list.

		Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,attr,omitempty"`

		Type QName `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
	} `xml:"SimpleItemDescription,omitempty"`

	ElementItemDescription []struct {

		// Item name. Must be unique within a list.

		Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,attr,omitempty"`

		// The type of the item. The Type must reference a defined type.

		Type QName `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
	} `xml:"ElementItemDescription,omitempty"`

	Extension ItemListDescriptionExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ItemListDescriptionExtension type
type ItemListDescriptionExtension struct {
}

// Polyline type
type Polyline struct {
	Point []Vector `xml:"http://www.onvif.org/ver10/schema Point,omitempty"`
}

// AnalyticsEngineConfiguration type
type AnalyticsEngineConfiguration struct {
	AnalyticsModule []Config `xml:"http://www.onvif.org/ver10/schema AnalyticsModule,omitempty"`

	Extension AnalyticsEngineConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AnalyticsEngineConfigurationExtension type
type AnalyticsEngineConfigurationExtension struct {
}

// RuleEngineConfiguration type
type RuleEngineConfiguration struct {
	Rule []Config `xml:"http://www.onvif.org/ver10/schema Rule,omitempty"`

	Extension RuleEngineConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// RuleEngineConfigurationExtension type
type RuleEngineConfigurationExtension struct {
}

// Config type
type Config struct {

	// List of configuration parameters as defined in the correspding description.
	Parameters ItemList `xml:"http://www.onvif.org/ver10/schema Parameters,omitempty"`

	// Name of the configuration.

	Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,attr,omitempty"`

	// The Type attribute specifies the type of rule and shall be equal to value of one of Name attributes of ConfigDescription elements returned by GetSupportedRules and GetSupportedAnalyticsModules command.

	Type QName `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
}

// ConfigDescription type
type ConfigDescription struct {

	//
	// List describing the configuration parameters. The names of the parameters must be unique. If possible SimpleItems
	// should be used to transport the information to ease parsing of dynamically defined messages by a client
	// application.
	//
	Parameters ItemListDescription `xml:"http://www.onvif.org/ver10/schema Parameters,omitempty"`

	Messages []struct {
		*MessageDescription

		//
		// The ParentTopic labels the message (e.g. "nn:RuleEngine/LineCrossing"). The real message can extend the ParentTopic
		// by for example the name of the instaniated rule (e.g. "nn:RuleEngine/LineCrossing/corssMyFirstLine").
		// Even without knowing the complete topic name, the subscriber will be able to distiguish the
		// messages produced by different rule instances of the same type via the Source fields of the message.
		// There the name of the rule instance, which produced the message, must be listed.
		//
		ParentTopic string `xml:"http://www.onvif.org/ver10/device/wsdl ParentTopic,omitempty"`
	} `xml:"Messages,omitempty"`

	Extension ConfigDescriptionExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// The Name attribute (e.g. "tt::LineDetector") uniquely identifies the type of rule, not a type definition in a schema.

	Name QName `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

	// The fixed attribute signals that it is not allowed to add or remove this type of configuration.

	Fixed bool `xml:"fixed,attr,omitempty"`

	// The maxInstances attribute signals the maximum number of instances per configuration.

	MaxInstances int32 `xml:"maxInstances,attr,omitempty"`
}

// ConfigDescriptionExtension type
type ConfigDescriptionExtension struct {
}

// SupportedRules type
type SupportedRules struct {

	// Lists the location of all schemas that are referenced in the rules.
	RuleContentSchemaLocation []AnyURI `xml:"http://www.onvif.org/ver10/schema RuleContentSchemaLocation,omitempty"`

	// List of rules supported by the Video Analytics configuration..
	RuleDescription []ConfigDescription `xml:"http://www.onvif.org/ver10/schema RuleDescription,omitempty"`

	Extension SupportedRulesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SupportedRulesExtension type
type SupportedRulesExtension struct {
}

// SupportedAnalyticsModules type
type SupportedAnalyticsModules struct {

	// It optionally contains a list of URLs that provide the location of schema files.
	// These schema files describe the types and elements used in the analytics module descriptions.
	// Analytics module descriptions that reference types or elements imported from any ONVIF defined schema files
	// need not explicitly list those schema files.
	AnalyticsModuleContentSchemaLocation []AnyURI `xml:"http://www.onvif.org/ver10/schema AnalyticsModuleContentSchemaLocation,omitempty"`

	AnalyticsModuleDescription []ConfigDescription `xml:"http://www.onvif.org/ver10/schema AnalyticsModuleDescription,omitempty"`

	Extension SupportedAnalyticsModulesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SupportedAnalyticsModulesExtension type
type SupportedAnalyticsModulesExtension struct {
}

// Removed PolygonConfiguration by fixgen.py

// PolylineArray type
type PolylineArray struct {

	// Contains array of Polyline
	Segment []Polyline `xml:"http://www.onvif.org/ver10/schema Segment,omitempty"`

	Extension PolylineArrayExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PolylineArrayExtension type
type PolylineArrayExtension struct {
}

// Removed PolylineArrayConfiguration by fixgen.py

// MotionExpression type
type MotionExpression struct {

	// Motion Expression data structure contains motion expression which is based on Scene Descriptor schema with XPATH syntax. The Type argument could allow introduction of different dialects
	Expression string `xml:"http://www.onvif.org/ver10/device/wsdl Expression,omitempty"`

	Type string `xml:"http://www.onvif.org/ver10/device/wsdl Type,attr,omitempty"`
}

// Removed MotionExpressionConfiguration by fixgen.py

// Removed CellLayout by fixgen.py

// Removed PaneConfiguration by fixgen.py

// PaneLayout type
type PaneLayout struct {

	// Reference to the configuration of the streaming and coding parameters.
	Pane ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Pane,omitempty"`

	// Describes the location and size of the area on the monitor. The area coordinate values are espressed in normalized units [-1.0, 1.0].
	Area Rectangle `xml:"http://www.onvif.org/ver10/schema Area,omitempty"`
}

// Layout type
type Layout struct {

	// List of panes assembling the display layout.
	PaneLayout []PaneLayout `xml:"http://www.onvif.org/ver10/schema PaneLayout,omitempty"`

	Extension LayoutExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// LayoutExtension type
type LayoutExtension struct {
}

// Removed CodingCapabilities by fixgen.py

// LayoutOptions type
type LayoutOptions struct {

	// Lists the possible Pane Layouts of the Video Output
	PaneLayoutOptions []PaneLayoutOptions `xml:"http://www.onvif.org/ver10/schema PaneLayoutOptions,omitempty"`

	Extension LayoutOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// LayoutOptionsExtension type
type LayoutOptionsExtension struct {
}

// PaneLayoutOptions type
type PaneLayoutOptions struct {

	// List of areas assembling a layout. Coordinate values are in the range [-1.0, 1.0].
	Area []Rectangle `xml:"http://www.onvif.org/ver10/schema Area,omitempty"`

	Extension PaneOptionExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PaneOptionExtension type
type PaneOptionExtension struct {
}

// Receiver type
type Receiver struct {

	// Unique identifier of the receiver.
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token,omitempty"`

	// Describes the configuration of the receiver.
	Configuration ReceiverConfiguration `xml:"http://www.onvif.org/ver10/schema Configuration,omitempty"`
}

// ReceiverConfiguration type
type ReceiverConfiguration struct {

	// The following connection modes are defined:
	Mode ReceiverMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Details of the URI to which the receiver should connect.
	MediaUri AnyURI `xml:"http://www.onvif.org/ver10/schema MediaUri,omitempty"`

	// Stream connection parameters.
	StreamSetup StreamSetup `xml:"http://www.onvif.org/ver10/schema StreamSetup,omitempty"`
}

// Removed ReceiverStateInformation by fixgen.py

// SourceReference type
type SourceReference struct {
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token,omitempty"`

	Type AnyURI `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
}

// DateTimeRange type
type DateTimeRange struct {
	From string `xml:"http://www.onvif.org/ver10/schema From,omitempty"`

	Until string `xml:"http://www.onvif.org/ver10/schema Until,omitempty"`
}

// Removed RecordingSummary by fixgen.py

// SearchScope type
type SearchScope struct {

	// A list of sources that are included in the scope. If this list is included, only data from one of these sources shall be searched.
	IncludedSources []SourceReference `xml:"http://www.onvif.org/ver10/schema IncludedSources,omitempty"`

	// A list of recordings that are included in the scope. If this list is included, only data from one of these recordings shall be searched.
	IncludedRecordings []RecordingReference `xml:"http://www.onvif.org/ver10/schema IncludedRecordings,omitempty"`

	// An xpath expression used to specify what recordings to search. Only those recordings with an RecordingInformation structure that matches the filter shall be searched.
	RecordingInformationFilter XPathExpression `xml:"http://www.onvif.org/ver10/schema RecordingInformationFilter,omitempty"`

	// Extension point
	Extension SearchScopeExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SearchScopeExtension type
type SearchScopeExtension struct {
}

// Removed EventFilter by fixgen.py

// Removed PTZPositionFilter by fixgen.py

// Removed MetadataFilter by fixgen.py

// Removed FindRecordingResultList by fixgen.py

// Removed FindEventResultList by fixgen.py

// Removed FindEventResult by fixgen.py

// Removed FindPTZPositionResultList by fixgen.py

// Removed FindPTZPositionResult by fixgen.py

// Removed FindMetadataResultList by fixgen.py

// Removed FindMetadataResult by fixgen.py

// Removed RecordingInformation by fixgen.py

// RecordingSourceInformation type
type RecordingSourceInformation struct {

	//
	// Identifier for the source chosen by the client that creates the structure.
	// This identifier is opaque to the device. Clients may use any type of URI for this field. A device shall support at least 128 characters.
	SourceId AnyURI `xml:"http://www.onvif.org/ver10/schema SourceId,omitempty"`

	// Informative user readable name of the source, e.g. "Camera23". A device shall support at least 20 characters.
	Name Name `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	// Informative description of the physical location of the source, e.g. the coordinates on a map.
	Location Description `xml:"http://www.onvif.org/ver10/schema Location,omitempty"`

	// Informative description of the source.
	Description Description `xml:"http://www.onvif.org/ver10/schema Description,omitempty"`

	// URI provided by the service supplying data to be recorded. A device shall support at least 128 characters.
	Address AnyURI `xml:"http://www.onvif.org/ver10/schema Address,omitempty"`
}

// TrackInformation type
type TrackInformation struct {
	TrackToken TrackReference `xml:"http://www.onvif.org/ver10/schema TrackToken,omitempty"`

	// Type of the track: "Video", "Audio" or "Metadata".
	// The track shall only be able to hold data of that type.
	TrackType TrackType `xml:"http://www.onvif.org/ver10/schema TrackType,omitempty"`

	// Informative description of the contents of the track.
	Description Description `xml:"http://www.onvif.org/ver10/schema Description,omitempty"`

	// The start date and time of the oldest recorded data in the track.
	DataFrom string `xml:"http://www.onvif.org/ver10/schema DataFrom,omitempty"`

	// The stop date and time of the newest recorded data in the track.
	DataTo string `xml:"http://www.onvif.org/ver10/schema DataTo,omitempty"`
}

// Removed MediaAttributes by fixgen.py

// TrackAttributes type
type TrackAttributes struct {

	// The basic information about the track. Note that a track may represent a single contiguous time span or consist of multiple slices.
	TrackInformation TrackInformation `xml:"http://www.onvif.org/ver10/schema TrackInformation,omitempty"`

	// If the track is a video track, exactly one of this structure shall be present and contain the video attributes.
	VideoAttributes VideoAttributes `xml:"http://www.onvif.org/ver10/schema VideoAttributes,omitempty"`

	// If the track is an audio track, exactly one of this structure shall be present and contain the audio attributes.
	AudioAttributes AudioAttributes `xml:"http://www.onvif.org/ver10/schema AudioAttributes,omitempty"`

	// If the track is an metadata track, exactly one of this structure shall be present and contain the metadata attributes.
	MetadataAttributes MetadataAttributes `xml:"http://www.onvif.org/ver10/schema MetadataAttributes,omitempty"`

	Extension TrackAttributesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// TrackAttributesExtension type
type TrackAttributesExtension struct {
}

// VideoAttributes type
type VideoAttributes struct {

	// Average bitrate in kbps.
	Bitrate int32 `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// The width of the video in pixels.
	Width int32 `xml:"http://www.onvif.org/ver10/schema Width,omitempty"`

	// The height of the video in pixels.
	Height int32 `xml:"http://www.onvif.org/ver10/schema Height,omitempty"`

	// Video encoding of the track.  Use value from tt:VideoEncoding for MPEG4. Otherwise use values from tt:VideoEncodingMimeNames and .
	Encoding string `xml:"http://www.onvif.org/ver10/device/wsdl Encoding,omitempty"`

	// Average framerate in frames per second.
	Framerate float32 `xml:"http://www.onvif.org/ver10/schema Framerate,omitempty"`
}

// AudioAttributes type
type AudioAttributes struct {

	// The bitrate in kbps.
	Bitrate int32 `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// Audio encoding of the track.  Use values from tt:AudioEncoding for G711 and AAC. Otherwise use values from tt:AudioEncodingMimeNames and .
	Encoding string `xml:"http://www.onvif.org/ver10/device/wsdl Encoding,omitempty"`

	// The sample rate in kHz.
	Samplerate int32 `xml:"http://www.onvif.org/ver10/schema Samplerate,omitempty"`
}

// MetadataAttributes type
type MetadataAttributes struct {

	// Indicates that there can be PTZ data in the metadata track in the specified time interval.
	CanContainPTZ bool `xml:"http://www.onvif.org/ver10/device/wsdl CanContainPTZ,omitempty"`

	// Indicates that there can be analytics data in the metadata track in the specified time interval.
	CanContainAnalytics bool `xml:"http://www.onvif.org/ver10/device/wsdl CanContainAnalytics,omitempty"`

	// Indicates that there can be notifications in the metadata track in the specified time interval.
	CanContainNotifications bool `xml:"http://www.onvif.org/ver10/device/wsdl CanContainNotifications,omitempty"`

	// List of all PTZ spaces active for recording. Note that events are only recorded on position changes and the actual point of recording may not necessarily contain an event of the specified type.

	PtzSpaces StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl PtzSpaces,attr,omitempty"`
}

// RecordingConfiguration type
type RecordingConfiguration struct {

	// Information about the source of the recording.
	Source RecordingSourceInformation `xml:"http://www.onvif.org/ver10/schema Source,omitempty"`

	// Informative description of the source.
	Content Description `xml:"http://www.onvif.org/ver10/schema Content,omitempty"`

	// Sspecifies the maximum time that data in any track within the
	// recording shall be stored. The device shall delete any data older than the maximum retention
	// time. Such data shall not be accessible anymore. If the MaximumRetentionPeriod is set to 0,
	// the device shall not limit the retention time of stored data, except by resource constraints.
	// Whatever the value of MaximumRetentionTime, the device may automatically delete
	// recordings to free up storage space for new recordings.
	MaximumRetentionTime Duration `xml:"http://www.onvif.org/ver10/schema MaximumRetentionTime,omitempty"`
}

// TrackConfiguration type
type TrackConfiguration struct {

	// Type of the track. It shall be equal to the strings “Video”,
	// “Audio” or “Metadata”. The track shall only be able to hold data of that type.
	TrackType TrackType `xml:"http://www.onvif.org/ver10/schema TrackType,omitempty"`

	// Informative description of the track.
	Description Description `xml:"http://www.onvif.org/ver10/schema Description,omitempty"`
}

// Removed GetRecordingsResponseItem by fixgen.py

// Removed GetTracksResponseList by fixgen.py

// Removed GetTracksResponseItem by fixgen.py

// RecordingJobConfiguration type
type RecordingJobConfiguration struct {

	// Identifies the recording to which this job shall store the received data.
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/schema RecordingToken,omitempty"`

	// The mode of the job. If it is idle, nothing shall happen. If it is active, the device shall try
	// to obtain data from the receivers. A client shall use GetRecordingJobState to determine if data transfer is really taking place.
	// The only valid values for Mode shall be “Idle” and “Active”.
	Mode RecordingJobMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// This shall be a non-negative number. If there are multiple recording jobs that store data to
	// the same track, the device will only store the data for the recording job with the highest
	// priority. The priority is specified per recording job, but the device shall determine the priority
	// of each track individually. If there are two recording jobs with the same priority, the device
	// shall record the data corresponding to the recording job that was activated the latest.
	Priority int32 `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	// Source of the recording.
	Source []RecordingJobSource `xml:"http://www.onvif.org/ver10/schema Source,omitempty"`

	Extension RecordingJobConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// This attribute adds an additional requirement for activating the recording job.
	// If this optional field is provided the job shall only record if the schedule exists and is active.
	//

	ScheduleToken string `xml:"http://www.onvif.org/ver10/device/wsdl ScheduleToken,attr,omitempty"`
}

// RecordingJobConfigurationExtension type
type RecordingJobConfigurationExtension struct {
}

// RecordingJobSource type
type RecordingJobSource struct {

	// This field shall be a reference to the source of the data. The type of the source
	// is determined by the attribute Type in the SourceToken structure. If Type is
	// http://www.onvif.org/ver10/schema/Receiver, the token is a ReceiverReference. In this case
	// the device shall receive the data over the network. If Type is
	// http://www.onvif.org/ver10/schema/Profile, the token identifies a media profile, instructing the
	// device to obtain data from a profile that exists on the local device.
	SourceToken SourceReference `xml:"http://www.onvif.org/ver10/schema SourceToken,omitempty"`

	// If this field is TRUE, and if the SourceToken is omitted, the device
	// shall create a receiver object (through the receiver service) and assign the
	// ReceiverReference to the SourceToken field. When retrieving the RecordingJobConfiguration
	// from the device, the AutoCreateReceiver field shall never be present.
	AutoCreateReceiver bool `xml:"http://www.onvif.org/ver10/device/wsdl AutoCreateReceiver,omitempty"`

	// List of tracks associated with the recording.
	Tracks []RecordingJobTrack `xml:"http://www.onvif.org/ver10/schema Tracks,omitempty"`

	Extension RecordingJobSourceExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// RecordingJobSourceExtension type
type RecordingJobSourceExtension struct {
}

// RecordingJobTrack type
type RecordingJobTrack struct {

	// If the received RTSP stream contains multiple tracks of the same type, the
	// SourceTag differentiates between those Tracks. This field can be ignored in case of recording a local source.
	SourceTag string `xml:"http://www.onvif.org/ver10/device/wsdl SourceTag,omitempty"`

	// The destination is the tracktoken of the track to which the device shall store the
	// received data.
	Destination TrackReference `xml:"http://www.onvif.org/ver10/schema Destination,omitempty"`
}

// RecordingJobStateInformation type
type RecordingJobStateInformation struct {

	// Identification of the recording that the recording job records to.
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/schema RecordingToken,omitempty"`

	// Holds the aggregated state over the whole RecordingJobInformation structure.
	State RecordingJobState `xml:"http://www.onvif.org/ver10/schema State,omitempty"`

	// Identifies the data source of the recording job.
	Sources []RecordingJobStateSource `xml:"http://www.onvif.org/ver10/schema Sources,omitempty"`

	Extension RecordingJobStateInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// RecordingJobStateInformationExtension type
type RecordingJobStateInformationExtension struct {
}

// RecordingJobStateSource type
type RecordingJobStateSource struct {

	// Identifies the data source of the recording job.
	SourceToken SourceReference `xml:"http://www.onvif.org/ver10/schema SourceToken,omitempty"`

	// Holds the aggregated state over all substructures of RecordingJobStateSource.
	State RecordingJobState `xml:"http://www.onvif.org/ver10/schema State,omitempty"`

	// List of track items.
	Tracks RecordingJobStateTracks `xml:"http://www.onvif.org/ver10/schema Tracks,omitempty"`
}

// RecordingJobStateTracks type
type RecordingJobStateTracks struct {
	Track []RecordingJobStateTrack `xml:"http://www.onvif.org/ver10/schema Track,omitempty"`
}

// RecordingJobStateTrack type
type RecordingJobStateTrack struct {

	// Identifies the track of the data source that provides the data.
	SourceTag string `xml:"http://www.onvif.org/ver10/device/wsdl SourceTag,omitempty"`

	// Indicates the destination track.
	Destination TrackReference `xml:"http://www.onvif.org/ver10/schema Destination,omitempty"`

	// Optionally holds an implementation defined string value that describes the error.
	// The string should be in the English language.
	Error string `xml:"http://www.onvif.org/ver10/device/wsdl Error,omitempty"`

	// Provides the job state of the track. The valid
	// values of state shall be “Idle”, “Active” and “Error”. If state equals “Error”, the Error field may be filled in with an implementation defined value.
	State RecordingJobState `xml:"http://www.onvif.org/ver10/schema State,omitempty"`
}

// Removed GetRecordingJobsResponseItem by fixgen.py

// Removed ReplayConfiguration by fixgen.py

// AnalyticsEngine type
type AnalyticsEngine struct {
	*ConfigurationEntity

	AnalyticsEngineConfiguration AnalyticsDeviceEngineConfiguration `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration,omitempty"`
}

// AnalyticsDeviceEngineConfiguration type
type AnalyticsDeviceEngineConfiguration struct {
	EngineConfiguration []EngineConfiguration `xml:"http://www.onvif.org/ver10/schema EngineConfiguration,omitempty"`

	Extension AnalyticsDeviceEngineConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AnalyticsDeviceEngineConfigurationExtension type
type AnalyticsDeviceEngineConfigurationExtension struct {
}

// EngineConfiguration type
type EngineConfiguration struct {
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `xml:"http://www.onvif.org/ver10/schema VideoAnalyticsConfiguration,omitempty"`

	AnalyticsEngineInputInfo AnalyticsEngineInputInfo `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineInputInfo,omitempty"`
}

// AnalyticsEngineInputInfo type
type AnalyticsEngineInputInfo struct {
	InputInfo Config `xml:"http://www.onvif.org/ver10/schema InputInfo,omitempty"`

	Extension AnalyticsEngineInputInfoExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AnalyticsEngineInputInfoExtension type
type AnalyticsEngineInputInfoExtension struct {
}

// AnalyticsEngineInput type
type AnalyticsEngineInput struct {
	*ConfigurationEntity

	SourceIdentification SourceIdentification `xml:"http://www.onvif.org/ver10/schema SourceIdentification,omitempty"`

	VideoInput VideoEncoderConfiguration `xml:"http://www.onvif.org/ver10/schema VideoInput,omitempty"`

	MetadataInput MetadataInput `xml:"http://www.onvif.org/ver10/schema MetadataInput,omitempty"`
}

// SourceIdentification type
type SourceIdentification struct {
	Name string `xml:"http://www.onvif.org/ver10/device/wsdl Name,omitempty"`

	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl Token,omitempty"`

	Extension SourceIdentificationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SourceIdentificationExtension type
type SourceIdentificationExtension struct {
}

// MetadataInput type
type MetadataInput struct {
	MetadataConfig []Config `xml:"http://www.onvif.org/ver10/schema MetadataConfig,omitempty"`

	Extension MetadataInputExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// MetadataInputExtension type
type MetadataInputExtension struct {
}

// Removed AnalyticsEngineControl by fixgen.py

// Removed AnalyticsStateInformation by fixgen.py

// Removed AnalyticsState by fixgen.py

// ActionEngineEventPayload type
type ActionEngineEventPayload struct {

	// Request Message
	RequestInfo Envelope `xml:"RequestInfo,omitempty"`

	// Response Message
	ResponseInfo Envelope `xml:"ResponseInfo,omitempty"`

	// Fault Message
	Fault Fault `xml:"Fault,omitempty"`

	Extension ActionEngineEventPayloadExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ActionEngineEventPayloadExtension type
type ActionEngineEventPayloadExtension struct {
}

// AudioClassCandidate type
type AudioClassCandidate struct {

	// Indicates audio class label
	Type AudioClassType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// A likelihood/probability that the corresponding audio event belongs to this class. The sum of the likelihoods shall NOT exceed 1
	Likelihood float32 `xml:"http://www.onvif.org/ver10/schema Likelihood,omitempty"`
}

// AudioClassDescriptor type
type AudioClassDescriptor struct {

	// Array of audio class label and class probability
	ClassCandidate []AudioClassCandidate `xml:"http://www.onvif.org/ver10/schema ClassCandidate,omitempty"`

	Extension AudioClassDescriptorExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AudioClassDescriptorExtension type
type AudioClassDescriptorExtension struct {
}

// ActiveConnection type
type ActiveConnection struct {
	CurrentBitrate float32 `xml:"http://www.onvif.org/ver10/schema CurrentBitrate,omitempty"`

	CurrentFps float32 `xml:"http://www.onvif.org/ver10/schema CurrentFps,omitempty"`
}

// ProfileStatus type
type ProfileStatus struct {
	ActiveConnections []ActiveConnection `xml:"http://www.onvif.org/ver10/schema ActiveConnections,omitempty"`

	Extension ProfileStatusExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ProfileStatusExtension type
type ProfileStatusExtension struct {
}

// OSDReference type
type OSDReference struct {
	Value ReferenceToken
}

// OSDPosConfiguration type
type OSDPosConfiguration struct {

	// For OSD position type, following are the pre-defined:
	//
	Type string `xml:"http://www.onvif.org/ver10/device/wsdl Type,omitempty"`

	Pos Vector `xml:"http://www.onvif.org/ver10/schema Pos,omitempty"`

	Extension OSDPosConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// OSDPosConfigurationExtension type
type OSDPosConfigurationExtension struct {
}

// OSDColor type
type OSDColor struct {
	Color Color `xml:"http://www.onvif.org/ver10/schema Color,omitempty"`

	Transparent int32 `xml:"http://www.onvif.org/ver10/schema Transparent,attr,omitempty"`
}

// OSDTextConfiguration type
type OSDTextConfiguration struct {

	//
	// The following OSD Text Type are defined:
	//
	Type string `xml:"http://www.onvif.org/ver10/device/wsdl Type,omitempty"`

	//
	// List of supported OSD date formats. This element shall be present when the value of Type field has Date or DateAndTime. The following DateFormat are defined:
	//
	DateFormat string `xml:"http://www.onvif.org/ver10/device/wsdl DateFormat,omitempty"`

	//
	// List of supported OSD time formats. This element shall be present when the value of Type field has Time or DateAndTime. The following TimeFormat are defined:
	//
	TimeFormat string `xml:"http://www.onvif.org/ver10/device/wsdl TimeFormat,omitempty"`

	// Font size of the text in pt.
	FontSize int32 `xml:"http://www.onvif.org/ver10/schema FontSize,omitempty"`

	// Font color of the text.
	FontColor OSDColor `xml:"http://www.onvif.org/ver10/schema FontColor,omitempty"`

	// Background color of the text.
	BackgroundColor OSDColor `xml:"http://www.onvif.org/ver10/schema BackgroundColor,omitempty"`

	// The content of text to be displayed.
	PlainText string `xml:"http://www.onvif.org/ver10/device/wsdl PlainText,omitempty"`

	Extension OSDTextConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// This flag is applicable for Type Plain and defaults to true. When set to false the PlainText content will not be persistent across device reboots.

	IsPersistentText bool `xml:"http://www.onvif.org/ver10/device/wsdl IsPersistentText,attr,omitempty"`
}

// OSDTextConfigurationExtension type
type OSDTextConfigurationExtension struct {
}

// OSDImgConfiguration type
type OSDImgConfiguration struct {

	// The URI of the image which to be displayed.
	ImgPath AnyURI `xml:"http://www.onvif.org/ver10/schema ImgPath,omitempty"`

	Extension OSDImgConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// OSDImgConfigurationExtension type
type OSDImgConfigurationExtension struct {
}

// ColorspaceRange type
type ColorspaceRange struct {
	X FloatRange `xml:"http://www.onvif.org/ver10/schema X,omitempty"`

	Y FloatRange `xml:"http://www.onvif.org/ver10/schema Y,omitempty"`

	Z FloatRange `xml:"http://www.onvif.org/ver10/schema Z,omitempty"`

	//
	// Acceptable values are the same as in tt:Color.
	//
	Colorspace AnyURI `xml:"http://www.onvif.org/ver10/schema Colorspace,omitempty"`
}

// ColorOptions type
type ColorOptions struct {

	// List the supported color.
	ColorList []Color `xml:"http://www.onvif.org/ver10/schema ColorList,omitempty"`

	// Define the range of color supported.
	ColorspaceRange []ColorspaceRange `xml:"http://www.onvif.org/ver10/schema ColorspaceRange,omitempty"`
}

// OSDColorOptions type
type OSDColorOptions struct {

	// Optional list of supported colors.
	Color ColorOptions `xml:"http://www.onvif.org/ver10/schema Color,omitempty"`

	// Range of the transparent level. Larger means more tranparent.
	Transparent IntRange `xml:"http://www.onvif.org/ver10/schema Transparent,omitempty"`

	Extension OSDColorOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// OSDColorOptionsExtension type
type OSDColorOptionsExtension struct {
}

// OSDTextOptions type
type OSDTextOptions struct {

	// List of supported OSD text type. When a device indicates the supported number relating to Text type in MaximumNumberOfOSDs, the type shall be presented.
	Type []string `xml:"http://www.onvif.org/ver10/device/wsdl Type,omitempty"`

	// Range of the font size value.
	FontSizeRange IntRange `xml:"http://www.onvif.org/ver10/schema FontSizeRange,omitempty"`

	// List of supported date format.
	DateFormat []string `xml:"http://www.onvif.org/ver10/device/wsdl DateFormat,omitempty"`

	// List of supported time format.
	TimeFormat []string `xml:"http://www.onvif.org/ver10/device/wsdl TimeFormat,omitempty"`

	// List of supported font color.
	FontColor OSDColorOptions `xml:"http://www.onvif.org/ver10/schema FontColor,omitempty"`

	// List of supported background color.
	BackgroundColor OSDColorOptions `xml:"http://www.onvif.org/ver10/schema BackgroundColor,omitempty"`

	Extension OSDTextOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// OSDTextOptionsExtension type
type OSDTextOptionsExtension struct {
}

// OSDImgOptions type
type OSDImgOptions struct {

	// List of available image URIs.
	ImagePath []AnyURI `xml:"http://www.onvif.org/ver10/schema ImagePath,omitempty"`

	Extension OSDImgOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// List of supported image MIME types, such as "image/png".

	FormatsSupported StringAttrList `xml:"http://www.onvif.org/ver10/device/wsdl FormatsSupported,attr,omitempty"`

	// The maximum size (in bytes) of the image that can be uploaded.

	MaxSize int32 `xml:"http://www.onvif.org/ver10/schema MaxSize,attr,omitempty"`

	// The maximum width (in pixels) of the image that can be uploaded.

	MaxWidth int32 `xml:"http://www.onvif.org/ver10/schema MaxWidth,attr,omitempty"`

	// The maximum height (in pixels) of the image that can be uploaded.

	MaxHeight int32 `xml:"http://www.onvif.org/ver10/schema MaxHeight,attr,omitempty"`
}

// OSDImgOptionsExtension type
type OSDImgOptionsExtension struct {
}

// OSDConfiguration type
type OSDConfiguration struct {
	*DeviceEntity

	// Reference to the video source configuration.
	VideoSourceConfigurationToken OSDReference `xml:"http://www.onvif.org/ver10/schema VideoSourceConfigurationToken,omitempty"`

	// Type of OSD.
	Type OSDType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// Position configuration of OSD.
	Position OSDPosConfiguration `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`

	// Text configuration of OSD. It shall be present when the value of Type field is Text.
	TextString OSDTextConfiguration `xml:"http://www.onvif.org/ver10/schema TextString,omitempty"`

	// Image configuration of OSD. It shall be present when the value of Type field is Image
	Image OSDImgConfiguration `xml:"http://www.onvif.org/ver10/schema Image,omitempty"`

	Extension OSDConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// OSDConfigurationExtension type
type OSDConfigurationExtension struct {
}

// MaximumNumberOfOSDs type
type MaximumNumberOfOSDs struct {
	Total int32 `xml:"http://www.onvif.org/ver10/schema Total,attr,omitempty"`

	Image int32 `xml:"http://www.onvif.org/ver10/schema Image,attr,omitempty"`

	PlainText int32 `xml:"http://www.onvif.org/ver10/schema PlainText,attr,omitempty"`

	Date int32 `xml:"http://www.onvif.org/ver10/schema Date,attr,omitempty"`

	Time int32 `xml:"http://www.onvif.org/ver10/schema Time,attr,omitempty"`

	DateAndTime int32 `xml:"http://www.onvif.org/ver10/schema DateAndTime,attr,omitempty"`
}

// OSDConfigurationOptions type
type OSDConfigurationOptions struct {

	// The maximum number of OSD configurations supported for the specified video source configuration. If the configuration does not support OSDs, this value shall be zero and the Type and PositionOption elements are ignored. If a device limits the number of instances by OSDType, it shall indicate the supported number for each type via the related attribute.
	MaximumNumberOfOSDs MaximumNumberOfOSDs `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfOSDs,omitempty"`

	// List supported type of OSD configuration. When a device indicates the supported number for each types in MaximumNumberOfOSDs, related type shall be presented. A device shall return Option element relating to listed type.
	Type []OSDType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// List available OSD position type. Following are the pre-defined:
	//
	PositionOption []string `xml:"http://www.onvif.org/ver10/device/wsdl PositionOption,omitempty"`

	// Option of the OSD text configuration. This element shall be returned if the device is signaling the support for Text.
	TextOption OSDTextOptions `xml:"http://www.onvif.org/ver10/schema TextOption,omitempty"`

	// Option of the OSD image configuration. This element shall be returned if the device is signaling the support for Image.
	ImageOption OSDImgOptions `xml:"http://www.onvif.org/ver10/schema ImageOption,omitempty"`

	Extension OSDConfigurationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// OSDConfigurationOptionsExtension type
type OSDConfigurationOptionsExtension struct {
}

// FileProgress type
type FileProgress struct {

	// Exported file name
	FileName string `xml:"http://www.onvif.org/ver10/device/wsdl FileName,omitempty"`

	// Normalized percentage completion for uploading the exported file
	Progress float32 `xml:"http://www.onvif.org/ver10/schema Progress,omitempty"`
}

// ArrayOfFileProgress type
type ArrayOfFileProgress struct {

	// Exported file name and export progress information
	FileProgress []FileProgress `xml:"http://www.onvif.org/ver10/schema FileProgress,omitempty"`

	Extension ArrayOfFileProgressExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// ArrayOfFileProgressExtension type
type ArrayOfFileProgressExtension struct {
}

// StorageReferencePath type
type StorageReferencePath struct {

	// identifier of an existing Storage Configuration.
	StorageToken ReferenceToken `xml:"http://www.onvif.org/ver10/device/wsdl StorageToken,omitempty"`

	// gives the relative directory path on the storage
	RelativePath string `xml:"http://www.onvif.org/ver10/device/wsdl RelativePath,omitempty"`

	Extension StorageReferencePathExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// StorageReferencePathExtension type
type StorageReferencePathExtension struct {
}

// Device type
type Device interface {

	/* Returns information about services on the device. */
	GetServices(request *GetServices) (*GetServicesResponse, error)

	GetServicesContext(ctx context.Context, request *GetServices) (*GetServicesResponse, error)

	/* Returns the capabilities of the device service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/* This operation gets basic device information from the device. */
	GetDeviceInformation(request *GetDeviceInformation) (*GetDeviceInformationResponse, error)

	GetDeviceInformationContext(ctx context.Context, request *GetDeviceInformation) (*GetDeviceInformationResponse, error)

	/* This operation sets the device system date and time. The device shall support the
	configuration of the daylight saving setting and of the manual system date and time (if
	applicable) or indication of NTP time (if applicable) through the SetSystemDateAndTime
	command.
	If system time and date are set manually, the client shall include UTCDateTime in the request.
	A TimeZone token which is not formed according to the rules of IEEE 1003.1 section 8.3 is considered as invalid timezone.
	The DayLightSavings flag should be set to true to activate any DST settings of the TimeZone string.
	Clear the DayLightSavings flag if the DST portion of the TimeZone settings should be ignored.
	*/
	SetSystemDateAndTime(request *SetSystemDateAndTime) (*SetSystemDateAndTimeResponse, error)

	SetSystemDateAndTimeContext(ctx context.Context, request *SetSystemDateAndTime) (*SetSystemDateAndTimeResponse, error)

	/* This operation gets the device system date and time. The device shall support the return of
	the daylight saving setting and of the manual system date and time (if applicable) or indication
	of NTP time (if applicable) through the GetSystemDateAndTime command.
	A device shall provide the UTCDateTime information. */
	GetSystemDateAndTime(request *GetSystemDateAndTime) (*GetSystemDateAndTimeResponse, error)

	GetSystemDateAndTimeContext(ctx context.Context, request *GetSystemDateAndTime) (*GetSystemDateAndTimeResponse, error)

	/* This operation reloads the parameters on the device to their factory default values. */
	SetSystemFactoryDefault(request *SetSystemFactoryDefault) (*SetSystemFactoryDefaultResponse, error)

	SetSystemFactoryDefaultContext(ctx context.Context, request *SetSystemFactoryDefault) (*SetSystemFactoryDefaultResponse, error)

	/* This operation upgrades a device firmware version. After a successful upgrade the response
	message is sent before the device reboots. The device should support firmware upgrade
	through the UpgradeSystemFirmware command. The exact format of the firmware data is
	outside the scope of this standard. */
	UpgradeSystemFirmware(request *UpgradeSystemFirmware) (*UpgradeSystemFirmwareResponse, error)

	UpgradeSystemFirmwareContext(ctx context.Context, request *UpgradeSystemFirmware) (*UpgradeSystemFirmwareResponse, error)

	/* This operation reboots the device. */
	SystemReboot(request *SystemReboot) (*SystemRebootResponse, error)

	SystemRebootContext(ctx context.Context, request *SystemReboot) (*SystemRebootResponse, error)

	/* This operation restores the system backup configuration files(s) previously retrieved from a
	device. The device should support restore of backup configuration file(s) through the
	RestoreSystem command. The exact format of the backup configuration file(s) is outside the
	scope of this standard. If the command is supported, it shall accept backup files returned by
	the GetSystemBackup command. */
	RestoreSystem(request *RestoreSystem) (*RestoreSystemResponse, error)

	RestoreSystemContext(ctx context.Context, request *RestoreSystem) (*RestoreSystemResponse, error)

	/* This operation is retrieves system backup configuration file(s) from a device. The device
	should support return of back up configuration file(s) through the GetSystemBackup command.
	The backup is returned with reference to a name and mime-type together with binary data.
	The exact format of the backup configuration files is outside the scope of this standard. */
	GetSystemBackup(request *GetSystemBackup) (*GetSystemBackupResponse, error)

	GetSystemBackupContext(ctx context.Context, request *GetSystemBackup) (*GetSystemBackupResponse, error)

	/* This operation gets a system log from the device. The exact format of the system logs is outside the scope of this standard. */
	GetSystemLog(request *GetSystemLog) (*GetSystemLogResponse, error)

	GetSystemLogContext(ctx context.Context, request *GetSystemLog) (*GetSystemLogResponse, error)

	/* This operation gets arbitary device diagnostics information from the device. */
	GetSystemSupportInformation(request *GetSystemSupportInformation) (*GetSystemSupportInformationResponse, error)

	GetSystemSupportInformationContext(ctx context.Context, request *GetSystemSupportInformation) (*GetSystemSupportInformationResponse, error)

	/* This operation requests the scope parameters of a device. The scope parameters are used in
	the device discovery to match a probe message, see Section 7. The Scope parameters are of
	two different types:
	Fixed scope parameters are permanent device characteristics and cannot be removed through the device management interface.
	The scope type is indicated in the scope list returned in the get scope parameters response. A device shall support
	retrieval of discovery scope parameters through the GetScopes command. As some scope parameters are mandatory,
	the device shall return a non-empty scope list in the response. */
	GetScopes(request *GetScopes) (*GetScopesResponse, error)

	GetScopesContext(ctx context.Context, request *GetScopes) (*GetScopesResponse, error)

	/* This operation sets the scope parameters of a device. The scope parameters are used in the
	device discovery to match a probe message.
	This operation replaces all existing configurable scope parameters (not fixed parameters). If
	this shall be avoided, one should use the scope add command instead. The device shall
	support configuration of discovery scope parameters through the SetScopes command. */
	SetScopes(request *SetScopes) (*SetScopesResponse, error)

	SetScopesContext(ctx context.Context, request *SetScopes) (*SetScopesResponse, error)

	/* This operation adds new configurable scope parameters to a device. The scope parameters
	are used in the device discovery to match a probe message. The device shall
	support addition of discovery scope parameters through the AddScopes command. */
	AddScopes(request *AddScopes) (*AddScopesResponse, error)

	AddScopesContext(ctx context.Context, request *AddScopes) (*AddScopesResponse, error)

	/* This operation deletes scope-configurable scope parameters from a device. The scope
	parameters are used in the device discovery to match a probe message, see Section 7. The
	device shall support deletion of discovery scope parameters through the RemoveScopes
	command.
	Table */
	RemoveScopes(request *RemoveScopes) (*RemoveScopesResponse, error)

	RemoveScopesContext(ctx context.Context, request *RemoveScopes) (*RemoveScopesResponse, error)

	/* This operation gets the discovery mode of a device. See Section 7.2 for the definition of the
	different device discovery modes. The device shall support retrieval of the discovery mode
	setting through the GetDiscoveryMode command. */
	GetDiscoveryMode(request *GetDiscoveryMode) (*GetDiscoveryModeResponse, error)

	GetDiscoveryModeContext(ctx context.Context, request *GetDiscoveryMode) (*GetDiscoveryModeResponse, error)

	/* This operation sets the discovery mode operation of a device. See Section 7.2 for the
	definition of the different device discovery modes. The device shall support configuration of
	the discovery mode setting through the SetDiscoveryMode command. */
	SetDiscoveryMode(request *SetDiscoveryMode) (*SetDiscoveryModeResponse, error)

	SetDiscoveryModeContext(ctx context.Context, request *SetDiscoveryMode) (*SetDiscoveryModeResponse, error)

	/* This operation gets the remote discovery mode of a device. See Section 7.4 for the definition
	of remote discovery extensions. A device that supports remote discovery shall support
	retrieval of the remote discovery mode setting through the GetRemoteDiscoveryMode
	command. */
	GetRemoteDiscoveryMode(request *GetRemoteDiscoveryMode) (*GetRemoteDiscoveryModeResponse, error)

	GetRemoteDiscoveryModeContext(ctx context.Context, request *GetRemoteDiscoveryMode) (*GetRemoteDiscoveryModeResponse, error)

	/* This operation sets the remote discovery mode of operation of a device. See Section 7.4 for
	the definition of remote discovery remote extensions. A device that supports remote discovery
	shall support configuration of the discovery mode setting through the
	SetRemoteDiscoveryMode command. */
	SetRemoteDiscoveryMode(request *SetRemoteDiscoveryMode) (*SetRemoteDiscoveryModeResponse, error)

	SetRemoteDiscoveryModeContext(ctx context.Context, request *SetRemoteDiscoveryMode) (*SetRemoteDiscoveryModeResponse, error)

	/* This operation gets the remote DP address or addresses from a device. If the device supports
	remote discovery, as specified in Section 7.4, the device shall support retrieval of the remote
	DP address(es) through the GetDPAddresses command. */
	GetDPAddresses(request *GetDPAddresses) (*GetDPAddressesResponse, error)

	GetDPAddressesContext(ctx context.Context, request *GetDPAddresses) (*GetDPAddressesResponse, error)

	/* This operation sets the remote DP address or addresses on a device. If the device supports
	remote discovery, as specified in Section 7.4, the device shall support configuration of the
	remote DP address(es) through the SetDPAddresses command. */
	SetDPAddresses(request *SetDPAddresses) (*SetDPAddressesResponse, error)

	SetDPAddressesContext(ctx context.Context, request *SetDPAddresses) (*SetDPAddressesResponse, error)

	/* A client can ask for the device service endpoint reference address property that can be used
	to derive the password equivalent for remote user operation. The device shall support the
	GetEndpointReference command returning the address property of the device service
	endpoint reference. */
	GetEndpointReference(request *GetEndpointReference) (*GetEndpointReferenceResponse, error)

	GetEndpointReferenceContext(ctx context.Context, request *GetEndpointReference) (*GetEndpointReferenceResponse, error)

	/* This operation returns the configured remote user (if any). A device supporting remote user
	handling shall support this operation. The user is only valid for the WS-UserToken profile or
	as a HTTP / RTSP user.
	The algorithm to use for deriving the password is described in section 5.12.2.1 of the core specification. */
	GetRemoteUser(request *GetRemoteUser) (*GetRemoteUserResponse, error)

	GetRemoteUserContext(ctx context.Context, request *GetRemoteUser) (*GetRemoteUserResponse, error)

	/* This operation sets the remote user. A device supporting remote user handling shall support this
	operation. The user is only valid for the WS-UserToken profile or as a HTTP / RTSP user.
	The password that is set shall always be the original (not derived) password.
	If UseDerivedPassword is set password derivation shall be done by the device when connecting to a
	remote device.The algorithm to use for deriving the password is described in section 5.12.2.1 of the core specification.
	To remove the remote user SetRemoteUser should be called without the RemoteUser parameter. */
	SetRemoteUser(request *SetRemoteUser) (*SetRemoteUserResponse, error)

	SetRemoteUserContext(ctx context.Context, request *SetRemoteUser) (*SetRemoteUserResponse, error)

	/* This operation lists the registered users and corresponding credentials on a device. The
	device shall support retrieval of registered device users and their credentials for the user
	token through the GetUsers command. */
	GetUsers(request *GetUsers) (*GetUsersResponse, error)

	GetUsersContext(ctx context.Context, request *GetUsers) (*GetUsersResponse, error)

	/* This operation creates new device users and corresponding credentials on a device for authentication purposes.
	The device shall support creation of device users and their credentials through the CreateUsers
	command. Either all users are created successfully or a fault message shall be returned
	without creating any user.
	ONVIF compliant devices are recommended to support password length of at least 28 bytes,
	as clients may follow the password derivation mechanism which results in 'password
	equivalent' of length 28 bytes, as described in section 3.1.2 of the ONVIF security white paper. */
	CreateUsers(request *CreateUsers) (*CreateUsersResponse, error)

	CreateUsersContext(ctx context.Context, request *CreateUsers) (*CreateUsersResponse, error)

	/* This operation deletes users on a device. The device shall support deletion of device users and their credentials
	through the DeleteUsers command. A device may have one or more fixed users
	that cannot be deleted to ensure access to the unit. Either all users are deleted successfully or a
	fault message shall be returned and no users be deleted. */
	DeleteUsers(request *DeleteUsers) (*DeleteUsersResponse, error)

	DeleteUsersContext(ctx context.Context, request *DeleteUsers) (*DeleteUsersResponse, error)

	/* This operation updates the settings for one or several users on a device for authentication purposes.
	The device shall support update of device users and their credentials through the SetUser command.
	Either all change requests are processed successfully or a fault message shall be returned and no change requests be processed. */
	SetUser(request *SetUser) (*SetUserResponse, error)

	SetUserContext(ctx context.Context, request *SetUser) (*SetUserResponse, error)

	/* It is possible for an endpoint to request a URL that can be used to retrieve the complete
	schema and WSDL definitions of a device. The command gives in return a URL entry point
	where all the necessary product specific WSDL and schema definitions can be retrieved. The
	device shall provide a URL for WSDL and schema download through the GetWsdlUrl command. */
	GetWsdlUrl(request *GetWsdlUrl) (*GetWsdlUrlResponse, error)

	GetWsdlUrlContext(ctx context.Context, request *GetWsdlUrl) (*GetWsdlUrlResponse, error)

	/* This method has been replaced by the more generic GetServices method.
	For capabilities of individual services refer to the GetServiceCapabilities methods. */
	GetCapabilities(request *GetCapabilities) (*GetCapabilitiesResponse, error)

	GetCapabilitiesContext(ctx context.Context, request *GetCapabilities) (*GetCapabilitiesResponse, error)

	/* This operation is used by an endpoint to get the hostname from a device. The device shall
	return its hostname configurations through the GetHostname command. */
	GetHostname(request *GetHostname) (*GetHostnameResponse, error)

	GetHostnameContext(ctx context.Context, request *GetHostname) (*GetHostnameResponse, error)

	/* This operation sets the hostname on a device. It shall be possible to set the device hostname
	configurations through the SetHostname command.
	A device shall accept string formated according to RFC 1123 section 2.1 or alternatively to RFC 952,
	other string shall be considered as invalid strings.
	*/
	SetHostname(request *SetHostname) (*SetHostnameResponse, error)

	SetHostnameContext(ctx context.Context, request *SetHostname) (*SetHostnameResponse, error)

	/* This operation controls whether the hostname is set manually or retrieved via DHCP. */
	SetHostnameFromDHCP(request *SetHostnameFromDHCP) (*SetHostnameFromDHCPResponse, error)

	SetHostnameFromDHCPContext(ctx context.Context, request *SetHostnameFromDHCP) (*SetHostnameFromDHCPResponse, error)

	/* This operation gets the DNS settings from a device. The device shall return its DNS
	configurations through the GetDNS command. */
	GetDNS(request *GetDNS) (*GetDNSResponse, error)

	GetDNSContext(ctx context.Context, request *GetDNS) (*GetDNSResponse, error)

	/* This operation sets the DNS settings on a device. It shall be possible to set the device DNS
	configurations through the SetDNS command. */
	SetDNS(request *SetDNS) (*SetDNSResponse, error)

	SetDNSContext(ctx context.Context, request *SetDNS) (*SetDNSResponse, error)

	/* This operation gets the NTP settings from a device. If the device supports NTP, it shall be
	possible to get the NTP server settings through the GetNTP command. */
	GetNTP(request *GetNTP) (*GetNTPResponse, error)

	GetNTPContext(ctx context.Context, request *GetNTP) (*GetNTPResponse, error)

	/* This operation sets the NTP settings on a device. If the device supports NTP, it shall be
	possible to set the NTP server settings through the SetNTP command.
	A device shall accept string formated according to RFC 1123 section 2.1 or alternatively to RFC 952,
	other string shall be considered as invalid strings.
	Changes to the NTP server list will not affect the clock mode DateTimeType. Use SetSystemDateAndTime to activate NTP operation.
	*/
	SetNTP(request *SetNTP) (*SetNTPResponse, error)

	SetNTPContext(ctx context.Context, request *SetNTP) (*SetNTPResponse, error)

	/* This operation gets the dynamic DNS settings from a device. If the device supports dynamic
	DNS as specified in [RFC 2136] and [RFC 4702], it shall be possible to get the type, name
	and TTL through the GetDynamicDNS command. */
	GetDynamicDNS(request *GetDynamicDNS) (*GetDynamicDNSResponse, error)

	GetDynamicDNSContext(ctx context.Context, request *GetDynamicDNS) (*GetDynamicDNSResponse, error)

	/* This operation sets the dynamic DNS settings on a device. If the device supports dynamic
	DNS as specified in [RFC 2136] and [RFC 4702], it shall be possible to set the type, name
	and TTL through the SetDynamicDNS command. */
	SetDynamicDNS(request *SetDynamicDNS) (*SetDynamicDNSResponse, error)

	SetDynamicDNSContext(ctx context.Context, request *SetDynamicDNS) (*SetDynamicDNSResponse, error)

	/* This operation gets the network interface configuration from a device. The device shall
	support return of network interface configuration settings as defined by the NetworkInterface
	type through the GetNetworkInterfaces command. */
	GetNetworkInterfaces(request *GetNetworkInterfaces) (*GetNetworkInterfacesResponse, error)

	GetNetworkInterfacesContext(ctx context.Context, request *GetNetworkInterfaces) (*GetNetworkInterfacesResponse, error)

	/* This operation sets the network interface configuration on a device. The device shall support
	network configuration of supported network interfaces through the SetNetworkInterfaces
	command.
	For interoperability with a client unaware of the IEEE 802.11 extension a device shall retain
	its IEEE 802.11 configuration if the IEEE 802.11 configuration element isn’t present in the
	request. */
	SetNetworkInterfaces(request *SetNetworkInterfaces) (*SetNetworkInterfacesResponse, error)

	SetNetworkInterfacesContext(ctx context.Context, request *SetNetworkInterfaces) (*SetNetworkInterfacesResponse, error)

	/* This operation gets defined network protocols from a device. The device shall support the
	GetNetworkProtocols command returning configured network protocols. */
	GetNetworkProtocols(request *GetNetworkProtocols) (*GetNetworkProtocolsResponse, error)

	GetNetworkProtocolsContext(ctx context.Context, request *GetNetworkProtocols) (*GetNetworkProtocolsResponse, error)

	/* This operation configures defined network protocols on a device. The device shall support
	configuration of defined network protocols through the SetNetworkProtocols command. */
	SetNetworkProtocols(request *SetNetworkProtocols) (*SetNetworkProtocolsResponse, error)

	SetNetworkProtocolsContext(ctx context.Context, request *SetNetworkProtocols) (*SetNetworkProtocolsResponse, error)

	/* This operation gets the default gateway settings from a device. The device shall support the
	GetNetworkDefaultGateway command returning configured default gateway address(es). */
	GetNetworkDefaultGateway(request *GetNetworkDefaultGateway) (*GetNetworkDefaultGatewayResponse, error)

	GetNetworkDefaultGatewayContext(ctx context.Context, request *GetNetworkDefaultGateway) (*GetNetworkDefaultGatewayResponse, error)

	/* This operation sets the default gateway settings on a device. The device shall support
	configuration of default gateway through the SetNetworkDefaultGateway command. */
	SetNetworkDefaultGateway(request *SetNetworkDefaultGateway) (*SetNetworkDefaultGatewayResponse, error)

	SetNetworkDefaultGatewayContext(ctx context.Context, request *SetNetworkDefaultGateway) (*SetNetworkDefaultGatewayResponse, error)

	/* This operation gets the zero-configuration from a device. If the device supports dynamic IP
		configuration according to [RFC3927], it shall support the return of IPv4 zero configuration
		address and status through the GetZeroConfiguration command.
	Devices supporting zero configuration on more than one interface shall use the extension to list the additional interface settings. */
	GetZeroConfiguration(request *GetZeroConfiguration) (*GetZeroConfigurationResponse, error)

	GetZeroConfigurationContext(ctx context.Context, request *GetZeroConfiguration) (*GetZeroConfigurationResponse, error)

	/* This operation sets the zero-configuration. Use GetCapalities to get if zero-zero-configuration is supported or not. */
	SetZeroConfiguration(request *SetZeroConfiguration) (*SetZeroConfigurationResponse, error)

	SetZeroConfigurationContext(ctx context.Context, request *SetZeroConfiguration) (*SetZeroConfigurationResponse, error)

	/* This operation gets the IP address filter settings from a device. If the device supports device
	access control based on IP filtering rules (denied or accepted ranges of IP addresses), the
	device shall support the GetIPAddressFilter command. */
	GetIPAddressFilter(request *GetIPAddressFilter) (*GetIPAddressFilterResponse, error)

	GetIPAddressFilterContext(ctx context.Context, request *GetIPAddressFilter) (*GetIPAddressFilterResponse, error)

	/* This operation sets the IP address filter settings on a device. If the device supports device
	access control based on IP filtering rules (denied or accepted ranges of IP addresses), the
	device shall support configuration of IP filtering rules through the SetIPAddressFilter
	command. */
	SetIPAddressFilter(request *SetIPAddressFilter) (*SetIPAddressFilterResponse, error)

	SetIPAddressFilterContext(ctx context.Context, request *SetIPAddressFilter) (*SetIPAddressFilterResponse, error)

	/* This operation adds an IP filter address to a device. If the device supports device access
	control based on IP filtering rules (denied or accepted ranges of IP addresses), the device
	shall support adding of IP filtering addresses through the AddIPAddressFilter command. */
	AddIPAddressFilter(request *AddIPAddressFilter) (*AddIPAddressFilterResponse, error)

	AddIPAddressFilterContext(ctx context.Context, request *AddIPAddressFilter) (*AddIPAddressFilterResponse, error)

	/* This operation deletes an IP filter address from a device. If the device supports device access
	control based on IP filtering rules (denied or accepted ranges of IP addresses), the device
	shall support deletion of IP filtering addresses through the RemoveIPAddressFilter command. */
	RemoveIPAddressFilter(request *RemoveIPAddressFilter) (*RemoveIPAddressFilterResponse, error)

	RemoveIPAddressFilterContext(ctx context.Context, request *RemoveIPAddressFilter) (*RemoveIPAddressFilterResponse, error)

	/* Access to different services and sub-sets of services should be subject to access control. The
	WS-Security framework gives the prerequisite for end-point authentication. Authorization
	decisions can then be taken using an access security policy. This standard does not mandate
	any particular policy description format or security policy but this is up to the device
	manufacturer or system provider to choose policy and policy description format of choice.
	However, an access policy (in arbitrary format) can be requested using this command. If the
	device supports access policy settings based on WS-Security authentication, then the device
	shall support this command. */
	GetAccessPolicy(request *GetAccessPolicy) (*GetAccessPolicyResponse, error)

	GetAccessPolicyContext(ctx context.Context, request *GetAccessPolicy) (*GetAccessPolicyResponse, error)

	/* This command sets the device access security policy (for more details on the access security
	policy see the Get command). If the device supports access policy settings
	based on WS-Security authentication, then the device shall support this command. */
	SetAccessPolicy(request *SetAccessPolicy) (*SetAccessPolicyResponse, error)

	SetAccessPolicyContext(ctx context.Context, request *SetAccessPolicy) (*SetAccessPolicyResponse, error)

	/* This operation generates a private/public key pair and also can create a self-signed device
	certificate as a result of key pair generation. The certificate is created using a suitable
	onboard key pair generation mechanism.
	If a device supports onboard key pair generation, the device that supports TLS shall support
	this certificate creation command. And also, if a device supports onboard key pair generation,
	the device that support IEEE 802.1X shall support this command for the purpose of key pair
	generation. Certificates and key pairs are identified using certificate IDs. These IDs are either
	chosen by the certificate generation requester or by the device (in case that no ID value is
	given). */
	CreateCertificate(request *CreateCertificate) (*CreateCertificateResponse, error)

	CreateCertificateContext(ctx context.Context, request *CreateCertificate) (*CreateCertificateResponse, error)

	/* This operation gets all device server certificates (including self-signed) for the purpose of TLS
	authentication and all device client certificates for the purpose of IEEE 802.1X authentication.
	This command lists only the TLS server certificates and IEEE 802.1X client certificates for the
	device (neither trusted CA certificates nor trusted root certificates). The certificates are
	returned as binary data. A device that supports TLS shall support this command and the
	certificates shall be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding
	rules. */
	GetCertificates(request *GetCertificates) (*GetCertificatesResponse, error)

	GetCertificatesContext(ctx context.Context, request *GetCertificates) (*GetCertificatesResponse, error)

	/* This operation is specific to TLS functionality. This operation gets the status
	(enabled/disabled) of the device TLS server certificates. A device that supports TLS shall
	support this command. */
	GetCertificatesStatus(request *GetCertificatesStatus) (*GetCertificatesStatusResponse, error)

	GetCertificatesStatusContext(ctx context.Context, request *GetCertificatesStatus) (*GetCertificatesStatusResponse, error)

	/* This operation is specific to TLS functionality. This operation sets the status (enable/disable)
	of the device TLS server certificates. A device that supports TLS shall support this command.
	Typically only one device server certificate is allowed to be enabled at a time. */
	SetCertificatesStatus(request *SetCertificatesStatus) (*SetCertificatesStatusResponse, error)

	SetCertificatesStatusContext(ctx context.Context, request *SetCertificatesStatus) (*SetCertificatesStatusResponse, error)

	/* This operation deletes a certificate or multiple certificates. The device MAY also delete a
	private/public key pair which is coupled with the certificate to be deleted. The device that
	support either TLS or IEEE 802.1X shall support the deletion of a certificate or multiple
	certificates through this command. Either all certificates are deleted successfully or a fault
	message shall be returned without deleting any certificate. */
	DeleteCertificates(request *DeleteCertificates) (*DeleteCertificatesResponse, error)

	DeleteCertificatesContext(ctx context.Context, request *DeleteCertificates) (*DeleteCertificatesResponse, error)

	/* This operation requests a PKCS #10 certificate signature request from the device. The
	returned information field shall be either formatted exactly as specified in [PKCS#10] or PEM
	encoded [PKCS#10] format. In order for this command to work, the device must already have
	a private/public key pair. This key pair should be referred by CertificateID as specified in the
	input parameter description. This CertificateID refers to the key pair generated using
	CreateCertificate command.
	A device that support onboard key pair generation that supports either TLS or IEEE 802.1X
	using client certificate shall support this command. */
	GetPkcs10Request(request *GetPkcs10Request) (*GetPkcs10RequestResponse, error)

	GetPkcs10RequestContext(ctx context.Context, request *GetPkcs10Request) (*GetPkcs10RequestResponse, error)

	/* TLS server certificate(s) or IEEE 802.1X client certificate(s) created using the PKCS#10
	certificate request command can be loaded into the device using this command (see Section
	8.4.13). The certificate ID in the request shall be present. The device may sort the received
	certificate(s) based on the public key and subject information in the certificate(s).
	The certificate ID in the request will be the ID value the client wish to have. The device is
	supposed to scan the generated key pairs present in the device to identify which is the
	correspondent key pair with the loaded certificate and then make the link between the
	certificate and the key pair.
	A device that supports onboard key pair generation that support either TLS or IEEE 802.1X
	shall support this command.
	The certificates shall be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding
	rules.
	This command is applicable to any device type, although the parameter name is called for
	historical reasons NVTCertificate. */
	LoadCertificates(request *LoadCertificates) (*LoadCertificatesResponse, error)

	LoadCertificatesContext(ctx context.Context, request *LoadCertificates) (*LoadCertificatesResponse, error)

	/* This operation is specific to TLS functionality. This operation gets the status
	(enabled/disabled) of the device TLS client authentication. A device that supports TLS shall
	support this command. */
	GetClientCertificateMode(request *GetClientCertificateMode) (*GetClientCertificateModeResponse, error)

	GetClientCertificateModeContext(ctx context.Context, request *GetClientCertificateMode) (*GetClientCertificateModeResponse, error)

	/* This operation is specific to TLS functionality. This operation sets the status
	(enabled/disabled) of the device TLS client authentication. A device that supports TLS shall
	support this command. */
	SetClientCertificateMode(request *SetClientCertificateMode) (*SetClientCertificateModeResponse, error)

	SetClientCertificateModeContext(ctx context.Context, request *SetClientCertificateMode) (*SetClientCertificateModeResponse, error)

	/* This operation gets a list of all available relay outputs and their settings.
	This method has been depricated with version 2.0. Refer to the DeviceIO service. */
	GetRelayOutputs(request *GetRelayOutputs) (*GetRelayOutputsResponse, error)

	GetRelayOutputsContext(ctx context.Context, request *GetRelayOutputs) (*GetRelayOutputsResponse, error)

	/* This operation sets the settings of a relay output.
	This method has been depricated with version 2.0. Refer to the DeviceIO service. */
	SetRelayOutputSettings(request *SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, error)

	SetRelayOutputSettingsContext(ctx context.Context, request *SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, error)

	/* This operation sets the state of a relay output.
	This method has been depricated with version 2.0. Refer to the DeviceIO service. */
	SetRelayOutputState(request *SetRelayOutputState) (*SetRelayOutputStateResponse, error)

	SetRelayOutputStateContext(ctx context.Context, request *SetRelayOutputState) (*SetRelayOutputStateResponse, error)

	/* Manage auxiliary commands supported by a device, such as controlling an Infrared (IR) lamp,
	a heater or a wiper or a thermometer that is connected to the device.
	The supported commands can be retrieved via the AuxiliaryCommands capability.
	Although the name of the auxiliary commands can be freely defined, commands starting with the prefix tt: are
	reserved to define frequently used commands and these reserved commands shall all share the "tt:command|parameter" syntax.

	A device that indicates auxiliary service capability shall support this command. */
	SendAuxiliaryCommand(request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error)

	SendAuxiliaryCommandContext(ctx context.Context, request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error)

	/* CA certificates will be loaded into a device and be used for the sake of following two cases.
	The one is for the purpose of TLS client authentication in TLS server function. The other one
	is for the purpose of Authentication Server authentication in IEEE 802.1X function. This
	operation gets all CA certificates loaded into a device. A device that supports either TLS client
	authentication or IEEE 802.1X shall support this command and the returned certificates shall
	be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding rules. */
	GetCACertificates(request *GetCACertificates) (*GetCACertificatesResponse, error)

	GetCACertificatesContext(ctx context.Context, request *GetCACertificates) (*GetCACertificatesResponse, error)

	/* There might be some cases that a Certificate Authority or some other equivalent creates a
	certificate without having PKCS#10 certificate signing request. In such cases, the certificate
	will be bundled in conjunction with its private key. This command will be used for such use
	case scenarios. The certificate ID in the request is optionally set to the ID value the client
	wish to have. If the certificate ID is not specified in the request, device can choose the ID
	accordingly.
	This operation imports a private/public key pair into the device.
	The certificates shall be encoded using ASN.1 [X.681], [X.682], [X.683] DER [X.690] encoding
	rules.
	A device that does not support onboard key pair generation and support either TLS or IEEE
	802.1X using client certificate shall support this command. A device that support onboard key
	pair generation MAY support this command. The security policy of a device that supports this
	operation should make sure that the private key is sufficiently protected. */
	LoadCertificateWithPrivateKey(request *LoadCertificateWithPrivateKey) (*LoadCertificateWithPrivateKeyResponse, error)

	LoadCertificateWithPrivateKeyContext(ctx context.Context, request *LoadCertificateWithPrivateKey) (*LoadCertificateWithPrivateKeyResponse, error)

	/* This operation requests the information of a certificate specified by certificate ID. The device
	should respond with its “Issuer DN”, “Subject DN”, “Key usage”, "Extended key usage”, “Key
	Length”, “Version”, “Serial Number”, “Signature Algorithm” and “Validity” data as the
	information of the certificate, as long as the device can retrieve such information from the
	specified certificate.
	A device that supports either TLS or IEEE 802.1X should support this command. */
	GetCertificateInformation(request *GetCertificateInformation) (*GetCertificateInformationResponse, error)

	GetCertificateInformationContext(ctx context.Context, request *GetCertificateInformation) (*GetCertificateInformationResponse, error)

	/* This command is used when it is necessary to load trusted CA certificates or trusted root
	certificates for the purpose of verification for its counterpart i.e. client certificate verification in
	TLS function or server certificate verification in IEEE 802.1X function.
	A device that support either TLS or IEEE 802.1X shall support this command. As for the
	supported certificate format, either DER format or PEM format is possible to be used. But a
	device that support this command shall support at least DER format as supported format type.
	The device may sort the received certificate(s) based on the public key and subject
	information in the certificate(s). Either all CA certificates are loaded successfully or a fault
	message shall be returned without loading any CA certificate. */
	LoadCACertificates(request *LoadCACertificates) (*LoadCACertificatesResponse, error)

	LoadCACertificatesContext(ctx context.Context, request *LoadCACertificates) (*LoadCACertificatesResponse, error)

	/* This operation newly creates IEEE 802.1X configuration parameter set of the device. The
	device shall support this command if it supports IEEE 802.1X. If the device receives this
	request with already existing configuration token (Dot1XConfigurationToken) specification, the
	device should respond with 'ter:ReferenceToken ' error to indicate there is some configuration
	conflict. */
	CreateDot1XConfiguration(request *CreateDot1XConfiguration) (*CreateDot1XConfigurationResponse, error)

	CreateDot1XConfigurationContext(ctx context.Context, request *CreateDot1XConfiguration) (*CreateDot1XConfigurationResponse, error)

	/* While the CreateDot1XConfiguration command is trying to create a new configuration
	parameter set, this operation modifies existing IEEE 802.1X configuration parameter set of
	the device. A device that support IEEE 802.1X shall support this command. */
	SetDot1XConfiguration(request *SetDot1XConfiguration) (*SetDot1XConfigurationResponse, error)

	SetDot1XConfigurationContext(ctx context.Context, request *SetDot1XConfiguration) (*SetDot1XConfigurationResponse, error)

	/* This operation gets one IEEE 802.1X configuration parameter set from the device by
	specifying the configuration token (Dot1XConfigurationToken).
	A device that supports IEEE 802.1X shall support this command.
	Regardless of whether the 802.1X method in the retrieved configuration has a password or
	not, the device shall not include the Password element in the response. */
	GetDot1XConfiguration(request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error)

	GetDot1XConfigurationContext(ctx context.Context, request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error)

	/* This operation gets all the existing IEEE 802.1X configuration parameter sets from the device.
	The device shall respond with all the IEEE 802.1X configurations so that the client can get to
	know how many IEEE 802.1X configurations are existing and how they are configured.
	A device that support IEEE 802.1X shall support this command.
	Regardless of whether the 802.1X method in the retrieved configuration has a password or
	not, the device shall not include the Password element in the response. */
	GetDot1XConfigurations(request *GetDot1XConfigurations) (*GetDot1XConfigurationsResponse, error)

	GetDot1XConfigurationsContext(ctx context.Context, request *GetDot1XConfigurations) (*GetDot1XConfigurationsResponse, error)

	/* This operation deletes an IEEE 802.1X configuration parameter set from the device. Which
	configuration should be deleted is specified by the 'Dot1XConfigurationToken' in the request.
	A device that support IEEE 802.1X shall support this command. */
	DeleteDot1XConfiguration(request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error)

	DeleteDot1XConfigurationContext(ctx context.Context, request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error)

	/* This operation returns the IEEE802.11 capabilities. The device shall support
	this operation. */
	GetDot11Capabilities(request *GetDot11Capabilities) (*GetDot11CapabilitiesResponse, error)

	GetDot11CapabilitiesContext(ctx context.Context, request *GetDot11Capabilities) (*GetDot11CapabilitiesResponse, error)

	/* This operation returns the status of a wireless network interface. The device shall support this
	command. */
	GetDot11Status(request *GetDot11Status) (*GetDot11StatusResponse, error)

	GetDot11StatusContext(ctx context.Context, request *GetDot11Status) (*GetDot11StatusResponse, error)

	/* This operation returns a lists of the wireless networks in range of the device. A device should
	support this operation. */
	ScanAvailableDot11Networks(request *ScanAvailableDot11Networks) (*ScanAvailableDot11NetworksResponse, error)

	ScanAvailableDot11NetworksContext(ctx context.Context, request *ScanAvailableDot11Networks) (*ScanAvailableDot11NetworksResponse, error)

	/* This operation is used to retrieve URIs from which system information may be downloaded
	using HTTP. URIs may be returned for the following system information:
	System Logs. Multiple system logs may be returned, of different types. The exact format of
	the system logs is outside the scope of this specification.
	Support Information. This consists of arbitrary device diagnostics information from a device.
	The exact format of the diagnostic information is outside the scope of this specification.
	System Backup. The received file is a backup file that can be used to restore the current
	device configuration at a later date. The exact format of the backup configuration file is
	outside the scope of this specification.
	If the device allows retrieval of system logs, support information or system backup data, it
	should make them available via HTTP GET. If it does, it shall support the GetSystemUris
	command. */
	GetSystemUris(request *GetSystemUris) (*GetSystemUrisResponse, error)

	GetSystemUrisContext(ctx context.Context, request *GetSystemUris) (*GetSystemUrisResponse, error)

	/* This operation initiates a firmware upgrade using the HTTP POST mechanism. The response
	to the command includes an HTTP URL to which the upgrade file may be uploaded. The
	actual upgrade takes place as soon as the HTTP POST operation has completed. The device
	should support firmware upgrade through the StartFirmwareUpgrade command. The exact
	format of the firmware data is outside the scope of this specification.
	Firmware upgrade over HTTP may be achieved using the following steps:
	If the firmware upgrade fails because the upgrade file was invalid, the HTTP POST response
	shall be “415 Unsupported Media Type”. If the firmware upgrade fails due to an error at the
	device, the HTTP POST response shall be “500 Internal Server Error”.
	The value of the Content-Type header in the HTTP POST request shall be “application/octetstream”. */
	StartFirmwareUpgrade(request *StartFirmwareUpgrade) (*StartFirmwareUpgradeResponse, error)

	StartFirmwareUpgradeContext(ctx context.Context, request *StartFirmwareUpgrade) (*StartFirmwareUpgradeResponse, error)

	/* This operation initiates a system restore from backed up configuration data using the HTTP
	POST mechanism. The response to the command includes an HTTP URL to which the backup
	file may be uploaded. The actual restore takes place as soon as the HTTP POST operation
	has completed. Devices should support system restore through the StartSystemRestore
	command. The exact format of the backup configuration data is outside the scope of this
	specification.
	System restore over HTTP may be achieved using the following steps:
	If the system restore fails because the uploaded file was invalid, the HTTP POST response
	shall be “415 Unsupported Media Type”. If the system restore fails due to an error at the
	device, the HTTP POST response shall be “500 Internal Server Error”.
	The value of the Content-Type header in the HTTP POST request shall be “application/octetstream”. */
	StartSystemRestore(request *StartSystemRestore) (*StartSystemRestoreResponse, error)

	StartSystemRestoreContext(ctx context.Context, request *StartSystemRestore) (*StartSystemRestoreResponse, error)

	/*
		This operation lists all existing storage configurations for the device.
	*/
	GetStorageConfigurations(request *GetStorageConfigurations) (*GetStorageConfigurationsResponse, error)

	GetStorageConfigurationsContext(ctx context.Context, request *GetStorageConfigurations) (*GetStorageConfigurationsResponse, error)

	/*
		This operation creates a new storage configuration.
		The configuration data shall be created in the device and shall be persistent (remain after reboot).
	*/
	CreateStorageConfiguration(request *CreateStorageConfiguration) (*CreateStorageConfigurationResponse, error)

	CreateStorageConfigurationContext(ctx context.Context, request *CreateStorageConfiguration) (*CreateStorageConfigurationResponse, error)

	/*
		This operation retrieves the Storage configuration associated with the given storage configuration token.
	*/
	GetStorageConfiguration(request *GetStorageConfiguration) (*GetStorageConfigurationResponse, error)

	GetStorageConfigurationContext(ctx context.Context, request *GetStorageConfiguration) (*GetStorageConfigurationResponse, error)

	/*
		This operation modifies an existing Storage configuration.
	*/
	SetStorageConfiguration(request *SetStorageConfiguration) (*SetStorageConfigurationResponse, error)

	SetStorageConfigurationContext(ctx context.Context, request *SetStorageConfiguration) (*SetStorageConfigurationResponse, error)

	/*
		This operation deletes the given storage configuration and configuration change shall always be persistent.
	*/
	DeleteStorageConfiguration(request *DeleteStorageConfiguration) (*DeleteStorageConfigurationResponse, error)

	DeleteStorageConfigurationContext(ctx context.Context, request *DeleteStorageConfiguration) (*DeleteStorageConfigurationResponse, error)

	/*
		This operation lists all existing geo location configurations for the device.
	*/
	GetGeoLocation(request *GetGeoLocation) (*GetGeoLocationResponse, error)

	GetGeoLocationContext(ctx context.Context, request *GetGeoLocation) (*GetGeoLocationResponse, error)

	/*
		This operation allows to modify one or more geo configuration entries.
	*/
	SetGeoLocation(request *SetGeoLocation) (*SetGeoLocationResponse, error)

	SetGeoLocationContext(ctx context.Context, request *SetGeoLocation) (*SetGeoLocationResponse, error)

	/*
		This operation deletes the given geo location entries.
	*/
	DeleteGeoLocation(request *DeleteGeoLocation) (*DeleteGeoLocationResponse, error)

	DeleteGeoLocationContext(ctx context.Context, request *DeleteGeoLocation) (*DeleteGeoLocationResponse, error)
}

// device type
type device struct {
	client *soap.Client
	xaddr  string
}

func NewDevice(client *soap.Client, xaddr string) Device {
	return &device{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *device) GetServicesContext(ctx context.Context, request *GetServices) (*GetServicesResponse, error) {
	response := new(GetServicesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetServices", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetServices(request *GetServices) (*GetServicesResponse, error) {
	return service.GetServicesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDeviceInformationContext(ctx context.Context, request *GetDeviceInformation) (*GetDeviceInformationResponse, error) {
	response := new(GetDeviceInformationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDeviceInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDeviceInformation(request *GetDeviceInformation) (*GetDeviceInformationResponse, error) {
	return service.GetDeviceInformationContext(
		context.Background(),
		request,
	)
}

func (service *device) SetSystemDateAndTimeContext(ctx context.Context, request *SetSystemDateAndTime) (*SetSystemDateAndTimeResponse, error) {
	response := new(SetSystemDateAndTimeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetSystemDateAndTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetSystemDateAndTime(request *SetSystemDateAndTime) (*SetSystemDateAndTimeResponse, error) {
	return service.SetSystemDateAndTimeContext(
		context.Background(),
		request,
	)
}

func (service *device) GetSystemDateAndTimeContext(ctx context.Context, request *GetSystemDateAndTime) (*GetSystemDateAndTimeResponse, error) {
	response := new(GetSystemDateAndTimeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetSystemDateAndTime", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetSystemDateAndTime(request *GetSystemDateAndTime) (*GetSystemDateAndTimeResponse, error) {
	return service.GetSystemDateAndTimeContext(
		context.Background(),
		request,
	)
}

func (service *device) SetSystemFactoryDefaultContext(ctx context.Context, request *SetSystemFactoryDefault) (*SetSystemFactoryDefaultResponse, error) {
	response := new(SetSystemFactoryDefaultResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetSystemFactoryDefault", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetSystemFactoryDefault(request *SetSystemFactoryDefault) (*SetSystemFactoryDefaultResponse, error) {
	return service.SetSystemFactoryDefaultContext(
		context.Background(),
		request,
	)
}

func (service *device) UpgradeSystemFirmwareContext(ctx context.Context, request *UpgradeSystemFirmware) (*UpgradeSystemFirmwareResponse, error) {
	response := new(UpgradeSystemFirmwareResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/UpgradeSystemFirmware", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) UpgradeSystemFirmware(request *UpgradeSystemFirmware) (*UpgradeSystemFirmwareResponse, error) {
	return service.UpgradeSystemFirmwareContext(
		context.Background(),
		request,
	)
}

func (service *device) SystemRebootContext(ctx context.Context, request *SystemReboot) (*SystemRebootResponse, error) {
	response := new(SystemRebootResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SystemReboot", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SystemReboot(request *SystemReboot) (*SystemRebootResponse, error) {
	return service.SystemRebootContext(
		context.Background(),
		request,
	)
}

func (service *device) RestoreSystemContext(ctx context.Context, request *RestoreSystem) (*RestoreSystemResponse, error) {
	response := new(RestoreSystemResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/RestoreSystem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) RestoreSystem(request *RestoreSystem) (*RestoreSystemResponse, error) {
	return service.RestoreSystemContext(
		context.Background(),
		request,
	)
}

func (service *device) GetSystemBackupContext(ctx context.Context, request *GetSystemBackup) (*GetSystemBackupResponse, error) {
	response := new(GetSystemBackupResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetSystemBackup", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetSystemBackup(request *GetSystemBackup) (*GetSystemBackupResponse, error) {
	return service.GetSystemBackupContext(
		context.Background(),
		request,
	)
}

func (service *device) GetSystemLogContext(ctx context.Context, request *GetSystemLog) (*GetSystemLogResponse, error) {
	response := new(GetSystemLogResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetSystemLog", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetSystemLog(request *GetSystemLog) (*GetSystemLogResponse, error) {
	return service.GetSystemLogContext(
		context.Background(),
		request,
	)
}

func (service *device) GetSystemSupportInformationContext(ctx context.Context, request *GetSystemSupportInformation) (*GetSystemSupportInformationResponse, error) {
	response := new(GetSystemSupportInformationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetSystemSupportInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetSystemSupportInformation(request *GetSystemSupportInformation) (*GetSystemSupportInformationResponse, error) {
	return service.GetSystemSupportInformationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetScopesContext(ctx context.Context, request *GetScopes) (*GetScopesResponse, error) {
	response := new(GetScopesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetScopes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetScopes(request *GetScopes) (*GetScopesResponse, error) {
	return service.GetScopesContext(
		context.Background(),
		request,
	)
}

func (service *device) SetScopesContext(ctx context.Context, request *SetScopes) (*SetScopesResponse, error) {
	response := new(SetScopesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetScopes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetScopes(request *SetScopes) (*SetScopesResponse, error) {
	return service.SetScopesContext(
		context.Background(),
		request,
	)
}

func (service *device) AddScopesContext(ctx context.Context, request *AddScopes) (*AddScopesResponse, error) {
	response := new(AddScopesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/AddScopes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) AddScopes(request *AddScopes) (*AddScopesResponse, error) {
	return service.AddScopesContext(
		context.Background(),
		request,
	)
}

func (service *device) RemoveScopesContext(ctx context.Context, request *RemoveScopes) (*RemoveScopesResponse, error) {
	response := new(RemoveScopesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/RemoveScopes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) RemoveScopes(request *RemoveScopes) (*RemoveScopesResponse, error) {
	return service.RemoveScopesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDiscoveryModeContext(ctx context.Context, request *GetDiscoveryMode) (*GetDiscoveryModeResponse, error) {
	response := new(GetDiscoveryModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDiscoveryMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDiscoveryMode(request *GetDiscoveryMode) (*GetDiscoveryModeResponse, error) {
	return service.GetDiscoveryModeContext(
		context.Background(),
		request,
	)
}

func (service *device) SetDiscoveryModeContext(ctx context.Context, request *SetDiscoveryMode) (*SetDiscoveryModeResponse, error) {
	response := new(SetDiscoveryModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetDiscoveryMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetDiscoveryMode(request *SetDiscoveryMode) (*SetDiscoveryModeResponse, error) {
	return service.SetDiscoveryModeContext(
		context.Background(),
		request,
	)
}

func (service *device) GetRemoteDiscoveryModeContext(ctx context.Context, request *GetRemoteDiscoveryMode) (*GetRemoteDiscoveryModeResponse, error) {
	response := new(GetRemoteDiscoveryModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetRemoteDiscoveryMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetRemoteDiscoveryMode(request *GetRemoteDiscoveryMode) (*GetRemoteDiscoveryModeResponse, error) {
	return service.GetRemoteDiscoveryModeContext(
		context.Background(),
		request,
	)
}

func (service *device) SetRemoteDiscoveryModeContext(ctx context.Context, request *SetRemoteDiscoveryMode) (*SetRemoteDiscoveryModeResponse, error) {
	response := new(SetRemoteDiscoveryModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetRemoteDiscoveryMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetRemoteDiscoveryMode(request *SetRemoteDiscoveryMode) (*SetRemoteDiscoveryModeResponse, error) {
	return service.SetRemoteDiscoveryModeContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDPAddressesContext(ctx context.Context, request *GetDPAddresses) (*GetDPAddressesResponse, error) {
	response := new(GetDPAddressesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDPAddresses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDPAddresses(request *GetDPAddresses) (*GetDPAddressesResponse, error) {
	return service.GetDPAddressesContext(
		context.Background(),
		request,
	)
}

func (service *device) SetDPAddressesContext(ctx context.Context, request *SetDPAddresses) (*SetDPAddressesResponse, error) {
	response := new(SetDPAddressesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetDPAddresses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetDPAddresses(request *SetDPAddresses) (*SetDPAddressesResponse, error) {
	return service.SetDPAddressesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetEndpointReferenceContext(ctx context.Context, request *GetEndpointReference) (*GetEndpointReferenceResponse, error) {
	response := new(GetEndpointReferenceResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetEndpointReference", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetEndpointReference(request *GetEndpointReference) (*GetEndpointReferenceResponse, error) {
	return service.GetEndpointReferenceContext(
		context.Background(),
		request,
	)
}

func (service *device) GetRemoteUserContext(ctx context.Context, request *GetRemoteUser) (*GetRemoteUserResponse, error) {
	response := new(GetRemoteUserResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetRemoteUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetRemoteUser(request *GetRemoteUser) (*GetRemoteUserResponse, error) {
	return service.GetRemoteUserContext(
		context.Background(),
		request,
	)
}

func (service *device) SetRemoteUserContext(ctx context.Context, request *SetRemoteUser) (*SetRemoteUserResponse, error) {
	response := new(SetRemoteUserResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetRemoteUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetRemoteUser(request *SetRemoteUser) (*SetRemoteUserResponse, error) {
	return service.SetRemoteUserContext(
		context.Background(),
		request,
	)
}

func (service *device) GetUsersContext(ctx context.Context, request *GetUsers) (*GetUsersResponse, error) {
	response := new(GetUsersResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetUsers(request *GetUsers) (*GetUsersResponse, error) {
	return service.GetUsersContext(
		context.Background(),
		request,
	)
}

func (service *device) CreateUsersContext(ctx context.Context, request *CreateUsers) (*CreateUsersResponse, error) {
	response := new(CreateUsersResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/CreateUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) CreateUsers(request *CreateUsers) (*CreateUsersResponse, error) {
	return service.CreateUsersContext(
		context.Background(),
		request,
	)
}

func (service *device) DeleteUsersContext(ctx context.Context, request *DeleteUsers) (*DeleteUsersResponse, error) {
	response := new(DeleteUsersResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/DeleteUsers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) DeleteUsers(request *DeleteUsers) (*DeleteUsersResponse, error) {
	return service.DeleteUsersContext(
		context.Background(),
		request,
	)
}

func (service *device) SetUserContext(ctx context.Context, request *SetUser) (*SetUserResponse, error) {
	response := new(SetUserResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetUser(request *SetUser) (*SetUserResponse, error) {
	return service.SetUserContext(
		context.Background(),
		request,
	)
}

func (service *device) GetWsdlUrlContext(ctx context.Context, request *GetWsdlUrl) (*GetWsdlUrlResponse, error) {
	response := new(GetWsdlUrlResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetWsdlUrl", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetWsdlUrl(request *GetWsdlUrl) (*GetWsdlUrlResponse, error) {
	return service.GetWsdlUrlContext(
		context.Background(),
		request,
	)
}

func (service *device) GetCapabilitiesContext(ctx context.Context, request *GetCapabilities) (*GetCapabilitiesResponse, error) {
	response := new(GetCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetCapabilities(request *GetCapabilities) (*GetCapabilitiesResponse, error) {
	return service.GetCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetHostnameContext(ctx context.Context, request *GetHostname) (*GetHostnameResponse, error) {
	response := new(GetHostnameResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetHostname", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetHostname(request *GetHostname) (*GetHostnameResponse, error) {
	return service.GetHostnameContext(
		context.Background(),
		request,
	)
}

func (service *device) SetHostnameContext(ctx context.Context, request *SetHostname) (*SetHostnameResponse, error) {
	response := new(SetHostnameResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetHostname", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetHostname(request *SetHostname) (*SetHostnameResponse, error) {
	return service.SetHostnameContext(
		context.Background(),
		request,
	)
}

func (service *device) SetHostnameFromDHCPContext(ctx context.Context, request *SetHostnameFromDHCP) (*SetHostnameFromDHCPResponse, error) {
	response := new(SetHostnameFromDHCPResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetHostnameFromDHCP", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetHostnameFromDHCP(request *SetHostnameFromDHCP) (*SetHostnameFromDHCPResponse, error) {
	return service.SetHostnameFromDHCPContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDNSContext(ctx context.Context, request *GetDNS) (*GetDNSResponse, error) {
	response := new(GetDNSResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDNS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDNS(request *GetDNS) (*GetDNSResponse, error) {
	return service.GetDNSContext(
		context.Background(),
		request,
	)
}

func (service *device) SetDNSContext(ctx context.Context, request *SetDNS) (*SetDNSResponse, error) {
	response := new(SetDNSResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetDNS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetDNS(request *SetDNS) (*SetDNSResponse, error) {
	return service.SetDNSContext(
		context.Background(),
		request,
	)
}

func (service *device) GetNTPContext(ctx context.Context, request *GetNTP) (*GetNTPResponse, error) {
	response := new(GetNTPResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetNTP", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetNTP(request *GetNTP) (*GetNTPResponse, error) {
	return service.GetNTPContext(
		context.Background(),
		request,
	)
}

func (service *device) SetNTPContext(ctx context.Context, request *SetNTP) (*SetNTPResponse, error) {
	response := new(SetNTPResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetNTP", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetNTP(request *SetNTP) (*SetNTPResponse, error) {
	return service.SetNTPContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDynamicDNSContext(ctx context.Context, request *GetDynamicDNS) (*GetDynamicDNSResponse, error) {
	response := new(GetDynamicDNSResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDynamicDNS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDynamicDNS(request *GetDynamicDNS) (*GetDynamicDNSResponse, error) {
	return service.GetDynamicDNSContext(
		context.Background(),
		request,
	)
}

func (service *device) SetDynamicDNSContext(ctx context.Context, request *SetDynamicDNS) (*SetDynamicDNSResponse, error) {
	response := new(SetDynamicDNSResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetDynamicDNS", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetDynamicDNS(request *SetDynamicDNS) (*SetDynamicDNSResponse, error) {
	return service.SetDynamicDNSContext(
		context.Background(),
		request,
	)
}

func (service *device) GetNetworkInterfacesContext(ctx context.Context, request *GetNetworkInterfaces) (*GetNetworkInterfacesResponse, error) {
	response := new(GetNetworkInterfacesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetNetworkInterfaces", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetNetworkInterfaces(request *GetNetworkInterfaces) (*GetNetworkInterfacesResponse, error) {
	return service.GetNetworkInterfacesContext(
		context.Background(),
		request,
	)
}

func (service *device) SetNetworkInterfacesContext(ctx context.Context, request *SetNetworkInterfaces) (*SetNetworkInterfacesResponse, error) {
	response := new(SetNetworkInterfacesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetNetworkInterfaces", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetNetworkInterfaces(request *SetNetworkInterfaces) (*SetNetworkInterfacesResponse, error) {
	return service.SetNetworkInterfacesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetNetworkProtocolsContext(ctx context.Context, request *GetNetworkProtocols) (*GetNetworkProtocolsResponse, error) {
	response := new(GetNetworkProtocolsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetNetworkProtocols", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetNetworkProtocols(request *GetNetworkProtocols) (*GetNetworkProtocolsResponse, error) {
	return service.GetNetworkProtocolsContext(
		context.Background(),
		request,
	)
}

func (service *device) SetNetworkProtocolsContext(ctx context.Context, request *SetNetworkProtocols) (*SetNetworkProtocolsResponse, error) {
	response := new(SetNetworkProtocolsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetNetworkProtocols", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetNetworkProtocols(request *SetNetworkProtocols) (*SetNetworkProtocolsResponse, error) {
	return service.SetNetworkProtocolsContext(
		context.Background(),
		request,
	)
}

func (service *device) GetNetworkDefaultGatewayContext(ctx context.Context, request *GetNetworkDefaultGateway) (*GetNetworkDefaultGatewayResponse, error) {
	response := new(GetNetworkDefaultGatewayResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetNetworkDefaultGateway", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetNetworkDefaultGateway(request *GetNetworkDefaultGateway) (*GetNetworkDefaultGatewayResponse, error) {
	return service.GetNetworkDefaultGatewayContext(
		context.Background(),
		request,
	)
}

func (service *device) SetNetworkDefaultGatewayContext(ctx context.Context, request *SetNetworkDefaultGateway) (*SetNetworkDefaultGatewayResponse, error) {
	response := new(SetNetworkDefaultGatewayResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetNetworkDefaultGateway", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetNetworkDefaultGateway(request *SetNetworkDefaultGateway) (*SetNetworkDefaultGatewayResponse, error) {
	return service.SetNetworkDefaultGatewayContext(
		context.Background(),
		request,
	)
}

func (service *device) GetZeroConfigurationContext(ctx context.Context, request *GetZeroConfiguration) (*GetZeroConfigurationResponse, error) {
	response := new(GetZeroConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetZeroConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetZeroConfiguration(request *GetZeroConfiguration) (*GetZeroConfigurationResponse, error) {
	return service.GetZeroConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) SetZeroConfigurationContext(ctx context.Context, request *SetZeroConfiguration) (*SetZeroConfigurationResponse, error) {
	response := new(SetZeroConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetZeroConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetZeroConfiguration(request *SetZeroConfiguration) (*SetZeroConfigurationResponse, error) {
	return service.SetZeroConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetIPAddressFilterContext(ctx context.Context, request *GetIPAddressFilter) (*GetIPAddressFilterResponse, error) {
	response := new(GetIPAddressFilterResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetIPAddressFilter", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetIPAddressFilter(request *GetIPAddressFilter) (*GetIPAddressFilterResponse, error) {
	return service.GetIPAddressFilterContext(
		context.Background(),
		request,
	)
}

func (service *device) SetIPAddressFilterContext(ctx context.Context, request *SetIPAddressFilter) (*SetIPAddressFilterResponse, error) {
	response := new(SetIPAddressFilterResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetIPAddressFilter", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetIPAddressFilter(request *SetIPAddressFilter) (*SetIPAddressFilterResponse, error) {
	return service.SetIPAddressFilterContext(
		context.Background(),
		request,
	)
}

func (service *device) AddIPAddressFilterContext(ctx context.Context, request *AddIPAddressFilter) (*AddIPAddressFilterResponse, error) {
	response := new(AddIPAddressFilterResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/AddIPAddressFilter", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) AddIPAddressFilter(request *AddIPAddressFilter) (*AddIPAddressFilterResponse, error) {
	return service.AddIPAddressFilterContext(
		context.Background(),
		request,
	)
}

func (service *device) RemoveIPAddressFilterContext(ctx context.Context, request *RemoveIPAddressFilter) (*RemoveIPAddressFilterResponse, error) {
	response := new(RemoveIPAddressFilterResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/RemoveIPAddressFilter", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) RemoveIPAddressFilter(request *RemoveIPAddressFilter) (*RemoveIPAddressFilterResponse, error) {
	return service.RemoveIPAddressFilterContext(
		context.Background(),
		request,
	)
}

func (service *device) GetAccessPolicyContext(ctx context.Context, request *GetAccessPolicy) (*GetAccessPolicyResponse, error) {
	response := new(GetAccessPolicyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetAccessPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetAccessPolicy(request *GetAccessPolicy) (*GetAccessPolicyResponse, error) {
	return service.GetAccessPolicyContext(
		context.Background(),
		request,
	)
}

func (service *device) SetAccessPolicyContext(ctx context.Context, request *SetAccessPolicy) (*SetAccessPolicyResponse, error) {
	response := new(SetAccessPolicyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetAccessPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetAccessPolicy(request *SetAccessPolicy) (*SetAccessPolicyResponse, error) {
	return service.SetAccessPolicyContext(
		context.Background(),
		request,
	)
}

func (service *device) CreateCertificateContext(ctx context.Context, request *CreateCertificate) (*CreateCertificateResponse, error) {
	response := new(CreateCertificateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/CreateCertificate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) CreateCertificate(request *CreateCertificate) (*CreateCertificateResponse, error) {
	return service.CreateCertificateContext(
		context.Background(),
		request,
	)
}

func (service *device) GetCertificatesContext(ctx context.Context, request *GetCertificates) (*GetCertificatesResponse, error) {
	response := new(GetCertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetCertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetCertificates(request *GetCertificates) (*GetCertificatesResponse, error) {
	return service.GetCertificatesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetCertificatesStatusContext(ctx context.Context, request *GetCertificatesStatus) (*GetCertificatesStatusResponse, error) {
	response := new(GetCertificatesStatusResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetCertificatesStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetCertificatesStatus(request *GetCertificatesStatus) (*GetCertificatesStatusResponse, error) {
	return service.GetCertificatesStatusContext(
		context.Background(),
		request,
	)
}

func (service *device) SetCertificatesStatusContext(ctx context.Context, request *SetCertificatesStatus) (*SetCertificatesStatusResponse, error) {
	response := new(SetCertificatesStatusResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetCertificatesStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetCertificatesStatus(request *SetCertificatesStatus) (*SetCertificatesStatusResponse, error) {
	return service.SetCertificatesStatusContext(
		context.Background(),
		request,
	)
}

func (service *device) DeleteCertificatesContext(ctx context.Context, request *DeleteCertificates) (*DeleteCertificatesResponse, error) {
	response := new(DeleteCertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/DeleteCertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) DeleteCertificates(request *DeleteCertificates) (*DeleteCertificatesResponse, error) {
	return service.DeleteCertificatesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetPkcs10RequestContext(ctx context.Context, request *GetPkcs10Request) (*GetPkcs10RequestResponse, error) {
	response := new(GetPkcs10RequestResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetPkcs10Request", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetPkcs10Request(request *GetPkcs10Request) (*GetPkcs10RequestResponse, error) {
	return service.GetPkcs10RequestContext(
		context.Background(),
		request,
	)
}

func (service *device) LoadCertificatesContext(ctx context.Context, request *LoadCertificates) (*LoadCertificatesResponse, error) {
	response := new(LoadCertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/LoadCertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) LoadCertificates(request *LoadCertificates) (*LoadCertificatesResponse, error) {
	return service.LoadCertificatesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetClientCertificateModeContext(ctx context.Context, request *GetClientCertificateMode) (*GetClientCertificateModeResponse, error) {
	response := new(GetClientCertificateModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetClientCertificateMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetClientCertificateMode(request *GetClientCertificateMode) (*GetClientCertificateModeResponse, error) {
	return service.GetClientCertificateModeContext(
		context.Background(),
		request,
	)
}

func (service *device) SetClientCertificateModeContext(ctx context.Context, request *SetClientCertificateMode) (*SetClientCertificateModeResponse, error) {
	response := new(SetClientCertificateModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetClientCertificateMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetClientCertificateMode(request *SetClientCertificateMode) (*SetClientCertificateModeResponse, error) {
	return service.SetClientCertificateModeContext(
		context.Background(),
		request,
	)
}

func (service *device) GetRelayOutputsContext(ctx context.Context, request *GetRelayOutputs) (*GetRelayOutputsResponse, error) {
	response := new(GetRelayOutputsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetRelayOutputs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetRelayOutputs(request *GetRelayOutputs) (*GetRelayOutputsResponse, error) {
	return service.GetRelayOutputsContext(
		context.Background(),
		request,
	)
}

func (service *device) SetRelayOutputSettingsContext(ctx context.Context, request *SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, error) {
	response := new(SetRelayOutputSettingsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetRelayOutputSettings", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetRelayOutputSettings(request *SetRelayOutputSettings) (*SetRelayOutputSettingsResponse, error) {
	return service.SetRelayOutputSettingsContext(
		context.Background(),
		request,
	)
}

func (service *device) SetRelayOutputStateContext(ctx context.Context, request *SetRelayOutputState) (*SetRelayOutputStateResponse, error) {
	response := new(SetRelayOutputStateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetRelayOutputState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetRelayOutputState(request *SetRelayOutputState) (*SetRelayOutputStateResponse, error) {
	return service.SetRelayOutputStateContext(
		context.Background(),
		request,
	)
}

func (service *device) SendAuxiliaryCommandContext(ctx context.Context, request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error) {
	response := new(SendAuxiliaryCommandResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SendAuxiliaryCommand", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SendAuxiliaryCommand(request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error) {
	return service.SendAuxiliaryCommandContext(
		context.Background(),
		request,
	)
}

func (service *device) GetCACertificatesContext(ctx context.Context, request *GetCACertificates) (*GetCACertificatesResponse, error) {
	response := new(GetCACertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetCACertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetCACertificates(request *GetCACertificates) (*GetCACertificatesResponse, error) {
	return service.GetCACertificatesContext(
		context.Background(),
		request,
	)
}

func (service *device) LoadCertificateWithPrivateKeyContext(ctx context.Context, request *LoadCertificateWithPrivateKey) (*LoadCertificateWithPrivateKeyResponse, error) {
	response := new(LoadCertificateWithPrivateKeyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/LoadCertificateWithPrivateKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) LoadCertificateWithPrivateKey(request *LoadCertificateWithPrivateKey) (*LoadCertificateWithPrivateKeyResponse, error) {
	return service.LoadCertificateWithPrivateKeyContext(
		context.Background(),
		request,
	)
}

func (service *device) GetCertificateInformationContext(ctx context.Context, request *GetCertificateInformation) (*GetCertificateInformationResponse, error) {
	response := new(GetCertificateInformationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetCertificateInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetCertificateInformation(request *GetCertificateInformation) (*GetCertificateInformationResponse, error) {
	return service.GetCertificateInformationContext(
		context.Background(),
		request,
	)
}

func (service *device) LoadCACertificatesContext(ctx context.Context, request *LoadCACertificates) (*LoadCACertificatesResponse, error) {
	response := new(LoadCACertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/LoadCACertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) LoadCACertificates(request *LoadCACertificates) (*LoadCACertificatesResponse, error) {
	return service.LoadCACertificatesContext(
		context.Background(),
		request,
	)
}

func (service *device) CreateDot1XConfigurationContext(ctx context.Context, request *CreateDot1XConfiguration) (*CreateDot1XConfigurationResponse, error) {
	response := new(CreateDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/CreateDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) CreateDot1XConfiguration(request *CreateDot1XConfiguration) (*CreateDot1XConfigurationResponse, error) {
	return service.CreateDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) SetDot1XConfigurationContext(ctx context.Context, request *SetDot1XConfiguration) (*SetDot1XConfigurationResponse, error) {
	response := new(SetDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetDot1XConfiguration(request *SetDot1XConfiguration) (*SetDot1XConfigurationResponse, error) {
	return service.SetDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDot1XConfigurationContext(ctx context.Context, request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error) {
	response := new(GetDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDot1XConfiguration(request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error) {
	return service.GetDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDot1XConfigurationsContext(ctx context.Context, request *GetDot1XConfigurations) (*GetDot1XConfigurationsResponse, error) {
	response := new(GetDot1XConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDot1XConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDot1XConfigurations(request *GetDot1XConfigurations) (*GetDot1XConfigurationsResponse, error) {
	return service.GetDot1XConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *device) DeleteDot1XConfigurationContext(ctx context.Context, request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error) {
	response := new(DeleteDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/DeleteDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) DeleteDot1XConfiguration(request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error) {
	return service.DeleteDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDot11CapabilitiesContext(ctx context.Context, request *GetDot11Capabilities) (*GetDot11CapabilitiesResponse, error) {
	response := new(GetDot11CapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDot11Capabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDot11Capabilities(request *GetDot11Capabilities) (*GetDot11CapabilitiesResponse, error) {
	return service.GetDot11CapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *device) GetDot11StatusContext(ctx context.Context, request *GetDot11Status) (*GetDot11StatusResponse, error) {
	response := new(GetDot11StatusResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetDot11Status", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetDot11Status(request *GetDot11Status) (*GetDot11StatusResponse, error) {
	return service.GetDot11StatusContext(
		context.Background(),
		request,
	)
}

func (service *device) ScanAvailableDot11NetworksContext(ctx context.Context, request *ScanAvailableDot11Networks) (*ScanAvailableDot11NetworksResponse, error) {
	response := new(ScanAvailableDot11NetworksResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/ScanAvailableDot11Networks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) ScanAvailableDot11Networks(request *ScanAvailableDot11Networks) (*ScanAvailableDot11NetworksResponse, error) {
	return service.ScanAvailableDot11NetworksContext(
		context.Background(),
		request,
	)
}

func (service *device) GetSystemUrisContext(ctx context.Context, request *GetSystemUris) (*GetSystemUrisResponse, error) {
	response := new(GetSystemUrisResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetSystemUris", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetSystemUris(request *GetSystemUris) (*GetSystemUrisResponse, error) {
	return service.GetSystemUrisContext(
		context.Background(),
		request,
	)
}

func (service *device) StartFirmwareUpgradeContext(ctx context.Context, request *StartFirmwareUpgrade) (*StartFirmwareUpgradeResponse, error) {
	response := new(StartFirmwareUpgradeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/StartFirmwareUpgrade", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) StartFirmwareUpgrade(request *StartFirmwareUpgrade) (*StartFirmwareUpgradeResponse, error) {
	return service.StartFirmwareUpgradeContext(
		context.Background(),
		request,
	)
}

func (service *device) StartSystemRestoreContext(ctx context.Context, request *StartSystemRestore) (*StartSystemRestoreResponse, error) {
	response := new(StartSystemRestoreResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/StartSystemRestore", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) StartSystemRestore(request *StartSystemRestore) (*StartSystemRestoreResponse, error) {
	return service.StartSystemRestoreContext(
		context.Background(),
		request,
	)
}

func (service *device) GetStorageConfigurationsContext(ctx context.Context, request *GetStorageConfigurations) (*GetStorageConfigurationsResponse, error) {
	response := new(GetStorageConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetStorageConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetStorageConfigurations(request *GetStorageConfigurations) (*GetStorageConfigurationsResponse, error) {
	return service.GetStorageConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *device) CreateStorageConfigurationContext(ctx context.Context, request *CreateStorageConfiguration) (*CreateStorageConfigurationResponse, error) {
	response := new(CreateStorageConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/CreateStorageConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) CreateStorageConfiguration(request *CreateStorageConfiguration) (*CreateStorageConfigurationResponse, error) {
	return service.CreateStorageConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetStorageConfigurationContext(ctx context.Context, request *GetStorageConfiguration) (*GetStorageConfigurationResponse, error) {
	response := new(GetStorageConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetStorageConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetStorageConfiguration(request *GetStorageConfiguration) (*GetStorageConfigurationResponse, error) {
	return service.GetStorageConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) SetStorageConfigurationContext(ctx context.Context, request *SetStorageConfiguration) (*SetStorageConfigurationResponse, error) {
	response := new(SetStorageConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetStorageConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetStorageConfiguration(request *SetStorageConfiguration) (*SetStorageConfigurationResponse, error) {
	return service.SetStorageConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) DeleteStorageConfigurationContext(ctx context.Context, request *DeleteStorageConfiguration) (*DeleteStorageConfigurationResponse, error) {
	response := new(DeleteStorageConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/DeleteStorageConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) DeleteStorageConfiguration(request *DeleteStorageConfiguration) (*DeleteStorageConfigurationResponse, error) {
	return service.DeleteStorageConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *device) GetGeoLocationContext(ctx context.Context, request *GetGeoLocation) (*GetGeoLocationResponse, error) {
	response := new(GetGeoLocationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/GetGeoLocation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) GetGeoLocation(request *GetGeoLocation) (*GetGeoLocationResponse, error) {
	return service.GetGeoLocationContext(
		context.Background(),
		request,
	)
}

func (service *device) SetGeoLocationContext(ctx context.Context, request *SetGeoLocation) (*SetGeoLocationResponse, error) {
	response := new(SetGeoLocationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/SetGeoLocation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) SetGeoLocation(request *SetGeoLocation) (*SetGeoLocationResponse, error) {
	return service.SetGeoLocationContext(
		context.Background(),
		request,
	)
}

func (service *device) DeleteGeoLocationContext(ctx context.Context, request *DeleteGeoLocation) (*DeleteGeoLocationResponse, error) {
	response := new(DeleteGeoLocationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/device/wsdl/DeleteGeoLocation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *device) DeleteGeoLocation(request *DeleteGeoLocation) (*DeleteGeoLocationResponse, error) {
	return service.DeleteGeoLocationContext(
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

// PositiveInteger type
type PositiveInteger int64

// NonPositiveInteger type
type NonPositiveInteger int64

// AnySimpleType type
type AnySimpleType string

// String type
type String string
