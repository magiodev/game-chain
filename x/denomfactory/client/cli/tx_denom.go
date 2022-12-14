package cli

import (
	"github.com/G4AL-Entertainment/g4al-chain/x/denomfactory/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom [symbol] [project] [max-supply] [can-change-max-supply] [name] [description] [precision] [uri] [uri_hash]",
		Short: "Create a new Denom",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argProject := args[1]
			argMaxSupply, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argCanChangeMaxSupply, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}
			argName := args[4]
			argDescription := args[5]
			argPrecision, err := cast.ToUint32E(args[6])
			if err != nil {
				return err
			}
			argUri := args[7]
			argUriHash := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDenom(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argProject,
				argMaxSupply,
				argCanChangeMaxSupply,
				argName,
				argDescription,
				argPrecision,
				argUri,
				argUriHash,
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

func CmdUpdateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-denom [symbol] [project] [max-supply] [name] [description] [uri] [uri_hash]",
		Short: "Update a Denom",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argProject := args[1]
			argMaxSupply, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argName := args[3]
			argDescription := args[4]
			argUri := args[5]
			argUriHash := args[6]

			msg := types.NewMsgUpdateDenom(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argProject,
				argMaxSupply,
				argName,
				argDescription,
				argUri,
				argUriHash,
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
