package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type TokenSentParser struct {
	BaseParser
}

func (p *TokenSentParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.TokenSent](block, attributes)
	config.Log.Debugf("[TokenSentParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type ContractCallSubmittedParser struct {
	BaseParser
}

func (p *ContractCallSubmittedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.ContractCallSubmitted](block, attributes)
	config.Log.Debugf("[ContractCallSubmittedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type ContractCallWithTokenSubmittedParser struct {
	BaseParser
}

func (p *ContractCallWithTokenSubmittedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.ContractCallWithTokenSubmitted](block, attributes)
	config.Log.Debugf("[ContractCallWithTokenSubmittedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type FeePaidParser struct {
	BaseParser
}

func (p *FeePaidParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.FeePaid](block, attributes)
	config.Log.Debugf("[FeePaidParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type FeeCollectedParser struct {
	BaseParser
}

func (p *FeeCollectedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.FeeCollected](block, attributes)
	config.Log.Debugf("[FeeCollectedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type ScalarTransferCompletedParser struct {
	BaseParser
}

func (p *ScalarTransferCompletedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.ScalarTransferCompleted](block, attributes)
	config.Log.Debugf("[ScalarTransferCompletedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}
