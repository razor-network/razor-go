package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"math"
	"math/big"
	"time"
)

var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some razors",
	Long: `Stake allows user to stake razors in the razor network
	For ex:
	stake -a <amount> --address <address> --password <password>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, gasMultiplier, err := getConfigData(cmd)
		if err != nil {
			log.Fatal(err)
		}

		client := connectToClient(provider)
		address, _ := cmd.Flags().GetString("address")
		balance := fetchBalance(client, address)

		amount, err := cmd.Flags().GetString("amount")
		if err != nil  {
			log.Fatal("Error in reading amount", err)
		}
		amountInWei, ok := new(big.Int).SetString(amount, 10)
		if !ok {
			log.Fatal("SetString: error")
		}
		if amountInWei.Cmp(balance) > 0 {
			log.Fatal("Not enough balance")
		}

		accountBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
		if err != nil {
			log.Fatalf("Error in fetching balance of the account: %s\n%s", address, err)
		}

		minEtherBalance := new(big.Int).SetInt64(1e16)
		if accountBalance.Cmp(minEtherBalance) < 0 {
			log.Fatal("Please make sure you hold at least 0.01 ether in your account")
		}
		// TODO: Send stake manager address in 'to'
		approve(client, address, "0x1C7Ccf3054bA60bA8Ec1fecC7E4E722b59bDD90b", amountInWei, gasMultiplier)
		stakeCoins(client, address, amountInWei, gasMultiplier)
		//password, _ := cmd.Flags().GetString("password")
	},
}

func approve(client *ethclient.Client, from string, to string, amount *big.Int, gasMultiplier float32)  {
	coinContract := getCoinContract(client)
	opts := getOptions(false, from, "")
	allowance, err := coinContract.Allowance(&opts, common.HexToAddress(from), common.HexToAddress(to))
	if err != nil {
		log.Fatal(err)
	}
	if allowance.Cmp(amount) >= 0 {
		log.Info("Sufficient allowance, no need to increase")
	} else {
		log.Info("Sending Approve transaction...")
		txnOpts := getTxnOpts(client, from, gasMultiplier)
		txn, err := coinContract.Approve(&txnOpts, common.HexToAddress(to), amount)
		if err != nil {
			log.Fatal("Error in approving: ",err)
		}
		log.Info("Approve transaction sent...\nTxn Hash: ", txn.Hash())
	}
}

func stakeCoins(client *ethclient.Client, account string, amount *big.Int, gasMultiplier float32) {
	stateManager := getStateManager(client)
	stakeManager := getStakeManager(client)
	// TODO: Get a better approach for assigning epoch
	var epoch *big.Int
	for true {
		callOpts := getOptions(false, account, "")
		_epoch, err := stateManager.GetEpoch(&callOpts)
		if err != nil {
			log.Fatal("Error in fetching epoch: ", err)
		}
		epoch = _epoch
		state := getDelayedState(client)
		log.Info("Epoch", epoch)
		log.Info("State", state)
		if state != 0 {
			log.Info("Can only stake during state 0 (commit). Retrying in 1 second...")
			time.Sleep(1000)
		} else {
			break
		}
	}
	log.Info("Sending stake transactions...")
	txnOpts := getTxnOpts(client, account, gasMultiplier)
	tx, err := stakeManager.Stake(&txnOpts, epoch, amount)
	if err != nil {
		log.Fatal("Error in staking: ", err)
	}
	log.Info("Staked\nTxn Hash: ", tx.Hash())
}

func getDelayedState(client *ethclient.Client) int64 {
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Error in fetching latest block number: ", err)
	}
	if blockNumber % 10 > 7 || blockNumber % 10 < 1 {
		return -1
	}
	state := math.Floor(float64(blockNumber/10))
	return int64(state) % 4
}

func init() {
	rootCmd.AddCommand(stakeCmd)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	var (
		Amount   string
		Address  string
		Password string
	)

	stakeCmd.Flags().StringVarP(&Amount, "amount", "a", "0", "amount to stake (in Wei)")
	stakeCmd.Flags().StringVarP(&Address, "address", "", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password to unlock account")

	stakeCmd.MarkFlagRequired("amount")
	stakeCmd.MarkFlagRequired("address")
	stakeCmd.MarkFlagRequired("password")

}
