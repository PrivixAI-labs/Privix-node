package init

import (
	"bytes"
	"fmt"

	"github.com/PrivixAI-labs/Privix-node/command/helper"
	"github.com/PrivixAI-labs/Privix-node/types"
)

type SecretsInitResult struct {
	Address      types.Address `json:"address"`
	BLSPubkey    string        `json:"bls_pubkey"`
	NodeID       string        `json:"node_id"`
	ValidatorKey string        `json:"validator_key"`
	// Insecure     bool          `json:"insecure"`
}

func (r *SecretsInitResult) GetOutput() string {
	var buffer bytes.Buffer

	vals := make([]string, 0, 3)

	vals = append(
		vals,
		fmt.Sprintf("Public key (address)|%s", r.Address.String()),
	)
	vals = append(
		vals,
		fmt.Sprintf("Private key| 0x%s", r.ValidatorKey),
	)

	// if r.BLSPubkey != "" {
	// 	vals = append(
	// 		vals,
	// 		fmt.Sprintf("BLS Public key|%s", r.BLSPubkey),
	// 	)
	// }

	vals = append(vals, fmt.Sprintf("Node ID|%s", r.NodeID))

	// if r.Insecure {
	// 	buffer.WriteString("\n[WARNING: INSECURE LOCAL SECRETS - SHOULD NOT BE RUN IN PRODUCTION]\n")
	// }

	buffer.WriteString("\n[SECRETS INIT]\n")
	buffer.WriteString(helper.FormatKV(vals))
	buffer.WriteString("\n")

	return buffer.String()
}
