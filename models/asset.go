package models

import "strconv"

type Asset struct {
	Unit     string `json:"unit" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
}

type Assets []Asset

func (a *Assets) PopAssetByUnit(unit string) Asset {
	var filteredAssets Assets
	filteredAssets = []Asset{}
	found := false
	var filteredAsset Asset
	for _, asset := range *a {
		if asset.Unit == unit && !found {
			filteredAsset = asset
			found = true
		} else {
			filteredAssets = append(filteredAssets, asset)
		}
	}
	a = &filteredAssets
	return filteredAsset
}

func (a *Assets) MergeAssets(assets []Asset) []Asset {
	mergedAssets := make(map[string]Asset)
	for _, asset := range assets {
		if existingAsset, ok := mergedAssets[asset.Unit]; ok {
			existingAsset.Quantity = AddQuantities(existingAsset.Quantity, asset.Quantity)
			mergedAssets[asset.Unit] = existingAsset
		} else {
			mergedAssets[asset.Unit] = asset
		}
	}
	result := make([]Asset, 0, len(mergedAssets))
	for _, asset := range mergedAssets {
		result = append(result, asset)
	}
	return result
}

func AddQuantities(quantity1 string, quantity2 string) string {
	intQuantity1, _ := strconv.Atoi(quantity1)
	intQuantity2, _ := strconv.Atoi(quantity2)
	sum := intQuantity1 + intQuantity2
	stringSum := strconv.Itoa(sum)
	return stringSum
}
