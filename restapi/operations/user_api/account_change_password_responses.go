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

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// AccountChangePasswordCreatedCode is the HTTP code returned for type AccountChangePasswordCreated
const AccountChangePasswordCreatedCode int = 201

/*AccountChangePasswordCreated A successful login.

swagger:response accountChangePasswordCreated
*/
type AccountChangePasswordCreated struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewAccountChangePasswordCreated creates AccountChangePasswordCreated with default headers values
func NewAccountChangePasswordCreated() *AccountChangePasswordCreated {

	return &AccountChangePasswordCreated{}
}

// WithPayload adds the payload to the account change password created response
func (o *AccountChangePasswordCreated) WithPayload(payload *models.LoginResponse) *AccountChangePasswordCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the account change password created response
func (o *AccountChangePasswordCreated) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AccountChangePasswordCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AccountChangePasswordDefault Generic error response.

swagger:response accountChangePasswordDefault
*/
type AccountChangePasswordDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAccountChangePasswordDefault creates AccountChangePasswordDefault with default headers values
func NewAccountChangePasswordDefault(code int) *AccountChangePasswordDefault {
	if code <= 0 {
		code = 500
	}

	return &AccountChangePasswordDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the account change password default response
func (o *AccountChangePasswordDefault) WithStatusCode(code int) *AccountChangePasswordDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the account change password default response
func (o *AccountChangePasswordDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the account change password default response
func (o *AccountChangePasswordDefault) WithPayload(payload *models.Error) *AccountChangePasswordDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the account change password default response
func (o *AccountChangePasswordDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AccountChangePasswordDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
