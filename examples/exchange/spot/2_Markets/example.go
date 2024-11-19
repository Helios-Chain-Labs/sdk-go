package main

import (
	"context"
	"encoding/json"
	"fmt"

	"sdk-go/client/common"
	exchangeclient "sdk-go/client/exchange"
	spotExchangePB "sdk-go/exchange/spot_exchange_rpc/pb"
)

func main() {
	//network := common.LoadNetwork("mainnet", "k8s")
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketStatus := "active"
	quoteDenom := "peggy0xdAC17F958D2ee523a2206206994597C13D831ec7"

	req := spotExchangePB.MarketsRequest{
		MarketStatus: marketStatus,
		QuoteDenom:   quoteDenom,
	}

	res, err := exchangeClient.GetSpotMarkets(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
