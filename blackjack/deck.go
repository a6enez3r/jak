package blackjack

import (
	// import using alias to avoid conflict with math/rand
	cryprand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// Deck : collection of cards
type Deck struct {
	Cards  []Card
	Dealt  []Card
	Suites []string `json:"suites"`
	Names  []string `json:"names"`
	Values []int    `json:"values"`
}

// Card : atomic component of deck
type Card struct {
	Suite string `json:"suite"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// Shuffle : unshuffled cards are no fun so randomise their order
func (d *Deck) Shuffle() {
	// if all cards have been dealt
	if len(d.Cards) == 0 {
		// restack deck with dealt cards
		d.Cards, d.Dealt = d.Dealt, d.Cards
	}
	// get cards from deck
	cards := d.Cards
	// for each card
	for i := range cards {
		// get random index
		j := rand.Intn(i + 1)
		// swap card with card at random index
		cards[i], cards[j] = cards[j], cards[i]
	}
}

// generateIndex : all knowing oracle that generates a random index/number
func generateIndex(m int) int {
	// init var to store seed
	var s int64
	// generate seed & store in var
	binary.Read(cryprand.Reader, binary.LittleEndian, &s)
	// use generated number as seed
	rand.Seed(s)
	// generate random index & return
	return rand.Intn(m)
}

// removeDrawn : why draw a card if you can redraw it
func removeDrawn(s []Card, i int) []Card {
	// remove elemnt at index
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	// return slice with removed element
	return s[:len(s)-1]
}

// Draw : draw a card from a deck
func (d *Deck) Draw() Card {
	// shuffle deck
	d.Shuffle()
	// generate random index
	cardIndex := generateIndex(len(d.Cards))
	// add card @ index to hand
	cardDrawn := d.Cards[cardIndex]
	// add card @ index to dealt
	d.Dealt = append(d.Dealt, d.Cards[cardIndex])
	// remove card from undrawn cards
	d.Cards = removeDrawn(d.Cards, cardIndex)
	// return drawn card
	return cardDrawn
}

// MakeDeck : makes a deck
func MakeDeck() Deck {
	// constants
	suites := []string{"spades", "hearts", "diamonds", "clubs"}
	names := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "Queen", "King", "Ace"}
	values := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10}
	// init deck
	deck := Deck{Suites: suites, Names: names, Values: values}
	// for each suite
	for _, suite := range suites {
		// add all values & names
		for idx := range values {
			deck.Cards = append(deck.Cards, Card{Suite: suite, Name: names[idx], Value: values[idx]})
		}
	}
	// shuffle deck
	deck.Shuffle()
	// return deck
	return deck
}
