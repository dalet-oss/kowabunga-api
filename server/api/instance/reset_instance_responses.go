// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ResetInstanceOKCode is the HTTP code returned for type ResetInstanceOK
const ResetInstanceOKCode int = 200

/*
ResetInstanceOK The virtual machine has been reseted successfully.

swagger:response resetInstanceOK
*/
type ResetInstanceOK struct {
}

// NewResetInstanceOK creates ResetInstanceOK with default headers values
func NewResetInstanceOK() *ResetInstanceOK {

	return &ResetInstanceOK{}
}

// WriteResponse to the client
func (o *ResetInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ResetInstanceNotFoundCode is the HTTP code returned for type ResetInstanceNotFound
const ResetInstanceNotFoundCode int = 404

/*
ResetInstanceNotFound Invalid instance ID was provided.

swagger:response resetInstanceNotFound
*/
type ResetInstanceNotFound struct {
}

// NewResetInstanceNotFound creates ResetInstanceNotFound with default headers values
func NewResetInstanceNotFound() *ResetInstanceNotFound {

	return &ResetInstanceNotFound{}
}

// WriteResponse to the client
func (o *ResetInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ResetInstanceInternalServerErrorCode is the HTTP code returned for type ResetInstanceInternalServerError
const ResetInstanceInternalServerErrorCode int = 500

/*
ResetInstanceInternalServerError An error occurred when trying to reset the virtual machine.

swagger:response resetInstanceInternalServerError
*/
type ResetInstanceInternalServerError struct {
}

// NewResetInstanceInternalServerError creates ResetInstanceInternalServerError with default headers values
func NewResetInstanceInternalServerError() *ResetInstanceInternalServerError {

	return &ResetInstanceInternalServerError{}
}

// WriteResponse to the client
func (o *ResetInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
