// Code generated by go-swagger; DO NOT EDIT.

package nfs

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

// NewDeleteNfsStorageParams creates a new DeleteNfsStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNfsStorageParams() *DeleteNfsStorageParams {
	return &DeleteNfsStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNfsStorageParamsWithTimeout creates a new DeleteNfsStorageParams object
// with the ability to set a timeout on a request.
func NewDeleteNfsStorageParamsWithTimeout(timeout time.Duration) *DeleteNfsStorageParams {
	return &DeleteNfsStorageParams{
		timeout: timeout,
	}
}

// NewDeleteNfsStorageParamsWithContext creates a new DeleteNfsStorageParams object
// with the ability to set a context for a request.
func NewDeleteNfsStorageParamsWithContext(ctx context.Context) *DeleteNfsStorageParams {
	return &DeleteNfsStorageParams{
		Context: ctx,
	}
}

// NewDeleteNfsStorageParamsWithHTTPClient creates a new DeleteNfsStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNfsStorageParamsWithHTTPClient(client *http.Client) *DeleteNfsStorageParams {
	return &DeleteNfsStorageParams{
		HTTPClient: client,
	}
}

/*
DeleteNfsStorageParams contains all the parameters to send to the API endpoint

	for the delete nfs storage operation.

	Typically these are written to a http.Request.
*/
type DeleteNfsStorageParams struct {

	/* NfsID.

	   The ID of the NFS storage to get.
	*/
	NfsID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete nfs storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNfsStorageParams) WithDefaults() *DeleteNfsStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete nfs storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNfsStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete nfs storage params
func (o *DeleteNfsStorageParams) WithTimeout(timeout time.Duration) *DeleteNfsStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nfs storage params
func (o *DeleteNfsStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nfs storage params
func (o *DeleteNfsStorageParams) WithContext(ctx context.Context) *DeleteNfsStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nfs storage params
func (o *DeleteNfsStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nfs storage params
func (o *DeleteNfsStorageParams) WithHTTPClient(client *http.Client) *DeleteNfsStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nfs storage params
func (o *DeleteNfsStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNfsID adds the nfsID to the delete nfs storage params
func (o *DeleteNfsStorageParams) WithNfsID(nfsID string) *DeleteNfsStorageParams {
	o.SetNfsID(nfsID)
	return o
}

// SetNfsID adds the nfsId to the delete nfs storage params
func (o *DeleteNfsStorageParams) SetNfsID(nfsID string) {
	o.NfsID = nfsID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNfsStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nfsId
	if err := r.SetPathParam("nfsId", o.NfsID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}