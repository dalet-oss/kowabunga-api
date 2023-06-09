// Code generated by go-swagger; DO NOT EDIT.

package kce

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateKCEOKCode is the HTTP code returned for type UpdateKCEOK
const UpdateKCEOKCode int = 200

/*
UpdateKCEOK Returns the updated KCE virtual machine object.

swagger:response updateKCEOK
*/
type UpdateKCEOK struct {

	/*
	  In: Body
	*/
	Payload *models.KCE `json:"body,omitempty"`
}

// NewUpdateKCEOK creates UpdateKCEOK with default headers values
func NewUpdateKCEOK() *UpdateKCEOK {

	return &UpdateKCEOK{}
}

// WithPayload adds the payload to the update k c e o k response
func (o *UpdateKCEOK) WithPayload(payload *models.KCE) *UpdateKCEOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update k c e o k response
func (o *UpdateKCEOK) SetPayload(payload *models.KCE) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateKCEOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateKCEBadRequestCode is the HTTP code returned for type UpdateKCEBadRequest
const UpdateKCEBadRequestCode int = 400

/*
UpdateKCEBadRequest Bad parameters were provided.

swagger:response updateKCEBadRequest
*/
type UpdateKCEBadRequest struct {
}

// NewUpdateKCEBadRequest creates UpdateKCEBadRequest with default headers values
func NewUpdateKCEBadRequest() *UpdateKCEBadRequest {

	return &UpdateKCEBadRequest{}
}

// WriteResponse to the client
func (o *UpdateKCEBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateKCENotFoundCode is the HTTP code returned for type UpdateKCENotFound
const UpdateKCENotFoundCode int = 404

/*
UpdateKCENotFound Invalid KCE virtual machine ID was provided.

swagger:response updateKCENotFound
*/
type UpdateKCENotFound struct {
}

// NewUpdateKCENotFound creates UpdateKCENotFound with default headers values
func NewUpdateKCENotFound() *UpdateKCENotFound {

	return &UpdateKCENotFound{}
}

// WriteResponse to the client
func (o *UpdateKCENotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateKCEInternalServerErrorCode is the HTTP code returned for type UpdateKCEInternalServerError
const UpdateKCEInternalServerErrorCode int = 500

/*
UpdateKCEInternalServerError Unable to update the KCE virtual machine.

swagger:response updateKCEInternalServerError
*/
type UpdateKCEInternalServerError struct {
}

// NewUpdateKCEInternalServerError creates UpdateKCEInternalServerError with default headers values
func NewUpdateKCEInternalServerError() *UpdateKCEInternalServerError {

	return &UpdateKCEInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateKCEInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// UpdateKCEInsufficientStorageCode is the HTTP code returned for type UpdateKCEInsufficientStorage
const UpdateKCEInsufficientStorageCode int = 507

/*
UpdateKCEInsufficientStorage Requested KCE characteristics are beyond associated project's quota in place.

swagger:response updateKCEInsufficientStorage
*/
type UpdateKCEInsufficientStorage struct {
}

// NewUpdateKCEInsufficientStorage creates UpdateKCEInsufficientStorage with default headers values
func NewUpdateKCEInsufficientStorage() *UpdateKCEInsufficientStorage {

	return &UpdateKCEInsufficientStorage{}
}

// WriteResponse to the client
func (o *UpdateKCEInsufficientStorage) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(507)
}
