package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/BiPwL/quiz-base-api/util"
)

func createRandomQuestion(t *testing.T) Question {
	arg := CreateQuestionParams{
		Text:     util.RandomStr(30),
		Hint:     util.RandomStr(15),
		Category: createRandomCategory(t).Key,
	}

	question, err := testQueries.CreateQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, question)

	require.Equal(t, arg.Text, question.Text)
	require.Equal(t, arg.Hint, question.Hint)
	require.Equal(t, arg.Category, question.Category)

	require.NotZero(t, question.ID)
	require.NotZero(t, question.Text)
	require.NotZero(t, question.Hint)
	require.NotZero(t, question.Category)

	return question
}

func TestCreateQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"questions", "categories"})

	createRandomQuestion(t)
}

func TestGetQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"questions", "categories"})

	question1 := createRandomQuestion(t)
	question2, err := testQueries.GetQuestion(context.Background(), question1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, question2)

	require.Equal(t, question1.ID, question2.ID)
	require.Equal(t, question1.Text, question2.Text)
	require.Equal(t, question1.Hint, question2.Hint)
	require.Equal(t, question1.Category, question2.Category)
}

func TestUpdateQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"questions", "categories"})

	question1 := createRandomQuestion(t)

	arg := UpdateQuestionParams{
		ID:   question1.ID,
		Text: util.RandomStr(10),
		Hint: util.RandomStr(5),
	}

	question2, err := testQueries.UpdateQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, question2)

	require.Equal(t, question1.ID, question2.ID)
	require.Equal(t, arg.Text, question2.Text)
	require.Equal(t, arg.Hint, question2.Hint)
}

func TestDeleteQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"questions", "categories"})

	question1 := createRandomQuestion(t)
	err := testQueries.DeleteQuestion(context.Background(), question1.ID)
	require.NoError(t, err)

	question2, err := testQueries.GetQuestion(context.Background(), question1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, question2)
}

func TestListQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"questions", "categories"})

	const numQuestions = 10

	for i := 0; i < numQuestions; i++ {
		createRandomQuestion(t)
	}
	arg := ListQuestionsParams{
		Limit:  numQuestions,
		Offset: 0,
	}

	questions, err := testQueries.ListQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, questions, numQuestions)

	for _, question := range questions {
		require.NotEmpty(t, question)
	}
}

func TestListQuestionAnswers(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"answers", "questions", "categories"})

	const numQuestions = 10
	question := createRandomQuestion(t)
	expectedAnswers := [numQuestions]Answer{}
	var err error

	for i := 0; i < numQuestions; i++ {
		answer := CreateAnswerParams{
			QuestionID: question.ID,
			Text:       util.RandomStr(8),
			IsCorrect:  util.RandomBool(),
		}
		expectedAnswers[i], err = testQueries.CreateAnswer(context.Background(), answer)
		require.NoError(t, err)
	}

	arg := ListQuestionAnswersParams{
		QuestionID: question.ID,
		Limit:      numQuestions,
		Offset:     0,
	}

	answers, err := testQueries.ListQuestionAnswers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answers, numQuestions)

	for i, answer := range answers {
		require.NotEmpty(t, question)
		require.Equal(t, expectedAnswers[i].ID, answer.ID)
		require.Equal(t, expectedAnswers[i].QuestionID, answer.QuestionID)
		require.Equal(t, expectedAnswers[i].Text, answer.Text)
		require.Equal(t, expectedAnswers[i].IsCorrect, answer.IsCorrect)
		require.WithinDuration(t, expectedAnswers[i].CreatedAt, answer.CreatedAt, time.Second)
	}
}

func TestGetQuestionAnswersCount(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"answers", "questions", "categories"})

	question := createRandomQuestion(t)

	for i := 0; i < 3; i++ {
		answer := CreateAnswerParams{
			QuestionID: question.ID,
			Text:       util.RandomStr(8),
			IsCorrect:  util.RandomBool(),
		}
		_, err := testQueries.CreateAnswer(context.Background(), answer)
		require.NoError(t, err)
	}

	count, err := testQueries.GetQuestionAnswersCount(context.Background(), question.ID)
	require.NoError(t, err)
	require.Equal(t, int64(3), count)

	nonExistentQuestion := "non_existent_question"
	nonExistentCount, err := testQueries.GetCategoryQuestionsCount(context.Background(), nonExistentQuestion)
	require.NoError(t, err)
	require.Equal(t, int64(0), nonExistentCount)
}

func TestGetQuestionsCount(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"questions", "categories"})

	for i := 0; i < 5; i++ {
		createRandomQuestion(t)
	}

	count, err := testQueries.GetQuestionsCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(numQuestions), count)
}
