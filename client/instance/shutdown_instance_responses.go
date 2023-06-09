// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ShutdownInstanceReader is a Reader for the ShutdownInstance structure.
type ShutdownInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShutdownInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewShutdownInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShutdownInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutdownInstanceOK creates a ShutdownInstanceOK with default headers values
func NewShutdownInstanceOK() *ShutdownInstanceOK {
	return &ShutdownInstanceOK{}
}

/*
ShutdownInstanceOK describes a response with status code 200, with default header values.

The virtual machine has been shut down successfully.
*/
type ShutdownInstanceOK struct {
}

// IsSuccess returns true when this shutdown instance o k response has a 2xx status code
func (o *ShutdownInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this shutdown instance o k response has a 3xx status code
func (o *ShutdownInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shutdown instance o k response has a 4xx status code
func (o *ShutdownInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this shutdown instance o k response has a 5xx status code
func (o *ShutdownInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this shutdown instance o k response a status code equal to that given
func (o *ShutdownInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the shutdown instance o k response
func (o *ShutdownInstanceOK) Code() int {
	return 200
}

func (o *ShutdownInstanceOK) Error() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/shutdown][%d] shutdownInstanceOK ", 200)
}

func (o *ShutdownInstanceOK) String() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/shutdown][%d] shutdownInstanceOK ", 200)
}

func (o *ShutdownInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShutdownInstanceNotFound creates a ShutdownInstanceNotFound with default headers values
func NewShutdownInstanceNotFound() *ShutdownInstanceNotFound {
	return &ShutdownInstanceNotFound{}
}

/*
ShutdownInstanceNotFound describes a response with status code 404, with default header values.

Invalid instance ID was provided.
*/
type ShutdownInstanceNotFound struct {
}

// IsSuccess returns true when this shutdown instance not found response has a 2xx status code
func (o *ShutdownInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this shutdown instance not found response has a 3xx status code
func (o *ShutdownInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shutdown instance not found response has a 4xx status code
func (o *ShutdownInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this shutdown instance not found response has a 5xx status code
func (o *ShutdownInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this shutdown instance not found response a status code equal to that given
func (o *ShutdownInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the shutdown instance not found response
func (o *ShutdownInstanceNotFound) Code() int {
	return 404
}

func (o *ShutdownInstanceNotFound) Error() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/shutdown][%d] shutdownInstanceNotFound ", 404)
}

func (o *ShutdownInstanceNotFound) String() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/shutdown][%d] shutdownInstanceNotFound ", 404)
}

func (o *ShutdownInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShutdownInstanceInternalServerError creates a ShutdownInstanceInternalServerError with default headers values
func NewShutdownInstanceInternalServerError() *ShutdownInstanceInternalServerError {
	return &ShutdownInstanceInternalServerError{}
}

/*
ShutdownInstanceInternalServerError describes a response with status code 500, with default header values.

An error occurred when trying to shut down the virtual machine.
*/
type ShutdownInstanceInternalServerError struct {
}

// IsSuccess returns true when this shutdown instance internal server error response has a 2xx status code
func (o *ShutdownInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this shutdown instance internal server error response has a 3xx status code
func (o *ShutdownInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this shutdown instance internal server error response has a 4xx status code
func (o *ShutdownInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this shutdown instance internal server error response has a 5xx status code
func (o *ShutdownInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this shutdown instance internal server error response a status code equal to that given
func (o *ShutdownInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the shutdown instance internal server error response
func (o *ShutdownInstanceInternalServerError) Code() int {
	return 500
}

func (o *ShutdownInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/shutdown][%d] shutdownInstanceInternalServerError ", 500)
}

func (o *ShutdownInstanceInternalServerError) String() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/shutdown][%d] shutdownInstanceInternalServerError ", 500)
}

func (o *ShutdownInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
