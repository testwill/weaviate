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

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// GraphqlBatchHandlerFunc turns a function with the right signature into a graphql batch handler
type GraphqlBatchHandlerFunc func(GraphqlBatchParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GraphqlBatchHandlerFunc) Handle(params GraphqlBatchParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GraphqlBatchHandler interface for that can handle valid graphql batch params
type GraphqlBatchHandler interface {
	Handle(GraphqlBatchParams, *models.Principal) middleware.Responder
}

// NewGraphqlBatch creates a new http.Handler for the graphql batch operation
func NewGraphqlBatch(ctx *middleware.Context, handler GraphqlBatchHandler) *GraphqlBatch {
	return &GraphqlBatch{Context: ctx, Handler: handler}
}

/*GraphqlBatch swagger:route POST /graphql/batch graphql graphqlBatch

Get a response based on GraphQL.

Perform a batched GraphQL query

*/
type GraphqlBatch struct {
	Context *middleware.Context
	Handler GraphqlBatchHandler
}

func (o *GraphqlBatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	Params := NewGraphqlBatchParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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
