// Code generated by go-swagger; DO NOT EDIT.

package subnet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSubnetReader is a Reader for the DeleteSubnet structure.
type DeleteSubnetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubnetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSubnetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSubnetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteSubnetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSubnetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSubnetOK creates a DeleteSubnetOK with default headers values
func NewDeleteSubnetOK() *DeleteSubnetOK {
	return &DeleteSubnetOK{}
}

/*
DeleteSubnetOK describes a response with status code 200, with default header values.

The subnet has been successfully removed.
*/
type DeleteSubnetOK struct {
}

// IsSuccess returns true when this delete subnet o k response has a 2xx status code
func (o *DeleteSubnetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete subnet o k response has a 3xx status code
func (o *DeleteSubnetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subnet o k response has a 4xx status code
func (o *DeleteSubnetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subnet o k response has a 5xx status code
func (o *DeleteSubnetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subnet o k response a status code equal to that given
func (o *DeleteSubnetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete subnet o k response
func (o *DeleteSubnetOK) Code() int {
	return 200
}

func (o *DeleteSubnetOK) Error() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetOK ", 200)
}

func (o *DeleteSubnetOK) String() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetOK ", 200)
}

func (o *DeleteSubnetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubnetNotFound creates a DeleteSubnetNotFound with default headers values
func NewDeleteSubnetNotFound() *DeleteSubnetNotFound {
	return &DeleteSubnetNotFound{}
}

/*
DeleteSubnetNotFound describes a response with status code 404, with default header values.

Invalid subnet ID was provided.
*/
type DeleteSubnetNotFound struct {
}

// IsSuccess returns true when this delete subnet not found response has a 2xx status code
func (o *DeleteSubnetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subnet not found response has a 3xx status code
func (o *DeleteSubnetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subnet not found response has a 4xx status code
func (o *DeleteSubnetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete subnet not found response has a 5xx status code
func (o *DeleteSubnetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subnet not found response a status code equal to that given
func (o *DeleteSubnetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete subnet not found response
func (o *DeleteSubnetNotFound) Code() int {
	return 404
}

func (o *DeleteSubnetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetNotFound ", 404)
}

func (o *DeleteSubnetNotFound) String() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetNotFound ", 404)
}

func (o *DeleteSubnetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubnetConflict creates a DeleteSubnetConflict with default headers values
func NewDeleteSubnetConflict() *DeleteSubnetConflict {
	return &DeleteSubnetConflict{}
}

/*
DeleteSubnetConflict describes a response with status code 409, with default header values.

The subnet is not empty or still being referenced.
*/
type DeleteSubnetConflict struct {
}

// IsSuccess returns true when this delete subnet conflict response has a 2xx status code
func (o *DeleteSubnetConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subnet conflict response has a 3xx status code
func (o *DeleteSubnetConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subnet conflict response has a 4xx status code
func (o *DeleteSubnetConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete subnet conflict response has a 5xx status code
func (o *DeleteSubnetConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subnet conflict response a status code equal to that given
func (o *DeleteSubnetConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete subnet conflict response
func (o *DeleteSubnetConflict) Code() int {
	return 409
}

func (o *DeleteSubnetConflict) Error() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetConflict ", 409)
}

func (o *DeleteSubnetConflict) String() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetConflict ", 409)
}

func (o *DeleteSubnetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubnetInternalServerError creates a DeleteSubnetInternalServerError with default headers values
func NewDeleteSubnetInternalServerError() *DeleteSubnetInternalServerError {
	return &DeleteSubnetInternalServerError{}
}

/*
DeleteSubnetInternalServerError describes a response with status code 500, with default header values.

Unable to delete subnet.
*/
type DeleteSubnetInternalServerError struct {
}

// IsSuccess returns true when this delete subnet internal server error response has a 2xx status code
func (o *DeleteSubnetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subnet internal server error response has a 3xx status code
func (o *DeleteSubnetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subnet internal server error response has a 4xx status code
func (o *DeleteSubnetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subnet internal server error response has a 5xx status code
func (o *DeleteSubnetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete subnet internal server error response a status code equal to that given
func (o *DeleteSubnetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete subnet internal server error response
func (o *DeleteSubnetInternalServerError) Code() int {
	return 500
}

func (o *DeleteSubnetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetInternalServerError ", 500)
}

func (o *DeleteSubnetInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /subnet/{subnetId}][%d] deleteSubnetInternalServerError ", 500)
}

func (o *DeleteSubnetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}