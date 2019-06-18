package controller

import (
	"app/domain/errors"
	"app/domain/service"
	"app/domain/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AlertController struct
type AlertController struct {
	Alerts service.Alerts
}

// GetAlert Consulta os detalhes do alerta
func (ac AlertController) GetAlert(c *gin.Context) {
	var alertURI validator.IDURI
	if err := c.ShouldBindUri(&alertURI); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	alert, err := ac.Alerts.GetAlert(alertURI.ID)

	switch {
	case err != nil:
		errors.AbortWithError(c, &err)
	case alert == nil:
		c.AbortWithStatus(http.StatusNotFound)
	default:
		c.JSON(http.StatusOK, alert)
	}

}

// ListAlerts Lista os alertas
func (ac AlertController) ListAlerts(c *gin.Context) {
	var q validator.AlertListRequest
	if err := c.ShouldBindQuery(&q); err != nil {
		validator.AbortWithValidation(c, &err)
		return
	}

	alerts, err := ac.Alerts.ListAlerts(&q)

	if err != nil {
		errors.AbortWithError(c, &err)
	} else {
		c.JSON(http.StatusOK, alerts)
	}

}
