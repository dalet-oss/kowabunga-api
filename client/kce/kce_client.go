// Code generated by go-swagger; DO NOT EDIT.

package kce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new kce API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kce API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteKCE(params *DeleteKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKCEOK, error)

	GetAllKCEs(params *GetAllKCEsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllKCEsOK, error)

	GetKCE(params *GetKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKCEOK, error)

	GetKCEState(params *GetKCEStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKCEStateOK, error)

	RebootKCE(params *RebootKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RebootKCEOK, error)

	ResetKCE(params *ResetKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetKCEOK, error)

	ResumeKCE(params *ResumeKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResumeKCEOK, error)

	ShutdownKCE(params *ShutdownKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShutdownKCEOK, error)

	StartKCE(params *StartKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartKCEOK, error)

	StopKCE(params *StopKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopKCEOK, error)

	SuspendKCE(params *SuspendKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SuspendKCEOK, error)

	UpdateKCE(params *UpdateKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateKCEOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteKCE Deletes an existing KCE virtual machine.
*/
func (a *Client) DeleteKCE(params *DeleteKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteKCE",
		Method:             "DELETE",
		PathPattern:        "/kce/{kceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteKCEReader{formats: a.formats},
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
	success, ok := result.(*DeleteKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllKCEs Returns the IDs of registered KCE virtual machines.
*/
func (a *Client) GetAllKCEs(params *GetAllKCEsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllKCEsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllKCEsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAllKCEs",
		Method:             "GET",
		PathPattern:        "/kce",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllKCEsReader{formats: a.formats},
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
	success, ok := result.(*GetAllKCEsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAllKCEs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetKCE Returns the description of the KCE virtual machine.
*/
func (a *Client) GetKCE(params *GetKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetKCE",
		Method:             "GET",
		PathPattern:        "/kce/{kceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetKCEReader{formats: a.formats},
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
	success, ok := result.(*GetKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetKCEState Returns the state of the KCE virtual machine.
*/
func (a *Client) GetKCEState(params *GetKCEStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKCEStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKCEStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetKCEState",
		Method:             "GET",
		PathPattern:        "/kce/{kceId}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetKCEStateReader{formats: a.formats},
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
	success, ok := result.(*GetKCEStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetKCEState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RebootKCE Perform a KCE virtual machine software reboot.
*/
func (a *Client) RebootKCE(params *RebootKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RebootKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RebootKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RebootKCEReader{formats: a.formats},
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
	success, ok := result.(*RebootKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RebootKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResetKCE Perform a KCE virtual machine hardware reset.
*/
func (a *Client) ResetKCE(params *ResetKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResetKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResetKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetKCEReader{formats: a.formats},
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
	success, ok := result.(*ResetKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResetKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ResumeKCE Perform a KCE virtual machine software PM resume.
*/
func (a *Client) ResumeKCE(params *ResumeKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResumeKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResumeKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResumeKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/resume",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResumeKCEReader{formats: a.formats},
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
	success, ok := result.(*ResumeKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ResumeKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ShutdownKCE Initiate a software shutdown of a KCE virtual machine.
*/
func (a *Client) ShutdownKCE(params *ShutdownKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ShutdownKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShutdownKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ShutdownKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/shutdown",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ShutdownKCEReader{formats: a.formats},
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
	success, ok := result.(*ShutdownKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ShutdownKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StartKCE Boot up a KCE virtual machine.
*/
func (a *Client) StartKCE(params *StartKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartKCEReader{formats: a.formats},
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
	success, ok := result.(*StartKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StartKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopKCE Initiate a hardware stop of a KCE virtual machine.
*/
func (a *Client) StopKCE(params *StopKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StopKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopKCEReader{formats: a.formats},
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
	success, ok := result.(*StopKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StopKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SuspendKCE Perform a KCE virtual machine software PM suspend.
*/
func (a *Client) SuspendKCE(params *SuspendKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SuspendKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSuspendKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SuspendKCE",
		Method:             "POST",
		PathPattern:        "/kce/{kceId}/suspend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SuspendKCEReader{formats: a.formats},
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
	success, ok := result.(*SuspendKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SuspendKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateKCE Updates a KCE virtual machine configuration.
*/
func (a *Client) UpdateKCE(params *UpdateKCEParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateKCEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateKCEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateKCE",
		Method:             "PUT",
		PathPattern:        "/kce/{kceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateKCEReader{formats: a.formats},
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
	success, ok := result.(*UpdateKCEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateKCE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
