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
		Use:   "create-denom [symbol] [project] [max-supply] [can-change-max-supply] [name] [description] [uri] [uri_hash]",
		Short: "Create a new Denom",
		Args:  cobra.ExactArgs(8),
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
		Use:   "update-denom [symbol] [max-supply] [name] [description] [uri] [uri_hash]",
		Short: "Update a Denom",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argMaxSupply, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			argName := args[2]
			argDescription := args[3]
			argUri := args[4]
			argUriHash := args[5]

			msg := types.NewMsgUpdateDenom(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
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

func CmdMintDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-denom [symbol] [amount] [receiver]",
		Short: "Broadcast message mint-denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSymbol := args[0]
			argAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argReceiver := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintDenom(
				clientCtx.GetFromAddress().String(),
				argSymbol,
				argAmount,
				argReceiver,
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

func CmdBurnDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-denom [symbol] [amount]",
		Short: "Broadcast message burn-denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSymbol := args[0]
			argAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnDenom(
				clientCtx.GetFromAddress().String(),
				argSymbol,
				argAmount,
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

func CmdTransferDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-denom [symbol] [amount] [receiver]",
		Short: "Broadcast message transfer-denom",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSymbol := args[0]
			argAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argReceiver := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferDenom(
				clientCtx.GetFromAddress().String(),
				argSymbol,
				argAmount,
				argReceiver,
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
