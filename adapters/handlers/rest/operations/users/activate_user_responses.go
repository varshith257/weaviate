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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ActivateUserOKCode is the HTTP code returned for type ActivateUserOK
const ActivateUserOKCode int = 200

/*
ActivateUserOK User successfully activated

swagger:response activateUserOK
*/
type ActivateUserOK struct {
}

// NewActivateUserOK creates ActivateUserOK with default headers values
func NewActivateUserOK() *ActivateUserOK {

	return &ActivateUserOK{}
}

// WriteResponse to the client
func (o *ActivateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ActivateUserBadRequestCode is the HTTP code returned for type ActivateUserBadRequest
const ActivateUserBadRequestCode int = 400

/*
ActivateUserBadRequest Malformed request.

swagger:response activateUserBadRequest
*/
type ActivateUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActivateUserBadRequest creates ActivateUserBadRequest with default headers values
func NewActivateUserBadRequest() *ActivateUserBadRequest {

	return &ActivateUserBadRequest{}
}

// WithPayload adds the payload to the activate user bad request response
func (o *ActivateUserBadRequest) WithPayload(payload *models.ErrorResponse) *ActivateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activate user bad request response
func (o *ActivateUserBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActivateUserUnauthorizedCode is the HTTP code returned for type ActivateUserUnauthorized
const ActivateUserUnauthorizedCode int = 401

/*
ActivateUserUnauthorized Unauthorized or invalid credentials.

swagger:response activateUserUnauthorized
*/
type ActivateUserUnauthorized struct {
}

// NewActivateUserUnauthorized creates ActivateUserUnauthorized with default headers values
func NewActivateUserUnauthorized() *ActivateUserUnauthorized {

	return &ActivateUserUnauthorized{}
}

// WriteResponse to the client
func (o *ActivateUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ActivateUserForbiddenCode is the HTTP code returned for type ActivateUserForbidden
const ActivateUserForbiddenCode int = 403

/*
ActivateUserForbidden Forbidden

swagger:response activateUserForbidden
*/
type ActivateUserForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActivateUserForbidden creates ActivateUserForbidden with default headers values
func NewActivateUserForbidden() *ActivateUserForbidden {

	return &ActivateUserForbidden{}
}

// WithPayload adds the payload to the activate user forbidden response
func (o *ActivateUserForbidden) WithPayload(payload *models.ErrorResponse) *ActivateUserForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activate user forbidden response
func (o *ActivateUserForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivateUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActivateUserNotFoundCode is the HTTP code returned for type ActivateUserNotFound
const ActivateUserNotFoundCode int = 404

/*
ActivateUserNotFound user not found

swagger:response activateUserNotFound
*/
type ActivateUserNotFound struct {
}

// NewActivateUserNotFound creates ActivateUserNotFound with default headers values
func NewActivateUserNotFound() *ActivateUserNotFound {

	return &ActivateUserNotFound{}
}

// WriteResponse to the client
func (o *ActivateUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ActivateUserUnprocessableEntityCode is the HTTP code returned for type ActivateUserUnprocessableEntity
const ActivateUserUnprocessableEntityCode int = 422

/*
ActivateUserUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous.

swagger:response activateUserUnprocessableEntity
*/
type ActivateUserUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActivateUserUnprocessableEntity creates ActivateUserUnprocessableEntity with default headers values
func NewActivateUserUnprocessableEntity() *ActivateUserUnprocessableEntity {

	return &ActivateUserUnprocessableEntity{}
}

// WithPayload adds the payload to the activate user unprocessable entity response
func (o *ActivateUserUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ActivateUserUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activate user unprocessable entity response
func (o *ActivateUserUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivateUserUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ActivateUserInternalServerErrorCode is the HTTP code returned for type ActivateUserInternalServerError
const ActivateUserInternalServerErrorCode int = 500

/*
ActivateUserInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response activateUserInternalServerError
*/
type ActivateUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewActivateUserInternalServerError creates ActivateUserInternalServerError with default headers values
func NewActivateUserInternalServerError() *ActivateUserInternalServerError {

	return &ActivateUserInternalServerError{}
}

// WithPayload adds the payload to the activate user internal server error response
func (o *ActivateUserInternalServerError) WithPayload(payload *models.ErrorResponse) *ActivateUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activate user internal server error response
func (o *ActivateUserInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
