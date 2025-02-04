package ibft

import (
	"github.com/PrivixAI-labs/Privix-node/command/NLG-ibft/candidates"
	"github.com/PrivixAI-labs/Privix-node/command/NLG-ibft/propose"
	"github.com/PrivixAI-labs/Privix-node/command/NLG-ibft/quorum"
	"github.com/PrivixAI-labs/Privix-node/command/NLG-ibft/snapshot"
	"github.com/PrivixAI-labs/Privix-node/command/NLG-ibft/status"
	_switch "github.com/PrivixAI-labs/Privix-node/command/NLG-ibft/switch"
	"github.com/PrivixAI-labs/Privix-node/command/helper"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "pribft",
		Short: "Top level privix ibft command for interacting with the privix ibft consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
