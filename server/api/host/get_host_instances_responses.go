// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHostInstancesOKCode is the HTTP code returned for type GetHostInstancesOK
const GetHostInstancesOKCode int = 200

/*
GetHostInstancesOK Returns the an array of virtual machine UUIDs.

swagger:response getHostInstancesOK
*/
type GetHostInstancesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetHostInstancesOK creates GetHostInstancesOK with default headers values
func NewGetHostInstancesOK() *GetHostInstancesOK {

	return &GetHostInstancesOK{}
}

// WithPayload adds the payload to the get host instances o k response
func (o *GetHostInstancesOK) WithPayload(payload []string) *GetHostInstancesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get host instances o k response
func (o *GetHostInstancesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHostInstancesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetHostInstancesNotFoundCode is the HTTP code returned for type GetHostInstancesNotFound
const GetHostInstancesNotFoundCode int = 404

/*
GetHostInstancesNotFound Invalid host ID was provided.

swagger:response getHostInstancesNotFound
*/
type GetHostInstancesNotFound struct {
}

// NewGetHostInstancesNotFound creates GetHostInstancesNotFound with default headers values
func NewGetHostInstancesNotFound() *GetHostInstancesNotFound {

	return &GetHostInstancesNotFound{}
}

// WriteResponse to the client
func (o *GetHostInstancesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}