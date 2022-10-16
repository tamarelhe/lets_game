package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tabbed/pqtype"
	"github.com/tamarelhe/lets_game/util"
)

func createRandomGameConstraints() pqtype.NullRawMessage {
	constraints := map[string]interface{}{
		"min_players": util.RandomInt(2, 4),
		"max_players": util.RandomInt(2, 10),
	}

	jsonData, err := json.Marshal(constraints)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return pqtype.NullRawMessage{}
	}

	return pqtype.NullRawMessage{
		RawMessage: jsonData,
		Valid:      true,
	}
}

func createRandomGameLocation() pqtype.NullRawMessage {
	location := map[string]interface{}{
		"latitude":  util.RandomInt(2, 4),
		"longitude": util.RandomInt(2, 10),
	}

	jsonData, err := json.Marshal(location)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return pqtype.NullRawMessage{}
	}

	return pqtype.NullRawMessage{
		RawMessage: jsonData,
		Valid:      true,
	}
}

func createRandomGame(t *testing.T) LgGame {
	group := createRandomGroup(t)

	membersJson := getGroupMembersJson(group)
	locationJson := createRandomGameLocation()
	constraintsJson := createRandomGameConstraints()

	arg := CreateGameParams{
		ID:       util.RandomUUID(),
		GroupID:  group.ID,
		TypeID:   util.RandomNullUUID(),
		Datetime: time.Now(),
		Members: pqtype.NullRawMessage{
			RawMessage: membersJson,
			Valid:      true,
		},
		Location:    locationJson,
		Constraints: constraintsJson,
		Message: sql.NullString{
			String: util.RandomString(60),
			Valid:  true,
		},
	}

	game, err := testQueries.CreateGame(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.Equal(t, arg.ID, game.ID)
	require.Equal(t, arg.GroupID, game.GroupID)
	require.Equal(t, arg.TypeID, game.TypeID)
	//require.Equal(t, arg.Datetime.UTC(), game.Datetime.UTC())
	require.WithinDuration(t, arg.Datetime.UTC(), game.Datetime.UTC(), time.Second)

	require.JSONEq(t, string(membersJson), string(game.Members.RawMessage))
	require.JSONEq(t, string(locationJson.RawMessage), string(game.Location.RawMessage))
	require.JSONEq(t, string(constraintsJson.RawMessage), string(game.Constraints.RawMessage))

	require.Equal(t, arg.Message, game.Message)

	require.NotZero(t, game.ID)
	require.NotZero(t, game.CreatedAt)

	return game
}

func TestCreateGame(t *testing.T) {
	createRandomGame(t)
}

func TestGetGame(t *testing.T) {
	game1 := createRandomGame(t)
	game2, err := testQueries.GetGame(context.Background(), game1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, game2)

	require.Equal(t, game1.ID, game2.ID)
	require.Equal(t, game1.GroupID, game2.GroupID)
	require.Equal(t, game1.TypeID, game2.TypeID)
	//require.Equal(t, game1.Datetime.UTC(), game2.Datetime.UTC())
	require.WithinDuration(t, game1.Datetime.UTC(), game2.Datetime.UTC(), time.Second)
	require.Equal(t, game1.Members.RawMessage, game2.Members.RawMessage)
	require.Equal(t, game1.Location.RawMessage, game2.Location.RawMessage)
	require.Equal(t, game1.Constraints.RawMessage, game2.Constraints.RawMessage)
	require.Equal(t, game1.Message, game2.Message)
	require.Equal(t, game1.CreatedAt, game2.CreatedAt)
}

func TestUpdateGame(t *testing.T) {
	game1 := createRandomGame(t)

	group := createRandomGroup(t)
	membersJson := getGroupMembersJson(group)
	locationJson := createRandomGameLocation()
	constraintsJson := createRandomGameConstraints()

	arg := UpdateGameParams{
		ID:       game1.ID,
		GroupID:  group.ID,
		TypeID:   util.RandomNullUUID(),
		Datetime: time.Now(),
		Members: pqtype.NullRawMessage{
			RawMessage: membersJson,
			Valid:      true,
		},
		Location:    locationJson,
		Constraints: constraintsJson,
		Message: sql.NullString{
			String: util.RandomString(60),
			Valid:  true,
		},
	}

	game2, err := testQueries.UpdateGame(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, game2)

	require.Equal(t, game1.ID, game2.ID)
	require.Equal(t, arg.GroupID, game2.GroupID)
	require.Equal(t, arg.TypeID, game2.TypeID)
	//require.Equal(t, arg.Datetime.UTC(), game2.Datetime.UTC())
	require.WithinDuration(t, arg.Datetime.UTC(), game2.Datetime.UTC(), time.Second)
	require.JSONEq(t, string(arg.Members.RawMessage), string(game2.Members.RawMessage))
	require.JSONEq(t, string(arg.Location.RawMessage), string(game2.Location.RawMessage))
	require.JSONEq(t, string(arg.Constraints.RawMessage), string(game2.Constraints.RawMessage))
	require.Equal(t, arg.Message, game2.Message)
}

func TestDeleteGame(t *testing.T) {
	game1 := createRandomGame(t)
	err := testQueries.DeleteGame(context.Background(), game1.ID)
	require.NoError(t, err)

	game2, err := testQueries.GetGame(context.Background(), game1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, game2)
}

func TestListGames(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomGame(t)
	}

	arg := ListGamesParams{
		Limit:  5,
		Offset: 5,
	}

	games, err := testQueries.ListGames(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, games, 5)

	for _, game := range games {
		require.NotEmpty(t, game)
	}
}
