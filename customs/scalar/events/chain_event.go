package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type ChainEventConfirmedParser struct {
	BaseParser
}

// Custom event parser
func (p *ChainEventConfirmedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.ChainEventConfirmed](block, attributes)
	config.Log.Debugf("[ChainEventConfirmedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type ChainEventCompletedParser struct {
	BaseParser
}

// Custom event parser
func (p *ChainEventCompletedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.ChainEventCompleted](block, attributes)
	config.Log.Debugf("[ChainEventCompletedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type ChainEventFailedParser struct {
	BaseParser
}

// Custom event parser
func (p *ChainEventFailedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.ChainEventFailed](block, attributes)
	config.Log.Debugf("[ChainEventFailedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type ChainEventRetryFailedParser struct {
	BaseParser
}

// Custom event parser
func (p *ChainEventRetryFailedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.ChainEventRetryFailed](block, attributes)
	config.Log.Debugf("[ChainEventRetryFailedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
