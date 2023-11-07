// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v2/pkg/utils"
	"net/http"
)

type GetTerminalByMobileNumberRequest struct {
	TerminalPhoneNo string  `pathParam:"style=simple,explode=false,name=terminal_phone_no"`
	XAPIVersion     *string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID       string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret   string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (g GetTerminalByMobileNumberRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTerminalByMobileNumberRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTerminalByMobileNumberRequest) GetTerminalPhoneNo() string {
	if o == nil {
		return ""
	}
	return o.TerminalPhoneNo
}

func (o *GetTerminalByMobileNumberRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *GetTerminalByMobileNumberRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *GetTerminalByMobileNumberRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type GetTerminalByMobileNumberResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	Headers       map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TerminalDetails *shared.TerminalDetails
}

func (o *GetTerminalByMobileNumberResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTerminalByMobileNumberResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *GetTerminalByMobileNumberResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetTerminalByMobileNumberResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTerminalByMobileNumberResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTerminalByMobileNumberResponse) GetTerminalDetails() *shared.TerminalDetails {
	if o == nil {
		return nil
	}
	return o.TerminalDetails
}
