package data

func CurrencySymbol(policyID string) ByteString {
	return NewByteString(policyID)
}

func TokenName(tokenName string) ByteString {
	return NewByteString(tokenName)
}

func AssetClass(policyID string, tokenName string) Constr {
	return NewConstr0([]PlutusData{
		CurrencySymbol(policyID),
		TokenName(tokenName),
	})
}

func TxOutRef(txHash string, index int64) Constr {
	return NewConstr0([]PlutusData{
		NewConstr0([]PlutusData{
			NewByteString(txHash),
		}),
		NewInteger(index),
	})
}

func OutputReference(txHash string, index int64) Constr {
	return NewConstr0([]PlutusData{
		NewByteString(txHash),
		NewInteger(index),
	})
}

func PosixTime(posixTime int64) Integer {
	return NewInteger(posixTime)
}
