package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type TokenConfirmationParser struct {
	BaseParser
}

func (p *TokenConfirmationParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}
	for _, attribute := range attributes {
		config.Log.Debugf("[TokenConfirmationParser] ParseBlockEvent# %s => %++v", attribute.BlockEventAttributeKey.Key, attribute.Value)
	}
	var model any = CreateEvent[indexer.TokenConfirmation](block, attributes)
	config.Log.Debugf("[TokenConfirmationParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
