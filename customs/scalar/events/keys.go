package events

import (
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
)

type KeygenStartedParser struct {
	BaseParser
}

func (p *KeygenStartedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeygenStarted](block, attributes)
	config.Log.Debugf("[KeygenStartedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type KeygenCompletedParser struct {
	BaseParser
}

func (p *KeygenCompletedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeygenCompleted](block, attributes)
	config.Log.Debugf("[KeygenCompletedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type KeygenExpiredParser struct {
	BaseParser
}

func (p *KeygenExpiredParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeygenExpired](block, attributes)
	config.Log.Debugf("[KeygenExpiredParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type KeyAssignedParser struct {
	BaseParser
}

func (p *KeyAssignedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeyAssigned](block, attributes)
	config.Log.Debugf("[KeyAssignedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type MultiSigKeyRotatedParser struct {
	BaseParser
}

func (p *MultiSigKeyRotatedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeyRotated](block, attributes)
	config.Log.Debugf("[MultiSigKeyRotatedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type KeyRotatedParser struct {
	BaseParser
}

func (p *KeyRotatedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeyRotated](block, attributes)
	config.Log.Debugf("[KeyRotatedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type SigningStartedParser struct {
	BaseParser
}

func (p *SigningStartedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.SigningStarted](block, attributes)
	config.Log.Debugf("[SigningStartedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type SigningCompletedParser struct {
	BaseParser
}

func (p *SigningCompletedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.SigningCompleted](block, attributes)
	config.Log.Debugf("[SigningCompletedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type SigningExpiredParser struct {
	BaseParser
}

func (p *SigningExpiredParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.SigningExpired](block, attributes)
	config.Log.Debugf("[SigningExpiredParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type SignatureSubmittedParser struct {
	BaseParser
}

func (p *SignatureSubmittedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.SignatureSubmitted](block, attributes)
	config.Log.Debugf("[SignatureSubmittedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type PubkeySubmittedParser struct {
	BaseParser
}

func (p *PubkeySubmittedParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.PubkeySubmitted](block, attributes)
	config.Log.Debugf("[PubkeySubmittedParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type KeygenOptOutParser struct {
	BaseParser
}

func (p *KeygenOptOutParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeygenOptOut](block, attributes)
	config.Log.Debugf("[KeygenOptOutParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}

type KeygenOptInParser struct {
	BaseParser
}

func (p *KeygenOptInParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	if _, err := p.BaseParser.ParseBlockEvent(block, event, attributes, conf); err != nil {
		return nil, err
	}
	var model any = CreateEvent[indexer.KeygenOptIn](block, attributes)
	config.Log.Debugf("[KeygenOptInParser] ParseBlockEvent# model: %++v", model)
	return &model, nil
}
