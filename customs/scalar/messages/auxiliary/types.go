package auxiliary

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
)

type BatchRequestMsg struct {
	Type     string      `gorm:"type:varchar(128)" json:"type"`
	Sender   string      `gorm:"type:varchar(64)" json:"sender"`
	Messages []types.Any `gorm:"type:json" json:"messages"`
}
