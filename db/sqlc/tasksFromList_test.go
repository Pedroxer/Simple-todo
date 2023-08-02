package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTaskToList(t *testing.T) (List, Task) {
	list := createRandomList(t)
	var task Task
	for i := 0; i < 10; i++ {
		task = createRandomTask(t)
		arg := AddTaskToListParams{
			int32(task.ID),
			int32(list.ID),
		}
		tasksList, err := testQueries.AddTaskToList(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, tasksList)

		require.Equal(t, int32(task.ID), tasksList.TaskID)
		require.Equal(t, int32(list.ID), tasksList.ListID)
	}
	return list, task
}

func TestListAllTasks(t *testing.T) {
	list, _ := createRandomTaskToList(t)

	taskList, err := testQueries.ListAllTasks(context.Background(), int32(list.ID))

	require.NoError(t, err)
	require.NotEmpty(t, taskList)
	require.Len(t, taskList, 10)
}

func TestChangeListForTask(t *testing.T) {
	_, task := createRandomTaskToList(t)

	arg := ChangeListForTaskParams{
		ListID: 1,
		TaskID: int32(task.ID),
	}

	err := testQueries.ChangeListForTask(context.Background(), arg)
	require.NoError(t, err)

}
