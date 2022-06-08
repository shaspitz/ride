package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/spf13/cobra"
)

func CmdShowNextRide() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-next-ride",
		Short: "shows nextRide",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetNextRideRequest{}

			res, err := queryClient.NextRide(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
