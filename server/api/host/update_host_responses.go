// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateHostOKCode is the HTTP code returned for type UpdateHostOK
const UpdateHostOKCode int = 200

/*
UpdateHostOK Returns the updated host object.

swagger:response updateHostOK
*/
type UpdateHostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Host `json:"body,omitempty"`
}

// NewUpdateHostOK creates UpdateHostOK with default headers values
func NewUpdateHostOK() *UpdateHostOK {

	return &UpdateHostOK{}
}

// WithPayload adds the payload to the update host o k response
func (o *UpdateHostOK) WithPayload(payload *models.Host) *UpdateHostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update host o k response
func (o *UpdateHostOK) SetPayload(payload *models.Host) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateHostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateHostBadRequestCode is the HTTP code returned for type UpdateHostBadRequest
const UpdateHostBadRequestCode int = 400

/*
UpdateHostBadRequest Bad parameters were provided.

swagger:response updateHostBadRequest
*/
type UpdateHostBadRequest struct {
}

// NewUpdateHostBadRequest creates UpdateHostBadRequest with default headers values
func NewUpdateHostBadRequest() *UpdateHostBadRequest {

	return &UpdateHostBadRequest{}
}

// WriteResponse to the client
func (o *UpdateHostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateHostNotFoundCode is the HTTP code returned for type UpdateHostNotFound
const UpdateHostNotFoundCode int = 404

/*
UpdateHostNotFound Invalid host ID was provided.

swagger:response updateHostNotFound
*/
type UpdateHostNotFound struct {
}

// NewUpdateHostNotFound creates UpdateHostNotFound with default headers values
func NewUpdateHostNotFound() *UpdateHostNotFound {

	return &UpdateHostNotFound{}
}

// WriteResponse to the client
func (o *UpdateHostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}