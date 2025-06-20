package data_test

import (
	"encoding/json"
	"testing"

	"github.com/sidan-lab/rum/common/data"
)

func TestConstr(t *testing.T) {
	correctConstr := `{"constructor":10,"fields":[{"bytes":"hello"}]}`
	data := data.NewConstr(10, []data.PlutusData{
		data.NewByteString("hello"),
	})
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctConstr {
		t.Errorf("Constr incorrect")
	}
}

func TestConstr0(t *testing.T) {
	correctConstr0 := `{"constructor":0,"fields":[{"bytes":"hello"}]}`
	data := data.NewConstr0([]data.PlutusData{
		data.NewByteString("hello"),
	})
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctConstr0 {
		t.Errorf("Constr0 incorrect")
	}
}

func TestConstr1(t *testing.T) {
	correctConstr1 := `{"constructor":1,"fields":[{"bytes":"hello"}]}`
	data := data.NewConstr1([]data.PlutusData{
		data.NewByteString("hello"),
	})
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctConstr1 {
		t.Errorf("Constr1 incorrect")
	}
}

func TestConstr2(t *testing.T) {
	correctConstr2 := `{"constructor":2,"fields":[{"bytes":"hello"}]}`
	data := data.NewConstr2([]data.PlutusData{
		data.NewByteString("hello"),
	})
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctConstr2 {
		t.Errorf("Constr2 incorrect")
	}
}
