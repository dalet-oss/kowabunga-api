// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetVNetSubnetsParams creates a new GetVNetSubnetsParams object
//
// There are no default values defined in the spec.
func NewGetVNetSubnetsParams() GetVNetSubnetsParams {

	return GetVNetSubnetsParams{}
}

// GetVNetSubnetsParams contains all the bound params for the get v net subnets operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetVNetSubnets
type GetVNetSubnetsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the virtual network to query.
	  Required: true
	  In: path
	*/
	VnetID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVNetSubnetsParams() beforehand.
func (o *GetVNetSubnetsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rVnetID, rhkVnetID, _ := route.Params.GetOK("vnetId")
	if err := o.bindVnetID(rVnetID, rhkVnetID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVnetID binds and validates parameter VnetID from path.
func (o *GetVNetSubnetsParams) bindVnetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.VnetID = raw

	return nil
}
