// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type LgGame struct {
	ID          uuid.UUID             `json:"id"`
	GroupID     uuid.UUID             `json:"group_id"`
	TypeID      uuid.NullUUID         `json:"type_id"`
	Datetime    time.Time             `json:"datetime"`
	Members     pqtype.NullRawMessage `json:"members"`
	Location    pqtype.NullRawMessage `json:"location"`
	Constraints pqtype.NullRawMessage `json:"constraints"`
	Message     sql.NullString        `json:"message"`
	CreatedAt   interface{}           `json:"created_at"`
}

type LgGroup struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Avatar    sql.NullString  `json:"avatar"`
	Members   json.RawMessage `json:"members"`
	CreatedAt time.Time       `json:"created_at"`
}

type LgUser struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Avatar    sql.NullString `json:"avatar"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	Groups    []string       `json:"groups"`
}
