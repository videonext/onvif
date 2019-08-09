package event

import (
	"context"
	"encoding/xml"
	"github.com/videonext/onvif/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// GetServiceCapabilities type
type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capabilities for the event service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"Capabilities,omitempty"`
}

// CreatePullPointSubscription type
type CreatePullPointSubscription struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl CreatePullPointSubscription"`

	// Optional XPATH expression to select specific topics.
	Filter FilterType `xml:"Filter,omitempty"`

	// Initial termination time.
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
}

// CreatePullPointSubscriptionResponse type
type CreatePullPointSubscriptionResponse struct {
	XMLName xml.Name `xml:"CreatePullPointSubscriptionResponse"`

	// Endpoint reference of the subscription to be used for pulling the messages.
	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference,omitempty"`

	CurrentTime CurrentTime `xml:"CurrentTime,omitempty"`

	TerminationTime TerminationTime `xml:"TerminationTime,omitempty"`
}

// PullMessages type
type PullMessages struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl PullMessages"`

	// Maximum time to block until this method returns.
	Timeout Duration `xml:"http://www.onvif.org/ver10/schema Timeout,omitempty"`

	// Upper limit for the number of messages to return at once. A server implementation may decide to return less messages.
	MessageLimit int32 `xml:"http://www.onvif.org/ver10/schema MessageLimit,omitempty"`
}

// PullMessagesResponse type
type PullMessagesResponse struct {
	XMLName xml.Name `xml:"PullMessagesResponse"`

	// The date and time when the messages have been delivered by the web server to the client.
	CurrentTime string `xml:"CurrentTime,omitempty"`

	// Date time when the PullPoint will be shut down without further pull requests.
	TerminationTime string `xml:"TerminationTime,omitempty"`

	NotificationMessage []NotificationMessage `xml:"NotificationMessage,omitempty"`
}

// Seek type
type Seek struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl Seek"`

	// The date and time to match against stored messages.
	UtcTime string `xml:"http://www.onvif.org/ver10/schema UtcTime,omitempty"`

	// Reverse the pull direction of PullMessages.
	Reverse bool `xml:"http://www.onvif.org/ver10/events/wsdl Reverse,omitempty"`
}

// SeekResponse type
type SeekResponse struct {
	XMLName xml.Name `xml:"SeekResponse"`
}

// SetSynchronizationPoint type
type SetSynchronizationPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl SetSynchronizationPoint"`
}

// SetSynchronizationPointResponse type
type SetSynchronizationPointResponse struct {
	XMLName xml.Name `xml:"SetSynchronizationPointResponse"`
}

// GetEventProperties type
type GetEventProperties struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl GetEventProperties"`
}

// GetEventPropertiesResponse type
type GetEventPropertiesResponse struct {
	XMLName xml.Name `xml:"GetEventPropertiesResponse"`

	// List of topic namespaces supported.
	TopicNamespaceLocation []AnyURI `xml:"TopicNamespaceLocation,omitempty"`

	FixedTopicSet FixedTopicSet `xml:"FixedTopicSet,omitempty"`

	TopicSet TopicSet `xml:"TopicSet,omitempty"`

	TopicExpressionDialect []TopicExpressionDialect `xml:"TopicExpressionDialect,omitempty"`

	//
	// Defines the XPath function set supported for message content filtering.
	// The following MessageContentFilterDialects should be returned if a device supports the message content filtering:
	//
	// A device that does not support any MessageContentFilterDialect returns a single empty url.
	//
	MessageContentFilterDialect []AnyURI `xml:"MessageContentFilterDialect,omitempty"`

	//
	// Optional ProducerPropertiesDialects. Refer to  for advanced filtering.
	ProducerPropertiesFilterDialect []AnyURI `xml:"ProducerPropertiesFilterDialect,omitempty"`

	//
	// The Message Content Description Language allows referencing
	// of vendor-specific types. In order to ease the integration of such types into a client application,
	// the GetEventPropertiesResponse shall list all URI locations to schema files whose types are
	// used in the description of notifications, with MessageContentSchemaLocation elements.
	// This list shall at least contain the URI of the ONVIF schema file.
	MessageContentSchemaLocation []AnyURI `xml:"MessageContentSchemaLocation,omitempty"`
}

