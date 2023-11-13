package request

import (
	"ApricusFuzz/core/Iterator"
	"ApricusFuzz/core/payload"
	"ApricusFuzz/core/response"
	"github.com/go-resty/resty/v2"
)

type FuzzRequest struct {
	Method        string
	TargetUrl     string
	TargetIp      string
	TargetPort    string
	TargetHeaders map[string]string
	TargetBody    string
	Payloads      []payload.Payload
	Iterator      Iterator.BaseIterator
	value         []string
	AgentProxy    string
}

type Request struct {
	Data     []string
	Request  *resty.Request
	Response *response.Response
}
