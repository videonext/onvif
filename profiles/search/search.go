package search

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
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetServiceCapabilities"`
}

// GetServiceCapabilitiesResponse type
type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"GetServiceCapabilitiesResponse"`

	// The capabilities for the search service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"Capabilities,omitempty"`
}

// GetRecordingSummary type
type GetRecordingSummary struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetRecordingSummary"`
}

// GetRecordingSummaryResponse type
type GetRecordingSummaryResponse struct {
	XMLName xml.Name `xml:"GetRecordingSummaryResponse"`

	Summary RecordingSummary `xml:"Summary,omitempty"`
}

// GetRecordingInformation type
type GetRecordingInformation struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetRecordingInformation"`

	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`
}

// GetRecordingInformationResponse type
type GetRecordingInformationResponse struct {
	XMLName xml.Name `xml:"GetRecordingInformationResponse"`

	RecordingInformation RecordingInformation `xml:"RecordingInformation,omitempty"`
}

// GetMediaAttributes type
type GetMediaAttributes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetMediaAttributes"`

	RecordingTokens []RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingTokens,omitempty"`

	Time string `xml:"http://www.onvif.org/ver10/schema Time,omitempty"`
}

// GetMediaAttributesResponse type
type GetMediaAttributesResponse struct {
	XMLName xml.Name `xml:"GetMediaAttributesResponse"`

	MediaAttributes []MediaAttributes `xml:"MediaAttributes,omitempty"`
}

// FindRecordings type
type FindRecordings struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl FindRecordings"`

	// Scope defines the dataset to consider for this search.
	Scope SearchScope `xml:"http://www.onvif.org/ver10/search/wsdl Scope,omitempty"`

	// The search will be completed after this many matches. If not specified, the search will continue until reaching the endpoint or until the session expires.
	MaxMatches int32 `xml:"http://www.onvif.org/ver10/schema MaxMatches,omitempty"`

	// The time the search session will be kept alive after responding to this and subsequent requests. A device shall support at least values up to ten seconds.
	KeepAliveTime Duration `xml:"http://www.onvif.org/ver10/schema KeepAliveTime,omitempty"`
}

// FindRecordingsResponse type
type FindRecordingsResponse struct {
	XMLName xml.Name `xml:"FindRecordingsResponse"`

	SearchToken JobToken `xml:"SearchToken,omitempty"`
}

// GetRecordingSearchResults type
type GetRecordingSearchResults struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetRecordingSearchResults"`

	// The search session to get results from.
	SearchToken JobToken `xml:"http://www.onvif.org/ver10/search/wsdl SearchToken,omitempty"`

	// The minimum number of results to return in one response.
	MinResults int32 `xml:"http://www.onvif.org/ver10/schema MinResults,omitempty"`

	// The maximum number of results to return in one response.
	MaxResults int32 `xml:"http://www.onvif.org/ver10/schema MaxResults,omitempty"`

	// The maximum time before responding to the request, even if the MinResults parameter is not fulfilled.
	WaitTime Duration `xml:"http://www.onvif.org/ver10/schema WaitTime,omitempty"`
}

// GetRecordingSearchResultsResponse type
type GetRecordingSearchResultsResponse struct {
	XMLName xml.Name `xml:"GetRecordingSearchResultsResponse"`

	ResultList FindRecordingResultList `xml:"ResultList,omitempty"`
}

// FindEvents type
type FindEvents struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl FindEvents"`

	// The point of time where the search will start.
	StartPoint string `xml:"http://www.onvif.org/ver10/schema StartPoint,omitempty"`

	// The point of time where the search will stop. This can be a time before the StartPoint, in which case the search is performed backwards in time.
	EndPoint string `xml:"http://www.onvif.org/ver10/schema EndPoint,omitempty"`

	Scope SearchScope `xml:"http://www.onvif.org/ver10/search/wsdl Scope,omitempty"`

	SearchFilter EventFilter `xml:"http://www.onvif.org/ver10/search/wsdl SearchFilter,omitempty"`

	// Setting IncludeStartState to true means that the server should return virtual events representing the start state for any recording included in the scope. Start state events are limited to the topics defined in the SearchFilter that have the IsProperty flag set to true.
	IncludeStartState bool `xml:"http://www.onvif.org/ver10/search/wsdl IncludeStartState,omitempty"`

	// The search will be completed after this many matches. If not specified, the search will continue until reaching the endpoint or until the session expires.
	MaxMatches int32 `xml:"http://www.onvif.org/ver10/schema MaxMatches,omitempty"`

	// The time the search session will be kept alive after responding to this and subsequent requests. A device shall support at least values up to ten seconds.
	KeepAliveTime Duration `xml:"http://www.onvif.org/ver10/schema KeepAliveTime,omitempty"`
}

// FindEventsResponse type
type FindEventsResponse struct {
	XMLName xml.Name `xml:"FindEventsResponse"`

	// A unique reference to the search session created by this request.
	SearchToken JobToken `xml:"SearchToken,omitempty"`
}

// GetEventSearchResults type
type GetEventSearchResults struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetEventSearchResults"`

	// The search session to get results from.
	SearchToken JobToken `xml:"http://www.onvif.org/ver10/search/wsdl SearchToken,omitempty"`

	// The minimum number of results to return in one response.
	MinResults int32 `xml:"http://www.onvif.org/ver10/schema MinResults,omitempty"`

	// The maximum number of results to return in one response.
	MaxResults int32 `xml:"http://www.onvif.org/ver10/schema MaxResults,omitempty"`

	// The maximum time before responding to the request, even if the MinResults parameter is not fulfilled.
	WaitTime Duration `xml:"http://www.onvif.org/ver10/schema WaitTime,omitempty"`
}

// GetEventSearchResultsResponse type
type GetEventSearchResultsResponse struct {
	XMLName xml.Name `xml:"GetEventSearchResultsResponse"`

	ResultList FindEventResultList `xml:"ResultList,omitempty"`
}

// FindPTZPosition type
type FindPTZPosition struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl FindPTZPosition"`

	// The point of time where the search will start.
	StartPoint string `xml:"http://www.onvif.org/ver10/schema StartPoint,omitempty"`

	// The point of time where the search will stop. This can be a time before the StartPoint, in which case the search is performed backwards in time.
	EndPoint string `xml:"http://www.onvif.org/ver10/schema EndPoint,omitempty"`

	Scope SearchScope `xml:"http://www.onvif.org/ver10/search/wsdl Scope,omitempty"`

	SearchFilter PTZPositionFilter `xml:"http://www.onvif.org/ver10/search/wsdl SearchFilter,omitempty"`

	// The search will be completed after this many matches. If not specified, the search will continue until reaching the endpoint or until the session expires.
	MaxMatches int32 `xml:"http://www.onvif.org/ver10/schema MaxMatches,omitempty"`

	// The time the search session will be kept alive after responding to this and subsequent requests. A device shall support at least values up to ten seconds.
	KeepAliveTime Duration `xml:"http://www.onvif.org/ver10/schema KeepAliveTime,omitempty"`
}

