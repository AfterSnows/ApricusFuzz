package payload

import (
	"ApricusFuzz/core/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type payloadLoader interface {
	Load(string) []string
}

type FilePayloadLoader struct{}

func (loader *FilePayloadLoader) Load(value string) []string {
	payloads := make([]string, 0)
	fp, err := utils.Open(value)
	if err != nil {
		log.Fatalf("Load File [%s] Failed: %s\n", value, err)
	}
	defer fp.Close()

	for {
		payload, ok := fp.ReadLine()
		if !ok {
			break
		}
		payloads = append(payloads, payload)
	}

	return payloads
}

type LongWordListPayloadLoader struct{}

func (loader *LongWordListPayloadLoader) Load(value string) []string {
	return strings.Split(value, "(+)")
}

type ListPayloadLoader struct{}

func (loader *ListPayloadLoader) Load(value string) []string {
	return strings.Split(value, "-")
}

type StdinPayloadLoader struct{}

func (loader *StdinPayloadLoader) Load(value string) []string {
	payloads := make([]string, 0)
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		payloads = append(payloads, scan.Text())
	}

	return payloads
}

type RangePayloadLoader struct{}

func (loader *RangePayloadLoader) Load(value string) []string {
	payloads := make([]string, 0)
	index := strings.Index(value, "-")

	if len(value[:index]) != 1 || len(value[index+1:]) != 1 {
		log.Fatalf("Syntax Error in Range List Loader: %s", value)
	}

	var (
		start = int(value[:index][0])
		end   = int(value[index+1:][0])
	)

	for i := start; i <= end; i++ {
		payloads = append(payloads, fmt.Sprintf("%c", i))
	}

	return payloads
}
