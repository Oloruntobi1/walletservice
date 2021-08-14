package wallet

import (
	"fmt"

	"github.com/Oloruntobi1/payload"
)

type Wallet interface {
	Create() string
	Payout() payload.WalletPayoutResponse
}

func CreateWallet(w Wallet) (string, error){
	if w == nil {
		return "", fmt.Errorf("no wallet provider currently")
	}

	return w.Create(), nil
}

func PayoutUser(w Wallet) (payload.WalletPayoutResponse, error){
	if w == nil {
		return payload.WalletPayoutResponse{}, fmt.Errorf("no wallet provider currently")
	}

	return w.Payout(), nil
}