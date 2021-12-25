package solanasdk

import (
	"context"
	"encoding/json"
)

type Meta struct {
	Fee               int64                   `json:"fee"`
	PostBalances      []int64                 `json:"postBalances"`
	PreBalances       []int64                 `json:"preBalances"`
	PostTokenBalances []PostTokenBalancesItem `json:"postTokenBalances"`
	PreTokenBalances  []PostTokenBalancesItem `json:"preTokenBalances"`
	Rewards           []int64                 `json:"rewards"`
	Status            map[string]interface{}  `json:"status"`
}

type PostTokenBalancesItem struct {
	AccountIndex  int64                              `json:"accountIndex"`
	Mint          string                             `json:"mint"`
	Owner         string                             `json:"owner"`
	UiTokenAmount PostTokenBalancesItemUiTokenAmount `json:"uiTokenAmount"`
}
type PostTokenBalancesItemUiTokenAmount struct {
	Amount         string  `json:"amount"`
	Decimals       int64   `json:"decimals"`
	UiAmount       float64 `json:"uiAmount"`
	UiAmountString string  `json:"uiAmountString"`
}

type TransactionItem struct {
	Message TransactionItemMessage `json:"message"`
}

type TransactionItemMessage struct {
	AccountKeys     []string                            `json:"accountKeys"`
	Header          TransactionItemMessageHedear        `json:"header"`
	Instructions    []TransactionItemMessageInstruction `json:"instructions"`
	TecentBlockhash string                              `json:"recentBlockhash"`
}

type TransactionItemMessageHedear struct {
	NumReadonlySignedAccounts   int64    `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int64    `json:"numReadonlyUnsignedAccounts"`
	NumRequiredSignatures       int64    `json:"numRequiredSignatures"`
	Signatures                  []string `json:"signatures"`
}

type TransactionItemMessageInstruction struct {
	Accounts       []int64 `json:"accounts"`
	Data           string  `json:"data"`
	ProgramIdIndex int32   `json:"programIdIndex"`
}

type Transaction struct {
	Meta        Meta            `json:"meta"`
	Transaction TransactionItem `json:"transaction"`
}

type Block struct {
	Blockhash         string        `json:"blockhash"`
	ParentSlot        int64         `json:"parentSlot"`
	PreviousBlockhash string        `json:"previousBlockhash"`
	Transactions      []Transaction `json:"transactions"`
	BlockHeight       int64         `json:"blockHeight"`
	BlockTime         int64         `json:"blockTime"`
}

type TransactionBind struct {
	Amount   float64 `json:"amount"`
	Address  string  `json:"address"`
	Currency string  `json:"currency"`
	Mint     string  `json:"mint"`
	Native   bool    `json:"native"`
	Txid     string  `json:"txid"`
}

type BlockQuery struct {
	Encoding           string `json:"encoding"`
	TransactionDetails string `json:"transactionDetails"`
	Rewards            bool   `json:"rewards"`
}

func (t *Client) GetBlock(ctx context.Context, number int32) (*Block, error) {
	var blockRef Block
	response, err := t.client.CallComand(ctx, "getBlock", number, &BlockQuery{
		Encoding:           "json",
		TransactionDetails: "full",
		Rewards:            false,
	})
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(response.Result)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &blockRef)
	if err != nil {
		return nil, err
	}
	return &blockRef, nil

}

func (t *Client) GetConfirmedBlockFindDeposit(ctx context.Context, number int32, address []string) ([]*TransactionBind, error) {
	var transactions []*TransactionBind
	block, err := t.GetBlock(context.TODO(), number)
	if err != nil {
		return nil, err
	}
	for _, item := range block.Transactions {
		if len(item.Transaction.Message.AccountKeys) > 2 {
			for _, itemRef := range address {
				if item.Transaction.Message.AccountKeys[1] == itemRef {
					amountIntPre := item.Meta.PreBalances[1]
					amountIntPos := item.Meta.PostBalances[1]
					totalAmountInt := amountIntPos - amountIntPre
					totalAmountFloat := float64(totalAmountInt / 1000000000)
					txRef := &TransactionBind{Address: itemRef, Amount: totalAmountFloat, Currency: "SOL", Native: true, Txid: item.Transaction.Message.Header.Signatures[0]}
					transactions = append(transactions, txRef)
				}
			}
		}
		if len(item.Meta.PostTokenBalances) > 0 {
			for _, itemBalance := range item.Meta.PostTokenBalances {
				for _, itemRef := range address {
					if itemBalance.Owner == itemRef {
						itemBalance.Owner = itemRef
						txRef := &TransactionBind{Address: itemRef, Amount: itemBalance.UiTokenAmount.UiAmount, Mint: itemBalance.Mint, Native: false, Txid: item.Transaction.Message.Header.Signatures[0]}
						transactions = append(transactions, txRef)
					}

				}
			}
		}

	}
	return transactions, nil
}
