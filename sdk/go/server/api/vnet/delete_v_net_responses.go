// Code generated by go-swagger; DO NOT EDIT.

package vnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteVNetOKCode is the HTTP code returned for type DeleteVNetOK
const DeleteVNetOKCode int = 200

/*
DeleteVNetOK The virtual network has been successfully removed.

swagger:response deleteVNetOK
*/
type DeleteVNetOK struct {
}

// NewDeleteVNetOK creates DeleteVNetOK with default headers values
func NewDeleteVNetOK() *DeleteVNetOK {

	return &DeleteVNetOK{}
}

// WriteResponse to the client
func (o *DeleteVNetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteVNetNotFoundCode is the HTTP code returned for type DeleteVNetNotFound
const DeleteVNetNotFoundCode int = 404

/*
DeleteVNetNotFound Invalid virtual network ID was provided.

swagger:response deleteVNetNotFound
*/
type DeleteVNetNotFound struct {
}

// NewDeleteVNetNotFound creates DeleteVNetNotFound with default headers values
func NewDeleteVNetNotFound() *DeleteVNetNotFound {

	return &DeleteVNetNotFound{}
}

// WriteResponse to the client
func (o *DeleteVNetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteVNetConflictCode is the HTTP code returned for type DeleteVNetConflict
const DeleteVNetConflictCode int = 409

/*
DeleteVNetConflict The virtual network is not empty or still being referenced.

swagger:response deleteVNetConflict
*/
type DeleteVNetConflict struct {
}

// NewDeleteVNetConflict creates DeleteVNetConflict with default headers values
func NewDeleteVNetConflict() *DeleteVNetConflict {

	return &DeleteVNetConflict{}
}

// WriteResponse to the client
func (o *DeleteVNetConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// DeleteVNetInternalServerErrorCode is the HTTP code returned for type DeleteVNetInternalServerError
const DeleteVNetInternalServerErrorCode int = 500

/*
DeleteVNetInternalServerError Unable to delete virtual network.

swagger:response deleteVNetInternalServerError
*/
type DeleteVNetInternalServerError struct {
}

// NewDeleteVNetInternalServerError creates DeleteVNetInternalServerError with default headers values
func NewDeleteVNetInternalServerError() *DeleteVNetInternalServerError {

	return &DeleteVNetInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteVNetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}