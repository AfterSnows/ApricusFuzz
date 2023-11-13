package payload

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestFilePayloadLoader_Load(t *testing.T) {

	loader := &LongWordListPayloadLoader{}
	split := loader.Load("List,")
	spew.Dump(split)
}
