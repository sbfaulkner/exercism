package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

var ordinals = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

func giftList(day int) (list string) {
	if day > 1 {
		list = strings.Join(gifts[12-day:11], ", ")
		list += ", and "
	}

	list += gifts[11]

	return
}

// Song returns the lyrics for the entire Twelve Days of Christmas song
func Song() string {
	verses := make([]string, 13)

	for day := 1; day < 13; day++ {
		verses[day-1] = Verse(day)
	}

	return strings.Join(verses, "\n")
}

// Verse returns the lyrics for a single verse of the Twelve Days of Christmas song
func Verse(day int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", ordinals[day-1], giftList(day))
}
