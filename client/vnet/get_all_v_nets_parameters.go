// Code generated by go-swagger; DO NOT EDIT.

package vnet

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

// NewGetAllVNetsParams creates a new GetAllVNetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllVNetsParams() *GetAllVNetsParams {
	return &GetAllVNetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllVNetsParamsWithTimeout creates a new GetAllVNetsParams object
// with the ability to set a timeout on a request.
func NewGetAllVNetsParamsWithTimeout(timeout time.Duration) *GetAllVNetsParams {
	return &GetAllVNetsParams{
		timeout: timeout,
	}
}

// NewGetAllVNetsParamsWithContext creates a new GetAllVNetsParams object
// with the ability to set a context for a request.
func NewGetAllVNetsParamsWithContext(ctx context.Context) *GetAllVNetsParams {
	return &GetAllVNetsParams{
		Context: ctx,
	}
}

// NewGetAllVNetsParamsWithHTTPClient creates a new GetAllVNetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllVNetsParamsWithHTTPClient(client *http.Client) *GetAllVNetsParams {
	return &GetAllVNetsParams{
		HTTPClient: client,
	}
}

/*
GetAllVNetsParams contains all the parameters to send to the API endpoint

	for the get all v nets operation.

	Typically these are written to a http.Request.
*/
type GetAllVNetsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all v nets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllVNetsParams) WithDefaults() *GetAllVNetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all v nets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllVNetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all v nets params
func (o *GetAllVNetsParams) WithTimeout(timeout time.Duration) *GetAllVNetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all v nets params
func (o *GetAllVNetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all v nets params
func (o *GetAllVNetsParams) WithContext(ctx context.Context) *GetAllVNetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all v nets params
func (o *GetAllVNetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all v nets params
func (o *GetAllVNetsParams) WithHTTPClient(client *http.Client) *GetAllVNetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all v nets params
func (o *GetAllVNetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllVNetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}