// Code generated by go-swagger; DO NOT EDIT.

package kgw

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteKGWParams creates a new DeleteKGWParams object
//
// There are no default values defined in the spec.
func NewDeleteKGWParams() DeleteKGWParams {

	return DeleteKGWParams{}
}

// DeleteKGWParams contains all the bound params for the delete k g w operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteKGW
type DeleteKGWParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the KGW to delete.
	  Required: true
	  In: path
	*/
	KgwID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteKGWParams() beforehand.
func (o *DeleteKGWParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rKgwID, rhkKgwID, _ := route.Params.GetOK("kgwId")
	if err := o.bindKgwID(rKgwID, rhkKgwID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindKgwID binds and validates parameter KgwID from path.
func (o *DeleteKGWParams) bindKgwID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.KgwID = raw

	return nil
}