// FindPTZPositionResponse type
type FindPTZPositionResponse struct {
	XMLName xml.Name `xml:"FindPTZPositionResponse"`

	// A unique reference to the search session created by this request.
	SearchToken JobToken `xml:"SearchToken,omitempty"`
}

// GetPTZPositionSearchResults type
type GetPTZPositionSearchResults struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetPTZPositionSearchResults"`

	// The search session to get results from.
	SearchToken JobToken `xml:"http://www.onvif.org/ver10/search/wsdl SearchToken,omitempty"`

	// The minimum number of results to return in one response.
	MinResults int32 `xml:"http://www.onvif.org/ver10/schema MinResults,omitempty"`

	// The maximum number of results to return in one response.
	MaxResults int32 `xml:"http://www.onvif.org/ver10/schema MaxResults,omitempty"`

	// The maximum time before responding to the request, even if the MinResults parameter is not fulfilled.
	WaitTime Duration `xml:"http://www.onvif.org/ver10/schema WaitTime,omitempty"`
}

// GetPTZPositionSearchResultsResponse type
type GetPTZPositionSearchResultsResponse struct {
	XMLName xml.Name `xml:"GetPTZPositionSearchResultsResponse"`

	ResultList FindPTZPositionResultList `xml:"ResultList,omitempty"`
}

// FindMetadata type
type FindMetadata struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl FindMetadata"`

	// The point of time where the search will start.
	StartPoint string `xml:"http://www.onvif.org/ver10/schema StartPoint,omitempty"`

	// The point of time where the search will stop. This can be a time before the StartPoint, in which case the search is performed backwards in time.
	EndPoint string `xml:"http://www.onvif.org/ver10/schema EndPoint,omitempty"`

	Scope SearchScope `xml:"http://www.onvif.org/ver10/search/wsdl Scope,omitempty"`

	MetadataFilter MetadataFilter `xml:"http://www.onvif.org/ver10/search/wsdl MetadataFilter,omitempty"`

	// The search will be completed after this many matches. If not specified, the search will continue until reaching the endpoint or until the session expires.
	MaxMatches int32 `xml:"http://www.onvif.org/ver10/schema MaxMatches,omitempty"`

	// The time the search session will be kept alive after responding to this and subsequent requests. A device shall support at least values up to ten seconds.
	KeepAliveTime Duration `xml:"http://www.onvif.org/ver10/schema KeepAliveTime,omitempty"`
}

// FindMetadataResponse type
type FindMetadataResponse struct {
	XMLName xml.Name `xml:"FindMetadataResponse"`

	// A unique reference to the search session created by this request.
	SearchToken JobToken `xml:"SearchToken,omitempty"`
}

// GetMetadataSearchResults type
type GetMetadataSearchResults struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetMetadataSearchResults"`

	// The search session to get results from.
	SearchToken JobToken `xml:"http://www.onvif.org/ver10/search/wsdl SearchToken,omitempty"`

	// The minimum number of results to return in one response.
	MinResults int32 `xml:"http://www.onvif.org/ver10/schema MinResults,omitempty"`

	// The maximum number of results to return in one response.
	MaxResults int32 `xml:"http://www.onvif.org/ver10/schema MaxResults,omitempty"`

	// The maximum time before responding to the request, even if the MinResults parameter is not fulfilled.
	WaitTime Duration `xml:"http://www.onvif.org/ver10/schema WaitTime,omitempty"`
}

// GetMetadataSearchResultsResponse type
type GetMetadataSearchResultsResponse struct {
	XMLName xml.Name `xml:"GetMetadataSearchResultsResponse"`

	ResultList FindMetadataResultList `xml:"ResultList,omitempty"`
}

// GetSearchState type
type GetSearchState struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl GetSearchState"`

	// The search session to get the state from.
	SearchToken JobToken `xml:"http://www.onvif.org/ver10/search/wsdl SearchToken,omitempty"`
}

// GetSearchStateResponse type
type GetSearchStateResponse struct {
	XMLName xml.Name `xml:"GetSearchStateResponse"`

	State SearchState `xml:"State,omitempty"`
}

// EndSearch type
type EndSearch struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/search/wsdl EndSearch"`

	// The search session to end.
	SearchToken JobToken `xml:"http://www.onvif.org/ver10/search/wsdl SearchToken,omitempty"`
}

// EndSearchResponse type
type EndSearchResponse struct {
	XMLName xml.Name `xml:"EndSearchResponse"`

	// The point of time the search had reached when it was ended. It is equal to the EndPoint specified in Find-operation if the search was completed.
	Endpoint string `xml:"Endpoint,omitempty"`
}

// Capabilities type
type Capabilities struct {
	MetadataSearch bool `xml:"http://www.onvif.org/ver10/search/wsdl MetadataSearch,attr,omitempty"`

	// Indicates support for general virtual property events in the FindEvents method.

	GeneralStartEvents bool `xml:"http://www.onvif.org/ver10/search/wsdl GeneralStartEvents,attr,omitempty"`
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

// Removed LocationEntity by fixgen.py

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

	Item []string `xml:"http://www.onvif.org/ver10/schema Item,omitempty"`
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
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema SourceToken,omitempty"`

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
	VideoSourceTokensAvailable []ReferenceToken `xml:"http://www.onvif.org/ver10/schema VideoSourceTokensAvailable,omitempty"`

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

	Reboot bool `xml:"http://www.onvif.org/ver10/search/wsdl Reboot,attr,omitempty"`
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

	GuaranteedFrameRate bool `xml:"http://www.onvif.org/ver10/search/wsdl GuaranteedFrameRate,attr,omitempty"`
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
	SourceToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema SourceToken,omitempty"`
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
	Analytics bool `xml:"http://www.onvif.org/ver10/search/wsdl Analytics,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast MulticastConfiguration `xml:"http://www.onvif.org/ver10/schema Multicast,omitempty"`

	// The rtsp session timeout for the related audio stream (when using Media2 Service, this value is deprecated and ignored)
	SessionTimeout Duration `xml:"http://www.onvif.org/ver10/schema SessionTimeout,omitempty"`

	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"http://www.onvif.org/ver10/schema AnalyticsEngineConfiguration,omitempty"`

	Extension MetadataConfigurationExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	// Optional parameter to configure compression type of Metadata payload. Use values from enumeration MetadataCompressionType.

	CompressionType string `xml:"http://www.onvif.org/ver10/schema CompressionType,attr,omitempty"`

	// Optional parameter to configure if the metadata stream shall contain the Geo Location coordinates of each target.

	GeoLocation bool `xml:"http://www.onvif.org/ver10/search/wsdl GeoLocation,attr,omitempty"`
}

