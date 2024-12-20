package rum

import (
	"strconv"
)

type Value struct {
	Value map[string]int64
}

func NewValue() *Value {
	return &Value{
		Value: make(map[string]int64),
	}
}

func NewValueFromAssets(assets *[]Asset) *Value {
	value := &Value{
		Value: make(map[string]int64),
	}
	if assets == nil {
		return value
	}
	return value.AddAssets(assets)
}

func (v *Value) AddAsset(asset *Asset) *Value {
	quantity, _ := strconv.ParseInt(asset.Quantity, 10, 64)
	if existingQuantity, exists := v.Value[asset.Unit]; exists {
		v.Value[asset.Unit] = existingQuantity + quantity
	} else {
		v.Value[asset.Unit] = quantity
	}
	return v
}

func (v *Value) AddAssets(assets *[]Asset) *Value {
	if assets == nil {
		return v
	}
	for _, asset := range *assets {
		v.AddAsset(&asset)
	}
	return v
}

func (v *Value) NegateAsset(asset *Asset) *Value {
	if asset == nil {
		return v
	}
	quantity, _ := strconv.ParseInt(asset.Quantity, 10, 64)
	if existingQuantity, exists := v.Value[asset.Unit]; exists {
		newQuantity := existingQuantity - quantity
		if newQuantity == 0 {
			delete(v.Value, asset.Unit)
		} else {
			v.Value[asset.Unit] = newQuantity
		}
	}
	return v
}

func (v *Value) NegateAssets(assets *[]Asset) *Value {
	if assets == nil {
		return v
	}
	for _, asset := range *assets {
		v.NegateAsset(&asset)
	}
	return v
}

func (v *Value) Merge(values ...*Value) *Value {
	for _, other := range values {
		if other == nil {
			continue
		}
		for unit, quantity := range other.Value {
			if existingQuantity, exists := v.Value[unit]; exists {
				v.Value[unit] = existingQuantity + quantity
			} else {
				v.Value[unit] = quantity
			}
		}
	}
	return v
}

func (v *Value) Get(unit string) int64 {
	if quantity, exists := v.Value[unit]; exists {
		return quantity
	}
	return 0
}

func (v *Value) Units() []string {
	units := make([]string, 0, len(v.Value))
	for unit := range v.Value {
		units = append(units, unit)
	}
	return units
}

func (v *Value) IsEmpty() bool {
	return len(v.Value) == 0
}

func (v *Value) ToAssets() *[]Asset {
	assets := make([]Asset, 0, len(v.Value))
	for unit, quantity := range v.Value {
		assets = append(assets, Asset{
			Unit:     unit,
			Quantity: strconv.FormatInt(quantity, 10),
		})
	}
	return &assets
}

func (v *Value) Geq(other *Value) bool {
	for unit, quantity := range other.Value {
		if existingQuantity, exists := v.Value[unit]; !exists || existingQuantity < quantity {
			return false
		}
	}
	return true
}
