package application

import (
	"app/domain/service"
	"app/domain/service/engine"
	customValidator "app/domain/validator"
	"app/resources/repository"
	"log"
	"net/http"

	"app/application/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-xorm/xorm"
	"gopkg.in/go-playground/validator.v8"
)

// NewRouter retorna um novo router.
func NewRouter(db *xorm.Engine) *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("alerttype", customValidator.AlertTypeValidator)
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}

	router.Use(cors.New(config))

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(gin.DefaultWriter)

	// Repositories
	userRepository := repository.UserRepository{
		DB: db,
	}
	CustomerRepository := repository.CustomerRepository{
		DB: db,
	}
	publicAgentRepository := repository.PublicAgentRepository{
		DB: db,
	}
	alertRepository := repository.AlertRepository{
		DB: db,
	}
	alertService := service.AlertService{
		Repository: &alertRepository,
	}

	// Services
	engineAlertService := engine.EngineAlertService{
		CustomerDB:   CustomerRepository,
		AlertDB:      alertRepository,
		UserDB:       userRepository,
		EmailChannel: controller.EmailChannel,
	}
	engineAlertService.Init()
	customerService := service.CustomerService{
		Repository: &CustomerRepository,
		Alert:      engineAlertService,
	}
	publicAgentService := service.PublicAgentService{
		Repository: &publicAgentRepository,
		Alert:      engineAlertService,
	}
	userService := service.UserService{
		Repository: userRepository,
		Alert:      engineAlertService,
	}

	// Controllers
	uc := controller.UserController{
		Users: &userService,
	}
	ac := controller.AlertController{
		Alerts: &alertService,
	}
	cc := controller.CustomerController{
		Customers: &customerService,
	}
	pac := controller.PublicAgentController{
		PublicAgents: &publicAgentService,
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
