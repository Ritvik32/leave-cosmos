package keeper

import (
	//"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"context"
	"leave-cosmos/x/leave/types"
)

var _ types.TxMsgServer = &msgServer{} ///without & its not coming

type msgServer struct {
	Keeper
	types.UnimplementedTxMsgServer
}

func NewMsgServerImpl(k Keeper) types.TxMsgServer {
	return &msgServer{
		Keeper: k,
	}
}
func (k msgServer) AddStudent(context.Context, *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	return &types.AddStudentResponse{}, nil
}

func (k msgServer) AddAdmin(context.Context, *types.AddAdminRequest) (*types.AddAdminResponse, error) {
	return &types.AddAdminResponse{}, nil
}

func (k msgServer) ApplyLeave(context.Context, *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	return &types.ApplyLeaveResponse{}, nil
}

func (k msgServer) AcceptLeave(context.Context, *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}
