package main

import (
	"encoding/hex"
	"io/ioutil"

	"github.com/meitu/go-ethereum/accounts/keystore"
	"github.com/meitu/go-ethereum/crypto"
)

func main() {
	keyjson, err := ioutil.ReadFile("keystore/account1.json")
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
