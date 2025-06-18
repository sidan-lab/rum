package models

type Protocol struct {
	Epoch                      int32
	MinFeeA                    uint64
	MinFeeB                    uint64
	MaxBlockSize               int32
	MaxTxSize                  uint32
	MaxBlockHeaderSize         int32
	KeyDeposit                 uint64
	PoolDeposit                uint64
	Decentralisation           float64
	MinPoolCost                string
	PriceMem                   float64
	PriceStep                  float64
	MaxTxExMem                 string
	MaxTxExSteps               string
	MaxBlockExMem              string
	MaxBlockExSteps            string
	MaxValSize                 uint32
	CollateralPercent          float64
	MaxCollateralInputs        int32
	CoinsPerUtxoSize           uint64
	MinFeeRefScriptCostPerByte uint64
}

func Default() *Protocol {
	return &Protocol{
		Epoch:                      0,
		MinFeeA:                    44,
		MinFeeB:                    155381,
		MaxBlockSize:               98304,
		MaxTxSize:                  16384,
		MaxBlockHeaderSize:         1100,
		KeyDeposit:                 2000000,
		PoolDeposit:                500000000,
		MinPoolCost:                "340000000",
		PriceMem:                   0.0577,
		PriceStep:                  0.0000721,
		MaxTxExMem:                 "16000000",
		MaxTxExSteps:               "10000000000",
		MaxBlockExMem:              "80000000",
		MaxBlockExSteps:            "40000000000",
		MaxValSize:                 5000,
		CollateralPercent:          150.0,
		MaxCollateralInputs:        3,
		CoinsPerUtxoSize:           4310,
		MinFeeRefScriptCostPerByte: 15,
		Decentralisation:           0.0,
	}
}
