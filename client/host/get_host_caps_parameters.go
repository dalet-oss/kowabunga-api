// Code generated by go-swagger; DO NOT EDIT.

package host

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

// NewGetHostCapsParams creates a new GetHostCapsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostCapsParams() *GetHostCapsParams {
	return &GetHostCapsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostCapsParamsWithTimeout creates a new GetHostCapsParams object
// with the ability to set a timeout on a request.
func NewGetHostCapsParamsWithTimeout(timeout time.Duration) *GetHostCapsParams {
	return &GetHostCapsParams{
		timeout: timeout,
	}
}

// NewGetHostCapsParamsWithContext creates a new GetHostCapsParams object
// with the ability to set a context for a request.
func NewGetHostCapsParamsWithContext(ctx context.Context) *GetHostCapsParams {
	return &GetHostCapsParams{
		Context: ctx,
	}
}

// NewGetHostCapsParamsWithHTTPClient creates a new GetHostCapsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostCapsParamsWithHTTPClient(client *http.Client) *GetHostCapsParams {
	return &GetHostCapsParams{
		HTTPClient: client,
	}
}

/*
GetHostCapsParams contains all the parameters to send to the API endpoint

	for the get host caps operation.

	Typically these are written to a http.Request.
*/
type GetHostCapsParams struct {

	/* HostID.

	   The ID of the host to query.
	*/
	HostID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host caps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostCapsParams) WithDefaults() *GetHostCapsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host caps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostCapsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host caps params
func (o *GetHostCapsParams) WithTimeout(timeout time.Duration) *GetHostCapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host caps params
func (o *GetHostCapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host caps params
func (o *GetHostCapsParams) WithContext(ctx context.Context) *GetHostCapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host caps params
func (o *GetHostCapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host caps params
func (o *GetHostCapsParams) WithHTTPClient(client *http.Client) *GetHostCapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host caps params
func (o *GetHostCapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHostID adds the hostID to the get host caps params
func (o *GetHostCapsParams) WithHostID(hostID string) *GetHostCapsParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the get host caps params
func (o *GetHostCapsParams) SetHostID(hostID string) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostCapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hostId
	if err := r.SetPathParam("hostId", o.HostID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
