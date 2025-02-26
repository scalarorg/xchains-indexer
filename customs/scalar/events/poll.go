package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type PollCompletedEventParser struct {
	BaseParser
}

func (p *PollCompletedEventParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.PollCompleted](block, attributes)
	config.Log.Debugf("[PollCompletedEventParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type PollExpiredEventParser struct {
	BaseParser
}

func (p *PollExpiredEventParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.PollExpired](block, attributes)
	config.Log.Debugf("[PollExpiredEventParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type PollFailedEventParser struct {
	BaseParser
}

func (p *PollFailedEventParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.PollFailed](block, attributes)
	config.Log.Debugf("[PollFailedEventParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type NoEventConfirmedParser struct {
	BaseParser
}

func (p *NoEventConfirmedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.NoEventConfirmed](block, attributes)
	config.Log.Debugf("[NoEventConfirmedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
