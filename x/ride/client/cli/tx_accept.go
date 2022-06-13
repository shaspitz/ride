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

func CmdAccept() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept [id-value]",
		Short: "Broadcast message accept",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIdValue := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAccept(
				clientCtx.GetFromAddress().String(),
				argIdValue,
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
