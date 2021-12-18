package blackjack

import (
	"log"
	"os"
	"testing"
)

func TestBlackjack(t *testing.T) {
	// create blackjack
	blackjack := Blackjack{}
	// assert fields are nil
	if !(blackjack.PlayerSum == 0) {
		t.Errorf("invalid blackjack player sum at init")
	}
	if !(blackjack.DealerSum == 0) {
		t.Errorf("invalid blackjack dealer sum at init")
	}
}

func TestBlackjackInit(t *testing.T) {
	// test player name
	playerName := "test"
	// init blackjack
	blackjack := Init(playerName)
	// assert fields are nil
	if !(blackjack.PlayerSum == 0) {
		t.Errorf("invalid blackjack player sum at init")
	}
	if !(blackjack.DealerSum == 0) {
		t.Errorf("invalid blackjack dealer sum at init")
	}
}

func benchmarkBlackjackInit(i int, b *testing.B) {
	// test player name
	playerName := "test"
	// run the init function b.N times
	for n := 0; n < b.N; n++ {
		// create blackjack
		Init(playerName)
	}
}

func BenchmarkBlackjackInit1(b *testing.B) {
	benchmarkBlackjackInit(1, b)
}

func BenchmarkBlackjackInit100(b *testing.B) {
	benchmarkBlackjackInit(100, b)
}

func BenchmarkBlackjackInit10000(b *testing.B) {
	benchmarkBlackjackInit(10000, b)
}

// todo: fix flaky deal test
func testBlackjackDeal(t *testing.T) {
	// test player name
	playerName := "test"
	// init blackjack
	blackjack := Init(playerName)
	// deal
	blackjack.Deal()
	// assert num drawn cards != 0
	if len(blackjack.Player.Drawn) < 3 {
		t.Errorf("invalid blackjack player drawn after deal")
	}
	if len(blackjack.Dealer.Drawn) < 3 {
		t.Errorf("invalid blackjack dealer drawn after deal")
	}
}

func benchmarkBlackjackDeal(i int, b *testing.B) {
	// test player name
	playerName := "test"
	// create blackjack
	blackjack := Init(playerName)
	// run the init function b.N times
	for n := 0; n < b.N; n++ {
		// deal
		blackjack.Deal()
	}
}

func BenchmarkBlackjackDeal1(b *testing.B) {
	benchmarkBlackjackDeal(1, b)
}

func BenchmarkBlackjackDeal100(b *testing.B) {
	benchmarkBlackjackDeal(100, b)
}

func BenchmarkBlackjackDeal10000(b *testing.B) {
	benchmarkBlackjackDeal(10000, b)
}

func TestBlackjackCheck(t *testing.T) {
	// test player name
	playerName := "test"
	// init blackjack
	blackjack := Init(playerName)
	// check
	blackjack.Check()
	// assert num drawn cards != 0
	if blackjack.PlayerSum == 0 {
		t.Errorf("invalid blackjack player sum after check")
	}
	if blackjack.DealerSum == 0 {
		t.Errorf("invalid blackjack dealer sum after check")
	}
	// repeatedly hit
	for i := 0; i < 10; i++ {
		blackjack.Player.Hit(blackjack.Deck)
	}
	// check
	blackjack.Check()
	// assert game state over
	if blackjack.State != "over" {
		t.Errorf("invalid blackjack game state after check")
	}
}

func benchmarkBlackjackCheck(i int, b *testing.B) {
	// test player name
	playerName := "test"
	// create blackjack
	blackjack := Init(playerName)
	// run the init function b.N times
	for n := 0; n < b.N; n++ {
		// check
		blackjack.Check()
	}
}

func BenchmarkBlackjackCheck1(b *testing.B) {
	benchmarkBlackjackCheck(1, b)
}

func BenchmarkBlackjackCheck100(b *testing.B) {
	benchmarkBlackjackCheck(100, b)
}

func BenchmarkBlackjackCheck10000(b *testing.B) {
	benchmarkBlackjackCheck(10000, b)
}

func TestBlackjackAnother(t *testing.T) {
	// test player name
	playerName := "test"
	// init blackjack
	blackjack := Init(playerName)
	// play another
	blackjack.Another()
	// assert num drawn cards != 0
	if !(blackjack.PlayerSum == 0) {
		t.Errorf("invalid blackjack player sum after another")
	}
	if !(blackjack.DealerSum == 0) {
		t.Errorf("invalid blackjack dealer sum after another")
	}
}

func benchmarkBlackjackAnother(i int, b *testing.B) {
	// test player name
	playerName := "test"
	// create blackjack
	blackjack := Init(playerName)
	// run the init function b.N times
	for n := 0; n < b.N; n++ {
		// play another
		blackjack.Another()
	}
}

func BenchmarkBlackjackAnother1(b *testing.B) {
	benchmarkBlackjackAnother(1, b)
}

func BenchmarkBlackjackAnother100(b *testing.B) {
	benchmarkBlackjackAnother(100, b)
}

func BenchmarkBlackjackAnother10000(b *testing.B) {
	benchmarkBlackjackAnother(10000, b)
}

func TestBlackjackPlay(t *testing.T) {
	// test player name
	playerName := "test"
	// init blackjack
	blackjack := Init(playerName)
	// play
	blackjack.Play("hit")
	// assert last action
	if blackjack.LastAction != "hit" {
		t.Errorf("invalid blackjack last action after play")
	}
	// play
	blackjack.Play("stand")
	// assert last action
	if blackjack.LastAction != "stand" {
		t.Errorf("invalid blackjack last action after play")
	}
	// play
	blackjack.Play("double")
	// assert last action
	if blackjack.LastAction != "double" {
		t.Errorf("invalid blackjack last action after play")
	}
}

func benchmarkBlackjackPlay(i int, b *testing.B) {
	// test player name
	playerName := "test"
	// create blackjack
	blackjack := Init(playerName)
	// run the init function b.N times
	for n := 0; n < b.N; n++ {
		// play
		blackjack.Play("hit")
	}
}

func BenchmarkBlackjackPlay1(b *testing.B) {
	benchmarkBlackjackPlay(1, b)
}

func BenchmarkBlackjackPlay100(b *testing.B) {
	benchmarkBlackjackPlay(100, b)
}

func BenchmarkBlackjackPlay10000(b *testing.B) {
	benchmarkBlackjackPlay(10000, b)
}

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

func TestBlackjackDisplay(t *testing.T) {
	// supress print
	defer quiet()()
	// test player name
	playerName := "test"
	// init blackjack
	blackjack := Init(playerName)
	// display
	blackjack.Display()
	// set bet
	blackjack.Player.Bet(10)
	// display
	blackjack.Display()
	// set game state to over
	blackjack.State = "over"
	// display
	blackjack.Display()
}

func benchmarkBlackjackDisplay(i int, b *testing.B) {
	// supress print
	defer quiet()()
	// test player name
	playerName := "test"
	// create blackjack
	blackjack := Init(playerName)
	// run the init function b.N times
	for n := 0; n < b.N; n++ {
		// display
		blackjack.Display()
	}
}

func BenchmarkBlackjackDisplay1(b *testing.B) {
	benchmarkBlackjackDisplay(1, b)
}

func BenchmarkBlackjackDisplay100(b *testing.B) {
	benchmarkBlackjackDisplay(100, b)
}

func BenchmarkBlackjackDisplay10000(b *testing.B) {
	benchmarkBlackjackDisplay(10000, b)
}
