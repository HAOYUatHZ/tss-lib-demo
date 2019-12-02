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
	// fmt.Println(party1)
	// fmt.Println(party2)

	parties := make([]*tss.PartyID, 2)
	parties[0] = party1
	parties[1] = party2

	partyIDMap := make(map[string]*tss.PartyID)
	for _, id := range parties {
		partyIDMap[id.Id] = id
	}
	fmt.Println(partyIDMap)

	ctx := tss.NewPeerContext(tss.SortPartyIDs(parties))
	params := tss.NewParameters(ctx, party1, len(parties), 1)
	preParams, _ := keygen.GeneratePreParams(1 * time.Minute)

	// TODO:
	party := keygen.NewLocalParty(params, outCh, endCh, preParams)
}
