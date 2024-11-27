package models

type AccountInfo struct {
	Active      bool
	PoolId      *string
	Balance     string
	Rewards     string
	Withdrawals string
}
