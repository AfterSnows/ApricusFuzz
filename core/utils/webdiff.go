package utils

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"unicode"
)

// todo 对于网页的动态检测没有实现
/*
*

用于对比响应相似度,

*
*/
const (
	matchRatio     = 0.98
	ratioTolerance = 0.05
)

// 返回html纯文本
func removeHTMLElse(text string) string {
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return ""
	}

	var result strings.Builder
	var extractText func(*html.Node)
	extractText = func(n *html.Node) {
		if n.Type == html.TextNode {
			result.WriteString(n.Data)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				extractText(c)
			}
		}
	}

	extractText(doc)
	return result.String()
}

// 去除字符串中的空格和标点符号
func removeSpacesAndPunctuation(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)
}

// 计算两个字符串的相似度比率
func CalculateSimilarityRatio(text1, text2 string) float64 {
	text1 = removeSpacesAndPunctuation(text1)
	text2 = removeSpacesAndPunctuation(text2)

	diffs := levenshteinDistance(text1, text2)
	maxLength := float64(max(len(text1), len(text2)))

	ratio := maxLength - float64(diffs)
	if maxLength > 0 {
		ratio /= maxLength
	} else {
		ratio = 0
	}

	return ratio
}

// 计算两个字符串的Levenshtein距离
func levenshteinDistance(text1, text2 string) int {
	len1, len2 := len(text1), len(text2)
	matrix := make([][]int, len1+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1]+1)
			}
		}
	}

	return matrix[len1][len2]
}

// 返回三个整数中的最小值
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else if b < c {
		return b
	}
	return c
}

// 返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 判断网页是否是真页面
func IsTruePage(baseResponse, currentResponse string) bool {
	baseText := removeHTMLElse(baseResponse)
	currentText := removeHTMLElse(currentResponse)

	ratio := CalculateSimilarityRatio(baseText, currentText)

	if ratio > matchRatio+ratioTolerance {
		return true
	}

	return false
}
