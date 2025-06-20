package data

func NewPaymentPubKeyHash(pubKeyHash string) ByteString {
	return NewByteString(pubKeyHash)
}

func NewPubKeyHash(pubKeyHash string) ByteString {
	return NewByteString(pubKeyHash)
}

func NewMaybeStakingHash(stakeCredential *string, isScriptStakeKey bool) Constr {
	if stakeCredential == nil {
		return NewConstr1([]PlutusData{})
	} else if isScriptStakeKey {
		return NewConstr0([]PlutusData{
			NewConstr0([]PlutusData{
				NewConstr1([]PlutusData{
					NewByteString(*stakeCredential),
				}),
			}),
		})
	} else {
		return NewConstr0([]PlutusData{
			NewConstr0([]PlutusData{
				NewConstr0([]PlutusData{
					NewByteString(*stakeCredential),
				}),
			}),
		})
	}
}

func NewPubKeyAddress(bytes string, stakeCredential *string, isScriptStakeKey bool) Constr {
	return NewConstr0([]PlutusData{
		NewConstr0([]PlutusData{
			NewByteString(bytes),
		}),
		NewMaybeStakingHash(stakeCredential, isScriptStakeKey),
	})
}

func NewScriptAddress(bytes string, stakeCredential *string, isScriptStakeKey bool) Constr {
	return NewConstr0([]PlutusData{
		NewConstr1([]PlutusData{
			NewByteString(bytes),
		}),
		NewMaybeStakingHash(stakeCredential, isScriptStakeKey),
	})
}
