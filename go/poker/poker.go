package poker

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const testVersion = 5

// rankings for hands
const (
	highCard int = iota
	onePair
	twoPair
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	fiveOfAKind
)

// constants to define value of ace
const (
	aceHigh = 14
	aceLow  = 1
)

// internal representation of a card to rank hands
type card struct {
	rank int
	suit rune
}

// internal type for comparing hand rank
type ranking []int

// regex for parsing card from string representation of hand
var cardRegex = regexp.MustCompile(`\A([0-9]+|[JQKA])([♤♡♢♧])\z`)

// parseCards converts a string representation of a hand into a slice of card types
func parseCards(hand string, ace int) ([]card, error) {
	h := strings.Fields(hand)
	if len(h) != 5 {
		return []card{}, errors.New("poker: incorrect number of cards")
	}

	cards := make([]card, 0, len(h))

	var rank int

	for _, c := range h {
		m := cardRegex.FindStringSubmatch(c)
		if len(m) != 3 {
			return []card{}, errors.New("poker: invalid card")
		}

		switch m[1] {
		case "J":
			rank = 11
		case "Q":
			rank = 12
		case "K":
			rank = 13
		case "A":
			rank = ace
		default:
			r, err := strconv.Atoi(m[1])
			if err != nil || r < 2 || r > 10 {
				return []card{}, errors.New("poker: invalid card rank")
			}
			rank = r
		}

		suit := []rune(m[2])[0]

		cards = append(cards, card{rank: rank, suit: suit})
	}

	return cards, nil
}

// maxValue finds the maximum value in a map
func maxValue(m map[int]int) (max int) {
	for _, f := range m {
		if f > max {
			max = f
		}
	}

	return
}

// rankHand calculates a ranking for a hand based on the hand and the values of the component cards
func rankHand(hand string, ace int) (ranking, error) {
	cards, err := parseCards(hand, ace)
	if err != nil {
		return []int{}, err
	}

	ranks := map[int]int{}
	suits := map[rune]int{}

	for _, c := range cards {
		ranks[c.rank] = ranks[c.rank] + 1
		suits[c.suit] = suits[c.suit] + 1
	}

	sort.Slice(
		cards,
		func(i, j int) bool {
			if ranks[cards[i].rank] != ranks[cards[j].rank] {
				return ranks[cards[i].rank] > ranks[cards[j].rank]
			}

			return cards[i].rank > cards[j].rank
		},
	)

	r := make(ranking, 6)

	switch len(ranks) {
	case 1:
		r[0] = fiveOfAKind
	case 2:
		if maxValue(ranks) == 4 {
			r[0] = fourOfAKind
		} else {
			r[0] = fullHouse
		}
	case 3:
		if maxValue(ranks) == 3 {
			r[0] = threeOfAKind
		} else {
			r[0] = twoPair
		}
	case 4:
		r[0] = onePair
	case 5:
		if cards[0].rank == cards[4].rank+4 {
			if len(suits) == 1 {
				r[0] = straightFlush
			} else {
				r[0] = straight
			}
		} else if len(suits) == 1 {
			r[0] = flush
		} else {
			r[0] = highCard
		}
	}

	for i, c := range cards {
		r[i+1] = c.rank
	}

	return r, nil
}

// compareRanking compares to ranking types
func compareRanking(r1, r2 ranking) int {
	for i := range r1 {
		if r1[i] != r2[i] {
			return r1[i] - r2[i]
		}
	}

	return 0
}

// BestHand finds the best poker hand in a slice of strings representing the hands
func BestHand(hands []string) ([]string, error) {
	rank := ranking{highCard, 0, 0, 0, 0, 0}
	best := []string{}

	for _, h := range hands {
		r, err := rankHand(h, aceHigh)
		if err != nil {
			return best, err
		}

		if (r[0] == highCard || r[0] == flush) && r[1] == aceHigh {
			r2, err := rankHand(h, aceLow)
			if err != nil {
				return best, err
			}

			if compareRanking(r2, rank) > 0 {
				r = r2
			}
		}

		if c := compareRanking(r, rank); c > 0 {
			best = []string{h}
			rank = r
		} else if c == 0 {
			best = append(best, h)
		}
	}

	return best, nil
}
