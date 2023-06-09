// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new instance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for instance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteInstance(params *DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstanceOK, error)

	GetAllInstances(params *GetAllInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllInstancesOK, error)

	GetInstance(params *GetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceOK, error)

	GetInstanceState(params *GetInstanceStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceStateOK, error)

	RebootInstance(params *RebootInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RebootInstanceOK, error)

	ResetInstance(params *ResetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetInstanceOK, error)

	ResumeInstance(params *ResumeInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResumeInstanceOK, error)

	ShutdownInstance(params *ShutdownInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShutdownInstanceOK, error)

	StartInstance(params *StartInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartInstanceOK, error)

	StopInstance(params *StopInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopInstanceOK, error)

	SuspendInstance(params *SuspendInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SuspendInstanceOK, error)

	UpdateInstance(params *UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInstanceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteInstance Deletes an existing virtual machine instance.
*/
func (a *Client) DeleteInstance(params *DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteInstance",
		Method:             "DELETE",
		PathPattern:        "/instance/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllInstances Returns the IDs of registered virtual machines.
*/
func (a *Client) GetAllInstances(params *GetAllInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAllInstances",
		Method:             "GET",
		PathPattern:        "/instance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAllInstances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInstance Returns the description of the virtual machine.
*/
func (a *Client) GetInstance(params *GetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInstance",
		Method:             "GET",
		PathPattern:        "/instance/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInstanceState Returns the state of the virtual machine.
*/
func (a *Client) GetInstanceState(params *GetInstanceStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInstanceState",
		Method:             "GET",
		PathPattern:        "/instance/{instanceId}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInstanceStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInstanceStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInstanceState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RebootInstance Perform a virtual machine software reboot.
*/
func (a *Client) RebootInstance(params *RebootInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RebootInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RebootInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RebootInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RebootInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RebootInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetInstance Perform a virtual machine hardware reset.
*/
func (a *Client) ResetInstance(params *ResetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResetInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResetInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResumeInstance Perform a virtual machine software PM resume.
*/
func (a *Client) ResumeInstance(params *ResumeInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResumeInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResumeInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResumeInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/resume",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResumeInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResumeInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResumeInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ShutdownInstance Initiate a software shutdown of a virtual machine.
*/
func (a *Client) ShutdownInstance(params *ShutdownInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShutdownInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShutdownInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ShutdownInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/shutdown",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ShutdownInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ShutdownInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ShutdownInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartInstance Boot up a virtual machine.
*/
func (a *Client) StartInstance(params *StartInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopInstance Initiate a hardware stop of a virtual machine.
*/
func (a *Client) StopInstance(params *StopInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SuspendInstance Perform a virtual machine software PM suspend.
*/
func (a *Client) SuspendInstance(params *SuspendInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SuspendInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSuspendInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SuspendInstance",
		Method:             "POST",
		PathPattern:        "/instance/{instanceId}/suspend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SuspendInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SuspendInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SuspendInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateInstance Updates a virtual machine configuration.
*/
func (a *Client) UpdateInstance(params *UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInstanceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateInstance",
		Method:             "PUT",
		PathPattern:        "/instance/{instanceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateInstanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateInstance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
