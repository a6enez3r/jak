package blackjack

import "testing"

func TestDealerHit(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// make test dealer
	dealer := Dealer{}
	// run dealer hit method
	dealer.Hit(&deck)
	// check if drawn cards changed
	if len(dealer.Drawn) == 0 {
		t.Errorf("dealer hit failed")
	}
}

func benchmarkDealerHit(i int, b *testing.B) {
	// make test deck
	deck := MakeDeck()
	// make test dealer
	dealer := Dealer{}
	// run the Hit function b.N times
	for n := 0; n < b.N; n++ {
		// run dealer hit method
		dealer.Hit(&deck)
	}
}

func BenchmarkDealerHit1(b *testing.B) {
	benchmarkDealerHit(1, b)
}

func BenchmarkDealerHit10(b *testing.B) {
	benchmarkDealerHit(10, b)
}

func BenchmarkDealerHit20(b *testing.B) {
	benchmarkDealerHit(20, b)
}

func BenchmarkDealerHit40(b *testing.B) {
	benchmarkDealerHit(40, b)
}

func TestDealerRefresh(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// make test dealer
	dealer := Dealer{}
	// run dealer hit method
	dealer.Hit(&deck)
	// check if drawn cards changed
	if len(dealer.Drawn) == 0 {
		t.Errorf("dealer hit failed")
	}
	// run dealer refresh method
	dealer.Refresh()
	// check if drawn cards changed
	if len(dealer.Drawn) != 0 {
		t.Errorf("dealer refresh failed")
	}
}

func benchmarkDealerRefresh(i int, b *testing.B) {
	// make test deck
	deck := MakeDeck()
	// make test dealer
	dealer := Dealer{}
	// run dealer hit method
	dealer.Hit(&deck)
	// check if drawn cards changed
	if len(dealer.Drawn) == 0 {
		b.Errorf("dealer hit failed")
	}
	// run the Refresh function b.N times
	for n := 0; n < b.N; n++ {
		// run dealer refresh method
		dealer.Refresh()
	}
}

func BenchmarkDealerRefresh1(b *testing.B) {
	benchmarkDealerRefresh(1, b)
}

func BenchmarkDealerRefresh10(b *testing.B) {
	benchmarkDealerRefresh(10, b)
}

func BenchmarkDealerRefresh20(b *testing.B) {
	benchmarkDealerRefresh(20, b)
}

func BenchmarkDealerRefresh40(b *testing.B) {
	benchmarkDealerRefresh(40, b)
}

func TestDealerSumHand(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// make test dealer
	dealer := Dealer{}
	// run dealer hit method
	dealer.Hit(&deck)
	// check if drawn cards changed
	if len(dealer.Drawn) == 0 {
		t.Errorf("dealer hit failed")
	}
	// check sum hand
	if dealer.sumHand() == 0 {
		t.Errorf("dealer sum hand failed")
	}
}

func benchmarkDealerSumHand(i int, b *testing.B) {
	// make test deck
	deck := MakeDeck()
	// make test dealer
	dealer := Dealer{}
	// run dealer hit method
	dealer.Hit(&deck)
	// check if drawn cards changed
	if len(dealer.Drawn) == 0 {
		b.Errorf("dealer hit failed")
	}
	// run the sumHand function b.N times
	for n := 0; n < b.N; n++ {
		// run dealer refresh method
		dealer.sumHand()
	}
}

func BenchmarkDealerSumHand1(b *testing.B) {
	benchmarkDealerSumHand(1, b)
}

func BenchmarkDealerSumHand10(b *testing.B) {
	benchmarkDealerSumHand(10, b)
}

func BenchmarkDealerSumHand100(b *testing.B) {
	benchmarkDealerSumHand(100, b)
}

func BenchmarkDealerSumHand1000(b *testing.B) {
	benchmarkDealerSumHand(1000, b)
}
