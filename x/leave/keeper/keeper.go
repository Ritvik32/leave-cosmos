package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	// "github.com/cosmos/gogoproto/codec"
	//storetypes "github.com/cosmos/cosmos-sdk/store/types"
	storetypes "cosmossdk.io/store/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey

	cdc codec.BinaryCodec
}

func NewKeeper(
	key storetypes.StoreKey,
	cdc codec.BinaryCodec,
	//ak leave.AccountKeeper,

) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: key,
	}
}
