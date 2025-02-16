// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtraildata

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
)

const opPutAuditEvents = "PutAuditEvents"

// PutAuditEventsRequest generates a "aws/request.Request" representing the
// client's request for the PutAuditEvents operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutAuditEvents for more information on using the PutAuditEvents
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the PutAuditEventsRequest method.
//	req, resp := client.PutAuditEventsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-data-2021-08-11/PutAuditEvents
func (c *CloudTrailData) PutAuditEventsRequest(input *PutAuditEventsInput) (req *request.Request, output *PutAuditEventsOutput) {
	op := &request.Operation{
		Name:       opPutAuditEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/PutAuditEvents",
	}

	if input == nil {
		input = &PutAuditEventsInput{}
	}

	output = &PutAuditEventsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// PutAuditEvents API operation for AWS CloudTrail Data Service.
//
// Ingests your application events into CloudTrail Lake. A required parameter,
// auditEvents, accepts the JSON records (also called payload) of events that
// you want CloudTrail to ingest. You can add up to 100 of these events (or
// up to 1 MB) per PutAuditEvents request.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS CloudTrail Data Service's
// API operation PutAuditEvents for usage and error information.
//
// Returned Error Types:
//
//   - ChannelInsufficientPermission
//     The caller's account ID must be the same as the channel owner's account ID.
//
//   - InvalidAuditEventId
//     An audit event ID must not be empty, and can only contain alphanumeric characters,
//     hyphens, and underscores. It can be a maximum of 128 characters long.
//
//   - ChannelNotFound
//     The channel could not be found.
//
//   - InvalidChannelARN
//     The specified channel ARN is not a valid channel ARN.
//
//   - ChannelUnsupportedSchema
//     The schema type of the event is not supported.
//
//   - DuplicatedAuditEventId
//     Two or more entries in the request have the same event ID.
//
//   - UnsupportedOperationException
//     The operation requested is not supported in this region or account.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-data-2021-08-11/PutAuditEvents
func (c *CloudTrailData) PutAuditEvents(input *PutAuditEventsInput) (*PutAuditEventsOutput, error) {
	req, out := c.PutAuditEventsRequest(input)
	return out, req.Send()
}

// PutAuditEventsWithContext is the same as PutAuditEvents with the addition of
// the ability to pass a context and additional request options.
//
// See PutAuditEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CloudTrailData) PutAuditEventsWithContext(ctx aws.Context, input *PutAuditEventsInput, opts ...request.Option) (*PutAuditEventsOutput, error) {
	req, out := c.PutAuditEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// An event from a source outside of Amazon Web Services that you want CloudTrail
// to log.
type AuditEvent struct {
	_ struct{} `type:"structure"`

	// The content of an audit event that comes from the event, such as userIdentity,
	// userAgent, and eventSource.
	//
	// EventData is a required field
	EventData *string `locationName:"eventData" type:"string" required:"true"`

	// A checksum is a base64-SHA256 algorithm that helps you verify that CloudTrail
	// receives the event that matches with the checksum. Calculate the checksum
	// by running a command like the following:
	//
	// printf %s $eventdata | openssl dgst -binary -sha256 | base64
	EventDataChecksum *string `locationName:"eventDataChecksum" type:"string"`

	// The original event ID from the source event.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuditEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuditEvent) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuditEvent) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AuditEvent"}
	if s.EventData == nil {
		invalidParams.Add(request.NewErrParamRequired("EventData"))
	}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEventData sets the EventData field's value.
func (s *AuditEvent) SetEventData(v string) *AuditEvent {
	s.EventData = &v
	return s
}

// SetEventDataChecksum sets the EventDataChecksum field's value.
func (s *AuditEvent) SetEventDataChecksum(v string) *AuditEvent {
	s.EventDataChecksum = &v
	return s
}

// SetId sets the Id field's value.
func (s *AuditEvent) SetId(v string) *AuditEvent {
	s.Id = &v
	return s
}

// A response that includes successful and failed event results.
type AuditEventResultEntry struct {
	_ struct{} `type:"structure"`

	// The CloudTrail matched checksum value.
	EventDataChecksum *string `locationName:"eventDataChecksum" type:"string"`

	// The event ID assigned by CloudTrail.
	//
	// EventID is a required field
	EventID *string `locationName:"eventID" min:"1" type:"string" required:"true"`

	// The original event ID from the source event.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuditEventResultEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuditEventResultEntry) GoString() string {
	return s.String()
}

// SetEventDataChecksum sets the EventDataChecksum field's value.
func (s *AuditEventResultEntry) SetEventDataChecksum(v string) *AuditEventResultEntry {
	s.EventDataChecksum = &v
	return s
}

// SetEventID sets the EventID field's value.
func (s *AuditEventResultEntry) SetEventID(v string) *AuditEventResultEntry {
	s.EventID = &v
	return s
}

// SetId sets the Id field's value.
func (s *AuditEventResultEntry) SetId(v string) *AuditEventResultEntry {
	s.Id = &v
	return s
}

// The caller's account ID must be the same as the channel owner's account ID.
type ChannelInsufficientPermission struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ChannelInsufficientPermission) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ChannelInsufficientPermission) GoString() string {
	return s.String()
}

