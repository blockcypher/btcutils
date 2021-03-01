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
	neo> relay {"type":"Neo.Network.P2P.Payloads.Transaction","hex":"0071c0992d42e2a62c8b763b5de5b0e1b2e239a7bbd2952a0c00e1f50500000000ac0c240000000000cb152300000142e2a62c8b763b5de5b0e1b2e239a7bbd2952a0c01550400c2eb0b146c93f190909dea8dfe3caeb2ee90530b4ef21e861442e2a62c8b763b5de5b0e1b2e239a7bbd2952a0c53c1087472616e73666572142582d1b275e86c8f0e93a9b0xdcf144d9ed2d64482fb5caafa719cf6706e9afd607ab043e8bfcb9018795e4d1a411305e7913d44ab56e388125ef597be46a8958b2ed8c5e298076c2d69ab3337c944f5356c462b"},{"type":"Signature","value":"d9ac57bac4260c60707e0b641585c70789e1a2eb5438c95de972af9aff99f5f4485b81cd2382218583b7f4950da54dbd8d1468f72b91809e14bb1c8139cca637"}]}}}
Data relay success, the hash is shown as follows:
0xdcf144d9ed2d64482fb5caafa719cf6706e9afd607ab043e8bfcb9018795e4d1
