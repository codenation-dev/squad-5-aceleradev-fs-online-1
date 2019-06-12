package webcrawler

import (
	"app/domain/builder"
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadPublicAgentsList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		assert.Equal(t, "POST", r.Method)

		buf := new(bytes.Buffer)
		writer := zip.NewWriter(buf)

		fw, err := writer.Create("test.txt")
		assert.Nil(t, err)

		_, err = fw.Write([]byte("test 123"))
		assert.Nil(t, err)
		err = writer.Close()
		assert.Nil(t, err)

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=\"test.zip\"")
		io.Copy(w, buf)
		w.Write(buf.Bytes())

	}))
	defer ts.Close()

	nome := "./temp" + builder.NewULID() + ".zip"

	f, err := os.Create(nome)
	assert.Nil(t, err)

	defer func() {
		f.Close()
		os.Remove(nome)
	}()

	err = DownloadPublicAgentsList(ts.URL, f)
	assert.Nil(t, err)
	stat, err := f.Stat()
	assert.Nil(t, err)
	assert.NotZero(t, stat.Size())

}
