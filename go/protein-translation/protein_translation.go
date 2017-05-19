package protein

const testVersion = 1

type polypeptideTranslation struct {
	codons      []string
	polypeptide string
}

var polypeptideTranslations = []polypeptideTranslation{
	{[]string{"AUG"}, "Methionine"},
	{[]string{"UUU", "UUC"}, "Phenylalanine"},
	{[]string{"UUA", "UUG"}, "Leucine"},
	{[]string{"UCU", "UCC", "UCA", "UCG"}, "Serine"},
	{[]string{"UAU", "UAC"}, "Tyrosine"},
	{[]string{"UGU", "UGC"}, "Cysteine"},
	{[]string{"UGG"}, "Tryptophan"},
	{[]string{"UAA", "UAG", "UGA"}, "STOP"},
}

// FromCodon converts a codon to a polypeptide.
func FromCodon(sequence string) string {
	for _, t := range polypeptideTranslations {
		for _, codon := range t.codons {
			if sequence == codon {
				return t.polypeptide
			}
		}
	}

	return ""
}

// FromRNA translates an RNA sequence into a protein.
func FromRNA(sequence string) (protein []string) {
	for i := 0; i < len(sequence); i += 3 {
		polypeptide := FromCodon(sequence[i : i+3])
		if polypeptide == "STOP" {
			break
		}
		protein = append(protein, polypeptide)
	}

	return
}
