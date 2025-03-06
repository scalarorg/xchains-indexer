package multisig

import (
	"encoding/json"
	"errors"

	stdTypes "github.com/cosmos/cosmos-sdk/types"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	"github.com/scalarorg/xchains-indexer/config"
	txTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/parsers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RotateKeyRequestParser struct {
	common.BaseMessageParser
}

func (p *RotateKeyRequestParser) ParseMessage(sdkMsg stdTypes.Msg, messageLog *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := sdkMsg.(*multisigTypes.RotateKeyRequest)
	if !ok {
		return nil, errors.New("not a RotateKeyRequest")
	}

	var parsedValue any = &MultisigKeyMsg{
		Type:   MSG_SCALAR_MULTISIG_ROTATE_KEY_REQUEST,
		Sender: parsedMsg.Sender.String(),
		Chain:  string(parsedMsg.Chain),
		KeyID:  parsedMsg.KeyID.String(),
	}

	return &parsedValue, nil
}

func (p *RotateKeyRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	config.Log.Debugf("RotateKeyRequestParser# IndexMessage# message: %++v, dataset: %T", message, *dataset)
	parsedMsg, ok := (*dataset).(*MultisigKeyMsg)
	if !ok {
		return errors.New("failed to cast dataset to RotateKeyMsg")
	}
	err := db.Create(&parsedMsg).Error
	if err != nil {
		config.Log.Debugf("RotateKeyRequestParser# Failed to save message %v", err)
	}
	jsonValue, err := json.Marshal(parsedMsg)
	if err == nil {
		txMessage := common.TxMessage{
			Tx:            message.Tx,
			TxID:          message.TxID,
			MessageID:     message.ID,
			MessageType:   parsedMsg.Type,
			BlockId:       message.Tx.BlockID,
			MessageDetail: string(jsonValue),
		}

		err = db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "tx_id"}, {Name: "message_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"message_detail"}),
		}).Create(&txMessage).Error
		if err != nil {
			config.Log.Debugf("RotateKeyRequestParser# Failed to save message event %v", err)
		}
	} else {
		config.Log.Debugf("RotateKeyRequestParser# Failed to marshal message event %v", err)
		return err
	}
	return nil
}
