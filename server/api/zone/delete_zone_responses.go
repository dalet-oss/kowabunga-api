// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteZoneOKCode is the HTTP code returned for type DeleteZoneOK
const DeleteZoneOKCode int = 200

/*
DeleteZoneOK The zone has been successfully removed.

swagger:response deleteZoneOK
*/
type DeleteZoneOK struct {
}

// NewDeleteZoneOK creates DeleteZoneOK with default headers values
func NewDeleteZoneOK() *DeleteZoneOK {

	return &DeleteZoneOK{}
}

// WriteResponse to the client
func (o *DeleteZoneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteZoneNotFoundCode is the HTTP code returned for type DeleteZoneNotFound
const DeleteZoneNotFoundCode int = 404

/*
DeleteZoneNotFound Invalid zone ID was provided.

swagger:response deleteZoneNotFound
*/
type DeleteZoneNotFound struct {
}

// NewDeleteZoneNotFound creates DeleteZoneNotFound with default headers values
func NewDeleteZoneNotFound() *DeleteZoneNotFound {

	return &DeleteZoneNotFound{}
}

// WriteResponse to the client
func (o *DeleteZoneNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteZoneConflictCode is the HTTP code returned for type DeleteZoneConflict
const DeleteZoneConflictCode int = 409

/*
DeleteZoneConflict The zone is not empty or still being referenced.

swagger:response deleteZoneConflict
*/
type DeleteZoneConflict struct {
}

// NewDeleteZoneConflict creates DeleteZoneConflict with default headers values
func NewDeleteZoneConflict() *DeleteZoneConflict {

	return &DeleteZoneConflict{}
}

// WriteResponse to the client
func (o *DeleteZoneConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// DeleteZoneInternalServerErrorCode is the HTTP code returned for type DeleteZoneInternalServerError
const DeleteZoneInternalServerErrorCode int = 500

/*
DeleteZoneInternalServerError Unable to delete zone.

swagger:response deleteZoneInternalServerError
*/
type DeleteZoneInternalServerError struct {
}

// NewDeleteZoneInternalServerError creates DeleteZoneInternalServerError with default headers values
func NewDeleteZoneInternalServerError() *DeleteZoneInternalServerError {

	return &DeleteZoneInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteZoneInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
