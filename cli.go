package main

import (
	"fmt"
	"strconv"

	"github.com/abmamo/cards/blackjack"
)

// BlackjackCli : play blackjack in terminal
func BlackjackCli() {
	// create var to get player name
	var playerName string
	// user prompt
	fmt.Println("name:- ")
	fmt.Println("-----------------")
	// take input from user
	fmt.Scanln(&playerName)
	// init blackjack
	b := blackjack.Init(playerName)
	// run until user exits
	for {
		// show game info
		b.Display()
		// if no bet placed
		if b.Player.CurrentBet == 0 {
			// create var to get player buy
			var betAmountStr string
			// user prompt
			fmt.Println("bet:- ")
			fmt.Println("-----------------")
			// take input from user & store
			fmt.Scanln(&betAmountStr)
			// convert to int
			betAmount, _ := strconv.Atoi(betAmountStr)
			// place bet
			b.Player.Bet(betAmount)
		}
		// create var to get player action
		var playerAction string
		// user prompt
		fmt.Println("action (hit/stand/double):- ")
		fmt.Println("-----------------")
		// take input from user & store
		fmt.Scanln(&playerAction)
		// start game play
		b.Play(playerAction)
		// check game state
		if b.State == "over" {
			// show game info
			b.Display()
			// user prompt
			fmt.Println("another hand? (y/n)")
			fmt.Println("-----------------")
			// take input from user & store
			fmt.Scanln(&playerAction)
			// if another hand
			if playerAction == "y" {
				// start another game
				b.Another()
				// show info
				b.Display()
			} else {
				// break
				break
			}
		}
	}
}
