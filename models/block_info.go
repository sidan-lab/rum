package models

type BlockInfo struct {
	Time                   int
	Hash                   string
	Slot                   string
	Epoch                  int
	EpochSlot              string
	SlotLeader             string
	Size                   int
	TxCount                int
	Output                 string
	Fees                   string
	PreviousBlock          string
	NextBlock              string
	Confirmations          int
	OperationalCertificate string
	VRFKey                 string
}