// MetadataConfigurationExtension type
type MetadataConfigurationExtension struct {
}

// PTZFilter type
type PTZFilter struct {

	// True if the metadata stream shall contain the PTZ status (IDLE, MOVING or UNKNOWN)
	Status bool `xml:"http://www.onvif.org/ver10/search/wsdl Status,omitempty"`

	// True if the metadata stream shall contain the PTZ position
	Position bool `xml:"http://www.onvif.org/ver10/search/wsdl Position,omitempty"`
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

	GeoLocation bool `xml:"http://www.onvif.org/ver10/search/wsdl GeoLocation,attr,omitempty"`
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
	PanTiltStatusSupported bool `xml:"http://www.onvif.org/ver10/search/wsdl PanTiltStatusSupported,omitempty"`

	// True if the device is able to stream zoom status inforamtion.
	ZoomStatusSupported bool `xml:"http://www.onvif.org/ver10/search/wsdl ZoomStatusSupported,omitempty"`

	// True if the device is able to stream the pan or tilt position.
	PanTiltPositionSupported bool `xml:"http://www.onvif.org/ver10/search/wsdl PanTiltPositionSupported,omitempty"`

	// True if the device is able to stream zoom position information.
	ZoomPositionSupported bool `xml:"http://www.onvif.org/ver10/search/wsdl ZoomPositionSupported,omitempty"`

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
	OutputToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema OutputToken,omitempty"`

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
	Address IPAddress `xml:"http://www.onvif.org/ver10/schema Address,omitempty"`

	// The RTP mutlicast destination port. A device may support RTCP. In this case the port value shall be even to allow the corresponding RTCP stream to be mapped to the next higher (odd) destination port number as defined in the RTSP specification.
	Port int32 `xml:"http://www.onvif.org/ver10/schema Port,omitempty"`

	// In case of IPv6 the TTL value is assumed as the hop limit. Note that for IPV6 and administratively scoped IPv4 multicast the primary use for hop limit / TTL is to prevent packets from (endlessly) circulating and not limiting scope. In these cases the address contains the scope.
	TTL int32 `xml:"http://www.onvif.org/ver10/schema TTL,omitempty"`

	// Read only property signalling that streaming is persistant. Use the methods StartMulticastStreaming and StopMulticastStreaming to switch its state.
	AutoStart bool `xml:"http://www.onvif.org/ver10/search/wsdl AutoStart,omitempty"`
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
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

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
	AutoNegotiation bool `xml:"http://www.onvif.org/ver10/search/wsdl AutoNegotiation,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

	// IPv6 configuration.
	Config IPv6Configuration `xml:"http://www.onvif.org/ver10/schema Config,omitempty"`
}

// IPv4NetworkInterface type
type IPv4NetworkInterface struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

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
	DHCP bool `xml:"http://www.onvif.org/ver10/search/wsdl DHCP,omitempty"`
}

// IPv6Configuration type
type IPv6Configuration struct {

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"http://www.onvif.org/ver10/search/wsdl AcceptRouterAdvert,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

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
	FromDHCP bool `xml:"http://www.onvif.org/ver10/search/wsdl FromDHCP,omitempty"`

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
	FromDHCP bool `xml:"http://www.onvif.org/ver10/search/wsdl FromDHCP,omitempty"`

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
	FromDHCP bool `xml:"http://www.onvif.org/ver10/search/wsdl FromDHCP,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

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
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"http://www.onvif.org/ver10/search/wsdl AcceptRouterAdvert,omitempty"`

	// List of manually added IPv6 addresses.
	Manual []PrefixedIPv6Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// DHCP configuration.
	DHCP IPv6DHCPConfiguration `xml:"http://www.onvif.org/ver10/schema DHCP,omitempty"`
}

// IPv4NetworkInterfaceSetConfiguration type
type IPv4NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

	// List of manually added IPv4 addresses.
	Manual []PrefixedIPv4Address `xml:"http://www.onvif.org/ver10/schema Manual,omitempty"`

	// Indicates whether or not DHCP is used.
	DHCP bool `xml:"http://www.onvif.org/ver10/search/wsdl DHCP,omitempty"`
}

// Removed NetworkGateway by fixgen.py

// NetworkZeroConfiguration type
type NetworkZeroConfiguration struct {

	// Unique identifier of network interface.
	InterfaceToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema InterfaceToken,omitempty"`

	// Indicates whether the zero-configuration is enabled or not.
	Enabled bool `xml:"http://www.onvif.org/ver10/search/wsdl Enabled,omitempty"`

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

	Alias Name `xml:"http://www.onvif.org/ver10/schema Alias,omitempty"`

	Priority NetworkInterfaceConfigPriority `xml:"http://www.onvif.org/ver10/schema Priority,omitempty"`

	Security Dot11SecurityConfiguration `xml:"http://www.onvif.org/ver10/schema Security,omitempty"`
}

// Dot11SecurityConfiguration type
type Dot11SecurityConfiguration struct {
	Mode Dot11SecurityMode `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	Algorithm Dot11Cipher `xml:"http://www.onvif.org/ver10/schema Algorithm,omitempty"`

	PSK Dot11PSKSet `xml:"http://www.onvif.org/ver10/schema PSK,omitempty"`

	Dot1X ReferenceToken `xml:"http://www.onvif.org/ver10/schema Dot1X,omitempty"`

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

// Removed Dot11Capabilities by fixgen.py

// Removed Dot11Status by fixgen.py

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
	RuleSupport bool `xml:"http://www.onvif.org/ver10/search/wsdl RuleSupport,omitempty"`

	// Indicates whether or not modules are supported.
	AnalyticsModuleSupport bool `xml:"http://www.onvif.org/ver10/search/wsdl AnalyticsModuleSupport,omitempty"`
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
	WSSubscriptionPolicySupport bool `xml:"http://www.onvif.org/ver10/search/wsdl WSSubscriptionPolicySupport,omitempty"`

	// Indicates whether or not WS Pull Point is supported.
	WSPullPointSupport bool `xml:"http://www.onvif.org/ver10/search/wsdl WSPullPointSupport,omitempty"`

	// Indicates whether or not WS Pausable Subscription Manager Interface is supported.
	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"http://www.onvif.org/ver10/search/wsdl WSPausableSubscriptionManagerInterfaceSupport,omitempty"`
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
	Auxiliary bool `xml:"http://www.onvif.org/ver10/search/wsdl Auxiliary,omitempty"`

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
	ProfileCapabilities ProfileCapabilities `xml:"http://www.onvif.org/ver10/schema ProfileCapabilities,omitempty"`
}

