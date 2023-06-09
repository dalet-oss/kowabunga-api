// Code generated by go-swagger; DO NOT EDIT.

package adapter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAdapterReader is a Reader for the DeleteAdapter structure.
type DeleteAdapterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdapterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAdapterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAdapterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAdapterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAdapterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAdapterOK creates a DeleteAdapterOK with default headers values
func NewDeleteAdapterOK() *DeleteAdapterOK {
	return &DeleteAdapterOK{}
}

/*
DeleteAdapterOK describes a response with status code 200, with default header values.

The network adapter has been successfully removed.
*/
type DeleteAdapterOK struct {
}

// IsSuccess returns true when this delete adapter o k response has a 2xx status code
func (o *DeleteAdapterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete adapter o k response has a 3xx status code
func (o *DeleteAdapterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete adapter o k response has a 4xx status code
func (o *DeleteAdapterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete adapter o k response has a 5xx status code
func (o *DeleteAdapterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete adapter o k response a status code equal to that given
func (o *DeleteAdapterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete adapter o k response
func (o *DeleteAdapterOK) Code() int {
	return 200
}

func (o *DeleteAdapterOK) Error() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterOK ", 200)
}

func (o *DeleteAdapterOK) String() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterOK ", 200)
}

func (o *DeleteAdapterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAdapterNotFound creates a DeleteAdapterNotFound with default headers values
func NewDeleteAdapterNotFound() *DeleteAdapterNotFound {
	return &DeleteAdapterNotFound{}
}

/*
DeleteAdapterNotFound describes a response with status code 404, with default header values.

Invalid network adapter ID was provided.
*/
type DeleteAdapterNotFound struct {
}

// IsSuccess returns true when this delete adapter not found response has a 2xx status code
func (o *DeleteAdapterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete adapter not found response has a 3xx status code
func (o *DeleteAdapterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete adapter not found response has a 4xx status code
func (o *DeleteAdapterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete adapter not found response has a 5xx status code
func (o *DeleteAdapterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete adapter not found response a status code equal to that given
func (o *DeleteAdapterNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete adapter not found response
func (o *DeleteAdapterNotFound) Code() int {
	return 404
}

func (o *DeleteAdapterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterNotFound ", 404)
}

func (o *DeleteAdapterNotFound) String() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterNotFound ", 404)
}

func (o *DeleteAdapterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAdapterConflict creates a DeleteAdapterConflict with default headers values
func NewDeleteAdapterConflict() *DeleteAdapterConflict {
	return &DeleteAdapterConflict{}
}

/*
DeleteAdapterConflict describes a response with status code 409, with default header values.

The network adapter is not empty or still being referenced.
*/
type DeleteAdapterConflict struct {
}

// IsSuccess returns true when this delete adapter conflict response has a 2xx status code
func (o *DeleteAdapterConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete adapter conflict response has a 3xx status code
func (o *DeleteAdapterConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete adapter conflict response has a 4xx status code
func (o *DeleteAdapterConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete adapter conflict response has a 5xx status code
func (o *DeleteAdapterConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete adapter conflict response a status code equal to that given
func (o *DeleteAdapterConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete adapter conflict response
func (o *DeleteAdapterConflict) Code() int {
	return 409
}

func (o *DeleteAdapterConflict) Error() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterConflict ", 409)
}

func (o *DeleteAdapterConflict) String() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterConflict ", 409)
}

func (o *DeleteAdapterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAdapterInternalServerError creates a DeleteAdapterInternalServerError with default headers values
func NewDeleteAdapterInternalServerError() *DeleteAdapterInternalServerError {
	return &DeleteAdapterInternalServerError{}
}

/*
DeleteAdapterInternalServerError describes a response with status code 500, with default header values.

Unable to delete network adapter.
*/
type DeleteAdapterInternalServerError struct {
}

// IsSuccess returns true when this delete adapter internal server error response has a 2xx status code
func (o *DeleteAdapterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete adapter internal server error response has a 3xx status code
func (o *DeleteAdapterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete adapter internal server error response has a 4xx status code
func (o *DeleteAdapterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete adapter internal server error response has a 5xx status code
func (o *DeleteAdapterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete adapter internal server error response a status code equal to that given
func (o *DeleteAdapterInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete adapter internal server error response
func (o *DeleteAdapterInternalServerError) Code() int {
	return 500
}

func (o *DeleteAdapterInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterInternalServerError ", 500)
}

func (o *DeleteAdapterInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /adapter/{adapterId}][%d] deleteAdapterInternalServerError ", 500)
}

func (o *DeleteAdapterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
