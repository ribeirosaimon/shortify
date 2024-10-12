package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FuzzBase62 realiza um fuzz test na função Encode
func FuzzBase62(f *testing.F) {
	f.Add("test1")
	f.Add("example")
	f.Add("fuzzing")
	f.Add("Base62")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		encoded := NewBase62().Encode()
		assert.NotNil(t, encoded)
	})
}