// Capabilities type
type Capabilities struct {

	// Indicates that the WS Subscription policy is supported.

	WSSubscriptionPolicySupport bool `xml:"http://www.onvif.org/ver10/events/wsdl WSSubscriptionPolicySupport,attr,omitempty"`

	// Indicates that the WS Pull Point is supported.

	WSPullPointSupport bool `xml:"http://www.onvif.org/ver10/events/wsdl WSPullPointSupport,attr,omitempty"`

	// Indicates that the WS Pausable Subscription Manager Interface is supported.

	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"http://www.onvif.org/ver10/events/wsdl WSPausableSubscriptionManagerInterfaceSupport,attr,omitempty"`

	// Maximum number of supported notification producers as defined by WS-BaseNotification.

	MaxNotificationProducers int32 `xml:"http://www.onvif.org/ver10/schema MaxNotificationProducers,attr,omitempty"`

	// Maximum supported number of notification pull points.

	MaxPullPoints int32 `xml:"http://www.onvif.org/ver10/schema MaxPullPoints,attr,omitempty"`

	// Indication if the device supports persistent notification storage.

	PersistentNotificationStorage bool `xml:"http://www.onvif.org/ver10/events/wsdl PersistentNotificationStorage,attr,omitempty"`
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

// CreatePullPoint type
type CreatePullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 CreatePullPoint"`
}

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

// EventPortType type
type EventPortType interface {

	/* Returns the capabilities of the event service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	// Error can be either of the following types:
	//
	//   - ResourceUnknownFault
	//   - InvalidFilterFault
	//   - TopicExpressionDialectUnknownFault
	//   - InvalidTopicExpressionFault
	//   - TopicNotSupportedFault
	//   - InvalidProducerPropertiesExpressionFault
	//   - InvalidMessageContentExpressionFault
	//   - UnacceptableInitialTerminationTimeFault
	//   - UnrecognizedPolicyRequestFault
	//   - UnsupportedPolicyRequestFault
	//   - NotifyMessageNotSupportedFault
	//   - SubscribeCreationFailedFault
	/* This method returns a PullPointSubscription that can be polled using PullMessages.
	This message contains the same elements as the SubscriptionRequest of the WS-BaseNotification without the ConsumerReference.
	If no Filter is specified the pullpoint notifies all occurring events to the client.
	This method is mandatory. */
	CreatePullPointSubscription(request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error)

	CreatePullPointSubscriptionContext(ctx context.Context, request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error)

	/* The WS-BaseNotification specification defines a set of OPTIONAL WS-ResouceProperties.
	This specification does not require the implementation of the WS-ResourceProperty interface.
	Instead, the subsequent direct interface shall be implemented by an ONVIF compliant device
	in order to provide information about the FilterDialects, Schema files and topics supported by
	the device. */
	GetEventProperties(request *GetEventProperties) (*GetEventPropertiesResponse, error)

	GetEventPropertiesContext(ctx context.Context, request *GetEventProperties) (*GetEventPropertiesResponse, error)
}

// eventPortType type
type eventPortType struct {
	client *soap.Client
	xaddr  string
}

