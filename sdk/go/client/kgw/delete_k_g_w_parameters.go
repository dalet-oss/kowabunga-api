// Code generated by go-swagger; DO NOT EDIT.

package kgw

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

// NewDeleteKGWParams creates a new DeleteKGWParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKGWParams() *DeleteKGWParams {
	return &DeleteKGWParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKGWParamsWithTimeout creates a new DeleteKGWParams object
// with the ability to set a timeout on a request.
func NewDeleteKGWParamsWithTimeout(timeout time.Duration) *DeleteKGWParams {
	return &DeleteKGWParams{
		timeout: timeout,
	}
}

// NewDeleteKGWParamsWithContext creates a new DeleteKGWParams object
// with the ability to set a context for a request.
func NewDeleteKGWParamsWithContext(ctx context.Context) *DeleteKGWParams {
	return &DeleteKGWParams{
		Context: ctx,
	}
}

// NewDeleteKGWParamsWithHTTPClient creates a new DeleteKGWParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKGWParamsWithHTTPClient(client *http.Client) *DeleteKGWParams {
	return &DeleteKGWParams{
		HTTPClient: client,
	}
}

/*
DeleteKGWParams contains all the parameters to send to the API endpoint

	for the delete k g w operation.

	Typically these are written to a http.Request.
*/
type DeleteKGWParams struct {

	/* KgwID.

	   The ID of the KGW to delete.
	*/
	KgwID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete k g w params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKGWParams) WithDefaults() *DeleteKGWParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete k g w params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKGWParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete k g w params
func (o *DeleteKGWParams) WithTimeout(timeout time.Duration) *DeleteKGWParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete k g w params
func (o *DeleteKGWParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete k g w params
func (o *DeleteKGWParams) WithContext(ctx context.Context) *DeleteKGWParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete k g w params
func (o *DeleteKGWParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete k g w params
func (o *DeleteKGWParams) WithHTTPClient(client *http.Client) *DeleteKGWParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete k g w params
func (o *DeleteKGWParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKgwID adds the kgwID to the delete k g w params
func (o *DeleteKGWParams) WithKgwID(kgwID string) *DeleteKGWParams {
	o.SetKgwID(kgwID)
	return o
}

// SetKgwID adds the kgwId to the delete k g w params
func (o *DeleteKGWParams) SetKgwID(kgwID string) {
	o.KgwID = kgwID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKGWParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kgwId
	if err := r.SetPathParam("kgwId", o.KgwID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}