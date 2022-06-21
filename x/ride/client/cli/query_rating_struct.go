package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/spf13/cobra"
)

func CmdListRatingStruct() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-rating-struct",
		Short: "list all ratingStruct",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllRatingStructRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.RatingStructAll(context.Background(), params)
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

func CmdShowRatingStruct() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-rating-struct [index]",
		Short: "shows a ratingStruct",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetRatingStructRequest{
				Index: argIndex,
			}

			res, err := queryClient.RatingStruct(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
