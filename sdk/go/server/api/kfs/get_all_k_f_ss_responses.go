// Code generated by go-swagger; DO NOT EDIT.

package kfs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAllKFSsOKCode is the HTTP code returned for type GetAllKFSsOK
const GetAllKFSsOKCode int = 200

/*
GetAllKFSsOK Returns the an array of KFS storage volume IDs.

swagger:response getAllKFSsOK
*/
type GetAllKFSsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetAllKFSsOK creates GetAllKFSsOK with default headers values
func NewGetAllKFSsOK() *GetAllKFSsOK {

	return &GetAllKFSsOK{}
}

// WithPayload adds the payload to the get all k f ss o k response
func (o *GetAllKFSsOK) WithPayload(payload []string) *GetAllKFSsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all k f ss o k response
func (o *GetAllKFSsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllKFSsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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