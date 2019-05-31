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

	//log.SetOutput(gin.DefaultWriter)

	userRepository := repository.UserRepository{db}
	userService := service.UserService{userRepository}

	uc := controller.UserController{&userService}

	router.GET("/", Index)
	router.POST("/users", controller.CreateUser(&userRepository))
	router.GET("/users/:userId", controller.GetUser(&userRepository))
	router.GET("/users", uc.ListUser)
	router.PUT("/users/:userId", uc.UpdateUser))
	
	return router
}

// Index rota raiz
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
