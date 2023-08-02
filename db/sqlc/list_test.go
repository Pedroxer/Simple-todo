package sqlc

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/Pedroxer/Simple-todo/util"
)

func createRandomList(t *testing.T) List {
	arg := util.RandomString(10)
	list, err := testQueries.CreateList(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, arg, list.Title)
	return list
}

func TestCreateList(t *testing.T) {
	createRandomList(t)
}

func TestGetList(t *testing.T) {
	list := createRandomList(t)

	list1, err := testQueries.GetList(context.Background(), list.Title)

	require.NoError(t, err)
	require.NotEmpty(t, list1)
	require.Equal(t, list.Title, list1.Title)
	require.Equal(t, list.ID, list1.ID)
	require.WithinDuration(t, list.CreatedAt.Time, list1.CreatedAt.Time, time.Second)
}

func TestChangeListName(t *testing.T) {
	list := createRandomList(t)
	arg := ChangeListNameParams{
		Title: util.RandomString(6),
		ID:    list.ID,
	}

	err := testQueries.ChangeListName(context.Background(), arg)
	require.NoError(t, err)

	list1, err := testQueries.GetList(context.Background(), arg.Title)

	require.NoError(t, err)
	require.NotEmpty(t, list1)
	require.Equal(t, arg.Title, list1.Title)
}

func TestDeleteList(t *testing.T) {
	list := createRandomList(t)
	err := testQueries.DeleteList(context.Background(), list.Title)
	require.NoError(t, err)

	list_del, err := testQueries.GetList(context.Background(), list.Title)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, list_del)
}
