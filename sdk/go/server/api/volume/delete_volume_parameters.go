// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteVolumeParams creates a new DeleteVolumeParams object
//
// There are no default values defined in the spec.
func NewDeleteVolumeParams() DeleteVolumeParams {

	return DeleteVolumeParams{}
}

// DeleteVolumeParams contains all the bound params for the delete volume operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteVolume
type DeleteVolumeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The ID of the storage volume to delete.
	  Required: true
	  In: path
	*/
	VolumeID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteVolumeParams() beforehand.
func (o *DeleteVolumeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rVolumeID, rhkVolumeID, _ := route.Params.GetOK("volumeId")
	if err := o.bindVolumeID(rVolumeID, rhkVolumeID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVolumeID binds and validates parameter VolumeID from path.
func (o *DeleteVolumeParams) bindVolumeID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.VolumeID = raw

	return nil
}