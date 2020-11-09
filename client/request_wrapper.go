package client

import (
	"bytes"
	"github.com/gr4c2-2000/go-lite-http-client/entity"
	"github.com/gr4c2-2000/go-lite-http-client/errors"
	"net/http"
)

func HttpRequestClient(request *entity.HttpRequest) (*http.Response, error) {

	req, err := http.NewRequest(request.Method, request.Address, bytes.NewBuffer(request.Query))
	if err != nil {
		return &http.Response{}, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	resp, err := request.Client.Do(req)
	if err != nil {
		return resp, err
	}
	if &request.ExpectedResponseCode != nil {
		if request.ExpectedResponseCode != resp.StatusCode {
			responseCodeErr := errors.WrongResponseCodeError{Response: resp}
			return &http.Response{}, &responseCodeErr
		}
	}
	defer resp.Body.Close()
	return resp, nil
}
