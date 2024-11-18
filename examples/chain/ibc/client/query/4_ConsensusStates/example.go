package main

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/query"

	"sdk-go/client"

	chainclient "sdk-go/client/chain"
	"sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	"os"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.heliosd",
		"heliosd",
		"file",
		"inj-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	if err != nil {
		panic(err)
	}

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		panic(err)
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	clientId := "07-tendermint-0"
	pagination := query.PageRequest{Offset: 2, Limit: 4}
	ctx := context.Background()

	res, err := chainClient.FetchIBCConsensusStates(ctx, clientId, &pagination)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(res)

}
