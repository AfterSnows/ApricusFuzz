package fuzz

import (
	"regexp"
	"strings"
)

var FuzzTag = "{FUZZ%d}"

func IfFuzz(value string) bool {
	IsFuzz, _ := regexp.MatchString("{FUZZ.*}", value)
	return IsFuzz
}
func Fuzz(value string, Tag string, data []string, i int) string {
	if IfFuzz(value) {
		return strings.ReplaceAll(value, Tag, data[i])
	}
	return value
}
