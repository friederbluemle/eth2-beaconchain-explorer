package rpc

import (
	"eth2-exporter/types"
	"github.com/sirupsen/logrus"
)

// RpcClient provides an interface for RPC clients
type RpcClient interface {
	GetChainHead() (*types.ChainHead, error)
	GetEpochData(epoch uint64) (*types.EpochData, error)
	GetValidatorQueue() (*types.ValidatorQueue, map[string]uint64, error)
	GetAttestationPool() ([]*types.Attestation, error)
	GetEpochAssignments(epoch uint64) (*types.EpochAssignments, error)
	GetBlocksBySlot(slot uint64) ([]*types.Block, error)
	GetValidatorParticipation(epoch uint64) (*types.ValidatorParticipation, error)
}

var logger = logrus.New().WithField("module", "rpc")