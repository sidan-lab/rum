package builder

import (
	"github.com/sidan-lab/rum/models"
	types "github.com/sidan-lab/rum/models/builder_types"
)

func (builder *TxBuilder) TxOut(address string, amount []models.Asset) *TxBuilder {
	if builder.TxOutput != nil {
		builder.TxBuilderBody.Outputs = append(builder.TxBuilderBody.Outputs, *builder.TxOutput)
		builder.TxOutput = nil
	}

	builder.TxOutput = &types.Output{
		Address:         address,
		Amount:          amount,
		Datum:           nil,
		ReferenceScript: nil,
	}
	return builder
}

func (builder *TxBuilder) TxOutDatumHashValue(data WData) *TxBuilder {
	if builder.TxOutput == nil {
		panic("Undefined output")
	}
	rawData, err := data.ToCbor()
	if err != nil {
		panic("Error converting datum to CBOR")
	}
	builder.TxOutput.Datum = &types.HashDatum{
		Inner: rawData,
	}
	return builder
}

func (builder *TxBuilder) TxOutDatumEmbedValue(data WData) *TxBuilder {
	if builder.TxOutput == nil {
		panic("Undefined output")
	}
	rawData, err := data.ToCbor()
	if err != nil {
		panic("Error converting datum to CBOR")
	}
	builder.TxOutput.Datum = &types.EmbeddedDatum{
		Inner: rawData,
	}
	return builder
}

func (builder *TxBuilder) TxOutDatumInlineValue(data WData) *TxBuilder {
	if builder.TxOutput == nil {
		panic("Undefined output")
	}
	rawData, err := data.ToCbor()
	if err != nil {
		panic("Error converting datum to CBOR")
	}
	builder.TxOutput.Datum = &types.InlineDatum{
		Inner: rawData,
	}
	return builder
}

func (builder *TxBuilder) TxOutReferenceScript(scriptCbor string, version *types.LanguageVersion) *TxBuilder {
	if builder.TxOutput == nil {
		panic("Undefined output")
	}
	if version != nil {
		builder.TxOutput.ReferenceScript = types.ProvidedScriptSource{
			ScriptCbor:      scriptCbor,
			LanguageVersion: *version,
		}
	} else {
		builder.TxOutput.ReferenceScript = types.ProvidedSimpleScriptSource{
			ScriptCbor: scriptCbor,
		}
	}
	return builder
}
