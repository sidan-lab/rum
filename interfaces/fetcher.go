package interfaces

import "github.com/sidan-lab/rum/models"

type IFetcher interface {
	// FetchAccountInfo(address string) (models.AccountInfo, error)
	// FetchAddressUTxOs(address string, asset *string) ([]models.UTxO, error)
	// FetchAssetAddresses(asset string) ([]struct {
	// 	Address  string
	// 	Quantity string
	// }, error)
	// FetchAssetMetadata(asset string) (models.AssetMetadata, error)
	// FetchBlockInfo(hash string) (models.BlockInfo, error)
	// FetchCollectionAssets(policyId string, cursor *interface{}) (struct {
	// 	Assets models.Assets
	// 	Next   *interface{}
	// }, error)
	// FetchHandle(handle string) (map[string]interface{}, error)
	// FetchHandleAddress(handle string) (string, error)
	// FetchProtocolParameters(epoch int) (models.Protocol, error)
	FetchTxInfo(hash string) (models.TransactionInfo, error)
	// FetchUTxOs(hash string, index *int) ([]models.UTxO, error)
	// Get(url string) (interface{}, error)
}
