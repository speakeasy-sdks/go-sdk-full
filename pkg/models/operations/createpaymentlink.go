// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/utils"
	"net/http"
)

type CreatePaymentLinkRequest struct {
	CreateLinkRequest *shared.CreateLinkRequest `request:"mediaType=application/json"`
	XAPIVersion       *string                   `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID         string                    `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret     string                    `header:"style=simple,explode=false,name=x-client-secret"`
}

func (c CreatePaymentLinkRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePaymentLinkRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePaymentLinkRequest) GetCreateLinkRequest() *shared.CreateLinkRequest {
	if o == nil {
		return nil
	}
	return o.CreateLinkRequest
}

func (o *CreatePaymentLinkRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *CreatePaymentLinkRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *CreatePaymentLinkRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type CreatePaymentLinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// Payment Link created
	LinkResponse *shared.LinkResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePaymentLinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePaymentLinkResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreatePaymentLinkResponse) GetLinkResponse() *shared.LinkResponse {
	if o == nil {
		return nil
	}
	return o.LinkResponse
}

func (o *CreatePaymentLinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePaymentLinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
