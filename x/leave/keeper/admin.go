package keeper

import (
	//"cosmossdk.io/store/prefix"

	"fmt"
	"leave-cosmos/x/leave/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetAdmin(ctx sdk.Context, a types.Admin) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&a)
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		return err
	} else {
		store.Set(types.AdminStoreKey(a.Address), bz)
		fmt.Println("cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc", a.Address)

	}
	return nil

}

func (k Keeper) GetAdmin(ctx sdk.Context, id string) (a types.Admin) {
	store := ctx.KVStore(k.storeKey)

	temp := store.Get(types.AdminStoreKey(id))
	//store.Has()
	k.cdc.Unmarshal(temp, &a)
	fmt.Printf("a: %v\n", a)
	return a

}
