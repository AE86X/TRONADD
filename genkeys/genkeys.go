package genkeys

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
	addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func GenerateKey() (wif string, address string) {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	if len(priv.D.Bytes()) != 32 {
		for {
			priv, err := btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(priv.D.Bytes()) == 32 {
				break
			}
		}
	}
	a := addr.PubkeyToAddress(priv.ToECDSA().PublicKey)
	address = a.String()
	wif = hex.EncodeToString(priv.D.Bytes())
	return
}
