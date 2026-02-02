package rum

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type HttpOgmiosProvider struct {
	*rpc.Client
}

func NewHttpOgmiosProvider(baseUrl string) ISubmitter {
	// assume correct url
	rpcClient, _ := rpc.Dial(baseUrl)

	return &HttpOgmiosProvider{
		rpcClient,
	}
}

type SubmitTxResult struct {
	Transaction struct {
		ID string `json:"id"`
	} `json:"transaction"`
}

func (o *HttpOgmiosProvider) SubmitTx(txCbor string) (string, error) {
	var result SubmitTxResult
	params := map[string]interface{}{
		"transaction": map[string]string{
			"cbor": txCbor,
		},
	}

	err := o.Client.Call(&result, "submitTransaction", params)
	if err != nil {
		return "", err
	}

	return result.Transaction.ID, nil
}
