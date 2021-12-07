package cmd

import (
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var razorUtils utilsInterface
var tokenManagerUtils tokenManagerInterface
var transactionUtils transactionInterface
var stakeManagerUtils stakeManagerInterface

//var utilsStructInterface structUtilsInterface

var stakeCmd = &cobra.Command{
	Use:   "stake",
	Short: "Stake some razors",
	Long: `Stake allows user to stake razors in the razor network

Example:
  ./razor stake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000`,
	Run: func(cmd *cobra.Command, args []string) {
		utilsStruct := UtilsStruct{
			razorUtils:        razorUtils,
			stakeManagerUtils: stakeManagerUtils,
			transactionUtils:  transactionUtils,
			tokenManagerUtils: tokenManagerUtils,
			flagSetUtils:      flagSetUtils,
		}
		password := utils.AssignPassword(cmd.Flags())
		utilsStruct.executeStake(cmd.Flags(), password)
	},
}

func (utilsStruct UtilsStruct) executeStake(flagSet *pflag.FlagSet, password string) {
	config, err := GetConfigData(utilsStruct)
	utils.CheckError("Error in getting config: ", err)

	address, _ := flagSet.GetString("address")
	client := utils.ConnectToClient(config.Provider)
	balance, err := utils.FetchBalance(client, address)
	utils.CheckError("Error in fetching balance for account: "+address, err)

	valueInWei, err := AssignAmountInWei(flagSet, utilsStruct)
	utils.CheckError("Error in getting amount: ", err)

	utils.CheckAmountAndBalance(valueInWei, balance)

	utils.CheckEthBalanceIsZero(client, address)

	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Password:       password,
		Amount:         valueInWei,
		ChainId:        core.ChainId,
		Config:         config,
	}

	approveTxnHash, err := utilsStruct.approve(txnArgs)
	utils.CheckError("Approve error: ", err)

	if approveTxnHash != core.NilHash {
		razorUtils.WaitForBlockCompletion(txnArgs.Client, approveTxnHash.String())
	}

	stakeTxnHash, err := utilsStruct.stakeCoins(txnArgs)
	utils.CheckError("Stake error: ", err)
	razorUtils.WaitForBlockCompletion(txnArgs.Client, stakeTxnHash.String())
}

func (utilsStruct UtilsStruct) stakeCoins(txnArgs types.TransactionOptions) (common.Hash, error) {
	epoch, err := utilsStruct.razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		return common.Hash{0x00}, err
	}

	log.Info("Sending stake transactions...")
	txnArgs.ContractAddress = core.StakeManagerAddress
	txnArgs.MethodName = "stake"
	txnArgs.Parameters = []interface{}{epoch, txnArgs.Amount}
	txnArgs.ABI = bindings.StakeManagerABI
	txnOpts := utilsStruct.razorUtils.GetTxnOpts(txnArgs)
	tx, err := utilsStruct.stakeManagerUtils.Stake(txnArgs.Client, txnOpts, epoch, txnArgs.Amount)
	if err != nil {
		return common.Hash{0x00}, err
	}
	log.Info("Txn Hash: ", utilsStruct.transactionUtils.Hash(tx).Hex())
	return utilsStruct.transactionUtils.Hash(tx), nil
}

func init() {
	razorUtils = Utils{}
	tokenManagerUtils = TokenManagerUtils{}
	transactionUtils = TransactionUtils{}
	stakeManagerUtils = StakeManagerUtils{}
	flagSetUtils = FlagSetUtils{}

	rootCmd.AddCommand(stakeCmd)
	var (
		Amount   string
		Address  string
		Password string
		Power    string
	)

	stakeCmd.Flags().StringVarP(&Amount, "value", "v", "0", "amount of Razors to stake")
	stakeCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	stakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	stakeCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")

	amountErr := stakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", amountErr)
	addrErr := stakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
