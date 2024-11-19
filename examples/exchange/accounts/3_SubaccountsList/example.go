package main

import (
	"context"
	"encoding/json"
	"fmt"

	"sdk-go/client/common"
	exchangeclient "sdk-go/client/exchange"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	accountAddress := "helios14au322k9munkmx5wrchz9q30juf5wjgz2cfqku"
	res, err := exchangeClient.GetSubaccountsList(ctx, accountAddress)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
