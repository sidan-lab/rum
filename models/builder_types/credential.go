package builder_types

type Credential interface {
	isCredential()
}

type ScriptHash struct {
	Inner string `json:"scriptHash"`
}

func (ScriptHash) isCredential() {}


type KeyHash struct {
	Inner string `json:"keyHash"`
}

func (KeyHash) isCredential() {}
