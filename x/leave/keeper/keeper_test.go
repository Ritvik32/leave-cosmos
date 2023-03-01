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
func (suite *TestSuite) T() *testing.T {
	suite.mu.RLock()
	defer suite.mu.RUnlock()
	return suite.t
}

// SetT sets the current *testing.T context.
func (suite *TestSuite) SetT(t *testing.T) {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	suite.t = t
	suite.Assertions = assert.New(t)
	suite.require = require.New(t)
}

// Require returns a require context for suite.
func (suite *TestSuite) Require() *require.Assertions {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	if suite.require == nil {
		suite.require = require.New(suite.T())
	}
	return suite.require
}
func (s *TestSuite) TestSetStudent() {
	student := types.Student{
		Id:      "01",
		Name:    "Ritvik",
		Address: sdk.AccAddress("a111").String(),
	}
	// r := types.AddStudentRequest{
	// Id:           "01",
	// Name:         "Ritvik",
	// Address:      sdk.AccAddress("a111").String(),
	// AdminAddress: "abcdef",
	// }
	res := s.skeeper.SetStudent(s.ctx, student)
	fmt.Println(res)
}

// func (s *TestSuite) TestGetStudent() {
// 
// }
func (s *TestSuite) ApplyLeave() {

	leave := types.Leave{
		Id:      "01",
		From:    "3-7-2023",
		To:      "5-7-2023",
		LeaveId: "l1",
	}
	res := s.skeeper.SetLeave(s.ctx, leave)
	fmt.Println(res)
}

func (s *TestSuite) RegisterAdmin() {
	admin := types.Admin{
		Id:      "a1",
		Address: "cosmosv10222",
	}
	res := s.skeeper.SetAdmin(s.ctx, admin)
	fmt.Println(res)

}
func (s *TestSuite) GetAdmin() {
	admin1 := types.Admin{
		Id:      "a1",
		Address: "cosmosv10222",
	}

	if admin1.Address != Gaddress {
		fmt.Println("not the admin")
	} else {
		fmt.Println("Admin successfully got")
	}

}
