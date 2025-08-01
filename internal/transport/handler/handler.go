package handler

import (
	"big_mall_api/internal/logic"
	"big_mall_api/pkg/storage"
)

type MallServerHandler struct {
	logic *logic.ServerLogic
}

func NewMallServerHandler(dbMgr *storage.DbManager) *MallServerHandler {
	serverLogic := logic.NewServerLogic(dbMgr)
	return &MallServerHandler{
		logic: serverLogic,
	}
}
