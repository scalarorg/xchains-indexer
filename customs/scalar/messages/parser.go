package messages

import (
	"log"

	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/chains"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/reward"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/vote"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
)

const (
	ERR_FAILED_TO_CREATE_REGEX_MESSAGE_TYPE_FILTER string = "failed to create regex message type filter. Err: %v"
	EVENT_TYPE_MESSAGE                             string = "message"
)

func ExtendMessagesIndexer(instance *indexer.Indexer) error {
	//Register all message types
	scalarFilter, err := filter.NewRegexMessageTypeFilter("^/scalar.*.v1beta1.*")
	if err != nil {
		log.Fatalf(ERR_FAILED_TO_CREATE_REGEX_MESSAGE_TYPE_FILTER, err)
		return err
	}
	instance.RegisterMessageTypeFilter(scalarFilter)
	reward.ExtendMessagesIndexerReward(instance)
	chains.ExtendMessagesIndexerChains(instance)
	vote.ExtendMessagesIndexerVote(instance)
	extendMessagesIndexerTokens(instance)

	instance.PostSetupCustomFunction = func(dataset indexer.PostSetupCustomDataset) error {
		reward.PostSetupCustomFunctionReward(instance, &dataset)
		chains.PostSetupCustomFunctionChains(instance, &dataset)
		vote.PostSetupCustomFunctionVote(instance, &dataset)
		return nil
	}
	return nil
}
