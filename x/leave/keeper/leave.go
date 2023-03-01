package keeper

import (
	//"cosmossdk.io/store/prefix"

	"leave-cosmos/x/leave/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetLeave(ctx sdk.Context, a types.Leave) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&a)
	if err != nil {
		return err
	}
	classID := a.LeaveId
	key := make([]byte, len(types.LeaveId)+len(classID))
	copy(key, types.LeaveId)
	copy(key[len(types.LeaveId):], []byte(classID))
	store.Set(key, bz)
	return nil
}

func (k Keeper) GetLeave(ctx sdk.Context, id string) (a types.Leave) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(id))
	classID := id
	key := make([]byte, len(types.PrefstudentId)+len(classID))
	copy(key, types.PrefstudentId)
	copy(key[len(types.PrefstudentId):], []byte(classID))
	store.Set(key, bz)

	k.cdc.MustUnmarshal(bz, &a)
	return a
}
