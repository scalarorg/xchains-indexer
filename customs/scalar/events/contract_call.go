package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type ContractCallApprovedParser struct {
	BaseParser
}

// Custom event parser
func (p *ContractCallApprovedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.EventContractCallApproved](block, attributes)
	config.Log.Debugf("[ContractCallApprovedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type ContractCallFailedParser struct {
	BaseParser
}

// Custom event parser
func (p *ContractCallFailedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.ContractCallFailed](block, attributes)
	config.Log.Debugf("[ContractCallFailedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type EventContractCallWithMintApprovedParser struct {
	BaseParser
}

// Custom event parser
func (p *EventContractCallWithMintApprovedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.EventContractCallWithMintApproved](block, attributes)
	config.Log.Debugf("[EventContractCallWithMintApprovedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
