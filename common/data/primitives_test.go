package data_test

import (
	"encoding/json"
	"testing"

	"github.com/sidan-lab/rum/common/data"
)

func TestInteger(t *testing.T) {
	correctInteger := `{"int":10}`
	data := data.NewInteger(10)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctInteger {
		t.Errorf("Integer incorrect")
	}
}

func TestByteString(t *testing.T) {
	correctByteString := `{"bytes":"hello"}`
	data := data.NewByteString("hello")
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctByteString {
		t.Errorf("ByteString incorrect")
	}
}

func TestBool(t *testing.T) {
	correctBool := `{"constructor":1,"fields":[]}`
	data := data.NewBool(false)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctBool {
		t.Errorf("Bool incorrect")
	}
}

func TestList(t *testing.T) {
	correctList := `{"list":[{"bytes":"a0bd47e8938e7c41d4c1d7c22033892319d28f86fdace791d45c51946553791b"}` +
		`,{"int":1000000},{"constructor":0,"fields":[]}]}`
	data := data.NewList([]data.PlutusData{
		data.NewByteString("a0bd47e8938e7c41d4c1d7c22033892319d28f86fdace791d45c51946553791b"),
		data.NewInteger(1000000),
		data.NewBool(true),
	})
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctList {
		t.Errorf("List incorrect")
	}
}

func TestMap(t *testing.T) {
	correctMap := `{"map":[{"k":{"bytes":"aa"},"v":{"int":1000000}},{"k":{"bytes":"bb"},"v":{"int":2000000}}]}`
	data := data.NewMap([][2]data.PlutusData{
		{data.NewByteString("aa"), data.NewInteger(1000000)},
		{data.NewByteString("bb"), data.NewInteger(2000000)},
	})
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctMap {
		t.Errorf("Map incorrect")
	}
}

func TestTuple(t *testing.T) {
	correctTuple := `{"constructor":0,"fields":[{"bytes":"hello"},{"bytes":"world"}]}`
	data := data.NewTuple(data.NewByteString("hello"), data.NewByteString("world"))
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctTuple {
		t.Errorf("Tuple incorrect")
	}
}
