package chains

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_CHAINS_FILTER                           string = "^/scalar.chains.v1beta1.*"
	MSG_SCALAR_CHAINS_CREATE_DEPLOY_TOKEN_REQUEST      string = "/scalar.chains.v1beta1.CreateDeployTokenRequest"
	MSG_SCALAR_CHAINS_CONFIRM_TOKEN_REQUEST            string = "/scalar.chains.v1beta1.ConfirmTokenRequest"
	MSG_SCALAR_CHAINS_SIGN_COMMANDS_REQUEST            string = "/scalar.chains.v1beta1.SignCommandsRequest"
	MSG_SCALAR_CHAINS_SIGN_BTC_COMMANDS_REQUEST        string = "/scalar.chains.v1beta1.SignBtcCommandsRequest"
	MSG_SCALAR_CHAINS_SIGN_PSBT_COMMANDS_REQUEST       string = "/scalar.chains.v1beta1.SignPsbtCommandsRequest"
	MSG_SCALAR_CHAINS_CONFIRM_SOURCE_TXS_REQUEST       string = "/scalar.chains.v1beta1.ConfirmSourceTxsRequest"
	MSG_SCALAR_CHAINS_CREATE_PENDING_TRANSFERS_REQUEST string = "/scalar.chains.v1beta1.CreatePendingTransfersRequest"
	//MSG_SOURCE_TXS_REQUEST   string = "ConfirmSourceTxsRequest"
)

func ExtendMessagesIndexerChains(instance *indexer.Indexer) {
	customParsers := []parsers.MessageParser{}
	customParsers = append(customParsers, &CreateDeployTokenRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_CREATE_DEPLOY_TOKEN_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &ConfirmTokenRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_CONFIRM_TOKEN_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &SignCommandsParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_SIGN_COMMANDS_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &SignBtcCommandsParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_SIGN_BTC_COMMANDS_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &SignPsbtCommandsParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_SIGN_PSBT_COMMANDS_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &ConfirmSourceTxsParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_CONFIRM_SOURCE_TXS_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &CreatePendingTransfersParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_CHAINS_CREATE_PENDING_TRANSFERS_REQUEST,
			Indexer: instance,
		},
	})
	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionChains(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate(
		&CreateDeployTokenMsg{},
		&ConfirmTokenValueMsg{},
		&ConfirmSourceTxsMsg{},
		// &CreatePendingTransfersMsg{},
		// &SignCommandsMsg{},
		// &SignBtcCommandsMsg{},
		// &SignPsbtCommandsMsg{},

	)
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_CREATE_DEPLOY_TOKEN_REQUEST, (*sdk.Msg)(nil), &chainsTypes.CreateDeployTokenRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_CONFIRM_TOKEN_REQUEST, (*sdk.Msg)(nil), &chainsTypes.ConfirmTokenRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_SIGN_COMMANDS_REQUEST, (*sdk.Msg)(nil), &chainsTypes.SignCommandsRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_SIGN_BTC_COMMANDS_REQUEST, (*sdk.Msg)(nil), &chainsTypes.SignBtcCommandsRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_SIGN_PSBT_COMMANDS_REQUEST, (*sdk.Msg)(nil), &chainsTypes.SignPsbtCommandRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_CONFIRM_SOURCE_TXS_REQUEST, (*sdk.Msg)(nil), &chainsTypes.ConfirmSourceTxsRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_CHAINS_CREATE_PENDING_TRANSFERS_REQUEST, (*sdk.Msg)(nil), &chainsTypes.CreatePendingTransfersRequest{})
	}
}
