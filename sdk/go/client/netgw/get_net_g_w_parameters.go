// Code generated by go-swagger; DO NOT EDIT.

package netgw

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetGWParams creates a new GetNetGWParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetGWParams() *GetNetGWParams {
	return &GetNetGWParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetGWParamsWithTimeout creates a new GetNetGWParams object
// with the ability to set a timeout on a request.
func NewGetNetGWParamsWithTimeout(timeout time.Duration) *GetNetGWParams {
	return &GetNetGWParams{
		timeout: timeout,
	}
}

// NewGetNetGWParamsWithContext creates a new GetNetGWParams object
// with the ability to set a context for a request.
func NewGetNetGWParamsWithContext(ctx context.Context) *GetNetGWParams {
	return &GetNetGWParams{
		Context: ctx,
	}
}

// NewGetNetGWParamsWithHTTPClient creates a new GetNetGWParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetGWParamsWithHTTPClient(client *http.Client) *GetNetGWParams {
	return &GetNetGWParams{
		HTTPClient: client,
	}
}

/*
GetNetGWParams contains all the parameters to send to the API endpoint

	for the get net g w operation.

	Typically these are written to a http.Request.
*/
type GetNetGWParams struct {

	/* NetgwID.

	   The ID of the network gateway to get.
	*/
	NetgwID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get net g w params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetGWParams) WithDefaults() *GetNetGWParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get net g w params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetGWParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get net g w params
func (o *GetNetGWParams) WithTimeout(timeout time.Duration) *GetNetGWParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get net g w params
func (o *GetNetGWParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get net g w params
func (o *GetNetGWParams) WithContext(ctx context.Context) *GetNetGWParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get net g w params
func (o *GetNetGWParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get net g w params
func (o *GetNetGWParams) WithHTTPClient(client *http.Client) *GetNetGWParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get net g w params
func (o *GetNetGWParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetgwID adds the netgwID to the get net g w params
func (o *GetNetGWParams) WithNetgwID(netgwID string) *GetNetGWParams {
	o.SetNetgwID(netgwID)
	return o
}

// SetNetgwID adds the netgwId to the get net g w params
func (o *GetNetGWParams) SetNetgwID(netgwID string) {
	o.NetgwID = netgwID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetGWParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param netgwId
	if err := r.SetPathParam("netgwId", o.NetgwID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}