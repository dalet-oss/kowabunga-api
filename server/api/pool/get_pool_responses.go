// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// GetPoolOKCode is the HTTP code returned for type GetPoolOK
const GetPoolOKCode int = 200

/*
GetPoolOK Returns the pool object.

swagger:response getPoolOK
*/
type GetPoolOK struct {

	/*
	  In: Body
	*/
	Payload *models.StoragePool `json:"body,omitempty"`
}

// NewGetPoolOK creates GetPoolOK with default headers values
func NewGetPoolOK() *GetPoolOK {

	return &GetPoolOK{}
}

// WithPayload adds the payload to the get pool o k response
func (o *GetPoolOK) WithPayload(payload *models.StoragePool) *GetPoolOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pool o k response
func (o *GetPoolOK) SetPayload(payload *models.StoragePool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPoolOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPoolNotFoundCode is the HTTP code returned for type GetPoolNotFound
const GetPoolNotFoundCode int = 404

/*
GetPoolNotFound Invalid pool ID was provided.

swagger:response getPoolNotFound
*/
type GetPoolNotFound struct {
}

// NewGetPoolNotFound creates GetPoolNotFound with default headers values
func NewGetPoolNotFound() *GetPoolNotFound {

	return &GetPoolNotFound{}
}

// WriteResponse to the client
func (o *GetPoolNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
