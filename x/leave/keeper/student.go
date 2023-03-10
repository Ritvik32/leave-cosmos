package keeper

import (
	//"cosmossdk.io/store/prefix"

	"fmt"
	"leave-cosmos/x/leave/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetStudent(ctx sdk.Context, a *types.Student) error {
	// panic(fmt.Sprintf("aQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQQq: %v\n", a))
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(a)
	if err != nil {
		return err
	} else {
		store.Set(types.StudentStoreKey(a.Address), bz)
	}
	return nil

}
func (k Keeper) GetStudent(ctx sdk.Context, id string) (a types.Student) {
	store := ctx.KVStore(k.storeKey)
	temp := store.Get(types.StudentStoreKey(id))
	k.cdc.MustUnmarshal(temp, &a)
	fmt.Printf("a!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!: %v\n", a)
	return a
}

func (k Keeper) GetStudents(ctx sdk.Context) []*types.Student {
	store := ctx.KVStore(k.storeKey)
	var students []*types.Student
	fmt.Println("11111111111111111111")
	itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	for ; itr.Valid(); itr.Next() {
		var t types.Student
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)
	}
	fmt.Println("222222222222222222222")
	return students
}
