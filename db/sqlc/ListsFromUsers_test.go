package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomListsForUsers(t *testing.T) (User, List) {
	user := createRandomUser(t)
	var list List
	for i := 0; i < 10; i++ {
		list = createRandomList(t)
		arg := AddListToUserParams{
			UserID: int32(user.ID),
			ListID: int32(list.ID),
		}
		listsToUser, err := testQueries.AddListToUser(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, listsToUser)

		require.Equal(t, arg.ListID, listsToUser.ListID)
		require.Equal(t, arg.UserID, listsToUser.UserID)
	}
	return user, list
}

func TestAddListToUser(t *testing.T) {
	createRandomListsForUsers(t)
}
func TestListAllUserLists(t *testing.T) {
	user, _ := createRandomListsForUsers(t)

	userLists, err := testQueries.ListAllUserLists(context.Background(), int32(user.ID))

	require.NoError(t, err)
	require.NotEmpty(t, userLists)
	require.Len(t, userLists, 10)
	for i := 0; i < 10; i++ {
		require.NotEmpty(t, userLists[i])
	}
}

func TestDeleteListFromUser(t *testing.T) {
	user, list := createRandomListsForUsers(t)
	arg := DeleteListFromUserParams{
		ListID: int32(list.ID),
		UserID: int32(user.ID),
	}
	err := testQueries.DeleteListFromUser(context.Background(), arg)
	require.NoError(t, err)

}
