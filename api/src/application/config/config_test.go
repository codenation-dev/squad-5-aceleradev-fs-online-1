package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("VAR_OK", "TEST")
	got := GetEnv("VAR_OK", "OK")
	assert.Equal(t, "TEST", got)
}

func TestGetEnv_Default(t *testing.T) {
	os.Clearenv()
	got := GetEnv("VAR_OK", "OK")
	assert.Equal(t, "OK", got)
}

func TestGetEnv_Error(t *testing.T) {
	os.Clearenv()

	assert.Panics(t, func() {
		GetEnv("VAR_OK")
	})
}
