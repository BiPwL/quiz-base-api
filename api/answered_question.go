package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

type createAnsweredQuestionRequest struct {
	UserID     int64 `json:"user_id" binding:"required"`
	QuestionID int64 `json:"question_id" binding:"required"`
}

func (server *Server) createAnsweredQuestion(ctx *gin.Context) {
	var req createAnsweredQuestionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAnsweredQuestionParams{
		UserID:     req.UserID,
		QuestionID: req.QuestionID,
	}

	answeredQuestion, err := server.store.CreateAnsweredQuestion(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Constraint == "idx_unique_user_question" {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answeredQuestion)
}

type getAnsweredQuestionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAnsweredQuestion(ctx *gin.Context) {
	var req getAnsweredQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	answeredQuestion, err := server.store.GetAnsweredQuestion(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answeredQuestion)
}
