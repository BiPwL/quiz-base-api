// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error)
	CreateAnsweredQuestion(ctx context.Context, arg CreateAnsweredQuestionParams) (AnsweredQuestion, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAnswer(ctx context.Context, id int64) error
	DeleteAnsweredQuestion(ctx context.Context, id int64) error
	DeleteCategory(ctx context.Context, key string) error
	DeleteQuestion(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetAnswer(ctx context.Context, id int64) (Answer, error)
	GetAnsweredQuestion(ctx context.Context, id int64) (AnsweredQuestion, error)
	GetAnswersCount(ctx context.Context) (int64, error)
	GetCategoriesCount(ctx context.Context) (int64, error)
	GetCategory(ctx context.Context, key string) (Category, error)
	GetCategoryQuestionsCount(ctx context.Context, category string) (int64, error)
	GetQuestion(ctx context.Context, id int64) (Question, error)
	GetQuestionAnswersCount(ctx context.Context, questionID int64) (int64, error)
	GetQuestionsCount(ctx context.Context) (int64, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserAnsweredQuestionsCount(ctx context.Context, arg GetUserAnsweredQuestionsCountParams) (int64, error)
	GetUsersCount(ctx context.Context) (int64, error)
	ListAnsweredQuestions(ctx context.Context, arg ListAnsweredQuestionsParams) ([]AnsweredQuestion, error)
	ListAnswers(ctx context.Context, arg ListAnswersParams) ([]Answer, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListCategoryQuestions(ctx context.Context, arg ListCategoryQuestionsParams) ([]Question, error)
	ListQuestionAnswers(ctx context.Context, arg ListQuestionAnswersParams) ([]Answer, error)
	ListQuestions(ctx context.Context, arg ListQuestionsParams) ([]Question, error)
	ListUserAnsweredQuestions(ctx context.Context, arg ListUserAnsweredQuestionsParams) ([]Question, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) (Answer, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) (Question, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
