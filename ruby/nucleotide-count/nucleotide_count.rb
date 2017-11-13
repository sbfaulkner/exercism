class Nucleotide
  def self.from_dna(strand)
    new(strand.chars)
  end

  NUCLEOTIDES = { 'A' => 0, 'C' => 0, 'G' => 0, 'T' => 0 }.freeze

  attr_reader :histogram

  def initialize(nucleotides)
    @histogram = nucleotides.each_with_object(NUCLEOTIDES.dup) do |nucleotide, histogram|
      histogram[nucleotide] = histogram.fetch(nucleotide) { raise ArgumentError } + 1
    end
  end

  def count(nucleotide)
    histogram[nucleotide]
  end
end