func NewEventPortType(client *soap.Client, xaddr string) EventPortType {
	return &eventPortType{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *eventPortType) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *eventPortType) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *eventPortType) CreatePullPointSubscriptionContext(ctx context.Context, request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error) {
	response := new(CreatePullPointSubscriptionResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/CreatePullPointSubscription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *eventPortType) CreatePullPointSubscription(request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error) {
	return service.CreatePullPointSubscriptionContext(
		context.Background(),
		request,
	)
}

func (service *eventPortType) GetEventPropertiesContext(ctx context.Context, request *GetEventProperties) (*GetEventPropertiesResponse, error) {
	response := new(GetEventPropertiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/GetEventProperties", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *eventPortType) GetEventProperties(request *GetEventProperties) (*GetEventPropertiesResponse, error) {
	return service.GetEventPropertiesContext(
		context.Background(),
		request,
	)
}

// PullPointSubscription type
type PullPointSubscription interface {

	// Error can be either of the following types:
	//
	//   - PullMessagesFaultResponse
	/*
		This method pulls one or more messages from a PullPoint.
		The device shall provide the following PullMessages command for all SubscriptionManager
		endpoints returned by the CreatePullPointSubscription command. This method shall not wait until
		the requested number of messages is available but return as soon as at least one message is available.
		The command shall at least support a Timeout of one minute. In case a device supports retrieval of less messages
		than requested it shall return these without generating a fault. */
	PullMessages(request *PullMessages) (*PullMessagesResponse, error)

	PullMessagesContext(ctx context.Context, request *PullMessages) (*PullMessagesResponse, error)

	/*
		This method readjusts the pull pointer into the past.
		A device supporting persistent notification storage shall provide the
		following Seek command for all SubscriptionManager endpoints returned by
		the CreatePullPointSubscription command. The optional Reverse argument can
		be used to reverse the pull direction of the PullMessages command.
		The UtcTime argument will be matched against the UtcTime attribute on a
		NotificationMessage.
	*/
	Seek(request *Seek) (*SeekResponse, error)

	SeekContext(ctx context.Context, request *Seek) (*SeekResponse, error)

	/* Properties inform a client about property creation, changes and
	deletion in a uniform way. When a client wants to synchronize its properties with the
	properties of the device, it can request a synchronization point which repeats the current
	status of all properties to which a client has subscribed. The PropertyOperation of all
	produced notifications is set to “Initialized”. The Synchronization Point is
	requested directly from the SubscriptionManager which was returned in either the
	SubscriptionResponse or in the CreatePullPointSubscriptionResponse. The property update is
	transmitted via the notification transportation of the notification interface. This method is mandatory.
	*/
	SetSynchronizationPoint(request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	SetSynchronizationPointContext(ctx context.Context, request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	// Error can be either of the following types:
	//
	//   - ResourceUnknownFault
	//   - UnableToDestroySubscriptionFault
	/* The device shall provide the following Unsubscribe command for all SubscriptionManager endpoints returned by the CreatePullPointSubscription command.
	This command shall terminate the lifetime of a pull point.
	*/
	Unsubscribe() error

	UnsubscribeContext(ctx context.Context) error
}

// pullPointSubscription type
type pullPointSubscription struct {
	client *soap.Client
	xaddr  string
}

func NewPullPointSubscription(client *soap.Client, xaddr string) PullPointSubscription {
	return &pullPointSubscription{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *pullPointSubscription) PullMessagesContext(ctx context.Context, request *PullMessages) (*PullMessagesResponse, error) {
	response := new(PullMessagesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/PullMessages", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pullPointSubscription) PullMessages(request *PullMessages) (*PullMessagesResponse, error) {
	return service.PullMessagesContext(
		context.Background(),
		request,
	)
}

func (service *pullPointSubscription) SeekContext(ctx context.Context, request *Seek) (*SeekResponse, error) {
	response := new(SeekResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/Seek", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pullPointSubscription) Seek(request *Seek) (*SeekResponse, error) {
	return service.SeekContext(
		context.Background(),
		request,
	)
}

func (service *pullPointSubscription) SetSynchronizationPointContext(ctx context.Context, request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	response := new(SetSynchronizationPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/SetSynchronizationPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pullPointSubscription) SetSynchronizationPoint(request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	return service.SetSynchronizationPointContext(
		context.Background(),
		request,
	)
}

func (service *pullPointSubscription) UnsubscribeContext(ctx context.Context) error {

	err := service.client.CallContext(ctx, service.xaddr, "''", nil, struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func (service *pullPointSubscription) Unsubscribe() error {
	return service.UnsubscribeContext(
		context.Background(),
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