func newErrorChannelInsufficientPermission(v protocol.ResponseMetadata) error {
	return &ChannelInsufficientPermission{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ChannelInsufficientPermission) Code() string {
	return "ChannelInsufficientPermission"
}

// Message returns the exception's message.
func (s *ChannelInsufficientPermission) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ChannelInsufficientPermission) OrigErr() error {
	return nil
}

func (s *ChannelInsufficientPermission) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ChannelInsufficientPermission) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ChannelInsufficientPermission) RequestID() string {
	return s.RespMetadata.RequestID
}

// The channel could not be found.
type ChannelNotFound struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ChannelNotFound) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ChannelNotFound) GoString() string {
	return s.String()
}

func newErrorChannelNotFound(v protocol.ResponseMetadata) error {
	return &ChannelNotFound{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ChannelNotFound) Code() string {
	return "ChannelNotFound"
}

// Message returns the exception's message.
func (s *ChannelNotFound) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ChannelNotFound) OrigErr() error {
	return nil
}

func (s *ChannelNotFound) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ChannelNotFound) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ChannelNotFound) RequestID() string {
	return s.RespMetadata.RequestID
}

// The schema type of the event is not supported.
type ChannelUnsupportedSchema struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ChannelUnsupportedSchema) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ChannelUnsupportedSchema) GoString() string {
	return s.String()
}

func newErrorChannelUnsupportedSchema(v protocol.ResponseMetadata) error {
	return &ChannelUnsupportedSchema{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ChannelUnsupportedSchema) Code() string {
	return "ChannelUnsupportedSchema"
}

// Message returns the exception's message.
func (s *ChannelUnsupportedSchema) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ChannelUnsupportedSchema) OrigErr() error {
	return nil
}

func (s *ChannelUnsupportedSchema) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ChannelUnsupportedSchema) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ChannelUnsupportedSchema) RequestID() string {
	return s.RespMetadata.RequestID
}

// Two or more entries in the request have the same event ID.
type DuplicatedAuditEventId struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DuplicatedAuditEventId) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DuplicatedAuditEventId) GoString() string {
	return s.String()
}

func newErrorDuplicatedAuditEventId(v protocol.ResponseMetadata) error {
	return &DuplicatedAuditEventId{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *DuplicatedAuditEventId) Code() string {
	return "DuplicatedAuditEventId"
}

// Message returns the exception's message.
func (s *DuplicatedAuditEventId) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *DuplicatedAuditEventId) OrigErr() error {
	return nil
}

func (s *DuplicatedAuditEventId) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *DuplicatedAuditEventId) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *DuplicatedAuditEventId) RequestID() string {
	return s.RespMetadata.RequestID
}

// An audit event ID must not be empty, and can only contain alphanumeric characters,
// hyphens, and underscores. It can be a maximum of 128 characters long.
type InvalidAuditEventId struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidAuditEventId) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidAuditEventId) GoString() string {
	return s.String()
}

func newErrorInvalidAuditEventId(v protocol.ResponseMetadata) error {
	return &InvalidAuditEventId{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidAuditEventId) Code() string {
	return "InvalidAuditEventId"
}

// Message returns the exception's message.
func (s *InvalidAuditEventId) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidAuditEventId) OrigErr() error {
	return nil
}

func (s *InvalidAuditEventId) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidAuditEventId) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidAuditEventId) RequestID() string {
	return s.RespMetadata.RequestID
}

// The specified channel ARN is not a valid channel ARN.
type InvalidChannelARN struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidChannelARN) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidChannelARN) GoString() string {
	return s.String()
}

func newErrorInvalidChannelARN(v protocol.ResponseMetadata) error {
	return &InvalidChannelARN{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidChannelARN) Code() string {
	return "InvalidChannelARN"
}

// Message returns the exception's message.
func (s *InvalidChannelARN) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidChannelARN) OrigErr() error {
	return nil
}

func (s *InvalidChannelARN) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidChannelARN) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidChannelARN) RequestID() string {
	return s.RespMetadata.RequestID
}

