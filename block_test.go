package solanasdk_test

import (
	"context"
	"testing"

	solanasdk "github.com/medeirosfalante/solana-sdk"
)

func TestGetByBlock(t *testing.T) {
	endpoint := solanasdk.TestNet_RPC
	client := solanasdk.New(endpoint)
	listAddress := []string{"2FdfSKnTK1yL35RMMpqBJBa4w29dcgHkuc2i7dJTuFPh"}
	tx, err := client.GetConfirmedBlockFindDeposit(context.TODO(), 109761983, listAddress)
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if len(tx) <= 0 {
		t.Errorf("invalid counter  need more than 0")
		return
	}

}
