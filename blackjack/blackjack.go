package blackjack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Blackjack : will have a player, a dealer and a deck
type Blackjack struct {
	Deck       *Deck   `json:"deck"`       // deck of cards
	Dealer     *Dealer `json:"dealer"`     // dealer struct
	Player     *Player `json:"player"`     // player struct
	State      string  `json:"state"`      // game state (over or in progress)
	PlayerSum  int     `json:"playerSum"`  // player hand sum
	DealerSum  int     `json:"dealerSum"`  // dealer hand sum
	LastAction string  `json:"lastAction"` // last player action
}

// Display : show game info
func (b *Blackjack) Display() {
	// print player wallet
	fmt.Println("-----------------")
	fmt.Println("wallet:- ", b.Player.Wallet)
	fmt.Println("-----------------")
	// show one dealer card
	fmt.Println("dealer has:- ", b.Dealer.Drawn[0])
	fmt.Println("-----------------")
	// show cards
	fmt.Println("hand:- ", b.Player.Drawn)
	fmt.Println("-----------------")
	// show bet if placed
	if b.Player.CurrentBet != 0 {
		fmt.Println("bet:- ", b.Player.CurrentBet)
		fmt.Println("-----------------")
	}
	// show game info if over
	if b.State == "over" {
		// winning conditions
		if b.PlayerSum > 21 || b.DealerSum > b.PlayerSum {
			// display results
			fmt.Println("you:- ", b.PlayerSum)
			fmt.Println("dealer:- ", b.DealerSum)
			fmt.Println("dealer wins:- ", b.Player.CurrentBet)
		}
		if b.DealerSum > 21 || b.DealerSum < b.PlayerSum {
			// update wallet
			b.Player.Wallet += 2 * b.Player.CurrentBet
			// display results
			fmt.Println("you:- ", b.PlayerSum)
			fmt.Println("dealer:- ", b.DealerSum)
			fmt.Println("you win:-", b.Player.CurrentBet)
		}
	}
}

// Init : initialize blackjack game
func Init(playerName string) Blackjack {
	// make deck
	deck := MakeDeck()
	// make dealer
	dealer := Dealer{}
	// make player
	player := Player{Wallet: 50, Name: playerName}
	// create blackjack struct
	b := Blackjack{Deck: &deck, Dealer: &dealer, Player: &player}
	// start game play
	b.Deal()
	// return black jack
	return b
}

// Deal : deal cards
func (b *Blackjack) Deal() {
	// get dealer sum
	dealerSum := b.Dealer.sumHand()
	// if dealer sum under 17 always hit
	if dealerSum < 17 {
		// hit dealer
		b.Dealer.Hit(b.Deck)
	}
	// hit player
	b.Player.Hit(b.Deck)
}

// Check : check if there is a winner
func (b *Blackjack) Check() {
	// get player hand sum
	playerSum := b.Player.sumHand()
	// get dealer sum
	dealerSum := b.Dealer.sumHand()
	// set sums
	b.PlayerSum = playerSum
	b.DealerSum = dealerSum
	// check if drawn cards add up to more than 21
	if b.PlayerSum > 21 || b.DealerSum > 21 {
		// if so end game
		b.State = "over"
	}
}

// Another : play another round of a given game
func (b *Blackjack) Another() {
	// reset drawn cards
	b.Player.Refresh()
	b.Dealer.Refresh()
	// reset sums
	b.PlayerSum, b.DealerSum = 0, 0
	// reset game state
	b.State = "in progress"
	// deal first round
	b.Deal()
}

// Play : blackjack game play
func (b *Blackjack) Play(playerAction string) {
	switch {
	case playerAction == "hit":
		// set last action
		b.LastAction = "hit"
		// hit player
		b.Deal()
		// check winner
		b.Check()
	case playerAction == "stand":
		// set last action
		b.LastAction = "stand"
		// check winner
		b.Check()
		// set game state to over
		b.State = "over"
	case playerAction == "double":
		// set last action
		b.LastAction = "double"
		// hit player
		b.Player.Hit(b.Deck)
		// double bet
		b.Player.Wallet -= b.Player.CurrentBet
		b.Player.CurrentBet *= 2
		// check winner
		b.Check()
		// set game state to over
		b.State = "over"
	}
}

// Save : save game to JSON
func (b *Blackjack) Save() {
	// marshal struct
	file, _ := json.MarshalIndent(b, "", " ")
	// write to file
	_ = ioutil.WriteFile("game.json", file, 0644)
}

// Load : load game from JSON
func Load() {
	fmt.Println(filepath.Abs("./"))
	// write to file
	dat, _ := ioutil.ReadFile("game.json")
	var blackjack Blackjack
	json.Unmarshal([]byte(dat), &blackjack)
	fmt.Println(blackjack)
}

// save and load functionality
// add API to run gameplay from a RESTful endpoint
// update main to start cli and optionally the API
