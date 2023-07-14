package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

type createCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Key  string `json:"key" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		Name: req.Name,
		Key:  req.Key,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}
