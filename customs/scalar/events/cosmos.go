package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type HeartBeatParser struct {
	BaseParser
}

// Custom event parser
func (p *HeartBeatParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	// for _, attribute := range attributes {
	// 	config.Log.Debugf("[HeartBeatParser] ParseBlockEvent# %s => %++v", attribute.BlockEventAttributeKey.Key, attribute.Value)
	// }
	var model any = CreateEvent[indexer.Heartbeat](block, attributes)
	config.Log.Debugf("[HeartBeatParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
