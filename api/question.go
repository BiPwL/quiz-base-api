package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

type createQuestionRequest struct {
	Text     string `json:"text" binding:"required"`
	Hint     string `json:"hint" binding:"required"`
	Category string `json:"category" binding:"required"`
}

func (server *Server) createQuestion(ctx *gin.Context) {
	var req createQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateQuestionParams{
		Text:     req.Text,
		Hint:     req.Hint,
		Category: req.Category,
	}

	question, err := server.store.CreateQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}
