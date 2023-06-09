// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// CreateVNetCreatedCode is the HTTP code returned for type CreateVNetCreated
const CreateVNetCreatedCode int = 201

/*
CreateVNetCreated Returns the newly created virtual network object.

swagger:response createVNetCreated
*/
type CreateVNetCreated struct {

	/*
	  In: Body
	*/
	Payload *models.VNet `json:"body,omitempty"`
}

// NewCreateVNetCreated creates CreateVNetCreated with default headers values
func NewCreateVNetCreated() *CreateVNetCreated {

	return &CreateVNetCreated{}
}

// WithPayload adds the payload to the create v net created response
func (o *CreateVNetCreated) WithPayload(payload *models.VNet) *CreateVNetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create v net created response
func (o *CreateVNetCreated) SetPayload(payload *models.VNet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVNetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVNetBadRequestCode is the HTTP code returned for type CreateVNetBadRequest
const CreateVNetBadRequestCode int = 400

/*
CreateVNetBadRequest Bad parameters were provided.

swagger:response createVNetBadRequest
*/
type CreateVNetBadRequest struct {
}

// NewCreateVNetBadRequest creates CreateVNetBadRequest with default headers values
func NewCreateVNetBadRequest() *CreateVNetBadRequest {

	return &CreateVNetBadRequest{}
}

// WriteResponse to the client
func (o *CreateVNetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateVNetNotFoundCode is the HTTP code returned for type CreateVNetNotFound
const CreateVNetNotFoundCode int = 404

/*
CreateVNetNotFound Invalid zone ID was provided.

swagger:response createVNetNotFound
*/
type CreateVNetNotFound struct {
}

// NewCreateVNetNotFound creates CreateVNetNotFound with default headers values
func NewCreateVNetNotFound() *CreateVNetNotFound {

	return &CreateVNetNotFound{}
}

// WriteResponse to the client
func (o *CreateVNetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// CreateVNetConflictCode is the HTTP code returned for type CreateVNetConflict
const CreateVNetConflictCode int = 409

/*
CreateVNetConflict Virtual network already exists.

swagger:response createVNetConflict
*/
type CreateVNetConflict struct {
}

// NewCreateVNetConflict creates CreateVNetConflict with default headers values
func NewCreateVNetConflict() *CreateVNetConflict {

	return &CreateVNetConflict{}
}

// WriteResponse to the client
func (o *CreateVNetConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// CreateVNetInternalServerErrorCode is the HTTP code returned for type CreateVNetInternalServerError
const CreateVNetInternalServerErrorCode int = 500

/*
CreateVNetInternalServerError Unable to create virtual network.

swagger:response createVNetInternalServerError
*/
type CreateVNetInternalServerError struct {
}

// NewCreateVNetInternalServerError creates CreateVNetInternalServerError with default headers values
func NewCreateVNetInternalServerError() *CreateVNetInternalServerError {

	return &CreateVNetInternalServerError{}
}

// WriteResponse to the client
func (o *CreateVNetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
