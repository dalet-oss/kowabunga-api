// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new template API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for template API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteTemplate(params *DeleteTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTemplateOK, error)

	GetAllTemplates(params *GetAllTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllTemplatesOK, error)

	GetTemplate(params *GetTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTemplateOK, error)

	UpdateTemplate(params *UpdateTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTemplateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteTemplate Deletes an existing volume template.
*/
func (a *Client) DeleteTemplate(params *DeleteTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTemplate",
		Method:             "DELETE",
		PathPattern:        "/template/{templateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTemplateReader{formats: a.formats},
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
	success, ok := result.(*DeleteTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllTemplates Returns the IDs of volume templates.
*/
func (a *Client) GetAllTemplates(params *GetAllTemplatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTemplatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAllTemplates",
		Method:             "GET",
		PathPattern:        "/template",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllTemplatesReader{formats: a.formats},
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
	success, ok := result.(*GetAllTemplatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAllTemplates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTemplate Returns a description of the volume template.
*/
func (a *Client) GetTemplate(params *GetTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTemplate",
		Method:             "GET",
		PathPattern:        "/template/{templateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTemplateReader{formats: a.formats},
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
	success, ok := result.(*GetTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateTemplate Updates a volume template configuration.
*/
func (a *Client) UpdateTemplate(params *UpdateTemplateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTemplate",
		Method:             "PUT",
		PathPattern:        "/template/{templateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTemplateReader{formats: a.formats},
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
	success, ok := result.(*UpdateTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateTemplate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}