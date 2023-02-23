package types

/*import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/kv"
)*/

const (
	// ModuleName is the name of the module
	ModuleName = "leave"

	// StoreKey is the store key string for slashing
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute is the querier route for slashing
	QuerierRoute = ModuleName
)

var (
	PrefstudentId = []byte{0x01}
	LeaveId       = []byte{0x02}
	AdminId       = []byte{0x03}
)
