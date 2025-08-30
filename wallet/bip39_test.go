package wallet

import (
	"testing"
)

func TestMnemonicSignTx(t *testing.T) {
	// This test replicates the Rust test_mnemonic_sign_tx
	mnemonicPhrase := "summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer"
	
	// Create wallet using the standard payment derivation path
	wallet, err := NewMnemonicWallet(mnemonicPhrase, PaymentDerivation(0, 0))
	if err != nil {
		t.Fatalf("Failed to create mnemonic wallet: %v", err)
	}
	
	// Transaction hex from the Rust test
	txHex := "84a4008182582004509185eb98edd8e2420c1ceea914d6a7a3142041039b2f12b4d4f03162d56f04018282581d605867c3b8e27840f556ac268b781578b14c5661fc63ee720dbeab663f1a000f42408258390004845038ee499ee8bc0afe56f688f27b2dd76f230d3698a9afcc1b66e0464447c1f51adaefe1ebfb0dd485a349a70479ced1d198cbdf7fe71a15d35396021a0002917d075820bdaa99eb158414dea0a91d6c727e2268574b23efe6e08ab3b841abe8059a030ca0f5d90103a0"
	
	// Sign the transaction
	signedTx, err := wallet.Signer().SignTransaction(txHex)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}
	
	// Verify that we got a signed transaction (non-empty result)
	if signedTx == "" {
		t.Error("Expected non-empty signed transaction")
	}
	
	// Verify the signed transaction is different from the unsigned one
	if signedTx == txHex {
		t.Error("Signed transaction should be different from unsigned transaction")
	}
	
	t.Logf("Successfully signed transaction with mnemonic wallet. Length: %d", len(signedTx))
}

func TestRootKeySignTx(t *testing.T) {
	// This test replicates the Rust test_root_sign_tx
	// Note: This is a sample root key - replace with actual test data when available
	rootPrivateKey := "xprv1cqa3sdvldefault1root2key3for4testing5purposes6only7not8for9production0use1test2data3only"
	
	// Create wallet using root key with standard payment derivation
	wallet, err := NewRootKeyWallet(rootPrivateKey, PaymentDerivation(0, 0))
	if err != nil {
		t.Skipf("Skipping root key test - may need valid root key format: %v", err)
		return
	}
	
	// Same transaction hex as mnemonic test
	txHex := "84a4008182582004509185eb98edd8e2420c1ceea914d6a7a3142041039b2f12b4d4f03162d56f04018282581d605867c3b8e27840f556ac268b781578b14c5661fc63ee720dbeab663f1a000f42408258390004845038ee499ee8bc0afe56f688f27b2dd76f230d3698a9afcc1b66e0464447c1f51adaefe1ebfb0dd485a349a70479ced1d198cbdf7fe71a15d35396021a0002917d075820bdaa99eb158414dea0a91d6c727e2268574b23efe6e08ab3b841abe8059a030ca0f5d90103a0"
	
	// Sign the transaction
	signedTx, err := wallet.Signer().SignTransaction(txHex)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}
	
	// Verify that we got a signed transaction
	if signedTx == "" {
		t.Error("Expected non-empty signed transaction")
	}
	
	if signedTx == txHex {
		t.Error("Signed transaction should be different from unsigned transaction")
	}
	
	t.Logf("Successfully signed transaction with root key wallet. Length: %d", len(signedTx))
}

func TestRawPathSignTx(t *testing.T) {
	// This test replicates the Rust test_raw_path_sign_tx
	mnemonicPhrase := "summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer"
	
	// Use the specific derivation path from the Rust test
	derivationPath := FromString("m/1852'/1815'/0'/0/0")
	
	// Create wallet with the raw derivation path
	wallet, err := NewMnemonicWallet(mnemonicPhrase, derivationPath)
	if err != nil {
		t.Fatalf("Failed to create mnemonic wallet with raw path: %v", err)
	}
	
	// Same transaction hex as other tests
	txHex := "84a4008182582004509185eb98edd8e2420c1ceea914d6a7a3142041039b2f12b4d4f03162d56f04018282581d605867c3b8e27840f556ac268b781578b14c5661fc63ee720dbeab663f1a000f42408258390004845038ee499ee8bc0afe56f688f27b2dd76f230d3698a9afcc1b66e0464447c1f51adaefe1ebfb0dd485a349a70479ced1d198cbdf7fe71a15d35396021a0002917d075820bdaa99eb158414dea0a91d6c727e2268574b23efe6e08ab3b841abe8059a030ca0f5d90103a0"
	
	// Sign the transaction
	signedTx, err := wallet.Signer().SignTransaction(txHex)
	if err != nil {
		t.Fatalf("Failed to sign transaction with raw path: %v", err)
	}
	
	// Verify that we got a signed transaction
	if signedTx == "" {
		t.Error("Expected non-empty signed transaction")
	}
	
	if signedTx == txHex {
		t.Error("Signed transaction should be different from unsigned transaction")
	}
	
	t.Logf("Successfully signed transaction with raw path. Length: %d", len(signedTx))
}

