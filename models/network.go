package models

type Network string

const (
	Testnet Network = "testnet"
	Preview Network = "preview"
	Preprod Network = "preprod"
	Mainnet Network = "mainnet"
)

var AllNetworks = []Network{Testnet, Preview, Preprod, Mainnet}

func (n Network) String() string {
	return string(n)
}

func IsNetwork(value string) bool {
	for _, network := range AllNetworks {
		if value == network.String() {
			return true
		}
	}
	return false
}
