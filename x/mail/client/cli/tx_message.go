package cli

import (
	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdSendMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-message [receiver] [body] [forwarded-id]",
		Short: "Create a new message",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argBody := args[1]

			argForwardedId := uint64(0)
			if len(args) == 3 {
				var err error
				argForwardedId, err = cast.ToUint64E(args[2])
				if err != nil {
					return err
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendMessage(clientCtx.GetFromAddress().String(), argReceiver, argBody, argForwardedId)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
