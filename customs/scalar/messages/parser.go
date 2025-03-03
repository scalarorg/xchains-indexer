package messages

import (
	"log"

	common "github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/chains"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/reward"
	"github.com/scalarorg/xchains-indexer/filter"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
	evmTypes "github.com/scalarorg/xchains-indexer/x/evm/types"
	voteTypes "github.com/scalarorg/xchains-indexer/x/vote/types"
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
	extendMessagesIndexerTokens(instance)
	var filters []filter.MessageTypeFilter
	customParsers := make(map[string]parsers.MessageParser)

	// basic
	basicRegexMessageTypeFilter, err := filter.NewRegexMessageTypeFilter("^/" + "cosmos.staking.v1beta1.MsgCreateValidator" + "$")
	if err != nil {
		log.Fatalf(ERR_FAILED_TO_CREATE_REGEX_MESSAGE_TYPE_FILTER, err)
		return err
	}
	filters = append(filters, basicRegexMessageTypeFilter)

	// Extend refund VoteRequest parser
	voteRequestFilter, err := filter.NewRegexMessageTypeFilter("^/" + voteTypes.MSG_VOTE_REQUEST + "$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return err
	}
	filters = append(filters, voteRequestFilter)
	customParsers["/"+voteTypes.MSG_VOTE_REQUEST] = &VoteRequestParser{
		Id:      "vote-request",
		Indexer: instance,
	}

	// Extend refund VoteEvents parser
	voteEventsFilter, err := filter.NewRegexMessageTypeFilter("^/" + evmTypes.MSG_EVM_VOTE_EVENTS + "$")
	if err != nil {
		log.Fatalf("Failed to create regex message type filter. Err: %v", err)
		return err
	}
	filters = append(filters, voteEventsFilter)
	customParsers["/"+evmTypes.MSG_EVM_VOTE_EVENTS] = &VoteEventsParser{
		Id:      "vote-events",
		Indexer: instance,
	}

	// Add RegisterChainMaintainerRequest parser
	// customParsers["/"+nexusTypes.MSG_NEXUS_REGISTER_CHAIN_MAINTAINER_REQUEST] = &RegisterChainMaintainerRequestParser{
	// 	Id:      "register-chain-maintainer-request",
	// 	Indexer: instance,
	// }

	for _, filter := range filters {
		instance.RegisterMessageTypeFilter(filter)
	}
	for key, parser := range customParsers {
		instance.RegisterCustomMessageParser(key, parser)
	}
	instance.PostSetupCustomFunction = func(dataset indexer.PostSetupCustomDataset) error {
		reward.PostSetupCustomFunctionReward(instance, &dataset)
		chains.PostSetupCustomFunctionChains(instance, &dataset)
		postSetupCustomFunctionTokens(instance, &dataset)
		// Register msg types for the custom messages
		dataset.DB.AutoMigrate(&common.TxMessage{})
		if instance.ChainClient != nil {
			// instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(voteTypes.MSG_VOTE_REQUEST, (*sdk.Msg)(nil), &voteTypes.VoteRequest{})
			// instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(evmTypes.MSG_EVM_VOTE_EVENTS, (*sdk.Msg)(nil), &evmTypes.VoteEvents{})
			// instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(nexusTypes.MSG_NEXUS_REGISTER_CHAIN_MAINTAINER_REQUEST, (*sdk.Msg)(nil), &nexusTypes.RegisterChainMaintainerRequest{})
		}
		return nil
	}
	return nil
}
