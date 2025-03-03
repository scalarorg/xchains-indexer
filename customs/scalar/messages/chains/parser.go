package chains

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
	chainsTypes "github.com/scalarorg/xchains-indexer/x/chains/types"
)

const (
	MSG_SCALAR_CHAINS_FILTER                string = "^/scalar.chains.v1beta1.*"
	MSG_SCALAR_CHAINS_CONFIRM_TOKEN_REQUEST string = "/scalar.chains.v1beta1.ConfirmTokenRequest"
	//MSG_SOURCE_TXS_REQUEST   string = "ConfirmSourceTxsRequest"
)

func ExtendMessagesIndexerChains(instance *indexer.Indexer) {
	customParsers := make(map[string]parsers.MessageParser)
	for key, parser := range customParsers {
		instance.RegisterCustomMessageParser(key, parser)
	}
}

func PostSetupCustomFunctionChains(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate()
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_CONFIRM_TOKEN_REQUEST, (*sdk.Msg)(nil), &chainsTypes.ConfirmTokenRequest{})
		// instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(evmTypes.MSG_EVM_VOTE_EVENTS, (*sdk.Msg)(nil), &evmTypes.VoteEvents{})
		// instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(nexusTypes.MSG_NEXUS_REGISTER_CHAIN_MAINTAINER_REQUEST, (*sdk.Msg)(nil), &nexusTypes.RegisterChainMaintainerRequest{})
	}
}
