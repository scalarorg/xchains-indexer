package indexer

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/parsers"
)

func (indexer *Indexer) RegisterCustomModuleBasics(basics []module.AppModuleBasic) {
	indexer.CustomModuleBasics = append(indexer.CustomModuleBasics, basics...)
}

func (indexer *Indexer) RegisterMessageTypeFilter(filter filter.MessageTypeFilter) {
	indexer.MessageTypeFilters = append(indexer.MessageTypeFilters, filter)
}

func (indexer *Indexer) RegisterCustomModels(models []any) {
	indexer.CustomModels = append(indexer.CustomModels, models...)
}

func (indexer *Indexer) RegisterCustomBlockEventParser(eventKey string, parser parsers.BlockEventParser) {
	var err error
	indexer.CustomBlockEventParserRegistry, indexer.CustomBlockParserTrackers, err = customBlockEventRegistration(
		indexer.CustomBlockEventParserRegistry,
		indexer.CustomBlockParserTrackers,
		eventKey,
		parser,
	)

	if err != nil {
		config.Log.Fatal("Error registering BeginBlock custom parser", err)
	}
}

func (indexer *Indexer) RegisterCustomMessageParser(messageKey string, parser parsers.MessageParser) {
	if indexer.CustomMessageParserRegistry == nil {
		indexer.CustomMessageParserRegistry = make(map[string][]parsers.MessageParser)
	}

	if indexer.CustomMessageParserTrackers == nil {
		indexer.CustomMessageParserTrackers = make(map[string]models.MessageParser)
	}

	indexer.CustomMessageParserRegistry[messageKey] = append(indexer.CustomMessageParserRegistry[messageKey], parser)

	if _, ok := indexer.CustomMessageParserTrackers[parser.Identifier()]; ok {
		config.Log.Fatalf("Found duplicate message parser with identifier \"%s\", parsers must be uniquely identified", parser.Identifier())
	}

	indexer.CustomMessageParserTrackers[parser.Identifier()] = models.MessageParser{
		Identifier: parser.Identifier(),
	}
}

func customBlockEventRegistration(registry map[string][]parsers.BlockEventParser, tracker map[string]models.BlockEventParser, eventKey string, parser parsers.BlockEventParser) (map[string][]parsers.BlockEventParser, map[string]models.BlockEventParser, error) {
	if registry == nil {
		registry = make(map[string][]parsers.BlockEventParser)
	}

	if tracker == nil {
		tracker = make(map[string]models.BlockEventParser)
	}

	registry[eventKey] = append(registry[eventKey], parser)

	if _, ok := tracker[parser.Identifier()]; ok {
		return registry, tracker, fmt.Errorf("found duplicate block event parser with identifier \"%s\", parsers must be uniquely identified", parser.Identifier())
	}

	tracker[parser.Identifier()] = models.BlockEventParser{
		Identifier: parser.Identifier(),
	}
	return registry, tracker, nil
}
