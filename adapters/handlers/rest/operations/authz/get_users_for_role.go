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

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/entities/models"
)

// GetUsersForRoleHandlerFunc turns a function with the right signature into a get users for role handler
type GetUsersForRoleHandlerFunc func(GetUsersForRoleParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersForRoleHandlerFunc) Handle(params GetUsersForRoleParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetUsersForRoleHandler interface for that can handle valid get users for role params
type GetUsersForRoleHandler interface {
	Handle(GetUsersForRoleParams, *models.Principal) middleware.Responder
}

// NewGetUsersForRole creates a new http.Handler for the get users for role operation
func NewGetUsersForRole(ctx *middleware.Context, handler GetUsersForRoleHandler) *GetUsersForRole {
	return &GetUsersForRole{Context: ctx, Handler: handler}
}

/*
	GetUsersForRole swagger:route GET /authz/roles/{id}/users/{userType} authz getUsersForRole

get users or a keys assigned to role
*/
type GetUsersForRole struct {
	Context *middleware.Context
	Handler GetUsersForRoleHandler
}

func (o *GetUsersForRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	Params := NewGetUsersForRoleParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)
}
