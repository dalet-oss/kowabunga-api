// Code generated by go-swagger; DO NOT EDIT.

package nfs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetNfsKfsOKCode is the HTTP code returned for type GetNfsKfsOK
const GetNfsKfsOKCode int = 200

/*
GetNfsKfsOK Returns an array of KFS storage volume IDs.

swagger:response getNfsKfsOK
*/
type GetNfsKfsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetNfsKfsOK creates GetNfsKfsOK with default headers values
func NewGetNfsKfsOK() *GetNfsKfsOK {

	return &GetNfsKfsOK{}
}

// WithPayload adds the payload to the get nfs kfs o k response
func (o *GetNfsKfsOK) WithPayload(payload []string) *GetNfsKfsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nfs kfs o k response
func (o *GetNfsKfsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNfsKfsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetNfsKfsNotFoundCode is the HTTP code returned for type GetNfsKfsNotFound
const GetNfsKfsNotFoundCode int = 404

/*
GetNfsKfsNotFound Invalid NFS storage ID was provided.

swagger:response getNfsKfsNotFound
*/
type GetNfsKfsNotFound struct {
}

// NewGetNfsKfsNotFound creates GetNfsKfsNotFound with default headers values
func NewGetNfsKfsNotFound() *GetNfsKfsNotFound {

	return &GetNfsKfsNotFound{}
}

// WriteResponse to the client
func (o *GetNfsKfsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}