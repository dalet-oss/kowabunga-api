// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/sdk/go/models"
)

// UpdateTemplateOKCode is the HTTP code returned for type UpdateTemplateOK
const UpdateTemplateOKCode int = 200

/*
UpdateTemplateOK Returns the updated volume template object.

swagger:response updateTemplateOK
*/
type UpdateTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Template `json:"body,omitempty"`
}

// NewUpdateTemplateOK creates UpdateTemplateOK with default headers values
func NewUpdateTemplateOK() *UpdateTemplateOK {

	return &UpdateTemplateOK{}
}

// WithPayload adds the payload to the update template o k response
func (o *UpdateTemplateOK) WithPayload(payload *models.Template) *UpdateTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update template o k response
func (o *UpdateTemplateOK) SetPayload(payload *models.Template) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateTemplateBadRequestCode is the HTTP code returned for type UpdateTemplateBadRequest
const UpdateTemplateBadRequestCode int = 400

/*
UpdateTemplateBadRequest Bad parameters were provided.

swagger:response updateTemplateBadRequest
*/
type UpdateTemplateBadRequest struct {
}

// NewUpdateTemplateBadRequest creates UpdateTemplateBadRequest with default headers values
func NewUpdateTemplateBadRequest() *UpdateTemplateBadRequest {

	return &UpdateTemplateBadRequest{}
}

// WriteResponse to the client
func (o *UpdateTemplateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateTemplateNotFoundCode is the HTTP code returned for type UpdateTemplateNotFound
const UpdateTemplateNotFoundCode int = 404

/*
UpdateTemplateNotFound Invalid volume template ID was provided.

swagger:response updateTemplateNotFound
*/
type UpdateTemplateNotFound struct {
}

// NewUpdateTemplateNotFound creates UpdateTemplateNotFound with default headers values
func NewUpdateTemplateNotFound() *UpdateTemplateNotFound {

	return &UpdateTemplateNotFound{}
}

// WriteResponse to the client
func (o *UpdateTemplateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}