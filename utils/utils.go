package utils

import "github.com/sidan-lab/rum/models"

func FindUtxoByIndex(utxos []models.UTxO, index int) *models.UTxO {
	for _, utxo := range utxos {
		if utxo.Input.OutputIndex == index {
			return &utxo
		}
	}
	return nil
}
