// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// NewResetInstanceParams creates a new ResetInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetInstanceParams() *ResetInstanceParams {
	return &ResetInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetInstanceParamsWithTimeout creates a new ResetInstanceParams object
// with the ability to set a timeout on a request.
func NewResetInstanceParamsWithTimeout(timeout time.Duration) *ResetInstanceParams {
	return &ResetInstanceParams{
		timeout: timeout,
	}
}

// NewResetInstanceParamsWithContext creates a new ResetInstanceParams object
// with the ability to set a context for a request.
func NewResetInstanceParamsWithContext(ctx context.Context) *ResetInstanceParams {
	return &ResetInstanceParams{
		Context: ctx,
	}
}

// NewResetInstanceParamsWithHTTPClient creates a new ResetInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetInstanceParamsWithHTTPClient(client *http.Client) *ResetInstanceParams {
	return &ResetInstanceParams{
		HTTPClient: client,
	}
}

/*
ResetInstanceParams contains all the parameters to send to the API endpoint

	for the reset instance operation.

	Typically these are written to a http.Request.
*/
type ResetInstanceParams struct {

	/* InstanceID.

	   The ID of the virtual machine to query.
	*/
	InstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetInstanceParams) WithDefaults() *ResetInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset instance params
func (o *ResetInstanceParams) WithTimeout(timeout time.Duration) *ResetInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset instance params
func (o *ResetInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset instance params
func (o *ResetInstanceParams) WithContext(ctx context.Context) *ResetInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset instance params
func (o *ResetInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset instance params
func (o *ResetInstanceParams) WithHTTPClient(client *http.Client) *ResetInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset instance params
func (o *ResetInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInstanceID adds the instanceID to the reset instance params
func (o *ResetInstanceParams) WithInstanceID(instanceID string) *ResetInstanceParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the reset instance params
func (o *ResetInstanceParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param instanceId
	if err := r.SetPathParam("instanceId", o.InstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}