// RealTimeStreamingCapabilities type
type RealTimeStreamingCapabilities struct {

	// Indicates whether or not RTP multicast is supported.
	RTPMulticast bool `xml:"http://www.onvif.org/ver10/search/wsdl RTPMulticast,omitempty"`

	// Indicates whether or not RTP over TCP is supported.
	RTP_TCP bool `xml:"http://www.onvif.org/ver10/search/wsdl RTP_TCP,omitempty"`

	// Indicates whether or not RTP/RTSP/TCP is supported.
	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/search/wsdl RTP_RTSP_TCP,omitempty"`

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

// NetworkCapabilities type
type NetworkCapabilities struct {

	// Indicates whether or not IP filtering is supported.
	IPFilter bool `xml:"http://www.onvif.org/ver10/search/wsdl IPFilter,omitempty"`

	// Indicates whether or not zeroconf is supported.
	ZeroConfiguration bool `xml:"http://www.onvif.org/ver10/search/wsdl ZeroConfiguration,omitempty"`

	// Indicates whether or not IPv6 is supported.
	IPVersion6 bool `xml:"http://www.onvif.org/ver10/search/wsdl IPVersion6,omitempty"`

	// Indicates whether or not  is supported.
	DynDNS bool `xml:"http://www.onvif.org/ver10/search/wsdl DynDNS,omitempty"`

	Extension NetworkCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// NetworkCapabilitiesExtension type
type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool `xml:"http://www.onvif.org/ver10/search/wsdl Dot11Configuration,omitempty"`

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
	OnboardKeyGeneration bool `xml:"http://www.onvif.org/ver10/search/wsdl OnboardKeyGeneration,omitempty"`

	// Indicates whether or not access policy configuration is supported.
	AccessPolicyConfig bool `xml:"http://www.onvif.org/ver10/search/wsdl AccessPolicyConfig,omitempty"`

	// Indicates whether or not WS-Security X.509 token is supported.
	X509Token bool `xml:"X.509Token,omitempty"`

	// Indicates whether or not WS-Security SAML token is supported.
	SAMLToken bool `xml:"http://www.onvif.org/ver10/search/wsdl SAMLToken,omitempty"`

	// Indicates whether or not WS-Security Kerberos token is supported.
	KerberosToken bool `xml:"http://www.onvif.org/ver10/search/wsdl KerberosToken,omitempty"`

	// Indicates whether or not WS-Security REL token is supported.
	RELToken bool `xml:"http://www.onvif.org/ver10/search/wsdl RELToken,omitempty"`

	Extension SecurityCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SecurityCapabilitiesExtension type
type SecurityCapabilitiesExtension struct {
	TLS10 bool `xml:"TLS1.0,omitempty"`

	Extension SecurityCapabilitiesExtension2 `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SecurityCapabilitiesExtension2 type
type SecurityCapabilitiesExtension2 struct {
	Dot1X bool `xml:"http://www.onvif.org/ver10/search/wsdl Dot1X,omitempty"`

	// EAP Methods supported by the device. The int values refer to the .
	SupportedEAPMethod []int32 `xml:"http://www.onvif.org/ver10/schema SupportedEAPMethod,omitempty"`

	RemoteUserHandling bool `xml:"http://www.onvif.org/ver10/search/wsdl RemoteUserHandling,omitempty"`
}

// SystemCapabilities type
type SystemCapabilities struct {

	// Indicates whether or not WS Discovery resolve requests are supported.
	DiscoveryResolve bool `xml:"http://www.onvif.org/ver10/search/wsdl DiscoveryResolve,omitempty"`

	// Indicates whether or not WS-Discovery Bye is supported.
	DiscoveryBye bool `xml:"http://www.onvif.org/ver10/search/wsdl DiscoveryBye,omitempty"`

	// Indicates whether or not remote discovery is supported.
	RemoteDiscovery bool `xml:"http://www.onvif.org/ver10/search/wsdl RemoteDiscovery,omitempty"`

	// Indicates whether or not system backup is supported.
	SystemBackup bool `xml:"http://www.onvif.org/ver10/search/wsdl SystemBackup,omitempty"`

	// Indicates whether or not system logging is supported.
	SystemLogging bool `xml:"http://www.onvif.org/ver10/search/wsdl SystemLogging,omitempty"`

	// Indicates whether or not firmware upgrade is supported.
	FirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/search/wsdl FirmwareUpgrade,omitempty"`

	// Indicates supported ONVIF version(s).
	SupportedVersions []OnvifVersion `xml:"http://www.onvif.org/ver10/schema SupportedVersions,omitempty"`

	Extension SystemCapabilitiesExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SystemCapabilitiesExtension type
type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade bool `xml:"http://www.onvif.org/ver10/search/wsdl HttpFirmwareUpgrade,omitempty"`

	HttpSystemBackup bool `xml:"http://www.onvif.org/ver10/search/wsdl HttpSystemBackup,omitempty"`

	HttpSystemLogging bool `xml:"http://www.onvif.org/ver10/search/wsdl HttpSystemLogging,omitempty"`

	HttpSupportInformation bool `xml:"http://www.onvif.org/ver10/search/wsdl HttpSupportInformation,omitempty"`

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
	FixedLayout bool `xml:"http://www.onvif.org/ver10/search/wsdl FixedLayout,omitempty"`
}

// RecordingCapabilities type
type RecordingCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	ReceiverSource bool `xml:"http://www.onvif.org/ver10/search/wsdl ReceiverSource,omitempty"`

	MediaProfileSource bool `xml:"http://www.onvif.org/ver10/search/wsdl MediaProfileSource,omitempty"`

	DynamicRecordings bool `xml:"http://www.onvif.org/ver10/search/wsdl DynamicRecordings,omitempty"`

	DynamicTracks bool `xml:"http://www.onvif.org/ver10/search/wsdl DynamicTracks,omitempty"`

	MaxStringLength int32 `xml:"http://www.onvif.org/ver10/schema MaxStringLength,omitempty"`
}

// SearchCapabilities type
type SearchCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	MetadataSearch bool `xml:"http://www.onvif.org/ver10/search/wsdl MetadataSearch,omitempty"`
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
	RTP_Multicast bool `xml:"http://www.onvif.org/ver10/search/wsdl RTP_Multicast,omitempty"`

	// Indicates whether the device can receive RTP/TCP streams
	RTP_TCP bool `xml:"http://www.onvif.org/ver10/search/wsdl RTP_TCP,omitempty"`

	// Indicates whether the device can receive RTP/RTSP/TCP streams.
	RTP_RTSP_TCP bool `xml:"http://www.onvif.org/ver10/search/wsdl RTP_RTSP_TCP,omitempty"`

	// The maximum number of receivers supported by the device.
	SupportedReceivers int32 `xml:"http://www.onvif.org/ver10/schema SupportedReceivers,omitempty"`

	// The maximum allowed length for RTSP URIs.
	MaximumRTSPURILength int32 `xml:"http://www.onvif.org/ver10/schema MaximumRTSPURILength,omitempty"`
}

