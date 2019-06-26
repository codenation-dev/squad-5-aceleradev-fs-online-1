package service

import (
	"app/domain/builder"
	"app/domain/service/engine"
	"app/resources/repository"
	"app/resources/webcrawler"
	"log"
	"os"
)

// PublicAgents interface
type PublicAgents interface {
	StartProcess()
}

// PublicAgentService struct
type PublicAgentService struct {
	Repository repository.PublicAgentDB
	Alert      engine.Alert
}

const urlPublicAgent string = "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/Remuneracao.aspx"
const maxConsumers int = 8

// StartProcess inicia o processo de webcrawler
func (us PublicAgentService) StartProcess() {
	log.Println("StartProcess Begin")
	filename := "./temp" + builder.NewULID() + ".zip"

	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		f.Close()
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

	webcrawler.StoreFile(filename)

	channels := make([]chan []string, 0)
	for i := 0; i < maxConsumers; i++ {
		channels = append(channels, make(chan []string))
	}

	errChannel := make(chan error)

	go webcrawler.CsvReader(content, &channels, errChannel)

	p := make([]chan bool, 0)

	for pos, c := range channels {
		done := make(chan bool)
		p = append(p, done)
		go us.processChannel(pos, c, done)
	}

	err = <-errChannel
	if err != nil {
		log.Println(err)
		return
	}

	for _, done := range p {
		<-done
	}

	log.Println("StartProcess End")
}

func (us PublicAgentService) processChannel(pos int, c chan []string, done chan bool) {
	log.Printf("processChannel Begin %v\n", pos)
	for line := range c {
		publicAgent := builder.PublicAgentFrom(&line)
		updated, err := us.Repository.CreateOrUpdatePublicAgent(publicAgent)
		if err != nil {
			log.Printf("processChannel %#v\n", publicAgent)
			log.Println("processChannel Error", err)
		}
		if updated {
			us.Alert.PublicAgents() <- *publicAgent
		}
	}
	log.Printf("processChannel End %v\n", pos)
	done <- true
}
