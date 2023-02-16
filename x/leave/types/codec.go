package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {

	legacy.RegisterAminoMsg(cdc, &AddAdminRequest{}, "leave-cosmos/CreateAdmin")
	legacy.RegisterAminoMsg(cdc, &AddStudentRequest{}, "leave-cosmos/CreateStudent")
	legacy.RegisterAminoMsg(cdc, &ApplyLeaveRequest{}, "leave-cosmos/ApplyLeave")
	legacy.RegisterAminoMsg(cdc, &AcceptLeaveRequest{}, "leave-cosmos/AcceptLeave")

	cdc.RegisterConcrete(&AddAdminRequest{}, "leave-cosmos/CreateAdmin", nil)
	cdc.RegisterConcrete(&AddStudentRequest{}, "leave-cosmos/CreateStudent", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "leave-cosmos/ApplyLeave", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "leave-cosmos/AcceptLeave", nil)

}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&AddAdminRequest{},
		&AddStudentRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_TxMsg_serviceDesc)

}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
