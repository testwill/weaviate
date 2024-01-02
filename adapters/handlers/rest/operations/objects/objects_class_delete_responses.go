//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsClassDeleteNoContentCode is the HTTP code returned for type ObjectsClassDeleteNoContent
const ObjectsClassDeleteNoContentCode int = 204

/*
ObjectsClassDeleteNoContent Successfully deleted.

swagger:response objectsClassDeleteNoContent
*/
type ObjectsClassDeleteNoContent struct {
}

// NewObjectsClassDeleteNoContent creates ObjectsClassDeleteNoContent with default headers values
func NewObjectsClassDeleteNoContent() *ObjectsClassDeleteNoContent {

	return &ObjectsClassDeleteNoContent{}
}

// WriteResponse to the client
func (o *ObjectsClassDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ObjectsClassDeleteBadRequestCode is the HTTP code returned for type ObjectsClassDeleteBadRequest
const ObjectsClassDeleteBadRequestCode int = 400

/*
ObjectsClassDeleteBadRequest Malformed request.

swagger:response objectsClassDeleteBadRequest
*/
type ObjectsClassDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassDeleteBadRequest creates ObjectsClassDeleteBadRequest with default headers values
func NewObjectsClassDeleteBadRequest() *ObjectsClassDeleteBadRequest {

	return &ObjectsClassDeleteBadRequest{}
}

// WithPayload adds the payload to the objects class delete bad request response
func (o *ObjectsClassDeleteBadRequest) WithPayload(payload *models.ErrorResponse) *ObjectsClassDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class delete bad request response
func (o *ObjectsClassDeleteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassDeleteUnauthorizedCode is the HTTP code returned for type ObjectsClassDeleteUnauthorized
const ObjectsClassDeleteUnauthorizedCode int = 401

/*
ObjectsClassDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassDeleteUnauthorized
*/
type ObjectsClassDeleteUnauthorized struct {
}

// NewObjectsClassDeleteUnauthorized creates ObjectsClassDeleteUnauthorized with default headers values
func NewObjectsClassDeleteUnauthorized() *ObjectsClassDeleteUnauthorized {

	return &ObjectsClassDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassDeleteForbiddenCode is the HTTP code returned for type ObjectsClassDeleteForbidden
const ObjectsClassDeleteForbiddenCode int = 403

/*
ObjectsClassDeleteForbidden Forbidden

swagger:response objectsClassDeleteForbidden
*/
type ObjectsClassDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassDeleteForbidden creates ObjectsClassDeleteForbidden with default headers values
func NewObjectsClassDeleteForbidden() *ObjectsClassDeleteForbidden {

	return &ObjectsClassDeleteForbidden{}
}

// WithPayload adds the payload to the objects class delete forbidden response
func (o *ObjectsClassDeleteForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class delete forbidden response
func (o *ObjectsClassDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassDeleteNotFoundCode is the HTTP code returned for type ObjectsClassDeleteNotFound
const ObjectsClassDeleteNotFoundCode int = 404

/*
ObjectsClassDeleteNotFound Successful query result but no resource was found.

swagger:response objectsClassDeleteNotFound
*/
type ObjectsClassDeleteNotFound struct {
}

// NewObjectsClassDeleteNotFound creates ObjectsClassDeleteNotFound with default headers values
func NewObjectsClassDeleteNotFound() *ObjectsClassDeleteNotFound {

	return &ObjectsClassDeleteNotFound{}
}

// WriteResponse to the client
func (o *ObjectsClassDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsClassDeleteUnprocessableEntityCode is the HTTP code returned for type ObjectsClassDeleteUnprocessableEntity
const ObjectsClassDeleteUnprocessableEntityCode int = 422

/*
ObjectsClassDeleteUnprocessableEntity Request is well-formed (i.e., syntactically correct), but erroneous.

swagger:response objectsClassDeleteUnprocessableEntity
*/
type ObjectsClassDeleteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassDeleteUnprocessableEntity creates ObjectsClassDeleteUnprocessableEntity with default headers values
func NewObjectsClassDeleteUnprocessableEntity() *ObjectsClassDeleteUnprocessableEntity {

	return &ObjectsClassDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the objects class delete unprocessable entity response
func (o *ObjectsClassDeleteUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsClassDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class delete unprocessable entity response
func (o *ObjectsClassDeleteUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassDeleteInternalServerErrorCode is the HTTP code returned for type ObjectsClassDeleteInternalServerError
const ObjectsClassDeleteInternalServerErrorCode int = 500

/*
ObjectsClassDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassDeleteInternalServerError
*/
type ObjectsClassDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassDeleteInternalServerError creates ObjectsClassDeleteInternalServerError with default headers values
func NewObjectsClassDeleteInternalServerError() *ObjectsClassDeleteInternalServerError {

	return &ObjectsClassDeleteInternalServerError{}
}

// WithPayload adds the payload to the objects class delete internal server error response
func (o *ObjectsClassDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class delete internal server error response
func (o *ObjectsClassDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
