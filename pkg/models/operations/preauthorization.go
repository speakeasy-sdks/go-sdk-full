// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/utils"
	"net/http"
)

type PreauthorizationRequest struct {
	AuthorizationRequest *shared.AuthorizationRequest `request:"mediaType=application/json"`
	OrderID              string                       `pathParam:"style=simple,explode=false,name=order_id"`
	XAPIVersion          *string                      `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID            string                       `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret        string                       `header:"style=simple,explode=false,name=x-client-secret"`
}

func (p PreauthorizationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PreauthorizationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PreauthorizationRequest) GetAuthorizationRequest() *shared.AuthorizationRequest {
	if o == nil {
		return nil
	}
	return o.AuthorizationRequest
}

func (o *PreauthorizationRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *PreauthorizationRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *PreauthorizationRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *PreauthorizationRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type PreauthorizationResponse struct {
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	// OK
	PaymentsEntity *shared.PaymentsEntity
	StatusCode     int
	RawResponse    *http.Response
}

func (o *PreauthorizationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PreauthorizationResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *PreauthorizationResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PreauthorizationResponse) GetPaymentsEntity() *shared.PaymentsEntity {
	if o == nil {
		return nil
	}
	return o.PaymentsEntity
}

func (o *PreauthorizationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PreauthorizationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
