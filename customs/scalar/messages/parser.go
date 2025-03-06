package messages

import (
	"log"

	"github.com/cosmos/cosmos-sdk/codec"
	gogoprototypes "github.com/gogo/protobuf/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covenantExported "github.com/scalarorg/scalar-core/x/covenant/exported"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/chains"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/covenant"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/multisig"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/reward"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages/tss"
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

	chains.ExtendMessagesIndexerChains(instance)
	covenant.ExtendMessagesIndexerCovenant(instance)
	multisig.ExtendMessagesIndexerMultisig(instance)
	reward.ExtendMessagesIndexerReward(instance)
	tss.ExtendMessagesIndexerTss(instance)
	vote.ExtendMessagesIndexerVote(instance)

	extendMessagesIndexerTokens(instance)

	instance.PostSetupCustomFunction = func(dataset indexer.PostSetupCustomDataset) error {
		chains.PostSetupCustomFunctionChains(instance, &dataset)
		covenant.PostSetupCustomFunctionCovenant(instance, &dataset)
		multisig.PostSetupCustomFunctionMultisig(instance, &dataset)
		reward.PostSetupCustomFunctionReward(instance, &dataset)
		tss.PostSetupCustomFunctionTss(instance, &dataset)
		vote.PostSetupCustomFunctionVote(instance, &dataset)
		instance.ChainClient.Codec.InterfaceRegistry.RegisterImplementations((*codec.ProtoMarshaler)(nil),
			&gogoprototypes.BoolValue{},
			&chainsTypes.SigMetadata{},
			&chainsTypes.Event{},
			&chainsTypes.VoteEvents{},
			&chainsTypes.PollMetadata{},
			&covenantTypes.PsbtMultiSig{},
			&covenantExported.TapScriptSigsList{},
			&multisigTypes.MultiSig{},
		)
		return nil
	}
	return nil
}
