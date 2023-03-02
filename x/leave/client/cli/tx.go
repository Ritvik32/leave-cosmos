package cli

import (
	"leave-cosmos/x/leave/types"
	"strconv"

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
		Use:   "to add an admin",
		Short: "id||name||address",
		Long:  "add admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			admin := types.AddAdminRequest{
				Id:           args[0],
				Name:         args[1],
				AdminAddress: args[2],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &admin)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewAddStudents() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "to add new students",
		Short: "student1[id||name||address||admin address]",
		Long:  "to add new students by that particular admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			student := types.AddStudentRequest{
				Id:           args[0],
				Name:         args[1],
				Address:      args[2],
				AdminAddress: args[3],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &student)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewApplyLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "to apply leave",
		Short: "student_address||Reason||from||to",
		Long:  "to apply leave  for the student",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			leave := types.ApplyLeaveRequest{
				Studentaddress: args[0],
				Reason:         args[1],
				From:           args[2],
				To:             args[3],
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &leave)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func NewAcceptLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "to accept leave",
		Short: "student_address||Reason||from||to",
		Long:  "to apply leave  for the student",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			num, _ := strconv.Atoi(args[1])
			accept := types.AcceptLeaveRequest{
				AdminAddress: args[0],
				Leaveid:      int64(num),
				Status:       types.LeaveStatus_STATUS_ACCEPTED,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &accept)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
