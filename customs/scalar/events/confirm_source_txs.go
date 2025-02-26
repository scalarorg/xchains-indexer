package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type EventConfirmSourceTxsParser struct {
	BaseParser
}

// Custom event parser
func (p *EventConfirmSourceTxsParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.EventConfirmSourceTxs](block, attributes)
	config.Log.Debugf("[EventConfirmSourceTxsParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
