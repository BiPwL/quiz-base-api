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

func TestListQuestionAnswers(t *testing.T) {
	tablesUsed := [2]string{"questions", "answers"}

	question := createRandomQuestion(t)
	expectedAnswers := [2]Answer{}
	var err error

	for i := 0; i < 2; i++ {
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
		Limit:      2,
		Offset:     0,
	}

	answers, err := testQueries.ListQuestionAnswers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answers, 2)

	for i, answer := range answers {
		require.NotEmpty(t, question)
		require.Equal(t, expectedAnswers[i].ID, answer.ID)
		require.Equal(t, expectedAnswers[i].QuestionID, answer.QuestionID)
		require.Equal(t, expectedAnswers[i].Text, answer.Text)
		require.Equal(t, expectedAnswers[i].IsCorrect, answer.IsCorrect)
		require.WithinDuration(t, expectedAnswers[i].CreatedAt, answer.CreatedAt, time.Second)
	}

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
	err = testQueries.CleanTable(context.Background(), tablesUsed[1])
	require.NoError(t, err)
}

func TestGetQuestionAnswersCount(t *testing.T) {
	tablesUsed := [2]string{"questions", "answers"}

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

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
	err = testQueries.CleanTable(context.Background(), tablesUsed[1])
	require.NoError(t, err)
}

func TestGetQuestionsCount(t *testing.T) {
	tablesUsed := [1]string{"questions"}

	for i := 0; i < 5; i++ {
		createRandomQuestion(t)
	}

	count, err := testQueries.GetQuestionsCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(5), count)

	err = testQueries.CleanTable(context.Background(), tablesUsed[0])
	require.NoError(t, err)
}
