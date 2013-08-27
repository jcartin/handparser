package handparser

import (
	"testing"
	"fmt"
	)

func TestCanParseASimpleHand(t *testing.T) {
	hand := "Black: 2H 3D 5S 9C KD  White: 2C 3H 4S 8C AH"
	p := &Parser{}
	player, err := p.Parse(hand)

	if err != nil {
		t.Error("parsing returned an error", err)
	}

	fmt.Println(player)

	if player == nil {
		t.Error("returned player object is nil")
	}
	
}