// AnalyticsDeviceCapabilities type
type AnalyticsDeviceCapabilities struct {
	XAddr AnyURI `xml:"http://www.onvif.org/ver10/schema XAddr,omitempty"`

	// Obsolete property.
	RuleSupport bool `xml:"http://www.onvif.org/ver10/search/wsdl RuleSupport,omitempty"`

	Extension AnalyticsDeviceExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// AnalyticsDeviceExtension type
type AnalyticsDeviceExtension struct {
}

// SystemLog type
type SystemLog struct {

	// The log information as attachment data.
	Binary AttachmentData `xml:"http://www.onvif.org/ver10/schema Binary,omitempty"`

	// The log information as character data.
	String string `xml:"http://www.onvif.org/ver10/schema String,omitempty"`
}

// Removed SupportInformation by fixgen.py

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

// Removed BackupFile by fixgen.py

// Removed SystemLogUriList by fixgen.py

// Removed SystemLogUri by fixgen.py

// SystemDateTime type
type SystemDateTime struct {

	// Indicates if the time is set manully or through NTP.
	DateTimeType SetDateTimeType `xml:"http://www.onvif.org/ver10/schema DateTimeType,omitempty"`

	// Informative indicator whether daylight savings is currently on/off.
	DaylightSavings bool `xml:"http://www.onvif.org/ver10/search/wsdl DaylightSavings,omitempty"`

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

// Removed RemoteUser by fixgen.py

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

// Removed CertificateStatus by fixgen.py

// Removed CertificateWithPrivateKey by fixgen.py

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

	Critical bool `xml:"http://www.onvif.org/ver10/search/wsdl Critical,attr,omitempty"`
}

// CertificateInformationExtension type
type CertificateInformationExtension struct {
}

// Dot1XConfiguration type
type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema Dot1XConfigurationToken,omitempty"`

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

	Properties RelayOutputSettings `xml:"http://www.onvif.org/ver10/schema Properties,omitempty"`
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
	HomeSupported bool `xml:"http://www.onvif.org/ver10/search/wsdl HomeSupported,omitempty"`

	//
	// A list of supported Auxiliary commands. If the list is not empty, the Auxiliary Operations MUST be available for this PTZ Node.
	//
	AuxiliaryCommands []AuxiliaryData `xml:"http://www.onvif.org/ver10/schema AuxiliaryCommands,omitempty"`

	Extension PTZNodeExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`

	//
	// Indication whether the HomePosition of a Node is fixed or it can be changed via the SetHomePosition command.
	//

	FixedHomePosition bool `xml:"http://www.onvif.org/ver10/search/wsdl FixedHomePosition,attr,omitempty"`

	//
	// Indication whether the Node supports the geo-referenced move command.
	//

	GeoMove bool `xml:"http://www.onvif.org/ver10/search/wsdl GeoMove,attr,omitempty"`
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
	NodeToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema NodeToken,omitempty"`

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
	PresetToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema PresetToken,omitempty"`

	// Option to specify the preset position with the home position of this PTZ Node. "False" to this parameter shall be treated as an invalid argument.
	Home bool `xml:"http://www.onvif.org/ver10/search/wsdl Home,omitempty"`

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

	RandomPresetOrder bool `xml:"http://www.onvif.org/ver10/search/wsdl RandomPresetOrder,attr,omitempty"`
}

// PTZPresetTourStartingConditionExtension type
type PTZPresetTourStartingConditionExtension struct {
}

// Removed PTZPresetTourOptions by fixgen.py

// Removed PTZPresetTourSpotOptions by fixgen.py

// PTZPresetTourPresetDetailOptions type
type PTZPresetTourPresetDetailOptions struct {

	// A list of available Preset Tokens for tour spots.
	PresetToken []ReferenceToken `xml:"http://www.onvif.org/ver10/schema PresetToken,omitempty"`

	// An option to indicate Home postion for tour spots.
	Home bool `xml:"http://www.onvif.org/ver10/search/wsdl Home,omitempty"`

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
	BoundaryOffset bool `xml:"http://www.onvif.org/ver10/search/wsdl BoundaryOffset,omitempty"`

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
	Level bool `xml:"http://www.onvif.org/ver10/search/wsdl Level,omitempty"`
}

// DefoggingOptions type
type DefoggingOptions struct {

	// Supported options for Defogging mode. Its options shall be chosen from tt:DefoggingMode Type.
	Mode []string `xml:"http://www.onvif.org/ver10/schema Mode,omitempty"`

	// Indicates whether or not support Level parameter for Defogging.
	Level bool `xml:"http://www.onvif.org/ver10/search/wsdl Level,omitempty"`
}

// NoiseReductionOptions type
type NoiseReductionOptions struct {

	// Indicates whether or not support Level parameter for NoiseReduction.
	Level bool `xml:"http://www.onvif.org/ver10/search/wsdl Level,omitempty"`
}

// MessageExtension type
type MessageExtension struct {
}

