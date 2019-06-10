package controller

import (
	"app/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PublicAgentController struct
type PublicAgentController struct {
	PublicAgents service.PublicAgents
}

// StartProcess inicia processo para carregar a lista de funcionários publicos de São Paulo
func (pac PublicAgentController) StartProcess(c *gin.Context) {
	go pac.PublicAgents.StartProcess()
	c.AbortWithStatus(http.StatusNoContent)
}
