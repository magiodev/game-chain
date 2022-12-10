package cli

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/assetfactory/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-class [symbol] [project] [max-supply] [can-change-max-supply]",
		Short: "Create a new Class",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argProject := args[1]
			argMaxSupply, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argCanChangeMaxSupply, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateClass(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argProject,
				argMaxSupply,
				argCanChangeMaxSupply,
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

func CmdUpdateClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-class [symbol] [project] [max-supply] [can-change-max-supply]",
		Short: "Update a Class",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argProject := args[1]
			argMaxSupply, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argCanChangeMaxSupply, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateClass(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argProject,
				argMaxSupply,
				argCanChangeMaxSupply,
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
