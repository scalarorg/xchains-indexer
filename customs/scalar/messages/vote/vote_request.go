package vote

import (
	"encoding/json"
	"errors"
	"fmt"

	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	stdTypes "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/vote/types"
	"github.com/scalarorg/xchains-indexer/config"
	txTypes "github.com/scalarorg/xchains-indexer/cosmos/modules/tx"
	"github.com/scalarorg/xchains-indexer/customs/scalar/common"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/parsers"
	"github.com/scalarorg/xchains-indexer/probe/client"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// This defines the custom message parser for the call contract approve message type
// It implements the MessageParser interface
type VoteRequestParser struct {
	common.BaseMessageParser
}

func (p *VoteRequestParser) ParseMessage(cosmosMsg stdTypes.Msg, logMsg *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	parsedMsg, ok := cosmosMsg.(*types.VoteRequest)
	if !ok {
		return nil, errors.New("not a Vote request")
	}
	// config.Log.Debugf("VoteRequest# PollId %d, Sender %s, Vote %+v", parsedMsg.PollID, parsedMsg.Sender, parsedMsg.Vote)
	var parsedValue any = &VoteRequestMsg{
		Type:   MSG_SCALAR_VOTE_REQUEST,
		PollID: parsedMsg.PollID,
		Sender: parsedMsg.Sender,
		Vote:   parsedMsg.Vote,
	}
	if parsedMsg.Vote != nil {
		msg, err := common.ParseInnerMessage(p.Indexer.ChainClient.Codec, parsedMsg.Vote, p.Indexer.CustomMessageParserRegistry, logMsg, cfg)
		if err == nil && msg != nil {
			config.Log.Debugf("VoteRequest# Success Parsed inner message")
			parsedValue.(*VoteRequestMsg).VoteMsg = *msg
		} else {
			msg, err = p.ParseVoteEvents(p.Indexer.ChainClient.Codec, parsedMsg.Vote, logMsg, cfg)
			if err == nil {
				config.Log.Debugf("VoteRequest# Successfully parsed VoteEvents")
				parsedValue.(*VoteRequestMsg).VoteMsg = *msg
			} else {
				config.Log.Debugf("VoteRequest# Failed to parse inner message")
			}
		}
	}
	return &parsedValue, nil
}
func (p *VoteRequestParser) ParseVoteEvents(codec client.Codec, vote *codecTypes.Any, logMsg *txTypes.LogMessage, cfg config.IndexConfig) (*any, error) {
	var events chainsTypes.VoteEvents
	err := codec.InterfaceRegistry.UnpackAny(vote, &events)
	if err == nil {
		config.Log.Debugf("Successfully Unpacked message of type %s", vote.TypeUrl)
		config.Log.Debugf("VoteEvents# Events %v", events)
	} else {
		config.Log.Debug(fmt.Sprintf("ParseVoteEvents error: %v", err))
	}
	return nil, err
}

// This method is called during database insertion. It is responsible for storing the parsed data in the database.
// The gorm db is wrapped in a transaction, so any errors will cause a rollback.
// Any errors returned will be saved as a parser error in the database as well for later debugging.
func (p *VoteRequestParser) IndexMessage(dataset *any, db *gorm.DB, message models.Message, messageEvents []parsers.MessageEventWithAttributes, cfg config.IndexConfig) error {
	config.Log.Debugf("VoteRequestParser# IndexMessage# message: %++v, dataset: %T", message, *dataset)
	parsedMsg, ok := (*dataset).(*VoteRequestMsg)
	if !ok {
		return errors.New("failed to cast dataset to VoteRequestEvent")
	}
	err := db.Create(&parsedMsg).Error
	if err != nil {
		config.Log.Debugf("StartKeygenRequestParser# Failed to save message %v", err)
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
			config.Log.Debugf("VoteRequestParser# Failed to save message event %v", err)
		}
	} else {
		config.Log.Debugf("VoteRequestParser# Failed to marshal message event %v", err)
		return err
	}

	return nil
}
