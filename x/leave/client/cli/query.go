package cli

import (
	"leave-cosmos/x/leave/types"

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
		GetStudent(),
		GetAllStudent(),
		GetAllLeave(),
		LeaveStatus(),
		GetAdmin(),
	)
	return cmd
}
func GetLeave() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Get-Student-Leave [StudentAddress]",
		Short:   "Used to list student leave",
		Long:    "Used to list student leave for specific user by specifying the student_id",
		Example: "./lmsd q leave Get-Student-Leave cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			//num, _ := strconv.Atoi(args[1])
			leave := types.ListStudentLeaveRequest{ //types.ListStudentLeaveRequest
				Stuaddress: args[0],
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
		Use:   "List-All-Leave",
		Short: "list all the leaves",
		Long:  "list all leaves for the student",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			leave := types.ListAllLeavesRequest{
				StudentId: args[0],
			}
			res, err := queryClient.ListAllLeave(cmd.Context(), &leave)
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
		Use:     "Get-Student [StudentAddress]",
		Short:   "used to retrieve student data",
		Long:    "get student details by student_id as string as a parameter",
		Example: "./lmsd q leave Get-Student cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua",
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

func GetAllStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Get-All-Student",
		Short:   "used to retrieve data regarding all the students",
		Long:    "get  all student details by student_id as string as a parameter",
		Example: "./lmsd q leave Get-All-Student",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			// student := types.GetStudentRequest{
			// Id: args[0],
			// }
			res, err := queryClient.GetAllStudentDetails(cmd.Context(), &types.GetStudentRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func LeaveStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "Leave-Status [StudentAddress]",
		Short:   "get status of student leave",
		Long:    "get leave status of the leave",
		Example: "./lmsd q leave Leave-Status cosmos1dcfxmfchss6r3rlqly8m3jc05538w7xgvtmyua",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			student := types.LeaveStatusRequest{
				StudentAddress: args[0],
			}
			res, err := queryClient.GetLeaveStatus(cmd.Context(), &student)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get-admin [AdminAddress]",
		Short:   "Used to get the admin details",
		Long:    "get leave status of the leave",
		Example: "./lmsd q leave get-admin cosmos14qypg0485c6mclvkqfwwwlryy0kqa3hyycdeul",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryMsgClient(clientCtx)
			admin := types.GetAdminRequest{
				Adminaddress: args[0],
			}
			res, err := queryClient.GetAdminDetails(cmd.Context(), &admin)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
