package providers_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/lpernett/godotenv"

	"github.com/sidan-lab/rum/interfaces"
	"github.com/sidan-lab/rum/models"
	"github.com/sidan-lab/rum/providers"
)

const confirmedTxHash = "5101ee2bac21b7e8d7409ee95aeec1444f294248ae85a98c88bbba640cbd132d"
const unconfirmedTxHash = "771ddc23cb2d9bd6a8afe9013d583367bb110090eee6532ad17b8ef31ba05462"

func Init(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
}

func TestMaestroProviderImplementsIFetcher(t *testing.T) {
	var _ interfaces.IFetcher = (*providers.MaestroProvider)(nil)
}
func TestMaestroProviderImplementsISubmitter(t *testing.T) {
	var _ interfaces.ISubmitter = (*providers.MaestroProvider)(nil)
}

func TestMaestroFetchTxInfo(t *testing.T) {
	Init(t)
	txHash := confirmedTxHash
	txHash = unconfirmedTxHash
	network := os.Getenv("NETWORK")
	maestro := providers.NewMaestroProvider(os.Getenv("MAESTRO_API_KEY"), models.Network(network))
	_, err := maestro.FetchTxInfo(txHash)
	if err != nil {
		fmt.Println("failed to fetch tx")
	}
	fmt.Println("success!")
}

func TestMaestroFetchUtxos(t *testing.T) {
	Init(t)
	txHash := "4d2545880f6a6518e6b273875882089c9f3f9955cb3623e9222047e98fc7d1fe"
	network := os.Getenv("NETWORK")
	maestro := providers.NewMaestroProvider(os.Getenv("MAESTRO_API_KEY"), models.Network(network))
	utxos, err := maestro.FetchUTxOs(txHash, nil)
	if err != nil {
		fmt.Println("failed to fetch utxos")
	}
	fmt.Println("success!")
	for _, utxo := range utxos {
		fmt.Printf("utxo: %v\n\n", utxo)
		fmt.Printf("datum: %v\n\n", utxo.Output.PlutusData)
	}
}
