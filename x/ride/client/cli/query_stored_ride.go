package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/spf13/cobra"
)

func CmdListStoredRide() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stored-ride",
		Short: "list all storedRide",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStoredRideRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StoredRideAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowStoredRide() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stored-ride [index]",
		Short: "shows a storedRide",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetStoredRideRequest{
				Index: argIndex,
			}

			res, err := queryClient.StoredRide(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
