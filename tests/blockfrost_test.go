package rum_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sidan-lab/rum"
)

func TestBlockfrostProviderImplementsIFetcher(t *testing.T) {
	var _ rum.IFetcher = (*rum.BlockfrostProvider)(nil)
}
func TestBlockfrostProviderImplementsISubmitter(t *testing.T) {
	var _ rum.ISubmitter = (*rum.BlockfrostProvider)(nil)
}

func TestBlockfrostFetchTxInfo(t *testing.T) {
	Init(t)
	txHash := confirmedTxHash
	// txHash = unconfirmedTxHash
	blockfrost := rum.NewBlockfrostProvider(os.Getenv("BLOCKFROST_PROJECT_ID"))
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
	blockfrost := rum.NewBlockfrostProvider(os.Getenv("BLOCKFROST_PROJECT_ID"))
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

func TestBlockfrostFetchAddressUtxos(t *testing.T) {
	Init(t)
	address := "addr_test1qpvx0sacufuypa2k4sngk7q40zc5c4npl337uusdh64kv0uafhxhu32dys6pvn6wlw8dav6cmp4pmtv7cc3yel9uu0nq93swx9"
	blockfrost := rum.NewBlockfrostProvider(os.Getenv("BLOCKFROST_PROJECT_ID"))
	utxos, err := blockfrost.FetchAddressUTxOs(address, nil)
	if err != nil {
		fmt.Printf("failed to fetch address utxos: %v\n", err)
		t.Fail()
		return
	}
	fmt.Printf("success! found %d utxos\n", len(utxos))
	for _, utxo := range utxos {
		fmt.Printf("utxo: %v\n\n", utxo)
	}
}

func TestBlockfrostFetchAddressUtxosWithAsset(t *testing.T) {
	Init(t)
	address := "addr_test1qpvx0sacufuypa2k4sngk7q40zc5c4npl337uusdh64kv0uafhxhu32dys6pvn6wlw8dav6cmp4pmtv7cc3yel9uu0nq93swx9"
	asset := "c69b981db7a65e339a6d783755f85a2e03afa1cece9714c55fe4c9135553444d"
	blockfrost := rum.NewBlockfrostProvider(os.Getenv("BLOCKFROST_PROJECT_ID"))
	utxos, err := blockfrost.FetchAddressUTxOs(address, &asset)
	if err != nil {
		fmt.Printf("failed to fetch address utxos with asset: %v\n", err)
		t.Fail()
		return
	}
	fmt.Printf("success! found %d utxos with asset %s\n", len(utxos), asset)
	for _, utxo := range utxos {
		fmt.Printf("utxo: %v\n\n", utxo)
	}
}
