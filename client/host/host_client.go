// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new host API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for host API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteHost(params *DeleteHostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHostOK, error)

	GetAllHosts(params *GetAllHostsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllHostsOK, error)

	GetHost(params *GetHostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostOK, error)

	GetHostCaps(params *GetHostCapsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostCapsOK, error)

	GetHostInstances(params *GetHostInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostInstancesOK, error)

	UpdateHost(params *UpdateHostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateHostOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteHost Deletes an existing host.
*/
func (a *Client) DeleteHost(params *DeleteHostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteHost",
		Method:             "DELETE",
		PathPattern:        "/host/{hostId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteHostReader{formats: a.formats},
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
	success, ok := result.(*DeleteHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllHosts Returns the IDs of registered hosts.
*/
func (a *Client) GetAllHosts(params *GetAllHostsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllHostsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllHostsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAllHosts",
		Method:             "GET",
		PathPattern:        "/host",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllHostsReader{formats: a.formats},
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
	success, ok := result.(*GetAllHostsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAllHosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHost Returns a description of the host
*/
func (a *Client) GetHost(params *GetHostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHost",
		Method:             "GET",
		PathPattern:        "/host/{hostId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostReader{formats: a.formats},
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
	success, ok := result.(*GetHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostCaps Returns the host capabilities.
*/
func (a *Client) GetHostCaps(params *GetHostCapsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostCapsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostCapsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostCaps",
		Method:             "GET",
		PathPattern:        "/host/{hostId}/caps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostCapsReader{formats: a.formats},
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
	success, ok := result.(*GetHostCapsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostCaps: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostInstances Returns the UUIDs of the virtual machines running on the host.
*/
func (a *Client) GetHostInstances(params *GetHostInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetHostInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostInstances",
		Method:             "GET",
		PathPattern:        "/host/{hostId}/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostInstancesReader{formats: a.formats},
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
	success, ok := result.(*GetHostInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostInstances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateHost Updates a host configuration.
*/
func (a *Client) UpdateHost(params *UpdateHostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateHostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateHost",
		Method:             "PUT",
		PathPattern:        "/host/{hostId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateHostReader{formats: a.formats},
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
	success, ok := result.(*UpdateHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}