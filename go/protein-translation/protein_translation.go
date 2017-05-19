package protein

const testVersion = 1

var polypeptideTranslations = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

// FromCodon converts a codon to a polypeptide.
func FromCodon(sequence string) string {
	return polypeptideTranslations[sequence]
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
