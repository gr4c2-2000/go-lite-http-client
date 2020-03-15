package client

import (
	"bytes"
	"go-lite-http-client/errors"
	"go-lite-http-client/types"
	"net/http"
)

func HttpRequestClient(request *types.HttpRequest) (*http.Response, error) {

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
	if &request.ExpectedResponseCode != nil{
		if request.ExpectedResponseCode != resp.StatusCode{
			error := errors.WrongResponseCodeError{Response: resp}
			return &http.Response{}, &error
		}
	}
	defer resp.Body.Close()
	return resp, nil
}
