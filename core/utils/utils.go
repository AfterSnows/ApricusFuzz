package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type File struct {
	Fp   *os.File
	Scan *bufio.Scanner
}

func Open(filename string) (*File, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	file := &File{
		Fp:   fp,
		Scan: bufio.NewScanner(fp),
	}
	return file, nil
}

func (file *File) Write(data string) (int, error) {
	return file.Fp.WriteString(data)
}

func (file *File) ReadLine() (string, bool) {
	if file.Scan.Scan() {
		return file.Scan.Text(), true
	}
	return "", false
}

func (file *File) Close() error {
	if err := file.Scan.Err(); err != nil {
		return err
	}
	return file.Fp.Close()
}

func stripPort(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return hostport
	}
	if i := strings.IndexByte(hostport, ']'); i != -1 {
		return strings.TrimPrefix(hostport[:i], "[")
	}
	return hostport[:colon]
}

func portOnly(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return ""
	}
	if i := strings.Index(hostport, "]:"); i != -1 {
		return hostport[i+len("]:"):]
	}
	if strings.Contains(hostport, "]") {
		return ""
	}
	return hostport[colon+len(":"):]
}
func ParseHostToAddrString(host string) string {
	ip := net.ParseIP(host)
	if ip == nil {
		return host
	}

	if ret := ip.To4(); ret == nil {
		return fmt.Sprintf("[%v]", ip.String())
	}

	return host
}
func ParseStringToHostPort(raw string) (host string, port int, err error) {
	if strings.Contains(raw, "://") {
		urlObject, _ := url.Parse(raw)
		if urlObject != nil {
			// 处理 URL
			portRaw := urlObject.Port()
			portInt64, err := strconv.ParseInt(portRaw, 10, 32)
			if err != nil || portInt64 <= 0 {
				switch urlObject.Scheme {
				case "http", "ws":
					port = 80
				case "https", "wss":
					port = 443
				}
			} else {
				port = int(portInt64)
			}

			host = urlObject.Hostname()
			err = nil
			return host, port, err
		}
	}

	host = stripPort(raw)
	portStr := portOnly(raw)
	if len(portStr) <= 0 {
		return host, 0, fmt.Errorf("unknown port for [%s]", raw)
	}

	portStr = strings.TrimSpace(portStr)
	portInt64, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return host, 0, fmt.Errorf("%s parse port(%s) failed: %s", raw, portStr, err)
	}

	port = int(portInt64)
	err = nil
	return
}

// GenerateFullURL todo 这里取巧了，先采取的https看看情况，需要修改
func GenerateFullURL(host string) string {
	// 检查是否已经包含协议前缀
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		return host
	}
	// 尝试使用 https 协议
	httpsURL := fmt.Sprintf("https://%s", host)
	// 发送 https 请求，检查是否可以连接
	_, err := http.Get(httpsURL)
	if err == nil {
		return httpsURL
	}
	// 尝试使用 http 协议
	httpURL := fmt.Sprintf("http://%s", host)
	// 发送 http 请求，检查是否可以连接
	_, err = http.Get(httpURL)
	if err == nil {
		return httpURL
	}

	// 如果都无法连接，则返回原始主机名
	return host
}

func CheckHostRight(Hostname string) bool {
	// 如果不是ip，那么就通过 socket 检测域名的连通性，不连通就报错
	// 测试 DNS查询是否正常
	IpMatch, _ := regexp.MatchString("\\A\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\Z", Hostname)
	// 没有匹配到那么就是域名
	if !IpMatch {
		_, err := net.LookupHost(Hostname)
		if err != nil {
			log.Fatalf("Unkown hostnaame %s", Hostname)
			return false
		}
	}
	// 测试网络连通性
	//code := 200
	//// 根据状态码来进行判断
	//if code < 200 || code >= 300 {
	//	return false
	//}
	return true
}

func HostPort(host string, port interface{}) string {
	return fmt.Sprintf("%v:%v", ParseHostToAddrString(host), port)
}
func RandomStr(strArr []string, num int) []string {
	rand.Seed(time.Now().UnixNano())

	result := make([]string, 0, num)
	length := len(strArr)

	for i := 0; i < num; i++ {
		randomStr := ""
		for j := 0; j < length; j++ {
			randomIndex := rand.Intn(length)
			randomStr += strArr[randomIndex]
		}
		result = append(result, randomStr)
	}

	return result
}
