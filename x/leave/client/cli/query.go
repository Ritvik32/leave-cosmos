package cli

import (
	"leave-cosmos/x/leave/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "query commands",
		Long:  "query commands",
		RunE:  client.ValidateCmd,
	}
	cmd.AddCommand(
		GetLeave(),
		GetAllLeave(),
		GetStudent(),
	)
	return cmd
}
func GetLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get user specified leave status",
		Short: "studentid||leaveid",
		Long:  "query commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			num, _ := strconv.Atoi(args[1])
			leave := types.ListStudentLeaveRequest{
				Id:      args[0],
				Leaveid: int64(num),
			}
			res, err := queryClient.ListStudentLeave(cmd.Context(), &leave)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func GetAllLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list all leaves",
		Short: "list all leaves",
		Long:  "list all leaves",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			res, err := queryClient.ListAllLeave(cmd.Context(), &types.ListAllLeavesRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func GetStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get student details",
		Short: "StudentId|| as a parameter",
		Long:  "get student details by student_id as string as a parameter",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			student := types.GetStudentRequest{
				Id: args[0],
			}
			res, err := queryClient.GetStudentDetails(cmd.Context(), &student)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
