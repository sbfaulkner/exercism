package allergies

const testVersion = 1

var allergyItems = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies lists the items that the allergy score indicates allergies to.
func Allergies(score uint) []string {
	allergies := []string{}

	for i, allergy := range allergyItems {
		if isScoreBit(score, i) {
			allergies = append(allergies, allergy)
		}
	}

	return allergies
}

// AllergicTo returns whether or not the provided allergy score indicates an allergy to the specified item.
func AllergicTo(score uint, item string) bool {
	for i, allergy := range allergyItems {
		if item == allergy {
			return isScoreBit(score, i)
		}
	}
	return false
}

func isScoreBit(score uint, b int) bool {
	return score&(1<<uint(b)) != 0
}
