package dna

import (
	"fmt"
	"strings"
)

const testVersion = 2

const nucleotides = "ACGT"

// DNA represents a DNA sequence
type DNA string

// Histogram holds the frequency count of the nucleotides in a DNA sequence
type Histogram map[byte]int

// Count determines the frequency of a nucleotide within a DNA sequence
func (d DNA) Count(nucleotide byte) (int, error) {
	if !strings.ContainsRune(nucleotides, rune(nucleotide)) {
		return 0, fmt.Errorf("dna: invalid nucleotide - %q", nucleotide)
	}

	count := 0

	for _, n := range []byte(d) {
		if n == nucleotide {
			count++
		}
	}

	return count, nil
}

// Counts determines the frequency of all nucleotide within a DNA sequence
func (d DNA) Counts() (Histogram, error) {
	counts := Histogram{}

	for _, n := range []byte(nucleotides) {
		counts[n] = 0
	}

	for _, n := range []byte(d) {
		c, err := d.Count(n)

		if err != nil {
			return nil, err
		}

		counts[n] = c
	}

	return counts, nil
}
