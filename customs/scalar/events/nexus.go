package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type RateLimitUpdatedParser struct {
	BaseParser
}

// Custom event parser
func (p *RateLimitUpdatedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	var result any = CreateEvent[indexer.RateLimitUpdated](block, attributes)
	config.Log.Debugf("[RateLimitUpdatedParser] ParseBlockEvent# result: %++v", result)
	return &result, nil
}

type MessageExecutedEventParser struct {
	BaseParser
}

func (p *MessageExecutedEventParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.MessageExecuted](block, attributes)
	config.Log.Debugf("[MessageExecutedEventParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
