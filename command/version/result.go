package version

import (
	"bytes"
	"fmt"

	"github.com/PrivixAI-labs/Privix-node/command/helper"
)

type VersionResult struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Branch    string `json:"branch"`
	BuildTime string `json:"buildTime"`
}

func (r *VersionResult) GetOutput() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n[VERSION INFO]\n")
	buffer.WriteString(helper.FormatKV([]string{
		fmt.Sprintf("Release version|%s", r.Version),
		fmt.Sprintf("Git branch|%s", r.Branch),
		//fmt.Sprintf("Commit hash|%s", r.Commit),
		fmt.Sprintf("Build time|%s", r.BuildTime),
	}))

	return buffer.String()
}

func main() {
	// Create an instance of VersionResult with manual values
	versionInfo := VersionResult{
		Version:   "1.0.2",
		Commit:    "0x0",
		Branch:    "dev",
		BuildTime: "11/14/2023",
	}

	// Call GetOutput() to get the formatted output
	output := versionInfo.GetOutput()

	// Print the output
	fmt.Println(output)
}
