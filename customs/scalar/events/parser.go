package events

import (
	"github.com/scalarorg/xchains-indexer/indexer"
	"github.com/scalarorg/xchains-indexer/parsers"
)

const (
	EVENT_TYPE_MESSAGE                                string = "message"
	EVENT_TYPE_EVENT_TOKEN_SENT                       string = "scalar.chains.v1beta1.EventTokenSent"
	EVENT_TYPE_CONTRACT_CALL_APPROVED                 string = "scalar.chains.v1beta1.ContractCallApproved"
	EVENT_TYPE_CONTRACT_CALL_FAILED                   string = "scalar.chains.v1beta1.ContractCallFailed"
	EVENT_TYPE_EVENT_CONTRACT_CALL_WITH_MINT_APPROVED string = "scalar.chains.v1beta1.EventContractCallWithMintApproved"
	EVENT_TYPE_POLL_COMPLETED                         string = "scalar.chains.v1beta1.PollCompleted"
	EVENT_TYPE_POLL_EXPIRED                           string = "scalar.chains.v1beta1.PollExpired"
	EVENT_TYPE_POLL_FAILED                            string = "scalar.chains.v1beta1.PollFailed"
	EVENT_TYPE_NO_EVENT_CONFIRMED                     string = "scalar.chains.v1beta1.NoEventConfirmed"
	EVENT_TYPE_CHAIN_EVENT_CONFIRMED                  string = "scalar.chains.v1beta1.ChainEventConfirmed"
	EVENT_TYPE_CHAIN_EVENT_COMPLETED                  string = "scalar.chains.v1beta1.ChainEventCompleted"
	EVENT_TYPE_CHAIN_EVENT_FAILED                     string = "scalar.chains.v1beta1.ChainEventFailed"
	EVENT_TYPE_CHAIN_EVENT_RETRY_FAILED               string = "scalar.chains.v1beta1.ChainEventRetryFailed"
	EVENT_TYPE_CONFIRM_TOKEN_STARTED                  string = "scalar.chains.v1beta1.ConfirmTokenStarted"
	EVENT_TYPE_CONFIRM_DEPOSIT_STARTED                string = "scalar.chains.v1beta1.ConfirmDepositStarted"
	EVENT_TYPE_CONFIRM_SOURCE_TXS                     string = "scalar.chains.v1beta1.EventConfirmSourceTxsStarted"
	EVENT_TYPE_CONFIRM_KEY_TRANSFER_STARTED           string = "scalar.chains.v1beta1.ConfirmKeyTransferStarted"
	EVENT_TYPE_COMMAND_BATCH_SIGNED                   string = "scalar.chains.v1beta1.CommandBatchSigned"
	EVENT_TYPE_COMMAND_BATCH_ABORTED                  string = "scalar.chains.v1beta1.CommandBatchAborted"
	EVENT_TYPE_MINT_COMMAND                           string = "scalar.chains.v1beta1.MintCommand"
	EVENT_TYPE_BURN_COMMAND                           string = "scalar.chains.v1beta1.BurnCommand"

	EVENT_TYPE_SIGNING_PSBT_STARTED      string = "scalar.covenant.v1beta1.SigningPsbtStarted"
	EVENT_TYPE_SIGNING_PSBT_COMPLETED    string = "scalar.covenant.v1beta1.SigningPsbtCompleted"
	EVENT_TYPE_SIGNING_PSBT_EXPIRED      string = "scalar.covenant.v1beta1.SigningPsbtExpired"
	EVENT_TYPE_TAP_SCRIPT_SIGS_SUBMITTED string = "scalar.covenant.v1beta1.TapScriptSigsSubmitted"
	EVENT_TYPE_KEY_ROTATED               string = "scalar.covenant.v1beta1.KeyRotated"

	EVENT_TYPE_KEY_ASSIGNED         string = "scalar.multisig.v1beta1.KeyAssigned"
	EVENT_TYPE_MILTISIG_KEY_ROTATED string = "scalar.multisig.v1beta1.KeyRotated"
	EVENT_TYPE_KEYGEN_STARTED       string = "scalar.multisig.v1beta1.KeygenStarted"
	EVENT_TYPE_KEYGEN_COMPLETED     string = "scalar.multisig.v1beta1.KeygenCompleted"
	EVENT_TYPE_KEYGEN_EXPIRED       string = "scalar.multisig.v1beta1.KeygenExpired"
	EVENT_TYPE_SIGNING_STARTED      string = "scalar.multisig.v1beta1.SigningStarted"
	EVENT_TYPE_SIGNING_COMPLETED    string = "scalar.multisig.v1beta1.SigningCompleted"
	EVENT_TYPE_SIGNING_EXPIRED      string = "scalar.multisig.v1beta1.SigningExpired"
	EVENT_TYPE_SIGNATURE_SUBMITTED  string = "scalar.multisig.v1beta1.SignatureSubmitted"
	EVENT_TYPE_PUBKEY_SUBMITTED     string = "scalar.multisig.v1beta1.PubkeySubmitted"
	EVENT_TYPE_KEYGEN_OPT_OUT       string = "scalar.multisig.v1beta1.KeygenOptOut"
	EVENT_TYPE_KEYGEN_OPT_IN        string = "scalar.multisig.v1beta1.KeygenOptIn"

	EVENT_TYPE_MESSAGE_PROCESSING string = "scalar.nexus.v1beta1.MessageProcessing"
	EVENT_TYPE_MESSAGE_EXECUTED   string = "scalar.nexus.v1beta1.MessageExecuted"
	EVENT_TYPE_RATE_LIMIT_UPDATED string = "scalar.nexus.v1beta1.RateLimitUpdated"

	EVENT_TYPE_CONTRACT_CALL_SUBMITTED            string = "scalar.scalarnet.v1beta1.ContractCallSubmitted"
	EVENT_TYPE_CONTRACT_CALL_WITH_TOKEN_SUBMITTED string = "scalar.scalarnet.v1beta1.ContractCallWithTokenSubmitted"
	EVENT_TYPE_TOKEN_SENT                         string = "scalar.scalarnet.v1beta1.TokenSent"
	EVENT_TYPE_FEE_PAID                           string = "scalar.scalarnet.v1beta1.FeePaid"

	EVENT_TYPE_VOTED string = "scalar.vote.v1beta1.Voted"
)

