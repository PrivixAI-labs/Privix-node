package main

import (
	_ "embed"

	"github.com/PrivixAI-labs/Privix-node/command/root"
	"github.com/PrivixAI-labs/Privix-node/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
