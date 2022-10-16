package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tamarelhe/lets_game/util"
)

func createRandomUser(t *testing.T, isActive bool) LgUser {
	arg := CreateUserParams{
		ID:       util.RandomUUID(),
		Name:     util.RandomString(30),
		Email:    util.RandomString(20),
		Password: util.RandomString(10),
		Avatar:   sql.NullString{},
		IsActive: isActive,
		Groups:   nil,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Avatar, user.Avatar)
	require.Equal(t, arg.IsActive, user.IsActive)
	require.Equal(t, arg.Groups, user.Groups)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t, true)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t, true)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Avatar, user2.Avatar)
	require.Equal(t, user1.IsActive, user2.IsActive)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.Groups, user2.Groups)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t, true)

	arg := UpdateUserParams{
		ID:     user1.ID,
		Name:   util.RandomString(30),
		Email:  util.RandomString(20),
		Avatar: sql.NullString{},
		Groups: nil,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Avatar, user2.Avatar)
	require.Equal(t, user1.IsActive, user2.IsActive)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.Groups, user2.Groups)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t, true)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t, true)
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
}

func TestActivateUser(t *testing.T) {
	user1 := createRandomUser(t, false)

	err := testQueries.ActivateUser(context.Background(), user1.ID)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEqual(t, user1.IsActive, user2.IsActive)
	require.Equal(t, user2.IsActive, true)
}

func TestInactivateUser(t *testing.T) {
	user1 := createRandomUser(t, true)

	err := testQueries.InactivateUser(context.Background(), user1.ID)
	require.NoError(t, err)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEqual(t, user1.IsActive, user2.IsActive)
	require.Equal(t, user2.IsActive, false)
}
