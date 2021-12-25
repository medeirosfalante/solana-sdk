package solanasdk_test

import (
	"context"
	"fmt"
	"testing"

	solanasdk "github.com/medeirosfalante/solana-sdk"
	"github.com/medeirosfalante/solana-sdk/ws"
)

func TestGetByBlockWs(t *testing.T) {
	endpoint := solanasdk.TestNet_WS
	client, err := ws.Connect(context.Background(), endpoint)
	if err != nil {
		panic(err)
	}
	t.Errorf("inital")
	sub, err := client.SlotSubscribe()
	if err != nil {
		panic(err)
	}

	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Printf("got  #%v\n", got)
	}

}
