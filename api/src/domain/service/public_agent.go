package service

import (
	"app/domain/builder"
	"app/resources/repository"
	"app/resources/webcrawler"
	"log"
	"os"
)

// PublicAgents interface
type PublicAgents interface {
	StartProcess() error
}

// PublicAgentService struct
type PublicAgentService struct {
	Repository repository.PublicAgentDB
}

const urlPublicAgent string = "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx"
const maxConsumers int = 4

// StartProcess inicia o processo de webcrawler
func (us PublicAgentService) StartProcess() {
	filename := "./temp" + builder.NewULID() + ".zip"

	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		f.Close()
		os.Remove(filename)
	}()
	err = webcrawler.DownloadPublicAgentsList(urlPublicAgent, f)
	if err != nil {
		log.Println(err)
		return
	}

	content, err := webcrawler.UnzipFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	channels := make([]chan []string, 0)
	for i := 0; i < maxConsumers; i++ {
		channels = append(channels, make(chan []string))
	}

	errChannel := make(chan error)

	go webcrawler.CsvReader(content, &channels, &errChannel)

	for _, c := range channels {
		go us.processChannel(&c)
	}
}

func (us PublicAgentService) processChannel(c *chan []string) {
	for line := range *c {
		us.Repository.CreateOrUpdatePublicAgent(line)
	}
}
