package utils

import (
	"fmt"
	"testing"
)

func Test_min(t *testing.T) {
	baseResponse := "<html><body><h1>Welcome to My Website</h1></body></html><html><body><h1>123123aasd aregsadf 123asfd qewgef</h1></body></html>"
	currentResponse := "<html><body><h1>Welcome to Our Website</h1></body></html><html><body><h1>adfhbwrt fgyjk serth DSFA</h1></body></html>"

	isTrue := IsTruePage(baseResponse, currentResponse)
	calculateSimilarityRatio1 := CalculateSimilarityRatio(baseResponse, currentResponse)
	fmt.Println(isTrue)
	fmt.Println(calculateSimilarityRatio1)

}
