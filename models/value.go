package models

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

func (mv *Value) AddAsset(asset Asset) {
	quantity, _ := strconv.ParseInt(asset.Quantity, 10, 64)
	if existingQuantity, exists := mv.Value[asset.Unit]; exists {
		mv.Value[asset.Unit] = existingQuantity + quantity
	} else {
		mv.Value[asset.Unit] = quantity
	}
}

func (mv *Value) AddAssets(assets []Asset) {
	for _, asset := range assets {
		mv.AddAsset(asset)
	}
}

func (mv *Value) NegateAsset(asset Asset) {
	quantity, _ := strconv.ParseInt(asset.Quantity, 10, 64)
	if existingQuantity, exists := mv.Value[asset.Unit]; exists {
		newQuantity := existingQuantity - quantity
		if newQuantity == 0 {
			delete(mv.Value, asset.Unit)
		} else {
			mv.Value[asset.Unit] = newQuantity
		}
	}
}

func (mv *Value) NegateAssets(assets []Asset) {
	for _, asset := range assets {
		mv.NegateAsset(asset)
	}
}

func (mv *Value) Get(unit string) int64 {
	if quantity, exists := mv.Value[unit]; exists {
		return quantity
	}
	return 0
}

func (mv *Value) Units() []string {
	units := make([]string, 0, len(mv.Value))
	for unit := range mv.Value {
		units = append(units, unit)
	}
	return units
}

func (mv *Value) IsEmpty() bool {
	return len(mv.Value) == 0
}

func (mv *Value) Merge(values ...*Value) {
	for _, other := range values {
		if other == nil {
			continue
		}
		for unit, quantity := range other.Value {
			if existingQuantity, exists := mv.Value[unit]; exists {
				mv.Value[unit] = existingQuantity + quantity
			} else {
				mv.Value[unit] = quantity
			}
		}
	}
}

func (mv *Value) ToAssets() *[]Asset {
	assets := make([]Asset, 0, len(mv.Value))
	for unit, quantity := range mv.Value {
		assets = append(assets, Asset{
			Unit:     unit,
			Quantity: strconv.FormatInt(quantity, 10),
		})
	}
	return &assets
}

func (mv *Value) Geq(other *Value) bool {
	for unit, quantity := range other.Value {
		if existingQuantity, exists := mv.Value[unit]; !exists || existingQuantity < quantity {
			return false
		}
	}
	return true
}

func FromAssets(assets []Asset) *Value {
	mv := NewValue()
	mv.AddAssets(assets)
	return mv
}
