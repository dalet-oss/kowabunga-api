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

// GetZoneNetGWsReader is a Reader for the GetZoneNetGWs structure.
type GetZoneNetGWsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZoneNetGWsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZoneNetGWsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetZoneNetGWsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZoneNetGWsOK creates a GetZoneNetGWsOK with default headers values
func NewGetZoneNetGWsOK() *GetZoneNetGWsOK {
	return &GetZoneNetGWsOK{}
}

/*
GetZoneNetGWsOK describes a response with status code 200, with default header values.

Returns an array of network gateway IDs.
*/
type GetZoneNetGWsOK struct {
	Payload []string
}

// IsSuccess returns true when this get zone net g ws o k response has a 2xx status code
func (o *GetZoneNetGWsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get zone net g ws o k response has a 3xx status code
func (o *GetZoneNetGWsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get zone net g ws o k response has a 4xx status code
func (o *GetZoneNetGWsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get zone net g ws o k response has a 5xx status code
func (o *GetZoneNetGWsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get zone net g ws o k response a status code equal to that given
func (o *GetZoneNetGWsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get zone net g ws o k response
func (o *GetZoneNetGWsOK) Code() int {
	return 200
}

func (o *GetZoneNetGWsOK) Error() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/netgws][%d] getZoneNetGWsOK  %+v", 200, o.Payload)
}

func (o *GetZoneNetGWsOK) String() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/netgws][%d] getZoneNetGWsOK  %+v", 200, o.Payload)
}

func (o *GetZoneNetGWsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetZoneNetGWsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZoneNetGWsNotFound creates a GetZoneNetGWsNotFound with default headers values
func NewGetZoneNetGWsNotFound() *GetZoneNetGWsNotFound {
	return &GetZoneNetGWsNotFound{}
}

/*
GetZoneNetGWsNotFound describes a response with status code 404, with default header values.

Invalid zone ID was provided.
*/
type GetZoneNetGWsNotFound struct {
}

// IsSuccess returns true when this get zone net g ws not found response has a 2xx status code
func (o *GetZoneNetGWsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get zone net g ws not found response has a 3xx status code
func (o *GetZoneNetGWsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get zone net g ws not found response has a 4xx status code
func (o *GetZoneNetGWsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get zone net g ws not found response has a 5xx status code
func (o *GetZoneNetGWsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get zone net g ws not found response a status code equal to that given
func (o *GetZoneNetGWsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get zone net g ws not found response
func (o *GetZoneNetGWsNotFound) Code() int {
	return 404
}

func (o *GetZoneNetGWsNotFound) Error() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/netgws][%d] getZoneNetGWsNotFound ", 404)
}

func (o *GetZoneNetGWsNotFound) String() string {
	return fmt.Sprintf("[GET /zone/{zoneId}/netgws][%d] getZoneNetGWsNotFound ", 404)
}

func (o *GetZoneNetGWsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
