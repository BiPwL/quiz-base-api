package api

import (
	"github.com/gin-gonic/gin"

	db "github.com/BiPwL/quiz-base-api/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// User handlers
	router.POST("users/new", server.createUser)
	router.GET("users/:id", server.getUser)
	router.GET("users", server.listUsers)
	router.DELETE("users/:id", server.deleteUser)
	router.POST("users", server.updateUser)
	// Category handlers
	router.POST("categories/new", server.createCategory)
	router.GET("categories/:key", server.getCategory)
	router.GET("categories", server.listCategories)
	router.DELETE("categories/:key", server.deleteCategory)
	router.POST("categories", server.updateCategory)
	router.GET("categories/questions", server.listCategoryQuestions)
	router.GET("categories/count/:key", server.getCategoryQuestionsCount)
	router.GET("categories/count", server.getCategoriesCount)
	// Question handlers
	router.POST("questions/new", server.createQuestion)
	router.GET("questions/:id", server.getQuestion)
	router.GET("questions", server.listQuestions)
	router.DELETE("questions/:id", server.deleteQuestion)
	router.POST("questions", server.updateQuestion)
	router.GET("questions/answers", server.listQuestionAnswers)
	router.GET("questions/count/:question_id", server.getQuestionAnswersCount)
	// Answer handlers
	router.POST("answers/new", server.createAnswer)
	router.GET("answers/:id", server.getAnswer)
	router.GET("answers", server.listAnswers)
	router.DELETE("answers/:id", server.deleteAnswer)
	router.POST("answers", server.updateAnswer)
	// Answered Questions handlers
	router.POST("answered_questions/new", server.createAnsweredQuestion)
	router.GET("answered_questions/:id", server.getAnsweredQuestion)
	router.GET("answered_questions", server.listAnsweredQuestions)
	router.DELETE("answered_questions/:id", server.deleteAnsweredQuestion)
	// dev operations handlers
	router.DELETE("clean/:name", server.cleanTable)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
