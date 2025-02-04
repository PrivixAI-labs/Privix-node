package peers

import (
	"github.com/PrivixAI-labs/Privix-node/command/helper"
	"github.com/PrivixAI-labs/Privix-node/command/peers/add"
	"github.com/PrivixAI-labs/Privix-node/command/peers/list"
	"github.com/PrivixAI-labs/Privix-node/command/peers/status"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	peersCmd := &cobra.Command{
		Use:   "peers",
		Short: "Top level command for interacting with the network peers. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(peersCmd)

	registerSubcommands(peersCmd)

	return peersCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// peers status
		status.GetCommand(),
		// peers list
		list.GetCommand(),
		// peers add
		add.GetCommand(),
	)
}
