// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteInstanceReader is a Reader for the DeleteInstance structure.
type DeleteInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInstanceOK creates a DeleteInstanceOK with default headers values
func NewDeleteInstanceOK() *DeleteInstanceOK {
	return &DeleteInstanceOK{}
}

/*
DeleteInstanceOK describes a response with status code 200, with default header values.

The instance has been successfully removed.
*/
type DeleteInstanceOK struct {
}

// IsSuccess returns true when this delete instance o k response has a 2xx status code
func (o *DeleteInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete instance o k response has a 3xx status code
func (o *DeleteInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance o k response has a 4xx status code
func (o *DeleteInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete instance o k response has a 5xx status code
func (o *DeleteInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance o k response a status code equal to that given
func (o *DeleteInstanceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete instance o k response
func (o *DeleteInstanceOK) Code() int {
	return 200
}

func (o *DeleteInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceOK ", 200)
}

func (o *DeleteInstanceOK) String() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceOK ", 200)
}

func (o *DeleteInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceNotFound creates a DeleteInstanceNotFound with default headers values
func NewDeleteInstanceNotFound() *DeleteInstanceNotFound {
	return &DeleteInstanceNotFound{}
}

/*
DeleteInstanceNotFound describes a response with status code 404, with default header values.

Invalid instance ID was provided.
*/
type DeleteInstanceNotFound struct {
}

// IsSuccess returns true when this delete instance not found response has a 2xx status code
func (o *DeleteInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance not found response has a 3xx status code
func (o *DeleteInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance not found response has a 4xx status code
func (o *DeleteInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete instance not found response has a 5xx status code
func (o *DeleteInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance not found response a status code equal to that given
func (o *DeleteInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete instance not found response
func (o *DeleteInstanceNotFound) Code() int {
	return 404
}

func (o *DeleteInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceNotFound ", 404)
}

func (o *DeleteInstanceNotFound) String() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceNotFound ", 404)
}

func (o *DeleteInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceConflict creates a DeleteInstanceConflict with default headers values
func NewDeleteInstanceConflict() *DeleteInstanceConflict {
	return &DeleteInstanceConflict{}
}

/*
DeleteInstanceConflict describes a response with status code 409, with default header values.

The instance is not empty or still being referenced.
*/
type DeleteInstanceConflict struct {
}

// IsSuccess returns true when this delete instance conflict response has a 2xx status code
func (o *DeleteInstanceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance conflict response has a 3xx status code
func (o *DeleteInstanceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance conflict response has a 4xx status code
func (o *DeleteInstanceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete instance conflict response has a 5xx status code
func (o *DeleteInstanceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance conflict response a status code equal to that given
func (o *DeleteInstanceConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete instance conflict response
func (o *DeleteInstanceConflict) Code() int {
	return 409
}

func (o *DeleteInstanceConflict) Error() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceConflict ", 409)
}

func (o *DeleteInstanceConflict) String() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceConflict ", 409)
}

func (o *DeleteInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInstanceInternalServerError creates a DeleteInstanceInternalServerError with default headers values
func NewDeleteInstanceInternalServerError() *DeleteInstanceInternalServerError {
	return &DeleteInstanceInternalServerError{}
}

/*
DeleteInstanceInternalServerError describes a response with status code 500, with default header values.

Unable to delete instance.
*/
type DeleteInstanceInternalServerError struct {
}

// IsSuccess returns true when this delete instance internal server error response has a 2xx status code
func (o *DeleteInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance internal server error response has a 3xx status code
func (o *DeleteInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance internal server error response has a 4xx status code
func (o *DeleteInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete instance internal server error response has a 5xx status code
func (o *DeleteInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete instance internal server error response a status code equal to that given
func (o *DeleteInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete instance internal server error response
func (o *DeleteInstanceInternalServerError) Code() int {
	return 500
}

func (o *DeleteInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceInternalServerError ", 500)
}

func (o *DeleteInstanceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /instance/{instanceId}][%d] deleteInstanceInternalServerError ", 500)
}

func (o *DeleteInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}