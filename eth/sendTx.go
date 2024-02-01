package eth

import (
	"context"
	"crypto/ecdsa"

	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const privateKey = "<private key>"

func SendOneEth(client *ethclient.Client, receiveAddress string) error {
	privateKeyHashed, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return errors.Wrap(err, "failed to convert private key to ECDSA")
	}

	publicKey := privateKeyHashed.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}

	//fmt.Println("From address: ", fromAddress.Hex())

	value := big.NewInt(1000000000000000000) // 1 ETH
	gasLimit := uint64(21000)                // gas limit for standard ETH transfer
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get gas price")
	}

	toAddress := common.HexToAddress(receiveAddress)
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	})

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKeyHashed)
	if err != nil {
		return errors.Wrap(err, "failed to sign tx")
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return errors.Wrap(err, "failed to send tx")
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	return nil
}