// ItemList type
type ItemList struct {
	SimpleItem []struct {

		// Item name.

		Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`

		// Item value. The type is defined in the corresponding description.

		Value AnySimpleType `xml:"Value,attr,omitempty"`
	} `xml:"SimpleItem,omitempty"`

	ElementItem []struct {

		// Item name.

		Name string `xml:"http://www.onvif.org/ver10/schema Name,attr,omitempty"`
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

	IsProperty bool `xml:"http://www.onvif.org/ver10/search/wsdl IsProperty,attr,omitempty"`
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
	Expression string `xml:"http://www.onvif.org/ver10/schema Expression,omitempty"`

	Type string `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
}

// Removed MotionExpressionConfiguration by fixgen.py

// Removed CellLayout by fixgen.py

// Removed PaneConfiguration by fixgen.py

// PaneLayout type
type PaneLayout struct {

	// Reference to the configuration of the streaming and coding parameters.
	Pane ReferenceToken `xml:"http://www.onvif.org/ver10/schema Pane,omitempty"`

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
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schema Token,omitempty"`

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
	Token ReferenceToken `xml:"http://www.onvif.org/ver10/schema Token,omitempty"`

	Type AnyURI `xml:"http://www.onvif.org/ver10/schema Type,attr,omitempty"`
}

// DateTimeRange type
type DateTimeRange struct {
	From string `xml:"http://www.onvif.org/ver10/schema From,omitempty"`

	Until string `xml:"http://www.onvif.org/ver10/schema Until,omitempty"`
}

// RecordingSummary type
type RecordingSummary struct {

	// The earliest point in time where there is recorded data on the device.
	DataFrom string `xml:"http://www.onvif.org/ver10/schema DataFrom,omitempty"`

	// The most recent point in time where there is recorded data on the device.
	DataUntil string `xml:"http://www.onvif.org/ver10/schema DataUntil,omitempty"`

	// The device contains this many recordings.
	NumberRecordings int32 `xml:"http://www.onvif.org/ver10/schema NumberRecordings,omitempty"`
}

// SearchScope type
type SearchScope struct {

	// A list of sources that are included in the scope. If this list is included, only data from one of these sources shall be searched.
	IncludedSources []SourceReference `xml:"http://www.onvif.org/ver10/schema IncludedSources,omitempty"`

	// A list of recordings that are included in the scope. If this list is included, only data from one of these recordings shall be searched.
	IncludedRecordings []RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl IncludedRecordings,omitempty"`

	// An xpath expression used to specify what recordings to search. Only those recordings with an RecordingInformation structure that matches the filter shall be searched.
	RecordingInformationFilter XPathExpression `xml:"http://www.onvif.org/ver10/schema RecordingInformationFilter,omitempty"`

	// Extension point
	Extension SearchScopeExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// SearchScopeExtension type
type SearchScopeExtension struct {
}

// EventFilter type
type EventFilter struct {
	*FilterType
}

// PTZPositionFilter type
type PTZPositionFilter struct {

	// The lower boundary of the PTZ volume to look for.
	MinPosition PTZVector `xml:"http://www.onvif.org/ver10/schema MinPosition,omitempty"`

	// The upper boundary of the PTZ volume to look for.
	MaxPosition PTZVector `xml:"http://www.onvif.org/ver10/schema MaxPosition,omitempty"`

	// If true, search for when entering the specified PTZ volume.
	EnterOrExit bool `xml:"http://www.onvif.org/ver10/search/wsdl EnterOrExit,omitempty"`
}

// MetadataFilter type
type MetadataFilter struct {
	MetadataStreamFilter XPathExpression `xml:"http://www.onvif.org/ver10/schema MetadataStreamFilter,omitempty"`
}

// FindRecordingResultList type
type FindRecordingResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState SearchState `xml:"http://www.onvif.org/ver10/search/wsdl SearchState,omitempty"`

	// A RecordingInformation structure for each found recording matching the search.
	RecordingInformation []RecordingInformation `xml:"http://www.onvif.org/ver10/search/wsdl RecordingInformation,omitempty"`
}

// FindEventResultList type
type FindEventResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState SearchState `xml:"http://www.onvif.org/ver10/search/wsdl SearchState,omitempty"`

	// A FindEventResult structure for each found event matching the search.
	Result []FindEventResult `xml:"http://www.onvif.org/ver10/search/wsdl Result,omitempty"`
}

// FindEventResult type
type FindEventResult struct {

	// The recording where this event was found. Empty string if no recording is associated with this event.
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

	// A reference to the track where this event was found. Empty string if no track is associated with this event.
	TrackToken TrackReference `xml:"http://www.onvif.org/ver10/schema TrackToken,omitempty"`

	// The time when the event occured.
	Time string `xml:"http://www.onvif.org/ver10/schema Time,omitempty"`

	// The description of the event.
	Event NotificationMessageHolderType `xml:"Event,omitempty"`

	// If true, indicates that the event is a virtual event generated for this particular search session to give the state of a property at the start time of the search.
	StartStateEvent bool `xml:"http://www.onvif.org/ver10/search/wsdl StartStateEvent,omitempty"`
}

// FindPTZPositionResultList type
type FindPTZPositionResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState SearchState `xml:"http://www.onvif.org/ver10/search/wsdl SearchState,omitempty"`

	// A FindPTZPositionResult structure for each found PTZ position matching the search.
	Result []FindPTZPositionResult `xml:"http://www.onvif.org/ver10/search/wsdl Result,omitempty"`
}

// FindPTZPositionResult type
type FindPTZPositionResult struct {

	// A reference to the recording containing the PTZ position.
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

	// A reference to the metadata track containing the PTZ position.
	TrackToken TrackReference `xml:"http://www.onvif.org/ver10/schema TrackToken,omitempty"`

	// The time when the PTZ position was valid.
	Time string `xml:"http://www.onvif.org/ver10/schema Time,omitempty"`

	// The PTZ position.
	Position PTZVector `xml:"http://www.onvif.org/ver10/schema Position,omitempty"`
}

// FindMetadataResultList type
type FindMetadataResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState SearchState `xml:"http://www.onvif.org/ver10/search/wsdl SearchState,omitempty"`

	// A FindMetadataResult structure for each found set of Metadata matching the search.
	Result []FindMetadataResult `xml:"http://www.onvif.org/ver10/search/wsdl Result,omitempty"`
}

// FindMetadataResult type
type FindMetadataResult struct {

	// A reference to the recording containing the metadata.
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

	// A reference to the metadata track containing the matching metadata.
	TrackToken TrackReference `xml:"http://www.onvif.org/ver10/schema TrackToken,omitempty"`

	// The point in time when the matching metadata occurs in the metadata track.
	Time string `xml:"http://www.onvif.org/ver10/schema Time,omitempty"`
}

// RecordingInformation type
type RecordingInformation struct {
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

	//
	// Information about the source of the recording. This gives a description of where the data in the recording comes from. Since a single
	// recording is intended to record related material, there is just one source. It is indicates the physical location or the
	// major data source for the recording. Currently the recordingconfiguration cannot describe each individual data source.
	//
	Source RecordingSourceInformation `xml:"http://www.onvif.org/ver10/schema Source,omitempty"`

	EarliestRecording string `xml:"http://www.onvif.org/ver10/schema EarliestRecording,omitempty"`

	LatestRecording string `xml:"http://www.onvif.org/ver10/schema LatestRecording,omitempty"`

	Content Description `xml:"http://www.onvif.org/ver10/schema Content,omitempty"`

	// Basic information about the track. Note that a track may represent a single contiguous time span or consist of multiple slices.
	Track []TrackInformation `xml:"http://www.onvif.org/ver10/schema Track,omitempty"`

	RecordingStatus RecordingStatus `xml:"http://www.onvif.org/ver10/schema RecordingStatus,omitempty"`
}

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

