// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"net/http"
)

type CreateOrderRequest struct {
	CreateOrderBackendRequest *shared.CreateOrderBackendRequest `request:"mediaType=application/json"`
	XAPIVersion               *string                           `header:"style=simple,explode=false,name=x-api-version"`
	XClientID                 string                            `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret             string                            `header:"style=simple,explode=false,name=x-client-secret"`
}

func (o *CreateOrderRequest) GetCreateOrderBackendRequest() *shared.CreateOrderBackendRequest {
	if o == nil {
		return nil
	}
	return o.CreateOrderBackendRequest
}

func (o *CreateOrderRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *CreateOrderRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *CreateOrderRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type CreateOrderResponse struct {
	// API related Errors
	APIError *shared.APIError
	// Authentication Error
	AuthenticationError *shared.AuthenticationError
	ContentType         string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	// OK
	OrdersEntity *shared.OrdersEntity
	// Rate Limit Error
	RateLimitError *shared.RateLimitError
	StatusCode     int
	RawResponse    *http.Response
}

func (o *CreateOrderResponse) GetAPIError() *shared.APIError {
	if o == nil {
		return nil
	}
	return o.APIError
}

func (o *CreateOrderResponse) GetAuthenticationError() *shared.AuthenticationError {
	if o == nil {
		return nil
	}
	return o.AuthenticationError
}

func (o *CreateOrderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateOrderResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *CreateOrderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateOrderResponse) GetOrdersEntity() *shared.OrdersEntity {
	if o == nil {
		return nil
	}
	return o.OrdersEntity
}

func (o *CreateOrderResponse) GetRateLimitError() *shared.RateLimitError {
	if o == nil {
		return nil
	}
	return o.RateLimitError
}

func (o *CreateOrderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateOrderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}