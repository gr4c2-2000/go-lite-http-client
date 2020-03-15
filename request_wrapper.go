package httpLiteClient

import (
	"bytes"
	"net/http"
)

func HttpRequestClient(request *HttpRequest) (*http.Response, error) {

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
	defer resp.Body.Close()
	return resp, nil
}
