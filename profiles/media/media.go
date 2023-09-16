package media

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// Indication which encodings are supported for this video source. The list may contain one or more enumeration values of tt:VideoEncoding.

// EncodingTypes type
type EncodingTypes []string

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capabilities for the media service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"Capabilities,omitempty"`
}

// GetVideoSources type
type GetVideoSources struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSources"`
}

// GetVideoSourcesResponse type
type GetVideoSourcesResponse struct {
	XMLName xml.Name `xml:"GetVideoSourcesResponse"`

	// List of existing Video Sources
	VideoSources []VideoSource `xml:"VideoSources,omitempty"`
}

// GetAudioSources type
type GetAudioSources struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSources"`
}

// GetAudioSourcesResponse type
type GetAudioSourcesResponse struct {
	XMLName xml.Name `xml:"GetAudioSourcesResponse"`

	// List of existing Audio Sources
	AudioSources []AudioSource `xml:"AudioSources,omitempty"`
}

// GetAudioOutputs type
type GetAudioOutputs struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputs"`
}

// GetAudioOutputsResponse type
type GetAudioOutputsResponse struct {
	XMLName xml.Name `xml:"GetAudioOutputsResponse"`

	// List of existing Audio Outputs
	AudioOutputs []AudioOutput `xml:"AudioOutputs,omitempty"`
}

// CreateProfile type
type CreateProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl CreateProfile"`

	// friendly name of the profile to be created
	Name Name `xml:"http://www.onvif.org/ver10/media/wsdl Name,omitempty"`

	// Optional token, specifying the unique identifier of the new profile. A device supports at least a token length of 12 characters and characters "A-Z" | "a-z" | "0-9" | "-.".
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Token,omitempty"`
}

// CreateProfileResponse type
type CreateProfileResponse struct {
	XMLName xml.Name `xml:"CreateProfileResponse"`

	// returns the new created profile
	Profile Profile `xml:"Profile,omitempty"`
}

// GetProfile type
type GetProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetProfile"`

	// this command requests a specific profile
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetProfileResponse type
type GetProfileResponse struct {
	XMLName xml.Name `xml:"GetProfileResponse"`

	// returns the requested media profile
	Profile Profile `xml:"Profile,omitempty"`
}

// GetProfiles type
type GetProfiles struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetProfiles"`
}

// GetProfilesResponse type
type GetProfilesResponse struct {
	XMLName xml.Name `xml:"GetProfilesResponse"`

	// lists all profiles that exist in the media service
	Profiles []Profile `xml:"Profiles,omitempty"`
}

// AddVideoEncoderConfiguration type
type AddVideoEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddVideoEncoderConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the VideoEncoderConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddVideoEncoderConfigurationResponse type
type AddVideoEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"AddVideoEncoderConfigurationResponse"`
}

// RemoveVideoEncoderConfiguration type
type RemoveVideoEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveVideoEncoderConfiguration"`

	// Contains a reference to the media profile from which the
	// VideoEncoderConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveVideoEncoderConfigurationResponse type
type RemoveVideoEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveVideoEncoderConfigurationResponse"`
}

// AddVideoSourceConfiguration type
type AddVideoSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddVideoSourceConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the VideoSourceConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddVideoSourceConfigurationResponse type
type AddVideoSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"AddVideoSourceConfigurationResponse"`
}

// RemoveVideoSourceConfiguration type
type RemoveVideoSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveVideoSourceConfiguration"`

	// Contains a reference to the media profile from which the
	// VideoSourceConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveVideoSourceConfigurationResponse type
type RemoveVideoSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveVideoSourceConfigurationResponse"`
}

// AddAudioEncoderConfiguration type
type AddAudioEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioEncoderConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the AudioEncoderConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddAudioEncoderConfigurationResponse type
type AddAudioEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"AddAudioEncoderConfigurationResponse"`
}

// RemoveAudioEncoderConfiguration type
type RemoveAudioEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioEncoderConfiguration"`

	// Contains a reference to the media profile from which the
	// AudioEncoderConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveAudioEncoderConfigurationResponse type
type RemoveAudioEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveAudioEncoderConfigurationResponse"`
}

// AddAudioSourceConfiguration type
type AddAudioSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioSourceConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the AudioSourceConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddAudioSourceConfigurationResponse type
type AddAudioSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"AddAudioSourceConfigurationResponse"`
}

// RemoveAudioSourceConfiguration type
type RemoveAudioSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioSourceConfiguration"`

	// Contains a reference to the media profile from which the
	// AudioSourceConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveAudioSourceConfigurationResponse type
type RemoveAudioSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveAudioSourceConfigurationResponse"`
}

// AddPTZConfiguration type
type AddPTZConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddPTZConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the PTZConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddPTZConfigurationResponse type
type AddPTZConfigurationResponse struct {
	XMLName xml.Name `xml:"AddPTZConfigurationResponse"`
}

// RemovePTZConfiguration type
type RemovePTZConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemovePTZConfiguration"`

	// Contains a reference to the media profile from which the
	// PTZConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemovePTZConfigurationResponse type
type RemovePTZConfigurationResponse struct {
	XMLName xml.Name `xml:"RemovePTZConfigurationResponse"`
}

// AddVideoAnalyticsConfiguration type
type AddVideoAnalyticsConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddVideoAnalyticsConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the VideoAnalyticsConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddVideoAnalyticsConfigurationResponse type
type AddVideoAnalyticsConfigurationResponse struct {
	XMLName xml.Name `xml:"AddVideoAnalyticsConfigurationResponse"`
}

// RemoveVideoAnalyticsConfiguration type
type RemoveVideoAnalyticsConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveVideoAnalyticsConfiguration"`

	// Contains a reference to the media profile from which the
	// VideoAnalyticsConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveVideoAnalyticsConfigurationResponse type
type RemoveVideoAnalyticsConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveVideoAnalyticsConfigurationResponse"`
}

// AddMetadataConfiguration type
type AddMetadataConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddMetadataConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the MetadataConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddMetadataConfigurationResponse type
type AddMetadataConfigurationResponse struct {
	XMLName xml.Name `xml:"AddMetadataConfigurationResponse"`
}

// RemoveMetadataConfiguration type
type RemoveMetadataConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveMetadataConfiguration"`

	// Contains a reference to the media profile from which the
	// MetadataConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveMetadataConfigurationResponse type
type RemoveMetadataConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveMetadataConfigurationResponse"`
}

// AddAudioOutputConfiguration type
type AddAudioOutputConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioOutputConfiguration"`

	// Reference to the profile where the configuration should be added
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// Contains a reference to the AudioOutputConfiguration to add
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddAudioOutputConfigurationResponse type
type AddAudioOutputConfigurationResponse struct {
	XMLName xml.Name `xml:"AddAudioOutputConfigurationResponse"`
}

// RemoveAudioOutputConfiguration type
type RemoveAudioOutputConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioOutputConfiguration"`

	// Contains a reference to the media profile from which the
	// AudioOutputConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveAudioOutputConfigurationResponse type
type RemoveAudioOutputConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveAudioOutputConfigurationResponse"`
}

// AddAudioDecoderConfiguration type
type AddAudioDecoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl AddAudioDecoderConfiguration"`

	// This element contains a reference to the profile where the configuration should be added.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`

	// This element contains a reference to the AudioDecoderConfiguration to add.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// AddAudioDecoderConfigurationResponse type
type AddAudioDecoderConfigurationResponse struct {
	XMLName xml.Name `xml:"AddAudioDecoderConfigurationResponse"`
}

// RemoveAudioDecoderConfiguration type
type RemoveAudioDecoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl RemoveAudioDecoderConfiguration"`

	// This element contains a  reference to the media profile from which the AudioDecoderConfiguration shall be removed.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// RemoveAudioDecoderConfigurationResponse type
type RemoveAudioDecoderConfigurationResponse struct {
	XMLName xml.Name `xml:"RemoveAudioDecoderConfigurationResponse"`
}

// DeleteProfile type
type DeleteProfile struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl DeleteProfile"`

	// This element contains a  reference to the profile that should be deleted.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// DeleteProfileResponse type
type DeleteProfileResponse struct {
	XMLName xml.Name `xml:"DeleteProfileResponse"`
}

// GetVideoEncoderConfigurations type
type GetVideoEncoderConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoEncoderConfigurations"`
}

// GetVideoEncoderConfigurationsResponse type
type GetVideoEncoderConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetVideoEncoderConfigurationsResponse"`

	// This element contains a list of video encoder configurations.
	Configurations []VideoEncoderConfiguration `xml:"Configurations,omitempty"`
}

// GetVideoSourceConfigurations type
type GetVideoSourceConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceConfigurations"`
}

// GetVideoSourceConfigurationsResponse type
type GetVideoSourceConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetVideoSourceConfigurationsResponse"`

	// This element contains a list of video source configurations.
	Configurations []VideoSourceConfiguration `xml:"Configurations,omitempty"`
}

// GetAudioEncoderConfigurations type
type GetAudioEncoderConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioEncoderConfigurations"`
}

// GetAudioEncoderConfigurationsResponse type
type GetAudioEncoderConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetAudioEncoderConfigurationsResponse"`

	// This element contains a list of audio encoder configurations.
	Configurations []AudioEncoderConfiguration `xml:"Configurations,omitempty"`
}

// GetAudioSourceConfigurations type
type GetAudioSourceConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSourceConfigurations"`
}

// GetAudioSourceConfigurationsResponse type
type GetAudioSourceConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetAudioSourceConfigurationsResponse"`

	// This element contains a list of audio source configurations.
	Configurations []AudioSourceConfiguration `xml:"Configurations,omitempty"`
}

// GetVideoAnalyticsConfigurations type
type GetVideoAnalyticsConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoAnalyticsConfigurations"`
}

// GetVideoAnalyticsConfigurationsResponse type
type GetVideoAnalyticsConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetVideoAnalyticsConfigurationsResponse"`

	// This element contains a list of VideoAnalytics configurations.
	Configurations []VideoAnalyticsConfiguration `xml:"Configurations,omitempty"`
}

// GetMetadataConfigurations type
type GetMetadataConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetMetadataConfigurations"`
}

// GetMetadataConfigurationsResponse type
type GetMetadataConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetMetadataConfigurationsResponse"`

	// This element contains a list of metadata configurations
	Configurations []MetadataConfiguration `xml:"Configurations,omitempty"`
}

// GetAudioOutputConfigurations type
type GetAudioOutputConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputConfigurations"`
}

// GetAudioOutputConfigurationsResponse type
type GetAudioOutputConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetAudioOutputConfigurationsResponse"`

	// This element contains a list of audio output configurations
	Configurations []AudioOutputConfiguration `xml:"Configurations,omitempty"`
}

// GetAudioDecoderConfigurations type
type GetAudioDecoderConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioDecoderConfigurations"`
}

// GetAudioDecoderConfigurationsResponse type
type GetAudioDecoderConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetAudioDecoderConfigurationsResponse"`

	// This element contains a list of audio decoder configurations
	Configurations []AudioDecoderConfiguration `xml:"Configurations,omitempty"`
}

// GetVideoSourceConfiguration type
type GetVideoSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceConfiguration"`

	// Token of the requested video source configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetVideoSourceConfigurationResponse type
type GetVideoSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"GetVideoSourceConfigurationResponse"`

	// The requested video source configuration.
	Configuration VideoSourceConfiguration `xml:"Configuration,omitempty"`
}

// GetVideoEncoderConfiguration type
type GetVideoEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoEncoderConfiguration"`

	// Token of the requested video encoder configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetVideoEncoderConfigurationResponse type
type GetVideoEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"GetVideoEncoderConfigurationResponse"`

	// The requested video encoder configuration.
	Configuration VideoEncoderConfiguration `xml:"Configuration,omitempty"`
}

// GetAudioSourceConfiguration type
type GetAudioSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSourceConfiguration"`

	// Token of the requested audio source configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetAudioSourceConfigurationResponse type
type GetAudioSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"GetAudioSourceConfigurationResponse"`

	// The requested audio source configuration.
	Configuration AudioSourceConfiguration `xml:"Configuration,omitempty"`
}

// GetAudioEncoderConfiguration type
type GetAudioEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioEncoderConfiguration"`

	// Token of the requested audio encoder configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetAudioEncoderConfigurationResponse type
type GetAudioEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"GetAudioEncoderConfigurationResponse"`

	// The requested audio encoder configuration
	Configuration AudioEncoderConfiguration `xml:"Configuration,omitempty"`
}

// GetVideoAnalyticsConfiguration type
type GetVideoAnalyticsConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoAnalyticsConfiguration"`

	// Token of the requested video analytics configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetVideoAnalyticsConfigurationResponse type
type GetVideoAnalyticsConfigurationResponse struct {
	XMLName xml.Name `xml:"GetVideoAnalyticsConfigurationResponse"`

	// The requested video analytics configuration.
	Configuration VideoAnalyticsConfiguration `xml:"Configuration,omitempty"`
}

// GetMetadataConfiguration type
type GetMetadataConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetMetadataConfiguration"`

	// Token of the requested metadata configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetMetadataConfigurationResponse type
type GetMetadataConfigurationResponse struct {
	XMLName xml.Name `xml:"GetMetadataConfigurationResponse"`

	// The requested metadata configuration.
	Configuration MetadataConfiguration `xml:"Configuration,omitempty"`
}

// GetAudioOutputConfiguration type
type GetAudioOutputConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputConfiguration"`

	// Token of the requested audio output configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetAudioOutputConfigurationResponse type
type GetAudioOutputConfigurationResponse struct {
	XMLName xml.Name `xml:"GetAudioOutputConfigurationResponse"`

	// The requested audio output configuration.
	Configuration AudioOutputConfiguration `xml:"Configuration,omitempty"`
}

// GetAudioDecoderConfiguration type
type GetAudioDecoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioDecoderConfiguration"`

	// Token of the requested audio decoder configuration.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetAudioDecoderConfigurationResponse type
type GetAudioDecoderConfigurationResponse struct {
	XMLName xml.Name `xml:"GetAudioDecoderConfigurationResponse"`

	// The requested audio decoder configuration
	Configuration AudioDecoderConfiguration `xml:"Configuration,omitempty"`
}

// GetCompatibleVideoEncoderConfigurations type
type GetCompatibleVideoEncoderConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleVideoEncoderConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleVideoEncoderConfigurationsResponse type
type GetCompatibleVideoEncoderConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleVideoEncoderConfigurationsResponse"`

	// Contains a list of video encoder configurations that are compatible with the specified media profile.
	Configurations []VideoEncoderConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleVideoSourceConfigurations type
type GetCompatibleVideoSourceConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleVideoSourceConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleVideoSourceConfigurationsResponse type
type GetCompatibleVideoSourceConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleVideoSourceConfigurationsResponse"`

	// Contains a list of video source configurations that are compatible with the specified media profile.
	Configurations []VideoSourceConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleAudioEncoderConfigurations type
type GetCompatibleAudioEncoderConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioEncoderConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleAudioEncoderConfigurationsResponse type
type GetCompatibleAudioEncoderConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleAudioEncoderConfigurationsResponse"`

	// Contains a list of audio encoder configurations that are compatible with the specified media profile.
	Configurations []AudioEncoderConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleAudioSourceConfigurations type
type GetCompatibleAudioSourceConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioSourceConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleAudioSourceConfigurationsResponse type
type GetCompatibleAudioSourceConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleAudioSourceConfigurationsResponse"`

	// Contains a list of audio source configurations that are compatible with the specified media profile.
	Configurations []AudioSourceConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleVideoAnalyticsConfigurations type
type GetCompatibleVideoAnalyticsConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleVideoAnalyticsConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleVideoAnalyticsConfigurationsResponse type
type GetCompatibleVideoAnalyticsConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleVideoAnalyticsConfigurationsResponse"`

	// Contains a list of video analytics configurations that are compatible with the specified media profile.
	Configurations []VideoAnalyticsConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleMetadataConfigurations type
type GetCompatibleMetadataConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleMetadataConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleMetadataConfigurationsResponse type
type GetCompatibleMetadataConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleMetadataConfigurationsResponse"`

	// Contains a list of metadata configurations that are compatible with the specified media profile.
	Configurations []MetadataConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleAudioOutputConfigurations type
type GetCompatibleAudioOutputConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioOutputConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleAudioOutputConfigurationsResponse type
type GetCompatibleAudioOutputConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleAudioOutputConfigurationsResponse"`

	// Contains a list of audio output configurations that are compatible with the specified media profile.
	Configurations []AudioOutputConfiguration `xml:"Configurations,omitempty"`
}

// GetCompatibleAudioDecoderConfigurations type
type GetCompatibleAudioDecoderConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetCompatibleAudioDecoderConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetCompatibleAudioDecoderConfigurationsResponse type
type GetCompatibleAudioDecoderConfigurationsResponse struct {
	XMLName xml.Name `xml:"GetCompatibleAudioDecoderConfigurationsResponse"`

	// Contains a list of audio decoder configurations that are compatible with the specified media profile.
	Configurations []AudioDecoderConfiguration `xml:"Configurations,omitempty"`
}

// SetVideoEncoderConfiguration type
type SetVideoEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoEncoderConfiguration"`

	// Contains the modified video encoder configuration. The configuration shall exist in the device.
	Configuration VideoEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetVideoEncoderConfigurationResponse type
type SetVideoEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"SetVideoEncoderConfigurationResponse"`
}

// SetVideoSourceConfiguration type
type SetVideoSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoSourceConfiguration"`

	// Contains the modified video source configuration. The configuration shall exist in the device.
	Configuration VideoSourceConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetVideoSourceConfigurationResponse type
type SetVideoSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"SetVideoSourceConfigurationResponse"`
}

// SetAudioEncoderConfiguration type
type SetAudioEncoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioEncoderConfiguration"`

	// Contains the modified audio encoder configuration. The configuration shall exist in the device.
	Configuration AudioEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetAudioEncoderConfigurationResponse type
type SetAudioEncoderConfigurationResponse struct {
	XMLName xml.Name `xml:"SetAudioEncoderConfigurationResponse"`
}

// SetAudioSourceConfiguration type
type SetAudioSourceConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioSourceConfiguration"`

	// Contains the modified audio source configuration. The configuration shall exist in the device.
	Configuration AudioSourceConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetAudioSourceConfigurationResponse type
type SetAudioSourceConfigurationResponse struct {
	XMLName xml.Name `xml:"SetAudioSourceConfigurationResponse"`
}

// SetVideoAnalyticsConfiguration type
type SetVideoAnalyticsConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoAnalyticsConfiguration"`

	// Contains the modified video analytics configuration. The configuration shall exist in the device.
	Configuration VideoAnalyticsConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetVideoAnalyticsConfigurationResponse type
type SetVideoAnalyticsConfigurationResponse struct {
	XMLName xml.Name `xml:"SetVideoAnalyticsConfigurationResponse"`
}

// SetMetadataConfiguration type
type SetMetadataConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetMetadataConfiguration"`

	// Contains the modified metadata configuration. The configuration shall exist in the device.
	Configuration MetadataConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetMetadataConfigurationResponse type
type SetMetadataConfigurationResponse struct {
	XMLName xml.Name `xml:"SetMetadataConfigurationResponse"`
}

// SetAudioOutputConfiguration type
type SetAudioOutputConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioOutputConfiguration"`

	// Contains the modified audio output configuration. The configuration shall exist in the device.
	Configuration AudioOutputConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetAudioOutputConfigurationResponse type
type SetAudioOutputConfigurationResponse struct {
	XMLName xml.Name `xml:"SetAudioOutputConfigurationResponse"`
}

// SetAudioDecoderConfiguration type
type SetAudioDecoderConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetAudioDecoderConfiguration"`

	// Contains the modified audio decoder configuration. The configuration shall exist in the device.
	Configuration AudioDecoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl Configuration,omitempty"`

	// The ForcePersistence element is obsolete and should always be assumed to be true.
	ForcePersistence bool `xml:"http://www.onvif.org/ver10/media/wsdl ForcePersistence,omitempty"`
}

// SetAudioDecoderConfigurationResponse type
type SetAudioDecoderConfigurationResponse struct {
	XMLName xml.Name `xml:"SetAudioDecoderConfigurationResponse"`
}

// GetVideoSourceConfigurationOptions type
type GetVideoSourceConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceConfigurationOptions"`

	// Optional video source configurationToken that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetVideoSourceConfigurationOptionsResponse type
type GetVideoSourceConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetVideoSourceConfigurationOptionsResponse"`

	// This message contains the video source configuration options. If a video source configuration is specified, the options shall concern that particular configuration. If a media profile is specified, the options shall be compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	Options VideoSourceConfigurationOptions `xml:"Options,omitempty"`
}

// GetVideoEncoderConfigurationOptions type
type GetVideoEncoderConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoEncoderConfigurationOptions"`

	// Optional video encoder configuration token that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetVideoEncoderConfigurationOptionsResponse type
type GetVideoEncoderConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetVideoEncoderConfigurationOptionsResponse"`

	Options VideoEncoderConfigurationOptions `xml:"Options,omitempty"`
}

// GetAudioSourceConfigurationOptions type
type GetAudioSourceConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioSourceConfigurationOptions"`

	// Optional audio source configuration token that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetAudioSourceConfigurationOptionsResponse type
type GetAudioSourceConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetAudioSourceConfigurationOptionsResponse"`

	// This message contains the audio source configuration options. If a audio source configuration is specified, the options shall concern that particular configuration. If a media profile is specified, the options shall be compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	Options AudioSourceConfigurationOptions `xml:"Options,omitempty"`
}

// GetAudioEncoderConfigurationOptions type
type GetAudioEncoderConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioEncoderConfigurationOptions"`

	// Optional audio encoder configuration token that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetAudioEncoderConfigurationOptionsResponse type
type GetAudioEncoderConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetAudioEncoderConfigurationOptionsResponse"`

	// This message contains the audio encoder configuration options. If a audio encoder configuration is specified, the options shall concern that particular configuration. If a media profile is specified, the options shall be compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	Options AudioEncoderConfigurationOptions `xml:"Options,omitempty"`
}

// GetMetadataConfigurationOptions type
type GetMetadataConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetMetadataConfigurationOptions"`

	// Optional metadata configuration token that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetMetadataConfigurationOptionsResponse type
type GetMetadataConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetMetadataConfigurationOptionsResponse"`

	// This message contains the metadata configuration options. If a metadata configuration is specified, the options shall concern that particular configuration. If a media profile is specified, the options shall be compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	Options MetadataConfigurationOptions `xml:"Options,omitempty"`
}

// GetAudioOutputConfigurationOptions type
type GetAudioOutputConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioOutputConfigurationOptions"`

	// Optional audio output configuration token that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetAudioOutputConfigurationOptionsResponse type
type GetAudioOutputConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetAudioOutputConfigurationOptionsResponse"`

	// This message contains the audio output configuration options. If a audio output configuration is specified, the options shall concern that particular configuration. If a media profile is specified, the options shall be compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	Options AudioOutputConfigurationOptions `xml:"Options,omitempty"`
}

// GetAudioDecoderConfigurationOptions type
type GetAudioDecoderConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetAudioDecoderConfigurationOptions"`

	// Optional audio decoder configuration token that specifies an existing configuration that the options are intended for.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`

	// Optional ProfileToken that specifies an existing media profile that the options shall be compatible with.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetAudioDecoderConfigurationOptionsResponse type
type GetAudioDecoderConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"GetAudioDecoderConfigurationOptionsResponse"`

	// This message contains the audio decoder configuration options. If a audio decoder configuration is specified, the options shall concern that particular configuration. If a media profile is specified, the options shall be compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	Options AudioDecoderConfigurationOptions `xml:"Options,omitempty"`
}

// GetGuaranteedNumberOfVideoEncoderInstances type
type GetGuaranteedNumberOfVideoEncoderInstances struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetGuaranteedNumberOfVideoEncoderInstances"`

	// Token of the video source configuration
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetGuaranteedNumberOfVideoEncoderInstancesResponse type
type GetGuaranteedNumberOfVideoEncoderInstancesResponse struct {
	XMLName xml.Name `xml:"GetGuaranteedNumberOfVideoEncoderInstancesResponse"`

	// The minimum guaranteed total number of encoder instances (applications) per VideoSourceConfiguration. The device is able to deliver the TotalNumber of streams
	TotalNumber int32 `xml:"TotalNumber,omitempty"`

	// If a device limits the number of instances for respective Video Codecs the response contains the information how many Jpeg streams can be set up at the same time per VideoSource.
	JPEG int32 `xml:"JPEG,omitempty"`

	// If a device limits the number of instances for respective Video Codecs the response contains the information how many H264 streams can be set up at the same time per VideoSource.
	H264 int32 `xml:"H264,omitempty"`

	// If a device limits the number of instances for respective Video Codecs the response contains the information how many Mpeg4 streams can be set up at the same time per VideoSource.
	MPEG4 int32 `xml:"MPEG4,omitempty"`
}

// GetStreamUri type
type GetStreamUri struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetStreamUri"`

	// Stream Setup that should be used with the uri
	StreamSetup StreamSetup `xml:"http://www.onvif.org/ver10/media/wsdl StreamSetup,omitempty"`

	// The ProfileToken element indicates the media profile to use and will define the configuration of the content of the stream.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetStreamUriResponse type
type GetStreamUriResponse struct {
	XMLName xml.Name `xml:"GetStreamUriResponse"`

	MediaUri MediaUri `xml:"MediaUri,omitempty"`
}

// StartMulticastStreaming type
type StartMulticastStreaming struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl StartMulticastStreaming"`

	// Contains the token of the Profile that is used to define the multicast stream.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// StartMulticastStreamingResponse type
type StartMulticastStreamingResponse struct {
	XMLName xml.Name `xml:"StartMulticastStreamingResponse"`
}

// StopMulticastStreaming type
type StopMulticastStreaming struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl StopMulticastStreaming"`

	// Contains the token of the Profile that is used to define the multicast stream.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// StopMulticastStreamingResponse type
type StopMulticastStreamingResponse struct {
	XMLName xml.Name `xml:"StopMulticastStreamingResponse"`
}

// SetSynchronizationPoint type
type SetSynchronizationPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetSynchronizationPoint"`

	// Contains a Profile reference for which a Synchronization Point is requested.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// SetSynchronizationPointResponse type
type SetSynchronizationPointResponse struct {
	XMLName xml.Name `xml:"SetSynchronizationPointResponse"`
}

// GetSnapshotUri type
type GetSnapshotUri struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetSnapshotUri"`

	// The ProfileToken element indicates the media profile to use and will define the source and dimensions of the snapshot.
	ProfileToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ProfileToken,omitempty"`
}

// GetSnapshotUriResponse type
type GetSnapshotUriResponse struct {
	XMLName xml.Name `xml:"GetSnapshotUriResponse"`

	MediaUri MediaUri `xml:"MediaUri,omitempty"`
}

// GetVideoSourceModes type
type GetVideoSourceModes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetVideoSourceModes"`

	// Contains a video source reference for which a video source mode is requested.
	VideoSourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceToken,omitempty"`
}

// GetVideoSourceModesResponse type
type GetVideoSourceModesResponse struct {
	XMLName xml.Name `xml:"GetVideoSourceModesResponse"`

	// Return the information for specified video source mode.
	VideoSourceModes []VideoSourceMode `xml:"VideoSourceModes,omitempty"`
}

// SetVideoSourceMode type
type SetVideoSourceMode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetVideoSourceMode"`

	// Contains a video source reference for which a video source mode is requested.
	VideoSourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceToken,omitempty"`

	// Indicate video source mode.
	VideoSourceModeToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceModeToken,omitempty"`
}

// SetVideoSourceModeResponse type
type SetVideoSourceModeResponse struct {
	XMLName xml.Name `xml:"SetVideoSourceModeResponse"`

	// The response contains information about rebooting after returning response. When Reboot is set true, a device will reboot automatically after setting mode.
	Reboot bool `xml:"Reboot,omitempty"`
}

// GetOSDs type
type GetOSDs struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetOSDs"`

	// Token of the Video Source Configuration, which has OSDs associated with are requested. If token not exist, request all available OSDs.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetOSDsResponse type
type GetOSDsResponse struct {
	XMLName xml.Name `xml:"GetOSDsResponse"`

	// This element contains a list of requested OSDs.
	OSDs []OSDConfiguration `xml:"OSDs,omitempty"`
}

// GetOSD type
type GetOSD struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetOSD"`

	// The GetOSD command fetches the OSD configuration if the OSD token is known.
	OSDToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl OSDToken,omitempty"`
}

// GetOSDResponse type
type GetOSDResponse struct {
	XMLName xml.Name `xml:"GetOSDResponse"`

	// The requested OSD configuration.
	OSD OSDConfiguration `xml:"OSD,omitempty"`
}

// SetOSD type
type SetOSD struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl SetOSD"`

	// Contains the modified OSD configuration.
	OSD OSDConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl OSD,omitempty"`
}

// SetOSDResponse type
type SetOSDResponse struct {
	XMLName xml.Name `xml:"SetOSDResponse"`
}

// GetOSDOptions type
type GetOSDOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl GetOSDOptions"`

	// Video Source Configuration Token that specifies an existing video source configuration that the options shall be compatible with.
	ConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl ConfigurationToken,omitempty"`
}

// GetOSDOptionsResponse type
type GetOSDOptionsResponse struct {
	XMLName xml.Name `xml:"GetOSDOptionsResponse"`

	OSDOptions OSDConfigurationOptions `xml:"OSDOptions,omitempty"`
}

// CreateOSD type
type CreateOSD struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl CreateOSD"`

	// Contain the initial OSD configuration for create.
	OSD OSDConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl OSD,omitempty"`
}

// CreateOSDResponse type
type CreateOSDResponse struct {
	XMLName xml.Name `xml:"CreateOSDResponse"`

	// Returns Token of the newly created OSD
	OSDToken ReferenceToken `xml:"OSDToken,omitempty"`
}

// DeleteOSD type
type DeleteOSD struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/media/wsdl DeleteOSD"`

	// This element contains a reference to the OSD configuration that should be deleted.
	OSDToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl OSDToken,omitempty"`
}

// DeleteOSDResponse type
type DeleteOSDResponse struct {
	XMLName xml.Name `xml:"DeleteOSDResponse"`
}

// Capabilities type
type Capabilities struct {

	// Media profile capabilities.
	ProfileCapabilities ProfileCapabilities `xml:"http://www.onvif.org/ver10/media/wsdl ProfileCapabilities,omitempty"`

	// Streaming capabilities.
	StreamingCapabilities StreamingCapabilities `xml:"http://www.onvif.org/ver10/media/wsdl StreamingCapabilities,omitempty"`

	// Indicates if GetSnapshotUri is supported.

	SnapshotUri bool `xml:"http://www.onvif.org/ver10/media/wsdl SnapshotUri,attr,omitempty"`

	// Indicates whether or not Rotation feature is supported.

	Rotation bool `xml:"http://www.onvif.org/ver10/media/wsdl Rotation,attr,omitempty"`

	// Indicates the support for changing video source mode.

	VideoSourceMode bool `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceMode,attr,omitempty"`

	// Indicates if OSD is supported.

	OSD bool `xml:"http://www.onvif.org/ver10/media/wsdl OSD,attr,omitempty"`

	// Indicates the support for temporary osd text configuration.

	TemporaryOSDText bool `xml:"http://www.onvif.org/ver10/media/wsdl TemporaryOSDText,attr,omitempty"`

	// Indicates the support for the Efficient XML Interchange (EXI) binary XML format.

	EXICompression bool `xml:"http://www.onvif.org/ver10/media/wsdl EXICompression,attr,omitempty"`
}

// ProfileCapabilities type
type ProfileCapabilities struct {

	// Maximum number of profiles supported.

	MaximumNumberOfProfiles int32 `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfProfiles,attr,omitempty"`
}

// StreamingCapabilities type
type StreamingCapabilities struct {

	// Indicates support for RTP multicast.

	RTPMulticast bool `xml:"http://www.onvif.org/ver10/media/wsdl RTPMulticast,attr,omitempty"`

	// Indicates support for RTP over TCP.

	RTP_TCP bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_TCP,attr,omitempty"`

	// Indicates support for RTP/RTSP/TCP.

	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_RTSP_TCP,attr,omitempty"`

	// Indicates support for non aggregate RTSP control.

	NonAggregateControl bool `xml:"http://www.onvif.org/ver10/media/wsdl NonAggregateControl,attr,omitempty"`

	// Indicates the device does not support live media streaming via RTSP.

	NoRTSPStreaming bool `xml:"http://www.onvif.org/ver10/media/wsdl NoRTSPStreaming,attr,omitempty"`
}

// VideoSourceMode type
type VideoSourceMode struct {

	// Max frame rate in frames per second for this video source mode.
	MaxFramerate float32 `xml:"http://www.onvif.org/ver10/schema MaxFramerate,omitempty"`

	// Max horizontal and vertical resolution for this video source mode.
	MaxResolution VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl MaxResolution,omitempty"`

	// Indication which encodings are supported for this video source. The list may contain one or more enumeration values of tt:VideoEncoding.
	Encodings EncodingTypes `xml:"http://www.onvif.org/ver10/media/wsdl Encodings,omitempty"`

	// After setting the mode if a device starts to reboot this value is true. If a device change the mode without rebooting this value is false. If true, configured parameters may not be guaranteed by the device after rebooting.
	Reboot bool `xml:"http://www.onvif.org/ver10/media/wsdl Reboot,omitempty"`

	// Informative description of this video source mode. This field should be described in English.
	Description Description `xml:"http://www.onvif.org/ver10/media/wsdl Description,omitempty"`

	Extension VideoSourceModeExtension `xml:"http://www.onvif.org/ver10/media/wsdl Extension,omitempty"`

	// Indicate token for video source mode.

	Token ReferenceToken `xml:"token,attr,omitempty"`

	// Indication of whether this mode is active. If active this value is true. In case of non-indication, it means as false. The value of true shall be had by only one video source mode.

	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,attr,omitempty"`
}

// VideoSourceModeExtension type
type VideoSourceModeExtension struct {
}

// FaultcodeEnum type
type FaultcodeEnum QName

const (
	// FaultcodeEnumTnsDataEncodingUnknown const
	FaultcodeEnumTnsDataEncodingUnknown FaultcodeEnum = "tns:DataEncodingUnknown"

	// FaultcodeEnumTnsMustUnderstand const
	FaultcodeEnumTnsMustUnderstand FaultcodeEnum = "tns:MustUnderstand"

	// FaultcodeEnumTnsReceiver const
	FaultcodeEnumTnsReceiver FaultcodeEnum = "tns:Receiver"

	// FaultcodeEnumTnsSender const
	FaultcodeEnumTnsSender FaultcodeEnum = "tns:Sender"

	// FaultcodeEnumTnsVersionMismatch const
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
	// RelationshipTypeHttpwwww3org200508addressingreply const
	RelationshipTypeHttpwwww3org200508addressingreply RelationshipType = "http://www.w3.org/2005/08/addressing/reply"
)

// FaultCodesType type
type FaultCodesType QName

const (
	// FaultCodesTypeTnsInvalidAddressingHeader const
	FaultCodesTypeTnsInvalidAddressingHeader FaultCodesType = "tns:InvalidAddressingHeader"

	// FaultCodesTypeTnsInvalidAddress const
	FaultCodesTypeTnsInvalidAddress FaultCodesType = "tns:InvalidAddress"

	// FaultCodesTypeTnsInvalidEPR const
	FaultCodesTypeTnsInvalidEPR FaultCodesType = "tns:InvalidEPR"

	// FaultCodesTypeTnsInvalidCardinality const
	FaultCodesTypeTnsInvalidCardinality FaultCodesType = "tns:InvalidCardinality"

	// FaultCodesTypeTnsMissingAddressInEPR const
	FaultCodesTypeTnsMissingAddressInEPR FaultCodesType = "tns:MissingAddressInEPR"

	// FaultCodesTypeTnsDuplicateMessageID const
	FaultCodesTypeTnsDuplicateMessageID FaultCodesType = "tns:DuplicateMessageID"

	// FaultCodesTypeTnsActionMismatch const
	FaultCodesTypeTnsActionMismatch FaultCodesType = "tns:ActionMismatch"

	// FaultCodesTypeTnsMessageAddressingHeaderRequired const
	FaultCodesTypeTnsMessageAddressingHeaderRequired FaultCodesType = "tns:MessageAddressingHeaderRequired"

	// FaultCodesTypeTnsDestinationUnreachable const
	FaultCodesTypeTnsDestinationUnreachable FaultCodesType = "tns:DestinationUnreachable"

	// FaultCodesTypeTnsActionNotSupported const
	FaultCodesTypeTnsActionNotSupported FaultCodesType = "tns:ActionNotSupported"

	// FaultCodesTypeTnsEndpointUnavailable const
	FaultCodesTypeTnsEndpointUnavailable FaultCodesType = "tns:EndpointUnavailable"
)

// EndpointReference type
type EndpointReference EndpointReferenceType

// Metadata type
type Metadata MetadataType

// RelatesTo type
type RelatesTo RelatesToType

// To type
type To AttributedURIType

// Action type
type Action AttributedURIType

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

// ConsumerReference type
type ConsumerReference EndpointReferenceType

// Filter type
type Filter FilterType

// SubscriptionPolicy type
type SubscriptionPolicy SubscriptionPolicyType

// CreationTime type
type CreationTime time.Time

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

// Subscribe type
type Subscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Subscribe"`

	ConsumerReference EndpointReferenceType `xml:"ConsumerReference,omitempty"`

	Filter FilterType `xml:"Filter,omitempty"`

	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
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

// UnableToGetMessagesFault type
type UnableToGetMessagesFault UnableToGetMessagesFaultType

// UnableToDestroyPullPointFault type
type UnableToDestroyPullPointFault UnableToDestroyPullPointFaultType

// UnableToCreatePullPointFault type
type UnableToCreatePullPointFault UnableToCreatePullPointFaultType

// UnacceptableTerminationTimeFault type
type UnacceptableTerminationTimeFault UnacceptableTerminationTimeFaultType

// UnableToDestroySubscriptionFault type
type UnableToDestroySubscriptionFault UnableToDestroySubscriptionFaultType

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
	XMLName xml.Name `xml:"Filter"`
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
	// MoveStatusIDLE const
	MoveStatusIDLE MoveStatus = "IDLE"

	// MoveStatusMOVING const
	MoveStatusMOVING MoveStatus = "MOVING"

	// MoveStatusUNKNOWN const
	MoveStatusUNKNOWN MoveStatus = "UNKNOWN"
)

// Entity type
type Entity string

