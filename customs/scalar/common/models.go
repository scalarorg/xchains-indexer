package common

import "github.com/scalarorg/xchains-indexer/db/models"

type MsgType int

const (
	MsgRecvPacket MsgType = iota
	MsgAcknowledgement
)

type TxMessage struct {
	ID            uint   `gorm:"primaryKey"`
	TxID          uint   `gorm:"uniqueIndex:txMessageIndex,priority:1"`
	MessageID     uint   `gorm:"uniqueIndex:txMessageIndex,priority:2"`
	MessageType   string `gorm:"type:varchar(128)" json:"message_type"`
	BlockId       uint
	Sender        string `gorm:"type:varchar(64)" json:"sender"`
	Chain         string `gorm:"type:varchar(64)" json:"chain"`
	MessageDetail string
	Tx            models.Tx `gorm:"foreignKey:TxID;references:ID"`
}
