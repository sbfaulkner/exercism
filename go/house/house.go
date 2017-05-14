package house

import "fmt"
import "strings"

const testVersion = 1

const numberOfVerses = 12

var details = []struct {
	subject string
	action  string
}{
	{"house", "Jack built"},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

// Song generates the text of The House That Jack Built
func Song() string {
	verses := make([]string, numberOfVerses)

	for i := 0; i < len(verses); i++ {
		verses[i] = Verse(i + 1)
	}

	return strings.Join(verses, "\n\n")
}

// Verse generates the specified verse of The House That Jack Built
func Verse(number int) string {
	if number == 1 {
		return fmt.Sprintf("This is the %s that %s.", details[0].subject, details[0].action)
	}

	number--
	verse := fmt.Sprintf("the %s\nthat %s the", details[number].subject, details[number].action)
	verse = strings.Replace(Verse(number), "the", verse, 1)

	return verse
}
