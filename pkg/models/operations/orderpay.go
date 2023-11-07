// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"net/http"
)

type OrderPayRequest struct {
	OrderPayRequest *shared.OrderPayRequest `request:"mediaType=application/json"`
	XAPIVersion     string                  `header:"style=simple,explode=false,name=x-api-version"`
}

func (o *OrderPayRequest) GetOrderPayRequest() *shared.OrderPayRequest {
	if o == nil {
		return nil
	}
	return o.OrderPayRequest
}

func (o *OrderPayRequest) GetXAPIVersion() string {
	if o == nil {
		return ""
	}
	return o.XAPIVersion
}

type OrderPayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	// OK
	OrderPayResponse *shared.OrderPayResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *OrderPayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *OrderPayResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *OrderPayResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *OrderPayResponse) GetOrderPayResponse() *shared.OrderPayResponse {
	if o == nil {
		return nil
	}
	return o.OrderPayResponse
}

func (o *OrderPayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *OrderPayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
