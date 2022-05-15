package netprivate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCidrCatchPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			assert.Fail(t, "no panic detected")
		}
	}()

	_ = parseCidr("test", "comment")
}
