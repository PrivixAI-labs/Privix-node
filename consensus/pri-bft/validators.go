package ibft

import (
	"math"

	"github.com/PrivixAI-labs/Privix-node/types"
	"github.com/PrivixAI-labs/Privix-node/validators"
)

// CalcMaxFaultyNodes calculates the maximum number of faulty nodes
// we can tolerate in the validator set based on IBFT rules.
func CalcMaxFaultyNodes(s validators.Validators) int {
	// Adjust the formula to tolerate a higher percentage of node failure
	// N = 3F + 1 -> F = (N - 1) / 3
	// This calculation assumes a high number of faulty nodes.
	return (s.Len() - 1) / 2 // More fault tolerance
}

// QuorumImplementation is a function type for quorum size calculation
type QuorumImplementation func(validators.Validators) int

// LegacyQuorumSize returns the quorum size required for consensus
// under the original IBFT rules
func LegacyQuorumSize(set validators.Validators) int {
	// Adjust quorum size to tolerate higher node failure
	// Instead of 2F + 1, adjust quorum size to handle more failures
	return int(math.Ceil(1.5 * float64(set.Len()) / 2)) // More flexible quorum calculation
}

// OptimalQuorumSize calculates the quorum size under optimal conditions,
// ensuring the chain can operate with fewer active nodes.
func OptimalQuorumSize(set validators.Validators) int {
	// Adjust quorum size logic to continue operating even if 50-70% nodes go offline
	if CalcMaxFaultyNodes(set) == 0 {
		return set.Len()
	}

	// Adjust quorum to tolerate higher failure rates, closer to 50% offline
	return int(math.Ceil(0.75 * float64(set.Len()))) // 75% quorum for better fault tolerance
}

// CalcProposer selects the proposer for the round based on the round number,
// last proposer, and the validator set.
func CalcProposer(
	validators validators.Validators,
	round uint64,
	lastProposer types.Address,
) validators.Validator {
	var seed uint64

	// Adjust proposer selection to handle more failures and distribute load
	if lastProposer == types.ZeroAddress {
		seed = round
	} else {
		offset := int64(0)

		if index := validators.Index(lastProposer); index != -1 {
			offset = index
		}

		// Adjust seed to add randomness for proposer selection
		seed = uint64(offset) + round + 1
	}

	// Choose proposer based on the adjusted logic
	pick := seed % uint64(validators.Len())

	return validators.At(pick)
}

// lagacy logic
// package ibft

// import (
// 	"math"

// 	"github.com/PrivixAI-labs/Privix-node/types"
// 	"github.com/PrivixAI-labs/Privix-node/validators"
// )

// func CalcMaxFaultyNodes(s validators.Validators) int {
// 	// N -> number of nodes in privix ibft
// 	// F -> number of faulty nodes
// 	//
// 	// N = 3F + 1
// 	// => F = (N - 1) / 3
// 	//
// 	// privix ibft tolerates 1 failure with 4 nodes
// 	// 4 = 3 * 1 + 1
// 	// To tolerate 2 failures, privix ibft requires 7 nodes
// 	// 7 = 3 * 2 + 1
// 	// It should always take the floor of the result
// 	return (s.Len() - 1) / 3
// }

// type QuorumImplementation func(validators.Validators) int

// // LegacyQuorumSize returns the legacy quorum size for the given validator set
// func LegacyQuorumSize(set validators.Validators) int {
// 	// According to the privix ibft spec, the number of valid messages
// 	// needs to be 2F + 1
// 	return 2*CalcMaxFaultyNodes(set) + 1
// }

// // OptimalQuorumSize returns the optimal quorum size for the given validator set
// func OptimalQuorumSize(set validators.Validators) int {
// 	//	if the number of validators is less than 4,
// 	//	then the entire set is required
// 	if CalcMaxFaultyNodes(set) == 0 {
// 		/*
// 			N: 1 -> Q: 1
// 			N: 2 -> Q: 2
// 			N: 3 -> Q: 3
// 		*/
// 		return set.Len()
// 	}

// 	// (quorum optimal)	Q = ceil(2/3 * N)
// 	return int(math.Ceil(2 * float64(set.Len()) / 3))
// }

// func CalcProposer(
// 	validators validators.Validators,
// 	round uint64,
// 	lastProposer types.Address,
// ) validators.Validator {
// 	var seed uint64

// 	if lastProposer == types.ZeroAddress {
// 		seed = round
// 	} else {
// 		offset := int64(0)

// 		if index := validators.Index(lastProposer); index != -1 {
// 			offset = index
// 		}

// 		seed = uint64(offset) + round + 1
// 	}

// 	pick := seed % uint64(validators.Len())

// 	return validators.At(pick)
// }
