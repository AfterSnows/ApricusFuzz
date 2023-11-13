package Inject

import (
	"fmt"
	"testing"
)

func Test_extractErrorFromBody(t *testing.T) {
	responseBody := "You have an error in your SQL syntax"
	errorType := ExtractErrorFromBody(responseBody)
	fmt.Println(errorType)
}
