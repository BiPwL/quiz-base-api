package api

import (
	"database/sql"
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

type getQuestionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getQuestion(ctx *gin.Context) {
	var req getQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	question, err := server.store.GetQuestion(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}

type listQuestionsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listQuestions(ctx *gin.Context) {
	var req listQuestionsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListQuestionsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	questions, err := server.store.ListQuestions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, questions)
}

type deleteQuestionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteQuestion(ctx *gin.Context) {
	var req deleteQuestionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetQuestion(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteQuestion(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type updateQuestionParams struct {
	ID   int64  `json:"id" binding:"required,min=1"`
	Text string `json:"text" binding:"required"`
	Hint string `json:"hint" binding:"required"`
}

func (server *Server) updateQuestion(ctx *gin.Context) {
	var req updateQuestionParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateQuestionParams{
		ID:   req.ID,
		Text: req.Text,
		Hint: req.Hint,
	}

	question, err := server.store.UpdateQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, question)
}

type listQuestionAnswersRequest struct {
	QuestionID int64 `form:"question_id" binding:"required"`
	PageID     int32 `form:"page_id" binding:"required,min=1"`
	PageSize   int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listQuestionAnswers(ctx *gin.Context) {
	var req listQuestionAnswersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListQuestionAnswersParams{
		QuestionID: req.QuestionID,
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}

	answers, err := server.store.ListQuestionAnswers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, answers)
}

type getQuestionAnswersCountRequest struct {
	QuestionID int64 `uri:"question_id" binding:"required"`
}

func (server *Server) getQuestionAnswersCount(ctx *gin.Context) {
	var req getQuestionAnswersCountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetQuestion(ctx, req.QuestionID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	questionsCount, err := server.store.GetQuestionAnswersCount(ctx, req.QuestionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, questionsCount)
}

func (server *Server) getQuestionsCount(ctx *gin.Context) {
	count, err := server.store.GetQuestionsCount(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, count)
}
