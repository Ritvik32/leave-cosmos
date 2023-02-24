package keeper

import (
	"context"
	"leave-cosmos/x/leave/types"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
func (k msgServer) AddStudent(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {

	if _, err := sdk.AccAddressFromBech32(req.AdminAddress); err != nil {
		log.Fatal("___here in admin register, error___", err)
	}
	return &types.AddStudentResponse{}, nil
}

func (k msgServer) AddAdmin(ctx context.Context, req *types.AddAdminRequest) (*types.AddAdminResponse, error) {
	admin := &types.Admin{
		Id:      req.Id,
		Address: req.AdminAddress,
		Name:    req.Name,
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	k.Keeper.SetAdmin(sdkctx, *admin)

	return &types.AddAdminResponse{}, nil
}

func (k msgServer) ApplyLeave(ctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	if _, err := sdk.AccAddressFromBech32(req.Studentaddress); err != nil {
		log.Fatal("not the right student", err)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leave := &types.Leave{
		To:      req.To,
		From:    req.From,
		Id:      req.Studentaddress,
		LeaveId: "012",
	}
	k.Keeper.SetLeave(sdkctx, *leave)
	return &types.ApplyLeaveResponse{}, nil

}

func (k msgServer) AcceptLeave(ctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	if _, err := sdk.AccAddressFromBech32(req.AdminAddress); err != nil {
		log.Fatal("not the right admin", err)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leave1 := &types.Leave{
		Id:      req.AdminAddress,
		LeaveId: "012",
		Status:  types.LeaveStatus_STATUS_ACCEPTED,
	}

	k.Keeper.SetLeave(sdkctx, *leave1)

	return &types.AcceptLeaveResponse{}, nil
}
