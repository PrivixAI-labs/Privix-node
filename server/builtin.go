package server

import (
	"github.com/PrivixAI-labs/Privix-node/chain"
	"github.com/PrivixAI-labs/Privix-node/consensus"
	consensusIBFT "github.com/PrivixAI-labs/Privix-node/consensus/pri-bft"
	consensusDev "github.com/PrivixAI-labs/Privix-node/consensus/dev"
	consensusDummy "github.com/PrivixAI-labs/Privix-node/consensus/dummy"

	//consensusPolyBFT
	"github.com/PrivixAI-labs/Privix-node/forkmanager"
	"github.com/PrivixAI-labs/Privix-node/secrets"
	"github.com/PrivixAI-labs/Privix-node/secrets/awsssm"
	"github.com/PrivixAI-labs/Privix-node/secrets/gcpssm"
	"github.com/PrivixAI-labs/Privix-node/secrets/hashicorpvault"
	"github.com/PrivixAI-labs/Privix-node/secrets/local"
	"github.com/PrivixAI-labs/Privix-node/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus  ConsensusType = "dev"
	IBFTConsensus ConsensusType = "ibft"
	//PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:  consensusDev.Factory,
	IBFTConsensus: consensusIBFT.Factory,
	//PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	// PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
