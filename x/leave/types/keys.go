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
	StudentKey      = []byte{0x01}
	LeaveKey        = []byte{0x02}
	AdminKey        = []byte{0x03}
	AcceptKey       = []byte{0x04}
	StudentLeaveKey = []byte{0x05}
	CounterKey      = []byte{0x06}
)

func StudentLeaveStoreKey(sid, lid string) []byte {
	key := make([]byte, 3+len(sid)+len(lid), 0)
	key = append(key, StudentLeaveKey...)
	key = append(key, byte(len(sid)))
	key = append(key, []byte(sid)...)
	key = append(key, byte(len(lid)))
	key = append(key, []byte(lid)...)
	// copy(key,StudentLeaveKey)
	// copy(key[len(StudentLeaveKey):],len(sid),byte(sid))
	return key
}
func StudentStoreKey(student string) []byte {
	key := make([]byte, len(StudentKey)+len(student))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], []byte(student))
	return key
}

func LeaveStoreKey(sid string) []byte {
	key := make([]byte, len(LeaveKey)+len(sid))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(sid))
	return key
}
func AdminStoreKey(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

func AcceptLeaveStoreKey(leaveid string) []byte {
	key := make([]byte, len(AcceptKey)+len(leaveid))
	copy(key, leaveid)
	copy(key[len(AcceptKey):], []byte(leaveid))
	return key
}
func StudentLeavesCounterKey(sid string) []byte {
	key := make([]byte, len(CounterKey)+len(sid))
	copy(key, CounterKey)
	copy(key[len(CounterKey):], []byte(sid))
	return key
}
