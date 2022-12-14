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
		Use:   "create-class [symbol] [project] [max-supply] [can-change-max-supply] [name] [description] [uri] [uri-hash] [data]",
		Short: "Create a new Class",
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
			argUri := args[6]
			argUriHash := args[7]
			argData := args[8]

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
				argName,
				argDescription,
				argUri,
				argUriHash,
				argData,
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
		Use:   "update-class [symbol] [max-supply] [name] [description] [uri] [uri-hash] [data]",
		Short: "Update a Class",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argMaxSupply, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argName := args[2]
			argDescription := args[3]
			argUri := args[4]
			argUriHash := args[5]
			argData := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateClass(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argMaxSupply,
				argName,
				argDescription,
				argUri,
				argUriHash,
				argData,
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
