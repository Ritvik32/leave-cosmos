package keeper

import (
	//"cosmossdk.io/store/prefix"

	"leave-cosmos/x/leave/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetAdmin(ctx sdk.Context, a types.Admin) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&a)
	if err != nil {
		return err
	}
	classID := a.Id
	key := make([]byte, len(types.LeaveId)+len(classID))
	copy(key, types.LeaveId)
	copy(key[len(types.LeaveId):], []byte(classID))
	store.Set(key, bz)
	return nil

}

func (k Keeper) GetAdmin(ctx sdk.Context, b types.Admin) (a types.Admin) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(b.Id))
	classID := b.Id
	key := make([]byte, len(types.PrefstudentId)+len(classID))
	copy(key, types.PrefstudentId)
	copy(key[len(types.PrefstudentId):], []byte(classID))
	store.Set(key, bz)

	k.cdc.MustUnmarshal(bz, &a)
	return a
}
