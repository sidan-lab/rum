package providers

import (
	"strconv"

	"github.com/maestro-org/go-sdk/client"
	"github.com/sidan-lab/rum/models"
)

type MaestroProvider struct {
	maestroClient *client.Client
}

func NewMaestroProvider(apiKey string, network models.Network) *MaestroProvider {
	maestroClient := client.NewClient(apiKey, string(network))
	return &MaestroProvider{
		maestroClient: maestroClient,
	}
}

func (ms *MaestroProvider) SubmitTx(txCbor string) (string, error) {
	txHash, err := ms.maestroClient.TxManagerSubmit(txCbor)
	return txHash, err
}

func (ms *MaestroProvider) FetchTxInfo(hash string) (models.TransactionInfo, error) {
	tx, err := ms.maestroClient.TransactionDetails(hash)
	if err != nil {
		return models.TransactionInfo{}, err
	}
	msTxInfo := tx.Data
	txInfo := models.TransactionInfo{
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

func formatOptionalInt(value *int64) string {
	if value != nil {
		return strconv.FormatInt(*value, 10)
	}
	return ""
}
