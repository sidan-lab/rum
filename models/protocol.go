package models

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
