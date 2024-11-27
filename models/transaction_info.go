package models

type TransactionInfo struct {
	Index         int
	Block         string
	Hash          string
	Slot          string
	Fees          string
	Size          int
	Deposit       string
	InvalidBefore string
	InvalidAfter  string
}
