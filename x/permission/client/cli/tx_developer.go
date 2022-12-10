package cli

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/permission/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateDeveloper() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-developer [address] [created-at] [updated-at] [blocked]",
		Short: "Create a new Developer",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexAddress := args[0]

			// Get value arguments
			argCreatedAt, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argUpdatedAt, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argBlocked, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDeveloper(
				clientCtx.GetFromAddress().String(),
				indexAddress,
				argCreatedAt,
				argUpdatedAt,
				argBlocked,
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

func CmdUpdateDeveloper() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-developer [address] [created-at] [updated-at] [blocked]",
		Short: "Update a Developer",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexAddress := args[0]

			// Get value arguments
			argCreatedAt, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argUpdatedAt, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argBlocked, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDeveloper(
				clientCtx.GetFromAddress().String(),
				indexAddress,
				argCreatedAt,
				argUpdatedAt,
				argBlocked,
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
