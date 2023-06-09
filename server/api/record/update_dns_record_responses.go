// Code generated by go-swagger; DO NOT EDIT.

package record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateDNSRecordOKCode is the HTTP code returned for type UpdateDNSRecordOK
const UpdateDNSRecordOKCode int = 200

/*
UpdateDNSRecordOK Returns the updated DNS record object.

swagger:response updateDnsRecordOK
*/
type UpdateDNSRecordOK struct {

	/*
	  In: Body
	*/
	Payload *models.DNSRecord `json:"body,omitempty"`
}

// NewUpdateDNSRecordOK creates UpdateDNSRecordOK with default headers values
func NewUpdateDNSRecordOK() *UpdateDNSRecordOK {

	return &UpdateDNSRecordOK{}
}

// WithPayload adds the payload to the update Dns record o k response
func (o *UpdateDNSRecordOK) WithPayload(payload *models.DNSRecord) *UpdateDNSRecordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Dns record o k response
func (o *UpdateDNSRecordOK) SetPayload(payload *models.DNSRecord) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateDNSRecordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateDNSRecordBadRequestCode is the HTTP code returned for type UpdateDNSRecordBadRequest
const UpdateDNSRecordBadRequestCode int = 400

/*
UpdateDNSRecordBadRequest Bad parameters were provided.

swagger:response updateDnsRecordBadRequest
*/
type UpdateDNSRecordBadRequest struct {
}

// NewUpdateDNSRecordBadRequest creates UpdateDNSRecordBadRequest with default headers values
func NewUpdateDNSRecordBadRequest() *UpdateDNSRecordBadRequest {

	return &UpdateDNSRecordBadRequest{}
}

// WriteResponse to the client
func (o *UpdateDNSRecordBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateDNSRecordNotFoundCode is the HTTP code returned for type UpdateDNSRecordNotFound
const UpdateDNSRecordNotFoundCode int = 404

/*
UpdateDNSRecordNotFound Invalid DNS record ID was provided.

swagger:response updateDnsRecordNotFound
*/
type UpdateDNSRecordNotFound struct {
}

// NewUpdateDNSRecordNotFound creates UpdateDNSRecordNotFound with default headers values
func NewUpdateDNSRecordNotFound() *UpdateDNSRecordNotFound {

	return &UpdateDNSRecordNotFound{}
}

// WriteResponse to the client
func (o *UpdateDNSRecordNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
