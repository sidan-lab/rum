package models_test

import (
	"reflect"
	"testing"

	"github.com/sidan-lab/rum/models"
)

// func (a *Assets) GetLovelace() uint64 {
// 	if a == nil {
// 		return 0
// 	}
// 	for _, asset := range *a {
// 		if asset.Unit == "lovelace" {
// 			quantity, _ := strconv.ParseUint(asset.Quantity, 10, 64)
// 			return quantity
// 		}
// 	}
// 	return 0
// }

func TestGetLovelace(t *testing.T) {
	assets := models.Assets{
		{Unit: "lovelace", Quantity: "100"},
		{Unit: "USD", Quantity: "100"},
	}
	if assets.GetLovelace() != 100 {
		t.Errorf("Expected 100, got %d", assets.GetLovelace())
	}
	assets = models.Assets{
		{Unit: "USD", Quantity: "100"},
	}
	if assets.GetLovelace() != 0 {
		t.Errorf("Expected 0, got %d", assets.GetLovelace())
	}

	var nilAssets models.Assets
	if nilAssets.GetLovelace() != 0 {
		t.Errorf("Expected 0, got %d", nilAssets.GetLovelace())
	}
}

func TestPopAssetByUnit(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		initialAssets  models.Assets
		unitToPop      string
		expectedAsset  *models.Asset
		expectedAssets models.Assets
	}{
		{
			name: "Pop existing asset",
			initialAssets: models.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset1", Quantity: "2000"},
				{Unit: "asset2", Quantity: "3000"},
			},
			unitToPop:     "asset1",
			expectedAsset: &models.Asset{Unit: "asset1", Quantity: "2000"},
			expectedAssets: models.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset2", Quantity: "3000"},
			},
		},
		{
			name: "Pop non-existing asset",
			initialAssets: models.Assets{
				{Unit: "lovelace", Quantity: "1000"},
				{Unit: "asset1", Quantity: "2000"},
				{Unit: "asset2", Quantity: "3000"},
			},
			unitToPop:     "asset3",
			expectedAsset: &models.Asset{},
			expectedAssets: models.Assets{
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
