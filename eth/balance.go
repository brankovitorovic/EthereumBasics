package eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(client *ethclient.Client, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		fmt.Println("Failed to get balance, err: ", err)
		return nil, err
	}

	return balance, nil
}

func PrintBalanceInEth(balance *big.Int) {
	weiToEth := new(big.Float)
	weiToEth.SetString("1e18")

	balanceFloat := new(big.Float).SetInt(balance)
	balanceInEth := new(big.Float).Quo(balanceFloat, weiToEth)
	fmt.Println("Balance in ETH: ", balanceInEth)
}
