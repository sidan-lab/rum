package data

import "encoding/json"

func NewBool(b bool) Constr {
	if b {
		return NewConstr0([]PlutusData{})
	} else {
		return NewConstr1([]PlutusData{})
	}
}

type ByteString struct {
	Inner string `json:"bytes"`
}

func (ByteString) isPlutusData() {}

func NewByteString(s string) ByteString {
	return ByteString{
		Inner: s,
	}
}

type Integer struct {
	Inner int64 `json:"int"`
}

func (Integer) isPlutusData() {}

func NewInteger(i int64) Integer {
	return Integer{
		Inner: i,
	}
}

type List struct {
	Inner []PlutusData `json:"list"`
}

func (List) isPlutusData() {}

func NewList(l []PlutusData) List {
	return List{
		Inner: l,
	}
}

type Map struct {
	Inner [][2]PlutusData `json:"map"`
}

func (Map) isPlutusData() {}

func NewMap(m [][2]PlutusData) Map {
	return Map{
		Inner: m,
	}
}

func (m Map) MarshalJSON() ([]byte, error) {
	items := make([]map[string]PlutusData, 0, len(m.Inner))
	for _, pair := range m.Inner {
		item := map[string]PlutusData{
			"k": pair[0],
			"v": pair[1],
		}
		items = append(items, item)
	}

	return json.Marshal(map[string]any{
		"map": items,
	})
}

func NewTuple(itemA PlutusData, itemB PlutusData) Constr {
	return NewConstr(0, []PlutusData{itemA, itemB})
}
