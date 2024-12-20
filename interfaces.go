package rum

type IProvider interface {
	IFetcher
	ISubmitter
}

type IFetcher interface {
	// FetchAccountInfo(address string) (AccountInfo, error)
	// FetchAddressUTxOs(address string, asset *string) ([]UTxO, error)
	// FetchAssetAddresses(asset string) ([]struct {
	// 	Address  string
	// 	Quantity string
	// }, error)
	// FetchAssetMetadata(asset string) (AssetMetadata, error)
	// FetchBlockInfo(hash string) (BlockInfo, error)
	// FetchCollectionAssets(policyId string, cursor *interface{}) (struct {
	// 	Assets Assets
	// 	Next   *interface{}
	// }, error)
	// FetchHandle(handle string) (map[string]interface{}, error)
	// FetchHandleAddress(handle string) (string, error)
	// FetchProtocolParameters(epoch int) (Protocol, error)
	FetchTxInfo(hash string) (TransactionInfo, error)
	FetchUTxOs(hash string, index *int) ([]UTxO, error)
	// Get(url string) (interface{}, error)
}

type ISubmitter interface {
	SubmitTx(txCbor string) (string, error)
}
