package builder_types

type MintItem interface {
	isMintItem()
}

type ScriptMint struct {
	Mint         MintParameter `json:"mint"`
	Redeemer     *Redeemer     `json:"redeemer"`
	ScriptSource ScriptSource  `json:"scriptSource"`
}

func (ScriptMint) isMintItem() {}

type SimpleScriptMint struct {
	Mint         MintParameter      `json:"mint"`
	ScriptSource SimpleScriptSource `json:"scriptSource"`
}

func (SimpleScriptMint) isMintItem() {}

type MintParameter struct {
	PolicyID  string `json:"policyId"`
	AssetName string `json:"assetName"`
	Amount    int64  `json:"amount"`
}
