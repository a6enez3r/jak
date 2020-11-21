package blackjack

import (
	"reflect"
	"testing"
)

// stringInSlice: check membership of a string in a slice
func stringInSlice(s string, sls []string) bool {
	// iterate through slice
	for _, elem := range sls {
		// compare if element matches string
		if elem == s {
			// if it does return true immediately
			return true
		}
	}
	// return false / elem not found
	return false
}

// intInSlice: check membership of a int in a slice
func intInSlice(i int, sls []int) bool {
	// iterate through slice
	for _, elem := range sls {
		// compare if element matches string
		if elem == i {
			// if it does return true immediately
			return true
		}
	}
	// return false / elem not found
	return false
}

func TestMakeDeck(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// check if type matches
	if reflect.TypeOf(deck) != reflect.TypeOf(Deck{}) {
		t.Errorf("make deck failed")
	}
}

func benchmarkMakeDeck(i int, b *testing.B) {
	// run the Make Deck function b.N times
	for n := 0; n < b.N; n++ {
		// make test card array
		MakeDeck()
	}
}

func BenchmarkMakeDeck1(b *testing.B) {
	benchmarkMakeDeck(1, b)
}

func BenchmarkMakeDeck100(b *testing.B) {
	benchmarkMakeDeck(100, b)
}

func BenchmarkMakeDeck10000(b *testing.B) {
	benchmarkMakeDeck(10000, b)
}

func BenchmarkMakeDeck1000000(b *testing.B) {
	benchmarkMakeDeck(1000000, b)
}

func TestDraw(t *testing.T) {
	// make test deck
	deck := MakeDeck()
	// draw card
	cardDrawn := deck.Draw()
	// check if type matches
	if reflect.TypeOf(cardDrawn) != reflect.TypeOf(Card{}) {
		t.Errorf("draw card failed")
	}
	// check suite
	if !stringInSlice(cardDrawn.Suite, deck.Suites) {
		t.Errorf("invalid drawn card suite")
	}
	// check value
	if !intInSlice(cardDrawn.Value, deck.Values) {
		t.Errorf("invalid drawn card suite")
	}
	// check name
	if !stringInSlice(cardDrawn.Name, deck.Names) {
		t.Errorf("invalid drawn card suite")
	}
}

func TestRemoveDrawn(t *testing.T) {
	// make test card array
	deck := MakeDeck()
	// remove first card
	removedDeck := removeDrawn(deck.Cards, 1)
	// assert len of two is not the same
	if len(deck.Cards) == len(removedDeck) {
		t.Errorf("remove drawn card failed")
	}
}

func benchmarkRemoveDrawn(i int, b *testing.B) {
	// make test card array
	deck := MakeDeck()
	// run the remove drawn function b.N times
	for n := 0; n < b.N; n++ {
		// remove first card
		removeDrawn(deck.Cards, 1)
	}
}

func BenchmarkRemoveDrawn1(b *testing.B) {
	benchmarkRemoveDrawn(1, b)
}

func BenchmarkRemoveDrawn10(b *testing.B) {
	benchmarkRemoveDrawn(10, b)
}

func BenchmarkRemoveDrawn20(b *testing.B) {
	benchmarkRemoveDrawn(20, b)
}

func BenchmarkRemoveDrawn40(b *testing.B) {
	benchmarkRemoveDrawn(40, b)
}

func TestGenerateIndex(t *testing.T) {
	// generate test index
	randIndex := generateIndex(5)
	// check type
	if !(reflect.TypeOf(randIndex) == reflect.TypeOf(5)) {
		t.Errorf("invalid random index generated")
	}
}

func benchmarkGenerateIndex(i int, b *testing.B) {
	// run the generate index function b.N times
	for n := 0; n < b.N; n++ {
		// remove first card
		generateIndex(i)
	}
}

func BenchmarkGenerateIndex1(b *testing.B) {
	benchmarkGenerateIndex(1, b)
}

func BenchmarkGenerateIndex100(b *testing.B) {
	benchmarkGenerateIndex(100, b)
}

func BenchmarkGenerateIndex10000(b *testing.B) {
	benchmarkGenerateIndex(10000, b)
}

func BenchmarkGenerateIndex1000000(b *testing.B) {
	benchmarkGenerateIndex(1000000, b)
}

func TestShuffle(t *testing.T) {
	// make test card array
	deck := MakeDeck()
	// get before shuffle card
	beforeShuffleCard := deck.Cards[0]
	// shuffle deck
	deck.Shuffle()
	// get after shuffle card
	afterShuffleCard := deck.Cards[0]
	// assert the two cards are not the same
	if beforeShuffleCard == afterShuffleCard {
		t.Errorf("shuffle deck failed")
	}
	// check card restacks when reaching 0 cards
	for i := 0; i < 52; i++ {
		// draw card
		deck.Draw()
	}
	// check all cards have been drawn
	if len(deck.Cards) != 0 {
		t.Errorf("invalid num cards detected before shuffling")
	}
	// check if shuffle reuses dealt cards
	deck.Shuffle()
	// check all cards have been drawn
	if len(deck.Cards) == 0 {
		t.Errorf("invalid num cards detected after shuffling")
	}
}

func benchmarkShuffle(i int, b *testing.B) {
	// make test card array
	deck := MakeDeck()
	// run the Shuffle function b.N times
	for n := 0; n < b.N; n++ {
		// shuffle deck
		deck.Shuffle()
	}
}

func BenchmarkShuffle1(b *testing.B) {
	benchmarkShuffle(1, b)
}

func BenchmarkShuffle100(b *testing.B) {
	benchmarkShuffle(100, b)
}

func BenchmarkShuffle10000(b *testing.B) {
	benchmarkShuffle(10000, b)
}

func BenchmarkShuffle1000000(b *testing.B) {
	benchmarkShuffle(1000000, b)
}
