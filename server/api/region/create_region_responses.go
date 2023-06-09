// Code generated by go-swagger; DO NOT EDIT.

package region

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// CreateRegionCreatedCode is the HTTP code returned for type CreateRegionCreated
const CreateRegionCreatedCode int = 201

/*
CreateRegionCreated Returns the newly created region object.

swagger:response createRegionCreated
*/
type CreateRegionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Region `json:"body,omitempty"`
}

// NewCreateRegionCreated creates CreateRegionCreated with default headers values
func NewCreateRegionCreated() *CreateRegionCreated {

	return &CreateRegionCreated{}
}

// WithPayload adds the payload to the create region created response
func (o *CreateRegionCreated) WithPayload(payload *models.Region) *CreateRegionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create region created response
func (o *CreateRegionCreated) SetPayload(payload *models.Region) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRegionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRegionBadRequestCode is the HTTP code returned for type CreateRegionBadRequest
const CreateRegionBadRequestCode int = 400

/*
CreateRegionBadRequest Bad parameters were provided.

swagger:response createRegionBadRequest
*/
type CreateRegionBadRequest struct {
}

// NewCreateRegionBadRequest creates CreateRegionBadRequest with default headers values
func NewCreateRegionBadRequest() *CreateRegionBadRequest {

	return &CreateRegionBadRequest{}
}

// WriteResponse to the client
func (o *CreateRegionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateRegionConflictCode is the HTTP code returned for type CreateRegionConflict
const CreateRegionConflictCode int = 409

/*
CreateRegionConflict Region already exists.

swagger:response createRegionConflict
*/
type CreateRegionConflict struct {
}

// NewCreateRegionConflict creates CreateRegionConflict with default headers values
func NewCreateRegionConflict() *CreateRegionConflict {

	return &CreateRegionConflict{}
}

// WriteResponse to the client
func (o *CreateRegionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// CreateRegionInternalServerErrorCode is the HTTP code returned for type CreateRegionInternalServerError
const CreateRegionInternalServerErrorCode int = 500

/*
CreateRegionInternalServerError Unable to create the region.

swagger:response createRegionInternalServerError
*/
type CreateRegionInternalServerError struct {
}

// NewCreateRegionInternalServerError creates CreateRegionInternalServerError with default headers values
func NewCreateRegionInternalServerError() *CreateRegionInternalServerError {

	return &CreateRegionInternalServerError{}
}

// WriteResponse to the client
func (o *CreateRegionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
