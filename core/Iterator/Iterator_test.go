package Iterator

import (
	"ApricusFuzz/core/payload"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestChainIteratorExec(t *testing.T) {
	payloads := []payload.Payload{
		{
			Data: []string{"1", "2", "3"},
		},
		{
			Data: []string{"a", "b", "c"},
		},
	}
	c := NewChainIterator()
	c.Exec(payloads)
	for c.Scan() {
		spew.Dump(c.Value())
	}
}

func TestZipIteratorExec(t *testing.T) {
	payloads := []payload.Payload{
		{
			Data: []string{"1", "2", "3"},
		},
		{
			Data: []string{"a", "b", "c"},
		},
		{
			Data: []string{"1", "2"},
		},
	}
	c := NewZipIterator()
	c.Exec(payloads)
	for c.Scan() {
		spew.Dump(c.Value())
	}
}

type Payload struct {
	Data []string
}

func crossCombine1(payloads []Payload, currentCombination []string, index int) {
	if index == len(payloads) {
		// 当索引达到数组末尾时，打印当前组合结果
		fmt.Println(currentCombination)
		return
	}

	// 遍历当前payload的data数组
	for i := 0; i < len(payloads[index].Data); i++ {
		// 将当前元素添加到组合中
		currentCombination = append(currentCombination, payloads[index].Data[i])

		// 递归调用，处理下一个payload
		crossCombine1(payloads, currentCombination, index+1)

		// 回溯，从组合中移除当前元素
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func TestNewProductIterator(t *testing.T) {
	payloads := []Payload{
		{
			Data: []string{"1", "2", "3"},
		},
		{
			Data: []string{"a", "b", "c"},
		},
		{
			Data: []string{"1", "2", "3"},
		},
	}

	currentCombination := []string{} // 用于存储当前的组合

	crossCombine1(payloads, currentCombination, 0)
}
func TestProductIteratorExec(t *testing.T) {
	payloads := []payload.Payload{
		{
			Data: []string{"1", "2", "3"},
		},
		{
			Data: []string{"a", "b", "c"},
		},
		{
			Data: []string{"1", "2", "3"},
		},
	}
	c := NewProductIterator()
	c.Exec(payloads)
	for c.Scan() {
		fmt.Println(c.Value())
	}
}
