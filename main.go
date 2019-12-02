package main

import (
	"fmt"
	"math/big"

	"github.com/binance-chain/tss-lib/tss"
)

func main() {

	p2pPubkey1 := big.NewInt(1)
	p2pPubkey2 := big.NewInt(2)

	fmt.Println(p2pPubkey1)
	fmt.Println(p2pPubkey2)

	party1 := tss.NewPartyID("1", "moniker1", p2pPubkey1)
	party2 := tss.NewPartyID("2", "moniker1", p2pPubkey2)
	fmt.Println(party1)
	fmt.Println(party2)
}
