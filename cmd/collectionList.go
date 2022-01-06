package cmd

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"razor/utils"
	"strconv"
	"strings"
)

// collectionListCmd represents the collectionList command
var collectionListCmd = &cobra.Command{
	Use:   "collectionList",
	Short: "list of all collections",
	Long: `Provides the list of all collections with their name, power, ID etc.
Example:
	./razor collectionList `,
	Run: func(cmd *cobra.Command, args []string) {

		utilsStruct := UtilsStruct{
			razorUtils:   razorUtils,
			flagSetUtils: flagSetUtils,
		}

		config, err := cmdUtilsMockery.GetConfigData()
		utils.CheckError("Error in getting config: ", err)

		client := utils.ConnectToClient(config.Provider)

		err = utilsStruct.GetCollectionList(client)

		if err != nil {
			log.Error("Error in getting collection list: ", err)
		}
	},
}

func (utilsStruct *UtilsStruct) GetCollectionList(client *ethclient.Client) error {
	collections, err := utilsStruct.razorUtils.GetCollections(client)

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Active", "Collection Id", "Asset Index", "Power", "Aggregation Method", "Job IDs", "Name"})
	for i := 0; i < len(collections); i++ {
		jobIDs, _ := json.Marshal(collections[i].JobIDs)

		table.Append([]string{
			strconv.FormatBool(collections[i].Active),
			strconv.Itoa(int(collections[i].Id)),
			strconv.Itoa(int(collections[i].AssetIndex)),
			strconv.Itoa(int(collections[i].Power)),
			strconv.Itoa(int(collections[i].AggregationMethod)),
			strings.Trim(string(jobIDs), "[]"),
			collections[i].Name,
		})

	}

	table.Render()
	return nil

}

func init() {
	razorUtils = Utils{}
	flagSetUtils = FlagSetUtils{}
	cmdUtilsMockery = &UtilsStructMockery{}

	rootCmd.AddCommand(collectionListCmd)

}