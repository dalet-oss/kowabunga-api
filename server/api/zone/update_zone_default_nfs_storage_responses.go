// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateZoneDefaultNfsStorageOKCode is the HTTP code returned for type UpdateZoneDefaultNfsStorageOK
const UpdateZoneDefaultNfsStorageOKCode int = 200

/*
UpdateZoneDefaultNfsStorageOK Returns the updated zone resources object.

swagger:response updateZoneDefaultNfsStorageOK
*/
type UpdateZoneDefaultNfsStorageOK struct {
}

// NewUpdateZoneDefaultNfsStorageOK creates UpdateZoneDefaultNfsStorageOK with default headers values
func NewUpdateZoneDefaultNfsStorageOK() *UpdateZoneDefaultNfsStorageOK {

	return &UpdateZoneDefaultNfsStorageOK{}
}

// WriteResponse to the client
func (o *UpdateZoneDefaultNfsStorageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateZoneDefaultNfsStorageNotFoundCode is the HTTP code returned for type UpdateZoneDefaultNfsStorageNotFound
const UpdateZoneDefaultNfsStorageNotFoundCode int = 404

/*
UpdateZoneDefaultNfsStorageNotFound Invalid zone ID or NFS storage ID was provided.

swagger:response updateZoneDefaultNfsStorageNotFound
*/
type UpdateZoneDefaultNfsStorageNotFound struct {
}

// NewUpdateZoneDefaultNfsStorageNotFound creates UpdateZoneDefaultNfsStorageNotFound with default headers values
func NewUpdateZoneDefaultNfsStorageNotFound() *UpdateZoneDefaultNfsStorageNotFound {

	return &UpdateZoneDefaultNfsStorageNotFound{}
}

// WriteResponse to the client
func (o *UpdateZoneDefaultNfsStorageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateZoneDefaultNfsStorageInternalServerErrorCode is the HTTP code returned for type UpdateZoneDefaultNfsStorageInternalServerError
const UpdateZoneDefaultNfsStorageInternalServerErrorCode int = 500

/*
UpdateZoneDefaultNfsStorageInternalServerError Unable to assign the requested NFS storage as zone's default.

swagger:response updateZoneDefaultNfsStorageInternalServerError
*/
type UpdateZoneDefaultNfsStorageInternalServerError struct {
}

// NewUpdateZoneDefaultNfsStorageInternalServerError creates UpdateZoneDefaultNfsStorageInternalServerError with default headers values
func NewUpdateZoneDefaultNfsStorageInternalServerError() *UpdateZoneDefaultNfsStorageInternalServerError {

	return &UpdateZoneDefaultNfsStorageInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateZoneDefaultNfsStorageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}