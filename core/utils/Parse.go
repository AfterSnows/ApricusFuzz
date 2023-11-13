package utils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"
)

// todo  差MULTIPART数据中的匹配替换，之后考虑
type ReplaceTag interface {
	ReplaceTag(string) string
}

var replacementValue = "{FUZZ0}"

func IsXMLData(data string) bool {
	var result xml.Token
	decoder := xml.NewDecoder(bytes.NewBufferString(data))

	// 解析 XML 数据，并获取第一个解析结果的标记
	for {
		t, err := decoder.Token()
		if err != nil {
			break
		}
		result = t
		break
	}

	_, ok := result.(xml.StartElement)

	return ok
}

type AfterAndDataReplaceTag struct{}

func (ReplaceTag *AfterAndDataReplaceTag) ReplaceTag(AfterAndData string) string {
	pattern := `=([^;&]+)`

	regex := regexp.MustCompile(pattern)

	result := regex.ReplaceAllStringFunc(AfterAndData, func(s string) string {
		return "=" + replacementValue + ""
	})
	return result
}

type XMLDataReplaceTag struct{}

func (ReplaceTag *XMLDataReplaceTag) ReplaceTag(xmlData string) string {
	re := regexp.MustCompile(`>([^<]+)<`)
	modifiedData := re.ReplaceAllString(xmlData, ">"+replacementValue+"<")
	return modifiedData
}

type ArrayDataReplaceTag struct{}

func (ReplaceTag *ArrayDataReplaceTag) ReplaceTag(ArrayData string) string {
	pattern := `(\w+)\[\]=\w+`
	// 编译正则表达式
	regex := regexp.MustCompile(pattern)
	result := regex.ReplaceAllStringFunc(ArrayData, func(s string) string {
		matches := regex.FindStringSubmatch(s)
		if len(matches) > 1 {
			return fmt.Sprintf("%s[]="+replacementValue+"", matches[1])
		}
		return s
	})
	return result
}

type JsonReplaceTag struct {
}

// ReplaceJSONData todo 正则匹配很明显的问题，不能识别其他数组等形式，很难的辣，如果用json来识别，还需要预定识别结构体，很头疼，所以只能搞搞简单的识别
func (ReplaceTag *JsonReplaceTag) ReplaceTag(jsonData string) string {
	re := regexp.MustCompile(`:\s*("[^"\\]*(?:\\.[^"\\]*)*"|\b[-+]?[0-9]*\.?[0-9]+\b)([,}\]])`) // 修正后的正则表达式，匹配冒号后面的字符串或数字类型
	updatedJSON := re.ReplaceAllString(jsonData, `: `+replacementValue+`$2`)
	return updatedJSON
}
