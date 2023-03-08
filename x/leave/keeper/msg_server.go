package keeper

import (
	"context"
	"fmt"
	"leave-cosmos/x/leave/types"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.TxMsgServer = Keeper{} ///without & its not coming

// type msgServer struct {
// 	Keeper
// 	types.UnimplementedTxMsgServer
// }

//	func NewMsgServerImpl(k Keeper) types.TxMsgServer {
//		return &msgServer{
//			Keeper: k,
//		}
//	}
func (k Keeper) AddStudent(ctx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	fmt.Printf("\"==========================================================================\": %v\n", "==========================================================================")
	fmt.Printf("req: %v\n", req)
	// panic("Iam panicked")

	// panic(fmt.Sprintf("\"MSG Server\": %v\n", "MSG Server"))
	student := &types.Student{
		Id:      req.Id,
		Name:    req.Name,
		Address: req.Address,
	}
	if _, err := sdk.AccAddressFromBech32(req.AdminAddress); err != nil {
		log.Fatal("___here in admin register, error___", err)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	if err := k.SetStudent(sdkctx, student); err != nil {
		return &types.AddStudentResponse{}, err
	}

	return &types.AddStudentResponse{}, nil
}

func (k Keeper) RegisterAdmin(ctx context.Context, req *types.AddAdminRequest) (*types.AddAdminResponse, error) {
	admin := &types.Admin{
		Id:      req.Id,
		Address: req.AdminAddress,
		Name:    req.Name,
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	k.SetAdmin(sdkctx, *admin)

	return &types.AddAdminResponse{}, nil
}

func (k Keeper) ApplyLeave(ctx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	if _, err := sdk.AccAddressFromBech32(req.Studentaddress); err != nil {
		log.Fatal("not the right student", err)
	}
	fmt.Println("eeeeeeeeeeeeeeeeee")
	sdkctx := sdk.UnwrapSDKContext(ctx)
	// leave := &types.ApplyLeaveRequest{
	// To:   req.To,
	// From: req.From,
	//Id:   req.Studentaddress,
	// }
	if err := k.SetLeave(sdkctx, req); err != nil {
		return &types.ApplyLeaveResponse{}, nil
	}
	return &types.ApplyLeaveResponse{}, nil

}

func (k Keeper) AcceptLeave(ctx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	if _, err := sdk.AccAddressFromBech32(req.Leave.Adminaddress); err != nil {
		log.Fatal("not the right admin", err)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)

	k.AcceptLeaveReq(sdkctx, req.Leave.Studentaddress)

	return &types.AcceptLeaveResponse{}, nil
}
