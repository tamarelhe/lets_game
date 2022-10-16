package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tamarelhe/lets_game/util"
)

func createRandomMembers(n int) []byte {
	var membersJson []map[string]interface{}

	for i := 0; i < n; i++ {
		member := map[string]interface{}{
			"id":   util.RandomUUID(),
			"role": util.RandomInt(1, 3),
		}
		membersJson = append(membersJson, member)
	}

	jsonData, err := json.Marshal(membersJson)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return nil
	}

	return jsonData
}

func createRandomGroup(t *testing.T) LgGroup {
	membersJson := createRandomMembers(2)

	arg := CreateGroupParams{
		ID:      util.RandomUUID(),
		Name:    util.RandomString(30),
		Avatar:  sql.NullString{},
		Members: membersJson,
	}

	group, err := testQueries.CreateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group)

	require.Equal(t, arg.ID, group.ID)
	require.Equal(t, arg.Name, group.Name)
	require.Equal(t, arg.Avatar, group.Avatar)

	require.JSONEq(t, string(membersJson), string(group.Members))

	require.NotZero(t, group.ID)
	require.NotZero(t, group.CreatedAt)

	return group
}

func TestCreateGroup(t *testing.T) {
	createRandomGroup(t)
}

func TestGetGroup(t *testing.T) {
	group1 := createRandomGroup(t)
	group2, err := testQueries.GetGroup(context.Background(), group1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, group2)

	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, group1.Name, group2.Name)
	require.Equal(t, group1.Avatar, group2.Avatar)
	require.Equal(t, group1.Members, group2.Members)
	require.Equal(t, group1.CreatedAt, group2.CreatedAt)
}

func TestUpdateGroup(t *testing.T) {
	group1 := createRandomGroup(t)

	newMembersJson := createRandomMembers(3)

	arg := UpdateGroupParams{
		ID:      group1.ID,
		Name:    util.RandomString(30),
		Avatar:  sql.NullString{},
		Members: newMembersJson,
	}

	group2, err := testQueries.UpdateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group2)

	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, arg.Name, group2.Name)
	require.Equal(t, arg.Avatar, group2.Avatar)
	require.JSONEq(t, string(arg.Members), string(group2.Members))
	require.Equal(t, group1.CreatedAt, group2.CreatedAt)
}

func TestDeleteGroup(t *testing.T) {
	group1 := createRandomGroup(t)
	err := testQueries.DeleteGroup(context.Background(), group1.ID)
	require.NoError(t, err)

	group2, err := testQueries.GetGroup(context.Background(), group1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, group2)
}

func TestListGroups(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomGroup(t)
	}

	arg := ListGroupsParams{
		Limit:  5,
		Offset: 5,
	}

	groups, err := testQueries.ListGroups(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, groups, 5)

	for _, group := range groups {
		require.NotEmpty(t, group)
	}
}
