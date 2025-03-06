package reward

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rewardExported "github.com/scalarorg/scalar-core/x/reward/exported"
	rewardTypes "github.com/scalarorg/scalar-core/x/reward/types"
	common "github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_REWARD_CHAINS_FILTER      string = "^/scalar.reward.v1beta1.*"
	MSG_SCALAR_REWARD_REFUNDABLE         string = "reward.v1beta1.Refundable"
	MSG_SCALAR_REWARD_REFUND_MSG_REQUEST string = "/scalar.reward.v1beta1.RefundMsgRequest"
)

func ExtendMessagesIndexerReward(instance *indexer.Indexer) {
	customParsers := []parsers.MessageParser{}
	// Extend RefundMsgRequest parser
	customParsers = append(customParsers, &RefundMsgRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_REWARD_REFUND_MSG_REQUEST,
			Indexer: instance,
		},
	})
	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionReward(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate()
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_REWARD_REFUNDABLE, (*rewardExported.Refundable)(nil), &rewardTypes.RefundMsgRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_REWARD_REFUND_MSG_REQUEST, (*sdk.Msg)(nil), &rewardTypes.RefundMsgRequest{})
		// instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*rewardExported.Refundable)(nil), &rewardTypes.RefundMsgRequest{})
	}
}
