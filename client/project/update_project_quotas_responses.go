// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateProjectQuotasReader is a Reader for the UpdateProjectQuotas structure.
type UpdateProjectQuotasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectQuotasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectQuotasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProjectQuotasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectQuotasOK creates a UpdateProjectQuotasOK with default headers values
func NewUpdateProjectQuotasOK() *UpdateProjectQuotasOK {
	return &UpdateProjectQuotasOK{}
}

/*
UpdateProjectQuotasOK describes a response with status code 200, with default header values.

Returns the updated project resources object.
*/
type UpdateProjectQuotasOK struct {
	Payload *models.ProjectResources
}

// IsSuccess returns true when this update project quotas o k response has a 2xx status code
func (o *UpdateProjectQuotasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project quotas o k response has a 3xx status code
func (o *UpdateProjectQuotasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project quotas o k response has a 4xx status code
func (o *UpdateProjectQuotasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project quotas o k response has a 5xx status code
func (o *UpdateProjectQuotasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project quotas o k response a status code equal to that given
func (o *UpdateProjectQuotasOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update project quotas o k response
func (o *UpdateProjectQuotasOK) Code() int {
	return 200
}

func (o *UpdateProjectQuotasOK) Error() string {
	return fmt.Sprintf("[PUT /project/{projectId}/quotas][%d] updateProjectQuotasOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectQuotasOK) String() string {
	return fmt.Sprintf("[PUT /project/{projectId}/quotas][%d] updateProjectQuotasOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectQuotasOK) GetPayload() *models.ProjectResources {
	return o.Payload
}

func (o *UpdateProjectQuotasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectResources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectQuotasNotFound creates a UpdateProjectQuotasNotFound with default headers values
func NewUpdateProjectQuotasNotFound() *UpdateProjectQuotasNotFound {
	return &UpdateProjectQuotasNotFound{}
}

/*
UpdateProjectQuotasNotFound describes a response with status code 404, with default header values.

Invalid project ID was provided.
*/
type UpdateProjectQuotasNotFound struct {
}

// IsSuccess returns true when this update project quotas not found response has a 2xx status code
func (o *UpdateProjectQuotasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project quotas not found response has a 3xx status code
func (o *UpdateProjectQuotasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project quotas not found response has a 4xx status code
func (o *UpdateProjectQuotasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project quotas not found response has a 5xx status code
func (o *UpdateProjectQuotasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project quotas not found response a status code equal to that given
func (o *UpdateProjectQuotasNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update project quotas not found response
func (o *UpdateProjectQuotasNotFound) Code() int {
	return 404
}

func (o *UpdateProjectQuotasNotFound) Error() string {
	return fmt.Sprintf("[PUT /project/{projectId}/quotas][%d] updateProjectQuotasNotFound ", 404)
}

func (o *UpdateProjectQuotasNotFound) String() string {
	return fmt.Sprintf("[PUT /project/{projectId}/quotas][%d] updateProjectQuotasNotFound ", 404)
}

func (o *UpdateProjectQuotasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}