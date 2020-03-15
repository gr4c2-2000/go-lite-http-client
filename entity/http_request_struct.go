package entity

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
	ExpectedResponseCode int
}

func (req *HttpRequest) AddHeader(key string ,value string){
	req.Headers[key]=value
}
func (req *HttpRequest) SetTimeout(seconds int){
	req.Client.Timeout = time.Duration(seconds)*time.Second
}

func (req *HttpRequest) SendJson(){
	req.AddHeader("content-type",	"application/json")
}
func (req *HttpRequest) SendContentType(ContentType string){
	req.AddHeader("content-type",	ContentType)
}
func (req *HttpRequest) SendXml(){
	req.AddHeader("content-type",	"application/xml")
}
func (req *HttpRequest) SendForm(){
	req.AddHeader("content-type",	"application/x-www-form-urlencoded")
}
func (req *HttpRequest) PostMethod(){
	req.Method = "POST"
}
func (req *HttpRequest) PutMethod(){
	req.Method = "PUT"
}
func (req *HttpRequest) DeleteMethod(){
	req.Method = "DELETE"
}
func (req *HttpRequest) PatchMethod(){
	req.Method = "PATCH"
}

func (req *HttpRequest) Expect200(){
	req.ExpectedResponseCode = 200
}
func (req *HttpRequest) Expect201(){
	req.ExpectedResponseCode = 201
}
func (req *HttpRequest) Expect202(){
	req.ExpectedResponseCode = 202
}

func NewHttpRequestStruct() *HttpRequest {
	req := HttpRequest{}
	req.Query = []byte{}
	req.Method = "GET"
	req.Headers = map[string]string{"Connection": "close"}
	req.Address = "http://localhost"
	req.Client = &http.Client{Timeout: time.Second * 10}
	return &req
}