const (
	// EntityDevice const
	EntityDevice Entity = "Device"

	// EntityVideoSource const
	EntityVideoSource Entity = "VideoSource"

	// EntityAudioSource const
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
	Error string `xml:"http://www.onvif.org/ver10/schema Error,omitempty"`

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

// Transformation type
type Transformation struct {
	Translate Vector `xml:"http://www.onvif.org/ver10/schema Translate,omitempty"`

	Scale Vector `xml:"http://www.onvif.org/ver10/schema Scale,omitempty"`

	Extension TransformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// TransformationExtension type
type TransformationExtension struct {
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
	// RotateModeOFF const
	RotateModeOFF RotateMode = "OFF"

	// Disable the Rotate feature.
	// RotateModeON const
	RotateModeON RotateMode = "ON"

	// Rotate feature is automatically activated by the device.
	// RotateModeAUTO const
	RotateModeAUTO RotateMode = "AUTO"
)

// SceneOrientationMode type
type SceneOrientationMode string

const (
	// SceneOrientationModeMANUAL const
	SceneOrientationModeMANUAL SceneOrientationMode = "MANUAL"

	// SceneOrientationModeAUTO const
	SceneOrientationModeAUTO SceneOrientationMode = "AUTO"
)

// SceneOrientationOption type
type SceneOrientationOption string

const (
	// SceneOrientationOptionBelow const
	SceneOrientationOptionBelow SceneOrientationOption = "Below"

	// SceneOrientationOptionHorizon const
	SceneOrientationOptionHorizon SceneOrientationOption = "Horizon"

	// SceneOrientationOptionAbove const
	SceneOrientationOptionAbove SceneOrientationOption = "Above"
)

// Source view modes supported by device.

// ViewModes type
type ViewModes string

const (

	// Undewarped viewmode from device supporting fisheye lens.
	// ViewModesTtFisheye const
	ViewModesTtFisheye ViewModes = "tt:Fisheye"

	// 360 degree panoramic view.
	// ViewModesTt360Panorama const
	ViewModesTt360Panorama ViewModes = "tt:360Panorama"

	// 180 degree panoramic view.
	// ViewModesTt180Panorama const
	ViewModesTt180Panorama ViewModes = "tt:180Panorama"

	// View mode combining four streams in single Quad, eg., applicable for devices supporting four heads.
	// ViewModesTtQuad const
	ViewModesTtQuad ViewModes = "tt:Quad"

	// Unaltered view from the sensor.
	// ViewModesTtOriginal const
	ViewModesTtOriginal ViewModes = "tt:Original"

	// Viewmode combining the left side sensors, applicable for devices supporting multiple sensors.
	// ViewModesTtLeftHalf const
	ViewModesTtLeftHalf ViewModes = "tt:LeftHalf"

	// Viewmode combining the right side sensors, applicable for devices supporting multiple sensors.
	// ViewModesTtRightHalf const
	ViewModesTtRightHalf ViewModes = "tt:RightHalf"

	// Dewarped view mode for device supporting fisheye lens.
	// ViewModesTtDewarp const
	ViewModesTtDewarp ViewModes = "tt:Dewarp"
)

// VideoEncoding type
type VideoEncoding string

const (
	// VideoEncodingJPEG const
	VideoEncodingJPEG VideoEncoding = "JPEG"

	// VideoEncodingMPEG4 const
	VideoEncodingMPEG4 VideoEncoding = "MPEG4"

	// VideoEncodingH264 const
	VideoEncodingH264 VideoEncoding = "H264"
)

// Mpeg4Profile type
type Mpeg4Profile string

const (
	// Mpeg4ProfileSP const
	Mpeg4ProfileSP Mpeg4Profile = "SP"

	// Mpeg4ProfileASP const
	Mpeg4ProfileASP Mpeg4Profile = "ASP"
)

// H264Profile type
type H264Profile string

const (
	// H264ProfileBaseline const
	H264ProfileBaseline H264Profile = "Baseline"

	// H264ProfileMain const
	H264ProfileMain H264Profile = "Main"

	// H264ProfileExtended const
	H264ProfileExtended H264Profile = "Extended"

	// H264ProfileHigh const
	H264ProfileHigh H264Profile = "High"
)

// Video Media Subtypes as referenced by IANA (without the leading "video/" Video Media Type).  See also .

// VideoEncodingMimeNames type
type VideoEncodingMimeNames string

const (
	// VideoEncodingMimeNamesJPEG const
	VideoEncodingMimeNamesJPEG VideoEncodingMimeNames = "JPEG"

	// VideoEncodingMimeNamesMPV4ES const
	VideoEncodingMimeNamesMPV4ES VideoEncodingMimeNames = "MPV4-ES"

	// VideoEncodingMimeNamesH264 const
	VideoEncodingMimeNamesH264 VideoEncodingMimeNames = "H264"

	// VideoEncodingMimeNamesH265 const
	VideoEncodingMimeNamesH265 VideoEncodingMimeNames = "H265"
)

// VideoEncodingProfiles type
type VideoEncodingProfiles string

const (
	// VideoEncodingProfilesSimple const
	VideoEncodingProfilesSimple VideoEncodingProfiles = "Simple"

	// VideoEncodingProfilesAdvancedSimple const
	VideoEncodingProfilesAdvancedSimple VideoEncodingProfiles = "AdvancedSimple"

	// VideoEncodingProfilesBaseline const
	VideoEncodingProfilesBaseline VideoEncodingProfiles = "Baseline"

	// VideoEncodingProfilesMain const
	VideoEncodingProfilesMain VideoEncodingProfiles = "Main"

	// VideoEncodingProfilesMain10 const
	VideoEncodingProfilesMain10 VideoEncodingProfiles = "Main10"

	// VideoEncodingProfilesExtended const
	VideoEncodingProfilesExtended VideoEncodingProfiles = "Extended"

	// VideoEncodingProfilesHigh const
	VideoEncodingProfilesHigh VideoEncodingProfiles = "High"
)

// AudioEncoding type
type AudioEncoding string

const (
	// AudioEncodingG711 const
	AudioEncodingG711 AudioEncoding = "G711"

	// AudioEncodingG726 const
	AudioEncodingG726 AudioEncoding = "G726"

	// AudioEncodingAAC const
	AudioEncodingAAC AudioEncoding = "AAC"
)

// Audio Media Subtypes as referenced by IANA (without the leading "audio/" Audio Media Type).  See also .

// AudioEncodingMimeNames type
type AudioEncodingMimeNames string

const (
	// AudioEncodingMimeNamesPCMU const
	AudioEncodingMimeNamesPCMU AudioEncodingMimeNames = "PCMU"

	// AudioEncodingMimeNamesG726 const
	AudioEncodingMimeNamesG726 AudioEncodingMimeNames = "G726"

	// AudioEncodingMimeNamesMP4ALATM const
	AudioEncodingMimeNamesMP4ALATM AudioEncodingMimeNames = "MP4A-LATM"

	// AudioEncodingMimeNamesMpeg4generic const
	AudioEncodingMimeNamesMpeg4generic AudioEncodingMimeNames = "mpeg4-generic"
)

// MetadataCompressionType type
type MetadataCompressionType string

const (
	// MetadataCompressionTypeNone const
	MetadataCompressionTypeNone MetadataCompressionType = "None"

	// MetadataCompressionTypeGZIP const
	MetadataCompressionTypeGZIP MetadataCompressionType = "GZIP"

	// MetadataCompressionTypeEXI const
	MetadataCompressionTypeEXI MetadataCompressionType = "EXI"
)

// StreamType type
type StreamType string

const (
	// StreamTypeRTPUnicast const
	StreamTypeRTPUnicast StreamType = "RTP-Unicast"

	// StreamTypeRTPMulticast const
	StreamTypeRTPMulticast StreamType = "RTP-Multicast"
)

// TransportProtocol type
type TransportProtocol string

const (
	// TransportProtocolUDP const
	TransportProtocolUDP TransportProtocol = "UDP"

	// This value is deprecated.
	// TransportProtocolTCP const
	TransportProtocolTCP TransportProtocol = "TCP"

	// TransportProtocolRTSP const
	TransportProtocolRTSP TransportProtocol = "RTSP"

	// TransportProtocolHTTP const
	TransportProtocolHTTP TransportProtocol = "HTTP"
)

// ScopeDefinition type
type ScopeDefinition string

const (
	// ScopeDefinitionFixed const
	ScopeDefinitionFixed ScopeDefinition = "Fixed"

	// ScopeDefinitionConfigurable const
	ScopeDefinitionConfigurable ScopeDefinition = "Configurable"
)

// DiscoveryMode type
type DiscoveryMode string

const (
	// DiscoveryModeDiscoverable const
	DiscoveryModeDiscoverable DiscoveryMode = "Discoverable"

	// DiscoveryModeNonDiscoverable const
	DiscoveryModeNonDiscoverable DiscoveryMode = "NonDiscoverable"
)

// NetworkInterfaceConfigPriority type
type NetworkInterfaceConfigPriority int32

// Duplex type
type Duplex string

const (
	// DuplexFull const
	DuplexFull Duplex = "Full"

	// DuplexHalf const
	DuplexHalf Duplex = "Half"
)

// IANAIfTypes type
type IANAIfTypes int32

// IPv6DHCPConfiguration type
type IPv6DHCPConfiguration string

const (
	// IPv6DHCPConfigurationAuto const
	IPv6DHCPConfigurationAuto IPv6DHCPConfiguration = "Auto"

	// IPv6DHCPConfigurationStateful const
	IPv6DHCPConfigurationStateful IPv6DHCPConfiguration = "Stateful"

	// IPv6DHCPConfigurationStateless const
	IPv6DHCPConfigurationStateless IPv6DHCPConfiguration = "Stateless"

	// IPv6DHCPConfigurationOff const
	IPv6DHCPConfigurationOff IPv6DHCPConfiguration = "Off"
)

// NetworkProtocolType type
type NetworkProtocolType string

const (
	// NetworkProtocolTypeHTTP const
	NetworkProtocolTypeHTTP NetworkProtocolType = "HTTP"

	// NetworkProtocolTypeHTTPS const
	NetworkProtocolTypeHTTPS NetworkProtocolType = "HTTPS"

	// NetworkProtocolTypeRTSP const
	NetworkProtocolTypeRTSP NetworkProtocolType = "RTSP"
)

// NetworkHostType type
type NetworkHostType string

const (
	// NetworkHostTypeIPv4 const
	NetworkHostTypeIPv4 NetworkHostType = "IPv4"

	// NetworkHostTypeIPv6 const
	NetworkHostTypeIPv6 NetworkHostType = "IPv6"

	// NetworkHostTypeDNS const
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
	// IPTypeIPv4 const
	IPTypeIPv4 IPType = "IPv4"

	// IPTypeIPv6 const
	IPTypeIPv6 IPType = "IPv6"
)

// DNSName type
type DNSName string

// IPAddressFilterType type
type IPAddressFilterType string

const (
	// IPAddressFilterTypeAllow const
	IPAddressFilterTypeAllow IPAddressFilterType = "Allow"

	// IPAddressFilterTypeDeny const
	IPAddressFilterTypeDeny IPAddressFilterType = "Deny"
)

// DynamicDNSType type
type DynamicDNSType string

const (
	// DynamicDNSTypeNoUpdate const
	DynamicDNSTypeNoUpdate DynamicDNSType = "NoUpdate"

	// DynamicDNSTypeClientUpdates const
	DynamicDNSTypeClientUpdates DynamicDNSType = "ClientUpdates"

	// DynamicDNSTypeServerUpdates const
	DynamicDNSTypeServerUpdates DynamicDNSType = "ServerUpdates"
)

// Dot11SSIDType type
type Dot11SSIDType []byte

// Dot11StationMode type
type Dot11StationMode string

const (
	// Dot11StationModeAdhoc const
	Dot11StationModeAdhoc Dot11StationMode = "Ad-hoc"

	// Dot11StationModeInfrastructure const
	Dot11StationModeInfrastructure Dot11StationMode = "Infrastructure"

	// Dot11StationModeExtended const
	Dot11StationModeExtended Dot11StationMode = "Extended"
)

// Dot11SecurityMode type
type Dot11SecurityMode string

const (
	// Dot11SecurityModeNone const
	Dot11SecurityModeNone Dot11SecurityMode = "None"

	// Dot11SecurityModeWEP const
	Dot11SecurityModeWEP Dot11SecurityMode = "WEP"

	// Dot11SecurityModePSK const
	Dot11SecurityModePSK Dot11SecurityMode = "PSK"

	// Dot11SecurityModeDot1X const
	Dot11SecurityModeDot1X Dot11SecurityMode = "Dot1X"

	// Dot11SecurityModeExtended const
	Dot11SecurityModeExtended Dot11SecurityMode = "Extended"
)

// Dot11Cipher type
type Dot11Cipher string

const (
	// Dot11CipherCCMP const
	Dot11CipherCCMP Dot11Cipher = "CCMP"

	// Dot11CipherTKIP const
	Dot11CipherTKIP Dot11Cipher = "TKIP"

	// Dot11CipherAny const
	Dot11CipherAny Dot11Cipher = "Any"

	// Dot11CipherExtended const
	Dot11CipherExtended Dot11Cipher = "Extended"
)

// Dot11PSK type
type Dot11PSK []byte

// Dot11PSKPassphrase type
type Dot11PSKPassphrase string

// Dot11SignalStrength type
type Dot11SignalStrength string

const (
	// Dot11SignalStrengthNone const
	Dot11SignalStrengthNone Dot11SignalStrength = "None"

	// Dot11SignalStrengthVeryBad const
	Dot11SignalStrengthVeryBad Dot11SignalStrength = "Very Bad"

	// Dot11SignalStrengthBad const
	Dot11SignalStrengthBad Dot11SignalStrength = "Bad"

	// Dot11SignalStrengthGood const
	Dot11SignalStrengthGood Dot11SignalStrength = "Good"

	// Dot11SignalStrengthVeryGood const
	Dot11SignalStrengthVeryGood Dot11SignalStrength = "Very Good"

	// Dot11SignalStrengthExtended const
	Dot11SignalStrengthExtended Dot11SignalStrength = "Extended"
)

// Dot11AuthAndMangementSuite type
type Dot11AuthAndMangementSuite string

const (
	// Dot11AuthAndMangementSuiteNone const
	Dot11AuthAndMangementSuiteNone Dot11AuthAndMangementSuite = "None"

	// Dot11AuthAndMangementSuiteDot1X const
	Dot11AuthAndMangementSuiteDot1X Dot11AuthAndMangementSuite = "Dot1X"

	// Dot11AuthAndMangementSuitePSK const
	Dot11AuthAndMangementSuitePSK Dot11AuthAndMangementSuite = "PSK"

	// Dot11AuthAndMangementSuiteExtended const
	Dot11AuthAndMangementSuiteExtended Dot11AuthAndMangementSuite = "Extended"
)

// CapabilityCategory type
type CapabilityCategory string

const (
	// CapabilityCategoryAll const
	CapabilityCategoryAll CapabilityCategory = "All"

	// CapabilityCategoryAnalytics const
	CapabilityCategoryAnalytics CapabilityCategory = "Analytics"

	// CapabilityCategoryDevice const
	CapabilityCategoryDevice CapabilityCategory = "Device"

	// CapabilityCategoryEvents const
	CapabilityCategoryEvents CapabilityCategory = "Events"

	// CapabilityCategoryImaging const
	CapabilityCategoryImaging CapabilityCategory = "Imaging"

	// CapabilityCategoryMedia const
	CapabilityCategoryMedia CapabilityCategory = "Media"

	// CapabilityCategoryPTZ const
	CapabilityCategoryPTZ CapabilityCategory = "PTZ"
)

// Enumeration describing the available system log modes.

// SystemLogType type
type SystemLogType string

const (

	// Indicates that a system log is requested.
	// SystemLogTypeSystem const
	SystemLogTypeSystem SystemLogType = "System"

	// Indicates that a access log is requested.
	// SystemLogTypeAccess const
	SystemLogTypeAccess SystemLogType = "Access"
)

// Enumeration describing the available factory default modes.

// FactoryDefaultType type
type FactoryDefaultType string

const (

	// Indicates that a hard factory default is requested.
	// FactoryDefaultTypeHard const
	FactoryDefaultTypeHard FactoryDefaultType = "Hard"

	// Indicates that a soft factory default is requested.
	// FactoryDefaultTypeSoft const
	FactoryDefaultTypeSoft FactoryDefaultType = "Soft"
)

// SetDateTimeType type
type SetDateTimeType string

const (

	// Indicates that the date and time are set manually.
	// SetDateTimeTypeManual const
	SetDateTimeTypeManual SetDateTimeType = "Manual"

	// Indicates that the date and time are set through NTP
	// SetDateTimeTypeNTP const
	SetDateTimeTypeNTP SetDateTimeType = "NTP"
)

// UserLevel type
type UserLevel string

const (
	// UserLevelAdministrator const
	UserLevelAdministrator UserLevel = "Administrator"

	// UserLevelOperator const
	UserLevelOperator UserLevel = "Operator"

	// UserLevelUser const
	UserLevelUser UserLevel = "User"

	// UserLevelAnonymous const
	UserLevelAnonymous UserLevel = "Anonymous"

	// UserLevelExtended const
	UserLevelExtended UserLevel = "Extended"
)

// RelayLogicalState type
type RelayLogicalState string

const (
	// RelayLogicalStateActive const
	RelayLogicalStateActive RelayLogicalState = "active"

	// RelayLogicalStateInactive const
	RelayLogicalStateInactive RelayLogicalState = "inactive"
)

// RelayIdleState type
type RelayIdleState string

const (
	// RelayIdleStateClosed const
	RelayIdleStateClosed RelayIdleState = "closed"

	// RelayIdleStateOpen const
	RelayIdleStateOpen RelayIdleState = "open"
)

// RelayMode type
type RelayMode string

const (
	// RelayModeMonostable const
	RelayModeMonostable RelayMode = "Monostable"

	// RelayModeBistable const
	RelayModeBistable RelayMode = "Bistable"
)

// DigitalIdleState type
type DigitalIdleState string

const (
	// DigitalIdleStateClosed const
	DigitalIdleStateClosed DigitalIdleState = "closed"

	// DigitalIdleStateOpen const
	DigitalIdleStateOpen DigitalIdleState = "open"
)

// EFlipMode type
type EFlipMode string

const (
	// EFlipModeOFF const
	EFlipModeOFF EFlipMode = "OFF"

	// EFlipModeON const
	EFlipModeON EFlipMode = "ON"

	// EFlipModeExtended const
	EFlipModeExtended EFlipMode = "Extended"
)

// ReverseMode type
type ReverseMode string

const (
	// ReverseModeOFF const
	ReverseModeOFF ReverseMode = "OFF"

	// ReverseModeON const
	ReverseModeON ReverseMode = "ON"

	// ReverseModeAUTO const
	ReverseModeAUTO ReverseMode = "AUTO"

	// ReverseModeExtended const
	ReverseModeExtended ReverseMode = "Extended"
)

// AuxiliaryData type
type AuxiliaryData string

// PTZPresetTourState type
type PTZPresetTourState string

const (
	// PTZPresetTourStateIdle const
	PTZPresetTourStateIdle PTZPresetTourState = "Idle"

	// PTZPresetTourStateTouring const
	PTZPresetTourStateTouring PTZPresetTourState = "Touring"

	// PTZPresetTourStatePaused const
	PTZPresetTourStatePaused PTZPresetTourState = "Paused"

	// PTZPresetTourStateExtended const
	PTZPresetTourStateExtended PTZPresetTourState = "Extended"
)

// PTZPresetTourDirection type
type PTZPresetTourDirection string

const (
	// PTZPresetTourDirectionForward const
	PTZPresetTourDirectionForward PTZPresetTourDirection = "Forward"

	// PTZPresetTourDirectionBackward const
	PTZPresetTourDirectionBackward PTZPresetTourDirection = "Backward"

	// PTZPresetTourDirectionExtended const
	PTZPresetTourDirectionExtended PTZPresetTourDirection = "Extended"
)

// PTZPresetTourOperation type
type PTZPresetTourOperation string

const (
	// PTZPresetTourOperationStart const
	PTZPresetTourOperationStart PTZPresetTourOperation = "Start"

	// PTZPresetTourOperationStop const
	PTZPresetTourOperationStop PTZPresetTourOperation = "Stop"

	// PTZPresetTourOperationPause const
	PTZPresetTourOperationPause PTZPresetTourOperation = "Pause"

	// PTZPresetTourOperationExtended const
	PTZPresetTourOperationExtended PTZPresetTourOperation = "Extended"
)

// AutoFocusMode type
type AutoFocusMode string

const (
	// AutoFocusModeAUTO const
	AutoFocusModeAUTO AutoFocusMode = "AUTO"

	// AutoFocusModeMANUAL const
	AutoFocusModeMANUAL AutoFocusMode = "MANUAL"
)

// AFModes type
type AFModes string

const (

	// Focus of a moving camera is updated only once after stopping a pan, tilt or zoom movement.
	// AFModesOnceAfterMove const
	AFModesOnceAfterMove AFModes = "OnceAfterMove"
)

// WideDynamicMode type
type WideDynamicMode string

const (
	// WideDynamicModeOFF const
	WideDynamicModeOFF WideDynamicMode = "OFF"

	// WideDynamicModeON const
	WideDynamicModeON WideDynamicMode = "ON"
)

// Enumeration describing the available backlight compenstation modes.

// BacklightCompensationMode type
type BacklightCompensationMode string

const (

	// Backlight compensation is disabled.
	// BacklightCompensationModeOFF const
	BacklightCompensationModeOFF BacklightCompensationMode = "OFF"

	// Backlight compensation is enabled.
	// BacklightCompensationModeON const
	BacklightCompensationModeON BacklightCompensationMode = "ON"
)

// ExposurePriority type
type ExposurePriority string

const (
	// ExposurePriorityLowNoise const
	ExposurePriorityLowNoise ExposurePriority = "LowNoise"

	// ExposurePriorityFrameRate const
	ExposurePriorityFrameRate ExposurePriority = "FrameRate"
)

// ExposureMode type
type ExposureMode string

const (
	// ExposureModeAUTO const
	ExposureModeAUTO ExposureMode = "AUTO"

	// ExposureModeMANUAL const
	ExposureModeMANUAL ExposureMode = "MANUAL"
)

// Enabled type
type Enabled string

const (
	// EnabledENABLED const
	EnabledENABLED Enabled = "ENABLED"

	// EnabledDISABLED const
	EnabledDISABLED Enabled = "DISABLED"
)

// WhiteBalanceMode type
type WhiteBalanceMode string

const (
	// WhiteBalanceModeAUTO const
	WhiteBalanceModeAUTO WhiteBalanceMode = "AUTO"

	// WhiteBalanceModeMANUAL const
	WhiteBalanceModeMANUAL WhiteBalanceMode = "MANUAL"
)

// IrCutFilterMode type
type IrCutFilterMode string

const (
	// IrCutFilterModeON const
	IrCutFilterModeON IrCutFilterMode = "ON"

	// IrCutFilterModeOFF const
	IrCutFilterModeOFF IrCutFilterMode = "OFF"

	// IrCutFilterModeAUTO const
	IrCutFilterModeAUTO IrCutFilterMode = "AUTO"
)

// ImageStabilizationMode type
type ImageStabilizationMode string

const (
	// ImageStabilizationModeOFF const
	ImageStabilizationModeOFF ImageStabilizationMode = "OFF"

	// ImageStabilizationModeON const
	ImageStabilizationModeON ImageStabilizationMode = "ON"

	// ImageStabilizationModeAUTO const
	ImageStabilizationModeAUTO ImageStabilizationMode = "AUTO"

	// ImageStabilizationModeExtended const
	ImageStabilizationModeExtended ImageStabilizationMode = "Extended"
)

// IrCutFilterAutoBoundaryType type
type IrCutFilterAutoBoundaryType string

const (
	// IrCutFilterAutoBoundaryTypeCommon const
	IrCutFilterAutoBoundaryTypeCommon IrCutFilterAutoBoundaryType = "Common"

	// IrCutFilterAutoBoundaryTypeToOn const
	IrCutFilterAutoBoundaryTypeToOn IrCutFilterAutoBoundaryType = "ToOn"

	// IrCutFilterAutoBoundaryTypeToOff const
	IrCutFilterAutoBoundaryTypeToOff IrCutFilterAutoBoundaryType = "ToOff"

	// IrCutFilterAutoBoundaryTypeExtended const
	IrCutFilterAutoBoundaryTypeExtended IrCutFilterAutoBoundaryType = "Extended"
)

// ToneCompensationMode type
type ToneCompensationMode string

const (
	// ToneCompensationModeOFF const
	ToneCompensationModeOFF ToneCompensationMode = "OFF"

	// ToneCompensationModeON const
	ToneCompensationModeON ToneCompensationMode = "ON"

	// ToneCompensationModeAUTO const
	ToneCompensationModeAUTO ToneCompensationMode = "AUTO"
)

// DefoggingMode type
type DefoggingMode string

const (
	// DefoggingModeOFF const
	DefoggingModeOFF DefoggingMode = "OFF"

	// DefoggingModeON const
	DefoggingModeON DefoggingMode = "ON"

	// DefoggingModeAUTO const
	DefoggingModeAUTO DefoggingMode = "AUTO"
)

// TopicNamespaceLocation type
type TopicNamespaceLocation AnyURI

// PropertyOperation type
type PropertyOperation string

const (
	// PropertyOperationInitialized const
	PropertyOperationInitialized PropertyOperation = "Initialized"

	// PropertyOperationDeleted const
	PropertyOperationDeleted PropertyOperation = "Deleted"

	// PropertyOperationChanged const
	PropertyOperationChanged PropertyOperation = "Changed"
)

// Direction type
type Direction string

const (
	// DirectionLeft const
	DirectionLeft Direction = "Left"

	// DirectionRight const
	DirectionRight Direction = "Right"

	// DirectionAny const
	DirectionAny Direction = "Any"
)

//
// Specifies a receiver connection mode.
//

// ReceiverMode type
type ReceiverMode string

const (

	// The receiver connects on demand, as required by consumers of the media streams.
	// ReceiverModeAutoConnect const
	ReceiverModeAutoConnect ReceiverMode = "AutoConnect"

	// The receiver attempts to maintain a persistent connection to the configured endpoint.
	// ReceiverModeAlwaysConnect const
	ReceiverModeAlwaysConnect ReceiverMode = "AlwaysConnect"

	// The receiver does not attempt to connect.
	// ReceiverModeNeverConnect const
	ReceiverModeNeverConnect ReceiverMode = "NeverConnect"

	// This case should never happen.
	// ReceiverModeUnknown const
	ReceiverModeUnknown ReceiverMode = "Unknown"
)

//
// Specifies the current connection state of the receiver.
//

// ReceiverState type
type ReceiverState string

const (

	// The receiver is not connected.
	// ReceiverStateNotConnected const
	ReceiverStateNotConnected ReceiverState = "NotConnected"

	// The receiver is attempting to connect.
	// ReceiverStateConnecting const
	ReceiverStateConnecting ReceiverState = "Connecting"

	// The receiver is connected.
	// ReceiverStateConnected const
	ReceiverStateConnected ReceiverState = "Connected"

	// This case should never happen.
	// ReceiverStateUnknown const
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
	// SearchStateQueued const
	SearchStateQueued SearchState = "Queued"

	// The search is underway and not yet completed.
	// SearchStateSearching const
	SearchStateSearching SearchState = "Searching"

	// The search has been completed and no new results will be found.
	// SearchStateCompleted const
	SearchStateCompleted SearchState = "Completed"

	// The state of the search is unknown. (This is not a valid response from GetSearchState.)
	// SearchStateUnknown const
	SearchStateUnknown SearchState = "Unknown"
)

// JobToken type
type JobToken ReferenceToken

// RecordingStatus type
type RecordingStatus string

const (
	// RecordingStatusInitiated const
	RecordingStatusInitiated RecordingStatus = "Initiated"

	// RecordingStatusRecording const
	RecordingStatusRecording RecordingStatus = "Recording"

	// RecordingStatusStopped const
	RecordingStatusStopped RecordingStatus = "Stopped"

	// RecordingStatusRemoving const
	RecordingStatusRemoving RecordingStatus = "Removing"

	// RecordingStatusRemoved const
	RecordingStatusRemoved RecordingStatus = "Removed"

	// This case should never happen.
	// RecordingStatusUnknown const
	RecordingStatusUnknown RecordingStatus = "Unknown"
)

// TrackType type
type TrackType string

const (
	// TrackTypeVideo const
	TrackTypeVideo TrackType = "Video"

	// TrackTypeAudio const
	TrackTypeAudio TrackType = "Audio"

	// TrackTypeMetadata const
	TrackTypeMetadata TrackType = "Metadata"

	// Placeholder for future extension.
	// TrackTypeExtended const
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
	// ModeOfOperationIdle const
	ModeOfOperationIdle ModeOfOperation = "Idle"

	// ModeOfOperationActive const
	ModeOfOperationActive ModeOfOperation = "Active"

	// This case should never happen.
	// ModeOfOperationUnknown const
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
	// OSDTypeText const
	OSDTypeText OSDType = "Text"

	// OSDTypeImage const
	OSDTypeImage OSDType = "Image"

	// OSDTypeExtended const
	OSDTypeExtended OSDType = "Extended"
)

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

// VideoSource type
type VideoSource struct {
	*DeviceEntity

	// Frame rate in frames per second.
	Framerate float32 `xml:"http://www.onvif.org/ver10/schema Framerate,omitempty"`

	// Horizontal and vertical resolution
	Resolution VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl Resolution,omitempty"`

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
	Name Name `xml:"http://www.onvif.org/ver10/media/wsdl Name,omitempty"`

	// Optional configuration of the Video input.
	VideoSourceConfiguration VideoSourceConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceConfiguration,omitempty"`

	// Optional configuration of the Audio input.
	AudioSourceConfiguration AudioSourceConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl AudioSourceConfiguration,omitempty"`

	// Optional configuration of the Video encoder.
	VideoEncoderConfiguration VideoEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl VideoEncoderConfiguration,omitempty"`

	// Optional configuration of the Audio encoder.
	AudioEncoderConfiguration AudioEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl AudioEncoderConfiguration,omitempty"`

	// Optional configuration of the video analytics module and rule engine.
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl VideoAnalyticsConfiguration,omitempty"`

	// Optional configuration of the pan tilt zoom unit.
	PTZConfiguration PTZConfiguration `xml:"http://www.onvif.org/ver10/schema PTZConfiguration,omitempty"`

	// Optional configuration of the metadata stream.
	MetadataConfiguration MetadataConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl MetadataConfiguration,omitempty"`

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
	AudioOutputConfiguration AudioOutputConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl AudioOutputConfiguration,omitempty"`

	// Optional configuration of the Audio decoder.
	AudioDecoderConfiguration AudioDecoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl AudioDecoderConfiguration,omitempty"`

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
	UseCount int32 `xml:"http://www.onvif.org/ver10/schema UseCount"`

	// Token that uniquely references this configuration. Length up to 64 characters.

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

// VideoSourceConfiguration type
type VideoSourceConfiguration struct {
	*ConfigurationEntity

	// Reference to the physical input.
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl SourceToken,omitempty"`

	// Rectangle specifying the Video capturing area. The capturing area shall not be larger than the whole Video source area.
	Bounds IntRectangle `xml:"http://www.onvif.org/ver10/schema Bounds,omitempty"`

	Extension VideoSourceConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Readonly parameter signalling Source configuration's view mode, for devices supporting different view modes as defined in tt:viewModes.

	ViewMode string `xml:"http://www.onvif.org/ver10/schema ViewMode,attr,omitempty"`
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
	VideoSourceTokensAvailable []ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl VideoSourceTokensAvailable,omitempty"`

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

	Reboot bool `xml:"http://www.onvif.org/ver10/media/wsdl Reboot,attr,omitempty"`
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
	Orientation string `xml:"http://www.onvif.org/ver10/schema Orientation,omitempty"`
}

// VideoEncoderConfiguration type
type VideoEncoderConfiguration struct {
	*ConfigurationEntity

	// Used video codec, either Jpeg, H.264 or Mpeg4
	Encoding VideoEncoding `xml:"http://www.onvif.org/ver10/schema Encoding,omitempty"`

	// Configured video resolution
	Resolution VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl Resolution,omitempty"`

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

	GuaranteedFrameRate bool `xml:"http://www.onvif.org/ver10/media/wsdl GuaranteedFrameRate,attr,omitempty"`
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

// VideoEncoderConfigurationOptions type
type VideoEncoderConfigurationOptions struct {

	// Range of the quality values. A high value means higher quality.
	QualityRange IntRange `xml:"http://www.onvif.org/ver10/schema QualityRange,omitempty"`

	// Optional JPEG encoder settings ranges (See also Extension element).
	JPEG JpegOptions `xml:"http://www.onvif.org/ver10/schema JPEG,omitempty"`

	// Optional MPEG-4 encoder settings ranges (See also Extension element).
	MPEG4 Mpeg4Options `xml:"http://www.onvif.org/ver10/schema MPEG4,omitempty"`

	// Optional H.264 encoder settings ranges (See also Extension element).
	H264 H264Options `xml:"http://www.onvif.org/ver10/schema H264,omitempty"`

	Extension VideoEncoderOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	//
	// Indicates the support for the GuaranteedFrameRate attribute on the VideoEncoderConfiguration element.
	//

	GuaranteedFrameRateSupported bool `xml:"http://www.onvif.org/ver10/media/wsdl GuaranteedFrameRateSupported,attr,omitempty"`
}

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
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl ResolutionsAvailable,omitempty"`

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
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl ResolutionsAvailable,omitempty"`

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
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl ResolutionsAvailable,omitempty"`

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

// AudioSourceConfiguration type
type AudioSourceConfiguration struct {
	*ConfigurationEntity

	// Token of the Audio Source the configuration applies to
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl SourceToken,omitempty"`
}

// AudioSourceConfigurationOptions type
type AudioSourceConfigurationOptions struct {

	// Tokens of the audio source the configuration can be used for.
	InputTokensAvailable []ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl InputTokensAvailable,omitempty"`

	Extension AudioSourceOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AudioSourceOptionsExtension type
type AudioSourceOptionsExtension struct {
}

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
	Options []AudioEncoderConfigurationOption `xml:"http://www.onvif.org/ver10/media/wsdl Options,omitempty"`
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
//	PTZStatus PTZFilter `xml:"http://www.onvif.org/ver10/schema PTZStatus,omitempty"`

	// Optional element to configure the streaming of events. A client might be interested in receiving all,
	// none or some of the events produced by the device:
	//
	Events EventSubscription `xml:"http://www.onvif.org/ver10/schema Events,omitempty"`

	// Defines whether the streamed metadata will include metadata from the analytics engines (video, cell motion, audio etc.)
	Analytics bool `xml:"http://www.onvif.org/ver10/schema Analytics,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast,omitempty"`

	// The rtsp session timeout for the related audio stream (when using Media2 Service, this value is deprecated and ignored)
	SessionTimeout Duration `xml:"http://www.onvif.org/ver10/schema SessionTimeout,omitempty"`

	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration,omitempty"`

	Extension MetadataConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Optional parameter to configure compression type of Metadata payload. Use values from enumeration MetadataCompressionType.

	CompressionType string `xml:"CompressionType,attr,omitempty"`

	// Optional parameter to configure if the metadata stream shall contain the Geo Location coordinates of each target.

	GeoLocation bool `xml:"GeoLocation,attr"`

	// Optional parameter to configure if the generated metadata stream should contain shape information as polygon.

	ShapePolygon bool `xml:"ShapePolygon,attr"`
}

// MetadataConfigurationExtension type
type MetadataConfigurationExtension struct {
}

// PTZFilter type
type PTZFilter struct {

	// True if the metadata stream shall contain the PTZ status (IDLE, MOVING or UNKNOWN)
	Status bool `xml:"http://www.onvif.org/ver10/media/wsdl Status,omitempty"`

	// True if the metadata stream shall contain the PTZ position
	Position bool `xml:"http://www.onvif.org/ver10/media/wsdl Position,omitempty"`
}

// EventSubscription type
type EventSubscription struct {
	Filter FilterType `xml:"http://www.onvif.org/ver10/schema Filter,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
}

// MetadataConfigurationOptions type
type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions `xml:"http://www.onvif.org/ver10/schema PTZStatusFilterOptions,omitempty"`

	Extension MetadataConfigurationOptionsExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// True if the device is able to stream the Geo Located positions of each target.

	GeoLocation bool `xml:"http://www.onvif.org/ver10/media/wsdl GeoLocation,attr,omitempty"`
}

// MetadataConfigurationOptionsExtension type
type MetadataConfigurationOptionsExtension struct {

	// List of supported metadata compression type. Its options shall be chosen from tt:MetadataCompressionType.
	CompressionType []string `xml:"http://www.onvif.org/ver10/schema CompressionType,omitempty"`

	Extension MetadataConfigurationOptionsExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// MetadataConfigurationOptionsExtension2 type
type MetadataConfigurationOptionsExtension2 struct {
}

// PTZStatusFilterOptions type
type PTZStatusFilterOptions struct {

	// True if the device is able to stream pan or tilt status information.
	PanTiltStatusSupported bool `xml:"http://www.onvif.org/ver10/media/wsdl PanTiltStatusSupported,omitempty"`

	// True if the device is able to stream zoom status inforamtion.
	ZoomStatusSupported bool `xml:"http://www.onvif.org/ver10/media/wsdl ZoomStatusSupported,omitempty"`

	// True if the device is able to stream the pan or tilt position.
	PanTiltPositionSupported bool `xml:"http://www.onvif.org/ver10/media/wsdl PanTiltPositionSupported,omitempty"`

	// True if the device is able to stream zoom position information.
	ZoomPositionSupported bool `xml:"http://www.onvif.org/ver10/media/wsdl ZoomPositionSupported,omitempty"`

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
	Resolution VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl Resolution,omitempty"`

	// Refresh rate of the display in Hertz.
	RefreshRate float32 `xml:"http://www.onvif.org/ver10/schema RefreshRate,omitempty"`

	// Aspect ratio of the display as physical extent of width divided by height.
	AspectRatio float32 `xml:"http://www.onvif.org/ver10/schema AspectRatio,omitempty"`

	Extension VideoOutputExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// VideoOutputExtension type
type VideoOutputExtension struct {
}

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
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl ResolutionsAvailable,omitempty"`

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
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl ResolutionsAvailable,omitempty"`

	// Supported Jpeg bitrate range in kbps
	SupportedInputBitrate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedInputBitrate,omitempty"`

	// Supported Jpeg framerate range in fps
	SupportedFrameRate IntRange `xml:"http://www.onvif.org/ver10/schema SupportedFrameRate,omitempty"`
}

// Mpeg4DecOptions type
type Mpeg4DecOptions struct {

	// List of supported Mpeg4 Video Resolutions
	ResolutionsAvailable []VideoResolution `xml:"http://www.onvif.org/ver10/media/wsdl ResolutionsAvailable,omitempty"`

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
	OutputToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl OutputToken,omitempty"`

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

// AudioOutputConfigurationOptions type
type AudioOutputConfigurationOptions struct {

	// Tokens of the physical Audio outputs (typically one).
	OutputTokensAvailable []ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl OutputTokensAvailable,omitempty"`

	//
	// An  channel MAY support different types of audio transmission. While for full duplex
	// operation no special handling is required, in half duplex operation the transmission direction
	// needs to be switched.
	// The optional SendPrimacy parameter inside the AudioOutputConfiguration indicates which
	// direction is currently active. An NVC can switch between different modes by setting the
	// AudioOutputConfiguration.
	// The following modes for the Send-Primacy are defined:
	// Acoustic echo cancellation is out of ONVIF scope.
	SendPrimacyOptions []AnyURI `xml:"http://www.onvif.org/ver10/schema SendPrimacyOptions,omitempty"`

	// Minimum and maximum level range supported for this Output.
	OutputLevelRange IntRange `xml:"http://www.onvif.org/ver10/schema OutputLevelRange,omitempty"`
}

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
	Address IPAddress `xml:"http://www.onvif.org/ver10/schema Address,omitempty"`

	// The RTP mutlicast destination port. A device may support RTCP. In this case the port value shall be even to allow the corresponding RTCP stream to be mapped to the next higher (odd) destination port number as defined in the RTSP specification.
	Port int32 `xml:"http://www.onvif.org/ver10/schema Port,omitempty"`

	// In case of IPv6 the TTL value is assumed as the hop limit. Note that for IPV6 and administratively scoped IPv4 multicast the primary use for hop limit / TTL is to prevent packets from (endlessly) circulating and not limiting scope. In these cases the address contains the scope.
	TTL int32 `xml:"http://www.onvif.org/ver10/schema TTL,omitempty"`

	// Read only property signalling that streaming is persistant. Use the methods StartMulticastStreaming and StopMulticastStreaming to switch its state.
	AutoStart bool `xml:"http://www.onvif.org/ver10/schema AutoStart"`
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

// MediaUri type
type MediaUri struct {

	// Stable Uri to be used for requesting the media stream
	Uri AnyURI `xml:"http://www.onvif.org/ver10/schema Uri,omitempty"`

	// Indicates if the Uri is only valid until the connection is established. The value shall be set to "false".
	InvalidAfterConnect bool `xml:"http://www.onvif.org/ver10/media/wsdl InvalidAfterConnect,omitempty"`

	// Indicates if the Uri is invalid after a reboot of the device. The value shall be set to "false".
	InvalidAfterReboot bool `xml:"http://www.onvif.org/ver10/media/wsdl InvalidAfterReboot,omitempty"`

	// Duration how long the Uri is valid. This parameter shall be set to PT0S to indicate that this stream URI is indefinitely valid even if the profile changes
	Timeout Duration `xml:"http://www.onvif.org/ver10/schema Timeout,omitempty"`
}

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
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

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
	AutoNegotiation bool `xml:"http://www.onvif.org/ver10/media/wsdl AutoNegotiation,omitempty"`

	// Speed.
	Speed int32 `xml:"http://www.onvif.org/ver10/schema Speed,omitempty"`

	// Duplex type, Half or Full.
	Duplex Duplex `xml:"http://www.onvif.org/ver10/schema Duplex,omitempty"`
}

// NetworkInterfaceInfo type
type NetworkInterfaceInfo struct {

	// Network interface name, for example eth0.
	Name string `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	// Network interface MAC address.
	HwAddress HwAddress `xml:"http://www.onvif.org/ver10/schema HwAddress,omitempty"`

	// Maximum transmission unit.
	MTU int32 `xml:"http://www.onvif.org/ver10/schema MTU,omitempty"`
}

// IPv6NetworkInterface type
type IPv6NetworkInterface struct {

	// Indicates whether or not IPv6 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

	// IPv6 configuration.
	Config IPv6Configuration `xml:"http://www.onvif.org/ver10/schema Config,omitempty"`
}

// IPv4NetworkInterface type
type IPv4NetworkInterface struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

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
	DHCP bool `xml:"http://www.onvif.org/ver10/media/wsdl DHCP,omitempty"`
}

// IPv6Configuration type
type IPv6Configuration struct {

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"http://www.onvif.org/ver10/media/wsdl AcceptRouterAdvert,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

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
	IPv4Address IPv4Address `xml:"http://www.onvif.org/ver10/schema IPv4Address,omitempty"`

	// IPv6 address.
	IPv6Address IPv6Address `xml:"http://www.onvif.org/ver10/schema IPv6Address,omitempty"`

	// DNS name.
	DNSname DNSName `xml:"http://www.onvif.org/ver10/schema DNSname,omitempty"`

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
	IPv4Address IPv4Address `xml:"http://www.onvif.org/ver10/schema IPv4Address,omitempty"`

	// IPv6 address
	IPv6Address IPv6Address `xml:"http://www.onvif.org/ver10/schema IPv6Address,omitempty"`
}

// PrefixedIPv4Address type
type PrefixedIPv4Address struct {

	// IPv4 address
	Address IPv4Address `xml:"http://www.onvif.org/ver10/schema Address,omitempty"`

	// Prefix/submask length
	PrefixLength int32 `xml:"http://www.onvif.org/ver10/schema PrefixLength,omitempty"`
}

// PrefixedIPv6Address type
type PrefixedIPv6Address struct {

	// IPv6 address
	Address IPv6Address `xml:"http://www.onvif.org/ver10/schema Address,omitempty"`

	// Prefix/submask length
	PrefixLength int32 `xml:"http://www.onvif.org/ver10/schema PrefixLength,omitempty"`
}

// HostnameInformation type
type HostnameInformation struct {

	// Indicates whether the hostname is obtained from DHCP or not.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/media/wsdl FromDHCP,omitempty"`

	// Indicates the hostname.
	Name string `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	Extension HostnameInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// HostnameInformationExtension type
type HostnameInformationExtension struct {
}

// DNSInformation type
type DNSInformation struct {

	// Indicates whether or not DNS information is retrieved from DHCP.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/media/wsdl FromDHCP,omitempty"`

	// Search domain.
	SearchDomain []string `xml:"http://www.onvif.org/ver10/schema SearchDomain,omitempty"`

	// List of DNS addresses received from DHCP.
	DNSFromDHCP []IPAddress `xml:"http://www.onvif.org/ver10/schema DNSFromDHCP,omitempty"`

	// List of manually entered DNS addresses.
	DNSManual []IPAddress `xml:"http://www.onvif.org/ver10/schema DNSManual,omitempty"`

	Extension DNSInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// DNSInformationExtension type
type DNSInformationExtension struct {
}

// NTPInformation type
type NTPInformation struct {

	// Indicates if NTP information is to be retrieved by using DHCP.
	FromDHCP bool `xml:"http://www.onvif.org/ver10/media/wsdl FromDHCP,omitempty"`

	// List of NTP addresses retrieved by using DHCP.
	NTPFromDHCP []NetworkHost `xml:"http://www.onvif.org/ver10/schema NTPFromDHCP,omitempty"`

	// List of manually entered NTP addresses.
	NTPManual []NetworkHost `xml:"http://www.onvif.org/ver10/schema NTPManual,omitempty"`

	Extension NTPInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NTPInformationExtension type
type NTPInformationExtension struct {
}

// DynamicDNSInformation type
type DynamicDNSInformation struct {

	// Dynamic DNS type.
	Type DynamicDNSType `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// DNS name.
	Name DNSName `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"http://www.onvif.org/ver10/media/wsdl AcceptRouterAdvert,omitempty"`

	// List of manually added IPv6 addresses.
	Manual []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// DHCP configuration.
	DHCP IPv6DHCPConfiguration `xml:"http://www.onvif.org/ver10/schema DHCP,omitempty"`
}

// IPv4NetworkInterfaceSetConfiguration type
type IPv4NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

	// List of manually added IPv4 addresses.
	Manual []PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// Indicates whether or not DHCP is used.
	DHCP bool `xml:"http://www.onvif.org/ver10/media/wsdl DHCP,omitempty"`
}

// NetworkZeroConfiguration type
type NetworkZeroConfiguration struct {

	// Unique identifier of network interface.
	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl InterfaceToken,omitempty"`

	// Indicates whether the zero-configuration is enabled or not.
	Enabled bool `xml:"http://www.onvif.org/ver10/media/wsdl Enabled,omitempty"`

	// The zero-configuration IPv4 address(es)
	Addresses []IPv4Address `xml:"http://www.onvif.org/ver10/schema Addresses,omitempty"`

	Extension *NetworkZeroConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkZeroConfigurationExtension type
type NetworkZeroConfigurationExtension struct {

	// Optional array holding the configuration for the second and possibly further interfaces.
	Additional []NetworkZeroConfiguration `xml:"http://www.onvif.org/ver10/schema Additional,omitempty"`

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

	Alias Name `xml:"http://www.onvif.org/ver10/media/wsdl Alias,omitempty"`

	Priority NetworkInterfaceConfigPriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	Security Dot11SecurityConfiguration `xml:"http://www.onvif.org/ver10/schema Security,omitempty"`
}

// Dot11SecurityConfiguration type
type Dot11SecurityConfiguration struct {
	Mode Dot11SecurityMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Algorithm Dot11Cipher `xml:"http://www.onvif.org/ver10/schema Algorithm,omitempty"`

	PSK Dot11PSKSet `xml:"http://www.onvif.org/ver10/schema PSK,omitempty"`

	Dot1X ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Dot1X,omitempty"`

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

// Dot11AvailableNetworks type
type Dot11AvailableNetworks struct {
	SSID Dot11SSIDType `xml:"http://www.onvif.org/ver10/schema SSID,omitempty"`

	BSSID string `xml:"http://www.onvif.org/ver10/schema BSSID,omitempty"`

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

/* Removed type Capabilities struct {

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
} Removed*/

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
	RuleSupport bool `xml:"http://www.onvif.org/ver10/media/wsdl RuleSupport,omitempty"`

	// Indicates whether or not modules are supported.
	AnalyticsModuleSupport bool `xml:"http://www.onvif.org/ver10/media/wsdl AnalyticsModuleSupport,omitempty"`
}

// DeviceCapabilities type
type DeviceCapabilities struct {

	// Device service URI.
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Network capabilities.
	Network NetworkCapabilities `xml:"http://www.onvif.org/ver10/schema Network,omitempty"`

	// System capabilities.
	System SystemCapabilities `xml:"http://www.onvif.org/ver10/schema System,omitempty"`

	// I/O capabilities.
	IO IOCapabilities `xml:"http://www.onvif.org/ver10/schema IO,omitempty"`

	// Security capabilities.
	Security SecurityCapabilities `xml:"http://www.onvif.org/ver10/schema Security,omitempty"`

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
	WSSubscriptionPolicySupport bool `xml:"http://www.onvif.org/ver10/media/wsdl WSSubscriptionPolicySupport,omitempty"`

	// Indicates whether or not WS Pull Point is supported.
	WSPullPointSupport bool `xml:"http://www.onvif.org/ver10/media/wsdl WSPullPointSupport,omitempty"`

	// Indicates whether or not WS Pausable Subscription Manager Interface is supported.
	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"http://www.onvif.org/ver10/media/wsdl WSPausableSubscriptionManagerInterfaceSupport,omitempty"`
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
	Auxiliary bool `xml:"http://www.onvif.org/ver10/media/wsdl Auxiliary,omitempty"`

	AuxiliaryCommands []AuxiliaryData `xml:"http://www.onvif.org/ver10/schema AuxiliaryCommands,omitempty"`

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
	ProfileCapabilities ProfileCapabilities `xml:"http://www.onvif.org/ver10/media/wsdl ProfileCapabilities,omitempty"`
}

// RealTimeStreamingCapabilities type
type RealTimeStreamingCapabilities struct {

	// Indicates whether or not RTP multicast is supported.
	RTPMulticast bool `xml:"http://www.onvif.org/ver10/media/wsdl RTPMulticast,omitempty"`

	// Indicates whether or not RTP over TCP is supported.
	RTP_TCP bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_TCP,omitempty"`

	// Indicates whether or not RTP/RTSP/TCP is supported.
	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_RTSP_TCP,omitempty"`

	Extension RealTimeStreamingCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// RealTimeStreamingCapabilitiesExtension type
type RealTimeStreamingCapabilitiesExtension struct {
}

// NetworkCapabilities type
type NetworkCapabilities struct {

	// Indicates whether or not IP filtering is supported.
	IPFilter bool `xml:"http://www.onvif.org/ver10/media/wsdl IPFilter,omitempty"`

	// Indicates whether or not zeroconf is supported.
	ZeroConfiguration bool `xml:"http://www.onvif.org/ver10/media/wsdl ZeroConfiguration,omitempty"`

	// Indicates whether or not IPv6 is supported.
	IPVersion6 bool `xml:"http://www.onvif.org/ver10/media/wsdl IPVersion6,omitempty"`

	// Indicates whether or not  is supported.
	DynDNS bool `xml:"http://www.onvif.org/ver10/media/wsdl DynDNS,omitempty"`

	Extension NetworkCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkCapabilitiesExtension type
type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool `xml:"http://www.onvif.org/ver10/media/wsdl Dot11Configuration,omitempty"`

	Extension NetworkCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkCapabilitiesExtension2 type
type NetworkCapabilitiesExtension2 struct {
}

// SecurityCapabilities type
type SecurityCapabilities struct {

	// Indicates whether or not TLS 1.1 is supported.
	TLS11 bool `xml:"TLS1.1,omitempty"`

	// Indicates whether or not TLS 1.2 is supported.
	TLS12 bool `xml:"TLS1.2,omitempty"`

	// Indicates whether or not onboard key generation is supported.
	OnboardKeyGeneration bool `xml:"http://www.onvif.org/ver10/media/wsdl OnboardKeyGeneration,omitempty"`

	// Indicates whether or not access policy configuration is supported.
	AccessPolicyConfig bool `xml:"http://www.onvif.org/ver10/media/wsdl AccessPolicyConfig,omitempty"`

	// Indicates whether or not WS-Security X.509 token is supported.
	X509Token bool `xml:"X.509Token,omitempty"`

	// Indicates whether or not WS-Security SAML token is supported.
	SAMLToken bool `xml:"http://www.onvif.org/ver10/media/wsdl SAMLToken,omitempty"`

	// Indicates whether or not WS-Security Kerberos token is supported.
	KerberosToken bool `xml:"http://www.onvif.org/ver10/media/wsdl KerberosToken,omitempty"`

	// Indicates whether or not WS-Security REL token is supported.
	RELToken bool `xml:"http://www.onvif.org/ver10/media/wsdl RELToken,omitempty"`

	Extension SecurityCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SecurityCapabilitiesExtension type
type SecurityCapabilitiesExtension struct {
	TLS10 bool `xml:"TLS1.0,omitempty"`

	Extension SecurityCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SecurityCapabilitiesExtension2 type
type SecurityCapabilitiesExtension2 struct {
	Dot1X bool `xml:"http://www.onvif.org/ver10/media/wsdl Dot1X,omitempty"`

	// EAP Methods supported by the device. The int values refer to the .
	SupportedEAPMethod []int32 `xml:"http://www.onvif.org/ver10/schema SupportedEAPMethod,omitempty"`

	RemoteUserHandling bool `xml:"http://www.onvif.org/ver10/media/wsdl RemoteUserHandling,omitempty"`
}

// SystemCapabilities type
type SystemCapabilities struct {

	// Indicates whether or not WS Discovery resolve requests are supported.
	DiscoveryResolve bool `xml:"http://www.onvif.org/ver10/media/wsdl DiscoveryResolve,omitempty"`

	// Indicates whether or not WS-Discovery Bye is supported.
	DiscoveryBye bool `xml:"http://www.onvif.org/ver10/media/wsdl DiscoveryBye,omitempty"`

	// Indicates whether or not remote discovery is supported.
	RemoteDiscovery bool `xml:"http://www.onvif.org/ver10/media/wsdl RemoteDiscovery,omitempty"`

	// Indicates whether or not system backup is supported.
	SystemBackup bool `xml:"http://www.onvif.org/ver10/media/wsdl SystemBackup,omitempty"`

	// Indicates whether or not system logging is supported.
	SystemLogging bool `xml:"http://www.onvif.org/ver10/media/wsdl SystemLogging,omitempty"`

	// Indicates whether or not firmware upgrade is supported.
	FirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/media/wsdl FirmwareUpgrade,omitempty"`

	// Indicates supported ONVIF version(s).
	SupportedVersions []OnvifVersion `xml:"http://www.onvif.org/ver10/schema SupportedVersions,omitempty"`

	Extension SystemCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SystemCapabilitiesExtension type
type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/media/wsdl HttpFirmwareUpgrade,omitempty"`

	HttpSystemBackup bool `xml:"http://www.onvif.org/ver10/media/wsdl HttpSystemBackup,omitempty"`

	HttpSystemLogging bool `xml:"http://www.onvif.org/ver10/media/wsdl HttpSystemLogging,omitempty"`

	HttpSupportInformation bool `xml:"http://www.onvif.org/ver10/media/wsdl HttpSupportInformation,omitempty"`

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
	FixedLayout bool `xml:"http://www.onvif.org/ver10/media/wsdl FixedLayout,omitempty"`
}

// RecordingCapabilities type
type RecordingCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	ReceiverSource bool `xml:"http://www.onvif.org/ver10/media/wsdl ReceiverSource,omitempty"`

	MediaProfileSource bool `xml:"http://www.onvif.org/ver10/media/wsdl MediaProfileSource,omitempty"`

	DynamicRecordings bool `xml:"http://www.onvif.org/ver10/media/wsdl DynamicRecordings,omitempty"`

	DynamicTracks bool `xml:"http://www.onvif.org/ver10/media/wsdl DynamicTracks,omitempty"`

	MaxStringLength int32 `xml:"http://www.onvif.org/ver10/schema MaxStringLength,omitempty"`
}

// SearchCapabilities type
type SearchCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	MetadataSearch bool `xml:"http://www.onvif.org/ver10/media/wsdl MetadataSearch,omitempty"`
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
	RTP_Multicast bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_Multicast,omitempty"`

	// Indicates whether the device can receive RTP/TCP streams
	RTP_TCP bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_TCP,omitempty"`

	// Indicates whether the device can receive RTP/RTSP/TCP streams.
	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/media/wsdl RTP_RTSP_TCP,omitempty"`

	// The maximum number of receivers supported by the device.
	SupportedReceivers int32 `xml:"http://www.onvif.org/ver10/schema SupportedReceivers,omitempty"`

	// The maximum allowed length for RTSP URIs.
	MaximumRTSPURILength int32 `xml:"http://www.onvif.org/ver10/schema MaximumRTSPURILength,omitempty"`
}

// AnalyticsDeviceCapabilities type
type AnalyticsDeviceCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Obsolete property.
	RuleSupport bool `xml:"http://www.onvif.org/ver10/media/wsdl RuleSupport,omitempty"`

	Extension AnalyticsDeviceExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AnalyticsDeviceExtension type
type AnalyticsDeviceExtension struct {
}

// BinaryData type
type BinaryData struct {

	// base64 encoded binary data.
	Data []byte `xml:"http://www.onvif.org/ver10/schema Data,omitempty"`

	ContentType string `xml:"contentType,attr,omitempty"`
}

// SystemDateTime type
type SystemDateTime struct {

	// Indicates if the time is set manully or through NTP.
	DateTimeType SetDateTimeType `xml:"http://www.onvif.org/ver10/schema DateTimeType,omitempty"`

	// Informative indicator whether daylight savings is currently on/off.
	DaylightSavings bool `xml:"http://www.onvif.org/ver10/media/wsdl DaylightSavings,omitempty"`

	// Timezone information in Posix format.
	TimeZone TimeZone `xml:"http://www.onvif.org/ver10/schema TimeZone,omitempty"`

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
	TZ string `xml:"http://www.onvif.org/ver10/schema TZ,omitempty"`
}

// User type
type User struct {

	// Username string.
	Username string `xml:"http://www.onvif.org/ver10/schema Username,omitempty"`

	// Password string.
	Password string `xml:"http://www.onvif.org/ver10/schema Password,omitempty"`

	// User level string.
	UserLevel UserLevel `xml:"http://www.onvif.org/ver10/schema UserLevel,omitempty"`

	Extension UserExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// UserExtension type
type UserExtension struct {
}

// CertificateGenerationParameters type
type CertificateGenerationParameters struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/schema CertificateID,omitempty"`

	Subject string `xml:"http://www.onvif.org/ver10/schema Subject,omitempty"`

	ValidNotBefore string `xml:"http://www.onvif.org/ver10/schema ValidNotBefore,omitempty"`

	ValidNotAfter string `xml:"http://www.onvif.org/ver10/schema ValidNotAfter,omitempty"`

	Extension CertificateGenerationParametersExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// CertificateGenerationParametersExtension type
type CertificateGenerationParametersExtension struct {
}

// Certificate type
type Certificate struct {

	// Certificate id.
	CertificateID string `xml:"http://www.onvif.org/ver10/schema CertificateID,omitempty"`

	// base64 encoded DER representation of certificate.
	Certificate BinaryData `xml:"http://www.onvif.org/ver10/schema Certificate,omitempty"`
}

// CertificateInformation type
type CertificateInformation struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/schema CertificateID,omitempty"`

	IssuerDN string `xml:"http://www.onvif.org/ver10/schema IssuerDN,omitempty"`

	SubjectDN string `xml:"http://www.onvif.org/ver10/schema SubjectDN,omitempty"`

	KeyUsage CertificateUsage `xml:"http://www.onvif.org/ver10/schema KeyUsage,omitempty"`

	ExtendedKeyUsage CertificateUsage `xml:"http://www.onvif.org/ver10/schema ExtendedKeyUsage,omitempty"`

	KeyLength int32 `xml:"http://www.onvif.org/ver10/schema KeyLength,omitempty"`

	Version string `xml:"http://www.onvif.org/ver10/schema Version,omitempty"`

	SerialNum string `xml:"http://www.onvif.org/ver10/schema SerialNum,omitempty"`

	// Validity Range is from "NotBefore" to "NotAfter"; the corresponding DateTimeRange is from "From" to "Until"
	SignatureAlgorithm string `xml:"http://www.onvif.org/ver10/schema SignatureAlgorithm,omitempty"`

	Validity DateTimeRange `xml:"http://www.onvif.org/ver10/schema Validity,omitempty"`

	Extension CertificateInformationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// CertificateUsage type
type CertificateUsage struct {
	Value string

	Critical bool `xml:"http://www.onvif.org/ver10/media/wsdl Critical,attr,omitempty"`
}

// CertificateInformationExtension type
type CertificateInformationExtension struct {
}

// Dot1XConfiguration type
type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Dot1XConfigurationToken,omitempty"`

	Identity string `xml:"http://www.onvif.org/ver10/schema Identity,omitempty"`

	AnonymousID string `xml:"http://www.onvif.org/ver10/schema AnonymousID,omitempty"`

	//
	// EAP Method type as defined in .
	//
	EAPMethod int32 `xml:"http://www.onvif.org/ver10/schema EAPMethod,omitempty"`

	CACertificateID []string `xml:"http://www.onvif.org/ver10/schema CACertificateID,omitempty"`

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
	Password string `xml:"http://www.onvif.org/ver10/schema Password,omitempty"`

	Extension EapMethodExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// EapMethodExtension type
type EapMethodExtension struct {
}

// TLSConfiguration type
type TLSConfiguration struct {
	CertificateID string `xml:"http://www.onvif.org/ver10/schema CertificateID,omitempty"`
}

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

	Properties RelayOutputSettings `xml:"http://www.onvif.org/ver10/schema Properties,omitempty"`
}

// PTZNode type
type PTZNode struct {
	*DeviceEntity

	//
	// A unique identifier that is used to reference PTZ Nodes.
	//
	Name Name `xml:"http://www.onvif.org/ver10/media/wsdl Name,omitempty"`

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
	HomeSupported bool `xml:"http://www.onvif.org/ver10/media/wsdl HomeSupported,omitempty"`

	//
	// A list of supported Auxiliary commands. If the list is not empty, the Auxiliary Operations MUST be available for this PTZ Node.
	//
	AuxiliaryCommands []AuxiliaryData `xml:"http://www.onvif.org/ver10/schema AuxiliaryCommands,omitempty"`

	Extension PTZNodeExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	//
	// Indication whether the HomePosition of a Node is fixed or it can be changed via the SetHomePosition command.
	//

	FixedHomePosition bool `xml:"http://www.onvif.org/ver10/media/wsdl FixedHomePosition,attr,omitempty"`

	//
	// Indication whether the Node supports the geo-referenced move command.
	//

	GeoMove bool `xml:"http://www.onvif.org/ver10/media/wsdl GeoMove,attr,omitempty"`
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
	NodeToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl NodeToken,omitempty"`

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
	Name Name `xml:"http://www.onvif.org/ver10/media/wsdl Name,omitempty"`

	//
	// A list of preset position.
	//
	PTZPosition PTZVector `xml:"http://www.onvif.org/ver10/schema PTZPosition,omitempty"`

	Token ReferenceToken `xml:"token,attr,omitempty"`
}

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
	PresetToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl PresetToken,omitempty"`

	// Option to specify the preset position with the home position of this PTZ Node. "False" to this parameter shall be treated as an invalid argument.
	Home bool `xml:"http://www.onvif.org/ver10/media/wsdl Home,omitempty"`

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

	RandomPresetOrder bool `xml:"http://www.onvif.org/ver10/media/wsdl RandomPresetOrder,attr,omitempty"`
}

// PTZPresetTourStartingConditionExtension type
type PTZPresetTourStartingConditionExtension struct {
}

// PTZPresetTourPresetDetailOptions type
type PTZPresetTourPresetDetailOptions struct {

	// A list of available Preset Tokens for tour spots.
	PresetToken []ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl PresetToken,omitempty"`

	// An option to indicate Home postion for tour spots.
	Home bool `xml:"http://www.onvif.org/ver10/media/wsdl Home,omitempty"`

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
	Error string `xml:"http://www.onvif.org/ver10/schema Error,omitempty"`
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
	Error string `xml:"http://www.onvif.org/ver10/schema Error,omitempty"`

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
	BoundaryType string `xml:"http://www.onvif.org/ver10/schema BoundaryType,omitempty"`

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
	Mode string `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

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
	Mode string `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

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
	BoundaryType []string `xml:"http://www.onvif.org/ver10/schema BoundaryType,omitempty"`

	// Indicates whether or not boundary offset for toggling Ir cut filter is supported.
	BoundaryOffset bool `xml:"http://www.onvif.org/ver10/media/wsdl BoundaryOffset,omitempty"`

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

	AFMode StringAttrList `xml:"http://www.onvif.org/ver10/schema AFMode,attr,omitempty"`
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
	AFModes StringAttrList `xml:"http://www.onvif.org/ver10/schema AFModes,omitempty"`
}

// ToneCompensationOptions type
type ToneCompensationOptions struct {

	// Supported options for Tone Compensation mode. Its options shall be chosen from tt:ToneCompensationMode Type.
	Mode []string `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Indicates whether or not support Level parameter for Tone Compensation.
	Level bool `xml:"http://www.onvif.org/ver10/media/wsdl Level,omitempty"`
}

// DefoggingOptions type
type DefoggingOptions struct {

	// Supported options for Defogging mode. Its options shall be chosen from tt:DefoggingMode Type.
	Mode []string `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Indicates whether or not support Level parameter for Defogging.
	Level bool `xml:"http://www.onvif.org/ver10/media/wsdl Level,omitempty"`
}

// NoiseReductionOptions type
type NoiseReductionOptions struct {

	// Indicates whether or not support Level parameter for NoiseReduction.
	Level bool `xml:"http://www.onvif.org/ver10/media/wsdl Level,omitempty"`
}

// MessageExtension type
type MessageExtension struct {
}

// ItemList type
type ItemList struct {
	SimpleItem []SimpleItem `xml:"http://www.onvif.org/ver10/schema SimpleItem,omitempty"`

	ElementItem []ElementItem `xml:"http://www.onvif.org/ver10/schema ElementItem,omitempty"`

	Extension ItemListExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

type SimpleItem struct {

	// Item name.

	Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

	// Item value. The type is defined in the corresponding description.

	Value AnySimpleType `xml:"http://www.onvif.org/ver10/schema Value,attr,omitempty"`
}

type ElementItem struct {

	// Item name.

	Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

	Polygon []Polyline `xml:"http://www.onvif.org/ver10/schema Polygon,omitempty"`

	Polyline []Polyline `xml:"http://www.onvif.org/ver10/schema Polyline,omitempty"`
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

	IsProperty bool `xml:"http://www.onvif.org/ver10/media/wsdl IsProperty,attr,omitempty"`
}

// MessageDescriptionExtension type
type MessageDescriptionExtension struct {
}

// ItemListDescription type
type ItemListDescription struct {
	SimpleItemDescription []struct {

		// Item name. Must be unique within a list.

		Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

		Type QName `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
	} `xml:"SimpleItemDescription,omitempty"`

	ElementItemDescription []struct {

		// Item name. Must be unique within a list.

		Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

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

	Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

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
		ParentTopic string `xml:"http://www.onvif.org/ver10/schema ParentTopic,omitempty"`
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

// PolylineArray type
type PolylineArray struct {

	// Contains array of Polyline
	Segment []Polyline `xml:"http://www.onvif.org/ver10/schema Segment,omitempty"`

	Extension PolylineArrayExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// PolylineArrayExtension type
type PolylineArrayExtension struct {
}

// PaneLayout type
type PaneLayout struct {

	// Reference to the configuration of the streaming and coding parameters.
	Pane ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Pane,omitempty"`

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
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Token,omitempty"`

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
	StreamSetup StreamSetup `xml:"http://www.onvif.org/ver10/media/wsdl StreamSetup,omitempty"`
}

// SourceReference type
type SourceReference struct {
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Token,omitempty"`

	Type AnyURI `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
}

// DateTimeRange type
type DateTimeRange struct {
	From string `xml:"http://www.onvif.org/ver10/schema From,omitempty"`

	Until string `xml:"http://www.onvif.org/ver10/schema Until,omitempty"`
}

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

// RecordingSourceInformation type
type RecordingSourceInformation struct {

	//
	// Identifier for the source chosen by the client that creates the structure.
	// This identifier is opaque to the device. Clients may use any type of URI for this field. A device shall support at least 128 characters.
	SourceId AnyURI `xml:"http://www.onvif.org/ver10/schema SourceId,omitempty"`

	// Informative user readable name of the source, e.g. "Camera23". A device shall support at least 20 characters.
	Name Name `xml:"http://www.onvif.org/ver10/media/wsdl Name,omitempty"`

	// Informative description of the physical location of the source, e.g. the coordinates on a map.
	Location Description `xml:"http://www.onvif.org/ver10/media/wsdl Location,omitempty"`

	// Informative description of the source.
	Description Description `xml:"http://www.onvif.org/ver10/media/wsdl Description,omitempty"`

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
	Description Description `xml:"http://www.onvif.org/ver10/media/wsdl Description,omitempty"`

	// The start date and time of the oldest recorded data in the track.
	DataFrom string `xml:"http://www.onvif.org/ver10/schema DataFrom,omitempty"`

	// The stop date and time of the newest recorded data in the track.
	DataTo string `xml:"http://www.onvif.org/ver10/schema DataTo,omitempty"`
}

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
	Encoding string `xml:"http://www.onvif.org/ver10/schema Encoding,omitempty"`

	// Average framerate in frames per second.
	Framerate float32 `xml:"http://www.onvif.org/ver10/schema Framerate,omitempty"`
}

// AudioAttributes type
type AudioAttributes struct {

	// The bitrate in kbps.
	Bitrate int32 `xml:"http://www.onvif.org/ver10/schema Bitrate,omitempty"`

	// Audio encoding of the track.  Use values from tt:AudioEncoding for G711 and AAC. Otherwise use values from tt:AudioEncodingMimeNames and .
	Encoding string `xml:"http://www.onvif.org/ver10/schema Encoding,omitempty"`

	// The sample rate in kHz.
	Samplerate int32 `xml:"http://www.onvif.org/ver10/schema Samplerate,omitempty"`
}

// MetadataAttributes type
type MetadataAttributes struct {

	// Indicates that there can be PTZ data in the metadata track in the specified time interval.
	CanContainPTZ bool `xml:"http://www.onvif.org/ver10/media/wsdl CanContainPTZ,omitempty"`

	// Indicates that there can be analytics data in the metadata track in the specified time interval.
	CanContainAnalytics bool `xml:"http://www.onvif.org/ver10/media/wsdl CanContainAnalytics,omitempty"`

	// Indicates that there can be notifications in the metadata track in the specified time interval.
	CanContainNotifications bool `xml:"http://www.onvif.org/ver10/media/wsdl CanContainNotifications,omitempty"`

	// List of all PTZ spaces active for recording. Note that events are only recorded on position changes and the actual point of recording may not necessarily contain an event of the specified type.

	PtzSpaces StringAttrList `xml:"http://www.onvif.org/ver10/schema PtzSpaces,attr,omitempty"`
}

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

	ScheduleToken string `xml:"http://www.onvif.org/ver10/schema ScheduleToken,attr,omitempty"`
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
	AutoCreateReceiver bool `xml:"http://www.onvif.org/ver10/media/wsdl AutoCreateReceiver,omitempty"`

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
	SourceTag string `xml:"http://www.onvif.org/ver10/schema SourceTag,omitempty"`

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
	SourceTag string `xml:"http://www.onvif.org/ver10/schema SourceTag,omitempty"`

	// Indicates the destination track.
	Destination TrackReference `xml:"http://www.onvif.org/ver10/schema Destination,omitempty"`

	// Optionally holds an implementation defined string value that describes the error.
	// The string should be in the English language.
	Error string `xml:"http://www.onvif.org/ver10/schema Error,omitempty"`

	// Provides the job state of the track. The valid
	// values of state shall be “Idle”, “Active” and “Error”. If state equals “Error”, the Error field may be filled in with an implementation defined value.
	State RecordingJobState `xml:"http://www.onvif.org/ver10/schema State,omitempty"`
}

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
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl VideoAnalyticsConfiguration,omitempty"`

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

	VideoInput VideoEncoderConfiguration `xml:"http://www.onvif.org/ver10/media/wsdl VideoInput,omitempty"`

	MetadataInput MetadataInput `xml:"http://www.onvif.org/ver10/schema MetadataInput,omitempty"`
}

// SourceIdentification type
type SourceIdentification struct {
	Name string `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl Token,omitempty"`

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
	Type string `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

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
	Type string `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	//
	// List of supported OSD date formats. This element shall be present when the value of Type field has Date or DateAndTime. The following DateFormat are defined:
	//
	DateFormat string `xml:"http://www.onvif.org/ver10/schema DateFormat,omitempty"`

	//
	// List of supported OSD time formats. This element shall be present when the value of Type field has Time or DateAndTime. The following TimeFormat are defined:
	//
	TimeFormat string `xml:"http://www.onvif.org/ver10/schema TimeFormat,omitempty"`

	// Font size of the text in pt.
	FontSize int32 `xml:"http://www.onvif.org/ver10/schema FontSize,omitempty"`

	// Font color of the text.
	FontColor OSDColor `xml:"http://www.onvif.org/ver10/schema FontColor,omitempty"`

	// Background color of the text.
	BackgroundColor OSDColor `xml:"http://www.onvif.org/ver10/schema BackgroundColor,omitempty"`

	// The content of text to be displayed.
	PlainText string `xml:"http://www.onvif.org/ver10/schema PlainText,omitempty"`

	Extension OSDTextConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// This flag is applicable for Type Plain and defaults to true. When set to false the PlainText content will not be persistent across device reboots.

	IsPersistentText bool `xml:"http://www.onvif.org/ver10/media/wsdl IsPersistentText,attr,omitempty"`
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
	Type []string `xml:"http://www.onvif.org/ver10/schema Type,omitempty"`

	// Range of the font size value.
	FontSizeRange IntRange `xml:"http://www.onvif.org/ver10/schema FontSizeRange,omitempty"`

	// List of supported date format.
	DateFormat []string `xml:"http://www.onvif.org/ver10/schema DateFormat,omitempty"`

	// List of supported time format.
	TimeFormat []string `xml:"http://www.onvif.org/ver10/schema TimeFormat,omitempty"`

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

	FormatsSupported StringAttrList `xml:"http://www.onvif.org/ver10/schema FormatsSupported,attr,omitempty"`

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
	PositionOption []string `xml:"http://www.onvif.org/ver10/schema PositionOption,omitempty"`

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
	FileName string `xml:"http://www.onvif.org/ver10/schema FileName,omitempty"`

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
	StorageToken ReferenceToken `xml:"http://www.onvif.org/ver10/media/wsdl StorageToken,omitempty"`

	// gives the relative directory path on the storage
	RelativePath string `xml:"http://www.onvif.org/ver10/schema RelativePath,omitempty"`

	Extension StorageReferencePathExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// StorageReferencePathExtension type
type StorageReferencePathExtension struct {
}

// Media type
type Media interface {

	/* Returns the capabilities of the media service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/* This command lists all available physical video inputs of the device. */
	GetVideoSources(request *GetVideoSources) (*GetVideoSourcesResponse, error)

	GetVideoSourcesContext(ctx context.Context, request *GetVideoSources) (*GetVideoSourcesResponse, error)

	/* This command lists all available physical audio inputs of the device. */
	GetAudioSources(request *GetAudioSources) (*GetAudioSourcesResponse, error)

	GetAudioSourcesContext(ctx context.Context, request *GetAudioSources) (*GetAudioSourcesResponse, error)

	/* This command lists all available physical audio outputs of the device. */
	GetAudioOutputs(request *GetAudioOutputs) (*GetAudioOutputsResponse, error)

	GetAudioOutputsContext(ctx context.Context, request *GetAudioOutputs) (*GetAudioOutputsResponse, error)

	/* This operation creates a new empty media profile. The media profile shall be created in the
	device and shall be persistent (remain after reboot). A created profile shall be deletable and a device shall set the “fixed” attribute to false in the
	returned Profile. */
	CreateProfile(request *CreateProfile) (*CreateProfileResponse, error)

	CreateProfileContext(ctx context.Context, request *CreateProfile) (*CreateProfileResponse, error)

	/* If the profile token is already known, a profile can be fetched through the GetProfile command. */
	GetProfile(request *GetProfile) (*GetProfileResponse, error)

	GetProfileContext(ctx context.Context, request *GetProfile) (*GetProfileResponse, error)

	/* Any endpoint can ask for the existing media profiles of a device using the GetProfiles
	command. Pre-configured or dynamically configured profiles can be retrieved using this
	command. This command lists all configured profiles in a device. The client does not need to
	know the media profile in order to use the command. */
	GetProfiles(request *GetProfiles) (*GetProfilesResponse, error)

	GetProfilesContext(ctx context.Context, request *GetProfiles) (*GetProfilesResponse, error)

	/* This operation adds a VideoEncoderConfiguration to an existing media profile. If a
	configuration exists in the media profile, it will be replaced. The change shall be persistent. A device shall
	support adding a compatible VideoEncoderConfiguration to a Profile containing a VideoSourceConfiguration and shall
	support streaming video data of such a profile.
	*/
	AddVideoEncoderConfiguration(request *AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, error)

	AddVideoEncoderConfigurationContext(ctx context.Context, request *AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, error)

	/* This operation removes a VideoEncoderConfiguration from an existing media profile. If the
	media profile does not contain a VideoEncoderConfiguration, the operation has no effect. The removal shall be persistent. */
	RemoveVideoEncoderConfiguration(request *RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, error)

	RemoveVideoEncoderConfigurationContext(ctx context.Context, request *RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, error)

	/* This operation adds a VideoSourceConfiguration to an existing media profile. If such a
	configuration exists in the media profile, it will be replaced. The change shall be persistent. */
	AddVideoSourceConfiguration(request *AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, error)

	AddVideoSourceConfigurationContext(ctx context.Context, request *AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, error)

	/* This operation removes a VideoSourceConfiguration from an existing media profile. If the
	media profile does not contain a VideoSourceConfiguration, the operation has no effect. The removal shall be persistent. Video source configurations should only be removed after removing a
	VideoEncoderConfiguration from the media profile. */
	RemoveVideoSourceConfiguration(request *RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, error)

	RemoveVideoSourceConfigurationContext(ctx context.Context, request *RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, error)

	/* This operation adds an AudioEncoderConfiguration to an existing media profile. If a
	configuration exists in the media profile, it will be replaced. The change shall be persistent. A device shall
	support adding a compatible AudioEncoderConfiguration to a profile containing an AudioSourceConfiguration and shall
	support streaming audio data of such a profile.
	*/
	AddAudioEncoderConfiguration(request *AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, error)

	AddAudioEncoderConfigurationContext(ctx context.Context, request *AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, error)

	/* This operation removes an AudioEncoderConfiguration from an existing media profile. If the
	media profile does not contain an AudioEncoderConfiguration, the operation has no effect.
	The removal shall be persistent. */
	RemoveAudioEncoderConfiguration(request *RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, error)

	RemoveAudioEncoderConfigurationContext(ctx context.Context, request *RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, error)

	/* This operation adds an AudioSourceConfiguration to an existing media profile. If a
	configuration exists in the media profile, it will be replaced. The change shall be persistent. */
	AddAudioSourceConfiguration(request *AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, error)

	AddAudioSourceConfigurationContext(ctx context.Context, request *AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, error)

	/* This operation removes an AudioSourceConfiguration from an existing media profile. If the
	media profile does not contain an AudioSourceConfiguration, the operation has no effect. The
	removal shall be persistent. Audio source configurations should only be removed after removing an
	AudioEncoderConfiguration from the media profile. */
	RemoveAudioSourceConfiguration(request *RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, error)

	RemoveAudioSourceConfigurationContext(ctx context.Context, request *RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, error)

	/* This operation adds a PTZConfiguration to an existing media profile. If a configuration exists
	in the media profile, it will be replaced. The change shall be persistent. Adding a PTZConfiguration to a media profile means that streams using that media profile can
	contain PTZ status (in the metadata), and that the media profile can be used for controlling
	PTZ movement. */
	AddPTZConfiguration(request *AddPTZConfiguration) (*AddPTZConfigurationResponse, error)

	AddPTZConfigurationContext(ctx context.Context, request *AddPTZConfiguration) (*AddPTZConfigurationResponse, error)

	/* This operation removes a PTZConfiguration from an existing media profile. If the media profile
	does not contain a PTZConfiguration, the operation has no effect. The removal shall be persistent. */
	RemovePTZConfiguration(request *RemovePTZConfiguration) (*RemovePTZConfigurationResponse, error)

	RemovePTZConfigurationContext(ctx context.Context, request *RemovePTZConfiguration) (*RemovePTZConfigurationResponse, error)

	/* This operation adds a VideoAnalytics configuration to an existing media profile. If a
	configuration exists in the media profile, it will be replaced. The change shall be persistent. Adding a VideoAnalyticsConfiguration to a media profile means that streams using that media
	profile can contain video analytics data (in the metadata) as defined by the submitted configuration reference. A profile containing only a video analytics configuration but no video source configuration is incomplete. Therefore, a client should first add a video source configuration to a profile before adding a video analytics configuration. The device can deny adding of a video analytics
	configuration before a video source configuration. */
	AddVideoAnalyticsConfiguration(request *AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, error)

	AddVideoAnalyticsConfigurationContext(ctx context.Context, request *AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, error)

	/* This operation removes a VideoAnalyticsConfiguration from an existing media profile. If the media profile does not contain a VideoAnalyticsConfiguration, the operation has no effect.
	The removal shall be persistent. */
	RemoveVideoAnalyticsConfiguration(request *RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, error)

	RemoveVideoAnalyticsConfigurationContext(ctx context.Context, request *RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, error)

	/* This operation adds a Metadata configuration to an existing media profile. If a configuration exists in the media profile, it will be replaced. The change shall be persistent. Adding a MetadataConfiguration to a Profile means that streams using that profile contain metadata. Metadata can consist of events, PTZ status, and/or video analytics data. */
	AddMetadataConfiguration(request *AddMetadataConfiguration) (*AddMetadataConfigurationResponse, error)

	AddMetadataConfigurationContext(ctx context.Context, request *AddMetadataConfiguration) (*AddMetadataConfigurationResponse, error)

	/* This operation removes a MetadataConfiguration from an existing media profile. If the media profile does not contain a MetadataConfiguration, the operation has no effect. The removal shall be persistent. */
	RemoveMetadataConfiguration(request *RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, error)

	RemoveMetadataConfigurationContext(ctx context.Context, request *RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, error)

	/* This operation adds an AudioOutputConfiguration to an existing media profile. If a configuration exists in the media profile, it will be replaced. The change shall be persistent. */
	AddAudioOutputConfiguration(request *AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, error)

	AddAudioOutputConfigurationContext(ctx context.Context, request *AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, error)

	/* This operation removes an AudioOutputConfiguration from an existing media profile. If the media profile does not contain an AudioOutputConfiguration, the operation has no effect. The removal shall be persistent. */
	RemoveAudioOutputConfiguration(request *RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, error)

	RemoveAudioOutputConfigurationContext(ctx context.Context, request *RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, error)

	/* This operation adds an AudioDecoderConfiguration to an existing media profile. If a configuration exists in the media profile, it shall be replaced. The change shall be persistent. */
	AddAudioDecoderConfiguration(request *AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, error)

	AddAudioDecoderConfigurationContext(ctx context.Context, request *AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, error)

	/* This operation removes an AudioDecoderConfiguration from an existing media profile. If the media profile does not contain an AudioDecoderConfiguration, the operation has no effect. The removal shall be persistent. */
	RemoveAudioDecoderConfiguration(request *RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, error)

	RemoveAudioDecoderConfigurationContext(ctx context.Context, request *RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, error)

	/* This operation deletes a profile. This change shall always be persistent. Deletion of a profile is only possible for non-fixed profiles */
	DeleteProfile(request *DeleteProfile) (*DeleteProfileResponse, error)

	DeleteProfileContext(ctx context.Context, request *DeleteProfile) (*DeleteProfileResponse, error)

	/* This operation lists all existing video source configurations for a device. The client need not know anything about the video source configurations in order to use the command. */
	GetVideoSourceConfigurations(request *GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, error)

	GetVideoSourceConfigurationsContext(ctx context.Context, request *GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, error)

	/* This operation lists all existing video encoder configurations of a device. This command lists all configured video encoder configurations in a device. The client need not know anything apriori about the video encoder configurations in order to use the command. */
	GetVideoEncoderConfigurations(request *GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, error)

	GetVideoEncoderConfigurationsContext(ctx context.Context, request *GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, error)

	/* This operation lists all existing audio source configurations of a device. This command lists all audio source configurations in a device. The client need not know anything apriori about the audio source configurations in order to use the command. */
	GetAudioSourceConfigurations(request *GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, error)

	GetAudioSourceConfigurationsContext(ctx context.Context, request *GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, error)

	/* This operation lists all existing device audio encoder configurations. The client need not know anything apriori about the audio encoder configurations in order to use the command. */
	GetAudioEncoderConfigurations(request *GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, error)

	GetAudioEncoderConfigurationsContext(ctx context.Context, request *GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, error)

	/* This operation lists all video analytics configurations of a device. This command lists all configured video analytics in a device. The client need not know anything apriori about the video analytics in order to use the command. */
	GetVideoAnalyticsConfigurations(request *GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, error)

	GetVideoAnalyticsConfigurationsContext(ctx context.Context, request *GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, error)

	/* This operation lists all existing metadata configurations. The client need not know anything apriori about the metadata in order to use the command. */
	GetMetadataConfigurations(request *GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, error)

	GetMetadataConfigurationsContext(ctx context.Context, request *GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, error)

	/* This command lists all existing AudioOutputConfigurations of a device. The NVC need not know anything apriori about the audio configurations to use this command. */
	GetAudioOutputConfigurations(request *GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, error)

	GetAudioOutputConfigurationsContext(ctx context.Context, request *GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, error)

	/* This command lists all existing AudioDecoderConfigurations of a device. The NVC need not know anything apriori about the audio decoder configurations in order to
	use this command. */
	GetAudioDecoderConfigurations(request *GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, error)

	GetAudioDecoderConfigurationsContext(ctx context.Context, request *GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, error)

	/* If the video source configuration token is already known, the video source configuration can be fetched through the GetVideoSourceConfiguration command. */
	GetVideoSourceConfiguration(request *GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, error)

	GetVideoSourceConfigurationContext(ctx context.Context, request *GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, error)

	/* If the video encoder configuration token is already known, the encoder configuration can be fetched through the GetVideoEncoderConfiguration command. */
	GetVideoEncoderConfiguration(request *GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, error)

	GetVideoEncoderConfigurationContext(ctx context.Context, request *GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, error)

	/* The GetAudioSourceConfiguration command fetches the audio source configurations if the audio source configuration token is already known. An */
	GetAudioSourceConfiguration(request *GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, error)

	GetAudioSourceConfigurationContext(ctx context.Context, request *GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, error)

	/* The GetAudioEncoderConfiguration command fetches the encoder configuration if the audio encoder configuration token is known. */
	GetAudioEncoderConfiguration(request *GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, error)

	GetAudioEncoderConfigurationContext(ctx context.Context, request *GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, error)

	/* The GetVideoAnalyticsConfiguration command fetches the video analytics configuration if the video analytics token is known. */
	GetVideoAnalyticsConfiguration(request *GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, error)

	GetVideoAnalyticsConfigurationContext(ctx context.Context, request *GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, error)

	/* The GetMetadataConfiguration command fetches the metadata configuration if the metadata token is known. */
	GetMetadataConfiguration(request *GetMetadataConfiguration) (*GetMetadataConfigurationResponse, error)

	GetMetadataConfigurationContext(ctx context.Context, request *GetMetadataConfiguration) (*GetMetadataConfigurationResponse, error)

	/* If the audio output configuration token is already known, the output configuration can be fetched through the GetAudioOutputConfiguration command. */
	GetAudioOutputConfiguration(request *GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, error)

	GetAudioOutputConfigurationContext(ctx context.Context, request *GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, error)

	/* If the audio decoder configuration token is already known, the decoder configuration can be fetched through the GetAudioDecoderConfiguration command. */
	GetAudioDecoderConfiguration(request *GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, error)

	GetAudioDecoderConfigurationContext(ctx context.Context, request *GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, error)

	/* This operation lists all the video encoder configurations of the device that are compatible with a certain media profile. Each of the returned configurations shall be a valid input parameter for the AddVideoEncoderConfiguration command on the media profile. The result will vary depending on the capabilities, configurations and settings in the device. */
	GetCompatibleVideoEncoderConfigurations(request *GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, error)

	GetCompatibleVideoEncoderConfigurationsContext(ctx context.Context, request *GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, error)

	/* This operation requests all the video source configurations of the device that are compatible
	with a certain media profile. Each of the returned configurations shall be a valid input
	parameter for the AddVideoSourceConfiguration command on the media profile. The result
	will vary depending on the capabilities, configurations and settings in the device. */
	GetCompatibleVideoSourceConfigurations(request *GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, error)

	GetCompatibleVideoSourceConfigurationsContext(ctx context.Context, request *GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, error)

	/* This operation requests all audio encoder configurations of a device that are compatible with a certain media profile. Each of the returned configurations shall be a valid input parameter for the AddAudioSourceConfiguration command on the media profile. The result varies depending on the capabilities, configurations and settings in the device. */
	GetCompatibleAudioEncoderConfigurations(request *GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, error)

	GetCompatibleAudioEncoderConfigurationsContext(ctx context.Context, request *GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, error)

	/* This operation requests all audio source configurations of the device that are compatible with a certain media profile. Each of the returned configurations shall be a valid input parameter for the AddAudioEncoderConfiguration command on the media profile. The result varies depending on the capabilities, configurations and settings in the device. */
	GetCompatibleAudioSourceConfigurations(request *GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, error)

	GetCompatibleAudioSourceConfigurationsContext(ctx context.Context, request *GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, error)

	/* This operation requests all video analytic configurations of the device that are compatible with a certain media profile. Each of the returned configurations shall be a valid input parameter for the AddVideoAnalyticsConfiguration command on the media profile. The result varies depending on the capabilities, configurations and settings in the device. */
	GetCompatibleVideoAnalyticsConfigurations(request *GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, error)

	GetCompatibleVideoAnalyticsConfigurationsContext(ctx context.Context, request *GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, error)

	/* This operation requests all the metadata configurations of the device that are compatible with a certain media profile. Each of the returned configurations shall be a valid input parameter for the AddMetadataConfiguration command on the media profile. The result varies depending on the capabilities, configurations and settings in the device. */
	GetCompatibleMetadataConfigurations(request *GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, error)

	GetCompatibleMetadataConfigurationsContext(ctx context.Context, request *GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, error)

	/* This command lists all audio output configurations of a device that are compatible with a certain media profile. Each returned configuration shall be a valid input for the
	AddAudioOutputConfiguration command. */
	GetCompatibleAudioOutputConfigurations(request *GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, error)

	GetCompatibleAudioOutputConfigurationsContext(ctx context.Context, request *GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, error)

	/* This operation lists all the audio decoder configurations of the device that are compatible with a certain media profile. Each of the returned configurations shall be a valid input parameter for the AddAudioDecoderConfiguration command on the media profile. */
	GetCompatibleAudioDecoderConfigurations(request *GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, error)

	GetCompatibleAudioDecoderConfigurationsContext(ctx context.Context, request *GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, error)

	/* This operation modifies a video source configuration. The ForcePersistence flag indicates if the changes shall remain after reboot of the device. Running streams using this configuration may be immediately updated according to the new settings. The changes are not guaranteed to take effect unless the client requests a new stream URI and restarts any affected stream. NVC methods for changing a running stream are out of scope for this specification. */
	SetVideoSourceConfiguration(request *SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, error)

	SetVideoSourceConfigurationContext(ctx context.Context, request *SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, error)

	/* This operation modifies a video encoder configuration. The ForcePersistence flag indicates if the changes shall remain after reboot of the device. Changes in the Multicast settings shall always be persistent. Running streams using this configuration may be immediately updated according to the new settings. The changes are not guaranteed to take effect unless the client requests a new stream URI and restarts any affected stream. NVC methods for changing a running stream are out of scope for this specification. SessionTimeout is provided as a hint for keeping rtsp session by a device. If necessary the device may adapt parameter values for SessionTimeout elements without returning an error. For the time between keep alive calls the client shall adhere to the timeout value signaled via RTSP. */
	SetVideoEncoderConfiguration(request *SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, error)

	SetVideoEncoderConfigurationContext(ctx context.Context, request *SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, error)

	/* This operation modifies an audio source configuration. The ForcePersistence flag indicates if
	the changes shall remain after reboot of the device. Running streams using this configuration
	may be immediately updated according to the new settings. The changes are not guaranteed
	to take effect unless the client requests a new stream URI and restarts any affected stream
	NVC methods for changing a running stream are out of scope for this specification. */
	SetAudioSourceConfiguration(request *SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, error)

	SetAudioSourceConfigurationContext(ctx context.Context, request *SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, error)

	/* This operation modifies an audio encoder configuration. The ForcePersistence flag indicates if
	the changes shall remain after reboot of the device. Running streams using this configuration may be immediately updated
	according to the new settings. The changes are not guaranteed to take effect unless the client
	requests a new stream URI and restarts any affected streams. NVC methods for changing a
	running stream are out of scope for this specification. */
	SetAudioEncoderConfiguration(request *SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, error)

	SetAudioEncoderConfigurationContext(ctx context.Context, request *SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, error)

	/* A video analytics configuration is modified using this command. The ForcePersistence flag
	indicates if the changes shall remain after reboot of the device or not. Running streams using
	this configuration shall be immediately updated according to the new settings. Otherwise
	inconsistencies can occur between the scene description processed by the rule engine and
	the notifications produced by analytics engine and rule engine which reference the very same
	video analytics configuration token. */
	SetVideoAnalyticsConfiguration(request *SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, error)

	SetVideoAnalyticsConfigurationContext(ctx context.Context, request *SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, error)

	/* This operation modifies a metadata configuration. The ForcePersistence flag indicates if the
	changes shall remain after reboot of the device. Changes in the Multicast settings shall
	always be persistent. Running streams using this configuration may be updated immediately
	according to the new settings. The changes are not guaranteed to take effect unless the client
	requests a new stream URI and restarts any affected streams. NVC methods for changing a
	running stream are out of scope for this specification. */
	SetMetadataConfiguration(request *SetMetadataConfiguration) (*SetMetadataConfigurationResponse, error)

	SetMetadataConfigurationContext(ctx context.Context, request *SetMetadataConfiguration) (*SetMetadataConfigurationResponse, error)

	/* This operation modifies an audio output configuration. The ForcePersistence flag indicates if
	the changes shall remain after reboot of the device. */
	SetAudioOutputConfiguration(request *SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, error)

	SetAudioOutputConfigurationContext(ctx context.Context, request *SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, error)

	/* This operation modifies an audio decoder configuration. The ForcePersistence flag indicates if
	the changes shall remain after reboot of the device. */
	SetAudioDecoderConfiguration(request *SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, error)

	SetAudioDecoderConfigurationContext(ctx context.Context, request *SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, error)

	/* This operation returns the available options  (supported values and ranges for video source configuration parameters) when the video source parameters are
	reconfigured If a video source configuration is specified, the options shall concern that
	particular configuration. If a media profile is specified, the options shall be compatible with
	that media profile. */
	GetVideoSourceConfigurationOptions(request *GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, error)

	GetVideoSourceConfigurationOptionsContext(ctx context.Context, request *GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, error)

	/* This operation returns the available options (supported values and ranges for video encoder
	configuration parameters) when the video encoder parameters are reconfigured.
	For JPEG, MPEG4 and H264 extension elements have been defined that provide additional information. A device must provide the
	XxxOption information for all encodings supported and should additionally provide the corresponding XxxOption2 information.
	This response contains the available video encoder configuration options. If a video encoder configuration is specified,
	the options shall concern that particular configuration. If a media profile is specified, the options shall be
	compatible with that media profile. If no tokens are specified, the options shall be considered generic for the device.
	*/
	GetVideoEncoderConfigurationOptions(request *GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, error)

	GetVideoEncoderConfigurationOptionsContext(ctx context.Context, request *GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, error)

	/* This operation returns the available options (supported values and ranges for audio source configuration parameters) when the audio source parameters are
	reconfigured. If an audio source configuration is specified, the options shall concern that
	particular configuration. If a media profile is specified, the options shall be compatible with
	that media profile. */
	GetAudioSourceConfigurationOptions(request *GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, error)

	GetAudioSourceConfigurationOptionsContext(ctx context.Context, request *GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, error)

	/* This operation returns the available options  (supported values and ranges for audio encoder configuration parameters) when the audio encoder parameters are
	reconfigured. */
	GetAudioEncoderConfigurationOptions(request *GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, error)

	GetAudioEncoderConfigurationOptionsContext(ctx context.Context, request *GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, error)

	/* This operation returns the available options (supported values and ranges for metadata configuration parameters) for changing the metadata configuration. */
	GetMetadataConfigurationOptions(request *GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, error)

	GetMetadataConfigurationOptionsContext(ctx context.Context, request *GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, error)

	/* This operation returns the available options (supported values and ranges for audio output configuration parameters) for configuring an audio output. */
	GetAudioOutputConfigurationOptions(request *GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, error)

	GetAudioOutputConfigurationOptionsContext(ctx context.Context, request *GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, error)

	/* This command list the audio decoding capabilities for a given profile and configuration of a
	device. */
	GetAudioDecoderConfigurationOptions(request *GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, error)

	GetAudioDecoderConfigurationOptionsContext(ctx context.Context, request *GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, error)

	/* The GetGuaranteedNumberOfVideoEncoderInstances command can be used to request the
	minimum number of guaranteed video encoder instances (applications) per Video Source
	Configuration. */
	GetGuaranteedNumberOfVideoEncoderInstances(request *GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error)

	GetGuaranteedNumberOfVideoEncoderInstancesContext(ctx context.Context, request *GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error)

	/* This operation requests a URI that can be used to initiate a live media stream using RTSP as
	the control protocol. The returned URI shall remain valid indefinitely even if the profile is
	changed. The ValidUntilConnect, ValidUntilReboot and Timeout Parameter shall be set
	accordingly (ValidUntilConnect=false, ValidUntilReboot=false, timeout=PT0S).
					The correct syntax for the StreamSetup element for these media stream setups defined in 5.1.1 of the streaming specification are as follows:


	If a multicast stream is requested at least one of VideoEncoderConfiguration, AudioEncoderConfiguration and MetadataConfiguration shall have a valid multicast setting.
	For full compatibility with other ONVIF services a device should not generate Uris longer than
	128 octets. */
	GetStreamUri(request *GetStreamUri) (*GetStreamUriResponse, error)

	GetStreamUriContext(ctx context.Context, request *GetStreamUri) (*GetStreamUriResponse, error)

	/* This command starts multicast streaming using a specified media profile of a device.
	Streaming continues until StopMulticastStreaming is called for the same Profile. The
	streaming shall continue after a reboot of the device until a StopMulticastStreaming request is
	received. The multicast address, port and TTL are configured in the
	VideoEncoderConfiguration, AudioEncoderConfiguration and MetadataConfiguration
	respectively. */
	StartMulticastStreaming(request *StartMulticastStreaming) (*StartMulticastStreamingResponse, error)

	StartMulticastStreamingContext(ctx context.Context, request *StartMulticastStreaming) (*StartMulticastStreamingResponse, error)

	/* This command stop multicast streaming using a specified media profile of a device */
	StopMulticastStreaming(request *StopMulticastStreaming) (*StopMulticastStreamingResponse, error)

	StopMulticastStreamingContext(ctx context.Context, request *StopMulticastStreaming) (*StopMulticastStreamingResponse, error)

	/* Synchronization points allow clients to decode and correctly use all data after the
	synchronization point.
	For example, if a video stream is configured with a large I-frame distance and a client loses a
	single packet, the client does not display video until the next I-frame is transmitted. In such
	cases, the client can request a Synchronization Point which enforces the device to add an I-Frame as soon as possible. Clients can request Synchronization Points for profiles. The device
	shall add synchronization points for all streams associated with this profile.
	Similarly, a synchronization point is used to get an update on full PTZ or event status through
	the metadata stream.
	If a video stream is associated with the profile, an I-frame shall be added to this video stream.
	If a PTZ metadata stream is associated to the profile,
	the PTZ position shall be repeated within the metadata stream. */
	SetSynchronizationPoint(request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	SetSynchronizationPointContext(ctx context.Context, request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	/* A client uses the GetSnapshotUri command to obtain a JPEG snapshot from the
	device. The returned URI shall remain valid indefinitely even if the profile is changed. The
	ValidUntilConnect, ValidUntilReboot and Timeout Parameter shall be set accordingly
	(ValidUntilConnect=false, ValidUntilReboot=false, timeout=PT0S). The URI can be used for
	acquiring a JPEG image through a HTTP GET operation. The image encoding will always be
	JPEG regardless of the encoding setting in the media profile. The Jpeg settings
	(like resolution or quality) may be taken from the profile if suitable. The provided
	image will be updated automatically and independent from calls to GetSnapshotUri. */
	GetSnapshotUri(request *GetSnapshotUri) (*GetSnapshotUriResponse, error)

	GetSnapshotUriContext(ctx context.Context, request *GetSnapshotUri) (*GetSnapshotUriResponse, error)

	/* A device returns the information for current video source mode and settable video source modes of specified video source. A device that indicates a capability of  VideoSourceModes shall support this command. */
	GetVideoSourceModes(request *GetVideoSourceModes) (*GetVideoSourceModesResponse, error)

	GetVideoSourceModesContext(ctx context.Context, request *GetVideoSourceModes) (*GetVideoSourceModesResponse, error)

	/* SetVideoSourceMode changes the media profile structure relating to video source for the specified video source mode. A device that indicates a capability of VideoSourceModes shall support this command. The behavior after changing the mode is not defined in this specification. */
	SetVideoSourceMode(request *SetVideoSourceMode) (*SetVideoSourceModeResponse, error)

	SetVideoSourceModeContext(ctx context.Context, request *SetVideoSourceMode) (*SetVideoSourceModeResponse, error)

	/* Get the OSDs. */
	GetOSDs(request *GetOSDs) (*GetOSDsResponse, error)

	GetOSDsContext(ctx context.Context, request *GetOSDs) (*GetOSDsResponse, error)

	/* Get the OSD. */
	GetOSD(request *GetOSD) (*GetOSDResponse, error)

	GetOSDContext(ctx context.Context, request *GetOSD) (*GetOSDResponse, error)

	/* Get the OSD Options. */
	GetOSDOptions(request *GetOSDOptions) (*GetOSDOptionsResponse, error)

	GetOSDOptionsContext(ctx context.Context, request *GetOSDOptions) (*GetOSDOptionsResponse, error)

	/* Set the OSD */
	SetOSD(request *SetOSD) (*SetOSDResponse, error)

	SetOSDContext(ctx context.Context, request *SetOSD) (*SetOSDResponse, error)

	/* Create the OSD. */
	CreateOSD(request *CreateOSD) (*CreateOSDResponse, error)

	CreateOSDContext(ctx context.Context, request *CreateOSD) (*CreateOSDResponse, error)

	/* Delete the OSD. */
	DeleteOSD(request *DeleteOSD) (*DeleteOSDResponse, error)

	DeleteOSDContext(ctx context.Context, request *DeleteOSD) (*DeleteOSDResponse, error)
}

// media type
type media struct {
	client *soap.Client
	xaddr  string
}

func NewMedia(client *soap.Client, xaddr string) Media {
	return &media{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *media) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoSourcesContext(ctx context.Context, request *GetVideoSources) (*GetVideoSourcesResponse, error) {
	response := new(GetVideoSourcesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoSources", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoSources(request *GetVideoSources) (*GetVideoSourcesResponse, error) {
	return service.GetVideoSourcesContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioSourcesContext(ctx context.Context, request *GetAudioSources) (*GetAudioSourcesResponse, error) {
	response := new(GetAudioSourcesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioSources", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioSources(request *GetAudioSources) (*GetAudioSourcesResponse, error) {
	return service.GetAudioSourcesContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioOutputsContext(ctx context.Context, request *GetAudioOutputs) (*GetAudioOutputsResponse, error) {
	response := new(GetAudioOutputsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioOutputs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioOutputs(request *GetAudioOutputs) (*GetAudioOutputsResponse, error) {
	return service.GetAudioOutputsContext(
		context.Background(),
		request,
	)
}

func (service *media) CreateProfileContext(ctx context.Context, request *CreateProfile) (*CreateProfileResponse, error) {
	response := new(CreateProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/CreateProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) CreateProfile(request *CreateProfile) (*CreateProfileResponse, error) {
	return service.CreateProfileContext(
		context.Background(),
		request,
	)
}

func (service *media) GetProfileContext(ctx context.Context, request *GetProfile) (*GetProfileResponse, error) {
	response := new(GetProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetProfile(request *GetProfile) (*GetProfileResponse, error) {
	return service.GetProfileContext(
		context.Background(),
		request,
	)
}

func (service *media) GetProfilesContext(ctx context.Context, request *GetProfiles) (*GetProfilesResponse, error) {
	response := new(GetProfilesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetProfiles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetProfiles(request *GetProfiles) (*GetProfilesResponse, error) {
	return service.GetProfilesContext(
		context.Background(),
		request,
	)
}

func (service *media) AddVideoEncoderConfigurationContext(ctx context.Context, request *AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, error) {
	response := new(AddVideoEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddVideoEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddVideoEncoderConfiguration(request *AddVideoEncoderConfiguration) (*AddVideoEncoderConfigurationResponse, error) {
	return service.AddVideoEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveVideoEncoderConfigurationContext(ctx context.Context, request *RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, error) {
	response := new(RemoveVideoEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveVideoEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveVideoEncoderConfiguration(request *RemoveVideoEncoderConfiguration) (*RemoveVideoEncoderConfigurationResponse, error) {
	return service.RemoveVideoEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddVideoSourceConfigurationContext(ctx context.Context, request *AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, error) {
	response := new(AddVideoSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddVideoSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddVideoSourceConfiguration(request *AddVideoSourceConfiguration) (*AddVideoSourceConfigurationResponse, error) {
	return service.AddVideoSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveVideoSourceConfigurationContext(ctx context.Context, request *RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, error) {
	response := new(RemoveVideoSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveVideoSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveVideoSourceConfiguration(request *RemoveVideoSourceConfiguration) (*RemoveVideoSourceConfigurationResponse, error) {
	return service.RemoveVideoSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddAudioEncoderConfigurationContext(ctx context.Context, request *AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, error) {
	response := new(AddAudioEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddAudioEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddAudioEncoderConfiguration(request *AddAudioEncoderConfiguration) (*AddAudioEncoderConfigurationResponse, error) {
	return service.AddAudioEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveAudioEncoderConfigurationContext(ctx context.Context, request *RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, error) {
	response := new(RemoveAudioEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveAudioEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveAudioEncoderConfiguration(request *RemoveAudioEncoderConfiguration) (*RemoveAudioEncoderConfigurationResponse, error) {
	return service.RemoveAudioEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddAudioSourceConfigurationContext(ctx context.Context, request *AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, error) {
	response := new(AddAudioSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddAudioSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddAudioSourceConfiguration(request *AddAudioSourceConfiguration) (*AddAudioSourceConfigurationResponse, error) {
	return service.AddAudioSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveAudioSourceConfigurationContext(ctx context.Context, request *RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, error) {
	response := new(RemoveAudioSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveAudioSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveAudioSourceConfiguration(request *RemoveAudioSourceConfiguration) (*RemoveAudioSourceConfigurationResponse, error) {
	return service.RemoveAudioSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddPTZConfigurationContext(ctx context.Context, request *AddPTZConfiguration) (*AddPTZConfigurationResponse, error) {
	response := new(AddPTZConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddPTZConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddPTZConfiguration(request *AddPTZConfiguration) (*AddPTZConfigurationResponse, error) {
	return service.AddPTZConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemovePTZConfigurationContext(ctx context.Context, request *RemovePTZConfiguration) (*RemovePTZConfigurationResponse, error) {
	response := new(RemovePTZConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemovePTZConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemovePTZConfiguration(request *RemovePTZConfiguration) (*RemovePTZConfigurationResponse, error) {
	return service.RemovePTZConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddVideoAnalyticsConfigurationContext(ctx context.Context, request *AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, error) {
	response := new(AddVideoAnalyticsConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddVideoAnalyticsConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddVideoAnalyticsConfiguration(request *AddVideoAnalyticsConfiguration) (*AddVideoAnalyticsConfigurationResponse, error) {
	return service.AddVideoAnalyticsConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveVideoAnalyticsConfigurationContext(ctx context.Context, request *RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, error) {
	response := new(RemoveVideoAnalyticsConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveVideoAnalyticsConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveVideoAnalyticsConfiguration(request *RemoveVideoAnalyticsConfiguration) (*RemoveVideoAnalyticsConfigurationResponse, error) {
	return service.RemoveVideoAnalyticsConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddMetadataConfigurationContext(ctx context.Context, request *AddMetadataConfiguration) (*AddMetadataConfigurationResponse, error) {
	response := new(AddMetadataConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddMetadataConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddMetadataConfiguration(request *AddMetadataConfiguration) (*AddMetadataConfigurationResponse, error) {
	return service.AddMetadataConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveMetadataConfigurationContext(ctx context.Context, request *RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, error) {
	response := new(RemoveMetadataConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveMetadataConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveMetadataConfiguration(request *RemoveMetadataConfiguration) (*RemoveMetadataConfigurationResponse, error) {
	return service.RemoveMetadataConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddAudioOutputConfigurationContext(ctx context.Context, request *AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, error) {
	response := new(AddAudioOutputConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddAudioOutputConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddAudioOutputConfiguration(request *AddAudioOutputConfiguration) (*AddAudioOutputConfigurationResponse, error) {
	return service.AddAudioOutputConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveAudioOutputConfigurationContext(ctx context.Context, request *RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, error) {
	response := new(RemoveAudioOutputConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveAudioOutputConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveAudioOutputConfiguration(request *RemoveAudioOutputConfiguration) (*RemoveAudioOutputConfigurationResponse, error) {
	return service.RemoveAudioOutputConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) AddAudioDecoderConfigurationContext(ctx context.Context, request *AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, error) {
	response := new(AddAudioDecoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/AddAudioDecoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) AddAudioDecoderConfiguration(request *AddAudioDecoderConfiguration) (*AddAudioDecoderConfigurationResponse, error) {
	return service.AddAudioDecoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) RemoveAudioDecoderConfigurationContext(ctx context.Context, request *RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, error) {
	response := new(RemoveAudioDecoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/RemoveAudioDecoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) RemoveAudioDecoderConfiguration(request *RemoveAudioDecoderConfiguration) (*RemoveAudioDecoderConfigurationResponse, error) {
	return service.RemoveAudioDecoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) DeleteProfileContext(ctx context.Context, request *DeleteProfile) (*DeleteProfileResponse, error) {
	response := new(DeleteProfileResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/DeleteProfile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) DeleteProfile(request *DeleteProfile) (*DeleteProfileResponse, error) {
	return service.DeleteProfileContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoSourceConfigurationsContext(ctx context.Context, request *GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, error) {
	response := new(GetVideoSourceConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoSourceConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoSourceConfigurations(request *GetVideoSourceConfigurations) (*GetVideoSourceConfigurationsResponse, error) {
	return service.GetVideoSourceConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoEncoderConfigurationsContext(ctx context.Context, request *GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, error) {
	response := new(GetVideoEncoderConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoEncoderConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoEncoderConfigurations(request *GetVideoEncoderConfigurations) (*GetVideoEncoderConfigurationsResponse, error) {
	return service.GetVideoEncoderConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioSourceConfigurationsContext(ctx context.Context, request *GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, error) {
	response := new(GetAudioSourceConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioSourceConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioSourceConfigurations(request *GetAudioSourceConfigurations) (*GetAudioSourceConfigurationsResponse, error) {
	return service.GetAudioSourceConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioEncoderConfigurationsContext(ctx context.Context, request *GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, error) {
	response := new(GetAudioEncoderConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioEncoderConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioEncoderConfigurations(request *GetAudioEncoderConfigurations) (*GetAudioEncoderConfigurationsResponse, error) {
	return service.GetAudioEncoderConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoAnalyticsConfigurationsContext(ctx context.Context, request *GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, error) {
	response := new(GetVideoAnalyticsConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoAnalyticsConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoAnalyticsConfigurations(request *GetVideoAnalyticsConfigurations) (*GetVideoAnalyticsConfigurationsResponse, error) {
	return service.GetVideoAnalyticsConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetMetadataConfigurationsContext(ctx context.Context, request *GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, error) {
	response := new(GetMetadataConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetMetadataConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetMetadataConfigurations(request *GetMetadataConfigurations) (*GetMetadataConfigurationsResponse, error) {
	return service.GetMetadataConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioOutputConfigurationsContext(ctx context.Context, request *GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, error) {
	response := new(GetAudioOutputConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioOutputConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioOutputConfigurations(request *GetAudioOutputConfigurations) (*GetAudioOutputConfigurationsResponse, error) {
	return service.GetAudioOutputConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioDecoderConfigurationsContext(ctx context.Context, request *GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, error) {
	response := new(GetAudioDecoderConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioDecoderConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioDecoderConfigurations(request *GetAudioDecoderConfigurations) (*GetAudioDecoderConfigurationsResponse, error) {
	return service.GetAudioDecoderConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoSourceConfigurationContext(ctx context.Context, request *GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, error) {
	response := new(GetVideoSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoSourceConfiguration(request *GetVideoSourceConfiguration) (*GetVideoSourceConfigurationResponse, error) {
	return service.GetVideoSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoEncoderConfigurationContext(ctx context.Context, request *GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, error) {
	response := new(GetVideoEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoEncoderConfiguration(request *GetVideoEncoderConfiguration) (*GetVideoEncoderConfigurationResponse, error) {
	return service.GetVideoEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioSourceConfigurationContext(ctx context.Context, request *GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, error) {
	response := new(GetAudioSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioSourceConfiguration(request *GetAudioSourceConfiguration) (*GetAudioSourceConfigurationResponse, error) {
	return service.GetAudioSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioEncoderConfigurationContext(ctx context.Context, request *GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, error) {
	response := new(GetAudioEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioEncoderConfiguration(request *GetAudioEncoderConfiguration) (*GetAudioEncoderConfigurationResponse, error) {
	return service.GetAudioEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoAnalyticsConfigurationContext(ctx context.Context, request *GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, error) {
	response := new(GetVideoAnalyticsConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoAnalyticsConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoAnalyticsConfiguration(request *GetVideoAnalyticsConfiguration) (*GetVideoAnalyticsConfigurationResponse, error) {
	return service.GetVideoAnalyticsConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetMetadataConfigurationContext(ctx context.Context, request *GetMetadataConfiguration) (*GetMetadataConfigurationResponse, error) {
	response := new(GetMetadataConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetMetadataConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetMetadataConfiguration(request *GetMetadataConfiguration) (*GetMetadataConfigurationResponse, error) {
	return service.GetMetadataConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioOutputConfigurationContext(ctx context.Context, request *GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, error) {
	response := new(GetAudioOutputConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioOutputConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioOutputConfiguration(request *GetAudioOutputConfiguration) (*GetAudioOutputConfigurationResponse, error) {
	return service.GetAudioOutputConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioDecoderConfigurationContext(ctx context.Context, request *GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, error) {
	response := new(GetAudioDecoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioDecoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioDecoderConfiguration(request *GetAudioDecoderConfiguration) (*GetAudioDecoderConfigurationResponse, error) {
	return service.GetAudioDecoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleVideoEncoderConfigurationsContext(ctx context.Context, request *GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, error) {
	response := new(GetCompatibleVideoEncoderConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleVideoEncoderConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleVideoEncoderConfigurations(request *GetCompatibleVideoEncoderConfigurations) (*GetCompatibleVideoEncoderConfigurationsResponse, error) {
	return service.GetCompatibleVideoEncoderConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleVideoSourceConfigurationsContext(ctx context.Context, request *GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, error) {
	response := new(GetCompatibleVideoSourceConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleVideoSourceConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleVideoSourceConfigurations(request *GetCompatibleVideoSourceConfigurations) (*GetCompatibleVideoSourceConfigurationsResponse, error) {
	return service.GetCompatibleVideoSourceConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleAudioEncoderConfigurationsContext(ctx context.Context, request *GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, error) {
	response := new(GetCompatibleAudioEncoderConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleAudioEncoderConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleAudioEncoderConfigurations(request *GetCompatibleAudioEncoderConfigurations) (*GetCompatibleAudioEncoderConfigurationsResponse, error) {
	return service.GetCompatibleAudioEncoderConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleAudioSourceConfigurationsContext(ctx context.Context, request *GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, error) {
	response := new(GetCompatibleAudioSourceConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleAudioSourceConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleAudioSourceConfigurations(request *GetCompatibleAudioSourceConfigurations) (*GetCompatibleAudioSourceConfigurationsResponse, error) {
	return service.GetCompatibleAudioSourceConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleVideoAnalyticsConfigurationsContext(ctx context.Context, request *GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	response := new(GetCompatibleVideoAnalyticsConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleVideoAnalyticsConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleVideoAnalyticsConfigurations(request *GetCompatibleVideoAnalyticsConfigurations) (*GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	return service.GetCompatibleVideoAnalyticsConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleMetadataConfigurationsContext(ctx context.Context, request *GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, error) {
	response := new(GetCompatibleMetadataConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleMetadataConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleMetadataConfigurations(request *GetCompatibleMetadataConfigurations) (*GetCompatibleMetadataConfigurationsResponse, error) {
	return service.GetCompatibleMetadataConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleAudioOutputConfigurationsContext(ctx context.Context, request *GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, error) {
	response := new(GetCompatibleAudioOutputConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleAudioOutputConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleAudioOutputConfigurations(request *GetCompatibleAudioOutputConfigurations) (*GetCompatibleAudioOutputConfigurationsResponse, error) {
	return service.GetCompatibleAudioOutputConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetCompatibleAudioDecoderConfigurationsContext(ctx context.Context, request *GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, error) {
	response := new(GetCompatibleAudioDecoderConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetCompatibleAudioDecoderConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetCompatibleAudioDecoderConfigurations(request *GetCompatibleAudioDecoderConfigurations) (*GetCompatibleAudioDecoderConfigurationsResponse, error) {
	return service.GetCompatibleAudioDecoderConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *media) SetVideoSourceConfigurationContext(ctx context.Context, request *SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, error) {
	response := new(SetVideoSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetVideoSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetVideoSourceConfiguration(request *SetVideoSourceConfiguration) (*SetVideoSourceConfigurationResponse, error) {
	return service.SetVideoSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetVideoEncoderConfigurationContext(ctx context.Context, request *SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, error) {
	response := new(SetVideoEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetVideoEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetVideoEncoderConfiguration(request *SetVideoEncoderConfiguration) (*SetVideoEncoderConfigurationResponse, error) {
	return service.SetVideoEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetAudioSourceConfigurationContext(ctx context.Context, request *SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, error) {
	response := new(SetAudioSourceConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetAudioSourceConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetAudioSourceConfiguration(request *SetAudioSourceConfiguration) (*SetAudioSourceConfigurationResponse, error) {
	return service.SetAudioSourceConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetAudioEncoderConfigurationContext(ctx context.Context, request *SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, error) {
	response := new(SetAudioEncoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetAudioEncoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetAudioEncoderConfiguration(request *SetAudioEncoderConfiguration) (*SetAudioEncoderConfigurationResponse, error) {
	return service.SetAudioEncoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetVideoAnalyticsConfigurationContext(ctx context.Context, request *SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, error) {
	response := new(SetVideoAnalyticsConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetVideoAnalyticsConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetVideoAnalyticsConfiguration(request *SetVideoAnalyticsConfiguration) (*SetVideoAnalyticsConfigurationResponse, error) {
	return service.SetVideoAnalyticsConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetMetadataConfigurationContext(ctx context.Context, request *SetMetadataConfiguration) (*SetMetadataConfigurationResponse, error) {
	response := new(SetMetadataConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetMetadataConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetMetadataConfiguration(request *SetMetadataConfiguration) (*SetMetadataConfigurationResponse, error) {
	return service.SetMetadataConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetAudioOutputConfigurationContext(ctx context.Context, request *SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, error) {
	response := new(SetAudioOutputConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetAudioOutputConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetAudioOutputConfiguration(request *SetAudioOutputConfiguration) (*SetAudioOutputConfigurationResponse, error) {
	return service.SetAudioOutputConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) SetAudioDecoderConfigurationContext(ctx context.Context, request *SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, error) {
	response := new(SetAudioDecoderConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetAudioDecoderConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetAudioDecoderConfiguration(request *SetAudioDecoderConfiguration) (*SetAudioDecoderConfigurationResponse, error) {
	return service.SetAudioDecoderConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoSourceConfigurationOptionsContext(ctx context.Context, request *GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, error) {
	response := new(GetVideoSourceConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoSourceConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoSourceConfigurationOptions(request *GetVideoSourceConfigurationOptions) (*GetVideoSourceConfigurationOptionsResponse, error) {
	return service.GetVideoSourceConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoEncoderConfigurationOptionsContext(ctx context.Context, request *GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, error) {
	response := new(GetVideoEncoderConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoEncoderConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoEncoderConfigurationOptions(request *GetVideoEncoderConfigurationOptions) (*GetVideoEncoderConfigurationOptionsResponse, error) {
	return service.GetVideoEncoderConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioSourceConfigurationOptionsContext(ctx context.Context, request *GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, error) {
	response := new(GetAudioSourceConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioSourceConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioSourceConfigurationOptions(request *GetAudioSourceConfigurationOptions) (*GetAudioSourceConfigurationOptionsResponse, error) {
	return service.GetAudioSourceConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioEncoderConfigurationOptionsContext(ctx context.Context, request *GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, error) {
	response := new(GetAudioEncoderConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioEncoderConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioEncoderConfigurationOptions(request *GetAudioEncoderConfigurationOptions) (*GetAudioEncoderConfigurationOptionsResponse, error) {
	return service.GetAudioEncoderConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetMetadataConfigurationOptionsContext(ctx context.Context, request *GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, error) {
	response := new(GetMetadataConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetMetadataConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetMetadataConfigurationOptions(request *GetMetadataConfigurationOptions) (*GetMetadataConfigurationOptionsResponse, error) {
	return service.GetMetadataConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioOutputConfigurationOptionsContext(ctx context.Context, request *GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, error) {
	response := new(GetAudioOutputConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioOutputConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioOutputConfigurationOptions(request *GetAudioOutputConfigurationOptions) (*GetAudioOutputConfigurationOptionsResponse, error) {
	return service.GetAudioOutputConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetAudioDecoderConfigurationOptionsContext(ctx context.Context, request *GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, error) {
	response := new(GetAudioDecoderConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetAudioDecoderConfigurationOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetAudioDecoderConfigurationOptions(request *GetAudioDecoderConfigurationOptions) (*GetAudioDecoderConfigurationOptionsResponse, error) {
	return service.GetAudioDecoderConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetGuaranteedNumberOfVideoEncoderInstancesContext(ctx context.Context, request *GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	response := new(GetGuaranteedNumberOfVideoEncoderInstancesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetGuaranteedNumberOfVideoEncoderInstances", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetGuaranteedNumberOfVideoEncoderInstances(request *GetGuaranteedNumberOfVideoEncoderInstances) (*GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	return service.GetGuaranteedNumberOfVideoEncoderInstancesContext(
		context.Background(),
		request,
	)
}

func (service *media) GetStreamUriContext(ctx context.Context, request *GetStreamUri) (*GetStreamUriResponse, error) {
	response := new(GetStreamUriResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetStreamUri", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetStreamUri(request *GetStreamUri) (*GetStreamUriResponse, error) {
	return service.GetStreamUriContext(
		context.Background(),
		request,
	)
}

func (service *media) StartMulticastStreamingContext(ctx context.Context, request *StartMulticastStreaming) (*StartMulticastStreamingResponse, error) {
	response := new(StartMulticastStreamingResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/StartMulticastStreaming", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) StartMulticastStreaming(request *StartMulticastStreaming) (*StartMulticastStreamingResponse, error) {
	return service.StartMulticastStreamingContext(
		context.Background(),
		request,
	)
}

func (service *media) StopMulticastStreamingContext(ctx context.Context, request *StopMulticastStreaming) (*StopMulticastStreamingResponse, error) {
	response := new(StopMulticastStreamingResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/StopMulticastStreaming", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) StopMulticastStreaming(request *StopMulticastStreaming) (*StopMulticastStreamingResponse, error) {
	return service.StopMulticastStreamingContext(
		context.Background(),
		request,
	)
}

func (service *media) SetSynchronizationPointContext(ctx context.Context, request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	response := new(SetSynchronizationPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetSynchronizationPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetSynchronizationPoint(request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	return service.SetSynchronizationPointContext(
		context.Background(),
		request,
	)
}

func (service *media) GetSnapshotUriContext(ctx context.Context, request *GetSnapshotUri) (*GetSnapshotUriResponse, error) {
	response := new(GetSnapshotUriResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetSnapshotUri", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetSnapshotUri(request *GetSnapshotUri) (*GetSnapshotUriResponse, error) {
	return service.GetSnapshotUriContext(
		context.Background(),
		request,
	)
}

func (service *media) GetVideoSourceModesContext(ctx context.Context, request *GetVideoSourceModes) (*GetVideoSourceModesResponse, error) {
	response := new(GetVideoSourceModesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetVideoSourceModes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetVideoSourceModes(request *GetVideoSourceModes) (*GetVideoSourceModesResponse, error) {
	return service.GetVideoSourceModesContext(
		context.Background(),
		request,
	)
}

func (service *media) SetVideoSourceModeContext(ctx context.Context, request *SetVideoSourceMode) (*SetVideoSourceModeResponse, error) {
	response := new(SetVideoSourceModeResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetVideoSourceMode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetVideoSourceMode(request *SetVideoSourceMode) (*SetVideoSourceModeResponse, error) {
	return service.SetVideoSourceModeContext(
		context.Background(),
		request,
	)
}

func (service *media) GetOSDsContext(ctx context.Context, request *GetOSDs) (*GetOSDsResponse, error) {
	response := new(GetOSDsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetOSDs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetOSDs(request *GetOSDs) (*GetOSDsResponse, error) {
	return service.GetOSDsContext(
		context.Background(),
		request,
	)
}

func (service *media) GetOSDContext(ctx context.Context, request *GetOSD) (*GetOSDResponse, error) {
	response := new(GetOSDResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetOSD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetOSD(request *GetOSD) (*GetOSDResponse, error) {
	return service.GetOSDContext(
		context.Background(),
		request,
	)
}

func (service *media) GetOSDOptionsContext(ctx context.Context, request *GetOSDOptions) (*GetOSDOptionsResponse, error) {
	response := new(GetOSDOptionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/GetOSDOptions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) GetOSDOptions(request *GetOSDOptions) (*GetOSDOptionsResponse, error) {
	return service.GetOSDOptionsContext(
		context.Background(),
		request,
	)
}

func (service *media) SetOSDContext(ctx context.Context, request *SetOSD) (*SetOSDResponse, error) {
	response := new(SetOSDResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/SetOSD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) SetOSD(request *SetOSD) (*SetOSDResponse, error) {
	return service.SetOSDContext(
		context.Background(),
		request,
	)
}

func (service *media) CreateOSDContext(ctx context.Context, request *CreateOSD) (*CreateOSDResponse, error) {
	response := new(CreateOSDResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/CreateOSD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) CreateOSD(request *CreateOSD) (*CreateOSDResponse, error) {
	return service.CreateOSDContext(
		context.Background(),
		request,
	)
}

func (service *media) DeleteOSDContext(ctx context.Context, request *DeleteOSD) (*DeleteOSDResponse, error) {
	response := new(DeleteOSDResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/media/wsdl/DeleteOSD", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *media) DeleteOSD(request *DeleteOSD) (*DeleteOSDResponse, error) {
	return service.DeleteOSDContext(
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

// AnySimpleType type
type AnySimpleType string

// String type
type String string
