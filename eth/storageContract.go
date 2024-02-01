package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// this contract exists only on tenderly fork
const contractAddress = "0x790d896b22f98950333da325f2dc1356ab17bb33"
const contractABI = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "initMessage",
				"type": "string"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [],
		"name": "message",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "newMessage",
				"type": "string"
			}
		],
		"name": "update",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

func GetContractMessage(client *ethclient.Client) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", errors.Wrap(err, "failed to parse contract ABI")
	}

	contractAddress := common.HexToAddress(contractAddress)

	getMessage, err := parsedABI.Pack("message")
	if err != nil {
		return "", errors.Wrap(err, "failed to pack 'getMessage' func for call")
	}

	msg := ethereum.CallMsg{To: &contractAddress, Data: getMessage}
	output, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to call contract")
	}

	// Dekodiranje izlaza
	var message string
	err = parsedABI.UnpackIntoInterface(&message, "message", output)
	if err != nil {
		return "", errors.Wrap(err, "failed to unpack output")
	}

	return message, nil
}

func SetContractMessage(client *ethclient.Client, message string) error {
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return errors.Wrap(err, "failed to convert private key to ECDSA")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return errors.Wrap(err, "failed to get nonce")
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get network ID")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return errors.Wrap(err, "failed to create transactor")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get gas price")
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)

	contractABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return errors.Wrap(err, "failed to parse contract ABI")
	}

	data, err := contractABI.Pack("update", message)
	if err != nil {
		return errors.Wrap(err, "failed to pack data for tx")
	}

	contractAddress := common.HexToAddress(contractAddress)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: auth.GasPrice,
		Gas:      auth.GasLimit,
		To:       &contractAddress,
		Value:    big.NewInt(0),
		Data:     data,
	})

	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return errors.Wrap(err, "failed to sign tx")
	}

	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		return errors.Wrap(err, "failed to send tx")
	}

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())

	return nil
}
