package models_test

import (
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
