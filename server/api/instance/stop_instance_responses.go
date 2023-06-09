// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// StopInstanceOKCode is the HTTP code returned for type StopInstanceOK
const StopInstanceOKCode int = 200

/*
StopInstanceOK The virtual machine has been stopped successfully.

swagger:response stopInstanceOK
*/
type StopInstanceOK struct {
}

// NewStopInstanceOK creates StopInstanceOK with default headers values
func NewStopInstanceOK() *StopInstanceOK {

	return &StopInstanceOK{}
}

// WriteResponse to the client
func (o *StopInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// StopInstanceNotFoundCode is the HTTP code returned for type StopInstanceNotFound
const StopInstanceNotFoundCode int = 404

/*
StopInstanceNotFound Invalid instance ID was provided.

swagger:response stopInstanceNotFound
*/
type StopInstanceNotFound struct {
}

// NewStopInstanceNotFound creates StopInstanceNotFound with default headers values
func NewStopInstanceNotFound() *StopInstanceNotFound {

	return &StopInstanceNotFound{}
}

// WriteResponse to the client
func (o *StopInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// StopInstanceInternalServerErrorCode is the HTTP code returned for type StopInstanceInternalServerError
const StopInstanceInternalServerErrorCode int = 500

/*
StopInstanceInternalServerError An error occurred when trying to stop the virtual machine.

swagger:response stopInstanceInternalServerError
*/
type StopInstanceInternalServerError struct {
}

// NewStopInstanceInternalServerError creates StopInstanceInternalServerError with default headers values
func NewStopInstanceInternalServerError() *StopInstanceInternalServerError {

	return &StopInstanceInternalServerError{}
}

// WriteResponse to the client
func (o *StopInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
