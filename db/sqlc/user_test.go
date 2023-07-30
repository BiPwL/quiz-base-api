package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/BiPwL/quiz-base-api/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:    util.RandomEmail(),
		Password: util.RandomPassword(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:       user1.ID,
		Password: util.RandomPassword(8),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	tablesUsed := [1]string{"users"}

	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestGetUsersCount(t *testing.T) {
	tablesUsed := [1]string{"users"}

	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}

	count, err := testQueries.GetUsersCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(5), count)

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestListUserAnsweredQuestionsWithCategory(t *testing.T) {
	tablesUsed := [4]string{"users", "categories", "questions", "answered_questions"}

	user := createRandomUser(t)
	category := createRandomCategory(t)
	const numQuestions = 5
	expectedQuestions := [numQuestions]Question{}
	var err error

	for i := 0; i < numQuestions; i++ {
		argQuestion := CreateQuestionParams{
			Text:     util.RandomStr(10),
			Hint:     util.RandomStr(8),
			Category: category.Key,
		}
		expectedQuestions[i], err = testQueries.CreateQuestion(context.Background(), argQuestion)
		require.NoError(t, err)

		argAnsweredQuestion := CreateAnsweredQuestionParams{
			UserID:     user.ID,
			QuestionID: expectedQuestions[i].ID,
		}
		_, err = testQueries.CreateAnsweredQuestion(context.Background(), argAnsweredQuestion)
		require.NoError(t, err)
	}

	arg := ListUserAnsweredQuestionsParams{
		Limit:    5,
		Offset:   0,
		UserID:   user.ID,
		Category: category.Key,
	}

	answeredQuestions, err := testQueries.ListUserAnsweredQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answeredQuestions, numQuestions)

	for i, answeredQuestion := range answeredQuestions {
		require.NotEmpty(t, answeredQuestion)
		require.Equal(t, expectedQuestions[i].ID, answeredQuestion.ID)
		require.Equal(t, expectedQuestions[i].Text, answeredQuestion.Text)
		require.Equal(t, expectedQuestions[i].Hint, answeredQuestion.Hint)
		require.Equal(t, expectedQuestions[i].Category, answeredQuestion.Category)
		require.WithinDuration(t, expectedQuestions[i].CreatedAt, answeredQuestion.CreatedAt, time.Second)
	}

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}

func TestListUserAnsweredQuestionsWithoutCategory(t *testing.T) {
	tablesUsed := [4]string{"users", "categories", "questions", "answered_questions"}

	user := createRandomUser(t)
	const numQuestions = 5
	expectedQuestions := [numQuestions]Question{}
	var err error

	for i := 0; i < numQuestions; i++ {
		expectedQuestions[i] = createRandomQuestion(t)
		argAnsweredQuestion := CreateAnsweredQuestionParams{
			UserID:     user.ID,
			QuestionID: expectedQuestions[i].ID,
		}
		_, err = testQueries.CreateAnsweredQuestion(context.Background(), argAnsweredQuestion)
		require.NoError(t, err)
	}

	arg := ListUserAnsweredQuestionsParams{
		Limit:  5,
		Offset: 0,
		UserID: user.ID,
	}

	answeredQuestions, err := testQueries.ListUserAnsweredQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answeredQuestions, numQuestions)

	for i, answeredQuestion := range answeredQuestions {
		require.NotEmpty(t, answeredQuestion)
		require.Equal(t, expectedQuestions[i].ID, answeredQuestion.ID)
		require.Equal(t, expectedQuestions[i].Text, answeredQuestion.Text)
		require.Equal(t, expectedQuestions[i].Hint, answeredQuestion.Hint)
		require.Equal(t, expectedQuestions[i].Category, answeredQuestion.Category)
		require.WithinDuration(t, expectedQuestions[i].CreatedAt, answeredQuestion.CreatedAt, time.Second)
	}

	for _, table := range tablesUsed {
		err = testQueries.CleanTable(context.Background(), table)
		require.NoError(t, err)
	}
}
