package main

import (
	"encoding/json"
	"fmt"
	"os"

	"cosmossdk.io/math"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	"github.com/InjectiveLabs/sdk-go/client"
	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	tmClient, err := rpchttp.New(network.TmEndpoint, "/websocket")
	if err != nil {
		panic(err)
	}

	senderAddress, cosmosKeyring, err := chainclient.InitCosmosKeyring(
		os.Getenv("HOME")+"/.injectived",
		"injectived",
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

	minPriceTickSize := math.LegacyMustNewDecFromStr("0.01")
	minQuantityTickSize := math.LegacyMustNewDecFromStr("0.001")
	minNotional := math.LegacyMustNewDecFromStr("1")

	msg := &exchangev2types.MsgInstantSpotMarketLaunch{
		Sender:              senderAddress.String(),
		Ticker:              "INJ/USDC",
		BaseDenom:           "inj",
		QuoteDenom:          "factory/inj17vytdwqczqz72j65saukplrktd4gyfme5agf6c/usdc",
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
		MinNotional:         minNotional,
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
