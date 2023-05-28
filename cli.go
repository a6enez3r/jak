package main

import (
	"fmt"
	"strconv"

	"github.com/a6enez3r/jak/blackjack"

	"github.com/fatih/color"
)

// Cli : cli wrapper over blackjack game
type Cli struct {
	Game blackjack.Blackjack
}

// Init : init cli game
func Init() Cli {
	// create var to get player name
	var playerName string
	// user prompt
	color.Green("-----------------")
	color.Yellow("your name")
	color.Green("-----------------")
	// take input from user
	fmt.Scanln(&playerName)
	// init blackjack
	b := blackjack.Init(playerName)
	// add to struct
	cli := Cli{Game: b}
	// return cli
	return cli
}

// Play : start cli game play
func (c *Cli) Play() {
	// run until user exits
	for {
		// show game info
		c.Game.Display()
		// if no bet placed
		if c.Game.Player.CurrentBet == 0 {
			// create var to get player buy
			var betAmountStr string
			// user prompt
			color.Green("-----------------")
			color.Yellow("your buy-in")
			color.Green("-----------------")
			// take input from user & store
			fmt.Scanln(&betAmountStr)
			// convert to int
			betAmount, _ := strconv.Atoi(betAmountStr)
			// place bet
			c.Game.Player.Bet(betAmount)
		}
		// create var to get player action
		var playerAction string
		// user prompt
		color.Green("-----------------")
		color.Yellow("action (hit/stand/double)")
		color.Green("-----------------")
		// take input from user & store
		fmt.Scanln(&playerAction)
		// start game play
		c.Game.Play(playerAction)
		// check game state
		if c.Game.State == "over" {
			// show game info
			c.Game.Display()
			// user prompt
			color.Green("-----------------")
			color.Yellow("another hand? (y/n)")
			color.Green("-----------------")
			// take input from user & store
			fmt.Scanln(&playerAction)
			// if another hand
			if playerAction == "y" {
				// start another game
				c.Game.Another()
				// show info
				c.Game.Display()
			} else {
				// break
				break
			}
		}
	}
}

// BlackjackCli : play blackjack in terminal
func BlackjackCli() {
	// init
	cli := Init()
	// play
	cli.Play()
}
