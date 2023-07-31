package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/BiPwL/quiz-base-api/util"
)

func createRandomAnsweredQuestion(t *testing.T) AnsweredQuestion {
	arg := CreateAnsweredQuestionParams{
		UserID:     createRandomUser(t).ID,
		QuestionID: createRandomQuestion(t).ID,
	}

	answeredQuestion, err := testQueries.CreateAnsweredQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, answeredQuestion)

	require.Equal(t, arg.UserID, answeredQuestion.UserID)
	require.Equal(t, arg.QuestionID, answeredQuestion.QuestionID)

	require.NotZero(t, answeredQuestion.ID)
	require.NotZero(t, answeredQuestion.AnsweredAt)

	return answeredQuestion
}

func createAnsweredQuestionsWithCategory(t *testing.T, num int, userID int64, categoryKey string) []Question {
	var expectedQuestions []Question
	for i := 0; i < num; i++ {
		argQuestion := CreateQuestionParams{
			Text: util.RandomStr(10),
			Hint: util.RandomStr(8),
		}
		if categoryKey != "" {
			argQuestion.Category = categoryKey
		} else {
			argQuestion.Category = createRandomCategory(t).Key
		}

		question, err := testQueries.CreateQuestion(context.Background(), argQuestion)
		require.NoError(t, err)

		argAnsweredQuestion := CreateAnsweredQuestionParams{
			UserID:     userID,
			QuestionID: question.ID,
		}
		_, err = testQueries.CreateAnsweredQuestion(context.Background(), argAnsweredQuestion)
		require.NoError(t, err)

		expectedQuestions = append(expectedQuestions, question)
	}

	return expectedQuestions
}

func createAnsweredQuestionsWithRandomCategory(t *testing.T, num int, userID int64) []Question {
	return createAnsweredQuestionsWithCategory(t, num, userID, "")
}

func TestCreateAnsweredQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"answered_questions", "users", "categories", "questions"})

	createRandomAnsweredQuestion(t)
}

func TestGetAnsweredQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"answered_questions", "users", "categories", "questions"})

	answeredQuestion1 := createRandomAnsweredQuestion(t)
	answeredQuestion2, err := testQueries.GetAnsweredQuestion(context.Background(), answeredQuestion1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, answeredQuestion2)

	require.Equal(t, answeredQuestion1.ID, answeredQuestion2.ID)
	require.Equal(t, answeredQuestion1.UserID, answeredQuestion2.UserID)
	require.Equal(t, answeredQuestion1.QuestionID, answeredQuestion2.QuestionID)
	require.WithinDuration(t, answeredQuestion1.AnsweredAt, answeredQuestion2.AnsweredAt, time.Second)
}

func TestDeleteAnsweredQuestion(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"answered_questions", "users", "categories", "questions"})

	answeredQuestion1 := createRandomAnsweredQuestion(t)
	err := testQueries.DeleteAnsweredQuestion(context.Background(), answeredQuestion1.ID)
	require.NoError(t, err)

	answeredQuestion2, err := testQueries.GetAnsweredQuestion(context.Background(), answeredQuestion1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, answeredQuestion2)
}

func TestListAnsweredQuestions(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"answered_questions", "users", "categories", "questions"})

	const numQuestions = 10

	for i := 0; i < numQuestions; i++ {
		createRandomAnsweredQuestion(t)
	}
	arg := ListAnsweredQuestionsParams{
		Limit:  numQuestions,
		Offset: 0,
	}

	answeredQuestions, err := testQueries.ListAnsweredQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answeredQuestions, numQuestions)

	for _, answeredQuestion := range answeredQuestions {
		require.NotEmpty(t, answeredQuestion)
	}
}
