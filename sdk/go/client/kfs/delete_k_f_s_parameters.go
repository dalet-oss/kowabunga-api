// Code generated by go-swagger; DO NOT EDIT.

package kfs

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

// NewDeleteKFSParams creates a new DeleteKFSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteKFSParams() *DeleteKFSParams {
	return &DeleteKFSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKFSParamsWithTimeout creates a new DeleteKFSParams object
// with the ability to set a timeout on a request.
func NewDeleteKFSParamsWithTimeout(timeout time.Duration) *DeleteKFSParams {
	return &DeleteKFSParams{
		timeout: timeout,
	}
}

// NewDeleteKFSParamsWithContext creates a new DeleteKFSParams object
// with the ability to set a context for a request.
func NewDeleteKFSParamsWithContext(ctx context.Context) *DeleteKFSParams {
	return &DeleteKFSParams{
		Context: ctx,
	}
}

// NewDeleteKFSParamsWithHTTPClient creates a new DeleteKFSParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteKFSParamsWithHTTPClient(client *http.Client) *DeleteKFSParams {
	return &DeleteKFSParams{
		HTTPClient: client,
	}
}

/*
DeleteKFSParams contains all the parameters to send to the API endpoint

	for the delete k f s operation.

	Typically these are written to a http.Request.
*/
type DeleteKFSParams struct {

	/* KfsID.

	   The ID of the KFS storage volume to delete.
	*/
	KfsID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete k f s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKFSParams) WithDefaults() *DeleteKFSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete k f s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteKFSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete k f s params
func (o *DeleteKFSParams) WithTimeout(timeout time.Duration) *DeleteKFSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete k f s params
func (o *DeleteKFSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete k f s params
func (o *DeleteKFSParams) WithContext(ctx context.Context) *DeleteKFSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete k f s params
func (o *DeleteKFSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete k f s params
func (o *DeleteKFSParams) WithHTTPClient(client *http.Client) *DeleteKFSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete k f s params
func (o *DeleteKFSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKfsID adds the kfsID to the delete k f s params
func (o *DeleteKFSParams) WithKfsID(kfsID string) *DeleteKFSParams {
	o.SetKfsID(kfsID)
	return o
}

// SetKfsID adds the kfsId to the delete k f s params
func (o *DeleteKFSParams) SetKfsID(kfsID string) {
	o.KfsID = kfsID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKFSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kfsId
	if err := r.SetPathParam("kfsId", o.KfsID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}