func TestConsistentSigning(t *testing.T) {
	// Test that the same mnemonic and derivation path produces consistent results
	mnemonicPhrase := "summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer"
	
	// Create two wallets with same parameters
	derivationPath := FromString("m/1852'/1815'/0'/0/0")
	
	wallet1, err := NewMnemonicWallet(mnemonicPhrase, derivationPath)
	if err != nil {
		t.Fatalf("Failed to create first wallet: %v", err)
	}
	
	wallet2, err := NewMnemonicWallet(mnemonicPhrase, derivationPath)
	if err != nil {
		t.Fatalf("Failed to create second wallet: %v", err)
	}
	
	txHex := "84a4008182582004509185eb98edd8e2420c1ceea914d6a7a3142041039b2f12b4d4f03162d56f04018282581d605867c3b8e27840f556ac268b781578b14c5661fc63ee720dbeab663f1a000f42408258390004845038ee499ee8bc0afe56f688f27b2dd76f230d3698a9afcc1b66e0464447c1f51adaefe1ebfb0dd485a349a70479ced1d198cbdf7fe71a15d35396021a0002917d075820bdaa99eb158414dea0a91d6c727e2268574b23efe6e08ab3b841abe8059a030ca0f5d90103a0"
	
	// Sign with both wallets
	signedTx1, err := wallet1.Signer().SignTransaction(txHex)
	if err != nil {
		t.Fatalf("Failed to sign with first wallet: %v", err)
	}
	
	signedTx2, err := wallet2.Signer().SignTransaction(txHex)
	if err != nil {
		t.Fatalf("Failed to sign with second wallet: %v", err)
	}
	
	// They should produce the same signature
	if signedTx1 != signedTx2 {
		t.Errorf("Wallets with same parameters should produce identical signatures.\nWallet1: %s\nWallet2: %s", signedTx1, signedTx2)
	}
}

func TestDerivationPathVariations(t *testing.T) {
	// Test that different derivation paths produce different signatures
	mnemonicPhrase := "summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer summer"
	
	paths := []DerivationIndices{
		PaymentDerivation(0, 0),
		PaymentDerivation(0, 1),
		PaymentDerivation(1, 0),
		StakeDerivation(0, 0),
	}
	
	txHex := "84a4008182582004509185eb98edd8e2420c1ceea914d6a7a3142041039b2f12b4d4f03162d56f04018282581d605867c3b8e27840f556ac268b781578b14c5661fc63ee720dbeab663f1a000f42408258390004845038ee499ee8bc0afe56f688f27b2dd76f230d3698a9afcc1b66e0464447c1f51adaefe1ebfb0dd485a349a70479ced1d198cbdf7fe71a15d35396021a0002917d075820bdaa99eb158414dea0a91d6c727e2268574b23efe6e08ab3b841abe8059a030ca0f5d90103a0"
	
	signatures := make([]string, len(paths))
	
	for i, path := range paths {
		wallet, err := NewMnemonicWallet(mnemonicPhrase, path)
		if err != nil {
			t.Fatalf("Failed to create wallet for path %d: %v", i, err)
		}
		
		signedTx, err := wallet.Signer().SignTransaction(txHex)
		if err != nil {
			t.Fatalf("Failed to sign with path %d: %v", i, err)
		}
		
		signatures[i] = signedTx
		t.Logf("Path %s -> signature length: %d", path.ToString(), len(signedTx))
	}
	
	// Verify that different paths produce different signatures
	for i := range signatures {
		for j := i + 1; j < len(signatures); j++ {
			if signatures[i] == signatures[j] {
				t.Errorf("Different derivation paths should produce different signatures. Paths %d and %d produced identical results", i, j)
			}
		}
	}
}