package keeper

import (
	//"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"context"
	"leave-cosmos/x/leave/types"
)

var _ types.QueryMsgServer = &querySever{}

type querySever struct {
	Keeper
	types.UnimplementedQueryMsgServer
}

func NewQueryServerImpl(k Keeper) types.QueryMsgServer {
	return &querySever{
		Keeper: k,
	}
}

func (k querySever) GetStudent(context.Context, *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	return &types.GetStudentResponse{}, nil
}
func (k querySever) GetAllLeaves(context.Context, *types.ListAllLeavesRequest) (*types.ListAllLeavesResponse, error) {
	return &types.ListAllLeavesResponse{}, nil
}
func (k querySever) GetStudentLeave(context.Context, *types.ListStudentLeaveRequest) (*types.ListStudentLeaveResponse, error) {
	return &types.ListStudentLeaveResponse{}, nil
}
