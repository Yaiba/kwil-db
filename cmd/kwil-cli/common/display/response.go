package display

import (
	"fmt"
	txTypes "kwil/x/types/transactions"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func PrintTxResponse(res *txTypes.Response) {
	if res.Hash != nil {
		fmt.Println("Success!")
	}
	fmt.Println("Response:")
	fmt.Println("  Hash:", hexutil.Encode(res.Hash))
	fmt.Println("  Fee:", res.Fee)
}

type ClientChainResponse struct {
	Chain string `json:"chain"`
	Tx    string `json:"tx"`
}

func PrintClientChainResponse(res *ClientChainResponse) {
	fmt.Println("Response:")
	fmt.Println("  Fund:", res.Chain)
	fmt.Println("  Tx:", res.Tx)
}
