// Code generated by go-swagger; DO NOT EDIT.

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateSubnetOKCode is the HTTP code returned for type UpdateSubnetOK
const UpdateSubnetOKCode int = 200

/*
UpdateSubnetOK Returns the updated subnet object.

swagger:response updateSubnetOK
*/
type UpdateSubnetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Subnet `json:"body,omitempty"`
}

// NewUpdateSubnetOK creates UpdateSubnetOK with default headers values
func NewUpdateSubnetOK() *UpdateSubnetOK {

	return &UpdateSubnetOK{}
}

// WithPayload adds the payload to the update subnet o k response
func (o *UpdateSubnetOK) WithPayload(payload *models.Subnet) *UpdateSubnetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update subnet o k response
func (o *UpdateSubnetOK) SetPayload(payload *models.Subnet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSubnetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateSubnetBadRequestCode is the HTTP code returned for type UpdateSubnetBadRequest
const UpdateSubnetBadRequestCode int = 400

/*
UpdateSubnetBadRequest Bad parameters were provided.

swagger:response updateSubnetBadRequest
*/
type UpdateSubnetBadRequest struct {
}

// NewUpdateSubnetBadRequest creates UpdateSubnetBadRequest with default headers values
func NewUpdateSubnetBadRequest() *UpdateSubnetBadRequest {

	return &UpdateSubnetBadRequest{}
}

// WriteResponse to the client
func (o *UpdateSubnetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateSubnetNotFoundCode is the HTTP code returned for type UpdateSubnetNotFound
const UpdateSubnetNotFoundCode int = 404

/*
UpdateSubnetNotFound Invalid subnet ID was provided.

swagger:response updateSubnetNotFound
*/
type UpdateSubnetNotFound struct {
}

// NewUpdateSubnetNotFound creates UpdateSubnetNotFound with default headers values
func NewUpdateSubnetNotFound() *UpdateSubnetNotFound {

	return &UpdateSubnetNotFound{}
}

// WriteResponse to the client
func (o *UpdateSubnetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
