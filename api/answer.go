package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

type createAnswerRequest struct {
	QuestionID int64  `json:"question_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
	IsCorrect  bool   `json:"is_correct"`
}

func (server *Server) createAnswer(ctx *gin.Context) {
	var req createAnswerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAnswerParams{
		QuestionID: req.QuestionID,
		Text:       req.Text,
		IsCorrect:  req.IsCorrect,
	}

	answer, err := server.store.CreateAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answer)
}
