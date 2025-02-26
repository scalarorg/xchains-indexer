package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type CommandBatchSignedParser struct {
	BaseParser
}

// Custom event parser
func (p *CommandBatchSignedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.CommandBatchSigned](block, attributes)
	config.Log.Debugf("[CommandBatchSignedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type CommandBatchAbortedParser struct {
	BaseParser
}

// Custom event parser
func (p *CommandBatchAbortedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.CommandBatchAborted](block, attributes)
	config.Log.Debugf("[CommandBatchAbortedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type MintCommandParser struct {
	BaseParser
}

// Custom event parser
func (p *MintCommandParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.MintCommand](block, attributes)
	config.Log.Debugf("[MintCommandParser] ParseBlockEvent# model: %++v", model)
	return nil, nil
}

type BurnCommandParser struct {
	BaseParser
}

// Custom event parser
func (p *BurnCommandParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.BurnCommand](block, attributes)
	config.Log.Debugf("[BurnCommandParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
