package controller

import (
	"app/domain/errors"
	"app/domain/service"
	"app/domain/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DashboardController struct
type DashboardController struct {
	Dashboads service.Dasboards
}

// GetAlerts Consulta os detalhes do alerta
func (dc DashboardController) GetAlerts(c *gin.Context) {
	var q validator.DashboardRequest
	if err := c.ShouldBindQuery(&q); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	db, err := dc.Dashboads.GetData(q)

	if err != nil {
		errors.AbortWithError(c, &err)
	} else if db == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, db)
	}

}
