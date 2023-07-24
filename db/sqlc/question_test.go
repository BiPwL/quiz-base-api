package db

import (
	"context"
	"database/sql"
	"testing"

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
	tablesUsed := [1]string{"questions"}

	createRandomQuestion(t)

	err := testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
}

func TestGetQuestion(t *testing.T) {
	tablesUsed := [1]string{"questions"}

	question1 := createRandomQuestion(t)
	question2, err := testQueries.GetQuestion(context.Background(), question1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, question2)

	require.Equal(t, question1.ID, question2.ID)
	require.Equal(t, question1.Text, question2.Text)
	require.Equal(t, question1.Hint, question2.Hint)
	require.Equal(t, question1.Category, question2.Category)

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
}

func TestUpdateQuestion(t *testing.T) {
	tablesUsed := [1]string{"questions"}

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

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
}

func TestDeleteQuestion(t *testing.T) {
	tablesUsed := [1]string{"questions"}

	question1 := createRandomQuestion(t)
	err := testQueries.DeleteQuestion(context.Background(), question1.ID)
	require.NoError(t, err)

	question2, err := testQueries.GetQuestion(context.Background(), question1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, question2)

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
}

func TestListQuestion(t *testing.T) {
	tablesUsed := [1]string{"questions"}

	for i := 0; i < 10; i++ {
		createRandomQuestion(t)
	}
	arg := ListQuestionsParams{
		Limit:  5,
		Offset: 5,
	}

	questions, err := testQueries.ListQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, questions, 5)

	for _, question := range questions {
		require.NotEmpty(t, question)
	}

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
}
