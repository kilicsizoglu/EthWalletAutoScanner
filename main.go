package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nanmu42/etherscan-api"
	"golang.org/x/crypto/sha3"
	"log"
)

func main() {

	for {
		key, err := crypto.GenerateKey()
		if err != nil {
			return
		}
		privateKeyBytes := crypto.FromECDSA(key)

		fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

		publicKey := key.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		fmt.Println("address : " + address)

		hash := sha3.NewLegacyKeccak256()
		hash.Write(publicKeyBytes[1:])
		fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

		client := etherscan.New(etherscan.Mainnet, "N9D77VUCTPSE98BDSJI1ZMWPK6IGYD41P9")

		balance, err := client.AccountBalance(address)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(balance.Int())

		if balance.Int().Int64() >= 1 {

			break

		}
	}

}
