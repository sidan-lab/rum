package wallet

import (
	signer "github.com/sidan-lab/cardano-golang-signing-module"
)

type WalletType int

const (
	WalletTypeCli WalletType = iota
	WalletTypeMnemonic
	WalletTypeRootKey
)

type Account struct {
	PrivateKey string
	PublicKey  string
}

type Wallet struct {
	WalletType WalletType
	Account    Account
	signer     *signer.Signer
}

func NewCliWallet(ed25519Key string) (*Wallet, error) {
	walletSigner, err := signer.NewCLISigner(ed25519Key)
	if err != nil {
		return nil, err
	}
	return &Wallet{
		WalletType: WalletTypeCli,
		Account:    Account{PrivateKey: ed25519Key},
		signer:     walletSigner,
	}, nil
}

func NewRootKeyWallet(rootPrivateKey string, derivationPath DerivationIndices) (*Wallet, error) {
	walletSigner, err := signer.NewBech32Signer(rootPrivateKey, derivationPath.ToString())
	if err != nil {
		return nil, err
	}
	return &Wallet{
		WalletType: WalletTypeRootKey,
		Account:    Account{PrivateKey: rootPrivateKey},
		signer:     walletSigner,
	}, nil
}

func NewMnemonicWallet(mnemonic string, derivationPath DerivationIndices) (*Wallet, error) {
	walletSigner, err := signer.NewMnemonicSigner(mnemonic, derivationPath.ToString())
	if err != nil {
		return nil, err
	}
	return &Wallet{
		WalletType: WalletTypeMnemonic,
		Account:    Account{PrivateKey: mnemonic},
		signer:     walletSigner,
	}, nil
}

func (w *Wallet) PaymentAccount(accountIndex, keyIndex uint32) *Wallet {
	derivationPath := PaymentDerivation(accountIndex, keyIndex)
	newWallet := *w

	switch w.WalletType {
	case WalletTypeMnemonic:
		if signer, err := signer.NewMnemonicSigner(w.Account.PrivateKey, derivationPath.ToString()); err == nil {
			newWallet.signer = signer
		}
	case WalletTypeRootKey:
		if signer, err := signer.NewBech32Signer(w.Account.PrivateKey, derivationPath.ToString()); err == nil {
			newWallet.signer = signer
		}
	}

	return &newWallet
}

func (w *Wallet) StakeAccount(accountIndex, keyIndex uint32) *Wallet {
	derivationPath := StakeDerivation(accountIndex, keyIndex)
	newWallet := *w

	switch w.WalletType {
	case WalletTypeMnemonic:
		if signer, err := signer.NewMnemonicSigner(w.Account.PrivateKey, derivationPath.ToString()); err == nil {
			newWallet.signer = signer
		}
	case WalletTypeRootKey:
		if signer, err := signer.NewBech32Signer(w.Account.PrivateKey, derivationPath.ToString()); err == nil {
			newWallet.signer = signer
		}
	}

	return &newWallet
}

func (w *Wallet) DRepAccount(accountIndex, keyIndex uint32) *Wallet {
	derivationPath := DRepDerivation(accountIndex, keyIndex)
	newWallet := *w

	switch w.WalletType {
	case WalletTypeMnemonic:
		if signer, err := signer.NewMnemonicSigner(w.Account.PrivateKey, derivationPath.ToString()); err == nil {
			newWallet.signer = signer
		}
	case WalletTypeRootKey:
		if signer, err := signer.NewBech32Signer(w.Account.PrivateKey, derivationPath.ToString()); err == nil {
			newWallet.signer = signer
		}
	}

	return &newWallet
}

func (w *Wallet) Signer() *signer.Signer {
	return w.signer
}
