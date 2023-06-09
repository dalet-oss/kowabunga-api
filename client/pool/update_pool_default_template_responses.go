// Code generated by go-swagger; DO NOT EDIT.

package pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdatePoolDefaultTemplateReader is a Reader for the UpdatePoolDefaultTemplate structure.
type UpdatePoolDefaultTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePoolDefaultTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePoolDefaultTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdatePoolDefaultTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePoolDefaultTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdatePoolDefaultTemplateOK creates a UpdatePoolDefaultTemplateOK with default headers values
func NewUpdatePoolDefaultTemplateOK() *UpdatePoolDefaultTemplateOK {
	return &UpdatePoolDefaultTemplateOK{}
}

/*
UpdatePoolDefaultTemplateOK describes a response with status code 200, with default header values.

Returns success.
*/
type UpdatePoolDefaultTemplateOK struct {
}

// IsSuccess returns true when this update pool default template o k response has a 2xx status code
func (o *UpdatePoolDefaultTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update pool default template o k response has a 3xx status code
func (o *UpdatePoolDefaultTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update pool default template o k response has a 4xx status code
func (o *UpdatePoolDefaultTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update pool default template o k response has a 5xx status code
func (o *UpdatePoolDefaultTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update pool default template o k response a status code equal to that given
func (o *UpdatePoolDefaultTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update pool default template o k response
func (o *UpdatePoolDefaultTemplateOK) Code() int {
	return 200
}

func (o *UpdatePoolDefaultTemplateOK) Error() string {
	return fmt.Sprintf("[PUT /pool/{poolId}/template/{templateId}/default][%d] updatePoolDefaultTemplateOK ", 200)
}

func (o *UpdatePoolDefaultTemplateOK) String() string {
	return fmt.Sprintf("[PUT /pool/{poolId}/template/{templateId}/default][%d] updatePoolDefaultTemplateOK ", 200)
}

func (o *UpdatePoolDefaultTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePoolDefaultTemplateNotFound creates a UpdatePoolDefaultTemplateNotFound with default headers values
func NewUpdatePoolDefaultTemplateNotFound() *UpdatePoolDefaultTemplateNotFound {
	return &UpdatePoolDefaultTemplateNotFound{}
}

/*
UpdatePoolDefaultTemplateNotFound describes a response with status code 404, with default header values.

Invalid storage pool ID or volume template ID was provided.
*/
type UpdatePoolDefaultTemplateNotFound struct {
}

// IsSuccess returns true when this update pool default template not found response has a 2xx status code
func (o *UpdatePoolDefaultTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update pool default template not found response has a 3xx status code
func (o *UpdatePoolDefaultTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update pool default template not found response has a 4xx status code
func (o *UpdatePoolDefaultTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update pool default template not found response has a 5xx status code
func (o *UpdatePoolDefaultTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update pool default template not found response a status code equal to that given
func (o *UpdatePoolDefaultTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update pool default template not found response
func (o *UpdatePoolDefaultTemplateNotFound) Code() int {
	return 404
}

func (o *UpdatePoolDefaultTemplateNotFound) Error() string {
	return fmt.Sprintf("[PUT /pool/{poolId}/template/{templateId}/default][%d] updatePoolDefaultTemplateNotFound ", 404)
}

func (o *UpdatePoolDefaultTemplateNotFound) String() string {
	return fmt.Sprintf("[PUT /pool/{poolId}/template/{templateId}/default][%d] updatePoolDefaultTemplateNotFound ", 404)
}

func (o *UpdatePoolDefaultTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePoolDefaultTemplateInternalServerError creates a UpdatePoolDefaultTemplateInternalServerError with default headers values
func NewUpdatePoolDefaultTemplateInternalServerError() *UpdatePoolDefaultTemplateInternalServerError {
	return &UpdatePoolDefaultTemplateInternalServerError{}
}

/*
UpdatePoolDefaultTemplateInternalServerError describes a response with status code 500, with default header values.

Unable to assign the requested volume template as storage pool default.
*/
type UpdatePoolDefaultTemplateInternalServerError struct {
}

// IsSuccess returns true when this update pool default template internal server error response has a 2xx status code
func (o *UpdatePoolDefaultTemplateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update pool default template internal server error response has a 3xx status code
func (o *UpdatePoolDefaultTemplateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update pool default template internal server error response has a 4xx status code
func (o *UpdatePoolDefaultTemplateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update pool default template internal server error response has a 5xx status code
func (o *UpdatePoolDefaultTemplateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update pool default template internal server error response a status code equal to that given
func (o *UpdatePoolDefaultTemplateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update pool default template internal server error response
func (o *UpdatePoolDefaultTemplateInternalServerError) Code() int {
	return 500
}

func (o *UpdatePoolDefaultTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pool/{poolId}/template/{templateId}/default][%d] updatePoolDefaultTemplateInternalServerError ", 500)
}

func (o *UpdatePoolDefaultTemplateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /pool/{poolId}/template/{templateId}/default][%d] updatePoolDefaultTemplateInternalServerError ", 500)
}

func (o *UpdatePoolDefaultTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
