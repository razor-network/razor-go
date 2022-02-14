package utils

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"razor/core"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

func (*UtilsStruct) ConnectToClient(provider string) *ethclient.Client {
	client, err := Options.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

func (*UtilsStruct) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	address := common.HexToAddress(accountAddress)
	coinContract := UtilsInterface.GetTokenManager(client)
	opts := UtilsInterface.GetOptions()
	return UtilsInterface.BalanceOf(coinContract, &opts, address)
}

func (*UtilsStruct) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	block, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		return -1, err
	}
	blockNumber := uint64(block.Number.Int64())
	lowerLimit := (core.StateLength * uint64(buffer)) / 100
	upperLimit := core.StateLength - (core.StateLength*uint64(buffer))/100
	if blockNumber%(core.StateLength) > upperLimit || blockNumber%(core.StateLength) < lowerLimit {
		return -1, nil
	}
	state := blockNumber / core.StateLength
	return int64(state) % core.NumberOfStates, nil
}

func (*UtilsStruct) CheckTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := Options.TransactionReceipt(client, context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

func (*UtilsStruct) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	timeout := core.BlockCompletionTimeout
	for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
		log.Debug("Checking if transaction is mined....")
		transactionStatus := UtilsInterface.CheckTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			log.Error("Transaction mining unsuccessful")
			return 0
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully")
			return 1
		}
		UtilsInterface.Sleep(3 * time.Second)
	}
	log.Info("Timeout Passed")
	return 0
}

func (*UtilsStruct) WaitTillNextNSecs(waitTime int32) {
	if waitTime <= 0 {
		waitTime = 1
	}
	UtilsInterface.Sleep(time.Duration(waitTime) * time.Second)
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg + err.Error())
	}
}

func (*UtilsStruct) IsFlagPassed(name string) bool {
	found := false
	for _, arg := range os.Args {
		if arg == "--"+name {
			found = true
		}
	}
	return found
}

func (*UtilsStruct) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	ethBalance, err := Options.BalanceAt(client, context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Fatalf("Error in fetching eth balance of the account: %s\n%s", address, err)
	}
	if ethBalance.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("Eth balance is 0, Aborting...")
	}
}

func (*UtilsStruct) GetStateName(stateNumber int64) string {
	var stateName string
	switch stateNumber {
	case 0:
		stateName = "Commit"
	case 1:
		stateName = "Reveal"
	case 2:
		stateName = "Propose"
	case 3:
		stateName = "Dispute"
	case 4:
		stateName = "Confirm"
	default:
		stateName = "-1"
	}
	return stateName
}

func (*UtilsStruct) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	if UtilsInterface.IsFlagPassed("stakerId") {
		return UtilsInterface.GetUint32(flagSet, "stakerId")
	}
	return UtilsInterface.GetStakerId(client, address)
}

func (*UtilsStruct) GetEpoch(client *ethclient.Client) (uint32, error) {
	latestHeader, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}
	epoch := latestHeader.Number.Int64() / core.EpochLength
	return uint32(epoch), nil
}

func (*UtilsStruct) SaveCommittedDataToFile(fileName string, epoch uint32, committedData []*big.Int) error {
	if len(committedData) == 0 {
		return errors.New("committed data is empty")
	}
	f, err := Options.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	fmt.Fprintln(f, epoch)
	for _, datum := range committedData {
		_, err := fmt.Fprintln(f, datum.String())
		if err != nil {
			return err
		}
	}
	defer f.Close()
	return nil
}

func (*UtilsStruct) ReadCommittedDataFromFile(fileName string) (uint32, []*big.Int, error) {
	var (
		committedData []*big.Int
		epoch         uint32
	)
	file, err := Options.Open(fileName)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	scanner := Options.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		if lineCount > 0 {
			data, ok := big.NewInt(0).SetString(scanner.Text(), 10)
			if ok {
				committedData = append(committedData, data)
			}
		} else {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return 0, nil, err
			}
			epoch = uint32(value)
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}
	return epoch, committedData, nil
}

func (*UtilsStruct) CalculateBlockTime(client *ethclient.Client) int64 {
	latestBlock, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Fatalf("Error in fetching latest Block: %s", err)
	}
	latestBlockNumber := latestBlock.Number
	lastSecondBlock, err := Options.HeaderByNumber(client, context.Background(), big.NewInt(1).Sub(latestBlockNumber, big.NewInt(1)))
	if err != nil {
		log.Fatalf("Error in fetching last second Block: %s", err)
	}
	return int64(latestBlock.Time - lastSecondBlock.Time)
}
