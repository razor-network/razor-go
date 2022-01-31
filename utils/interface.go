package utils

import (
	"context"
	"crypto/ecdsa"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"razor/accounts"
	"razor/core/types"
	"razor/pkg/bindings"
)

//go:generate mockery --name OptionUtils --output ./mocks/ --case=underscore
//go:generate mockery --name Utils --output ./mocks/ --case=underscore

var Options OptionUtils
var UtilsInterface Utils

type OptionUtils interface {
	Parse(io.Reader) (abi.ABI, error)
	Pack(abi.ABI, string, ...interface{}) ([]byte, error)
	GetDefaultPath() (string, error)
	GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey
	NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error)
	RetryAttempts(uint) retry.Option
	PendingNonceAt(*ethclient.Client, context.Context, common.Address) (uint64, error)
	HeaderByNumber(*ethclient.Client, context.Context, *big.Int) (*Types.Header, error)
	SuggestGasPrice(*ethclient.Client, context.Context) (*big.Int, error)
	EstimateGas(*ethclient.Client, context.Context, ethereum.CallMsg) (uint64, error)
	FilterLogs(*ethclient.Client, context.Context, ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAt(*ethclient.Client, context.Context, common.Address, *big.Int) (*big.Int, error)
	GetNumProposedBlocks(*ethclient.Client, *bind.CallOpts, uint32) (uint8, error)
	GetProposedBlock(*ethclient.Client, *bind.CallOpts, uint32, uint32) (bindings.StructsBlock, error)
	GetBlock(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsBlock, error)
	MinStake(*ethclient.Client, *bind.CallOpts) (*big.Int, error)
	MaxAltBlocks(*ethclient.Client, *bind.CallOpts) (uint8, error)
	SortedProposedBlockIds(*ethclient.Client, *bind.CallOpts, uint32, *big.Int) (uint32, error)
	Commitments(*ethclient.Client, *bind.CallOpts, uint32) (types.Commitment, error)
	GetVoteValue(*ethclient.Client, *bind.CallOpts, uint16, uint32) (*big.Int, error)
	GetVote(*ethclient.Client, *bind.CallOpts, uint32) (bindings.StructsVote, error)
	GetInfluenceSnapshot(*ethclient.Client, *bind.CallOpts, uint32, uint32) (*big.Int, error)
	GetStakeSnapshot(*ethclient.Client, *bind.CallOpts, uint32, uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(*ethclient.Client, *bind.CallOpts, uint32) (*big.Int, error)
	GetRandaoHash(*ethclient.Client, *bind.CallOpts) ([32]byte, error)
	GetEpochLastCommitted(*ethclient.Client, *bind.CallOpts, uint32) (uint32, error)
	GetEpochLastRevealed(*ethclient.Client, *bind.CallOpts, uint32) (uint32, error)
}

type Utils interface {
	SuggestGasPriceWithRetry(*ethclient.Client) (*big.Int, error)
	MultiplyFloatAndBigInt(*big.Int, float64) *big.Int
	GetPendingNonceAtWithRetry(*ethclient.Client, common.Address) (uint64, error)
	GetGasPrice(*ethclient.Client, types.Configurations) *big.Int
	GetTxnOpts(types.TransactionOptions) *bind.TransactOpts
	GetGasLimit(types.TransactionOptions, *bind.TransactOpts) (uint64, error)
	EstimateGasWithRetry(*ethclient.Client, ethereum.CallMsg) (uint64, error)
	IncreaseGasLimitValue(*ethclient.Client, uint64, float32) (uint64, error)
	GetLatestBlockWithRetry(*ethclient.Client) (*Types.Header, error)
	FilterLogsWithRetry(*ethclient.Client, ethereum.FilterQuery) ([]Types.Log, error)
	BalanceAtWithRetry(*ethclient.Client, common.Address) (*big.Int, error)
	GetBlockManager(*ethclient.Client) *bindings.BlockManager
	GetOptions() bind.CallOpts
	GetNumberOfProposedBlocks(*ethclient.Client, uint32) (uint8, error)
	GetSortedProposedBlockId(*ethclient.Client, uint32, *big.Int) (uint32, error)
	FetchPreviousValue(*ethclient.Client, uint32, uint16) (uint32, error)
	GetMaxAltBlocks(*ethclient.Client) (uint8, error)
	GetMinStakeAmount(*ethclient.Client) (*big.Int, error)
	GetProposedBlock(*ethclient.Client, uint32, uint32) (bindings.StructsBlock, error)
	GetSortedProposedBlockIds(*ethclient.Client, uint32) ([]uint32, error)
	GetBlockManagerWithOpts(*ethclient.Client) (*bindings.BlockManager, bind.CallOpts)
	GetVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts)
	GetCommitments(client *ethclient.Client, address string) ([32]byte, error)
	GetVoteValue(client *ethclient.Client, assetId uint16, stakerId uint32) (*big.Int, error)
	GetVotes(client *ethclient.Client, stakerId uint32) (bindings.StructsVote, error)
	GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error)
	GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error)
	GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32) (*big.Int, error)
	GetRandaoHash(client *ethclient.Client) ([32]byte, error)
	GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error)
	GetVoteManager(client *ethclient.Client) *bindings.VoteManager
}

type OptionsStruct struct{}
type UtilsStruct struct{}

type OptionsPackageStruct struct {
	Options        OptionUtils
	UtilsInterface Utils
}
