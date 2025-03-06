package vote

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	rewardExported "github.com/scalarorg/scalar-core/x/reward/exported"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_VOTE_CHAINS_FILTER string = "^/scalar.vote.v1beta1.*"
	MSG_SCALAR_VOTE_REQUEST       string = "/scalar.vote.v1beta1.VoteRequest"
	//MSG_SOURCE_TXS_REQUEST   string = "ConfirmSourceTxsRequest"
)

func ExtendMessagesIndexerVote(instance *indexer.Indexer) {
	messageTypeFilter, err := filter.NewRegexMessageTypeFilter(MSG_SCALAR_VOTE_CHAINS_FILTER)
	if err == nil {
		instance.RegisterMessageTypeFilter(messageTypeFilter)
	} else {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
	}
	customParsers := []parsers.MessageParser{}
	// Extend VoteRequest parser
	customParsers = append(customParsers, &VoteRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_VOTE_REQUEST,
			Indexer: instance,
		},
	})
	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionVote(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate(
		&VoteRequestMsg{},
	)
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*rewardExported.Refundable)(nil), &voteTypes.VoteRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_VOTE_REQUEST, (*sdk.Msg)(nil), &voteTypes.VoteRequest{})
	}
}
