package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	todo "github.com/idirall22/go-solidity-learn/gen"
)

func main() {
	b, err := ioutil.ReadFile("wallet/UTC--2021-05-24T16-47-26.459903259Z--c393967d7b4b7fd02e697d13085d645c9412af11")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://kovan.infura.io/v3/188b9ff61e6f46dc8c66b6ed36487bbe")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress("0x48B63160d3Bbc7B566aa596768a19523278C1525")
	t, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice

	// tra, err := t.Add(tx, "First Task")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tra.Hash())

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := t.List(&bind.CallOpts{
		From: add,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)

	// tra, err := t.Update(tx, big.NewInt(0), "update task content")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Toggle tx", tra.Hash())

	// tra, err := t.Remove(tx, big.NewInt(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Toggle tx", tra.Hash())

}
