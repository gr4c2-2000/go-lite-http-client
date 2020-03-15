package main

import (
	"net/http"
	"time"
)

type HttpRequest struct {
	Query   []byte
	Headers map[string]string
	Address string
	Method  string
	Client  *http.Client
}

func (req *HttpRequest) AddHeader(key string ,value string){
	req.Headers[key]=value
}

func (req *HttpRequest) SetTimeout(seconds int){
	req.Client.Timeout = time.Duration(seconds)*time.Second
}

func NewHttpRequestStruct() HttpRequest {
	req := HttpRequest{}
	req.Query = []byte{}
	req.Method = "GET"
	req.Headers = map[string]string{"Connection": "close"}
	req.Address = "http://localhost"
	req.Client = &http.Client{Timeout: time.Second * 10}
	return req
}
