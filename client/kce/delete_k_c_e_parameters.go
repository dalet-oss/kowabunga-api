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

// NewDeleteKCEParams creates a new DeleteKCEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKCEParams() *DeleteKCEParams {
	return &DeleteKCEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKCEParamsWithTimeout creates a new DeleteKCEParams object
// with the ability to set a timeout on a request.
func NewDeleteKCEParamsWithTimeout(timeout time.Duration) *DeleteKCEParams {
	return &DeleteKCEParams{
		timeout: timeout,
	}
}

// NewDeleteKCEParamsWithContext creates a new DeleteKCEParams object
// with the ability to set a context for a request.
func NewDeleteKCEParamsWithContext(ctx context.Context) *DeleteKCEParams {
	return &DeleteKCEParams{
		Context: ctx,
	}
}

// NewDeleteKCEParamsWithHTTPClient creates a new DeleteKCEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKCEParamsWithHTTPClient(client *http.Client) *DeleteKCEParams {
	return &DeleteKCEParams{
		HTTPClient: client,
	}
}

/*
DeleteKCEParams contains all the parameters to send to the API endpoint

	for the delete k c e operation.

	Typically these are written to a http.Request.
*/
type DeleteKCEParams struct {

	/* KceID.

	   The ID of the KCE virtual machine to delete.
	*/
	KceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete k c e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKCEParams) WithDefaults() *DeleteKCEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete k c e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKCEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete k c e params
func (o *DeleteKCEParams) WithTimeout(timeout time.Duration) *DeleteKCEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete k c e params
func (o *DeleteKCEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete k c e params
func (o *DeleteKCEParams) WithContext(ctx context.Context) *DeleteKCEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete k c e params
func (o *DeleteKCEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete k c e params
func (o *DeleteKCEParams) WithHTTPClient(client *http.Client) *DeleteKCEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete k c e params
func (o *DeleteKCEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKceID adds the kceID to the delete k c e params
func (o *DeleteKCEParams) WithKceID(kceID string) *DeleteKCEParams {
	o.SetKceID(kceID)
	return o
}

// SetKceID adds the kceId to the delete k c e params
func (o *DeleteKCEParams) SetKceID(kceID string) {
	o.KceID = kceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKCEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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