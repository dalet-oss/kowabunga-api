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

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 507:
		result := NewCreateProjectInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/*
CreateProjectCreated describes a response with status code 201, with default header values.

Returns the newly created project object.
*/
type CreateProjectCreated struct {
	Payload *models.Project
}

// IsSuccess returns true when this create project created response has a 2xx status code
func (o *CreateProjectCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create project created response has a 3xx status code
func (o *CreateProjectCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project created response has a 4xx status code
func (o *CreateProjectCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project created response has a 5xx status code
func (o *CreateProjectCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create project created response a status code equal to that given
func (o *CreateProjectCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create project created response
func (o *CreateProjectCreated) Code() int {
	return 201
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[POST /project][%d] createProjectCreated  %+v", 201, o.Payload)
}

func (o *CreateProjectCreated) String() string {
	return fmt.Sprintf("[POST /project][%d] createProjectCreated  %+v", 201, o.Payload)
}

func (o *CreateProjectCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectBadRequest creates a CreateProjectBadRequest with default headers values
func NewCreateProjectBadRequest() *CreateProjectBadRequest {
	return &CreateProjectBadRequest{}
}

/*
CreateProjectBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type CreateProjectBadRequest struct {
}

// IsSuccess returns true when this create project bad request response has a 2xx status code
func (o *CreateProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project bad request response has a 3xx status code
func (o *CreateProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project bad request response has a 4xx status code
func (o *CreateProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create project bad request response has a 5xx status code
func (o *CreateProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create project bad request response a status code equal to that given
func (o *CreateProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create project bad request response
func (o *CreateProjectBadRequest) Code() int {
	return 400
}

func (o *CreateProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /project][%d] createProjectBadRequest ", 400)
}

func (o *CreateProjectBadRequest) String() string {
	return fmt.Sprintf("[POST /project][%d] createProjectBadRequest ", 400)
}

func (o *CreateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectConflict creates a CreateProjectConflict with default headers values
func NewCreateProjectConflict() *CreateProjectConflict {
	return &CreateProjectConflict{}
}

/*
CreateProjectConflict describes a response with status code 409, with default header values.

Project already exists.
*/
type CreateProjectConflict struct {
}

// IsSuccess returns true when this create project conflict response has a 2xx status code
func (o *CreateProjectConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project conflict response has a 3xx status code
func (o *CreateProjectConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project conflict response has a 4xx status code
func (o *CreateProjectConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create project conflict response has a 5xx status code
func (o *CreateProjectConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create project conflict response a status code equal to that given
func (o *CreateProjectConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create project conflict response
func (o *CreateProjectConflict) Code() int {
	return 409
}

func (o *CreateProjectConflict) Error() string {
	return fmt.Sprintf("[POST /project][%d] createProjectConflict ", 409)
}

func (o *CreateProjectConflict) String() string {
	return fmt.Sprintf("[POST /project][%d] createProjectConflict ", 409)
}

func (o *CreateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectInternalServerError creates a CreateProjectInternalServerError with default headers values
func NewCreateProjectInternalServerError() *CreateProjectInternalServerError {
	return &CreateProjectInternalServerError{}
}

/*
CreateProjectInternalServerError describes a response with status code 500, with default header values.

Unable to create the requested project.
*/
type CreateProjectInternalServerError struct {
}

// IsSuccess returns true when this create project internal server error response has a 2xx status code
func (o *CreateProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project internal server error response has a 3xx status code
func (o *CreateProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project internal server error response has a 4xx status code
func (o *CreateProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project internal server error response has a 5xx status code
func (o *CreateProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create project internal server error response a status code equal to that given
func (o *CreateProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create project internal server error response
func (o *CreateProjectInternalServerError) Code() int {
	return 500
}

func (o *CreateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /project][%d] createProjectInternalServerError ", 500)
}

func (o *CreateProjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /project][%d] createProjectInternalServerError ", 500)
}

func (o *CreateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectInsufficientStorage creates a CreateProjectInsufficientStorage with default headers values
func NewCreateProjectInsufficientStorage() *CreateProjectInsufficientStorage {
	return &CreateProjectInsufficientStorage{}
}

/*
CreateProjectInsufficientStorage describes a response with status code 507, with default header values.

The expected VPC subnet size cannot be assigned.
*/
type CreateProjectInsufficientStorage struct {
}

// IsSuccess returns true when this create project insufficient storage response has a 2xx status code
func (o *CreateProjectInsufficientStorage) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project insufficient storage response has a 3xx status code
func (o *CreateProjectInsufficientStorage) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project insufficient storage response has a 4xx status code
func (o *CreateProjectInsufficientStorage) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project insufficient storage response has a 5xx status code
func (o *CreateProjectInsufficientStorage) IsServerError() bool {
	return true
}

// IsCode returns true when this create project insufficient storage response a status code equal to that given
func (o *CreateProjectInsufficientStorage) IsCode(code int) bool {
	return code == 507
}

// Code gets the status code for the create project insufficient storage response
func (o *CreateProjectInsufficientStorage) Code() int {
	return 507
}

func (o *CreateProjectInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /project][%d] createProjectInsufficientStorage ", 507)
}

func (o *CreateProjectInsufficientStorage) String() string {
	return fmt.Sprintf("[POST /project][%d] createProjectInsufficientStorage ", 507)
}

func (o *CreateProjectInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
