package rum

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type KupoProvider struct {
	baseUrl    string
	httpClient *http.Client
}

type kupoApiError struct {
	Hint string `json:"hint"`
}

type kupoUtxo struct {
	TransactionIndex int    `json:"transaction_index"`
	TransactionId    string `json:"transaction_id"`
	OutputIndex      int    `json:"output_index"`
	Address          string `json:"address"`
	Value            struct {
		Coins  int64            `json:"coins"`
		Assets map[string]int64 `json:"assets"`
	} `json:"value"`
	DatumHash  string `json:"datum_hash"`
	Datum      string `json:"datum"`
	DatumType  string `json:"datum_type"`
	ScriptHash string `json:"script_hash"`
	Script     struct {
		Language string `json:"language"`
		Script   string `json:"script"`
	} `json:"script"`
	CreatedAt struct {
		SlotNo     int64  `json:"slot_no"`
		HeaderHash string `json:"header_hash"`
	} `json:"created_at"`
	SpentAt struct {
		SlotNo int64 `json:"slot_no"`
		Header struct {
			Hash string `json:"hash"`
		} `json:"header"`
		TransactionId string `json:"transaction_id"`
		InputIndex    int    `json:"input_index"`
		Redeemer      string `json:"redeemer"`
	} `json:"spent_at"`
}

func NewKupoProvider(baseUrl string, httpClient *http.Client) IFetcher {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &KupoProvider{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}
}

func kupoUtxoToUtxo(kupoUtxo kupoUtxo) UTxO {
	length := len(kupoUtxo.Value.Assets)
	if kupoUtxo.Value.Coins != 0 {
		length++
	}
	amount := make([]Asset, 0, length)
	if kupoUtxo.Value.Coins != 0 {
		amount = append(amount, Asset{Quantity: fmt.Sprintf("%d", kupoUtxo.Value.Coins), Unit: "lovelace"})
	}
	if kupoUtxo.Value.Assets != nil {
		amount = append(amount, kupoAssetsToAssets(kupoUtxo.Value.Assets)...)
	}

	return UTxO{
		Input: Input{
			TxHash:      kupoUtxo.TransactionId,
			OutputIndex: kupoUtxo.OutputIndex,
		},
		Output: Output{
			Amount:     amount,
			Address:    kupoUtxo.Address,
			DataHash:   kupoUtxo.DatumHash,
			PlutusData: kupoUtxo.Datum,
			// TODO: parse script and normalize script
			ScriptRef:  "",
			ScriptHash: kupoUtxo.ScriptHash,
		},
	}
}

func kupoAssetsToAssets(kupoAssets map[string]int64) []Asset {
	assets := make([]Asset, 0, len(kupoAssets))
	for unit, quantity := range kupoAssets {
		assets = append(assets, Asset{
			Quantity: fmt.Sprintf("%d", quantity),
			Unit:     unit,
		})
	}
	return assets
}

func (kp *KupoProvider) FetchAddressUTxOs(address string, asset *string) ([]UTxO, error) {
	var utxos []UTxO
	var url string

	if asset != nil && *asset != "" {
		url = fmt.Sprintf("%s/matches/%s?policy_id=%s&unspent&resolve_hashes", kp.baseUrl, address, *asset)
	} else {
		url = fmt.Sprintf("%s/matches/%s?unspent&resolve_hashes", kp.baseUrl, address)
	}

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := kp.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var kupoErr kupoApiError
		if err := json.NewDecoder(resp.Body).Decode(&kupoErr); err != nil || kupoErr.Hint == "" {
			return nil, fmt.Errorf("unexpected error, status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("Kupo API error: %s", kupoErr.Hint)
	}

	var kupoUtxos []kupoUtxo
	if err := json.NewDecoder(resp.Body).Decode(&kupoUtxos); err != nil {
		return nil, err
	}
	utxos = make([]UTxO, 0, len(kupoUtxos))

	for _, kupoUtxo := range kupoUtxos {
		utxos = append(utxos, kupoUtxoToUtxo(kupoUtxo))
	}

	return utxos, nil
}

func (kp *KupoProvider) FetchTxInfo(hash string) (TransactionInfo, error) {
	url := fmt.Sprintf("%s/matches/*@%s?&unspent&resolve_hashes", kp.baseUrl, hash)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return TransactionInfo{}, err
	}

	resp, err := kp.httpClient.Do(req)
	if err != nil {
		return TransactionInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var kupoErr kupoApiError
		if err := json.NewDecoder(resp.Body).Decode(&kupoErr); err != nil || kupoErr.Hint == "" {
			return TransactionInfo{}, fmt.Errorf("unexpected error, status code: %d", resp.StatusCode)
		}
		return TransactionInfo{}, fmt.Errorf("Kupo API error: %s", kupoErr.Hint)
	}

	var kupoUtxos []kupoUtxo
	if err := json.NewDecoder(resp.Body).Decode(&kupoUtxos); err != nil {
		return TransactionInfo{}, err
	}

	if len(kupoUtxos) == 0 {
		return TransactionInfo{}, fmt.Errorf("transaction not found")
	}

	return TransactionInfo{
		Hash: kupoUtxos[0].CreatedAt.HeaderHash,
		Slot: fmt.Sprintf("%d", kupoUtxos[0].CreatedAt.SlotNo),
	}, nil
}

func (kp *KupoProvider) FetchUTxOs(hash string, index *int) ([]UTxO, error) {
	var utxos []UTxO
	url := fmt.Sprintf("%s/matches/*@%s?&unspent&resolve_hashes", kp.baseUrl, hash)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := kp.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var kupoErr kupoApiError
		if err := json.NewDecoder(resp.Body).Decode(&kupoErr); err != nil || kupoErr.Hint == "" {
			return nil, fmt.Errorf("unexpected error, status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("Kupo API error: %s", kupoErr.Hint)
	}

	var kupoUtxos []kupoUtxo
	if err := json.NewDecoder(resp.Body).Decode(&kupoUtxos); err != nil {
		return nil, err
	}

	utxos = make([]UTxO, 0, len(kupoUtxos))

	for _, kupoUtxo := range kupoUtxos {
		utxos = append(utxos, kupoUtxoToUtxo(kupoUtxo))
	}

	return utxos, nil
}
