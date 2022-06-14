package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdFinish() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "finish [id-value] [end-location]",
		Short: "Broadcast message finish",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIdValue := args[0]
			argEndLocation := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgFinish(
				clientCtx.GetFromAddress().String(),
				argIdValue,
				argEndLocation,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
