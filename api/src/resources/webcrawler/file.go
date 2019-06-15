package webcrawler

import (
	"app/application/config/store"
	"log"
	"os"
)

// StoreFile guarda o arquivo processo em uma novo local
func StoreFile(name string) error {
	log.Println("StoreFile from ", name, " to ", store.StorePath)
	return os.Rename(name, store.StorePath+name)
}
