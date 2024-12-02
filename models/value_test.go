package models

import (
	"testing"
)

func TestAddAsset(t *testing.T) {
	mv := NewValue()
	asset := Asset{Unit: "USD", Quantity: "100"}
	mv.AddAsset(asset)

	if mv.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv.Get("USD"))
	}
}

func TestAddAssets(t *testing.T) {
	mv := NewValue()
	assets := []Asset{
		{Unit: "USD", Quantity: "100"},
		{Unit: "EUR", Quantity: "200"},
	}
	mv.AddAssets(assets)

	if mv.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv.Get("USD"))
	}
	if mv.Get("EUR") != 200 {
		t.Errorf("Expected 200, got %d", mv.Get("EUR"))
	}
}

func TestNegateAsset(t *testing.T) {
	mv := NewValue()
	asset := Asset{Unit: "USD", Quantity: "100"}
	mv.AddAsset(asset)
	mv.NegateAsset(asset)

	if mv.Get("USD") != 0 {
		t.Errorf("Expected 0, got %d", mv.Get("USD"))
	}
}

func TestNegateAssets(t *testing.T) {
	mv := NewValue()
	assets := []Asset{
		{Unit: "USD", Quantity: "100"},
		{Unit: "EUR", Quantity: "200"},
	}
	mv.AddAssets(assets)
	mv.NegateAssets(assets)

	if mv.Get("USD") != 0 {
		t.Errorf("Expected 0, got %d", mv.Get("USD"))
	}
	if mv.Get("EUR") != 0 {
		t.Errorf("Expected 0, got %d", mv.Get("EUR"))
	}
}

func TestMerge(t *testing.T) {
	mv1 := NewValue()
	mv2 := NewValue()
	mv1.AddAsset(Asset{Unit: "USD", Quantity: "100"})
	mv2.AddAsset(Asset{Unit: "USD", Quantity: "200"})
	mv1.Merge(mv2)

	if mv1.Get("USD") != 300 {
		t.Errorf("Expected 300, got %d", mv1.Get("USD"))
	}
}

func TestMergeNil(t *testing.T) {
	mv1 := NewValue()
	mv1.AddAsset(Asset{Unit: "USD", Quantity: "100"})
	mv1.Merge(nil)

	if mv1.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv1.Get("USD"))
	}
}

func TestMergeFromNewMap(t *testing.T) {
	mv1 := NewValue()
	mv1.AddAsset(Asset{Unit: "USD", Quantity: "100"})

	mv2 := NewValue()
	mv2.Merge(mv1)

	if mv1.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv1.Get("USD"))
	}
}

func TestToAssets(t *testing.T) {
	mv := NewValue()
	mv.AddAsset(Asset{Unit: "USD", Quantity: "100"})
	assets := mv.ToAssets()

	if len(assets) != 1 {
		t.Errorf("Expected 1 asset, got %d", len(assets))
	}
	if assets[0].Unit != "USD" || assets[0].Quantity != "100" {
		t.Errorf("Expected asset with Unit USD and Quantity 100, got Unit %s and Quantity %s", assets[0].Unit, assets[0].Quantity)
	}
}

func TestFromAssets(t *testing.T) {
	assets := []Asset{
		{Unit: "USD", Quantity: "100"},
		{Unit: "EUR", Quantity: "200"},
	}
	mv := FromAssets(assets)

	if mv.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv.Get("USD"))
	}
	if mv.Get("EUR") != 200 {
		t.Errorf("Expected 200, got %d", mv.Get("EUR"))
	}
}
