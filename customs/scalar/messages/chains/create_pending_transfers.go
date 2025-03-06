package chains

import (
	"encoding/json"
	"errors"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/xchains-indexer/config"
	indexerTxTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/parsers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreatePendingTransfersParser struct {
	common.BaseMessageParser
}

func (p *CreatePendingTransfersParser) ParseMessage(sdkMsg sdkTypes.Msg, messageLog *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := sdkMsg.(*chainsTypes.CreatePendingTransfersRequest)
	if !ok {
		return nil, errors.New("not a CreatePendingTransfersRequest message request")
	}
	var parsedValue any = &CreatePendingTransfersMsg{
		Type:   MSG_SCALAR_CHAINS_CREATE_PENDING_TRANSFERS_REQUEST,
		Sender: parsedMsg.Sender.String(),
		Chain:  string(parsedMsg.Chain),
	}
	return &parsedValue, nil
}

func (p *CreatePendingTransfersParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	config.Log.Debugf("CreatePendingTransfersParser# IndexMessage# message: %++v, dataset: %T", message, *dataset)
	parsedMsg, ok := (*dataset).(*CreatePendingTransfersMsg)
	if !ok {
		return errors.New("failed to cast dataset to CreatePendingTransfersMsg")
	}
	// err := db.Create(&parsedMsg).Error
	// if err != nil {
	// 	config.Log.Debugf("CreatePendingTransfersParser# Failed to save message %v", err)
	// }
	jsonValue, err := json.Marshal(parsedMsg)
	if err == nil {
		txMessage := common.TxMessage{
			Tx:            message.Tx,
			TxID:          message.TxID,
			MessageID:     message.ID,
			BlockId:       message.Tx.BlockID,
			MessageType:   parsedMsg.Type,
			Sender:        parsedMsg.Sender,
			Chain:         parsedMsg.Chain,
			MessageDetail: string(jsonValue),
		}

		err = db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_id"}, {Name: "message_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"message_detail"}),
		}).Create(&txMessage).Error
		if err != nil {
			config.Log.Debugf("CreatePendingTransfersParser# Failed to save message event %v", err)
		}
	} else {
		config.Log.Debugf("CreatePendingTransfersParser# Failed to marshal message event %v", err)
		return err
	}
	return nil
}
