package handler

import "big_mall_api/internal/logic"

type MallServerHandler struct {
	logic *logic.ServerLogic
}

func NewMallServerHandler(serverLogic *logic.ServerLogic) *MallServerHandler {
	return &MallServerHandler{
		logic: serverLogic,
	}
}