// MediaAttributes type
type MediaAttributes struct {

	// A reference to the recording that has these attributes.
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

	// A set of attributes for each track.
	TrackAttributes []TrackAttributes `xml:"http://www.onvif.org/ver10/schema TrackAttributes,omitempty"`

	// The attributes are valid from this point in time in the recording.
	From string `xml:"http://www.onvif.org/ver10/schema From,omitempty"`

	// The attributes are valid until this point in time in the recording. Can be equal to 'From' to indicate that the attributes are only known to be valid for this particular point in time.
	Until string `xml:"http://www.onvif.org/ver10/schema Until,omitempty"`
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
	CanContainPTZ bool `xml:"http://www.onvif.org/ver10/search/wsdl CanContainPTZ,omitempty"`

	// Indicates that there can be analytics data in the metadata track in the specified time interval.
	CanContainAnalytics bool `xml:"http://www.onvif.org/ver10/search/wsdl CanContainAnalytics,omitempty"`

	// Indicates that there can be notifications in the metadata track in the specified time interval.
	CanContainNotifications bool `xml:"http://www.onvif.org/ver10/search/wsdl CanContainNotifications,omitempty"`

	// List of all PTZ spaces active for recording. Note that events are only recorded on position changes and the actual point of recording may not necessarily contain an event of the specified type.

	PtzSpaces StringAttrList `xml:"http://www.onvif.org/ver10/schema PtzSpaces,attr,omitempty"`
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
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

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
	AutoCreateReceiver bool `xml:"http://www.onvif.org/ver10/search/wsdl AutoCreateReceiver,omitempty"`

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
	RecordingToken RecordingReference `xml:"http://www.onvif.org/ver10/search/wsdl RecordingToken,omitempty"`

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
	Name string `xml:"http://www.onvif.org/ver10/schema Name,omitempty"`

	Token []ReferenceToken `xml:"http://www.onvif.org/ver10/schema Token,omitempty"`

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

	IsPersistentText bool `xml:"http://www.onvif.org/ver10/search/wsdl IsPersistentText,attr,omitempty"`
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
	StorageToken ReferenceToken `xml:"http://www.onvif.org/ver10/schema StorageToken,omitempty"`

	// gives the relative directory path on the storage
	RelativePath string `xml:"http://www.onvif.org/ver10/schema RelativePath,omitempty"`

	Extension StorageReferencePathExtension `xml:"http://www.onvif.org/ver10/schema Extension,omitempty"`
}

// StorageReferencePathExtension type
type StorageReferencePathExtension struct {
}

// SearchPort type
type SearchPort interface {

	/* Returns the capabilities of the search service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/* GetRecordingSummary is used to get a summary description of all recorded data. This
	operation is mandatory to support for a device implementing the recording search service. */
	GetRecordingSummary(request *GetRecordingSummary) (*GetRecordingSummaryResponse, error)

	GetRecordingSummaryContext(ctx context.Context, request *GetRecordingSummary) (*GetRecordingSummaryResponse, error)

	/* Returns information about a single Recording specified by a RecordingToken. This operation
	is mandatory to support for a device implementing the recording search service. */
	GetRecordingInformation(request *GetRecordingInformation) (*GetRecordingInformationResponse, error)

	GetRecordingInformationContext(ctx context.Context, request *GetRecordingInformation) (*GetRecordingInformationResponse, error)

	/* Returns a set of media attributes for all tracks of the specified recordings at a specified point
	in time. Clients using this operation shall be able to use it as a non blocking operation. A
	device shall set the starttime and endtime of the MediaAttributes structure to equal values if
	calculating this range would causes this operation to block. See MediaAttributes structure for
	more information. This operation is mandatory to support for a device implementing the
	recording search service. */
	GetMediaAttributes(request *GetMediaAttributes) (*GetMediaAttributesResponse, error)

	GetMediaAttributesContext(ctx context.Context, request *GetMediaAttributes) (*GetMediaAttributesResponse, error)

	/* FindRecordings starts a search session, looking for recordings that matches the scope (See
	20.2.4) defined in the request. Results from the search are acquired using the
	GetRecordingSearchResults request, specifying the search token returned from this request.
	The device shall continue searching until one of the following occurs:
	The order of the results is undefined, to allow the device to return results in any order they
	are found. This operation is mandatory to support for a device implementing the recording
	search service. */
	FindRecordings(request *FindRecordings) (*FindRecordingsResponse, error)

	FindRecordingsContext(ctx context.Context, request *FindRecordings) (*FindRecordingsResponse, error)

	/* GetRecordingSearchResults acquires the results from a recording search session previously
	initiated by a FindRecordings operation. The response shall not include results already
	returned in previous requests for the same session. If MaxResults is specified, the response
	shall not contain more than MaxResults results. The number of results relates to the number of recordings.
	For viewing individual recorded data for a signal track use the FindEvents method.
	GetRecordingSearchResults shall block until:
	This operation is mandatory to support for a device implementing the recording search service. */
	GetRecordingSearchResults(request *GetRecordingSearchResults) (*GetRecordingSearchResultsResponse, error)

	GetRecordingSearchResultsContext(ctx context.Context, request *GetRecordingSearchResults) (*GetRecordingSearchResultsResponse, error)

	/* FindEvents starts a search session, looking for recording events (in the scope that
	matches the search filter defined in the request. Results from the search are
	acquired using the GetEventSearchResults request, specifying the search token returned from
	this request.
	The device shall continue searching until one of the following occurs:
	Results shall be ordered by time, ascending in case of forward search, or descending in case
	of backward search. This operation is mandatory to support for a device implementing the
	recording search service. */
	FindEvents(request *FindEvents) (*FindEventsResponse, error)

	FindEventsContext(ctx context.Context, request *FindEvents) (*FindEventsResponse, error)

	/* GetEventSearchResults acquires the results from a recording event search session previously
	initiated by a FindEvents operation. The response shall not include results already returned in
	previous requests for the same session. If MaxResults is specified, the response shall not
	contain more than MaxResults results.
	GetEventSearchResults shall block until:
	This operation is mandatory to support for a device implementing the recording search service. */
	GetEventSearchResults(request *GetEventSearchResults) (*GetEventSearchResultsResponse, error)

	GetEventSearchResultsContext(ctx context.Context, request *GetEventSearchResults) (*GetEventSearchResultsResponse, error)

	/* FindPTZPosition starts a search session, looking for ptz positions in the scope (See 20.2.4)
	that matches the search filter defined in the request. Results from the search are acquired
	using the GetPTZPositionSearchResults request, specifying the search token returned from
	this request.
	The device shall continue searching until one of the following occurs:
	This operation is mandatory to support whenever CanContainPTZ is true for any metadata
	track in any recording on the device. */
	FindPTZPosition(request *FindPTZPosition) (*FindPTZPositionResponse, error)

	FindPTZPositionContext(ctx context.Context, request *FindPTZPosition) (*FindPTZPositionResponse, error)

	/* GetPTZPositionSearchResults acquires the results from a ptz position search session
	previously initiated by a FindPTZPosition operation. The response shall not include results
	already returned in previous requests for the same session. If MaxResults is specified, the
	response shall not contain more than MaxResults results.
	GetPTZPositionSearchResults shall block until:
	This operation is mandatory to support whenever CanContainPTZ is true for any metadata
	track in any recording on the device. */
	GetPTZPositionSearchResults(request *GetPTZPositionSearchResults) (*GetPTZPositionSearchResultsResponse, error)

	GetPTZPositionSearchResultsContext(ctx context.Context, request *GetPTZPositionSearchResults) (*GetPTZPositionSearchResultsResponse, error)

	/* GetSearchState returns the current state of the specified search session. This command is deprecated . */
	GetSearchState(request *GetSearchState) (*GetSearchStateResponse, error)

	GetSearchStateContext(ctx context.Context, request *GetSearchState) (*GetSearchStateResponse, error)

	/* EndSearch stops and ongoing search session, causing any blocking result request to return
	and the SearchToken to become invalid. If the search was interrupted before completion, the
	point in time that the search had reached shall be returned. If the search had not yet begun,
	the StartPoint shall be returned. If the search was completed the original EndPoint supplied
	by the Find operation shall be returned. When issuing EndSearch on a FindRecordings request the
	EndPoint is undefined and shall not be used since the FindRecordings request doesn't have
	StartPoint/EndPoint. This operation is mandatory to support for a device implementing the
	recording search service.
	*/
	EndSearch(request *EndSearch) (*EndSearchResponse, error)

	EndSearchContext(ctx context.Context, request *EndSearch) (*EndSearchResponse, error)

	/* FindMetadata starts a search session, looking for metadata in the scope (See 20.2.4) that
	matches the search filter defined in the request. Results from the search are acquired using
	the GetMetadataSearchResults request, specifying the search token returned from this
	request.
	The device shall continue searching until one of the following occurs:
	This operation is mandatory to support if the MetaDataSearch capability is set to true in the
	SearchCapabilities structure return by the GetCapabilities command in the Device service. */
	FindMetadata(request *FindMetadata) (*FindMetadataResponse, error)

	FindMetadataContext(ctx context.Context, request *FindMetadata) (*FindMetadataResponse, error)

	/* GetMetadataSearchResults acquires the results from a recording search session previously
	initiated by a FindMetadata operation. The response shall not include results already returned
	in previous requests for the same session. If MaxResults is specified, the response shall not
	contain more than MaxResults results.
	GetMetadataSearchResults shall block until:
	This operation is mandatory to support if the MetaDataSearch capability is set to true in the
	SearchCapabilities structure return by the GetCapabilities command in the Device service. */
	GetMetadataSearchResults(request *GetMetadataSearchResults) (*GetMetadataSearchResultsResponse, error)

	GetMetadataSearchResultsContext(ctx context.Context, request *GetMetadataSearchResults) (*GetMetadataSearchResultsResponse, error)
}

