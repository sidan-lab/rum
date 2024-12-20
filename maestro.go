package rum

import (
	"encoding/json"
	"strconv"

	"github.com/maestro-org/go-sdk/client"
	msModel "github.com/maestro-org/go-sdk/models"
)

type MaestroProvider struct {
	maestroClient *client.Client
}

func NewMaestroProvider(apiKey string, network Network) *MaestroProvider {
	maestroClient := client.NewClient(apiKey, string(network))
	return &MaestroProvider{
		maestroClient: maestroClient,
	}
}

func (ms *MaestroProvider) SubmitTx(txCbor string) (string, error) {
	txHash, err := ms.maestroClient.TxManagerSubmit(txCbor)
	return txHash, err
}

func (ms *MaestroProvider) FetchTxInfo(hash string) (TransactionInfo, error) {
	tx, err := ms.maestroClient.TransactionDetails(hash)
	if err != nil {
		return TransactionInfo{}, err
	}
	msTxInfo := tx.Data
	txInfo := TransactionInfo{
		Index:         int(msTxInfo.BlockTxIndex),
		Block:         msTxInfo.BlockHash,
		Hash:          msTxInfo.TxHash,
		Slot:          strconv.FormatInt(msTxInfo.BlockAbsoluteSlot, 10),
		Fees:          strconv.FormatInt(msTxInfo.Fee, 10),
		Size:          int(msTxInfo.Size),
		Deposit:       strconv.FormatInt(msTxInfo.Deposit, 10),
		InvalidBefore: formatOptionalInt(msTxInfo.InvalidBefore),
		InvalidAfter:  formatOptionalInt(msTxInfo.InvalidHereafter),
	}
	return txInfo, nil
}

type MsDatum struct {
	Hash  string `json:"hash"`
	Bytes string `json:"bytes"`
	Json  any    `json:"json"`
}

func (ms *MaestroProvider) FetchUTxOs(hash string, index *int) ([]UTxO, error) {
	res, err := ms.maestroClient.TransactionDetails(hash)
	if err != nil {
		return nil, err
	}
	msOutputs := res.Data.Outputs
	utxos := MsToUtxos(msOutputs)
	if index != nil {
		utxo := FindUtxoByIndex(utxos, *index)
		if utxo != nil {
			return []UTxO{*utxo}, nil
		}
		return []UTxO{}, nil
	}
	return utxos, nil
}

func MsToUtxos(msUtxos []msModel.Utxo) []UTxO {
	utxos := make([]UTxO, len(msUtxos))
	for i, msUtxo := range msUtxos {
		utxos[i] = MsToUtxo(msUtxo)
	}
	return utxos
}

func MsToUtxo(msUtxo msModel.Utxo) UTxO {
	var datum MsDatum
	if msUtxo.Datum != nil {
		datumBytes, err := json.Marshal(msUtxo.Datum)
		if err != nil {
		} else {
			json.Unmarshal(datumBytes, &datum)
		}
	}
	return UTxO{
		Input: Input{
			TxHash:      msUtxo.TxHash,
			OutputIndex: int(msUtxo.Index),
		},
		Output: Output{
			Amount:     MsToAssets(msUtxo.Assets),
			Address:    msUtxo.Address,
			DataHash:   datum.Hash,
			PlutusData: datum.Bytes,
			ScriptRef:  "", // TODO: add script ref
			ScriptHash: msUtxo.ReferenceScript.Hash,
		},
	}
}

func MsToAssets(msAssets []msModel.Asset) []Asset {
	assets := make([]Asset, len(msAssets))
	for i, msAsset := range msAssets {
		assets[i] = MsToAsset(msAsset)
	}
	return assets
}

func MsToAsset(msAsset msModel.Asset) Asset {
	return Asset{
		Quantity: strconv.FormatInt(msAsset.Amount, 10),
		Unit:     msAsset.Unit,
	}
}

func formatOptionalInt(value *int64) string {
	if value != nil {
		return strconv.FormatInt(*value, 10)
	}
	return ""
}
