// Code generated by go-swagger; DO NOT EDIT.

package nfs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dalet-oss/kowabunga-api/models"
)

// UpdateNfsStorageReader is a Reader for the UpdateNfsStorage structure.
type UpdateNfsStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNfsStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNfsStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNfsStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNfsStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /nfs/{nfsId}] UpdateNfsStorage", response, response.Code())
	}
}

// NewUpdateNfsStorageOK creates a UpdateNfsStorageOK with default headers values
func NewUpdateNfsStorageOK() *UpdateNfsStorageOK {
	return &UpdateNfsStorageOK{}
}

/*
UpdateNfsStorageOK describes a response with status code 200, with default header values.

Returns the updated NFS storage object.
*/
type UpdateNfsStorageOK struct {
	Payload *models.StorageNFS
}

// IsSuccess returns true when this update nfs storage o k response has a 2xx status code
func (o *UpdateNfsStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update nfs storage o k response has a 3xx status code
func (o *UpdateNfsStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update nfs storage o k response has a 4xx status code
func (o *UpdateNfsStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update nfs storage o k response has a 5xx status code
func (o *UpdateNfsStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update nfs storage o k response a status code equal to that given
func (o *UpdateNfsStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update nfs storage o k response
func (o *UpdateNfsStorageOK) Code() int {
	return 200
}

func (o *UpdateNfsStorageOK) Error() string {
	return fmt.Sprintf("[PUT /nfs/{nfsId}][%d] updateNfsStorageOK  %+v", 200, o.Payload)
}

func (o *UpdateNfsStorageOK) String() string {
	return fmt.Sprintf("[PUT /nfs/{nfsId}][%d] updateNfsStorageOK  %+v", 200, o.Payload)
}

func (o *UpdateNfsStorageOK) GetPayload() *models.StorageNFS {
	return o.Payload
}

func (o *UpdateNfsStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageNFS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNfsStorageBadRequest creates a UpdateNfsStorageBadRequest with default headers values
func NewUpdateNfsStorageBadRequest() *UpdateNfsStorageBadRequest {
	return &UpdateNfsStorageBadRequest{}
}

/*
UpdateNfsStorageBadRequest describes a response with status code 400, with default header values.

Bad parameters were provided.
*/
type UpdateNfsStorageBadRequest struct {
}

// IsSuccess returns true when this update nfs storage bad request response has a 2xx status code
func (o *UpdateNfsStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update nfs storage bad request response has a 3xx status code
func (o *UpdateNfsStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update nfs storage bad request response has a 4xx status code
func (o *UpdateNfsStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update nfs storage bad request response has a 5xx status code
func (o *UpdateNfsStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update nfs storage bad request response a status code equal to that given
func (o *UpdateNfsStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update nfs storage bad request response
func (o *UpdateNfsStorageBadRequest) Code() int {
	return 400
}

func (o *UpdateNfsStorageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /nfs/{nfsId}][%d] updateNfsStorageBadRequest ", 400)
}

func (o *UpdateNfsStorageBadRequest) String() string {
	return fmt.Sprintf("[PUT /nfs/{nfsId}][%d] updateNfsStorageBadRequest ", 400)
}

func (o *UpdateNfsStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNfsStorageNotFound creates a UpdateNfsStorageNotFound with default headers values
func NewUpdateNfsStorageNotFound() *UpdateNfsStorageNotFound {
	return &UpdateNfsStorageNotFound{}
}

/*
UpdateNfsStorageNotFound describes a response with status code 404, with default header values.

Invalid NFS storage ID was provided.
*/
type UpdateNfsStorageNotFound struct {
}

// IsSuccess returns true when this update nfs storage not found response has a 2xx status code
func (o *UpdateNfsStorageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update nfs storage not found response has a 3xx status code
func (o *UpdateNfsStorageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update nfs storage not found response has a 4xx status code
func (o *UpdateNfsStorageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update nfs storage not found response has a 5xx status code
func (o *UpdateNfsStorageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update nfs storage not found response a status code equal to that given
func (o *UpdateNfsStorageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update nfs storage not found response
func (o *UpdateNfsStorageNotFound) Code() int {
	return 404
}

func (o *UpdateNfsStorageNotFound) Error() string {
	return fmt.Sprintf("[PUT /nfs/{nfsId}][%d] updateNfsStorageNotFound ", 404)
}

func (o *UpdateNfsStorageNotFound) String() string {
	return fmt.Sprintf("[PUT /nfs/{nfsId}][%d] updateNfsStorageNotFound ", 404)
}

func (o *UpdateNfsStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}