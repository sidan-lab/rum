// Package data
package data

type PlutusData interface {
	isPlutusData()
}

type Constr struct {
	Tag    uint         `json:"constructor"`
	Fields []PlutusData `json:"fields"`
}

func (Constr) isPlutusData() {}

func NewConstr(tag uint, fields []PlutusData) Constr {
	if len(fields) == 0 {
		fields = make([]PlutusData, 0)
	}
	return Constr{
		Tag:    tag,
		Fields: fields,
	}
}

func NewConstr0(fileds []PlutusData) Constr {
	return NewConstr(0, fileds)
}

func NewConstr1(fileds []PlutusData) Constr {
	return NewConstr(1, fileds)
}

func NewConstr2(fileds []PlutusData) Constr {
	return NewConstr(2, fileds)
}
