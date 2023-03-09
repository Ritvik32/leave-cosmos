package cli

import (
	"leave-cosmos/x/leave/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "lms commands",
		Long:  "lms module commands",
		RunE:  client.ValidateCmd,
	}
	txCmd.AddCommand(
		NewAddAdmin(),
		NewAddStudents(),
		NewApplyLeave(),
		NewAcceptLeave(),
	)
	return txCmd
}
func NewAddAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Add-Admin [Id] [Name]",
		Short:   "used to add admin",
		Long:    "used to add admin",
		Example: "./lmsd tx leave Add-Admin 1 adminname  --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()
			admin := types.AddAdminRequest{
				Id:           args[0],
				Name:         args[1],
				AdminAddress: fromAddress,
			}
			msg := types.NewAddAdminReq(admin)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewAddStudents() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Add-Student [Id] [Name] [Address]",
		Short:   "used to add new students",
		Long:    "to add new students by that particular admin",
		Example: "./lmsd tx leave Add-Student s1 stuname cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()
			msg := types.AddStudentRequest{
				Id:           args[0],
				Name:         args[1],
				Address:      args[2],
				AdminAddress: fromAddress,
			}
			// msg := types.NewAddStudentReq(student)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewApplyLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Apply-Leave [Studentaddress] [Reason] [From] [To]",
		Short:   "used to apply new leave",
		Long:    "to apply leave  for the student",
		Example: "./lmsd tx leave Apply-Leave cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua cold 13-2-23 14-2-23  --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()
			leave := types.ApplyLeaveRequest{
				Studentaddress: args[0],
				Reason:         args[1],
				From:           args[2],
				To:             args[3],
				Adminaddress:   fromAddress,
			}
			msg := types.NewApplyReq(leave)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewAcceptLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Accept-Leave [StudentAddress]",
		Short:   "used to accept student leave",
		Long:    "to accept leave  for the student",
		Example: "./lmsd tx leave Accept-Leave cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua --from cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul --chain-id testnet",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress().String()
			AceptLeave := types.AcceptLeaveRequest{
				Leave: &types.Leave{
					Studentaddress: args[0],
					Adminaddress:   fromAddress,
					Status:         types.LeaveStatus_STATUS_ACCEPTED,
				},
			}
			msg := types.NewAcceptLeaveReq(AceptLeave)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
