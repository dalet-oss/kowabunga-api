// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dalet-oss/kowabunga-api/models"
)

// CreateProjectZoneVolumeCreatedCode is the HTTP code returned for type CreateProjectZoneVolumeCreated
const CreateProjectZoneVolumeCreatedCode int = 201

/*
CreateProjectZoneVolumeCreated Returns the newly created storage volume object.

swagger:response createProjectZoneVolumeCreated
*/
type CreateProjectZoneVolumeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Volume `json:"body,omitempty"`
}

// NewCreateProjectZoneVolumeCreated creates CreateProjectZoneVolumeCreated with default headers values
func NewCreateProjectZoneVolumeCreated() *CreateProjectZoneVolumeCreated {

	return &CreateProjectZoneVolumeCreated{}
}

// WithPayload adds the payload to the create project zone volume created response
func (o *CreateProjectZoneVolumeCreated) WithPayload(payload *models.Volume) *CreateProjectZoneVolumeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create project zone volume created response
func (o *CreateProjectZoneVolumeCreated) SetPayload(payload *models.Volume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProjectZoneVolumeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateProjectZoneVolumeBadRequestCode is the HTTP code returned for type CreateProjectZoneVolumeBadRequest
const CreateProjectZoneVolumeBadRequestCode int = 400

/*
CreateProjectZoneVolumeBadRequest Bad parameters were provided.

swagger:response createProjectZoneVolumeBadRequest
*/
type CreateProjectZoneVolumeBadRequest struct {
}

// NewCreateProjectZoneVolumeBadRequest creates CreateProjectZoneVolumeBadRequest with default headers values
func NewCreateProjectZoneVolumeBadRequest() *CreateProjectZoneVolumeBadRequest {

	return &CreateProjectZoneVolumeBadRequest{}
}

// WriteResponse to the client
func (o *CreateProjectZoneVolumeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateProjectZoneVolumeNotFoundCode is the HTTP code returned for type CreateProjectZoneVolumeNotFound
const CreateProjectZoneVolumeNotFoundCode int = 404

/*
CreateProjectZoneVolumeNotFound Invalid project or zone ID was provided.

swagger:response createProjectZoneVolumeNotFound
*/
type CreateProjectZoneVolumeNotFound struct {
}

// NewCreateProjectZoneVolumeNotFound creates CreateProjectZoneVolumeNotFound with default headers values
func NewCreateProjectZoneVolumeNotFound() *CreateProjectZoneVolumeNotFound {

	return &CreateProjectZoneVolumeNotFound{}
}

// WriteResponse to the client
func (o *CreateProjectZoneVolumeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// CreateProjectZoneVolumeConflictCode is the HTTP code returned for type CreateProjectZoneVolumeConflict
const CreateProjectZoneVolumeConflictCode int = 409

/*
CreateProjectZoneVolumeConflict Storage volume already exists.

swagger:response createProjectZoneVolumeConflict
*/
type CreateProjectZoneVolumeConflict struct {
}

// NewCreateProjectZoneVolumeConflict creates CreateProjectZoneVolumeConflict with default headers values
func NewCreateProjectZoneVolumeConflict() *CreateProjectZoneVolumeConflict {

	return &CreateProjectZoneVolumeConflict{}
}

// WriteResponse to the client
func (o *CreateProjectZoneVolumeConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// CreateProjectZoneVolumeInternalServerErrorCode is the HTTP code returned for type CreateProjectZoneVolumeInternalServerError
const CreateProjectZoneVolumeInternalServerErrorCode int = 500

/*
CreateProjectZoneVolumeInternalServerError Unable to create the storage volume.

swagger:response createProjectZoneVolumeInternalServerError
*/
type CreateProjectZoneVolumeInternalServerError struct {
}

// NewCreateProjectZoneVolumeInternalServerError creates CreateProjectZoneVolumeInternalServerError with default headers values
func NewCreateProjectZoneVolumeInternalServerError() *CreateProjectZoneVolumeInternalServerError {

	return &CreateProjectZoneVolumeInternalServerError{}
}

// WriteResponse to the client
func (o *CreateProjectZoneVolumeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

// CreateProjectZoneVolumeInsufficientStorageCode is the HTTP code returned for type CreateProjectZoneVolumeInsufficientStorage
const CreateProjectZoneVolumeInsufficientStorageCode int = 507

/*
CreateProjectZoneVolumeInsufficientStorage Requested volume characteristics are beyond project's quota in place.

swagger:response createProjectZoneVolumeInsufficientStorage
*/
type CreateProjectZoneVolumeInsufficientStorage struct {
}

// NewCreateProjectZoneVolumeInsufficientStorage creates CreateProjectZoneVolumeInsufficientStorage with default headers values
func NewCreateProjectZoneVolumeInsufficientStorage() *CreateProjectZoneVolumeInsufficientStorage {

	return &CreateProjectZoneVolumeInsufficientStorage{}
}

// WriteResponse to the client
func (o *CreateProjectZoneVolumeInsufficientStorage) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(507)
}
