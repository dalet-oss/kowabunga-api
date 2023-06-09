// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopInstanceReader is a Reader for the StopInstance structure.
type StopInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewStopInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopInstanceOK creates a StopInstanceOK with default headers values
func NewStopInstanceOK() *StopInstanceOK {
	return &StopInstanceOK{}
}

/*
StopInstanceOK describes a response with status code 200, with default header values.

The virtual machine has been stopped successfully.
*/
type StopInstanceOK struct {
}

// IsSuccess returns true when this stop instance o k response has a 2xx status code
func (o *StopInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop instance o k response has a 3xx status code
func (o *StopInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop instance o k response has a 4xx status code
func (o *StopInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop instance o k response has a 5xx status code
func (o *StopInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop instance o k response a status code equal to that given
func (o *StopInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop instance o k response
func (o *StopInstanceOK) Code() int {
	return 200
}

func (o *StopInstanceOK) Error() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/stop][%d] stopInstanceOK ", 200)
}

func (o *StopInstanceOK) String() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/stop][%d] stopInstanceOK ", 200)
}

func (o *StopInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopInstanceNotFound creates a StopInstanceNotFound with default headers values
func NewStopInstanceNotFound() *StopInstanceNotFound {
	return &StopInstanceNotFound{}
}

/*
StopInstanceNotFound describes a response with status code 404, with default header values.

Invalid instance ID was provided.
*/
type StopInstanceNotFound struct {
}

// IsSuccess returns true when this stop instance not found response has a 2xx status code
func (o *StopInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop instance not found response has a 3xx status code
func (o *StopInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop instance not found response has a 4xx status code
func (o *StopInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop instance not found response has a 5xx status code
func (o *StopInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop instance not found response a status code equal to that given
func (o *StopInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop instance not found response
func (o *StopInstanceNotFound) Code() int {
	return 404
}

func (o *StopInstanceNotFound) Error() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/stop][%d] stopInstanceNotFound ", 404)
}

func (o *StopInstanceNotFound) String() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/stop][%d] stopInstanceNotFound ", 404)
}

func (o *StopInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopInstanceInternalServerError creates a StopInstanceInternalServerError with default headers values
func NewStopInstanceInternalServerError() *StopInstanceInternalServerError {
	return &StopInstanceInternalServerError{}
}

/*
StopInstanceInternalServerError describes a response with status code 500, with default header values.

An error occurred when trying to stop the virtual machine.
*/
type StopInstanceInternalServerError struct {
}

// IsSuccess returns true when this stop instance internal server error response has a 2xx status code
func (o *StopInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop instance internal server error response has a 3xx status code
func (o *StopInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop instance internal server error response has a 4xx status code
func (o *StopInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop instance internal server error response has a 5xx status code
func (o *StopInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop instance internal server error response a status code equal to that given
func (o *StopInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stop instance internal server error response
func (o *StopInstanceInternalServerError) Code() int {
	return 500
}

func (o *StopInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/stop][%d] stopInstanceInternalServerError ", 500)
}

func (o *StopInstanceInternalServerError) String() string {
	return fmt.Sprintf("[POST /instance/{instanceId}/stop][%d] stopInstanceInternalServerError ", 500)
}

func (o *StopInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
