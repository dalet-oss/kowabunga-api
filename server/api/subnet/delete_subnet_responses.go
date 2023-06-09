// Code generated by go-swagger; DO NOT EDIT.

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteSubnetOKCode is the HTTP code returned for type DeleteSubnetOK
const DeleteSubnetOKCode int = 200

/*
DeleteSubnetOK The subnet has been successfully removed.

swagger:response deleteSubnetOK
*/
type DeleteSubnetOK struct {
}

// NewDeleteSubnetOK creates DeleteSubnetOK with default headers values
func NewDeleteSubnetOK() *DeleteSubnetOK {

	return &DeleteSubnetOK{}
}

// WriteResponse to the client
func (o *DeleteSubnetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteSubnetNotFoundCode is the HTTP code returned for type DeleteSubnetNotFound
const DeleteSubnetNotFoundCode int = 404

/*
DeleteSubnetNotFound Invalid subnet ID was provided.

swagger:response deleteSubnetNotFound
*/
type DeleteSubnetNotFound struct {
}

// NewDeleteSubnetNotFound creates DeleteSubnetNotFound with default headers values
func NewDeleteSubnetNotFound() *DeleteSubnetNotFound {

	return &DeleteSubnetNotFound{}
}

// WriteResponse to the client
func (o *DeleteSubnetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteSubnetConflictCode is the HTTP code returned for type DeleteSubnetConflict
const DeleteSubnetConflictCode int = 409

/*
DeleteSubnetConflict The subnet is not empty or still being referenced.

swagger:response deleteSubnetConflict
*/
type DeleteSubnetConflict struct {
}

// NewDeleteSubnetConflict creates DeleteSubnetConflict with default headers values
func NewDeleteSubnetConflict() *DeleteSubnetConflict {

	return &DeleteSubnetConflict{}
}

// WriteResponse to the client
func (o *DeleteSubnetConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// DeleteSubnetInternalServerErrorCode is the HTTP code returned for type DeleteSubnetInternalServerError
const DeleteSubnetInternalServerErrorCode int = 500

/*
DeleteSubnetInternalServerError Unable to delete subnet.

swagger:response deleteSubnetInternalServerError
*/
type DeleteSubnetInternalServerError struct {
}

// NewDeleteSubnetInternalServerError creates DeleteSubnetInternalServerError with default headers values
func NewDeleteSubnetInternalServerError() *DeleteSubnetInternalServerError {

	return &DeleteSubnetInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteSubnetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
