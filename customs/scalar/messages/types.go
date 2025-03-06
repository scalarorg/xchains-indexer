package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainTypes "github.com/scalarorg/scalar-core/x/chains/types"
	github_com_axelarnetwork_axelar_core_x_nexus_exported "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/vote/exported"
)

type RefundMsgRequestValue struct {
	Type         string         `json:"@type,omitempty"`
	Sender       sdk.AccAddress `json:"sender,omitempty"`
	InnerMessage any            `json:"inner_message,omitempty"`
}

type VoteRequestValue struct {
	Type   string          `json:"@type,omitempty"`
	Sender sdk.AccAddress  `json:"sender,omitempty"`
	PollID exported.PollID `json:"poll_id,omitempty"`
	Vote   any             `json:"vote,omitempty"`
}

type VoteEvents struct {
	Type   string                                                          `json:"@type,omitempty"`
	Chain  github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `json:"chain,omitempty"`
	Events []Event                                                         `json:"events,omitempty"`
}

type RegisterChainMaintainerRequestValue struct {
	Type   string                                                            `json:"@type,omitempty"`
	Sender sdk.AccAddress                                                    `json:"sender,omitempty"`
	Chains []github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `json:"chains,omitempty"`
}

// type Event struct {
// 	Chain  github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/axelarnetwork/axelar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
// 	TxID   Hash                                                            `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
// 	Index  uint64                                                          `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
// 	Status Event_Status                                                    `protobuf:"varint,4,opt,name=status,proto3,enum=axelar.evm.v1beta1.Event_Status" json:"status,omitempty"`
// 	// Types that are valid to be assigned to Event:
// 	//	*Event_TokenSent
// 	//	*Event_ContractCall
// 	//	*Event_ContractCallWithToken
// 	//	*Event_Transfer
// 	//	*Event_TokenDeployed
// 	//	*Event_MultisigOwnershipTransferred
// 	//	*Event_MultisigOperatorshipTransferred
// 	Event isEvent_Event `protobuf_oneof:"event"`
// }

type EventBase struct {
	Chain  github_com_axelarnetwork_axelar_core_x_nexus_exported.ChainName `json:"chain,omitempty"`
	Index  uint64                                                          `json:"index,omitempty"`
	Status string                                                          `json:"status,omitempty"`
	TxID   chainTypes.Hash                                                 `json:"tx_id,omitempty"`
}

type isEvent_Event interface {
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event struct {
	EventBase
	Event isEvent_Event `json:"event,omitempty"`
}
