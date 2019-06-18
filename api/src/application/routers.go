package application

import (
	"app/domain/service"
	"app/resources/repository"
	"log"
	"net/http"

	"app/application/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// NewRouter retorna um novo router.
func NewRouter(db *xorm.Engine) *gin.Engine {
	router := gin.Default()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(gin.DefaultWriter)

	userRepository := repository.UserRepository{
		DB: db,
	}
	userService := service.UserService{
		Repository: userRepository,
	}

	uc := controller.UserController{
		Users: &userService,
	}

	CustomerRepository := repository.CustomerRepository{
		DB: db,
	}

	customerService := service.CustomerService{
		Repository: &CustomerRepository,
	}
	cc := controller.CustomerController{
		Customers: &customerService,
	}

	publicAgentRepository := repository.PublicAgentRepository{
		DB: db,
	}

	publicAgentService := service.PublicAgentService{
		Repository: &publicAgentRepository,
	}

	pac := controller.PublicAgentController{
		PublicAgents: &publicAgentService,
	}

	alertRepository := repository.AlertRepository{
		DB: db,
	}
	alertService := service.AlertService{
		Repository: &alertRepository,
	}
	ac := controller.AlertController{
		Alerts: &alertService,
	}

	// Users
	router.GET("/", Index)
	router.POST("/users", uc.CreateUser)
	router.GET("/users/:userId", uc.GetUser)
	router.GET("/users", uc.ListUser)
	router.PUT("/users/:userId", uc.UpdateUser)

	// Customers
	router.POST("/customers", cc.UploadCustomer)

	// Public Agents
	router.GET("/webcrawler", pac.StartProcess)

	// Alerts
	router.GET("/alerts/:id", ac.GetAlert)
	router.GET("/alerts", ac.ListAlerts)

	return router
}

// Index rota raiz
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
