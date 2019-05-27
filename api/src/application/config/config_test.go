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

func TestGetEnvDefault(t *testing.T) {
	os.Clearenv()
	got := GetEnv("VAR_OK", "OK")
	assert.Equal(t, "OK", got)
}
