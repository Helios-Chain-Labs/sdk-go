package main

import (
	"context"
	"encoding/json"
	"fmt"

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
		os.Getenv("HOME")+"/.d",
		"d",
		"file",
		"helios-user",
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

	address := "helios14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	denom := "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5"
	ctx := context.Background()

	res, err := chainClient.GetBankSpendableBalancesByDenom(ctx, address, denom)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))

}
