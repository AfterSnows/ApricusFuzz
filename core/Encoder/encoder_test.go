package Encoder

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestEncodeLoader(t *testing.T) {
	result := EncodeLoader("Base64@md5@base64", "fuckyou")
	spew.Dump(result)
}
