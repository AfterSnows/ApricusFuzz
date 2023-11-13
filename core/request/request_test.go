package request

import (
	"ApricusFuzz/core/Iterator"
	"ApricusFuzz/core/payload"
	"ApricusFuzz/core/response"
	"testing"
)

func TestDoRequest(t *testing.T) {
	pool := InitPool(10000)
	payloads := []payload.Payload{}
	iter := Iterator.NewProductIterator()
	temp := NewRequest(
		"POST",
		"https://www.taobao.com/",
		[]string{
			"Cookie:cna=bBMGHcc7hGsCAbfdTIs4BNGX;",
		},
		"a=1111",
		payloads,
		iter,
	)
	resp := response.NewResponse()
	resp.NewFuzzResponses()
	pool.Start(temp, resp)
	pool.Wait()
	pool.Close()

}
