package payload

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewPayload(t *testing.T) {
	data := "file,D:\\Code Check\\Go\\ApricusFuzz\\wordlist\\general\\test.txt,base64"
	spew.Dump(NewPayload(data))
}