func ExtendEventsIndexer(indexer *indexer.Indexer) error {
	// blockEventParsers := []parsers.BlockEventParser{
	// 	&MessageEventParser{
	// 		BaseParser: BaseParser{
	// 			Id:      EVENT_TYPE_MESSAGE,
	// 			Indexer: indexer,
	// 		},
	// 	},
	// }
	// for _, parser := range blockEventParsers {
	// 	indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	// }
	registerBlockEventParsersCosmos(indexer)
	registerBlockEventParsersNexus(indexer)
	registerBlockEventParsersCovenant(indexer)
	registerBlockEventParsersChains(indexer)
	registerBlockEventParsersMultisig(indexer)
	registerBlockEventParsersScalarnet(indexer)
	registerBlockEventParsersVoted(indexer)
	return nil
}
func registerBlockEventParsersCosmos(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		&HeartBeatParser{
			BaseParser: BaseParser{
				Id:      "heartbeat",
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}
func registerBlockEventParsersChains(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		&TokenConfirmationParser{
			BaseParser: BaseParser{
				Id:      "tokenConfirmation",
				Indexer: indexer,
			},
		},
		&ChainEventConfirmedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CHAIN_EVENT_CONFIRMED,
				Indexer: indexer,
			},
		},
		&ChainEventCompletedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CHAIN_EVENT_COMPLETED,
				Indexer: indexer,
			},
		},
		&ChainEventFailedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CHAIN_EVENT_FAILED,
				Indexer: indexer,
			},
		},
		&ChainEventRetryFailedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CHAIN_EVENT_RETRY_FAILED,
				Indexer: indexer,
			},
		},
		&CommandBatchSignedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_COMMAND_BATCH_SIGNED,
				Indexer: indexer,
			},
		},
		&CommandBatchAbortedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_COMMAND_BATCH_ABORTED,
				Indexer: indexer,
			},
		},
		&MintCommandParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_MINT_COMMAND,
				Indexer: indexer,
			},
		},
		&BurnCommandParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_BURN_COMMAND,
				Indexer: indexer,
			},
		},
		&ConfirmTokenStartedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONFIRM_TOKEN_STARTED,
				Indexer: indexer,
			},
		},
		&ConfirmDepositStartedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONFIRM_DEPOSIT_STARTED,
				Indexer: indexer,
			},
		},
		&ConfirmKeyTransferStartedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONFIRM_KEY_TRANSFER_STARTED,
				Indexer: indexer,
			},
		},
		&EventConfirmSourceTxsParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONFIRM_SOURCE_TXS,
				Indexer: indexer,
			},
		},
		&EventTokenSentParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_EVENT_TOKEN_SENT,
				Indexer: indexer,
			},
		},
		&ContractCallApprovedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONTRACT_CALL_APPROVED,
				Indexer: indexer,
			},
		},
		&ContractCallFailedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONTRACT_CALL_FAILED,
				Indexer: indexer,
			},
		},
		&PollCompletedEventParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_POLL_COMPLETED,
				Indexer: indexer,
			},
		},
		&PollExpiredEventParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_POLL_EXPIRED,
				Indexer: indexer,
			},
		},
		&PollFailedEventParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_POLL_FAILED,
				Indexer: indexer,
			},
		},
		&NoEventConfirmedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_NO_EVENT_CONFIRMED,
				Indexer: indexer,
			},
		},
		&EventContractCallWithMintApprovedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_EVENT_CONTRACT_CALL_WITH_MINT_APPROVED,
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}

