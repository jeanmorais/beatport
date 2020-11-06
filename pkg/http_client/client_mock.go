package httpclient

import (
	"io"
	"net/http"
)

// ClientMock ia a mock of HTTPClient interface
type ClientMock struct {
	RequestBody    string
	ResponseBody   io.ReadCloser
	ResponseStatus int
	Error          error
}

// Do indicates a call of Do
func (hcm *ClientMock) Do(req *http.Request) (*http.Response, error) {

	if hcm.Error != nil {
		return nil, hcm.Error
	}

	return &http.Response{
		StatusCode: hcm.ResponseStatus,
		Body:       hcm.ResponseBody,
	}, nil
}
