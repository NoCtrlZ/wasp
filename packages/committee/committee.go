package committee

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/vm"
)

type Committee interface {
	Address() *address.Address
	OwnerAddress() *address.Address
	Color() *balance.Color
	Size() uint16
	OwnPeerIndex() uint16
	NumPeers() uint16
	SendMsg(targetPeerIndex uint16, msgType byte, msgData []byte) error
	SendMsgToCommitteePeers(msgType byte, msgData []byte) (uint16, int64)
	SendMsgInSequence(msgType byte, msgData []byte, seqIndex uint16, seq []uint16) (uint16, error)
	IsAlivePeer(peerIndex uint16) bool
	ReceiveMessage(msg interface{})
	InitTestRound()
	//
	SetReadyStateManager()
	SetReadyConsensus()
	Dismiss()
	IsDismissed() bool
}

type StateManager interface {
	EvidenceStateIndex(idx uint32)
	EventGetBatchMsg(msg *GetBatchMsg)
	EventBatchHeaderMsg(msg *BatchHeaderMsg)
	EventStateUpdateMsg(msg *StateUpdateMsg)
	EventStateTransactionMsg(msg StateTransactionMsg)
	EventPendingBatchMsg(msg PendingBatchMsg)
	EventTimerMsg(msg TimerTick)
}

type Operator interface {
	EventProcessorReady(msg ProcessorIsReady)
	EventStateTransitionMsg(msg *StateTransitionMsg)
	EventBalancesMsg(reqMsg BalancesMsg)
	EventRequestMsg(reqMsg RequestMsg)
	EventNotifyReqMsg(msg *NotifyReqMsg)
	EventStartProcessingBatchMsg(msg *StartProcessingBatchMsg)
	EventResultCalculated(result *vm.VMTask)
	EventSignedHashMsg(msg *SignedHashMsg)
	EventTimerMsg(msg TimerTick)
}

var ConstructorNew func(bootupData *registry.BootupData, log *logger.Logger) Committee

func New(bootupData *registry.BootupData, log *logger.Logger) Committee {
	return ConstructorNew(bootupData, log)
}
