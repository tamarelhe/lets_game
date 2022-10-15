package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		ID:       uuid.New(),
		Name:     "Test",
		Email:    "email@test.com",
		Password: "Test2",
		Avatar:   sql.NullString{},
		IsActive: true,
		Groups:   nil,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
