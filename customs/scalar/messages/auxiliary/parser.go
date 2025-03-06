package auxiliary

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	auxiliaryTypes "github.com/scalarorg/scalar-core/x/auxiliary/types"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_AUXILIARY_BATCH_REQUEST_FILTER string = "^/scalar.auxiliary.v1beta1.*"
	MSG_SCALAR_AUXILIARY_BATCH_REQUEST        string = "/scalar.auxiliary.v1beta1.BatchRequest"
)

func ExtendMessagesIndexerAuxiliary(instance *indexer.Indexer) {
	messageTypeFilter, err := filter.NewRegexMessageTypeFilter(MSG_SCALAR_AUXILIARY_BATCH_REQUEST_FILTER)
	if err == nil {
		instance.RegisterMessageTypeFilter(messageTypeFilter)
	} else {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
	}
	customParsers := []parsers.MessageParser{}

	customParsers = append(customParsers, &BatchRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_AUXILIARY_BATCH_REQUEST,
			Indexer: instance,
		},
	})
	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionAuxiliary(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate(
		&BatchRequestMsg{},
	)
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_AUXILIARY_BATCH_REQUEST, (*sdk.Msg)(nil), &auxiliaryTypes.BatchRequest{})
	}
}
