package rpc

import (
	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
)

type ResultBlockResults struct {
	Height                int64                     `json:"height"`
	TxsResults            []*abci.ExecTxResult      `json:"txs_results"`
	BeginBlockEvents      []abci.Event              `json:"begin_block_events"`
	EndBlockEvents        []abci.Event              `json:"end_block_events"`
	ValidatorUpdates      []abci.ValidatorUpdate    `json:"validator_updates"`
	ConsensusParamUpdates *cmtproto.ConsensusParams `json:"consensus_param_updates"`
	AppHash               []byte                    `json:"app_hash"`
}

func (r *ResultBlockResults) ToCometbftBlockResult() *ctypes.ResultBlockResults {
	return &ctypes.ResultBlockResults{
		Height:                r.Height,
		TxsResults:            r.TxsResults,
		FinalizeBlockEvents:   append(r.BeginBlockEvents, r.EndBlockEvents...),
		ValidatorUpdates:      r.ValidatorUpdates,
		ConsensusParamUpdates: r.ConsensusParamUpdates,
		AppHash:               r.AppHash,
	}
}
