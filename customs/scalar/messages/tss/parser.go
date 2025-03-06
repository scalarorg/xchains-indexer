package tss

import (
	"log"

	rewardExported "github.com/scalarorg/scalar-core/x/reward/exported"
	tssTypes "github.com/scalarorg/scalar-core/x/tss/types"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_TSS_CHAINS_FILTER     string = "^/scalar.tss.v1beta1.*"
	MSG_SCALAR_TSS_HEARTBEAT_REQUEST string = "/scalar.tss.v1beta1.HeartBeatRequest"
)

func ExtendMessagesIndexerTss(instance *indexer.Indexer) {
	messageTypeFilter, err := filter.NewRegexMessageTypeFilter(MSG_SCALAR_TSS_CHAINS_FILTER)
	if err == nil {
		instance.RegisterMessageTypeFilter(messageTypeFilter)
	} else {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
	}
	customParsers := []parsers.MessageParser{}

	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionTss(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate()
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*rewardExported.Refundable)(nil), &tssTypes.HeartBeatRequest{})

	}
}
