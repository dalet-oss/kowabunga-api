// Code generated by go-swagger; DO NOT EDIT.

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetSubnetAdaptersOKCode is the HTTP code returned for type GetSubnetAdaptersOK
const GetSubnetAdaptersOKCode int = 200

/*
GetSubnetAdaptersOK Returns an array of network adapter IDs.

swagger:response getSubnetAdaptersOK
*/
type GetSubnetAdaptersOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetSubnetAdaptersOK creates GetSubnetAdaptersOK with default headers values
func NewGetSubnetAdaptersOK() *GetSubnetAdaptersOK {

	return &GetSubnetAdaptersOK{}
}

// WithPayload adds the payload to the get subnet adapters o k response
func (o *GetSubnetAdaptersOK) WithPayload(payload []string) *GetSubnetAdaptersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subnet adapters o k response
func (o *GetSubnetAdaptersOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubnetAdaptersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetSubnetAdaptersNotFoundCode is the HTTP code returned for type GetSubnetAdaptersNotFound
const GetSubnetAdaptersNotFoundCode int = 404

/*
GetSubnetAdaptersNotFound Invalid subnet ID was provided.

swagger:response getSubnetAdaptersNotFound
*/
type GetSubnetAdaptersNotFound struct {
}

// NewGetSubnetAdaptersNotFound creates GetSubnetAdaptersNotFound with default headers values
func NewGetSubnetAdaptersNotFound() *GetSubnetAdaptersNotFound {

	return &GetSubnetAdaptersNotFound{}
}

// WriteResponse to the client
func (o *GetSubnetAdaptersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
