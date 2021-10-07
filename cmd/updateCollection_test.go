package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"testing"
)

func Test_updateCollection(t *testing.T) {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	txnOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))

	var client *ethclient.Client
	var flagSet *pflag.FlagSet
	var config types.Configurations

	razorUtils := UtilsMock{}
	assetManagerUtils := AssetManagerMock{}
	transactionUtils := TransactionMock{}
	flagSetUtils := FlagSetMock{}

	type args struct {
		password            string
		collectionId        uint8
		collectionIdErr     error
		address             string
		addressErr          error
		aggregation         uint32
		aggregationErr      error
		power               int8
		powerErr            error
		txnOpts             *bind.TransactOpts
		updateCollectionTxn *Types.Transaction
		updateCollectionErr error
		hash                common.Hash
	}

	tests := []struct {
		name    string
		args    args
		want    common.Hash
		wantErr error
	}{
		{
			name: "Test 1: When updateCollection function executes successfully",
			args: args{
				password:            "test",
				collectionId:        3,
				collectionIdErr:     nil,
				address:             "0x000000000000000000000000000000000000dead",
				addressErr:          nil,
				aggregation:         1,
				aggregationErr:      nil,
				power:               0,
				powerErr:            nil,
				txnOpts:             txnOpts,
				updateCollectionTxn: &Types.Transaction{},
				updateCollectionErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    common.BigToHash(big.NewInt(1)),
			wantErr: nil,
		},
		{
			name: "Test 2: When there is an error in getting collection id from flags",
			args: args{
				password:            "test",
				collectionIdErr:     errors.New("collectionIdErr error"),
				address:             "0x000000000000000000000000000000000000dead",
				addressErr:          nil,
				aggregation:         1,
				aggregationErr:      nil,
				power:               0,
				powerErr:            nil,
				txnOpts:             txnOpts,
				updateCollectionTxn: &Types.Transaction{},
				updateCollectionErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("collectionIdErr error"),
		},
		{
			name: "Test 3: When there is an error in getting address from flags",
			args: args{
				password:            "test",
				collectionId:        3,
				collectionIdErr:     nil,
				address:             "",
				addressErr:          errors.New("address error"),
				aggregation:         1,
				aggregationErr:      nil,
				power:               0,
				powerErr:            nil,
				txnOpts:             txnOpts,
				updateCollectionTxn: &Types.Transaction{},
				updateCollectionErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("address error"),
		},
		{
			name: "Test 4: When there is an error in getting aggregation method from flags",
			args: args{
				password:            "test",
				collectionId:        3,
				collectionIdErr:     nil,
				address:             "0x000000000000000000000000000000000000dead",
				addressErr:          nil,
				aggregationErr:      errors.New("aggregation error"),
				power:               0,
				powerErr:            nil,
				txnOpts:             txnOpts,
				updateCollectionTxn: &Types.Transaction{},
				updateCollectionErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("aggregation error"),
		},
		{
			name: "Test 5: When there is an error in getting power from flags",
			args: args{
				password:            "test",
				collectionId:        3,
				collectionIdErr:     nil,
				address:             "0x000000000000000000000000000000000000dead",
				addressErr:          nil,
				aggregation:         1,
				aggregationErr:      nil,
				powerErr:            errors.New("power error"),
				txnOpts:             txnOpts,
				updateCollectionTxn: &Types.Transaction{},
				updateCollectionErr: nil,
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("power error"),
		},
		{
			name: "Test 6: When updateCollection transaction fails",
			args: args{
				password:            "test",
				collectionId:        3,
				collectionIdErr:     nil,
				address:             "0x000000000000000000000000000000000000dead",
				addressErr:          nil,
				aggregation:         1,
				aggregationErr:      nil,
				power:               0,
				powerErr:            nil,
				txnOpts:             txnOpts,
				updateCollectionTxn: &Types.Transaction{},
				updateCollectionErr: errors.New("updateCollection error"),
				hash:                common.BigToHash(big.NewInt(1)),
			},
			want:    core.NilHash,
			wantErr: errors.New("updateCollection error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssignPasswordMock = func(*pflag.FlagSet) string {
				return tt.args.password
			}

			GetUint8CollectionIdMock = func(*pflag.FlagSet) (uint8, error) {
				return tt.args.collectionId, tt.args.collectionIdErr
			}

			GetStringAddressMock = func(*pflag.FlagSet) (string, error) {
				return tt.args.address, tt.args.addressErr
			}

			GetUint32AggregationMock = func(*pflag.FlagSet) (uint32, error) {
				return tt.args.aggregation, tt.args.aggregationErr
			}

			GetInt8PowerMock = func(*pflag.FlagSet) (int8, error) {
				return tt.args.power, tt.args.powerErr
			}

			ConnectToClientMock = func(string) *ethclient.Client {
				return client
			}

			GetTxnOptsMock = func(types.TransactionOptions) *bind.TransactOpts {
				return tt.args.txnOpts
			}

			UpdateCollectionMock = func(*ethclient.Client, *bind.TransactOpts, uint8, uint32, int8) (*Types.Transaction, error) {
				return tt.args.updateCollectionTxn, tt.args.updateCollectionErr
			}

			HashMock = func(*Types.Transaction) common.Hash {
				return tt.args.hash
			}

			got, err := updateCollection(flagSet, config, razorUtils, assetManagerUtils, transactionUtils, flagSetUtils)
			if got != tt.want {
				t.Errorf("Txn hash for updateCollection function, got = %v, want %v", got, tt.want)
			}
			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for updateCollection function, got = %v, want %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for updateCollection function, got = %v, want %v", got, tt.wantErr)
				}
			}
		})
	}
}
