package builder_types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type LanguageVersion string

const (
	LanguageVersionV1 LanguageVersion = "v1"
	LanguageVersionV2 LanguageVersion = "v2"
	LanguageVersionV3 LanguageVersion = "v3"
)

func (lv *LanguageVersion) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "v1", "v2", "v3":
		*lv = LanguageVersion(s)
		return nil
	default:
		return fmt.Errorf("invalid LanguageVersion: %s", s)
	}
}

type SimpleScriptSource interface {
	isSimpleScriptSource()
}

type ProvidedSimpleScriptSource struct {
	ScriptCbor string `json:"scriptCbor"`
}

func (ProvidedSimpleScriptSource) isSimpleScriptSource() {}

type InlineSimpleScriptSource struct {
	RefTxIn          RefTxIn `json:"refTxIn"`
	SimpleScriptHash string  `json:"simpleScriptHash"`
	ScriptSize       uint    `json:"scriptSize"`
}

func (InlineSimpleScriptSource) isSimpleScriptSource() {}
func (InlineSimpleScriptSource) isVote()               {}

type ScriptSource interface {
	isScriptSource()
}

type ProvidedScriptSource struct {
	ScriptCbor      string          `json:"scriptCbor"`
	LanguageVersion LanguageVersion `json:"languageVersion"`
}

func (ProvidedScriptSource) isScriptSource() {}

type InlineScriptSource struct {
	RefTxIn         RefTxIn         `json:"refTxIn"`
	ScriptHash      string          `json:"scriptHash"`
	LanguageVersion LanguageVersion `json:"languageVersion"`
	ScriptSize      uint            `json:"scriptSize"`
}

func (InlineScriptSource) isScriptSource() {}

type DatumSource interface {
	isDatumSource()
}

type ProvidedDatumSource struct {
	Data string `json:"data"`
}

func (ProvidedDatumSource) isDatumSource() {}

type InlineDatumSource struct {
	TxHash  string `json:"txHash"`
	TxIndex uint32 `json:"txIndex"`
}

func (InlineDatumSource) isDatumSource() {}
