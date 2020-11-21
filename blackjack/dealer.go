package blackjack

// Dealer : struct to maintain dealer drawn cards
type Dealer struct {
	Drawn []Card `json:"drawn"`
}

// sumHand : sum a given set of cards
func (d *Dealer) sumHand() int {
	// init var to store sum
	sum := 0
	// iteratively sum cards drawn
	for _, card := range d.Drawn {
		// add to sum
		sum += card.Value
	}
	// return sum
	return sum
}

// Refresh : remove any drawn cards from dealer
func (d *Dealer) Refresh() {
	// set dealer drawn to empty
	d.Drawn = []Card{}
}

// Hit : draw a card from a deck
func (d *Dealer) Hit(deck *Deck) {
	// init variable to store number of cards
	var numCards int
	// check if dealer has drawn
	if len(d.Drawn) == 0 {
		// set num drawn cards
		numCards = 2
	} else {
		// set num drawn cards
		numCards = 1
	}
	// check current sum
	if d.sumHand() < 17 {
		// iteratively draw cards
		for i := 0; i < numCards; i++ {
			// draw card from deck
			cardDrawn := deck.Draw()
			// add drawn card to dealer drawn list
			d.Drawn = append(d.Drawn, cardDrawn)
		}
	}
}
