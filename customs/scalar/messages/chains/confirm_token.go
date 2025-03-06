package chains

import (
	"encoding/json"
	"errors"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/xchains-indexer/config"
	indexerTxTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	common "github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/parsers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type ConfirmTokenRequestParser struct {
	common.BaseMessageParser
}

func (p *ConfirmTokenRequestParser) ParseMessage(sdkMsg sdkTypes.Msg, messageLog *indexerTxTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	config.Log.Debugf("[ConfirmTokenRequestParser] ParseMessage# msg: %++v", sdkMsg)
	parsedMsg, ok := sdkMsg.(*chainsTypes.ConfirmTokenRequest)
	if !ok {
		return nil, errors.New("not a ConfirmTokenRequest message request")
	}
	var parsedMessageValue any = &ConfirmTokenValueMsg{
		Type:        MSG_SCALAR_CHAINS_CONFIRM_TOKEN_REQUEST,
		Sender:      parsedMsg.Sender.String(),
		Chain:       string(parsedMsg.Chain),
		TxID:        parsedMsg.TxID.Hex(),
		AssetChain:  string(parsedMsg.Asset.Chain),
		AssetSymbol: parsedMsg.Asset.Symbol,
	}
	return &parsedMessageValue, nil
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (p *ConfirmTokenRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	config.Log.Debugf("ConfirmTokenRequestParser# IndexMessage# message: %++v, dataset: %T", message, *dataset)
	parsedMsg, ok := (*dataset).(*ConfirmTokenValueMsg)
	if !ok {
		return errors.New("failed to cast dataset to ConfirmTokenValueMsg")
	}
	err := db.Create(&parsedMsg).Error
	if err != nil {
		config.Log.Debugf("ConfirmTokenRequestParser# Failed to save message %v", err)
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
			config.Log.Debugf("ConfirmTokenRequestParser# Failed to save message event %v", err)
		}
	} else {
		config.Log.Debugf("ConfirmTokenRequestParser# Failed to marshal message event %v", err)
		return err
	}

	return err
}
