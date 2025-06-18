package builder_types

type TxBuilderBody struct {
	Inputs             []TxIn        `json:"inputs"`
	Outputs            []Output      `json:"outputs"`
	Collaterals        []PubKeyTxIn  `json:"collaterals"`
	RequiredSignatures []string      `json:"requiredSignatures"`
	ReferenceInputs    []RefTxIn     `json:"referenceInputs"`
	Withdrawals        []Withdrawal  `json:"withdrawals"`
	Mints              []MintItem    `json:"mints"`
	ChangeAddress      string        `json:"changeAddress"`
	ChangeDatum        Datum         `json:"changeDatum"`
	Metadata           []Metadata    `json:"metadata"`
	ValidityRange      ValidityRange `json:"validityRange"`
	Certificates       []Certificate `json:"certificates"`
	Votes              []Vote        `json:"votes"`
	SigningKey         []string      `json:"signingKey"`
	Fee                *string       `json:"fee"`
	Network            *Network      `json:"network"`
}

func NewTxBuilderBody() *TxBuilderBody {
	return &TxBuilderBody{
		Inputs:             []TxIn{},
		Outputs:            []Output{},
		Collaterals:        []PubKeyTxIn{},
		RequiredSignatures: []string{},
		ReferenceInputs:    []RefTxIn{},
		Withdrawals:        []Withdrawal{},
		Mints:              []MintItem{},
		ChangeAddress:      "",
		ChangeDatum:        nil,
		Certificates:       []Certificate{},
		Votes:              []Vote{},
		Metadata:           []Metadata{},
		ValidityRange: ValidityRange{
			InvalidBefore:    nil,
			InvalidHereafter: nil,
		},
		SigningKey: []string{},
		Fee:        nil,
		Network:    nil,
	}
}
