// Code generated by go-swagger; DO NOT EDIT.

package region

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// CreateRegionReader is a Reader for the CreateRegion structure.
type CreateRegionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRegionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRegionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRegionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRegionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRegionCreated creates a CreateRegionCreated with default headers values
func NewCreateRegionCreated() *CreateRegionCreated {
	return &CreateRegionCreated{}
}

/*
CreateRegionCreated describes a response with status code 201, with default header values.

Returns the newly created region object.
*/
type CreateRegionCreated struct {
	Payload *models.Region
}

// IsSuccess returns true when this create region created response has a 2xx status code
func (o *CreateRegionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create region created response has a 3xx status code
func (o *CreateRegionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create region created response has a 4xx status code
func (o *CreateRegionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create region created response has a 5xx status code
func (o *CreateRegionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create region created response a status code equal to that given
func (o *CreateRegionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create region created response
func (o *CreateRegionCreated) Code() int {
	return 201
}

func (o *CreateRegionCreated) Error() string {
	return fmt.Sprintf("[POST /region][%d] createRegionCreated  %+v", 201, o.Payload)
}

func (o *CreateRegionCreated) String() string {
	return fmt.Sprintf("[POST /region][%d] createRegionCreated  %+v", 201, o.Payload)
}

func (o *CreateRegionCreated) GetPayload() *models.Region {
	return o.Payload
}

func (o *CreateRegionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegionBadRequest creates a CreateRegionBadRequest with default headers values
func NewCreateRegionBadRequest() *CreateRegionBadRequest {
	return &CreateRegionBadRequest{}
}

/*
CreateRegionBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type CreateRegionBadRequest struct {
}

// IsSuccess returns true when this create region bad request response has a 2xx status code
func (o *CreateRegionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create region bad request response has a 3xx status code
func (o *CreateRegionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create region bad request response has a 4xx status code
func (o *CreateRegionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create region bad request response has a 5xx status code
func (o *CreateRegionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create region bad request response a status code equal to that given
func (o *CreateRegionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create region bad request response
func (o *CreateRegionBadRequest) Code() int {
	return 400
}

func (o *CreateRegionBadRequest) Error() string {
	return fmt.Sprintf("[POST /region][%d] createRegionBadRequest ", 400)
}

func (o *CreateRegionBadRequest) String() string {
	return fmt.Sprintf("[POST /region][%d] createRegionBadRequest ", 400)
}

func (o *CreateRegionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRegionConflict creates a CreateRegionConflict with default headers values
func NewCreateRegionConflict() *CreateRegionConflict {
	return &CreateRegionConflict{}
}

/*
CreateRegionConflict describes a response with status code 409, with default header values.

Region already exists.
*/
type CreateRegionConflict struct {
}

// IsSuccess returns true when this create region conflict response has a 2xx status code
func (o *CreateRegionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create region conflict response has a 3xx status code
func (o *CreateRegionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create region conflict response has a 4xx status code
func (o *CreateRegionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create region conflict response has a 5xx status code
func (o *CreateRegionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create region conflict response a status code equal to that given
func (o *CreateRegionConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create region conflict response
func (o *CreateRegionConflict) Code() int {
	return 409
}

func (o *CreateRegionConflict) Error() string {
	return fmt.Sprintf("[POST /region][%d] createRegionConflict ", 409)
}

func (o *CreateRegionConflict) String() string {
	return fmt.Sprintf("[POST /region][%d] createRegionConflict ", 409)
}

func (o *CreateRegionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRegionInternalServerError creates a CreateRegionInternalServerError with default headers values
func NewCreateRegionInternalServerError() *CreateRegionInternalServerError {
	return &CreateRegionInternalServerError{}
}

/*
CreateRegionInternalServerError describes a response with status code 500, with default header values.

Unable to create the region.
*/
type CreateRegionInternalServerError struct {
}

// IsSuccess returns true when this create region internal server error response has a 2xx status code
func (o *CreateRegionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create region internal server error response has a 3xx status code
func (o *CreateRegionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create region internal server error response has a 4xx status code
func (o *CreateRegionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create region internal server error response has a 5xx status code
func (o *CreateRegionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create region internal server error response a status code equal to that given
func (o *CreateRegionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create region internal server error response
func (o *CreateRegionInternalServerError) Code() int {
	return 500
}

func (o *CreateRegionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /region][%d] createRegionInternalServerError ", 500)
}

func (o *CreateRegionInternalServerError) String() string {
	return fmt.Sprintf("[POST /region][%d] createRegionInternalServerError ", 500)
}

func (o *CreateRegionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