// searchPort type
type searchPort struct {
	client *soap.Client
	xaddr  string
}

func NewSearchPort(client *soap.Client, xaddr string) SearchPort {
	return &searchPort{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *searchPort) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetRecordingSummaryContext(ctx context.Context, request *GetRecordingSummary) (*GetRecordingSummaryResponse, error) {
	response := new(GetRecordingSummaryResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetRecordingSummary", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetRecordingSummary(request *GetRecordingSummary) (*GetRecordingSummaryResponse, error) {
	return service.GetRecordingSummaryContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetRecordingInformationContext(ctx context.Context, request *GetRecordingInformation) (*GetRecordingInformationResponse, error) {
	response := new(GetRecordingInformationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetRecordingInformation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetRecordingInformation(request *GetRecordingInformation) (*GetRecordingInformationResponse, error) {
	return service.GetRecordingInformationContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetMediaAttributesContext(ctx context.Context, request *GetMediaAttributes) (*GetMediaAttributesResponse, error) {
	response := new(GetMediaAttributesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetMediaAttributes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetMediaAttributes(request *GetMediaAttributes) (*GetMediaAttributesResponse, error) {
	return service.GetMediaAttributesContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) FindRecordingsContext(ctx context.Context, request *FindRecordings) (*FindRecordingsResponse, error) {
	response := new(FindRecordingsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/FindRecordings", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) FindRecordings(request *FindRecordings) (*FindRecordingsResponse, error) {
	return service.FindRecordingsContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetRecordingSearchResultsContext(ctx context.Context, request *GetRecordingSearchResults) (*GetRecordingSearchResultsResponse, error) {
	response := new(GetRecordingSearchResultsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetRecordingSearchResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetRecordingSearchResults(request *GetRecordingSearchResults) (*GetRecordingSearchResultsResponse, error) {
	return service.GetRecordingSearchResultsContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) FindEventsContext(ctx context.Context, request *FindEvents) (*FindEventsResponse, error) {
	response := new(FindEventsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/FindEvents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) FindEvents(request *FindEvents) (*FindEventsResponse, error) {
	return service.FindEventsContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetEventSearchResultsContext(ctx context.Context, request *GetEventSearchResults) (*GetEventSearchResultsResponse, error) {
	response := new(GetEventSearchResultsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetEventSearchResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetEventSearchResults(request *GetEventSearchResults) (*GetEventSearchResultsResponse, error) {
	return service.GetEventSearchResultsContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) FindPTZPositionContext(ctx context.Context, request *FindPTZPosition) (*FindPTZPositionResponse, error) {
	response := new(FindPTZPositionResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/FindPTZPosition", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) FindPTZPosition(request *FindPTZPosition) (*FindPTZPositionResponse, error) {
	return service.FindPTZPositionContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetPTZPositionSearchResultsContext(ctx context.Context, request *GetPTZPositionSearchResults) (*GetPTZPositionSearchResultsResponse, error) {
	response := new(GetPTZPositionSearchResultsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetPTZPositionSearchResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetPTZPositionSearchResults(request *GetPTZPositionSearchResults) (*GetPTZPositionSearchResultsResponse, error) {
	return service.GetPTZPositionSearchResultsContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetSearchStateContext(ctx context.Context, request *GetSearchState) (*GetSearchStateResponse, error) {
	response := new(GetSearchStateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetSearchState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetSearchState(request *GetSearchState) (*GetSearchStateResponse, error) {
	return service.GetSearchStateContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) EndSearchContext(ctx context.Context, request *EndSearch) (*EndSearchResponse, error) {
	response := new(EndSearchResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/EndSearch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) EndSearch(request *EndSearch) (*EndSearchResponse, error) {
	return service.EndSearchContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) FindMetadataContext(ctx context.Context, request *FindMetadata) (*FindMetadataResponse, error) {
	response := new(FindMetadataResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/FindMetadata", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) FindMetadata(request *FindMetadata) (*FindMetadataResponse, error) {
	return service.FindMetadataContext(
		context.Background(),
		request,
	)
}

func (service *searchPort) GetMetadataSearchResultsContext(ctx context.Context, request *GetMetadataSearchResults) (*GetMetadataSearchResultsResponse, error) {
	response := new(GetMetadataSearchResultsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/search/wsdl/GetMetadataSearchResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *searchPort) GetMetadataSearchResults(request *GetMetadataSearchResults) (*GetMetadataSearchResultsResponse, error) {
	return service.GetMetadataSearchResultsContext(
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
