package builder_types

type Withdrawal interface {
	isWithdrawal()
}

type PubKeyWithdrawal struct {
	Address string `json:"address"`
	Coin    uint64 `json:"coin"`
}

func (PubKeyWithdrawal) isWithdrawal() {}

type PlutusScriptWithdrawal struct {
	Address      string       `json:"address"`
	Coin         uint64       `json:"coin"`
	ScriptSource ScriptSource `json:"scriptSource"`
	Redeemer     *Redeemer    `json:"redeemer"`
}

func (PlutusScriptWithdrawal) isWithdrawal() {}

type SimpleScriptWithdrawal struct {
	Address      string             `json:"address"`
	Coin         uint64             `json:"coin"`
	ScriptSource SimpleScriptSource `json:"scriptSource"`
}

func (SimpleScriptWithdrawal) isWithdrawal() {}
