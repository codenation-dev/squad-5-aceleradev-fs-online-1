package webcrawler

import (
	"bytes"
	"encoding/csv"
	"io"
)

// CsvReader le o csv e envia para os channels
func CsvReader(body *[]byte, channels *[]chan []string, errOutput *chan error) {
	defer func() {
		for _, c := range *channels {
			close(c)
		}
		close(*errOutput)
	}()

	br := bytes.NewReader(*body)

	r := csv.NewReader(br)
	r.Comma = ';'

	size := len(*channels)
	pos := 0

	//for index := 0; index < 10; index++ {

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			*errOutput <- err
			return
		}

		(*channels)[pos] <- record

		pos++

		if pos >= size {
			pos = 0
		}
	}

	*errOutput <- nil
}
