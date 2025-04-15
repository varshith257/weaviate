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

// ListAllUsersOKCode is the HTTP code returned for type ListAllUsersOK
const ListAllUsersOKCode int = 200

/*
ListAllUsersOK Info about the users

swagger:response listAllUsersOK
*/
type ListAllUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.DBUserInfo `json:"body,omitempty"`
}

// NewListAllUsersOK creates ListAllUsersOK with default headers values
func NewListAllUsersOK() *ListAllUsersOK {

	return &ListAllUsersOK{}
}

// WithPayload adds the payload to the list all users o k response
func (o *ListAllUsersOK) WithPayload(payload []*models.DBUserInfo) *ListAllUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all users o k response
func (o *ListAllUsersOK) SetPayload(payload []*models.DBUserInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.DBUserInfo, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListAllUsersUnauthorizedCode is the HTTP code returned for type ListAllUsersUnauthorized
const ListAllUsersUnauthorizedCode int = 401

/*
ListAllUsersUnauthorized Unauthorized or invalid credentials.

swagger:response listAllUsersUnauthorized
*/
type ListAllUsersUnauthorized struct {
}

// NewListAllUsersUnauthorized creates ListAllUsersUnauthorized with default headers values
func NewListAllUsersUnauthorized() *ListAllUsersUnauthorized {

	return &ListAllUsersUnauthorized{}
}

// WriteResponse to the client
func (o *ListAllUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ListAllUsersForbiddenCode is the HTTP code returned for type ListAllUsersForbidden
const ListAllUsersForbiddenCode int = 403

/*
ListAllUsersForbidden Forbidden

swagger:response listAllUsersForbidden
*/
type ListAllUsersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListAllUsersForbidden creates ListAllUsersForbidden with default headers values
func NewListAllUsersForbidden() *ListAllUsersForbidden {

	return &ListAllUsersForbidden{}
}

// WithPayload adds the payload to the list all users forbidden response
func (o *ListAllUsersForbidden) WithPayload(payload *models.ErrorResponse) *ListAllUsersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all users forbidden response
func (o *ListAllUsersForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllUsersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAllUsersInternalServerErrorCode is the HTTP code returned for type ListAllUsersInternalServerError
const ListAllUsersInternalServerErrorCode int = 500

/*
ListAllUsersInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response listAllUsersInternalServerError
*/
type ListAllUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewListAllUsersInternalServerError creates ListAllUsersInternalServerError with default headers values
func NewListAllUsersInternalServerError() *ListAllUsersInternalServerError {

	return &ListAllUsersInternalServerError{}
}

// WithPayload adds the payload to the list all users internal server error response
func (o *ListAllUsersInternalServerError) WithPayload(payload *models.ErrorResponse) *ListAllUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all users internal server error response
func (o *ListAllUsersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
