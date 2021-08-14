package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateWalletRequest struct {
	Name string `json:"name" example:"tobi"`
}

func(s *Server) HandleWalletCreate(ctx *gin.Context) {
	var req CreateWalletRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var rem string

	ctx.JSON(http.StatusOK, rem)

}
