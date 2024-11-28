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
