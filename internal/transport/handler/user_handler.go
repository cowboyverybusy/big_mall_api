package handler

import (
	"big_mall_api/internal/model"
	"big_mall_api/internal/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"strconv"
)

func (s *MallServerHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, 400, "Invalid request", err)
		return
	}

	user, err := s.logic.CreateUser(ctx.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(ctx, 500, "Failed to create user", err)
		return
	}

	utils.SuccessResponse(ctx, user)
}

func (s *MallServerHandler) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, 400, "Invalid user ID", err)
		return
	}

	user, err := s.logic.GetUser(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, 404, "User not found", err)
		return
	}

	utils.SuccessResponse(ctx, user)
}

func (s *MallServerHandler) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, 400, "Invalid user ID", err)
		return
	}

	if err := s.logic.DeleteUser(ctx.Request.Context(), uint(id)); err != nil {
		utils.ErrorResponse(ctx, 500, "Failed to delete user", err)
		return
	}

	utils.SuccessResponse(ctx, gin.H{"message": "User deleted successfully"})
}

func (s *MallServerHandler) ListUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	users, err := s.logic.ListUsers(ctx.Request.Context(), page, pageSize)
	if err != nil {
		utils.ErrorResponse(ctx, 500, "Failed to list users", err)
		return
	}

	utils.SuccessResponse(ctx, gin.H{
		"users":     users,
		"page":      page,
		"page_size": pageSize,
	})
}
