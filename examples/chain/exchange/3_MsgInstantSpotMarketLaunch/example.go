package main

import (
	"encoding/json"
	"fmt"
	"os"

	"cosmossdk.io/math"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"

	exchangetypes "sdk-go/chain/exchange/types"
	"sdk-go/client"
	chainclient "sdk-go/client/chain"
	"sdk-go/client/common"
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

	minPriceTickSize := math.LegacyMustNewDecFromStr("0.01")
	minQuantityTickSize := math.LegacyMustNewDecFromStr("0.001")
	minNotional := math.LegacyMustNewDecFromStr("1")

	chainMinPriceTickSize := minPriceTickSize.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(6)))
	chainMinPriceTickSize = chainMinPriceTickSize.Quo(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(18)))

	chainMinQuantityTickSize := minQuantityTickSize.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(18)))
	chainMinNotional := minNotional.Mul(math.LegacyNewDecFromIntWithPrec(math.NewInt(1), int64(6)))

	msg := &exchangetypes.MsgInstantSpotMarketLaunch{
		Sender:              senderAddress.String(),
		Ticker:              "INJ/USDC",
		BaseDenom:           "helios",
		QuoteDenom:          "factory/inj17vytdwqczqz72j65saukplrktd4gyfme5agf6c/usdc",
		MinPriceTickSize:    chainMinPriceTickSize,
		MinQuantityTickSize: chainMinQuantityTickSize,
		MinNotional:         chainMinNotional,
	}

	// AsyncBroadcastMsg, SyncBroadcastMsg, QueueBroadcastMsg
	response, err := chainClient.AsyncBroadcastMsg(msg)

	if err != nil {
		panic(err)
	}

	str, _ := json.MarshalIndent(response, "", " ")
	fmt.Print(string(str))
}
