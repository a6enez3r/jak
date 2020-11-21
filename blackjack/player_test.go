package blackjack

import "testing"

func TestPlayer(t *testing.T) {
	// create player
	player := Player{}
	// assert fields are nil
	if !(player.Wallet == 0) {
		t.Errorf("invalid player wallet at init")
	}
	if !(player.CurrentBet == 0) {
		t.Errorf("invalid player bet at init")
	}
	if !(player.Name == "") {
		t.Errorf("invalid player name at init")
	}
}

func TestPlayerHit(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// make test player
	player := Player{}
	// run player hit method
	player.Hit(&deck)
	// check if drawn cards changed
	if len(player.Drawn) == 0 {
		t.Errorf("player hit failed")
	}
}

func benchmarkPlayerHit(i int, b *testing.B) {
	// make test deck
	deck := MakeDeck()
	// make test player
	player := Player{}
	// run the Hit function b.N times
	for n := 0; n < b.N; n++ {
		// run player hit method
		player.Hit(&deck)
	}
}

func BenchmarkPlayerHit1(b *testing.B) {
	benchmarkPlayerHit(1, b)
}

func BenchmarkPlayerHit10(b *testing.B) {
	benchmarkPlayerHit(10, b)
}

func BenchmarkPlayerHit20(b *testing.B) {
	benchmarkPlayerHit(20, b)
}

func BenchmarkPlayerHit40(b *testing.B) {
	benchmarkPlayerHit(40, b)
}

func TestPlayerBet(t *testing.T) {
	// wallet value
	wallet := 60
	// make test player
	player := Player{Wallet: wallet}
	// bet value
	bet := 50
	// run player bet method
	player.Bet(bet)
	// check if bet subtracted from wallet
	if player.CurrentBet != bet {
		t.Errorf("invalid current bet for test player")
	}
	if player.Wallet != wallet-bet {
		t.Errorf("invalid wallet value for test player")
	}
	// run player bet method
	err := player.Bet(bet)
	// check err is nil
	if err == nil {
		t.Errorf("insufficient wallet expected error")
	}
}

func benchmarkPlayerBet(i int, b *testing.B) {
	// wallet value
	wallet := 10000000
	// make test player
	player := Player{Wallet: wallet}
	// bet value
	bet := 50
	// run the Bet function b.N times
	for n := 0; n < b.N; n++ {
		// run player bet method
		player.Bet(bet)
	}
}

func BenchmarkPlayerBet1(b *testing.B) {
	benchmarkPlayerBet(1, b)
}

func BenchmarkPlayerBet10(b *testing.B) {
	benchmarkPlayerBet(10, b)
}

func BenchmarkPlayerBet20(b *testing.B) {
	benchmarkPlayerBet(20, b)
}

func BenchmarkPlayerBet40(b *testing.B) {
	benchmarkPlayerBet(40, b)
}

func TestPlayerRefresh(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// make test player
	player := Player{}
	// run player hit method
	player.Hit(&deck)
	// check if drawn cards changed
	if len(player.Drawn) == 0 {
		t.Errorf("player hit failed")
	}
	// run player refresh method
	player.Refresh()
	// check if drawn cards changed
	if len(player.Drawn) != 0 {
		t.Errorf("player refresh failed")
	}
}

func benchmarkPlayerRefresh(i int, b *testing.B) {
	// make test deck
	deck := MakeDeck()
	// make test player
	player := Player{}
	// run player hit method
	player.Hit(&deck)
	// check if drawn cards changed
	if len(player.Drawn) == 0 {
		b.Errorf("player hit failed")
	}
	// run the Refresh function b.N times
	for n := 0; n < b.N; n++ {
		// run player refresh method
		player.Refresh()
	}
}

func BenchmarkPlayerRefresh1(b *testing.B) {
	benchmarkPlayerRefresh(1, b)
}

func BenchmarkPlayerRefresh10(b *testing.B) {
	benchmarkPlayerRefresh(10, b)
}

func BenchmarkPlayerRefresh20(b *testing.B) {
	benchmarkPlayerRefresh(20, b)
}

func BenchmarkPlayerRefresh40(b *testing.B) {
	benchmarkPlayerRefresh(40, b)
}

func TestPlayerSumHand(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// make test player
	player := Player{}
	// run player hit method
	player.Hit(&deck)
	// check if drawn cards changed
	if len(player.Drawn) == 0 {
		t.Errorf("player hit failed")
	}
	// check sum hand
	if player.sumHand() == 0 {
		t.Errorf("player sum hand failed")
	}
}

func benchmarkPlayerSumHand(i int, b *testing.B) {
	// make test deck
	deck := MakeDeck()
	// make test player
	player := Player{}
	// run player hit method
	player.Hit(&deck)
	// check if drawn cards changed
	if len(player.Drawn) == 0 {
		b.Errorf("player hit failed")
	}
	// run the sumHand function b.N times
	for n := 0; n < b.N; n++ {
		// run player refresh method
		player.sumHand()
	}
}

func BenchmarkPlayerSumHand1(b *testing.B) {
	benchmarkPlayerSumHand(1, b)
}

func BenchmarkPlayerSumHand10(b *testing.B) {
	benchmarkPlayerSumHand(10, b)
}

func BenchmarkPlayerSumHand100(b *testing.B) {
	benchmarkPlayerSumHand(100, b)
}

func BenchmarkPlayerSumHand1000(b *testing.B) {
	benchmarkPlayerSumHand(1000, b)
}
