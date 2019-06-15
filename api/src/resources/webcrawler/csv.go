package webcrawler

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
)

// CsvReader le o csv e envia para os channels
func CsvReader(body *[]byte, channels *[]chan []string, errOutput chan error) {
	log.Println("CsvReader Begin")
	defer func() {
		log.Println("CsvReader Defer")
		for _, c := range *channels {
			close(c)
		}
		close(errOutput)
		log.Println("CsvReader End")
	}()

	br := bytes.NewReader(*body)

	r := csv.NewReader(br)
	r.Comma = ';'

	size := len(*channels)
	log.Printf("CsvReader Channels = %v\n", size)
	pos := 0

	r.Read() // pula a primeira linha

	// for index := 0; index < 10; index++ {

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("CsvReader Error = %v\n", err)
			errOutput <- err
			return
		}

		(*channels)[pos] <- record

		pos++

		if pos >= size {
			pos = 0
		}
	}

	log.Println("CsvReader Ending")
	errOutput <- nil
}
