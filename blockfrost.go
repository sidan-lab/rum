package rum

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/blockfrost/blockfrost-go"
)

var NetworkMap = map[string]string{
	"testnet": blockfrost.CardanoTestNet,
	"preview": blockfrost.CardanoPreview,
	"preprod": blockfrost.CardanoPreProd,
	"mainnet": blockfrost.CardanoMainNet,
}

func GetServerURL(projectID string) string {
	if len(projectID) >= 7 {
		prefix := projectID[:7]
		if url, exists := NetworkMap[prefix]; exists {
			return url
		}
	}
	return NetworkMap["mainnet"]
}

type BlockfrostProvider struct {
	blockfrostClient blockfrost.APIClient
}

func NewBlockfrostProvider(projectID string) *BlockfrostProvider {
	blockfrostClient := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{
			ProjectID: projectID,
			Server:    GetServerURL(projectID),
		},
	)
	return &BlockfrostProvider{
		blockfrostClient: blockfrostClient,
	}
}

func (bf *BlockfrostProvider) SubmitTx(txCbor string) (string, error) {
	txBuffer, err := hex.DecodeString(txCbor)
	if err != nil {
		return "", err
	}

	txHash, err := bf.blockfrostClient.TransactionSubmit(context.TODO(), txBuffer)
	return txHash, err
}

func (bf *BlockfrostProvider) FetchTxInfo(hash string) (TransactionInfo, error) {
	bfTxInfo, err := bf.blockfrostClient.Transaction(context.TODO(), hash)
	if err != nil {
		return TransactionInfo{}, err
	}

	invalidHereafter := ""
	if bfTxInfo.InvalidHereafter != nil {
		invalidHereafter = *bfTxInfo.InvalidHereafter
	}
	invalidBefore := ""
	if bfTxInfo.InvalidBefore != nil {
		invalidBefore = *bfTxInfo.InvalidBefore
	}

	txInfo := TransactionInfo{
		Block:         bfTxInfo.Block,
		Deposit:       bfTxInfo.Deposit,
		Fees:          bfTxInfo.Fees,
		Hash:          bfTxInfo.Hash,
		Index:         int(bfTxInfo.Index),
		InvalidAfter:  invalidHereafter,
		InvalidBefore: invalidBefore,
		Slot:          strconv.FormatInt(int64(bfTxInfo.Slot), 10),
		Size:          int(bfTxInfo.Size),
	}
	return txInfo, nil
}

func (bf *BlockfrostProvider) FetchUTxOs(hash string, index *int) ([]UTxO, error) {
	res, err := bf.blockfrostClient.TransactionUTXOs(context.TODO(), hash)
	if err != nil {
		return nil, err
	}
	bfOutputs := res.Outputs
	utxos := BfToUtxos(bfOutputs, hash)
	if index != nil {
		utxo := FindUtxoByIndex(utxos, *index)
		if utxo != nil {
			return []UTxO{*utxo}, nil
		}
		return []UTxO{}, nil
	}
	return utxos, nil
}

func BfToUtxos(bfUtxos []blockfrost.TransactionOutput, hash string) []UTxO {
	utxos := make([]UTxO, len(bfUtxos))
	for i, bfUtxo := range bfUtxos {
		utxos[i] = BfToUtxo(bfUtxo, hash)
	}
	return utxos
}

func BfToUtxo(bfUtxo blockfrost.TransactionOutput, hash string) UTxO {
	dataHash := ""
	if bfUtxo.DataHash != nil {
		dataHash = *bfUtxo.DataHash
	}
	inlineDatum := ""
	if bfUtxo.InlineDatum != nil {
		inlineDatum = *bfUtxo.InlineDatum
	}
	referenceScriptHash := ""
	if bfUtxo.ReferenceScriptHash != nil {
		referenceScriptHash = *bfUtxo.ReferenceScriptHash
	}

	return UTxO{
		Input: Input{
			TxHash:      hash,
			OutputIndex: int(bfUtxo.OutputIndex),
		},
		Output: Output{
			Amount:     BfToAssets(bfUtxo.Amount),
			Address:    bfUtxo.Address,
			DataHash:   dataHash,
			PlutusData: inlineDatum,
			ScriptRef:  "", // TODO: add script ref
			ScriptHash: referenceScriptHash,
		},
	}
}

func BfToAssets(bfAssets []blockfrost.TxAmount) []Asset {
	assets := make([]Asset, len(bfAssets))
	for i, bfAsset := range bfAssets {
		assets[i] = BfToAsset(bfAsset)
	}
	return assets
}

func BfToAsset(bfAsset blockfrost.TxAmount) Asset {
	return Asset{
		Quantity: bfAsset.Quantity,
		Unit:     bfAsset.Unit,
	}
}
