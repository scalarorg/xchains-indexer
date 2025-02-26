package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type SigningPsbtStartedParser struct {
	BaseParser
}

func (p *SigningPsbtStartedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.SigningPsbtStarted](block, attributes)
	config.Log.Debugf("[SigningPsbtStartedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type SigningPsbtCompletedParser struct {
	BaseParser
}

func (p *SigningPsbtCompletedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.SigningPsbtCompleted](block, attributes)
	config.Log.Debugf("[SigningPsbtCompletedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type SigningPsbtExpiredParser struct {
	BaseParser
}

func (p *SigningPsbtExpiredParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.SigningPsbtExpired](block, attributes)
	config.Log.Debugf("[SigningPsbtExpiredParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type TapScriptSigsSubmittedParser struct {
	BaseParser
}

func (p *TapScriptSigsSubmittedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	_, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf)
	if err != nil {
		return nil, err
	}

	var model any = CreateEvent[indexer.TapScriptSigsSubmitted](block, attributes)
	config.Log.Debugf("[TapScriptSigsSubmittedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
