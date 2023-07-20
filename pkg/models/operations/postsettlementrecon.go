// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"net/http"
)

type PostSettlementReconRequest struct {
	FetchSettlementReconRequest *shared.FetchSettlementReconRequest `request:"mediaType=application/json"`
	XAPIVersion                 *string                             `header:"style=simple,explode=false,name=x-api-version"`
	XClientID                   string                              `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret               string                              `header:"style=simple,explode=false,name=x-client-secret"`
}

func (o *PostSettlementReconRequest) GetFetchSettlementReconRequest() *shared.FetchSettlementReconRequest {
	if o == nil {
		return nil
	}
	return o.FetchSettlementReconRequest
}

func (o *PostSettlementReconRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *PostSettlementReconRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *PostSettlementReconRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type PostSettlementReconResponse struct {
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	// OK
	FetchSettlementRecon *shared.FetchSettlementRecon
	Headers              map[string][]string
	StatusCode           int
	RawResponse          *http.Response
}

func (o *PostSettlementReconResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostSettlementReconResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *PostSettlementReconResponse) GetFetchSettlementRecon() *shared.FetchSettlementRecon {
	if o == nil {
		return nil
	}
	return o.FetchSettlementRecon
}

func (o *PostSettlementReconResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostSettlementReconResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostSettlementReconResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
