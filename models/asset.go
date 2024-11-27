package models

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
