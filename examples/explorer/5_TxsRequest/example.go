package main

import (
	"context"
	"encoding/json"
	"fmt"

	explorerPB "sdk-go/exchange/explorer_rpc/pb"

	"sdk-go/client/common"
	explorerclient "sdk-go/client/explorer"
)

func main() {
	network := common.LoadNetwork("testnet", "lb")
	explorerClient, err := explorerclient.NewExplorerClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	before := uint64(7158400)

	req := explorerPB.GetTxsRequest{
		Before: before,
	}

	res, err := explorerClient.GetTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
	}

	str, _ := json.MarshalIndent(res, "", " ")
	fmt.Print(string(str))
}
