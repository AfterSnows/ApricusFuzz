package utils

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_replaceHeaderFuzzTag(t *testing.T) {
	TestPostBaiduPacket := "POST /mcp/pc/pcsearch HTTP/1.1\nHost: ug.baidu.com\nAccept: */*\nAccept-Encoding: gzip, deflate, br\nAccept-Language: zh-CN,zh;q=0.9\nContent-Length: 56\nContent-Type: application/json\nCookie: BIDUPSID=85443E956DE07BD0886B7DC0AE322B97; PSTM=1699585088; BAIDUID=85443E956DE07BD0EDC984D4E5012052:FG=1; BA_HECTOR=2h802ha4240h2gak042080241ikr71v1r; BAIDUID_BFESS=85443E956DE07BD0EDC984D4E5012052:FG=1; ZFY=26UfSY4JFTMm1msBZXkBCbY3VO9fODjK:BoBUcODdpgU:C; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=39648_39668_39663_39683_39695; PSINO=1; delPer=0\nOrigin: https://www.baidu.com\nReferer: https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&rsv_idx=1&tn=baidu&wd=123&fenlei=256&rsv_pq=0xe459ed9700059173&rsv_t=4236rCFWM46mcavC7%2B1ZXjCSgmvZSwILWYoS%2FN39m%2B%2BB84LxGzdpq5VfCmgK&rqlang=en&rsv_enter=0&rsv_dl=tb&rsv_sug3=7&rsv_sug1=1&rsv_sug7=100&rsv_btype=i&inputT=5179&rsv_sug4=5179\nSec-Fetch-Dest: empty\nSec-Fetch-Mode: cors\nSec-Fetch-Site: same-site\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36\nsec-ch-ua: \"Google Chrome\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"\nsec-ch-ua-mobile: ?0\nsec-ch-ua-platform: \"Windows\"\n\n{\"invoke_info\":{\"pos_1\":[{}],\"pos_2\":[{}],\"pos_3\":[{}]}}"
	PostBaiduPacket, _ := ParsePacket(TestPostBaiduPacket)
	spew.Dump(ReplaceHeaderFuzzTag(PostBaiduPacket).Headers)
	spew.Dump(PostBaiduPacket.Headers)
}
