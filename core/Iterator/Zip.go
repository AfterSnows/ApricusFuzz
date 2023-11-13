package Iterator

import (
	"ApricusFuzz/core/payload"
	"log"
)

type ZipIterator struct {
	isEnd bool
	token []string

	channel chan []string
}

func NewZipIterator() *ZipIterator {
	return &ZipIterator{
		isEnd:   false,
		channel: make(chan []string),
	}
}
func MinDataLen(Payloads []payload.Payload) int {
	minLen := len(Payloads[0].Data)
	for _, payload := range Payloads {
		if minLen > len(payload.Data) {
			minLen = len(payload.Data)
		}
	}
	return minLen
}
func (iter *ZipIterator) Exec(Payloads []payload.Payload) {
	if len(Payloads) == 0 {
		log.Fatalf("Payloads Num is 0")
	}
	go func() {
		defer close(iter.channel)
		minLen := MinDataLen(Payloads)
		for i := 0; i < minLen; i++ {
			var data []string
			for _, payload := range Payloads {
				data = append(data, payload.Data[i])
			}
			iter.channel <- data
		}

	}()
}

func (iter *ZipIterator) IfEnd() bool {
	return iter.isEnd
}

func (iter *ZipIterator) Scan() bool {
	if iter.IfEnd() {
		return false
	}

	if data, ok := <-iter.channel; ok {
		iter.token = data
		return true
	} else {
		iter.isEnd = true
		return false
	}
}

func (iter *ZipIterator) Value() []string {
	return iter.token
}

func (iter *ZipIterator) Channel() chan []string {
	return iter.channel
}
