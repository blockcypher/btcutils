package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Signs the provided data (hex encoded) with the provided private key (hex encoded) " +
			"using the Bitcoin ECDSA curve.")
		fmt.Println("Usage: signer [datahex] [privatehex]")
		return
	}
	data, err := hex.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	priv, err := hex.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	sig := Sign(priv, data)
	if sig == nil {
		err = errors.New("error during signing process")
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(sig))
}

func Sign(private, data []byte) []byte {
	privkey, _ := btcec.PrivKeyFromBytes(private)
	sig := ecdsa.Sign(privkey, data)
	if sig == nil {
		return nil
	}
	return sig.Serialize()
}
