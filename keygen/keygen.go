package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
)

const (
	ALPHABET = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	bigRadix = big.NewInt(58)
	bigZero  = big.NewInt(0)
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Generates a new address, public key and private key for the provided network key.")
		fmt.Println("Usage: keygen [netkey]")
		return
	}
	netkey, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	pub, priv, err := NewKey()
	if err != nil {
		log.Fatal(err)
	}

	addr := EncodeAddress(hash160(pub), byte(netkey))
	fmt.Println("addr:", addr)
	fmt.Println("pubk:", hex.EncodeToString(pub))
	fmt.Println("priv:", hex.EncodeToString(priv))
}

func NewKey() ([]byte, []byte, error) {
	priv, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubkey := btcec.PublicKey(priv.PublicKey)
	pubkeyaddr := &pubkey
	return pubkeyaddr.SerializeCompressed(), priv.D.Bytes(), nil
}

func EncodeAddress(hash160 []byte, key byte) string {
	tosum := make([]byte, 21)
	tosum[0] = key
	copy(tosum[1:], hash160)
	cksum := doubleHash(tosum)

	// Address before base58 encoding is 1 byte for netID, ripemd160 hash
	// size, plus 4 bytes of checksum (total 25).
	b := make([]byte, 25)
	b[0] = key
	copy(b[1:], hash160)
	copy(b[21:], cksum[:4])

	return base58Encode(b)
}

func hash160(data []byte) []byte {
	if len(data) == 1 && data[0] == 0 {
		data = []byte{}
	}
	h1 := sha256.Sum256(data)
	h2 := ripemd160.New()
	h2.Write(h1[:])
	return h2.Sum(nil)
}

func doubleHash(data []byte) []byte {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

// Base58Encode encodes a byte slice to a modified base58 string.
func base58Encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, ALPHABET[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, ALPHABET[0])
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
