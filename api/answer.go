package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

type createAnswerRequest struct {
	QuestionID int64  `json:"question_id" binding:"required"`
	Text       string `json:"text" binding:"required"`
	IsCorrect  bool   `json:"is_correct" binding:"omitempty"`
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

type getAnswerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAnswer(ctx *gin.Context) {
	var req getAnswerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	answer, err := server.store.GetAnswer(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answer)
}

type listAnswersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAnswers(ctx *gin.Context) {
	var req listAnswersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAnswersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	answers, err := server.store.ListAnswers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answers)
}

type deleteAnswerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAnswer(ctx *gin.Context) {
	var req deleteAnswerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAnswer(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type updateAnswerParams struct {
	ID        int64  `json:"id" binding:"required,min=1"`
	Text      string `json:"text" binding:"omitempty"`
	IsCorrect *bool  `json:"is_correct" binding:"omitempty"`
}

func (server *Server) updateAnswer(ctx *gin.Context) {
	var req updateAnswerParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	currentAnswer, err := server.store.GetAnswer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.Text == "" {
		req.Text = currentAnswer.Text
	}

	if req.IsCorrect == nil {
		req.IsCorrect = &currentAnswer.IsCorrect
	}

	arg := db.UpdateAnswerParams{
		ID:        req.ID,
		Text:      req.Text,
		IsCorrect: *req.IsCorrect,
	}

	answer, err := server.store.UpdateAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answer)
}
