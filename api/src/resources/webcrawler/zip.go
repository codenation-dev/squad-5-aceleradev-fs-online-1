package webcrawler

import (
	"archive/zip"
	"io/ioutil"
	"log"
)

// UnzipFile descompacta o arquivo zip
func UnzipFile(nome string) (*[]byte, error) {
	log.Println("UnzipFile Begin")
	rzip, err := zip.OpenReader(nome)
	if err != nil {
		return nil, err
	}
	defer rzip.Close()
	defer log.Println("UnzipFile End")

	for _, f := range rzip.File {
		log.Println(f.Name)

		fc, err := f.Open()
		if err != nil {
			return nil, err
		}

		content, err := ioutil.ReadAll(fc)
		if err != nil {
			return nil, err
		}

		err = fc.Close()
		if err != nil {
			return nil, err
		}

		return &content, nil
	}

	return nil, err
}
