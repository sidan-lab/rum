package wallet

import (
	"testing"
)

func TestNewDerivationIndices(t *testing.T) {
	indices := NewDerivationIndices()
	expected := []uint32{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart,
		0,
		0,
	}
	
	if len(indices) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(indices))
	}
	
	for i, val := range expected {
		if indices[i] != val {
			t.Errorf("Expected indices[%d] = %d, got %d", i, val, indices[i])
		}
	}
}

func TestPaymentDerivation(t *testing.T) {
	indices := PaymentDerivation(0, 5)
	expected := []uint32{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart + 0,
		0,
		5,
	}
	
	for i, val := range expected {
		if indices[i] != val {
			t.Errorf("Expected indices[%d] = %d, got %d", i, val, indices[i])
		}
	}
}

func TestStakeDerivation(t *testing.T) {
	indices := StakeDerivation(1, 3)
	expected := []uint32{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart + 1,
		2,
		3,
	}
	
	for i, val := range expected {
		if indices[i] != val {
			t.Errorf("Expected indices[%d] = %d, got %d", i, val, indices[i])
		}
	}
}

func TestDRepDerivation(t *testing.T) {
	indices := DRepDerivation(2, 7)
	expected := []uint32{
		HardenedKeyStart + 1852,
		HardenedKeyStart + 1815,
		HardenedKeyStart + 2,
		3,
		7,
	}
	
	for i, val := range expected {
		if indices[i] != val {
			t.Errorf("Expected indices[%d] = %d, got %d", i, val, indices[i])
		}
	}
}

func TestFromString(t *testing.T) {
	testCases := []struct {
		input    string
		expected []uint32
	}{
		{
			"m/1852'/1815'/0'/0/0",
			[]uint32{HardenedKeyStart + 1852, HardenedKeyStart + 1815, HardenedKeyStart + 0, 0, 0},
		},
		{
			"1852'/1815'/0'/2/5",
			[]uint32{HardenedKeyStart + 1852, HardenedKeyStart + 1815, HardenedKeyStart + 0, 2, 5},
		},
		{
			"m/44'/0'/0'/0/1",
			[]uint32{HardenedKeyStart + 44, HardenedKeyStart + 0, HardenedKeyStart + 0, 0, 1},
		},
	}
	
	for _, tc := range testCases {
		indices := FromString(tc.input)
		if len(indices) != len(tc.expected) {
			t.Errorf("For input %s, expected length %d, got %d", tc.input, len(tc.expected), len(indices))
			continue
		}
		
		for i, val := range tc.expected {
			if indices[i] != val {
				t.Errorf("For input %s, expected indices[%d] = %d, got %d", tc.input, i, val, indices[i])
			}
		}
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		indices  DerivationIndices
		expected string
	}{
		{
			DerivationIndices{HardenedKeyStart + 1852, HardenedKeyStart + 1815, HardenedKeyStart + 0, 0, 0},
			"m/1852'/1815'/0'/0/0",
		},
		{
			DerivationIndices{HardenedKeyStart + 44, HardenedKeyStart + 0, HardenedKeyStart + 0, 0, 1},
			"m/44'/0'/0'/0/1",
		},
		{
			DerivationIndices{},
			"m",
		},
	}
	
	for _, tc := range testCases {
		result := tc.indices.ToString()
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	original := "m/1852'/1815'/0'/2/5"
	indices := FromString(original)
	result := indices.ToString()
	
	if result != original {
		t.Errorf("Round trip failed: original=%s, result=%s", original, result)
	}
}

func TestHelperMethods(t *testing.T) {
	indices := PaymentDerivation(0, 5)
	
	if indices.Len() != 5 {
		t.Errorf("Expected length 5, got %d", indices.Len())
	}
	
	if indices.Get(4) != 5 {
		t.Errorf("Expected Get(4) = 5, got %d", indices.Get(4))
	}
	
	if indices.Get(10) != 0 {
		t.Errorf("Expected Get(10) = 0 (out of bounds), got %d", indices.Get(10))
	}
	
	slice := indices.ToSlice()
	if len(slice) != 5 {
		t.Errorf("Expected slice length 5, got %d", len(slice))
	}
}