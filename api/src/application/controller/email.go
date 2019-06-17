package controller

import (
	"app/domain/model"
	"app/domain/service"
	"app/domain/validator"
	"app/resources/repository"
	"app/resources/sendemail"
	"log"

	"github.com/go-xorm/xorm"
)

// EmailChannel canal que irá receber todos os alertas.
var EmailChannel chan model.Email

// EmailController struct
type EmailController struct {
	EmailService service.EmailService
}

func init() {
	log.Println("Init package")
	EmailChannel = make(chan model.Email, 2)
}

// Send envia o email
func (ec *EmailController) Send(e model.Email) error {

	err := ec.EmailService.Send(e)

	return err

}

// InitSendEmail inicia o servidor de email
func InitSendEmail(db *xorm.Engine) {

	log.Println("inicializado servidor de email")

	ec := newEmailController()

	go func(emailChannel chan model.Email) {

		for {
			e, _ := <-emailChannel
			e.Recipients = getAllUserEmail(db)
			ec.Send(e)
			log.Println(e)
		}

	}(EmailChannel)
}

// newEmailController cria um novo controller de email
func newEmailController() *EmailController {

	ce := sendemail.ClientEmail{}
	se := service.EmailService{ce}
	ec := EmailController{se}
	return &ec

}

func getAllUserEmail(db *xorm.Engine) []string {

	ur := repository.UserRepository{
		DB: db,
	}
	us := service.UserService{
		Repository: ur,
	}

	q := &validator.UserListRequest{Limit: -1}

	users, _ := us.ListUsers(q)

	var emails []string

	for _, u := range users.Data {
		emails = append(emails, u.Email)
	}
	return emails
}
