// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateTemplateReader is a Reader for the UpdateTemplate structure.
type UpdateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTemplateOK creates a UpdateTemplateOK with default headers values
func NewUpdateTemplateOK() *UpdateTemplateOK {
	return &UpdateTemplateOK{}
}

/*
UpdateTemplateOK describes a response with status code 200, with default header values.

Returns the updated volume template object.
*/
type UpdateTemplateOK struct {
	Payload *models.Template
}

// IsSuccess returns true when this update template o k response has a 2xx status code
func (o *UpdateTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update template o k response has a 3xx status code
func (o *UpdateTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update template o k response has a 4xx status code
func (o *UpdateTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update template o k response has a 5xx status code
func (o *UpdateTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update template o k response a status code equal to that given
func (o *UpdateTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update template o k response
func (o *UpdateTemplateOK) Code() int {
	return 200
}

func (o *UpdateTemplateOK) Error() string {
	return fmt.Sprintf("[PUT /template/{templateId}][%d] updateTemplateOK  %+v", 200, o.Payload)
}

func (o *UpdateTemplateOK) String() string {
	return fmt.Sprintf("[PUT /template/{templateId}][%d] updateTemplateOK  %+v", 200, o.Payload)
}

func (o *UpdateTemplateOK) GetPayload() *models.Template {
	return o.Payload
}

func (o *UpdateTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Template)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTemplateBadRequest creates a UpdateTemplateBadRequest with default headers values
func NewUpdateTemplateBadRequest() *UpdateTemplateBadRequest {
	return &UpdateTemplateBadRequest{}
}

/*
UpdateTemplateBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type UpdateTemplateBadRequest struct {
}

// IsSuccess returns true when this update template bad request response has a 2xx status code
func (o *UpdateTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update template bad request response has a 3xx status code
func (o *UpdateTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update template bad request response has a 4xx status code
func (o *UpdateTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update template bad request response has a 5xx status code
func (o *UpdateTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update template bad request response a status code equal to that given
func (o *UpdateTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update template bad request response
func (o *UpdateTemplateBadRequest) Code() int {
	return 400
}

func (o *UpdateTemplateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /template/{templateId}][%d] updateTemplateBadRequest ", 400)
}

func (o *UpdateTemplateBadRequest) String() string {
	return fmt.Sprintf("[PUT /template/{templateId}][%d] updateTemplateBadRequest ", 400)
}

func (o *UpdateTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTemplateNotFound creates a UpdateTemplateNotFound with default headers values
func NewUpdateTemplateNotFound() *UpdateTemplateNotFound {
	return &UpdateTemplateNotFound{}
}

/*
UpdateTemplateNotFound describes a response with status code 404, with default header values.

Invalid volume template ID was provided.
*/
type UpdateTemplateNotFound struct {
}

// IsSuccess returns true when this update template not found response has a 2xx status code
func (o *UpdateTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update template not found response has a 3xx status code
func (o *UpdateTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update template not found response has a 4xx status code
func (o *UpdateTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update template not found response has a 5xx status code
func (o *UpdateTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update template not found response a status code equal to that given
func (o *UpdateTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update template not found response
func (o *UpdateTemplateNotFound) Code() int {
	return 404
}

func (o *UpdateTemplateNotFound) Error() string {
	return fmt.Sprintf("[PUT /template/{templateId}][%d] updateTemplateNotFound ", 404)
}

func (o *UpdateTemplateNotFound) String() string {
	return fmt.Sprintf("[PUT /template/{templateId}][%d] updateTemplateNotFound ", 404)
}

func (o *UpdateTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}