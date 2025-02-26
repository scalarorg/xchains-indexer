package main

import (
	"log"

	models "github.com/scalarorg/data-models/indexer"
	dataUtil "github.com/scalarorg/data-models/util"
	"github.com/scalarorg/xchains-indexer/cmd"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	events "github.com/scalarorg/xchains-indexer/customs/scalar/events"
	"github.com/scalarorg/xchains-indexer/customs/scalar/messages"
	"github.com/scalarorg/xchains-indexer/indexer"
	"gorm.io/gorm/schema"
)

func main() {
	// Get buildin indexer for extending with custom indexers
	indexer := cmd.GetBuiltinIndexer()
	// Register db serializers
	schema.RegisterSerializer("HexSerializer", dataUtil.HexSerializer{})
	// Register the custom database models. They will be migrated and included in the database when the indexer finishes setup.
	registerCustomModels(indexer)
	// indexer.RegisterCustomModuleBasics([]module.AppModuleBasic{
	// 	&nexus.AppModuleBasic{},
	// })
	messages.ExtendMessagesIndexer(indexer)
	events.ExtendEventsIndexer(indexer)

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute. Err: %v", err)
	}
}

func registerCustomModels(indexer *indexer.Indexer) {
	//Cosmos
	indexer.RegisterCustomModels([]any{
		&models.Heartbeat{},
	})
	//Chains models
	indexer.RegisterCustomModels([]any{
		&models.ChainEventConfirmed{},
		&models.ChainEventCompleted{},
		&models.ChainEventFailed{},
		&models.ChainEventRetryFailed{},
		&models.CommandBatchSigned{},
		&models.CommandBatchAborted{},
		&models.ConfirmTokenStarted{},
		&models.ConfirmDepositStarted{},
		&models.ConfirmKeyTransferStarted{},
		&models.EventConfirmSourceTxs{},
		&models.EventTokenSent{},
		&models.EventContractCallApproved{},
		&models.EventContractCallWithMintApproved{},
		&models.ContractCallFailed{},
		&models.MintCommand{},
		&models.BurnCommand{},
		&models.PollCompleted{},
		&models.PollExpired{},
		&models.PollFailed{},
		&models.NoEventConfirmed{},
		&models.TokenConfirmation{},
	})
	//Nexus models
	indexer.RegisterCustomModels([]any{
		&models.RateLimitUpdated{},
		&models.MessageExecuted{},
	})
	//Covenant models
	indexer.RegisterCustomModels([]any{
		&models.SigningPsbtStarted{},
		&models.SigningPsbtCompleted{},
		&models.SigningPsbtExpired{},
		&models.TapScriptSigsSubmitted{},
		&models.KeyRotated{},
	})
	//Multisig models
	indexer.RegisterCustomModels([]any{
		&models.KeygenStarted{},
		&models.KeygenCompleted{},
		&models.KeygenExpired{},
		&models.KeyAssigned{},
		&models.MultiSigKeyRotated{},
		&models.SigningStarted{},
		&models.SigningCompleted{},
		&models.SigningExpired{},
		&models.SignatureSubmitted{},
		&models.PubkeySubmitted{},
		&models.KeygenOptOut{},
		&models.KeygenOptIn{},
	})
	//ScalarNet
	indexer.RegisterCustomModels([]any{
		&models.ContractCallSubmitted{},
		&models.ContractCallWithTokenSubmitted{},
		&models.TokenSent{},
		&models.FeePaid{},
		&models.ScalarTransferCompleted{},
		&models.FeeCollected{},
	})
	//Vote\
	indexer.RegisterCustomModels([]any{
		&models.Voted{},
	})
	//Messages
	indexer.RegisterCustomModels([]any{
		&common.TxMessage{},
	})
}
