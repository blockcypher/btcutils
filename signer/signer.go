package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec"
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

	sig, err := Sign(priv, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(sig))
}

func Sign(private, data []byte) ([]byte, error) {
	privkey, _ := btcec.PrivKeyFromBytes(btcec.S256(), private)
	sig, err := privkey.Sign(data)
	if err != nil {
		return nil, err
	}
	return sig.Serialize(), nil
}
