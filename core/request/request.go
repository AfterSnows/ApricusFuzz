package request

import (
	"ApricusFuzz/core/Iterator"
	"ApricusFuzz/core/fuzz"
	"ApricusFuzz/core/payload"
	"ApricusFuzz/core/response"
	"ApricusFuzz/core/utils"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"log"
	"strings"
)

func NewRequest(method string, url string, headersString []string, body string, payloads []payload.Payload, iter Iterator.BaseIterator) *FuzzRequest {
	var (
		headersMap = make(map[string]string)
	)

	for _, header := range headersString {
		h := strings.Split(header, ": ")
		if len(h) == 2 {
			headersMap[h[0]] = h[1]
		} else {
			log.Printf("Format error HEADER(Has been ignored): %s\n", header)
			continue
		}
	}

	return &FuzzRequest{
		Method:        method,
		TargetUrl:     url,
		TargetHeaders: headersMap,
		TargetBody:    body,
		Payloads:      payloads,
		Iterator:      iter,
	}
}
func NewProxyRequest(AgentProxy string, method string, url string, headersString []string, body string, payloads []payload.Payload, iter Iterator.BaseIterator) *FuzzRequest {
	var (
		headersMap = make(map[string]string)
	)

	for _, header := range headersString {
		h := strings.Split(header, ": ")
		if len(h) == 2 {
			headersMap[h[0]] = h[1]
		} else {
			log.Printf("Format error HEADER(Has been ignored): %s\n", header)
			continue
		}
	}

	return &FuzzRequest{
		Method:        method,
		TargetUrl:     url,
		TargetHeaders: headersMap,
		TargetBody:    body,
		Payloads:      payloads,
		Iterator:      iter,
		AgentProxy:    AgentProxy,
	}
}

// NewParseRequest 配合parse packet包的合成fuzzrequest对象
func NewParseRequest(method string, url string, headersMap map[string]string, body string, payloads []payload.Payload, iter Iterator.BaseIterator) *FuzzRequest {
	return &FuzzRequest{
		Method:        method,
		TargetUrl:     url,
		TargetHeaders: headersMap,
		TargetBody:    body,
		Payloads:      payloads,
		Iterator:      iter,
	}
}

// NewPreInitRequest  配合初始原始包
func NewPreInitRequest(method string, url string, headersMap map[string]string, body string) *FuzzRequest {
	return &FuzzRequest{
		Method:        method,
		TargetUrl:     url,
		TargetHeaders: headersMap,
		TargetBody:    body,
	}
}

// DoRequest todo 差多header设置
func (fr *FuzzRequest) DoRequest(Data []string, response *response.Response) *Request {
	url := fr.TargetUrl
	header := fr.TargetHeaders
	body := fr.TargetBody
	for i := 0; i < len(Data); i++ {
		Tag := fmt.Sprintf(fuzz.FuzzTag, i)
		url = fuzz.Fuzz(url, Tag, Data, i)
		body = fuzz.Fuzz(body, Tag, Data, i)
		for k, v := range header {
			k = fuzz.Fuzz(k, Tag, Data, i)
			v = fuzz.Fuzz(v, Tag, Data, i)
		}
	}
	restClient := resty.New()

	if fr.AgentProxy != "" {
		restClient.SetProxy(fr.AgentProxy)
		log.Print("代理模式启动，代理地址:" + fr.AgentProxy)
	}
	request := resty.New().R()
	request.URL = url
	request.Method = fr.Method
	request.SetHeaders(header)
	request.SetBody(body)
	spew.Dump(request.Body)
	spew.Dump(request.URL)
	//request.SetHeader("User-Agent", UA_LIST[rand.Intn(4)])
	return &Request{
		Data:     Data,
		Request:  request,
		Response: response,
	}
}

// SimplePreInitRequest 用于前置分析请求
func SimplePreInitRequest(fr *FuzzRequest) *resty.Response {
	url := utils.GenerateFullURL(fr.TargetUrl)
	fmt.Println(url)

	header := fr.TargetHeaders
	body := fr.TargetBody
	request := resty.New().R()
	request.URL = url
	request.Method = fr.Method
	request.SetHeaders(header)
	request.SetBody(body)
	resp, err := request.Send()
	if err != nil {
		fmt.Println(err)
	}
	return resp
}

// ParseNewInitRequest 获取初始返回
func ParseNewInitRequest(p *utils.ParsedPacket) *resty.Response {
	TheRequest := NewPreInitRequest(
		p.Method,
		p.Host+p.URL,
		p.Headers,
		p.Body,
	)
	return SimplePreInitRequest(TheRequest)
}
