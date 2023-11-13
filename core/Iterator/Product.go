package Iterator

import (
	"ApricusFuzz/core/payload"
	"sync"
)

var lock sync.Mutex

type ProductIterator struct {
	isEnd bool
	token []string

	channel chan []string
}

func NewProductIterator() *ProductIterator {
	return &ProductIterator{
		isEnd:   false,
		channel: make(chan []string),
	}
}
func (iter *ProductIterator) crossCombine(payloads []payload.Payload, currentCombination []string, index int) {
	if index == len(payloads) {
		iter.channel <- append([]string{}, currentCombination...)
		return
	}
	for i := 0; i < len(payloads[index].Data); i++ {
		currentCombination = append(currentCombination, payloads[index].Data[i])
		iter.crossCombine(payloads, currentCombination, index+1)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
func (iter *ProductIterator) Exec(Payloads []payload.Payload) {
	go func() {
		defer close(iter.channel)
		var currentCombination []string

		iter.crossCombine(Payloads, currentCombination, 0)
	}()
}

func (iter *ProductIterator) IfEnd() bool {
	return iter.isEnd
}

func (iter *ProductIterator) Scan() bool {
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

func (iter *ProductIterator) Value() []string {
	return iter.token
}

func (iter *ProductIterator) Channel() chan []string {
	return iter.channel
}
