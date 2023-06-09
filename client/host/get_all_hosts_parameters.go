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

// NewGetAllHostsParams creates a new GetAllHostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllHostsParams() *GetAllHostsParams {
	return &GetAllHostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllHostsParamsWithTimeout creates a new GetAllHostsParams object
// with the ability to set a timeout on a request.
func NewGetAllHostsParamsWithTimeout(timeout time.Duration) *GetAllHostsParams {
	return &GetAllHostsParams{
		timeout: timeout,
	}
}

// NewGetAllHostsParamsWithContext creates a new GetAllHostsParams object
// with the ability to set a context for a request.
func NewGetAllHostsParamsWithContext(ctx context.Context) *GetAllHostsParams {
	return &GetAllHostsParams{
		Context: ctx,
	}
}

// NewGetAllHostsParamsWithHTTPClient creates a new GetAllHostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllHostsParamsWithHTTPClient(client *http.Client) *GetAllHostsParams {
	return &GetAllHostsParams{
		HTTPClient: client,
	}
}

/*
GetAllHostsParams contains all the parameters to send to the API endpoint

	for the get all hosts operation.

	Typically these are written to a http.Request.
*/
type GetAllHostsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllHostsParams) WithDefaults() *GetAllHostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllHostsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all hosts params
func (o *GetAllHostsParams) WithTimeout(timeout time.Duration) *GetAllHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all hosts params
func (o *GetAllHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all hosts params
func (o *GetAllHostsParams) WithContext(ctx context.Context) *GetAllHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all hosts params
func (o *GetAllHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all hosts params
func (o *GetAllHostsParams) WithHTTPClient(client *http.Client) *GetAllHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all hosts params
func (o *GetAllHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
