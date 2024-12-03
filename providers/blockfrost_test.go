package providers_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sidan-lab/rum/interfaces"
	"github.com/sidan-lab/rum/providers"
)

func TestBlockfrostProviderImplementsIFetcher(t *testing.T) {
	var _ interfaces.IFetcher = (*providers.BlockfrostProvider)(nil)
}
func TestBlockfrostProviderImplementsISubmitter(t *testing.T) {
	var _ interfaces.ISubmitter = (*providers.BlockfrostProvider)(nil)
}

func TestBlockfrostFetchTxInfo(t *testing.T) {
	Init(t)
	txHash := confirmedTxHash
	// txHash = unconfirmedTxHash
	blockfrost := providers.NewBlockfrostProvider(os.Getenv("BLOCKFROST_PROJECT_ID"))
	res, err := blockfrost.FetchTxInfo(txHash)
	if err != nil {
		fmt.Println("failed to fetch tx")
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("success! %v", res)
}

func TestBlockfrostFetchUtxos(t *testing.T) {
	Init(t)
	txHash := "4d2545880f6a6518e6b273875882089c9f3f9955cb3623e9222047e98fc7d1fe"
	blockfrost := providers.NewBlockfrostProvider(os.Getenv("BLOCKFROST_PROJECT_ID"))
	utxos, err := blockfrost.FetchUTxOs(txHash, nil)
	if err != nil {
		fmt.Println("failed to fetch utxos")
	}
	fmt.Println("success!")
	for _, utxo := range utxos {
		fmt.Printf("utxo: %v\n\n", utxo)
		fmt.Printf("datum: %v\n\n", utxo.Output.PlutusData)
	}
}
