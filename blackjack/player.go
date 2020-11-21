package blackjack

import "fmt"

// Player : struct to maintain player wallet, bets and drawn cards
type Player struct {
	Wallet     int    `json:"wallet"`
	Drawn      []Card `json:"drawn"`
	Name       string `json:"name"`
	CurrentBet int    `json:"currentBet"`
}

// sumHand : sum a given set of cards
func (p *Player) sumHand() int {
	// init var to store sum
	sum := 0
	// iteratively sum cards drawn
	for _, card := range p.Drawn {
		// add to sum
		sum += card.Value
	}
	// return sum
	return sum
}

// Refresh : remove any drawn cards from player
func (p *Player) Refresh() {
	// set player drawn to empty
	p.Drawn = []Card{}
	// set current bet to 0
	p.CurrentBet = 0
}

// Bet : place a bet & subtract amount from Deck
func (p *Player) Bet(betAmount int) error {
	// check if player has enough in wallet
	if p.Wallet < betAmount {
		return fmt.Errorf("insufficient wallet: %v", p.Wallet)
	}
	// set current bet
	p.CurrentBet = betAmount
	// subtract bet amount from wallet
	p.Wallet -= p.CurrentBet
	// return nil
	return nil
}

// Hit : draw a card from a deck
func (p *Player) Hit(d *Deck) {
	// variable to store number of cards
	var numCards int
	// check if player has drawn
	if len(p.Drawn) == 0 {
		// set num drawn cards
		numCards = 2
	} else {
		// set num drawn cards
		numCards = 1
	}
	// iteratively draw cards
	for i := 0; i < numCards; i++ {
		// draw card from deck
		cardDrawn := d.Draw()
		// add drawn card to player drawn list
		p.Drawn = append(p.Drawn, cardDrawn)
	}
}
