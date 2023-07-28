package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) heathCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "one-time secret is healthy/Running.")
}
