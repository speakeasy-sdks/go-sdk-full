// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"net/http"
)

type GetPaymentLinkDetailsRequest struct {
	LinkID        string  `pathParam:"style=simple,explode=false,name=link_id"`
	XAPIVersion   *string `header:"style=simple,explode=false,name=x-api-version"`
	XClientID     string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (o *GetPaymentLinkDetailsRequest) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

func (o *GetPaymentLinkDetailsRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *GetPaymentLinkDetailsRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *GetPaymentLinkDetailsRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type GetPaymentLinkDetailsResponse struct {
	ContentType string
	Headers     map[string][]string
	// OK
	LinkResponse *shared.LinkResponse
	StatusCode   int
	RawResponse  *http.Response
}

func (o *GetPaymentLinkDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentLinkDetailsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetPaymentLinkDetailsResponse) GetLinkResponse() *shared.LinkResponse {
	if o == nil {
		return nil
	}
	return o.LinkResponse
}

func (o *GetPaymentLinkDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentLinkDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
