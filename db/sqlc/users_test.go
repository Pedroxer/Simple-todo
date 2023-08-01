package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Pedroxer/Simple-todo/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomString(10),
		Password: util.RandomPassword(6),
		Email:    util.RandomEmail(5),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user1.Username, user.Username)
	require.Equal(t, user1.Password, user.Password)
	require.Equal(t, user1.Email, user.Email)
	require.WithinDuration(t, user1.CreatedAt.Time, user.CreatedAt.Time, time.Millisecond)
}

func TestChangePassword(t *testing.T){
	user := createRandomUser(t)

	arg := ChangePasswordParams{
		Password:util.RandomPassword(6),
		Username: user.Username,
	}
	err := testQueries.ChangePassword(context.Background(), arg)

	require.NoError(t, err)

	user1, err := testQueries.GetUser(context.Background(), user.Username)
	
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.Username, user.Username)
	require.Equal(t, user1.Password, arg.Password)
	require.Equal(t, user1.Email, user.Email)
	require.WithinDuration(t, user1.CreatedAt.Time, user.CreatedAt.Time, time.Millisecond)

}

func TestChangeEmail(t *testing.T){
	user := createRandomUser(t)

	arg := ChangeEmailParams{
		Email:util.RandomEmail(6),
		Username: user.Username,
	}
	err := testQueries.ChangeEmail(context.Background(), arg)

	require.NoError(t, err)

	user1, err := testQueries.GetUser(context.Background(), user.Username)
	
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.Username, user.Username)
	require.Equal(t, user1.Email, arg.Email)
	require.Equal(t, user1.Password, user.Password)
	require.WithinDuration(t, user1.CreatedAt.Time, user.CreatedAt.Time, time.Millisecond)

}
func TestDeleteUser(t *testing.T){
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.Username)

	require.NoError(t, err)

	del_user, err := testQueries.GetUser(context.Background(), user1.Username)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, del_user)
}