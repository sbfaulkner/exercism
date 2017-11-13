require 'forwardable'

class Nucleotide
  extend Forwardable

  NUCLEOTIDES = %w(A C G T)

  def self.from_dna(strand)
    new(strand.chars)
  end

  def_delegator :@nucleotides, :count

  def initialize(nucleotides)
    raise ArgumentError if (nucleotides - NUCLEOTIDES).any?
    @nucleotides = nucleotides
  end

  def histogram
    NUCLEOTIDES.each_with_object({}) { |n, h| h[n] = count(n) }
  end
end
