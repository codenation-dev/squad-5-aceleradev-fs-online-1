package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newULID(t *testing.T) {
	got := NewULID()
	assert.Len(t, got, 26)

	got2 := NewULID()
	assert.Len(t, got2, 26)
	assert.NotEqual(t, got, got2)
}
