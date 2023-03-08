package keeper

import (
	//"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"context"
	"fmt"
	"leave-cosmos/x/leave/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryMsgServer = Keeper{}

// type queryServer struct {
// 	Keeper
// 	types.UnimplementedQueryMsgServer
// }

// func NewQueryServerImpl(k Keeper) types.QueryMsgServer {
// 	return &queryServer{
// 		Keeper: k,
// 	}
// }

func (k Keeper) GetStudentDetails(ctx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	//return &types.GetStudentResponse{}, nil
	sdkctx := sdk.UnwrapSDKContext(ctx)
	a := k.GetStudent(sdkctx, req.Id)
	fmt.Printf("aaaaaaaaaaaaa: %v\n", a)
	return &types.GetStudentResponse{Student: &a}, nil
}
func (k Keeper) ListAllLeave(ctx context.Context, a *types.ListAllLeavesRequest) (*types.ListAllLeavesResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leaves := k.GetStudentLeaves(sdkctx, a.StudentId)
	return &types.ListAllLeavesResponse{Leave: leaves}, nil

	//return &types.ListAllLeavesResponse{}, nil
}

// func (k Keeper) ListStudentLeave(ctx context.Context, req *types.ListStudentLeaveRequest) (*types.ListStudentLeaveResponse, error) {
// sdkctx := sdk.UnwrapSDKContext(ctx)
// k.GetStudent(sdkctx, req.Id, req.Leaveid)
// a := k.GetLeave(sdkctx, req.Stuaddress)
// var temp =[] types.Leave{}
// return &types.ListStudentLeaveResponse{Leave: a}, nil
// }
func (k Keeper) ListStudentLeave(ctx context.Context, req *types.ListStudentLeaveRequest) (*types.ListStudentLeaveResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	//k.GetStudent(sdkctx, req.Id, req.Leaveid)
	a := k.GetLeave(sdkctx, req.Stuaddress)
	//var temp =[] types.Leave{}
	return &types.ListStudentLeaveResponse{Leave: &a}, nil
}

// func (k Keeper) ListStudents(ctx context.Context,req *types.GetStudentRequest)[]*types.Student{
// sdkctx:=sdk.UnwrapSDKContext(ctx)
// store := ctx.KVStore(k.storeKey)
// var students []*types.Student
// itr:=sdk.KVStorePrefixIterator(store)
// }

func (k Keeper) GetAllStudentDetails(ctx context.Context, _ *types.GetStudentRequest) (*types.GetAllStudentResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	students := k.GetStudents(sdkctx)
	return &types.GetAllStudentResponse{Student: students}, nil
}

// func (k Keeper) LeaveStatus(ctx context.Context, a *types.ApplyLeaveRequest)string {
// sdkctx := sdk.UnwrapSDKContext(ctx)
//
// }

func (k Keeper) GetLeaveStatus(ctx context.Context, a *types.LeaveStatusRequest) (*types.LeaveStatusResponse, error) {
	// store := ctx.KVStore(k.storeKey)
	// temp := store.Get(types.LeaveStoreKey(a.Studentaddress))
	// var leave types.AcceptLeaveRequest
	// k.cdc.Unmarshal(temp, &leave)
	// status := leave.Leave.Status.String()
	// return status
	sdkctx := sdk.UnwrapSDKContext(ctx)
	leaveStatus := k.CheckLeaveStatus(sdkctx, a.StudentAddress)
	return &types.LeaveStatusResponse{Status: leaveStatus.Status}, nil
}

func (k Keeper) GetAdminDetails(ctx context.Context, a *types.GetAdminRequest) (*types.GetAdminResponse, error) {
	sdkctx := sdk.UnwrapSDKContext(ctx)
	admin := k.GetAdmin(sdkctx, a.Adminaddress)
	return &types.GetAdminResponse{Admin: &admin}, nil
}
