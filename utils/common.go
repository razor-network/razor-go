package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"time"
)

func ConnectToClient(provider string) *ethclient.Client {
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	return client
}

func FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	address := common.HexToAddress(accountAddress)
	coinContract := GetCoinContract(client)
	opts := GetOptions(false, accountAddress, "")
	return coinContract.BalanceOf(&opts, address)
}

func GetDefaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	defaultPath := home + "/.razor"
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		os.Mkdir(defaultPath, 0777)
	}
	return defaultPath
}

func GetEpoch(client *ethclient.Client, address string) (*big.Int, error) {
	stateManager := GetStateManager(client)
	callOpts := GetOptions(false, address, "")
	return stateManager.GetEpoch(&callOpts)
}

func GetStakerId(client *ethclient.Client, address string) (*big.Int, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	return stakeManager.GetStakerId(&callOpts, common.HexToAddress(address))
}

func GetDelayedState(client *ethclient.Client) (int64, error) {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return -1, err
	}
	// TODO: Check error message
	if blockNumber%(core.BlockDivider) > 7 || blockNumber%(core.BlockDivider) < 1 {
		return -1, nil
	}
	state := math.Floor(float64(blockNumber / core.BlockDivider))
	return int64(state) % core.NumberOfStates, nil
}

func GetStake(client *ethclient.Client, address string, stakerId *big.Int) (*big.Int, error) {
	stakeManager := GetStakeManager(client)
	callOpts := GetOptions(false, address, "")
	stake, err := stakeManager.Stakers(&callOpts, stakerId)
	if err != nil {
		return nil, err
	}
	return stake.Stake, nil
}

func GetMinStakeAmount(client *ethclient.Client, address string) (*big.Int, error) {
	constantsManager := GetConstantsManager(client)
	callOpts := GetOptions(false, address, "")
	return constantsManager.MinStake(&callOpts)
}

func GetActiveJobs(client *ethclient.Client, address string) ([]types.Job, error) {
	var jobs []types.Job
	jobManager := GetJobManager(client)
	callOpts := GetOptions(false, address, "")
	numOfJobs, err := jobManager.GetNumJobs(&callOpts)
	if err != nil {
		return jobs, err
	}
	epoch, err := GetEpoch(client, address)
	if err != nil {
		return jobs, err
	}
	for jobIndex := 0; jobIndex < int(numOfJobs.Int64()); jobIndex++ {
		callOpts = GetOptions(false, address, "")
		job, err := jobManager.Jobs(&callOpts, big.NewInt(int64(jobIndex)))
		if err != nil {
			log.Error("Error in fetching job", err)
		} else {
			if !job.Fulfilled && job.Epoch.Cmp(epoch) < 0 {
				jobs = append(jobs, job)
			}
		}
	}
	return jobs, nil
}

func checkTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

func WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	for {
		log.Info("Checking if transaction is mined....\n")
		transactionStatus := checkTransactionReceipt(client, hashToRead)
		if transactionStatus == 0 {
			log.Info("Transaction mining unsuccessful")
			return 0
		} else if transactionStatus == 1 {
			log.Info("Transaction mined successfully\n")
			return 1
		}
		time.Sleep(2 * time.Second)
	}
}