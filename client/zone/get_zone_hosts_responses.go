// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetZoneHostsReader is a Reader for the GetZoneHosts structure.
type GetZoneHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZoneHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZoneHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetZoneHostsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZoneHostsOK creates a GetZoneHostsOK with default headers values
func NewGetZoneHostsOK() *GetZoneHostsOK {
	return &GetZoneHostsOK{}
}

/*
GetZoneHostsOK describes a response with status code 200, with default header values.

Returns an array of host IDs.
*/
type GetZoneHostsOK struct {
	Payload []string
}

// IsSuccess returns true when this get zone hosts o k response has a 2xx status code
func (o *GetZoneHostsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get zone hosts o k response has a 3xx status code
func (o *GetZoneHostsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get zone hosts o k response has a 4xx status code
func (o *GetZoneHostsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get zone hosts o k response has a 5xx status code
func (o *GetZoneHostsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get zone hosts o k response a status code equal to that given
func (o *GetZoneHostsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get zone hosts o k response
func (o *GetZoneHostsOK) Code() int {
	return 200
}

func (o *GetZoneHostsOK) Error() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/hosts][%d] getZoneHostsOK  %+v", 200, o.Payload)
}

func (o *GetZoneHostsOK) String() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/hosts][%d] getZoneHostsOK  %+v", 200, o.Payload)
}

func (o *GetZoneHostsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetZoneHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZoneHostsNotFound creates a GetZoneHostsNotFound with default headers values
func NewGetZoneHostsNotFound() *GetZoneHostsNotFound {
	return &GetZoneHostsNotFound{}
}

/*
GetZoneHostsNotFound describes a response with status code 404, with default header values.

Invalid zone ID was provided.
*/
type GetZoneHostsNotFound struct {
}

// IsSuccess returns true when this get zone hosts not found response has a 2xx status code
func (o *GetZoneHostsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get zone hosts not found response has a 3xx status code
func (o *GetZoneHostsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get zone hosts not found response has a 4xx status code
func (o *GetZoneHostsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get zone hosts not found response has a 5xx status code
func (o *GetZoneHostsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get zone hosts not found response a status code equal to that given
func (o *GetZoneHostsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get zone hosts not found response
func (o *GetZoneHostsNotFound) Code() int {
	return 404
}

func (o *GetZoneHostsNotFound) Error() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/hosts][%d] getZoneHostsNotFound ", 404)
}

func (o *GetZoneHostsNotFound) String() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/hosts][%d] getZoneHostsNotFound ", 404)
}

func (o *GetZoneHostsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
