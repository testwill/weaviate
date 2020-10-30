//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsDeleteOKCode is the HTTP code returned for type SchemaActionsDeleteOK
const SchemaActionsDeleteOKCode int = 200

/*SchemaActionsDeleteOK Removed the Action class from the schema.

swagger:response schemaActionsDeleteOK
*/
type SchemaActionsDeleteOK struct {
}

// NewSchemaActionsDeleteOK creates SchemaActionsDeleteOK with default headers values
func NewSchemaActionsDeleteOK() *SchemaActionsDeleteOK {
	return &SchemaActionsDeleteOK{}
}

// WriteResponse to the client
func (o *SchemaActionsDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SchemaActionsDeleteBadRequestCode is the HTTP code returned for type SchemaActionsDeleteBadRequest
const SchemaActionsDeleteBadRequestCode int = 400

/*SchemaActionsDeleteBadRequest Could not delete the Action class.

swagger:response schemaActionsDeleteBadRequest
*/
type SchemaActionsDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsDeleteBadRequest creates SchemaActionsDeleteBadRequest with default headers values
func NewSchemaActionsDeleteBadRequest() *SchemaActionsDeleteBadRequest {
	return &SchemaActionsDeleteBadRequest{}
}

// WithPayload adds the payload to the schema actions delete bad request response
func (o *SchemaActionsDeleteBadRequest) WithPayload(payload *models.ErrorResponse) *SchemaActionsDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions delete bad request response
func (o *SchemaActionsDeleteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsDeleteUnauthorizedCode is the HTTP code returned for type SchemaActionsDeleteUnauthorized
const SchemaActionsDeleteUnauthorizedCode int = 401

/*SchemaActionsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response schemaActionsDeleteUnauthorized
*/
type SchemaActionsDeleteUnauthorized struct {
}

// NewSchemaActionsDeleteUnauthorized creates SchemaActionsDeleteUnauthorized with default headers values
func NewSchemaActionsDeleteUnauthorized() *SchemaActionsDeleteUnauthorized {
	return &SchemaActionsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *SchemaActionsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SchemaActionsDeleteForbiddenCode is the HTTP code returned for type SchemaActionsDeleteForbidden
const SchemaActionsDeleteForbiddenCode int = 403

/*SchemaActionsDeleteForbidden Forbidden

swagger:response schemaActionsDeleteForbidden
*/
type SchemaActionsDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsDeleteForbidden creates SchemaActionsDeleteForbidden with default headers values
func NewSchemaActionsDeleteForbidden() *SchemaActionsDeleteForbidden {
	return &SchemaActionsDeleteForbidden{}
}

// WithPayload adds the payload to the schema actions delete forbidden response
func (o *SchemaActionsDeleteForbidden) WithPayload(payload *models.ErrorResponse) *SchemaActionsDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions delete forbidden response
func (o *SchemaActionsDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SchemaActionsDeleteInternalServerErrorCode is the HTTP code returned for type SchemaActionsDeleteInternalServerError
const SchemaActionsDeleteInternalServerErrorCode int = 500

/*SchemaActionsDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response schemaActionsDeleteInternalServerError
*/
type SchemaActionsDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSchemaActionsDeleteInternalServerError creates SchemaActionsDeleteInternalServerError with default headers values
func NewSchemaActionsDeleteInternalServerError() *SchemaActionsDeleteInternalServerError {
	return &SchemaActionsDeleteInternalServerError{}
}

// WithPayload adds the payload to the schema actions delete internal server error response
func (o *SchemaActionsDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *SchemaActionsDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the schema actions delete internal server error response
func (o *SchemaActionsDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SchemaActionsDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
