package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeAddAdmin    = "add-admin"
	TypeAddStudent  = "add-student"
	TypeApplyLeave  = "apply-leave"
	TypeAcceptLeave = "accept-leave"
)

var _ sdk.Msg = &AddAdminRequest{}
var _ sdk.Msg = &AddStudentRequest{}
var _ sdk.Msg = &ApplyLeaveRequest{}
var _ sdk.Msg = &AcceptLeaveRequest{}

func NewAddAdminReq(accAddr sdk.AccAddress) *AddAdminRequest {
	return &AddAdminRequest{
		AdminAddress: accAddr.String(),
	}

}
func (msg AddAdminRequest) Route() string {
	return RouterKey
}
func (msg AddAdminRequest) Type() string {
	return TypeAddAdmin
}
func (msg AddAdminRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)

}
func (msg AddAdminRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.AdminAddress) //ValAddressFromBech32(msg.AdminAddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AddAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////////////
func NewAddStudentReq(accountAddr sdk.AccAddress) *AddStudentRequest {
	return &AddStudentRequest{
		AdminAddress: accountAddr.String(),
	}

}
func (msg AddStudentRequest) Route() string {
	return RouterKey
}
func (msg AddStudentRequest) Type() string {
	return TypeAddStudent
}
func (msg AddStudentRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)

}
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.AdminAddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil

}

////////////////////////////////////////////////////////////////////////////////////
func NewApplyReq(accountAddr sdk.AccAddress) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Studentaddress: accountAddr.String(),
	}

}
func (msg ApplyLeaveRequest) Route() string {
	return RouterKey
}
func (msg ApplyLeaveRequest) Type() string {
	return TypeApplyLeave
}
func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)

}
func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Studentaddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Studentaddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil

}

///////////////////////////////////////////////////////////////////////////////////

func NewAcceptLeaveReq(accountAddr sdk.AccAddress) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{
		AdminAddress: accountAddr.String(),
	}

}
func (msg AcceptLeaveRequest) Route() string {
	return RouterKey
}
func (msg AcceptLeaveRequest) Type() string {
	return TypeAcceptLeave
}
func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)

}
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.AdminAddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil

}
