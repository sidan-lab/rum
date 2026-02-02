package rum

import (
	"context"
	"fmt"

	"github.com/SundaeSwap-finance/ogmigo"
)

type OgmiosProvider struct {
	*ogmigo.Client
}

func NewOgmiosProvider(baseUrl string, logger ogmigo.Logger) ISubmitter {
	options := make([]ogmigo.Option, 0)

	options = append(options, ogmigo.WithEndpoint(baseUrl))

	if logger != nil {
		options = append(options, ogmigo.WithLogger(logger))
	}

	return &OgmiosProvider{
		Client: ogmigo.New(options...),
	}
}

func (o *OgmiosProvider) SubmitTx(txCbor string) (string, error) {
	resp, err := o.Client.SubmitTx(context.TODO(), txCbor)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", fmt.Errorf("Submit Tx Error: %s", resp.Error)
	}

	return resp.ID, nil
}