type PutAuditEventsInput struct {
	_ struct{} `type:"structure"`

	// The JSON payload of events that you want to ingest. You can also point to
	// the JSON event payload in a file.
	//
	// AuditEvents is a required field
	AuditEvents []*AuditEvent `locationName:"auditEvents" min:"1" type:"list" required:"true"`

	// The ARN or ID (the ARN suffix) of a channel.
	//
	// ChannelArn is a required field
	ChannelArn *string `location:"querystring" locationName:"channelArn" type:"string" required:"true"`

	// A unique identifier that is conditionally required when the channel's resource
	// policy includes an external ID. This value can be any string, such as a passphrase
	// or account number.
	ExternalId *string `location:"querystring" locationName:"externalId" min:"2" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutAuditEventsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutAuditEventsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutAuditEventsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutAuditEventsInput"}
	if s.AuditEvents == nil {
		invalidParams.Add(request.NewErrParamRequired("AuditEvents"))
	}
	if s.AuditEvents != nil && len(s.AuditEvents) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("AuditEvents", 1))
	}
	if s.ChannelArn == nil {
		invalidParams.Add(request.NewErrParamRequired("ChannelArn"))
	}
	if s.ExternalId != nil && len(*s.ExternalId) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("ExternalId", 2))
	}
	if s.AuditEvents != nil {
		for i, v := range s.AuditEvents {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AuditEvents", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAuditEvents sets the AuditEvents field's value.
func (s *PutAuditEventsInput) SetAuditEvents(v []*AuditEvent) *PutAuditEventsInput {
	s.AuditEvents = v
	return s
}

// SetChannelArn sets the ChannelArn field's value.
func (s *PutAuditEventsInput) SetChannelArn(v string) *PutAuditEventsInput {
	s.ChannelArn = &v
	return s
}

// SetExternalId sets the ExternalId field's value.
func (s *PutAuditEventsInput) SetExternalId(v string) *PutAuditEventsInput {
	s.ExternalId = &v
	return s
}

type PutAuditEventsOutput struct {
	_ struct{} `type:"structure"`

	// Lists events in the provided event payload that could not be ingested into
	// CloudTrail, and includes the error code and error message returned for events
	// that could not be ingested.
	//
	// Failed is a required field
	Failed []*ResultErrorEntry `locationName:"failed" type:"list" required:"true"`

	// Lists events in the provided event payload that were successfully ingested
	// into CloudTrail.
	//
	// Successful is a required field
	Successful []*AuditEventResultEntry `locationName:"successful" type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutAuditEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutAuditEventsOutput) GoString() string {
	return s.String()
}

// SetFailed sets the Failed field's value.
func (s *PutAuditEventsOutput) SetFailed(v []*ResultErrorEntry) *PutAuditEventsOutput {
	s.Failed = v
	return s
}

// SetSuccessful sets the Successful field's value.
func (s *PutAuditEventsOutput) SetSuccessful(v []*AuditEventResultEntry) *PutAuditEventsOutput {
	s.Successful = v
	return s
}

// Includes the error code and error message for events that could not be ingested
// by CloudTrail.
type ResultErrorEntry struct {
	_ struct{} `type:"structure"`

	// The error code for events that could not be ingested by CloudTrail.
	//
	// ErrorCode is a required field
	ErrorCode *string `locationName:"errorCode" min:"1" type:"string" required:"true"`

	// The message that describes the error for events that could not be ingested
	// by CloudTrail.
	//
	// ErrorMessage is a required field
	ErrorMessage *string `locationName:"errorMessage" min:"1" type:"string" required:"true"`

	// The original event ID from the source event that could not be ingested by
	// CloudTrail.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResultErrorEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResultErrorEntry) GoString() string {
	return s.String()
}

// SetErrorCode sets the ErrorCode field's value.
func (s *ResultErrorEntry) SetErrorCode(v string) *ResultErrorEntry {
	s.ErrorCode = &v
	return s
}

// SetErrorMessage sets the ErrorMessage field's value.
func (s *ResultErrorEntry) SetErrorMessage(v string) *ResultErrorEntry {
	s.ErrorMessage = &v
	return s
}

// SetId sets the Id field's value.
func (s *ResultErrorEntry) SetId(v string) *ResultErrorEntry {
	s.Id = &v
	return s
}

// The operation requested is not supported in this region or account.
type UnsupportedOperationException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s UnsupportedOperationException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s UnsupportedOperationException) GoString() string {
	return s.String()
}

func newErrorUnsupportedOperationException(v protocol.ResponseMetadata) error {
	return &UnsupportedOperationException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *UnsupportedOperationException) Code() string {
	return "UnsupportedOperationException"
}

// Message returns the exception's message.
func (s *UnsupportedOperationException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *UnsupportedOperationException) OrigErr() error {
	return nil
}

func (s *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *UnsupportedOperationException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *UnsupportedOperationException) RequestID() string {
	return s.RespMetadata.RequestID
}
