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

// CreateProjectZoneVolumeReader is a Reader for the CreateProjectZoneVolume structure.
type CreateProjectZoneVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectZoneVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectZoneVolumeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectZoneVolumeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProjectZoneVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProjectZoneVolumeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectZoneVolumeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 507:
		result := NewCreateProjectZoneVolumeInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectZoneVolumeCreated creates a CreateProjectZoneVolumeCreated with default headers values
func NewCreateProjectZoneVolumeCreated() *CreateProjectZoneVolumeCreated {
	return &CreateProjectZoneVolumeCreated{}
}

/*
CreateProjectZoneVolumeCreated describes a response with status code 201, with default header values.

Returns the newly created storage volume object.
*/
type CreateProjectZoneVolumeCreated struct {
	Payload *models.Volume
}

// IsSuccess returns true when this create project zone volume created response has a 2xx status code
func (o *CreateProjectZoneVolumeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create project zone volume created response has a 3xx status code
func (o *CreateProjectZoneVolumeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project zone volume created response has a 4xx status code
func (o *CreateProjectZoneVolumeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project zone volume created response has a 5xx status code
func (o *CreateProjectZoneVolumeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create project zone volume created response a status code equal to that given
func (o *CreateProjectZoneVolumeCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create project zone volume created response
func (o *CreateProjectZoneVolumeCreated) Code() int {
	return 201
}

func (o *CreateProjectZoneVolumeCreated) Error() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeCreated  %+v", 201, o.Payload)
}

func (o *CreateProjectZoneVolumeCreated) String() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeCreated  %+v", 201, o.Payload)
}

func (o *CreateProjectZoneVolumeCreated) GetPayload() *models.Volume {
	return o.Payload
}

func (o *CreateProjectZoneVolumeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectZoneVolumeBadRequest creates a CreateProjectZoneVolumeBadRequest with default headers values
func NewCreateProjectZoneVolumeBadRequest() *CreateProjectZoneVolumeBadRequest {
	return &CreateProjectZoneVolumeBadRequest{}
}

/*
CreateProjectZoneVolumeBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type CreateProjectZoneVolumeBadRequest struct {
}

// IsSuccess returns true when this create project zone volume bad request response has a 2xx status code
func (o *CreateProjectZoneVolumeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project zone volume bad request response has a 3xx status code
func (o *CreateProjectZoneVolumeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project zone volume bad request response has a 4xx status code
func (o *CreateProjectZoneVolumeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create project zone volume bad request response has a 5xx status code
func (o *CreateProjectZoneVolumeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create project zone volume bad request response a status code equal to that given
func (o *CreateProjectZoneVolumeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create project zone volume bad request response
func (o *CreateProjectZoneVolumeBadRequest) Code() int {
	return 400
}

func (o *CreateProjectZoneVolumeBadRequest) Error() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeBadRequest ", 400)
}

func (o *CreateProjectZoneVolumeBadRequest) String() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeBadRequest ", 400)
}

func (o *CreateProjectZoneVolumeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectZoneVolumeNotFound creates a CreateProjectZoneVolumeNotFound with default headers values
func NewCreateProjectZoneVolumeNotFound() *CreateProjectZoneVolumeNotFound {
	return &CreateProjectZoneVolumeNotFound{}
}

/*
CreateProjectZoneVolumeNotFound describes a response with status code 404, with default header values.

Invalid project or zone ID was provided.
*/
type CreateProjectZoneVolumeNotFound struct {
}

// IsSuccess returns true when this create project zone volume not found response has a 2xx status code
func (o *CreateProjectZoneVolumeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project zone volume not found response has a 3xx status code
func (o *CreateProjectZoneVolumeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project zone volume not found response has a 4xx status code
func (o *CreateProjectZoneVolumeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create project zone volume not found response has a 5xx status code
func (o *CreateProjectZoneVolumeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create project zone volume not found response a status code equal to that given
func (o *CreateProjectZoneVolumeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create project zone volume not found response
func (o *CreateProjectZoneVolumeNotFound) Code() int {
	return 404
}

func (o *CreateProjectZoneVolumeNotFound) Error() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeNotFound ", 404)
}

func (o *CreateProjectZoneVolumeNotFound) String() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeNotFound ", 404)
}

func (o *CreateProjectZoneVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectZoneVolumeConflict creates a CreateProjectZoneVolumeConflict with default headers values
func NewCreateProjectZoneVolumeConflict() *CreateProjectZoneVolumeConflict {
	return &CreateProjectZoneVolumeConflict{}
}

/*
CreateProjectZoneVolumeConflict describes a response with status code 409, with default header values.

Storage volume already exists.
*/
type CreateProjectZoneVolumeConflict struct {
}

// IsSuccess returns true when this create project zone volume conflict response has a 2xx status code
func (o *CreateProjectZoneVolumeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project zone volume conflict response has a 3xx status code
func (o *CreateProjectZoneVolumeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project zone volume conflict response has a 4xx status code
func (o *CreateProjectZoneVolumeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create project zone volume conflict response has a 5xx status code
func (o *CreateProjectZoneVolumeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create project zone volume conflict response a status code equal to that given
func (o *CreateProjectZoneVolumeConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create project zone volume conflict response
func (o *CreateProjectZoneVolumeConflict) Code() int {
	return 409
}

func (o *CreateProjectZoneVolumeConflict) Error() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeConflict ", 409)
}

func (o *CreateProjectZoneVolumeConflict) String() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeConflict ", 409)
}

func (o *CreateProjectZoneVolumeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectZoneVolumeInternalServerError creates a CreateProjectZoneVolumeInternalServerError with default headers values
func NewCreateProjectZoneVolumeInternalServerError() *CreateProjectZoneVolumeInternalServerError {
	return &CreateProjectZoneVolumeInternalServerError{}
}

/*
CreateProjectZoneVolumeInternalServerError describes a response with status code 500, with default header values.

Unable to create the storage volume.
*/
type CreateProjectZoneVolumeInternalServerError struct {
}

// IsSuccess returns true when this create project zone volume internal server error response has a 2xx status code
func (o *CreateProjectZoneVolumeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project zone volume internal server error response has a 3xx status code
func (o *CreateProjectZoneVolumeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project zone volume internal server error response has a 4xx status code
func (o *CreateProjectZoneVolumeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project zone volume internal server error response has a 5xx status code
func (o *CreateProjectZoneVolumeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create project zone volume internal server error response a status code equal to that given
func (o *CreateProjectZoneVolumeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create project zone volume internal server error response
func (o *CreateProjectZoneVolumeInternalServerError) Code() int {
	return 500
}

func (o *CreateProjectZoneVolumeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeInternalServerError ", 500)
}

func (o *CreateProjectZoneVolumeInternalServerError) String() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeInternalServerError ", 500)
}

func (o *CreateProjectZoneVolumeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProjectZoneVolumeInsufficientStorage creates a CreateProjectZoneVolumeInsufficientStorage with default headers values
func NewCreateProjectZoneVolumeInsufficientStorage() *CreateProjectZoneVolumeInsufficientStorage {
	return &CreateProjectZoneVolumeInsufficientStorage{}
}

/*
CreateProjectZoneVolumeInsufficientStorage describes a response with status code 507, with default header values.

Requested volume characteristics are beyond project's quota in place.
*/
type CreateProjectZoneVolumeInsufficientStorage struct {
}

// IsSuccess returns true when this create project zone volume insufficient storage response has a 2xx status code
func (o *CreateProjectZoneVolumeInsufficientStorage) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create project zone volume insufficient storage response has a 3xx status code
func (o *CreateProjectZoneVolumeInsufficientStorage) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create project zone volume insufficient storage response has a 4xx status code
func (o *CreateProjectZoneVolumeInsufficientStorage) IsClientError() bool {
	return false
}

// IsServerError returns true when this create project zone volume insufficient storage response has a 5xx status code
func (o *CreateProjectZoneVolumeInsufficientStorage) IsServerError() bool {
	return true
}

// IsCode returns true when this create project zone volume insufficient storage response a status code equal to that given
func (o *CreateProjectZoneVolumeInsufficientStorage) IsCode(code int) bool {
	return code == 507
}

// Code gets the status code for the create project zone volume insufficient storage response
func (o *CreateProjectZoneVolumeInsufficientStorage) Code() int {
	return 507
}

func (o *CreateProjectZoneVolumeInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeInsufficientStorage ", 507)
}

func (o *CreateProjectZoneVolumeInsufficientStorage) String() string {
	return fmt.Sprintf("[POST /project/{projectId}/zone/{zoneId}/volume][%d] createProjectZoneVolumeInsufficientStorage ", 507)
}

func (o *CreateProjectZoneVolumeInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}