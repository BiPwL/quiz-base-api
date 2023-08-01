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
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomPasswordStr(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func innerTestListUserAnsweredQuestions(t *testing.T, numQuestions int, userID int64, categoryKey string, expectedQuestions []Question) {
	defer testQueries.CleanTables(context.Background(), []string{"categories", "questions", "answered_questions"})

	var err error

	arg := ListUserAnsweredQuestionsParams{
		Limit:    int32(numQuestions),
		Offset:   0,
		UserID:   userID,
		Category: categoryKey,
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
}

func innerTestGetUserAnsweredQuestionsCount(t *testing.T, numQuestions int, userID int64, categoryKey string) {
	defer testQueries.CleanTables(context.Background(), []string{"categories", "questions", "answered_questions"})

	arg := GetUserAnsweredQuestionsCountParams{
		UserID:   userID,
		Category: categoryKey,
	}

	count, err := testQueries.GetUserAnsweredQuestionsCount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, int64(numQuestions), count)
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
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:             user1.ID,
		HashedPassword: util.RandomPasswordStr(8),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
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
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	const numQuestions = 10

	for i := 0; i < numQuestions; i++ {
		createRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  numQuestions,
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, numQuestions)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestGetUsersCount(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users"})

	const numQuestions = 10

	for i := 0; i < numQuestions; i++ {
		createRandomUser(t)
	}

	count, err := testQueries.GetUsersCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(numQuestions), count)
}

func TestListUserAnsweredQuestions(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users", "categories"})

	const numQuestions = 5
	var expectedQuestions []Question
	user := createRandomUser(t)
	category := createRandomCategory(t)

	// test with category in request
	expectedQuestions = createAnsweredQuestionsWithCategory(t, numQuestions, user.ID, category.Key)
	innerTestListUserAnsweredQuestions(t, numQuestions, user.ID, category.Key, expectedQuestions)

	// test "without" category in request
	expectedQuestions = createAnsweredQuestionsWithRandomCategory(t, numQuestions, user.ID)
	innerTestListUserAnsweredQuestions(t, numQuestions, user.ID, "", expectedQuestions)
}

func TestGetUserAnsweredQuestionsCount(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"users", "categories"})

	const numQuestions = 5
	user := createRandomUser(t)
	category := createRandomCategory(t)

	// test with category in request
	createAnsweredQuestionsWithCategory(t, numQuestions, user.ID, category.Key)
	innerTestGetUserAnsweredQuestionsCount(t, numQuestions, user.ID, category.Key)

	// test "without" category in request
	createAnsweredQuestionsWithRandomCategory(t, numQuestions, user.ID)
	innerTestGetUserAnsweredQuestionsCount(t, numQuestions, user.ID, "")
}
