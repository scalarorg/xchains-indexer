package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type MessageEventParser struct {
	BaseParser
}

func (p *MessageEventParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	for _, attribute := range attributes {
		config.Log.Debugf("[MessageEventParser] ParseBlockEvent# %s => %++v", attribute.BlockEventAttributeKey.Key, attribute.Value)
	}
	config.Log.Debugf("[MessageEventParser] ParseBlockEvent# model: %++v", event)
	return nil, nil
}

type MessageProcessingEventParser struct {
	BaseParser
}

func (p *MessageProcessingEventParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	config.Log.Debugf("MessageProcessingEventParser# ParseBlockEvent: %v", event)

	return nil, nil
}
