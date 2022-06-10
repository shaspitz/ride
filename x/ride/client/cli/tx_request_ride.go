package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/smarshall-spitzbart/ride/x/ride/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRequestRide() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-ride [start-location] [destination] [mutual-stake] [hourly-pay] [distance-tip]",
		Short: "Broadcast message requestRide",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStartLocation := args[0]
			argDestination := args[1]
			argMutualStake, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argHourlyPay, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argDistanceTip, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestRide(
				clientCtx.GetFromAddress().String(),
				argStartLocation,
				argDestination,
				argMutualStake,
				argHourlyPay,
				argDistanceTip,
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
