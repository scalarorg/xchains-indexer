package messages

import (
	"log"

	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_TOKEN_MINT_MSG string = "scalar.scalarnet.v1beta1.MsgMint"
)

func extendMessagesIndexerTokens(instance *indexer.Indexer) {
	var customFilter filter.MessageTypeFilter
	customParsers := make(map[string]parsers.MessageParser)
	customFilter, err := filter.NewRegexMessageTypeFilter("^/" + MSG_SCALAR_TOKEN_MINT_MSG + ".*")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return
	}
	instance.RegisterMessageTypeFilter(customFilter)
	for key, parser := range customParsers {
		instance.RegisterCustomMessageParser(key, parser)
	}
}

func postSetupCustomFunctionTokens(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate()
}
