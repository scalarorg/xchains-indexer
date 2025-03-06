package chains

import (
	"encoding/json"
	"errors"

	stdTypes "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/xchains-indexer/config"
	txTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/parsers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateDeployTokenRequestParser struct {
	common.BaseMessageParser
}

func (p *CreateDeployTokenRequestParser) ParseMessage(sdkMsg stdTypes.Msg, messageLog *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := sdkMsg.(*chainsTypes.CreateDeployTokenRequest)
	if !ok {
		return nil, errors.New("not a CreateDeployTokenRequest")
	}

	var parsedValue any = &CreateDeployTokenMsg{
		Type:         MSG_SCALAR_CHAINS_CREATE_DEPLOY_TOKEN_REQUEST,
		Sender:       parsedMsg.Sender.String(),
		Chain:        string(parsedMsg.Chain),
		TokenSymbol:  string(parsedMsg.TokenSymbol),
		AliasedToken: string(parsedMsg.AliasedTokenName),
		Address:      parsedMsg.Address.String(),
	}

	return &parsedValue, nil
}

func (p *CreateDeployTokenRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	config.Log.Debugf("CreateDeployTokenRequestParser# IndexMessage# message: %++v, dataset: %T", message, *dataset)
	parsedMsg, ok := (*dataset).(*CreateDeployTokenMsg)
	if !ok {
		return errors.New("failed to cast dataset to CreateDeployTokenMsg")
	}
	err := db.Create(&parsedMsg).Error
	if err != nil {
		config.Log.Debugf("CreateDeployTokenRequestParser# Failed to save message %v", err)
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
			config.Log.Debugf("CreateDeployTokenRequestParser# Failed to save message event %v", err)
		}
	} else {
		config.Log.Debugf("CreateDeployTokenRequestParser# Failed to marshal message event %v", err)
		return err
	}
	return nil
}
