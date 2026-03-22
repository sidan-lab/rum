package rum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpOgmiosProvider struct {
	baseUrl    string
	httpClient *http.Client
}

func NewHttpOgmiosProvider(baseUrl string) ISubmitter {
	return &HttpOgmiosProvider{
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
	}
}

type ogmiosRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type SubmitTxResult struct {
	Transaction struct {
		ID string `json:"id"`
	} `json:"transaction"`
}

type ogmiosResponse struct {
	Result SubmitTxResult `json:"result"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (o *HttpOgmiosProvider) SubmitTx(txCbor string) (string, error) {
	reqBody := ogmiosRequest{
		Jsonrpc: "2.0",
		Method:  "submitTransaction",
		Params: map[string]interface{}{
			"transaction": map[string]string{
				"cbor": txCbor,
			},
		},
		ID: 1,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := o.httpClient.Post(o.baseUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ogmiosResp ogmiosResponse
	if err := json.Unmarshal(body, &ogmiosResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w, body: %s", err, string(body))
	}

	if ogmiosResp.Error != nil {
		return "", fmt.Errorf("%d: %s", ogmiosResp.Error.Code, ogmiosResp.Error.Message)
	}

	return ogmiosResp.Result.Transaction.ID, nil
}
