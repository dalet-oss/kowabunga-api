// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/sdk/go/models"
)

// GetTemplateOKCode is the HTTP code returned for type GetTemplateOK
const GetTemplateOKCode int = 200

/*
GetTemplateOK Returns the volume template object.

swagger:response getTemplateOK
*/
type GetTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Template `json:"body,omitempty"`
}

// NewGetTemplateOK creates GetTemplateOK with default headers values
func NewGetTemplateOK() *GetTemplateOK {

	return &GetTemplateOK{}
}

// WithPayload adds the payload to the get template o k response
func (o *GetTemplateOK) WithPayload(payload *models.Template) *GetTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get template o k response
func (o *GetTemplateOK) SetPayload(payload *models.Template) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTemplateNotFoundCode is the HTTP code returned for type GetTemplateNotFound
const GetTemplateNotFoundCode int = 404

/*
GetTemplateNotFound Invalid volume template ID was provided.

swagger:response getTemplateNotFound
*/
type GetTemplateNotFound struct {
}

// NewGetTemplateNotFound creates GetTemplateNotFound with default headers values
func NewGetTemplateNotFound() *GetTemplateNotFound {

	return &GetTemplateNotFound{}
}

// WriteResponse to the client
func (o *GetTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}