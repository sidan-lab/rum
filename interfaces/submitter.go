package interfaces

type ISubmitter interface {
	SubmitTx(txCbor string) (string, error)
}
