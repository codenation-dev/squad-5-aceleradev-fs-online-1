package webcrawler

import (
	"app/domain/builder"
	"archive/zip"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnzipFile(t *testing.T) {
	bytesExpected := []byte("test 123")
	nome := "./temp" + builder.NewULID() + ".zip"

	fzip, err := os.Create(nome)
	assert.Nil(t, err)

	w := zip.NewWriter(fzip)

	f, err := w.Create("test.txt")
	assert.Nil(t, err)

	_, err = f.Write(bytesExpected)
	assert.Nil(t, err)

	w.Close()

	fzip.Close()

	b, err := UnzipFile(nome)
	assert.Nil(t, err)
	assert.Equal(t, bytesExpected, *b)

	os.Remove(nome)

}
