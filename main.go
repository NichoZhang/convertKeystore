package main

import (
	"encoding/hex"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

const filePath = "keystore/account1.json"

func main() {
	keyjson, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	password := ""

	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		panic(err)
	}
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	println(privateKey)
}
