package strand

const testVersion = 3

var translations = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts a sequence of dna to its rna complement.
func ToRNA(dna string) string {
	sequence := []rune(dna)

	for i, d := range sequence {
		sequence[i] = translations[d]
	}

	return string(sequence)
}
