package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/BiPwL/quiz-base-api/util"
)

func createRandomAnswer(t *testing.T) Answer {
	arg := CreateAnswerParams{
		QuestionID: createRandomQuestion(t).ID,
		Text:       util.RandomStr(10),
		IsCorrect:  util.RandomBool(),
	}

	answer, err := testQueries.CreateAnswer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, answer)

	require.Equal(t, arg.QuestionID, answer.QuestionID)
	require.Equal(t, arg.Text, answer.Text)
	require.Equal(t, arg.IsCorrect, answer.IsCorrect)

	require.NotZero(t, answer.ID)

	return answer
}

func TestCreateAnswer(t *testing.T) {
	tablesUsed := [3]string{"answers", "questions", "categories"}

	createRandomAnswer(t)

	for _, table := range tablesUsed {
		err := testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestGetAnswer(t *testing.T) {
	tablesUsed := [3]string{"answers", "questions", "categories"}

	answer1 := createRandomAnswer(t)
	answer2, err := testQueries.GetAnswer(context.Background(), answer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, answer2)

	require.Equal(t, answer1.ID, answer2.ID)
	require.Equal(t, answer1.QuestionID, answer2.QuestionID)
	require.Equal(t, answer1.Text, answer2.Text)
	require.Equal(t, answer1.IsCorrect, answer2.IsCorrect)

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestUpdateAnswer(t *testing.T) {
	tablesUsed := [3]string{"answers", "questions", "categories"}

	answer1 := createRandomAnswer(t)

	arg := UpdateAnswerParams{
		ID:        answer1.ID,
		Text:      util.RandomStr(20),
		IsCorrect: util.RandomBool(),
	}

	answer2, err := testQueries.UpdateAnswer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, answer2)

	require.Equal(t, answer1.ID, answer2.ID)
	require.Equal(t, answer1.QuestionID, answer2.QuestionID)
	require.Equal(t, arg.Text, answer2.Text)
	require.Equal(t, arg.IsCorrect, answer2.IsCorrect)

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestDeleteAnswer(t *testing.T) {
	tablesUsed := [3]string{"answers", "questions", "categories"}

	answer1 := createRandomAnswer(t)
	err := testQueries.DeleteAnswer(context.Background(), answer1.ID)
	require.NoError(t, err)

	answer2, err := testQueries.GetAnswer(context.Background(), answer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, answer2)

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestListAnswers(t *testing.T) {
	tablesUsed := [3]string{"answers", "questions", "categories"}

	const numQuestions = 10

	for i := 0; i < numQuestions; i++ {
		createRandomAnswer(t)
	}
	arg := ListAnswersParams{
		Limit:  numQuestions,
		Offset: 0,
	}

	answers, err := testQueries.ListAnswers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answers, numQuestions)

	for _, answer := range answers {
		require.NotEmpty(t, answer)
	}

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestGetAnswersCount(t *testing.T) {
	tablesUsed := [3]string{"answers", "questions", "categories"}

	const numQuestions = 10

	for i := 0; i < numQuestions; i++ {
		createRandomAnswer(t)
	}

	count, err := testQueries.GetAnswersCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(numQuestions), count)

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}
