package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"math/big"
	"strings"
	"testing"
)

type ExternalApiMock struct{}

var (
// TODO: resolve internal api mock problem
//_ core.ExternalAPI = ExternalApiMock{}
)

func (externalApiMock ExternalApiMock) New(ctx context.Context) (address common.Address, err error) {
	return
}
func (externalApiMock ExternalApiMock) List(ctx context.Context) (addresses []common.Address, err error) {
	return
}
func (externalApiMock ExternalApiMock) SignTransaction(ctx context.Context, args core.SendTxArgs, methodSelector *string) (result interface{}, err error) {
	return
}
func (externalApiMock ExternalApiMock) SignData(ctx context.Context) (addresses []common.Address, err error) {
	return
}
func (externalApiMock ExternalApiMock) SignTypedData(ctx context.Context) (addresses []common.Address, err error) {
	return
}
func (externalApiMock ExternalApiMock) EcRecover(ctx context.Context) (addresses []common.Address, err error) {
	return
}
func (externalApiMock ExternalApiMock) Version(ctx context.Context) (addresses []common.Address, err error) {
	return
}

func TestPrepareTransactionsForPool(t *testing.T) {
	// it must be real endpoint, IPC is misleading because it does not need to be ipc.
	ipcLocation := "http://34.91.133.193:8545"

	// Star client on a server
	client := newClient(ipcLocation)

	privateKey, err := crypto.HexToECDSA(strings.ToLower(DefaultPrivateKey))
	assert.Nil(t, err)

	t.Run("Prepare 50 transactions", func(t *testing.T) {
		expectedLen := 50
		transactionsLen := big.NewInt(int64(expectedLen))
		err, transactions := PrepareTransactionsForPool(transactionsLen, client, privateKey)
		assert.Nil(t, err)
		assert.NotEmpty(t, transactions)
		assert.Len(t, transactions, expectedLen)
	})
}

// I leave it here with error: `method eth_syncing` does not exist. I do not want to waste time now for mocking it.
// Possible solution: add another mock for api for eth and assign method "eth_syncing"
// and other methods that are missing
func possibleMockForEthIPC(t *testing.T, ipcLocation string) {
	myApi := ExternalApiMock{}

	rpcAPI := []rpc.API{
		{
			Namespace: "account",
			Public:    true,
			Service:   myApi,
			Version:   "1.0",
		},
		{
			Namespace: "eth",
			Public:    true,
			Service:   myApi,
			Version:   "1.0",
		},
	}

	listener, server, err := rpc.StartIPCEndpoint(ipcLocation, rpcAPI)
	assert.Nil(t, err)

	defer func() {
		server.Stop()
		_ = listener.Close()
	}()
}
