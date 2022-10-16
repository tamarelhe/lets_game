package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/tamarelhe/lets_game/db/sqlc"
	"github.com/tamarelhe/lets_game/util"
)

// create user request type
type createUserRequest struct {
	Name     string         `json:"name" binding:"required"`
	Email    string         `json:"email" binding:"required,email"`
	Password string         `json:"password" binding:"required,min=8"`
	Avatar   sql.NullString `json:"avatar"`
}

// create user request type
type userResponse struct {
	ID       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Avatar   sql.NullString `json:"avatar"`
	IsActive bool           `json:"is_active"`
	Groups   []string       `json:"groups"`
}

// create user method
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		ID:       util.RandomUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Avatar:   req.Avatar,
		IsActive: true,
		Groups:   nil,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := userResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		IsActive: user.IsActive,
		Groups:   user.Groups,
	}

	ctx.JSON(http.StatusOK, resp)
}

// get user request type
type getUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

// get user method
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userid, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, userid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := userResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		IsActive: user.IsActive,
		Groups:   user.Groups,
	}

	ctx.JSON(http.StatusOK, resp)
}

// get users list request type
type getUsersListRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=50"`
}

// get users list method
func (server *Server) getUsersList(ctx *gin.Context) {
	var req getUsersListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var respList []userResponse
	for _, user := range users {
		resp := userResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Avatar:   user.Avatar,
			IsActive: user.IsActive,
			Groups:   user.Groups,
		}

		respList = append(respList, resp)
	}

	ctx.JSON(http.StatusOK, respList)
}

// delete user request type
type deleteUserRequest struct {
	ID string `uri:"id" binding:"required"`
}

// delete user method
func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userid, err := uuid.Parse(req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteUser(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
