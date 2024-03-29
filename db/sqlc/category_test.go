package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/BiPwL/quiz-base-api/util"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Key:  util.RandomStr(5),
		Name: util.RandomStr(5),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Key, category.Key)
	require.Equal(t, arg.Name, category.Name)

	require.NotZero(t, category.Key)
	require.NotZero(t, category.Name)

	return category
}

func TestCreateCategory(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories"})

	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories"})

	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.Key)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.Key, category2.Key)
	require.Equal(t, category1.Name, category2.Name)
}

func TestUpdateCategory(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories"})

	category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		Key:  category1.Key,
		Name: util.RandomStr(5),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.Key, category2.Key)
	require.Equal(t, arg.Name, category2.Name)
}

func TestDeleteCategory(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories"})

	category1 := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.Key)
	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.Key)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestListCategories(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories"})

	const numCategories = 10

	for i := 0; i < numCategories; i++ {
		createRandomCategory(t)
	}
	arg := ListCategoriesParams{
		Limit:  numCategories,
		Offset: 0,
	}

	categories, err := testQueries.ListCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, numCategories)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}

func TestListCategoryQuestions(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories", "questions"})

	const numQuestions = 10
	category := createRandomCategory(t)
	expectedQuestions := [numQuestions]Question{}
	var err error

	for i := 0; i < numQuestions; i++ {
		question := CreateQuestionParams{
			Text:     util.RandomStr(8),
			Hint:     util.RandomStr(6),
			Category: category.Key,
		}
		expectedQuestions[i], err = testQueries.CreateQuestion(context.Background(), question)
		require.NoError(t, err)
	}

	arg := ListCategoryQuestionsParams{
		Category: category.Key,
		Limit:    numQuestions,
		Offset:   0,
	}

	questions, err := testQueries.ListCategoryQuestions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, questions, numQuestions)

	for i, question := range questions {
		require.NotEmpty(t, question)
		require.Equal(t, expectedQuestions[i].ID, question.ID)
		require.Equal(t, expectedQuestions[i].Text, question.Text)
		require.Equal(t, expectedQuestions[i].Hint, question.Hint)
		require.Equal(t, expectedQuestions[i].Category, question.Category)
		require.WithinDuration(t, expectedQuestions[i].CreatedAt, question.CreatedAt, time.Second)
	}
}

func TestGetCategoryQuestionsCount(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories", "questions"})

	const numQuestions = 10
	category := createRandomCategory(t)

	for i := 0; i < numQuestions; i++ {
		question := CreateQuestionParams{
			Text:     util.RandomStr(8),
			Hint:     util.RandomStr(6),
			Category: category.Key,
		}
		_, err := testQueries.CreateQuestion(context.Background(), question)
		require.NoError(t, err)
	}

	count, err := testQueries.GetCategoryQuestionsCount(context.Background(), category.Key)
	require.NoError(t, err)
	require.Equal(t, int64(numQuestions), count)

	nonExistentCategory := "non_existent_category"
	nonExistentCount, err := testQueries.GetCategoryQuestionsCount(context.Background(), nonExistentCategory)
	require.NoError(t, err)
	require.Equal(t, int64(0), nonExistentCount)
}

func TestGetCategoriesCount(t *testing.T) {
	defer testQueries.CleanTables(context.Background(), []string{"categories"})

	const numCategories = 10

	for i := 0; i < numCategories; i++ {
		createRandomCategory(t)
	}

	count, err := testQueries.GetCategoriesCount(context.Background())
	require.NoError(t, err)
	require.Equal(t, int64(numCategories), count)
}
