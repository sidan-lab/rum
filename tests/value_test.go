package rum_test

import (
	"testing"

	"github.com/sidan-lab/rum"
)

func TestAddAsset(t *testing.T) {
	asset := rum.Asset{Unit: "USD", Quantity: "100"}
	mv := rum.NewValue().AddAsset(&asset)
	if mv.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv.Get("USD"))
	}
}

func TestAddAssets(t *testing.T) {
	assets := []rum.Asset{
		{Unit: "USD", Quantity: "100"},
		{Unit: "EUR", Quantity: "200"},
	}
	mv := rum.NewValue().AddAssets(&assets)
	if mv.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv.Get("USD"))
	}
	if mv.Get("EUR") != 200 {
		t.Errorf("Expected 200, got %d", mv.Get("EUR"))
	}
}

func TestNegateAsset(t *testing.T) {
	asset := rum.Asset{Unit: "USD", Quantity: "100"}
	mv := rum.NewValue().AddAsset(&asset).NegateAsset(&asset)

	if mv.Get("USD") != 0 {
		t.Errorf("Expected 0, got %d", mv.Get("USD"))
	}
}

func TestNegateAssets(t *testing.T) {
	assets := []rum.Asset{
		{Unit: "USD", Quantity: "100"},
		{Unit: "EUR", Quantity: "200"},
	}
	mv := rum.NewValue().AddAssets(&assets).NegateAssets(&assets)
	if mv.Get("USD") != 0 {
		t.Errorf("Expected 0, got %d", mv.Get("USD"))
	}
	if mv.Get("EUR") != 0 {
		t.Errorf("Expected 0, got %d", mv.Get("EUR"))
	}
}

func TestMerge(t *testing.T) {
	mv1 := rum.NewValue().AddAsset(&rum.Asset{Unit: "USD", Quantity: "100"})
	mv2 := rum.NewValue().AddAsset(&rum.Asset{Unit: "USD", Quantity: "200"}).Merge(mv1)
	if mv2.Get("USD") != 300 {
		t.Errorf("Expected 300, got %d", mv1.Get("USD"))
	}
}

func TestMergeNil(t *testing.T) {
	mv1 := rum.NewValue()
	mv1.AddAsset(&rum.Asset{Unit: "USD", Quantity: "100"})
	mv1.Merge(nil)

	if mv1.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv1.Get("USD"))
	}
}

func TestMergeFromNewMap(t *testing.T) {
	mv1 := rum.NewValue()
	mv1.AddAsset(&rum.Asset{Unit: "USD", Quantity: "100"})

	mv2 := rum.NewValue()
	mv2.Merge(mv1)

	if mv1.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv1.Get("USD"))
	}
}

func TestToAssets(t *testing.T) {
	mv := rum.NewValue()
	mv.AddAsset(&rum.Asset{Unit: "USD", Quantity: "100"})
	assets := *mv.ToAssets()

	if len(assets) != 1 {
		t.Errorf("Expected 1 asset, got %d", len(assets))
	}
	if assets[0].Unit != "USD" || assets[0].Quantity != "100" {
		t.Errorf("Expected asset with Unit USD and Quantity 100, got Unit %s and Quantity %s", assets[0].Unit, assets[0].Quantity)
	}
}

func TestFromAssets(t *testing.T) {
	assets := []rum.Asset{
		{Unit: "USD", Quantity: "100"},
		{Unit: "EUR", Quantity: "200"},
	}
	mv := rum.NewValueFromAssets(&assets)

	if mv.Get("USD") != 100 {
		t.Errorf("Expected 100, got %d", mv.Get("USD"))
	}
	if mv.Get("EUR") != 200 {
		t.Errorf("Expected 200, got %d", mv.Get("EUR"))
	}
}

func TestGeq(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		value1   *rum.Value
		value2   *rum.Value
		expected bool
	}{
		{
			name: "Value1 greater than or equal to Value2",
			value1: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
				},
			},
			value2: &rum.Value{
				Value: map[string]int64{
					"asset1": 500,
					"asset2": 1500,
				},
			},
			expected: true,
		},
		{
			name: "Value1 less than Value2",
			value1: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 1000,
				},
			},
			value2: &rum.Value{
				Value: map[string]int64{
					"asset1": 1500,
					"asset2": 1500,
				},
			},
			expected: false,
		},
		{
			name: "Value1 equal to Value2",
			value1: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
				},
			},
			value2: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
				},
			},
			expected: true,
		},
		{
			name: "Value1 has more assets than Value2",
			value1: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
					"asset3": 3000,
				},
			},
			value2: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
				},
			},
			expected: true,
		},
		{
			name: "Value1 has fewer assets than Value2",
			value1: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
				},
			},
			value2: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
				},
			},
			expected: false,
		},
		{
			name: "Value1 has fewer assets than Value2",
			value1: &rum.Value{
				Value: map[string]int64{
					"asset1": 1500,
					"asset2": 1500,
				},
			},
			value2: &rum.Value{
				Value: map[string]int64{
					"asset1": 1000,
					"asset2": 2000,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.value1.Geq(tt.value2)
			if result != tt.expected {
				t.Errorf("Geq() = %v, want %v", result, tt.expected)
			}
		})
	}
}
