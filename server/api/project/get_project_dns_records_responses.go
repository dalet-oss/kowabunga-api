// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetProjectDNSRecordsOKCode is the HTTP code returned for type GetProjectDNSRecordsOK
const GetProjectDNSRecordsOKCode int = 200

/*
GetProjectDNSRecordsOK Returns an array of DNS record IDs.

swagger:response getProjectDnsRecordsOK
*/
type GetProjectDNSRecordsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetProjectDNSRecordsOK creates GetProjectDNSRecordsOK with default headers values
func NewGetProjectDNSRecordsOK() *GetProjectDNSRecordsOK {

	return &GetProjectDNSRecordsOK{}
}

// WithPayload adds the payload to the get project Dns records o k response
func (o *GetProjectDNSRecordsOK) WithPayload(payload []string) *GetProjectDNSRecordsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project Dns records o k response
func (o *GetProjectDNSRecordsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDNSRecordsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectDNSRecordsNotFoundCode is the HTTP code returned for type GetProjectDNSRecordsNotFound
const GetProjectDNSRecordsNotFoundCode int = 404

/*
GetProjectDNSRecordsNotFound Invalid project ID was provided.

swagger:response getProjectDnsRecordsNotFound
*/
type GetProjectDNSRecordsNotFound struct {
}

// NewGetProjectDNSRecordsNotFound creates GetProjectDNSRecordsNotFound with default headers values
func NewGetProjectDNSRecordsNotFound() *GetProjectDNSRecordsNotFound {

	return &GetProjectDNSRecordsNotFound{}
}

// WriteResponse to the client
func (o *GetProjectDNSRecordsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
