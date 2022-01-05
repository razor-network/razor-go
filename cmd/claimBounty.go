package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/spf13/cobra"
)

var claimBountyCmd = &cobra.Command{
	Use:   "claimBounty",
	Short: "claim earned bounty",
	Long: `ClaimBounty allows the users who are bountyHunter to redeem their bounty in razor network

Example:
  ./razor claimBounty --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --bountyId 2`,
	Run: initialiseClaimBounty,
}

func initialiseClaimBounty(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteClaimBounty(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteClaimBounty(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	bountyId, err := flagSetUtilsMockery.GetUint32BountyId(flagSet)
	utils.CheckError("Error in getting bountyId: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	redeemBountyInput := types.RedeemBountyInput{
		Address:  address,
		Password: password,
		BountyId: bountyId,
	}

	txn, err := cmdUtilsMockery.ClaimBounty(config, client, redeemBountyInput)
	utils.CheckError("ClaimBounty error: ", err)

	if txn != core.NilHash {
		razorUtilsMockery.WaitForBlockCompletion(client, txn.String())
	}
}

func (*UtilsStructMockery) ClaimBounty(config types.Configurations, client *ethclient.Client, redeemBountyInput types.RedeemBountyInput) (common.Hash, error) {
	txnArgs := types.TransactionOptions{
		Client:          client,
		AccountAddress:  redeemBountyInput.Address,
		Password:        redeemBountyInput.Password,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerABI,
		MethodName:      "redeemBounty",
		Parameters:      []interface{}{redeemBountyInput.BountyId},
	}
	epoch, err := razorUtilsMockery.GetEpoch(txnArgs.Client)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return common.Hash{0x00}, err
	}

	callOpts := razorUtilsMockery.GetOptions()
	bountyLock, err := stakeManagerUtilsMockery.GetBountyLock(txnArgs.Client, &callOpts, redeemBountyInput.BountyId)
	if err != nil {
		log.Error("Error in getting bounty lock: ", err)
		return core.NilHash, err
	}

	if bountyLock.Amount.Cmp(big.NewInt(0)) == 0 {
		err = errors.New("bounty amount is 0")
		log.Error(err)
		return core.NilHash, err
	}

	log.Info("Claiming bounty transaction...")
	waitFor := bountyLock.RedeemAfter - epoch
	if waitFor > 0 {
		log.Debug("Waiting for lock period to get over....")

		//waiting till epoch reaches redeemAfter
		razorUtilsMockery.Sleep(time.Duration(int64(waitFor)*core.EpochLength*razorUtilsMockery.CalculateBlockTime(client)) * time.Second)
	}

	txnOpts := razorUtilsMockery.GetTxnOpts(txnArgs)

	for retry := 1; retry <= int(core.MaxRetries); retry++ {
		tx, err := stakeManagerUtilsMockery.RedeemBounty(txnArgs.Client, txnOpts, redeemBountyInput.BountyId)
		if err == nil {
			log.Info("Txn Hash: ", transactionUtilsMockery.Hash(tx).Hex())
			return transactionUtilsMockery.Hash(tx), nil
		}
		log.Error("Error while claiming bounty: ", err)
		if retry != int(core.MaxRetries) {
			log.Info("Retrying again...")
			log.Info("Waiting for 1 more epoch...")
			razorUtilsMockery.Sleep(time.Duration(core.EpochLength) * time.Second)
		}
	}
	return core.NilHash, err
}

func init() {
	razorUtilsMockery = &UtilsMockery{}
	cmdUtilsMockery = &UtilsStructMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}

	rootCmd.AddCommand(claimBountyCmd)
	var (
		Address  string
		Password string
		BountyId uint32
	)

	claimBountyCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	claimBountyCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	claimBountyCmd.Flags().Uint32VarP(&BountyId, "bountyId", "", 0, "bountyId of the bounty hunter")

	addrErr := claimBountyCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	bountyIdErr := claimBountyCmd.MarkFlagRequired("bountyId")
	utils.CheckError("BountyId error: ", bountyIdErr)

}
