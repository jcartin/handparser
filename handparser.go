package handparser

type Parser struct {}

type Card struct {
	suit string 		// the suit value (C,D,S,H) 
	value string 		// the card value (1,2,3,4,5,6,7,8,9,T,J,Q,K,A)
}

type Player struct {
	name string
	card1 Card
	card2 Card
	card3 Card
	card4 Card
	card5 Card
}

var players = make([]Player, 0)

// this function will take in a string of format:
//	Player C1 C2 C3 C4 C5 Player C1 C2 C3 C4 C5
// and parse the contents into the a player plus 
// a set of cards
func (h *Parser) Parse(play string) (*Player, error) {



	return nil, nil
}