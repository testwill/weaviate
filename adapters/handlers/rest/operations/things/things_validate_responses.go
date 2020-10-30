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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ThingsValidateOKCode is the HTTP code returned for type ThingsValidateOK
const ThingsValidateOKCode int = 200

/*ThingsValidateOK Successfully validated.

swagger:response thingsValidateOK
*/
type ThingsValidateOK struct {
}

// NewThingsValidateOK creates ThingsValidateOK with default headers values
func NewThingsValidateOK() *ThingsValidateOK {
	return &ThingsValidateOK{}
}

// WriteResponse to the client
func (o *ThingsValidateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ThingsValidateUnauthorizedCode is the HTTP code returned for type ThingsValidateUnauthorized
const ThingsValidateUnauthorizedCode int = 401

/*ThingsValidateUnauthorized Unauthorized or invalid credentials.

swagger:response thingsValidateUnauthorized
*/
type ThingsValidateUnauthorized struct {
}

// NewThingsValidateUnauthorized creates ThingsValidateUnauthorized with default headers values
func NewThingsValidateUnauthorized() *ThingsValidateUnauthorized {
	return &ThingsValidateUnauthorized{}
}

// WriteResponse to the client
func (o *ThingsValidateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ThingsValidateForbiddenCode is the HTTP code returned for type ThingsValidateForbidden
const ThingsValidateForbiddenCode int = 403

/*ThingsValidateForbidden Forbidden

swagger:response thingsValidateForbidden
*/
type ThingsValidateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsValidateForbidden creates ThingsValidateForbidden with default headers values
func NewThingsValidateForbidden() *ThingsValidateForbidden {
	return &ThingsValidateForbidden{}
}

// WithPayload adds the payload to the things validate forbidden response
func (o *ThingsValidateForbidden) WithPayload(payload *models.ErrorResponse) *ThingsValidateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things validate forbidden response
func (o *ThingsValidateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsValidateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsValidateUnprocessableEntityCode is the HTTP code returned for type ThingsValidateUnprocessableEntity
const ThingsValidateUnprocessableEntityCode int = 422

/*ThingsValidateUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response thingsValidateUnprocessableEntity
*/
type ThingsValidateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsValidateUnprocessableEntity creates ThingsValidateUnprocessableEntity with default headers values
func NewThingsValidateUnprocessableEntity() *ThingsValidateUnprocessableEntity {
	return &ThingsValidateUnprocessableEntity{}
}

// WithPayload adds the payload to the things validate unprocessable entity response
func (o *ThingsValidateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ThingsValidateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things validate unprocessable entity response
func (o *ThingsValidateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsValidateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ThingsValidateInternalServerErrorCode is the HTTP code returned for type ThingsValidateInternalServerError
const ThingsValidateInternalServerErrorCode int = 500

/*ThingsValidateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response thingsValidateInternalServerError
*/
type ThingsValidateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewThingsValidateInternalServerError creates ThingsValidateInternalServerError with default headers values
func NewThingsValidateInternalServerError() *ThingsValidateInternalServerError {
	return &ThingsValidateInternalServerError{}
}

// WithPayload adds the payload to the things validate internal server error response
func (o *ThingsValidateInternalServerError) WithPayload(payload *models.ErrorResponse) *ThingsValidateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the things validate internal server error response
func (o *ThingsValidateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ThingsValidateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
