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
