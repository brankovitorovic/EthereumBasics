package main

import (
	"ethereumBasics/eth"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	rpc                   = "<rpc>"
	sendAccountAddress    = "<address>"
	receiveAccountAddress = "<address>"
	// wei to eth decimal
	weiToEth = 1000000000000000000
)

func main() {

	client, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client, err: ", err)
		panic(err)
	}

	//number, err := client.BlockByNumber(context.Background(), nil)
	//if err != nil {
	//	fmt.Println("Failed to get block number, err: ", err)
	//	panic(err)
	//}
	//
	//fmt.Println("Block number: ", number.Number().Uint64())
	//
	//senderBalance, err := eth.GetBalance(client, sendAccountAddress)
	//if err != nil {
	//	fmt.Println("Failed to get balance for address ", sendAccountAddress, ", err: ", err)
	//	panic(err)
	//}
	//
	//receiverBalance, err := eth.GetBalance(client, receiveAccountAddress)
	//if err != nil {
	//	fmt.Println("Failed to get balance for address ", receiveAccountAddress, ", err: ", err)
	//	panic(err)
	//}

	//fmt.Println("Sender balance: ", senderBalance)
	//fmt.Println("Receiver balance: ", receiverBalance)
	//eth.PrintBalanceInEth(senderBalance)
	//eth.PrintBalanceInEth(receiverBalance)
	//
	//err = eth.SendOneEth(client, receiveAccountAddress)
	//if err != nil {
	//	fmt.Println("Failed to send 1 ETH, err: ", err)
	//	panic(err)
	//}
	//
	//eth.PrintBalanceInEth(senderBalance)
	//eth.PrintBalanceInEth(receiverBalance)

	message, err := eth.GetContractMessage(client)
	if err != nil {
		fmt.Println("Failed to get contract message, err: ", err)
		panic(err)
	}

	fmt.Println("Contract message: ", message)

	err = eth.SetContractMessage(client, "test message")
	if err != nil {
		fmt.Println("Failed to set contract message, err: ", err)
		panic(err)
	}

	message, err = eth.GetContractMessage(client)
	if err != nil {
		fmt.Println("Failed to get contract message, err: ", err)
		panic(err)
	}

	fmt.Println("Contract message: ", message)
}
