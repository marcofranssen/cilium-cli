// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PostIpamReader is a Reader for the PostIpam structure.
type PostIpamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIpamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostIpamCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPostIpamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewPostIpamFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ipam] PostIpam", response, response.Code())
	}
}

// NewPostIpamCreated creates a PostIpamCreated with default headers values
func NewPostIpamCreated() *PostIpamCreated {
	return &PostIpamCreated{}
}

/*
PostIpamCreated describes a response with status code 201, with default header values.

Success
*/
type PostIpamCreated struct {
	Payload *models.IPAMResponse
}

// IsSuccess returns true when this post ipam created response has a 2xx status code
func (o *PostIpamCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post ipam created response has a 3xx status code
func (o *PostIpamCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post ipam created response has a 4xx status code
func (o *PostIpamCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post ipam created response has a 5xx status code
func (o *PostIpamCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post ipam created response a status code equal to that given
func (o *PostIpamCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post ipam created response
func (o *PostIpamCreated) Code() int {
	return 201
}

func (o *PostIpamCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam][%d] postIpamCreated %s", 201, payload)
}

func (o *PostIpamCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam][%d] postIpamCreated %s", 201, payload)
}

func (o *PostIpamCreated) GetPayload() *models.IPAMResponse {
	return o.Payload
}

func (o *PostIpamCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIpamForbidden creates a PostIpamForbidden with default headers values
func NewPostIpamForbidden() *PostIpamForbidden {
	return &PostIpamForbidden{}
}

/*
PostIpamForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostIpamForbidden struct {
}

// IsSuccess returns true when this post ipam forbidden response has a 2xx status code
func (o *PostIpamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post ipam forbidden response has a 3xx status code
func (o *PostIpamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post ipam forbidden response has a 4xx status code
func (o *PostIpamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post ipam forbidden response has a 5xx status code
func (o *PostIpamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post ipam forbidden response a status code equal to that given
func (o *PostIpamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post ipam forbidden response
func (o *PostIpamForbidden) Code() int {
	return 403
}

func (o *PostIpamForbidden) Error() string {
	return fmt.Sprintf("[POST /ipam][%d] postIpamForbidden", 403)
}

func (o *PostIpamForbidden) String() string {
	return fmt.Sprintf("[POST /ipam][%d] postIpamForbidden", 403)
}

func (o *PostIpamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIpamFailure creates a PostIpamFailure with default headers values
func NewPostIpamFailure() *PostIpamFailure {
	return &PostIpamFailure{}
}

/*
PostIpamFailure describes a response with status code 502, with default header values.

Allocation failure
*/
type PostIpamFailure struct {
	Payload models.Error
}

// IsSuccess returns true when this post ipam failure response has a 2xx status code
func (o *PostIpamFailure) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post ipam failure response has a 3xx status code
func (o *PostIpamFailure) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post ipam failure response has a 4xx status code
func (o *PostIpamFailure) IsClientError() bool {
	return false
}

// IsServerError returns true when this post ipam failure response has a 5xx status code
func (o *PostIpamFailure) IsServerError() bool {
	return true
}

// IsCode returns true when this post ipam failure response a status code equal to that given
func (o *PostIpamFailure) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the post ipam failure response
func (o *PostIpamFailure) Code() int {
	return 502
}

func (o *PostIpamFailure) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam][%d] postIpamFailure %s", 502, payload)
}

func (o *PostIpamFailure) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ipam][%d] postIpamFailure %s", 502, payload)
}

func (o *PostIpamFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *PostIpamFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
