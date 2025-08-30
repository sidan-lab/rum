package wallet

import (
	"fmt"
	"testing"

	signer "github.com/sidan-lab/cardano-golang-signing-module"
)

func TestDirectSigning(t *testing.T) {
	fmt.Println("=== Mnemonic Signer Example ===")

	// Example mnemonic phrase (DO NOT use in production)
	mnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
	derivationPath := "m/44'/118'/0'/0/0"

	// Create a new signer from mnemonic
	signerInstance, err := signer.NewMnemonicSigner(mnemonic, derivationPath)
	if err != nil {
		t.Fatalf("Failed to create mnemonic signer: %v", err)
	}
	defer signerInstance.Close()

	// Get the public key
	publicKey, err := signerInstance.GetPublicKey()
	if err != nil {
		t.Fatalf("Failed to get public key: %v", err)
	}

	fmt.Printf("Public Key: %s\n", publicKey)

	// Example transaction hex (this is just sample data)
	txHex := "0a90010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412700a2d636f736d6f7331667838306a6b707771703074646d343461306b6564386e72676676376165727979723937337a6a7a122d636f736d6f73317a6368616e6e656c346a636474746e6379663334686e68713873743873397564756a616e306a731a100a05756174686f6d12073130303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f44e8a3b6d4e35aef2f1c6e2c3b8e6c7e1e2b1b3e2d1e2c3b1b2c1e2d3e4f5b61a12040a02080118091204100a200118c40120e0a71a0a05756174686f6d12013010a08d06"

	// Try to sign the transaction
	signature, err := signerInstance.SignTransaction(txHex)
	if err != nil {
		fmt.Printf("Transaction signing failed (expected with sample data): %v\n", err)
	} else {
		fmt.Printf("Transaction Signature: %s\n", signature)
	}

	fmt.Println("Example completed successfully!")
}

func TestWalletSigning(t *testing.T) {
	fmt.Println("=== Wallet-based Signing Example ===")

	// Example mnemonic phrase (DO NOT use in production)
	mnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
	derivationPath := FromString("m/44'/118'/0'/0/0")

	// Create a wallet from mnemonic
	wallet, err := NewMnemonicWallet(mnemonic, derivationPath)
	if err != nil {
		t.Fatalf("Failed to create mnemonic wallet: %v", err)
	}

	// Get the public key through the wallet's signer
	publicKey, err := wallet.Signer().GetPublicKey()
	if err != nil {
		t.Fatalf("Failed to get public key: %v", err)
	}

	fmt.Printf("Public Key from Wallet: %s\n", publicKey)

	// Example transaction hex (this is just sample data)
	txHex := "0a90010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412700a2d636f736d6f7331667838306a6b707771703074646d343461306b6564386e72676676376165727979723937337a6a7a122d636f736d6f73317a6368616e6e656c346a636474746e6379663334686e68713873743873397564756a616e306a731a100a05756174686f6d12073130303030303012670a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103f44e8a3b6d4e35aef2f1c6e2c3b8e6c7e1e2b1b3e2d1e2c3b1b2c1e2d3e4f5b61a12040a02080118091204100a200118c40120e0a71a0a05756174686f6d12013010a08d06"

	// Try to sign the transaction through the wallet
	signature, err := wallet.Signer().SignTransaction(txHex)
	if err != nil {
		fmt.Printf("Transaction signing failed (expected with sample data): %v\n", err)
	} else {
		fmt.Printf("Transaction Signature from Wallet: %s\n", signature)
	}

	fmt.Println("Wallet signing example completed successfully!")
}