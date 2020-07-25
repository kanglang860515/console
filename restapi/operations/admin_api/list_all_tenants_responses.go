// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/mcs/models"
)

// ListAllTenantsOKCode is the HTTP code returned for type ListAllTenantsOK
const ListAllTenantsOKCode int = 200

/*ListAllTenantsOK A successful response.

swagger:response listAllTenantsOK
*/
type ListAllTenantsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListTenantsResponse `json:"body,omitempty"`
}

// NewListAllTenantsOK creates ListAllTenantsOK with default headers values
func NewListAllTenantsOK() *ListAllTenantsOK {

	return &ListAllTenantsOK{}
}

// WithPayload adds the payload to the list all tenants o k response
func (o *ListAllTenantsOK) WithPayload(payload *models.ListTenantsResponse) *ListAllTenantsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all tenants o k response
func (o *ListAllTenantsOK) SetPayload(payload *models.ListTenantsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllTenantsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListAllTenantsDefault Generic error response.

swagger:response listAllTenantsDefault
*/
type ListAllTenantsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListAllTenantsDefault creates ListAllTenantsDefault with default headers values
func NewListAllTenantsDefault(code int) *ListAllTenantsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListAllTenantsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list all tenants default response
func (o *ListAllTenantsDefault) WithStatusCode(code int) *ListAllTenantsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list all tenants default response
func (o *ListAllTenantsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list all tenants default response
func (o *ListAllTenantsDefault) WithPayload(payload *models.Error) *ListAllTenantsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all tenants default response
func (o *ListAllTenantsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllTenantsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}