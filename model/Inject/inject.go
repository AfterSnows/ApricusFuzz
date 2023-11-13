package Inject

import (
	"ApricusFuzz/core/Iterator"
	"ApricusFuzz/core/payload"
	"ApricusFuzz/core/request"
	"ApricusFuzz/core/response"
	"ApricusFuzz/core/utils"
	"github.com/go-resty/resty/v2"
	"log"
)

func WafCheck(initresp string, checkpacket *utils.ParsedPacket) bool {
	utils.ReplaceHeaderFuzzTag(checkpacket)
	utils.ReplaceBodyFuzzTag(checkpacket)
	pool := request.InitPool(10)
	payloads := []payload.Payload{
		payload.NewPayload("file,./Asset/InjectPayload.txt"), // waf检测
	}
	iter := Iterator.NewZipIterator()
	temp := request.NewParseRequest(
		checkpacket.Method,
		checkpacket.URL,
		checkpacket.Headers,
		checkpacket.Body,
		payloads,
		iter,
	)
	resp := response.NewResponse()
	resp.NewFuzzResponses()
	pool.Start(temp, resp)
	pool.Wait()
	pool.Close()
	curresp := resp.FuzzResponses.Responses[0]
	Ratio := utils.CalculateSimilarityRatio(initresp, string(curresp.Body()))
	if Ratio < 0.5 {
		return true
	}
	return false

}
func heuristicsGetInject(checkpacket *utils.ParsedPacket) bool {
	// todo url中所有可能情况注入

	pool := request.InitPool(10)
	payloads := []payload.Payload{
		payload.Payload{
			Data: utils.RandomStr([]string{",", "'", "\"", "(", ")", "."}, 10),
		},
	}
	iter := Iterator.NewZipIterator()
	temp := request.NewParseRequest(
		checkpacket.Method,
		checkpacket.URL,
		checkpacket.Headers,
		checkpacket.Body,
		payloads,
		iter,
	)
	resps := response.NewResponse()
	resps.NewFuzzResponses()
	pool.Start(temp, resps)
	pool.Wait()
	pool.Close()
	for _, resp := range resps.FuzzResponses.Responses {
		if ExtractErrorFromBody(string(resp.Body())) {
			log.Print("匹配相应dms报错信息，目标可能存在潜在sql注入威胁，建议使用sqlmap进行进一步测试")
			return true
		}
	}
	return false
}

func heuristicsPostInject(checkpacket *utils.ParsedPacket) bool {
	utils.ReplaceHeaderFuzzTag(checkpacket)
	utils.ReplaceBodyFuzzTag(checkpacket)
	pool := request.InitPool(10)
	payloads := []payload.Payload{
		payload.Payload{
			Data: utils.RandomStr([]string{",", "'", "\"", "(", ")", "."}, 10),
		},
	}
	iter := Iterator.NewZipIterator()
	temp := request.NewParseRequest(
		checkpacket.Method,
		checkpacket.URL,
		checkpacket.Headers,
		checkpacket.Body,
		payloads,
		iter,
	)
	resps := response.NewResponse()
	resps.NewFuzzResponses()
	pool.Start(temp, resps)
	pool.Wait()
	pool.Close()
	for _, resp := range resps.FuzzResponses.Responses {
		if ExtractErrorFromBody(string(resp.Body())) {
			log.Print("匹配相应dms报错信息，目标可能存在潜在sql注入威胁，建议使用sqlmap进行进一步测试")
			return true
		}
	}
	return false
}

// InitCheck todo 网站连通性检测 WAF探测 网页稳定性检测 参数动态性检测 启发式注入检测 误报检测 :(太难了
func InitCheck(packet string) {
	var initresp *resty.Response
	var parsedpacket *utils.ParsedPacket
	parsedpacket, _ = utils.ParsePacket(packet)
	initresp = request.ParseNewInitRequest(parsedpacket)
	if string(initresp.Body()) == "" || initresp.Status() != "200" { //网站连通性检测
		log.Fatal("website can't touchable")
		return
	}
	secondresp := request.ParseNewInitRequest(parsedpacket)
	if !utils.IsTruePage(string(initresp.Body()), string(secondresp.Body())) { //网页稳定性检测
		log.Fatal("website 不稳定")
		return
	}
	use0, _ := utils.ParsePacket(packet)
	if WafCheck(string(initresp.Body()), use0) {
		log.Fatal("website 存在waf")
		return
	}
	switch parsedpacket.Method {
	case "GET":
		use1, _ := utils.ParsePacket(packet)
		heuristicsGetInject(use1)
		break
	case "POST":
		use2, _ := utils.ParsePacket(packet)
		heuristicsPostInject(use2)
		break
	}

}
