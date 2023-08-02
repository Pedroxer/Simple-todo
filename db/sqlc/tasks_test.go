package sqlc

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/Pedroxer/Simple-todo/util"
	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) Task {
	arg := CreateTaskParams{
		Name: util.RandomString(6),
		Description: sql.NullString{
			String: util.RandomString(20),
			Valid:  true,
		},
		Important: sql.NullInt32{
			Int32: int32(rand.Intn(1)),
			Valid: true,
		},
		Done: sql.NullInt32{
			Int32: int32(rand.Intn(1)),
			Valid: true,
		},
		Deadline: time.Now().Add(5 * time.Hour),
	}

	task, err := testQueries.CreateTask(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Name, task.Name)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Important, task.Important)
	require.Equal(t, arg.Done, task.Done)
	require.WithinDuration(t, arg.Deadline, task.Deadline, 181*time.Minute)
	require.NotZero(t, task.CreatedAt)
	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := createRandomTask(t)
	task, err := testQueries.GetTask(context.Background(), task1.Name)

	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, task1.Name, task.Name)
	require.Equal(t, task1.Description, task.Description)
	require.Equal(t, task1.Important, task.Important)
	require.Equal(t, task1.Done, task.Done)
	require.WithinDuration(t, task.Deadline, task1.Deadline, 3*time.Hour)
	require.NotZero(t, task1.CreatedAt)
}

func TestChangeDescription(t *testing.T) {
	task := createRandomTask(t)
	arg := ChangeDescriptionParams{
		Description: sql.NullString{
			String: util.RandomString(20),
			Valid:  true,
		},
		Name: task.Name,
	}
	err := testQueries.ChangeDescription(context.Background(), arg)
	require.NoError(t, err)

	task1, err := testQueries.GetTask(context.Background(), task.Name)
	require.NoError(t, err)
	require.Equal(t, arg.Description, task1.Description)

}

func TestChangeTaskDeadLine(t *testing.T) {
	task := createRandomTask(t)
	arg := ChangeTaskDeadlineParams{
		Deadline: util.RandomTime(3),
		Name:     task.Name,
	}
	err := testQueries.ChangeTaskDeadline(context.Background(), arg)
	require.NoError(t, err)

	task1, err := testQueries.GetTask(context.Background(), task.Name)
	require.NoError(t, err)
	require.WithinDuration(t, arg.Deadline, task1.Deadline, 181*time.Minute)

}

func TestChangeTaskDone(t *testing.T) {
	task := createRandomTask(t)
	arg := ChangeTaskDoneParams{
		Done: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Name: task.Name,
	}
	err := testQueries.ChangeTaskDone(context.Background(), arg)
	require.NoError(t, err)

	task1, err := testQueries.GetTask(context.Background(), task.Name)
	require.NoError(t, err)
	require.Equal(t, arg.Done, task1.Done)

}

func TestChangeTaskName(t *testing.T) {
	task := createRandomTask(t)
	arg := ChangeTaskNameParams{
		Name: util.RandomString(10),
		ID:   task.ID,
	}
	err := testQueries.ChangeTaskName(context.Background(), arg)
	require.NoError(t, err)

	task1, err := testQueries.GetTask(context.Background(), arg.Name)
	require.NoError(t, err)
	require.Equal(t, arg.Name, task1.Name)

}

func TestChangeTaskOrder(t *testing.T) {
	task := createRandomTask(t)
	arg := ChangeTaskOrderParams{
		Important: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Name: task.Name,
	}
	err := testQueries.ChangeTaskOrder(context.Background(), arg)
	require.NoError(t, err)

	task1, err := testQueries.GetTask(context.Background(), task.Name)
	require.NoError(t, err)
	require.Equal(t, arg.Important, task1.Important)

}

func TestDeleteTask(t *testing.T) {
	task := createRandomTask(t)
	err := testQueries.DeleteTask(context.Background(), task.Name)
	require.NoError(t, err)

	task_del, err := testQueries.GetTask(context.Background(), task.Name)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task_del)
}
