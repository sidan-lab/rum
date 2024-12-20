package rum_test

import (
	"reflect"
	"testing"

	"github.com/sidan-lab/rum"
)

func TestGetLovelace(t *testing.T) {
	assets := rum.Assets{
		{Unit: "lovelace", Quantity: "100"},
		{Unit: "USD", Quantity: "100"},
	}
	if assets.GetLovelace() != 100 {
		t.Errorf("Expected 100, got %d", assets.GetLovelace())
	}
	assets = rum.Assets{
		{Unit: "USD", Quantity: "100"},
	}
	if assets.GetLovelace() != 0 {
		t.Errorf("Expected 0, got %d", assets.GetLovelace())
	}

	var nilAssets rum.Assets
	if nilAssets.GetLovelace() != 0 {
		t.Errorf("Expected 0, got %d", nilAssets.GetLovelace())
	}
}

func TestPopAssetByUnit(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		initialAssets  rum.Assets
		unitToPop      string
		expectedAsset  *rum.Asset
		expectedAssets rum.Assets
	}{
		{
			name: "Pop existing asset",
			initialAssets: rum.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset1", Quantity: "2000"},
				{Unit: "asset2", Quantity: "3000"},
			},
			unitToPop:     "asset1",
			expectedAsset: &rum.Asset{Unit: "asset1", Quantity: "2000"},
			expectedAssets: rum.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset2", Quantity: "3000"},
			},
		},
		{
			name: "Pop non-existing asset",
			initialAssets: rum.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset1", Quantity: "2000"},
				{Unit: "asset2", Quantity: "3000"},
			},
			unitToPop:     "asset3",
			expectedAsset: &rum.Asset{},
			expectedAssets: rum.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset1", Quantity: "2000"},
				{Unit: "asset2", Quantity: "3000"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assets := tt.initialAssets
			poppedAsset := assets.PopAssetByUnit(tt.unitToPop)

			if !reflect.DeepEqual(poppedAsset, tt.expectedAsset) {
				t.Errorf("PopAssetByUnit() = %v, want %v", poppedAsset, tt.expectedAsset)
			}

			if !reflect.DeepEqual(assets, tt.expectedAssets) {
				t.Errorf("Remaining assets = %v, want %v", assets, tt.expectedAssets)
			}
		})
	}
}
