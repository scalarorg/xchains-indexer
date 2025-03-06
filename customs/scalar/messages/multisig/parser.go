package multisig

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	rewardExported "github.com/scalarorg/scalar-core/x/reward/exported"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	MSG_SCALAR_MULTISIG_FILTER               string = "^/scalar.multisig.v1beta1.*"
	MSG_SCALAR_MULTISIG_START_KEYGEN_REQUEST string = "/scalar.multisig.v1beta1.StartKeygenRequest"
	MSG_SCALAR_MULTISIG_ROTATE_KEY_REQUEST   string = "/scalar.multisig.v1beta1.RotateKeyRequest"
)

func ExtendMessagesIndexerMultisig(instance *indexer.Indexer) {
	customParsers := []parsers.MessageParser{}
	customParsers = append(customParsers, &StartKeygenRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_MULTISIG_START_KEYGEN_REQUEST,
			Indexer: instance,
		},
	})
	customParsers = append(customParsers, &RotateKeyRequestParser{
		BaseMessageParser: common.BaseMessageParser{
			Id:      MSG_SCALAR_MULTISIG_ROTATE_KEY_REQUEST,
			Indexer: instance,
		},
	})
	for _, parser := range customParsers {
		instance.RegisterCustomMessageParser(parser.Identifier(), parser)
	}
}

func PostSetupCustomFunctionMultisig(instance *indexer.Indexer, dataset *indexer.PostSetupCustomDataset) {
	dataset.DB.AutoMigrate(
		&MultisigKeyMsg{},
	)
	if instance.ChainClient != nil {
		instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*rewardExported.Refundable)(nil), &multisigTypes.SubmitPubKeyRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*rewardExported.Refundable)(nil), &multisigTypes.SubmitSignatureRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_MULTISIG_START_KEYGEN_REQUEST, (*sdk.Msg)(nil), &multisigTypes.StartKeygenRequest{})
		instance.ChainClient.Codec.InterfaceRegistry.RegisterInterface(MSG_SCALAR_MULTISIG_ROTATE_KEY_REQUEST, (*sdk.Msg)(nil), &multisigTypes.RotateKeyRequest{})
	}
}
