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

// NewResetKCEParams creates a new ResetKCEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetKCEParams() *ResetKCEParams {
	return &ResetKCEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetKCEParamsWithTimeout creates a new ResetKCEParams object
// with the ability to set a timeout on a request.
func NewResetKCEParamsWithTimeout(timeout time.Duration) *ResetKCEParams {
	return &ResetKCEParams{
		timeout: timeout,
	}
}

// NewResetKCEParamsWithContext creates a new ResetKCEParams object
// with the ability to set a context for a request.
func NewResetKCEParamsWithContext(ctx context.Context) *ResetKCEParams {
	return &ResetKCEParams{
		Context: ctx,
	}
}

// NewResetKCEParamsWithHTTPClient creates a new ResetKCEParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetKCEParamsWithHTTPClient(client *http.Client) *ResetKCEParams {
	return &ResetKCEParams{
		HTTPClient: client,
	}
}

/*
ResetKCEParams contains all the parameters to send to the API endpoint

	for the reset k c e operation.

	Typically these are written to a http.Request.
*/
type ResetKCEParams struct {

	/* KceID.

	   The ID of the KCE virtual machine to query.
	*/
	KceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset k c e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetKCEParams) WithDefaults() *ResetKCEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset k c e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetKCEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset k c e params
func (o *ResetKCEParams) WithTimeout(timeout time.Duration) *ResetKCEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset k c e params
func (o *ResetKCEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset k c e params
func (o *ResetKCEParams) WithContext(ctx context.Context) *ResetKCEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset k c e params
func (o *ResetKCEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset k c e params
func (o *ResetKCEParams) WithHTTPClient(client *http.Client) *ResetKCEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset k c e params
func (o *ResetKCEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKceID adds the kceID to the reset k c e params
func (o *ResetKCEParams) WithKceID(kceID string) *ResetKCEParams {
	o.SetKceID(kceID)
	return o
}

// SetKceID adds the kceId to the reset k c e params
func (o *ResetKCEParams) SetKceID(kceID string) {
	o.KceID = kceID
}

// WriteToRequest writes these params to a swagger request
func (o *ResetKCEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
