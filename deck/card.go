//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Suit represents suit of a card
type Suit uint8

const (
	//Spade represents spade suit
	Spade Suit = iota
	//Diamond represents diamond suit
	Diamond
	//Club represents club suit
	Club
	//Heart represents heart suit
	Heart
	//Joker represents joker card
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

//Rank represents rank of a card
type Rank uint8

const (
	_ Rank = iota
	//Ace represents ace rank
	Ace
	//Two represents two rank
	Two
	//Three represents three rank
	Three
	//Four represents four rank
	Four
	//Five represents five rank
	Five
	//Six represents six rank
	Six
	//Seven represents seven rank
	Seven
	//Eight represents eight rank
	Eight
	//Nine represents nine rank
	Nine
	//Ten represents ten rank
	Ten
	//Jack represents jack rank
	Jack
	//Queen represents queen rank
	Queen
	//King represents king rank
	King
)

const (
	minRank = Ace
	maxRank = King
)

//Card represents a card of real life
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

//New creates a new deck of cards
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

//DefaultSort sorts the cards
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Sort implements any custom sorting needed
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

//Less function implementation for sorting
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

//Permer provides an interdace for Permutation function for randomness
type Permer interface {
	Perm(n int) []int
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

//Shuffle returns a new shuffled deck
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

//Jokers return a new deck with specified jokers in it
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

//Filter returns a function which filters the card based on the given condition
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

//Deck returns the given amount of new deck
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
