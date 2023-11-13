package response

import (
	"github.com/go-resty/resty/v2"
)

func NewResponse() *Response {
	return &Response{
		RespChannel:   make(chan *resty.Response),
		FuzzResponses: new(FuzzResponses),
	}
}
func (resp *Response) NewFuzzResponses() {
	go func() {
		for TheResp := range resp.RespChannel {
			resp.FuzzResponses.Responses = append(resp.FuzzResponses.Responses, *TheResp)
		}
	}()
}

//func (resp *Response) NewFuzzResponse(TheResp *resty.Response) {
//	resp.FuzzResponses.RawResponse = TheResp.RawResponse
//	resp.FuzzResponses.ReceivedAt = TheResp.ReceivedAt()
//	resp.FuzzResponses.ContentLength = TheResp.Size()
//	resp.FuzzResponses.StatusCode = TheResp.StatusCode()
//	resp.FuzzResponses.Time = TheResp.Time()
//	spew.Dump(resp.FuzzResponses)
//
//}
