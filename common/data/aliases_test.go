package data_test

import (
	"encoding/json"
	"testing"

	"github.com/sidan-lab/rum/common/data"
)


func TestCurrenySymbol(t *testing.T) {
	correctCurrencySymbol := `{"bytes":"hello"}`
	data := data.CurrencySymbol("hello")
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctCurrencySymbol{
		t.Errorf("CurrenySymbol incorrect")
	}
}

func TestTokenName(t *testing.T) {
	correctTokenName := `{"bytes":"hello"}`
	data := data.TokenName("hello")
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctTokenName{
		t.Errorf("TokenName incorrect")
	}
}

func TestAssetClass(t *testing.T) {
	correctAssetClass := `{"constructor":0,"fields":[{"bytes":"hello"},{"bytes":"world"}]}`
	data := data.AssetClass("hello", "world")
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctAssetClass {
		t.Errorf("AssetClass incorrect")
	}
}

func TestTxOutRef(t *testing.T) {
	correctTxOutRef := `{"constructor":0,"fields":[{"constructor":0,"fields":[{"bytes":"hello"}]},{"int":12}]}`
	data := data.TxOutRef("hello", 12)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctTxOutRef {
		t.Errorf("TxOutRef incorrect")
	}
}

func TestOutputReference(t *testing.T) {
	correctOutputReference := `{"constructor":0,"fields":[{"bytes":"hello"},{"int":12}]}`
	data := data.OutputReference("hello", 12)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctOutputReference {
		t.Errorf("OutputReference incorrect")
	}
}

func TestPosixTime(t *testing.T) {
	correctPosixTime := `{"int":12}`
	data := data.PosixTime(12)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctPosixTime {
		t.Errorf("PosixTime incorrect")
	}
}
