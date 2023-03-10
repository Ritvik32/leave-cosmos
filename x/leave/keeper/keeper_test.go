package keeper_test

import (
	"fmt"
	"sync"
	"testing"

	"leave-cosmos/x/leave/keeper"
	"leave-cosmos/x/leave/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

var Gaddress string = "cosmosv10222"

type TestSuite struct {
	suite.Suite
	ctx     sdk.Context
	skeeper keeper.Keeper
	*assert.Assertions
	mu      sync.RWMutex
	require *require.Assertions
	t       *testing.T
}

var expected error

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.skeeper = keeper
	s.ctx = ctx
}

func (s *TestSuite) TestSetStudent() {
	//require:=s.require()
	student := types.Student{
		Id:      "01",
		Name:    "Ritvik",
		Address: sdk.AccAddress("a111").String(),
	}
	//expected = types.ErrStudentDoesNotExist
	err := s.skeeper.SetStudent(s.ctx, &student)
	s.Require().Nil(err)
}

func (s *TestSuite) TestGetStudent() {
	student := types.Student{
		Id:      "01",
		Name:    "Ritvik",
		Address: "a111",
	}
	//expected = types.ErrStudentDoesNotExist
	err := s.skeeper.SetStudent(s.ctx, &student)
	s.Require().Nil(err)

	stud := s.skeeper.GetStudent(s.ctx, student.Id)
	fmt.Printf("stud: %v\n", stud)
	s.Require().Equal(stud.Id, student.Id)
	//s.Require().Nil(err)

}
func (s *TestSuite) TestApplyLeave() {

	leave := types.ApplyLeaveRequest{
		//Id:      "01",
		From: "3-7-2023",
		To:   "5-7-2023",
		//LeaveId: "l1",
		//Status:  types.LeaveStatus_STATUS_UNDEFINED,
		Studentaddress: "aaaaaaaaaaaaaa",
		Reason:         "cold",
	}
	res := s.skeeper.SetLeave(s.ctx, &leave)
	s.Require().Nil(res)

}

func (s *TestSuite) TestAcceptleave() {
	leave1 := types.Leave{

		Studentaddress: "aaaaaaaaaaa",
		Reason:         "cold",
		From:           "3-7-2023",
		To:             "5-7-2023",
		Adminaddress:   "comos1",
		Status:         types.LeaveStatus_STATUS_ACCEPTED,
	}
	// Leave: &types.Leave{Stude "01",
	// From:    "3-7-2023",
	// To:      "5-7-2023",
	// LeaveId: "l1",
	// Status:  types.LeaveStatus_STATUS_ACCEPTED},

	res := s.skeeper.AcceptLeaveReq(s.ctx, leave1.Studentaddress)
	fmt.Println(res)
	s.Require().Nil(res)

}

func (s *TestSuite) TestRegisterAdmin() {
	admin := types.Admin{
		Id:      "a1",
		Address: "cosmosv10222",
	}
	//expected = types.ErrAdminDoesNotExist
	err := s.skeeper.SetAdmin(s.ctx, admin)
	s.Require().Nil(err)

}
func (s *TestSuite) TestGetAdmin() {
	admin1 := types.Admin{
		Id:      "a1",
		Address: "cosmosv10222",
	}
	//expected = types.ErrAdminDoesNotExist
	err := s.skeeper.SetAdmin(s.ctx, admin1)
	s.Require().Nil(err)

	stud := s.skeeper.GetAdmin(s.ctx, admin1.Id)
	s.Require().Equal(admin1, stud)

}

func (s *TestSuite) TestGetLeave() {
	// leave := types.Leave{
	// Id:      "a1",
	// From:    "10-3-23",
	// To:      "14-3-23",
	// LeaveId: "L01",
	// Status:  types.LeaveStatus_STATUS_ACCEPTED,
	// }
	err := s.skeeper.SetLeave(s.ctx, &types.ApplyLeaveRequest{
		Studentaddress: "1111",
		Reason:         "cold",
		From:           "1-1-1",
		To:             "3-1-1",
		Adminaddress:   "cosmos1",
	})
	s.Require().NoError(err)
	res := s.skeeper.GetLeave(s.ctx, "1111")
	fmt.Printf("res: %v\n", res)
	leave1 := s.skeeper.GetLeave(s.ctx, "1111")
	fmt.Printf("leave1: %v\n", leave1)
	//fmt.Println(leave)
	fmt.Println(leave1)
	s.Require().Equal(res, leave1)
}
