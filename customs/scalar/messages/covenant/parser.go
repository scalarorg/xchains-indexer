package covenant

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	rewardExported "github.com/scalarorg/scalar-core/x/reward/exported"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_COVENANT_CHAINS_FILTER                  string = "^/scalar.covenant.v1beta1.*"
	MSG_SCALAR_COVENANT_SUBMIT_TAP_SCRIPT_SIGS_REQUEST string = "/scalar.covenant.v1beta1.SubmitTapScriptSigsRequest"
	MSG_SCALAR_COVENANT_ROTATE_KEY_REQUEST             string = "/scalar.covenant.v1beta1.RotateKeyRequest"
)

func ExtendMessagesIndexerCovenant(instance *indexer.Indexer) {
	messageTypeFilter, err := filter.NewRegexMessageTypeFilter(MSG_SCALAR_COVENANT_CHAINS_FILTER)
	if err == nil {
		instance.RegisterMessageTypeFilter(messageTypeFilter)
	} else {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
	}
	customParsers := []parsers.MessageParser{}
	customParsers = append(customParsers, &RotateKeyRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_COVENANT_ROTATE_KEY_REQUEST,
			Indexer: instance,
		},
	})
	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionCovenant(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate(
		&CovenantKeyMsg{},
	)
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*rewardExported.Refundable)(nil), &covenantTypes.SubmitTapScriptSigsRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_COVENANT_ROTATE_KEY_REQUEST, (*sdk.Msg)(nil), &covenantTypes.RotateKeyRequest{})
	}
}
