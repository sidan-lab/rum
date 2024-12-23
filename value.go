package rum

import (
	"strconv"
)

/*
Value provide utility to handle the Cardano value manipulation. It offers certain axioms:
1. No duplication of asset - adding assets with same asset name will increase the quantity of the asset in the same record.
2. No zero and negative entry - the quantity of the asset should not be zero or negative.
3. Sanitization of lovelace asset name - the class handle back and forth conversion of lovelace asset name to empty string.
4. Easy convertion to Cardano data - offer utility to convert into either Mesh Data type and JSON type for its Cardano data representation.
*/
type Value struct {
	Value map[string]int64
}

// NewValue create a new Value instance with empty value.
func NewValue() *Value {
	return &Value{
		Value: make(map[string]int64),
	}
}

// NewValueFromAssets - create a new Value instance with the given assets.
func NewValueFromAssets(assets *[]Asset) *Value {
	value := &Value{
		Value: make(map[string]int64),
	}
	if assets == nil {
		return value
	}
	return value.AddAssets(assets)
}

// AddAsset - Add an asset to the Value class's value record.
func (v *Value) AddAsset(asset *Asset) *Value {
	quantity, _ := strconv.ParseInt(asset.Quantity, 10, 64)
	if existingQuantity, exists := v.Value[asset.Unit]; exists {
		v.Value[asset.Unit] = existingQuantity + quantity
	} else {
		v.Value[asset.Unit] = quantity
	}
	return v
}

// AddAssets - add multiple assets to the Value class's value record.
func (v *Value) AddAssets(assets *[]Asset) *Value {
	if assets == nil {
		return v
	}
	for _, asset := range *assets {
		v.AddAsset(&asset)
	}
	return v
}

// NegateAsset - deduct the value amount of an asset from the Value class's value record.
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

// NegateAssets - deduct the value amount of multiple assets from the Value class's value record.
func (v *Value) NegateAssets(assets *[]Asset) *Value {
	if assets == nil {
		return v
	}
	for _, asset := range *assets {
		v.NegateAsset(&asset)
	}
	return v
}

// Merge - merge multiple Value class's value record into the current Value class's value record.
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

// Get - get the quantity of an asset in the Value class's value record.
func (v *Value) Get(unit string) int64 {
	if quantity, exists := v.Value[unit]; exists {
		return quantity
	}
	return 0
}

// Units - get the list of asset names in the Value class's value record.
func (v *Value) Units() []string {
	units := make([]string, 0, len(v.Value))
	for unit := range v.Value {
		units = append(units, unit)
	}
	return units
}

// IsEmpty - check if the Value class's value record is empty.
func (v *Value) IsEmpty() bool {
	return len(v.Value) == 0
}

// ToAssets - convert the Value class's value record into a list of Asset.
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

// Geq - check if the value is greater than or equal to another value
func (v *Value) Geq(other *Value) bool {
	for unit, quantity := range other.Value {
		if existingQuantity, exists := v.Value[unit]; !exists || existingQuantity < quantity {
			return false
		}
	}
	return true
}
