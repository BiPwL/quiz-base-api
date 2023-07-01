package db

import (
	"context"
	"database/sql"
	"github.com/BiPwL/quiz-base-api/util"
	"github.com/stretchr/testify/require"
	"testing"
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
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.Key)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.Key, category2.Key)
	require.Equal(t, category1.Name, category2.Name)
}

func TestUpdateCategory(t *testing.T) {
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
	category1 := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category1.Key)
	require.NoError(t, err)

	category2, err := testQueries.GetCategory(context.Background(), category1.Key)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category2)
}

func TestListCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}
	arg := ListCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	categories, err := testQueries.ListCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, categories, 5)

	for _, category := range categories {
		require.NotEmpty(t, category)
	}
}
