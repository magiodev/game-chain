package cli

import (
	"strconv"
	"strings"

	"github.com/G4AL-Entertainment/g4al-chain/x/game/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [symbol] [name] [description]",
		Short: "Create a new Project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argName := args[1]
			argDescription := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProject(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argName,
				argDescription,
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

func CmdUpdateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-project [symbol] [name] [description]",
		Short: "Update a Project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexSymbol := args[0]

			// Get value arguments
			argName := args[1]
			argDescription := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProject(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argName,
				argDescription,
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

var _ = strconv.Itoa(0)

func CmdAddDelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-delegate [symbol] [delegate]",
		Short: "Broadcast message AddDelegate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexSymbol := args[0]

			argDelegate := strings.Split(args[1], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddDelegate(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argDelegate,
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

var _ = strconv.Itoa(0)

func CmdRemoveDelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-delegate [delegate] [symbol]",
		Short: "Broadcast message RemoveDelegate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexSymbol := args[0]

			argDelegate := strings.Split(args[1], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveDelegate(
				clientCtx.GetFromAddress().String(),
				indexSymbol,
				argDelegate,
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
