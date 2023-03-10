package keeper

import (
	//"cosmossdk.io/store/prefix"
	"fmt"
	"leave-cosmos/x/leave/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLeave(ctx sdk.Context, a *types.ApplyLeaveRequest) error {
	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR")
	store := ctx.KVStore(k.storeKey)

	leave := types.Leave{
		Studentaddress: a.Studentaddress,
		Reason:         a.Reason,
		From:           a.From,
		Adminaddress:   a.Adminaddress,
		To:             a.To,
		Status:         types.LeaveStatus_STATUS_UNDEFINED,
	}
	bz, err := k.cdc.Marshal(&leave)
	if err != nil {
		return err
	}

	store.Set(types.LeaveStoreKey(leave.Studentaddress), bz)
	fmt.Println(leave.Studentaddress)
	//store.Set(types.StudentLeaveStoreKey(s))
	// id := types.LeaveStoreKey(applyLeave.Address, string(counter))
	//return "Leave Applied Successfully"
	return nil
}
func (k Keeper) GetLeave(ctx sdk.Context, id string) (a types.Leave) {
	store := ctx.KVStore(k.storeKey)
	temp := store.Get(types.LeaveStoreKey(id))
	k.cdc.Unmarshal(temp, &a)
	fmt.Printf("a: %v\n", a)
	return a
}

func (k Keeper) AcceptLeaveReq(ctx sdk.Context, studentaddr string) error {
	store := ctx.KVStore(k.storeKey)
	// bz, err := k.cdc.Marshal(a)
	// if err != nil {
	// 	return err
	// }
	//store.Set(types.LeaveStoreKey(a.LeaveId), bz)
	//addr := types.StudentLeavesCounterKey(applyLeave.Address)
	// addr := types.StudentLeavesCounterKey(sdk.AccAddress(string(applyLeave.Address)).String())
	// counter := store.Get(addr)

	//counter = store.Get(addr)
	//handleError(err)
	byteArray := store.Get(types.LeaveStoreKey(studentaddr))
	var leave types.Leave
	k.cdc.Unmarshal(byteArray, &leave)
	leave.Status = types.LeaveStatus_STATUS_ACCEPTED
	byteArray, _ = k.cdc.Marshal(&leave)

	store.Set(types.LeaveStoreKey(studentaddr), byteArray)
	// id := types.LeaveStoreKey(applyLeave.Address, string(counter))
	//return "Leave Applied Successfully"
	return nil
}

// func (k Keeper) GetLeave(ctx sdk.Context, id string) (a types.ApplyLeaveResponse) {
// store := ctx.KVStore(k.storeKey)
// temp := store.Get(types.LeaveStoreKey(id))
// k.cdc.Unmarshal(temp, &a)
// fmt.Printf("a: %v\n", a)
// return a
// }

func (k Keeper) GetStudentLeaves(ctx sdk.Context, sid string) []*types.Leave {
	store := ctx.KVStore(k.storeKey)
	var leaves []*types.Leave

	key := make([]byte, 2+len(sid), 0)
	key = append(key, types.StudentLeaveKey...)
	key = append(key, byte(len(sid)))
	key = append(key, []byte(sid)...)
	// key = append(key, []byte(lid)...)

	itr := sdk.KVStorePrefixIterator(store, key)
	for ; itr.Valid(); itr.Next() {
		var t types.Leave
		k.cdc.Unmarshal(itr.Value(), &t)
		leaves = append(leaves, &t)
	}
	return leaves

}

// func (k Keeper) SetLeave(ctx sdk.Context, a *types.ApplyLeaveRequest) error {
// fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR")
// store := ctx.KVStore(k.storeKey)
// bz, err := k.cdc.Marshal(a)
// if err != nil {
// return err
// }
// leaveid := store.Get(types.StudentLeavesCounterKey(a.Studentaddress))
// leaveId, err := strconv.Atoi(string(leaveid))
// if err != nil {
// fmt.Println(err)
// }
// if leaveid == nil {
// leaveId = 0
// }
//
// leaveId++
// store.Set(types.StudentLeavesCounterKey(a.Studentaddress), []byte(strconv.Itoa(leaveId)))
// store.Set(types.StudentLeaveStoreKey(a.Studentaddress, string(leaveid)), bz)
func (k Keeper) CheckLeaveStatus(ctx sdk.Context, studentAddress string) (a types.Leave) {
	store := ctx.KVStore(k.storeKey)
	temp := store.Get(types.LeaveStoreKey(studentAddress))
	//var leave types.AcceptLeaveRequest
	k.cdc.Unmarshal(temp, &a)
	status1 := a.Status
	return types.Leave{Status: status1}

}

// func (k Keeper) GetLeave(ctx sdk.Context, id string) (a types.Leave) {
// store := ctx.KVStore(k.storeKey)
// temp := store.Get(types.LeaveStoreKey(id))
// k.cdc.Unmarshal(temp, &a)
// fmt.Printf("a: %v\n", a)
// return a
// }
