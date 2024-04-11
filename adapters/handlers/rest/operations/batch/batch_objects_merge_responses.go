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

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// BatchObjectsMergeOKCode is the HTTP code returned for type BatchObjectsMergeOK
const BatchObjectsMergeOKCode int = 200

/*
BatchObjectsMergeOK Request succeeded, see response body to get detailed information about each batched item.

swagger:response batchObjectsMergeOK
*/
type BatchObjectsMergeOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ObjectsGetResponse `json:"body,omitempty"`
}

// NewBatchObjectsMergeOK creates BatchObjectsMergeOK with default headers values
func NewBatchObjectsMergeOK() *BatchObjectsMergeOK {

	return &BatchObjectsMergeOK{}
}

// WithPayload adds the payload to the batch objects merge o k response
func (o *BatchObjectsMergeOK) WithPayload(payload []*models.ObjectsGetResponse) *BatchObjectsMergeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects merge o k response
func (o *BatchObjectsMergeOK) SetPayload(payload []*models.ObjectsGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsMergeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.ObjectsGetResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// BatchObjectsMergeBadRequestCode is the HTTP code returned for type BatchObjectsMergeBadRequest
const BatchObjectsMergeBadRequestCode int = 400

/*
BatchObjectsMergeBadRequest Malformed request.

swagger:response batchObjectsMergeBadRequest
*/
type BatchObjectsMergeBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsMergeBadRequest creates BatchObjectsMergeBadRequest with default headers values
func NewBatchObjectsMergeBadRequest() *BatchObjectsMergeBadRequest {

	return &BatchObjectsMergeBadRequest{}
}

// WithPayload adds the payload to the batch objects merge bad request response
func (o *BatchObjectsMergeBadRequest) WithPayload(payload *models.ErrorResponse) *BatchObjectsMergeBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects merge bad request response
func (o *BatchObjectsMergeBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsMergeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsMergeUnauthorizedCode is the HTTP code returned for type BatchObjectsMergeUnauthorized
const BatchObjectsMergeUnauthorizedCode int = 401

/*
BatchObjectsMergeUnauthorized Unauthorized or invalid credentials.

swagger:response batchObjectsMergeUnauthorized
*/
type BatchObjectsMergeUnauthorized struct {
}

// NewBatchObjectsMergeUnauthorized creates BatchObjectsMergeUnauthorized with default headers values
func NewBatchObjectsMergeUnauthorized() *BatchObjectsMergeUnauthorized {

	return &BatchObjectsMergeUnauthorized{}
}

// WriteResponse to the client
func (o *BatchObjectsMergeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BatchObjectsMergeForbiddenCode is the HTTP code returned for type BatchObjectsMergeForbidden
const BatchObjectsMergeForbiddenCode int = 403

/*
BatchObjectsMergeForbidden Forbidden

swagger:response batchObjectsMergeForbidden
*/
type BatchObjectsMergeForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsMergeForbidden creates BatchObjectsMergeForbidden with default headers values
func NewBatchObjectsMergeForbidden() *BatchObjectsMergeForbidden {

	return &BatchObjectsMergeForbidden{}
}

// WithPayload adds the payload to the batch objects merge forbidden response
func (o *BatchObjectsMergeForbidden) WithPayload(payload *models.ErrorResponse) *BatchObjectsMergeForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects merge forbidden response
func (o *BatchObjectsMergeForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsMergeForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsMergeUnprocessableEntityCode is the HTTP code returned for type BatchObjectsMergeUnprocessableEntity
const BatchObjectsMergeUnprocessableEntityCode int = 422

/*
BatchObjectsMergeUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response batchObjectsMergeUnprocessableEntity
*/
type BatchObjectsMergeUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsMergeUnprocessableEntity creates BatchObjectsMergeUnprocessableEntity with default headers values
func NewBatchObjectsMergeUnprocessableEntity() *BatchObjectsMergeUnprocessableEntity {

	return &BatchObjectsMergeUnprocessableEntity{}
}

// WithPayload adds the payload to the batch objects merge unprocessable entity response
func (o *BatchObjectsMergeUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *BatchObjectsMergeUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects merge unprocessable entity response
func (o *BatchObjectsMergeUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsMergeUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsMergeInternalServerErrorCode is the HTTP code returned for type BatchObjectsMergeInternalServerError
const BatchObjectsMergeInternalServerErrorCode int = 500

/*
BatchObjectsMergeInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response batchObjectsMergeInternalServerError
*/
type BatchObjectsMergeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsMergeInternalServerError creates BatchObjectsMergeInternalServerError with default headers values
func NewBatchObjectsMergeInternalServerError() *BatchObjectsMergeInternalServerError {

	return &BatchObjectsMergeInternalServerError{}
}

// WithPayload adds the payload to the batch objects merge internal server error response
func (o *BatchObjectsMergeInternalServerError) WithPayload(payload *models.ErrorResponse) *BatchObjectsMergeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects merge internal server error response
func (o *BatchObjectsMergeInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsMergeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
