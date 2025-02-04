package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	ibft "github.com/PrivixAI-labs/Privix-node/command/NLG-ibft"
	"github.com/PrivixAI-labs/Privix-node/command/backup"

	//"github.com/PrivixAI-labs/Privix-node/command/bridge"
	"github.com/PrivixAI-labs/Privix-node/command/genesis"
	"github.com/PrivixAI-labs/Privix-node/command/helper"
	"github.com/PrivixAI-labs/Privix-node/command/license"
	"github.com/PrivixAI-labs/Privix-node/command/monitor"
	"github.com/PrivixAI-labs/Privix-node/command/peers"

	//"github.com/PrivixAI-labs/Privix-node/command/polybft"
	//"github.com/PrivixAI-labs/Privix-node/command/polybftsecrets"
	//"github.com/PrivixAI-labs/Privix-node/command/regenesis"
	//"github.com/PrivixAI-labs/Privix-node/command/rootchain"
	"github.com/PrivixAI-labs/Privix-node/command/secrets"
	"github.com/PrivixAI-labs/Privix-node/command/server"
	"github.com/PrivixAI-labs/Privix-node/command/status"
	"github.com/PrivixAI-labs/Privix-node/command/txpool"
	"github.com/PrivixAI-labs/Privix-node/command/version"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "The go implementation of privix core",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		//	rootchain.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
	//	polybftsecrets.GetCommand(),
	//	polybft.GetCommand(),
	//	bridge.GetCommand(),
	//regenesis.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
