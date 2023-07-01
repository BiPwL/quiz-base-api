package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

func TestCreateAnsweredQuestion(t *testing.T) {
	createRandomAnsweredQuestion(t)
}

func TestGetAnsweredQuestion(t *testing.T) {
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
	answeredQuestion1 := createRandomAnsweredQuestion(t)
	err := testQueries.DeleteAnsweredQuestion(context.Background(), answeredQuestion1.ID)
	require.NoError(t, err)

	answeredQuestion2, err := testQueries.GetAnsweredQuestion(context.Background(), answeredQuestion1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, answeredQuestion2)
}

func TestListAnsweredQuestions(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAnsweredQuestion(t)
	}
	arg := ListAnsweredQuestionsParams{
		Limit:  5,
		Offset: 5,
	}

	answeredQuestions, err := testQueries.ListAnsweredQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answeredQuestions, 5)

	for _, answeredQuestion := range answeredQuestions {
		require.NotEmpty(t, answeredQuestion)
	}
}
