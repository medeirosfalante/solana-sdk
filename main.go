package solanasdk

import (
	"context"
	"fmt"

	"github.com/medeirosfalante/solana-sdk/ws"
)

func main() {
	endpoint := TestNet_WS
	client, err := ws.Connect(context.Background(), endpoint)
	if err != nil {
		panic(err)
	}
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
		fmt.Printf("got  #%v", got)
	}

}
