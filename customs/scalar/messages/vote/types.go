package vote

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/vote/exported"
)

type VoteRequestMsg struct {
	Type    string              `gorm:"type:varchar(128)" json:"type"`
	Sender  sdkTypes.AccAddress `gorm:"type:varchar(64)" json:"sender"`
	PollID  exported.PollID     `gorm:"type:varchar(64)" json:"poll_id"`
	Vote    *types.Any          `gorm:"type:jsonb" json:"vote"`
	VoteMsg any                 `gorm:"type:jsonb" json:"vote_msg"`
}
