package main

import (
	"fmt"

	"sdk-go/client"

	chainclient "sdk-go/client/chain"
	"sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("mainnet", "lb")

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}

	chainClient, err := chainclient.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	failEventCh := make(chan map[string]uint, 10000)
	go chainClient.StreamEventOrderFail("helios1rwv4zn3jptsqs7l8lpa3uvzhs57y8duemete9e", failEventCh)
	for {
		e := <-failEventCh
		fmt.Println(e)
	}
}
