package reward

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rewardTypes "github.com/scalarorg/scalar-core/x/reward/types"
	common "github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_REWARD_CHAINS_FILTER      string = "^/scalar.reward.v1beta1.*"
	MSG_SCALAR_REWARD_REFUND_MSG_REQUEST string = "/scalar.reward.v1beta1.RefundMsgRequest"
	//MSG_SOURCE_TXS_REQUEST   string = "ConfirmSourceTxsRequest"
)

func ExtendMessagesIndexerReward(instance *indexer.Indexer) {
	customParsers := make(map[string]parsers.MessageParser)
	// Extend RefundMsgRequest parser
	customParsers[MSG_SCALAR_REWARD_REFUND_MSG_REQUEST] = &RefundMsgRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      "refund-msg-request",
			Indexer: instance,
		},
	}
	for key, parser := range customParsers {
		instance.RegisterCustomMessageParser(key, parser)
	}
}

func PostSetupCustomFunctionReward(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate()
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_REWARD_REFUND_MSG_REQUEST, (*sdk.Msg)(nil), &rewardTypes.RefundMsgRequest{})
		//instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(voteTypes.MSG_VOTE_REQUEST, (*sdk.Msg)(nil), &voteTypes.VoteRequest{})
		//instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(evmTypes.MSG_EVM_VOTE_EVENTS, (*sdk.Msg)(nil), &evmTypes.VoteEvents{})
		//instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(nexusTypes.MSG_NEXUS_REGISTER_CHAIN_MAINTAINER_REQUEST, (*sdk.Msg)(nil), &nexusTypes.RegisterChainMaintainerRequest{})
	}
}
