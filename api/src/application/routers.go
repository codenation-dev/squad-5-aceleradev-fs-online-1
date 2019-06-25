package application

import (
	"app/domain/service"
	"app/domain/service/engine"
	customValidator "app/domain/validator"
	"app/resources/repository"
	"log"
	"net/http"

	"app/application/controller"
	"app/application/middleware"

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
	dasboardRepository := repository.DasboardRepository{
		DB: db,
	}
	loginRepository := repository.LoginRepository{
		DB: db,
	}

	// Services
	dasboardService := service.DasboardService{
		Repository:      dasboardRepository,
		AlertRepository: alertRepository,
	}
	alertService := service.AlertService{
		Repository: &alertRepository,
	}
	engineAlertService := engine.AlertService{
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
	loginService := service.LoginService{
		Repository: loginRepository,
	}

	// Controllers
	dc := controller.DashboardController{
		Dashboads: &dasboardService,
	}
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
	lc := controller.LoginController{
		Login: loginService,
	}

	// Users
	router.GET("/", Index)

	basicAuth := router.Group("/")
	basicAuth.Use(middleware.AuthenticationRequired)
	{
		basicAuth.POST("/users", uc.CreateUser)
		basicAuth.GET("/users/:userId", uc.GetUser)
		basicAuth.GET("/users", uc.ListUser)
		basicAuth.PUT("/users/:userId", uc.UpdateUser)

		// Customers
		basicAuth.POST("/customers", cc.UploadCustomer)
		basicAuth.POST("/customer", cc.CreateCustomer)
		basicAuth.PUT("/customer/:customerId", cc.UpdateCustomer)
		basicAuth.GET("/customers", cc.ListCustomer)

		// Alerts
		basicAuth.GET("/alerts/:id", ac.GetAlert)
		basicAuth.GET("/alerts", ac.ListAlerts)

		// Dashboards
		basicAuth.GET("/dashboard/alerts", dc.GetAlerts)
		basicAuth.GET("/dashboard/customer", dc.ListCustomers)
	}

	// Public Agents
	router.GET("/webcrawler", pac.StartProcess)

	// Login
	router.POST("/auth", lc.Authorization)

	return router
}

// Index rota raiz
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