func registerBlockEventParsersMultisig(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		&KeygenStartedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEYGEN_STARTED,
				Indexer: indexer,
			},
		},
		&KeygenCompletedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEYGEN_COMPLETED,
				Indexer: indexer,
			},
		},
		&KeygenExpiredParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEYGEN_EXPIRED,
				Indexer: indexer,
			},
		},
		&KeyAssignedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEY_ASSIGNED,
				Indexer: indexer,
			},
		},
		&MultiSigKeyRotatedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_MILTISIG_KEY_ROTATED,
				Indexer: indexer,
			},
		},
		&SigningStartedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNING_STARTED,
				Indexer: indexer,
			},
		},
		&SigningCompletedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNING_COMPLETED,
				Indexer: indexer,
			},
		},
		&SigningExpiredParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNING_EXPIRED,
				Indexer: indexer,
			},
		},
		&SignatureSubmittedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNATURE_SUBMITTED,
				Indexer: indexer,
			},
		},
		&PubkeySubmittedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_PUBKEY_SUBMITTED,
				Indexer: indexer,
			},
		},
		&KeygenOptInParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEYGEN_OPT_IN,
				Indexer: indexer,
			},
		},
		&KeygenOptOutParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEYGEN_OPT_OUT,
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}

func registerBlockEventParsersCovenant(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		&SigningPsbtStartedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNING_PSBT_STARTED,
				Indexer: indexer,
			},
		},
		&SigningPsbtCompletedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNING_PSBT_COMPLETED,
				Indexer: indexer,
			},
		},
		&SigningPsbtExpiredParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_SIGNING_PSBT_EXPIRED,
				Indexer: indexer,
			},
		},
		&TapScriptSigsSubmittedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_TAP_SCRIPT_SIGS_SUBMITTED,
				Indexer: indexer,
			},
		},
		&KeyRotatedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_KEY_ROTATED,
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}

func registerBlockEventParsersScalarnet(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		&ContractCallSubmittedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONTRACT_CALL_SUBMITTED,
				Indexer: indexer,
			},
		},
		&ContractCallWithTokenSubmittedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_CONTRACT_CALL_WITH_TOKEN_SUBMITTED,
				Indexer: indexer,
			},
		},
		&TokenSentParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_TOKEN_SENT,
				Indexer: indexer,
			},
		},
		&FeePaidParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_FEE_PAID,
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}

func registerBlockEventParsersNexus(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		// &MessageProcessingEventParser{
		// 	BaseParser: BaseParser{
		// 		Id:      EVENT_TYPE_MESSAGE_PROCESSING,
		// 		Indexer: indexer,
		// 	},
		// },
		// &MessageExecutedEventParser{
		// 	BaseParser: BaseParser{
		// 		Id:      EVENT_TYPE_MESSAGE_EXECUTED,
		// 		Indexer: indexer,
		// 	},
		// },
		&RateLimitUpdatedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_RATE_LIMIT_UPDATED,
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}

func registerBlockEventParsersVoted(indexer *indexer.Indexer) {
	parsers := []parsers.BlockEventParser{
		&VotedParser{
			BaseParser: BaseParser{
				Id:      EVENT_TYPE_VOTED,
				Indexer: indexer,
			},
		},
	}
	for _, parser := range parsers {
		indexer.RegisterCustomBlockEventParser(parser.Identifier(), parser)
	}
}
