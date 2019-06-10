package webcrawler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCsvReader(t *testing.T) {
	body := []byte(`c1;c2;c3
1;test;123
2;2test;456`)
	channels := make([]chan []string, 0)
	channels = append(channels, make(chan []string))
	channels = append(channels, make(chan []string))

	errChannel := make(chan error)
	go CsvReader(&body, &channels, errChannel)

	select {
	case r1 := <-channels[0]:
		assert.Equal(t, []string{"1", "test", "123"}, r1)
	case r2 := <-channels[1]:
		assert.Equal(t, []string{"2", "2test", "456"}, r2)
	case err := <-errChannel:
		assert.Nil(t, err)
	case <-time.After(5 * time.Second):
		assert.Fail(t, "Timeout error")
	}

}
