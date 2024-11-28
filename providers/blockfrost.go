package providers

import (
	"context"
	"encoding/hex"
	"strconv"

	"github.com/blockfrost/blockfrost-go"
	"github.com/sidan-lab/rum/models"
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

func (bf *BlockfrostProvider) FetchTxInfo(hash string) (models.TransactionInfo, error) {
	bfTxInfo, err := bf.blockfrostClient.Transaction(context.TODO(), hash)
	if err != nil {
		return models.TransactionInfo{}, err
	}

	invalidHereafter := ""
	if bfTxInfo.InvalidHereafter != nil {
		invalidHereafter = *bfTxInfo.InvalidHereafter
	}
	invalidBefore := ""
	if bfTxInfo.InvalidBefore != nil {
		invalidBefore = *bfTxInfo.InvalidBefore
	}

	txInfo := models.TransactionInfo{
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
