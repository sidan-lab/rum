package wallet

import (
	"strconv"
	"strings"
)

const HardenedKeyStart = 0x80000000

type DerivationIndices []uint32

func NewDerivationIndices() DerivationIndices {
	return DerivationIndices{
		HardenedKeyStart + 1852, // purpose
		HardenedKeyStart + 1815, // coin type
		HardenedKeyStart,        // account
		0,                       // payment
		0,                       // key index
	}
}

func PaymentDerivation(accountIndex, keyIndex uint32) DerivationIndices {
	return DerivationIndices{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart + accountIndex,
		0,
		keyIndex,
	}
}

func StakeDerivation(accountIndex, keyIndex uint32) DerivationIndices {
	return DerivationIndices{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart + accountIndex,
		2,
		keyIndex,
	}
}

func DRepDerivation(accountIndex, keyIndex uint32) DerivationIndices {
	return DerivationIndices{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart + accountIndex,
		3,
		keyIndex,
	}
}

func FromString(derivationPathStr string) DerivationIndices {
	derivationPathStr = strings.TrimPrefix(derivationPathStr, "m/")
	
	parts := strings.Split(derivationPathStr, "/")
	indices := make([]uint32, 0, len(parts))
	
	for _, part := range parts {
		if part == "" {
			continue
		}
		
		var value uint64
		var err error
		
		if strings.HasSuffix(part, "'") {
			part = strings.TrimSuffix(part, "'")
			value, err = strconv.ParseUint(part, 10, 32)
			if err != nil {
				continue
			}
			indices = append(indices, uint32(value)+HardenedKeyStart)
		} else {
			value, err = strconv.ParseUint(part, 10, 32)
			if err != nil {
				continue
			}
			indices = append(indices, uint32(value))
		}
	}
	
	return DerivationIndices(indices)
}

func (d DerivationIndices) ToString() string {
	if len(d) == 0 {
		return "m"
	}
	
	parts := make([]string, 0, len(d)+1)
	parts = append(parts, "m")
	
	for _, index := range d {
		if index >= HardenedKeyStart {
			parts = append(parts, strconv.FormatUint(uint64(index-HardenedKeyStart), 10)+"'")
		} else {
			parts = append(parts, strconv.FormatUint(uint64(index), 10))
		}
	}
	
	return strings.Join(parts, "/")
}

func (d DerivationIndices) ToSlice() []uint32 {
	return []uint32(d)
}

func (d DerivationIndices) Len() int {
	return len(d)
}

func (d DerivationIndices) Get(index int) uint32 {
	if index < 0 || index >= len(d) {
		return 0
	}
	return d[index]
}