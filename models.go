package rum

type AccountInfo struct {
	Active      bool
	PoolId      *string
	Balance     string
	Rewards     string
	Withdrawals string
}

type AssetMetadata = map[string]interface{}

type BlockInfo struct {
	Time                   int
	Hash                   string
	Slot                   string
	Epoch                  int
	EpochSlot              string
	SlotLeader             string
	Size                   int
	TxCount                int
	Output                 string
	Fees                   string
	PreviousBlock          string
	NextBlock              string
	Confirmations          int
	OperationalCertificate string
	VRFKey                 string
}

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

type Protocol struct {
	Epoch                      int
	MinFeeA                    int
	MinFeeB                    int
	MaxBlockSize               int
	MaxTxSize                  int
	MaxBlockHeaderSize         int
	KeyDeposit                 int
	PoolDeposit                int
	Decentralisation           float64
	MinPoolCost                string
	PriceMem                   float64
	PriceStep                  float64
	MaxTxExMem                 string
	MaxTxExSteps               string
	MaxBlockExMem              string
	MaxBlockExSteps            string
	MaxValSize                 int
	CollateralPercent          int
	MaxCollateralInputs        int
	CoinsPerUtxoSize           int
	MinFeeRefScriptCostPerByte int
}

type TransactionInfo struct {
	Index         int
	Block         string
	Hash          string
	Slot          string
	Fees          string
	Size          int
	Deposit       string
	InvalidBefore string
	InvalidAfter  string
}

type Input struct {
	OutputIndex int    `json:"output_index" binding:"required"`
	TxHash      string `json:"tx_hash" binding:"required"`
}

type Output struct {
	Address    string  `json:"address" binding:"required"`
	Amount     []Asset `json:"amount" binding:"required"`
	DataHash   string  `json:"data_hash,omitempty"`
	PlutusData string  `json:"plutus_data,omitempty"`
	ScriptRef  string  `json:"script_ref,omitempty"`
	ScriptHash string  `json:"script_hash,omitempty"`
}

type UTxO struct {
	Input  Input  `json:"input" binding:"required"`
	Output Output `json:"output" binding:"required"`
}

func MakeScriptUtxo(txHash string, outputIndex int, address string, amount []Asset, plutusData string, dataHash string) UTxO {
	return UTxO{
		Input: Input{
			OutputIndex: outputIndex,
			TxHash:      txHash,
		},
		Output: Output{
			Address:    address,
			Amount:     amount,
			DataHash:   dataHash,
			PlutusData: plutusData,
		},
	}
}
