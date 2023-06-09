// Code generated by go-swagger; DO NOT EDIT.

package kce

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

// NewStopKCEParams creates a new StopKCEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopKCEParams() *StopKCEParams {
	return &StopKCEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopKCEParamsWithTimeout creates a new StopKCEParams object
// with the ability to set a timeout on a request.
func NewStopKCEParamsWithTimeout(timeout time.Duration) *StopKCEParams {
	return &StopKCEParams{
		timeout: timeout,
	}
}

// NewStopKCEParamsWithContext creates a new StopKCEParams object
// with the ability to set a context for a request.
func NewStopKCEParamsWithContext(ctx context.Context) *StopKCEParams {
	return &StopKCEParams{
		Context: ctx,
	}
}

// NewStopKCEParamsWithHTTPClient creates a new StopKCEParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopKCEParamsWithHTTPClient(client *http.Client) *StopKCEParams {
	return &StopKCEParams{
		HTTPClient: client,
	}
}

/*
StopKCEParams contains all the parameters to send to the API endpoint

	for the stop k c e operation.

	Typically these are written to a http.Request.
*/
type StopKCEParams struct {

	/* KceID.

	   The ID of the KCE virtual machine to query.
	*/
	KceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop k c e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopKCEParams) WithDefaults() *StopKCEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop k c e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopKCEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop k c e params
func (o *StopKCEParams) WithTimeout(timeout time.Duration) *StopKCEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop k c e params
func (o *StopKCEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop k c e params
func (o *StopKCEParams) WithContext(ctx context.Context) *StopKCEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop k c e params
func (o *StopKCEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop k c e params
func (o *StopKCEParams) WithHTTPClient(client *http.Client) *StopKCEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop k c e params
func (o *StopKCEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKceID adds the kceID to the stop k c e params
func (o *StopKCEParams) WithKceID(kceID string) *StopKCEParams {
	o.SetKceID(kceID)
	return o
}

// SetKceID adds the kceId to the stop k c e params
func (o *StopKCEParams) SetKceID(kceID string) {
	o.KceID = kceID
}

// WriteToRequest writes these params to a swagger request
func (o *StopKCEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kceId
	if err := r.SetPathParam("kceId", o.KceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
