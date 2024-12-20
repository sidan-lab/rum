package rum

func FindUtxoByIndex(utxos []UTxO, index int) *UTxO {
	for _, utxo := range utxos {
		if utxo.Input.OutputIndex == index {
			return &utxo
		}
	}
	return nil
}
