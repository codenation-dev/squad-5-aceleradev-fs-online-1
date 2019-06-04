package application

import (
	"app/domain/service"
	"app/resources/repository"
	"net/http"

	"app/application/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

// NewRouter retorna um novo router.
func NewRouter(db *xorm.Engine) *gin.Engine {
	router := gin.Default()

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

	// Users
	router.GET("/", Index)
	router.POST("/users", uc.CreateUser)
	router.GET("/users/:userId", uc.GetUser)
	router.GET("/users", uc.ListUser)
	router.PUT("/users/:userId", uc.UpdateUser)

	// Customers
	router.POST("/customers", cc.UploadCustomer)

	return router
}

// Index rota raiz
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
