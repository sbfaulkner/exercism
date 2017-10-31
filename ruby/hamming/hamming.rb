# Hamming module to determine distance between two DNA strands
module Hamming
  def self.compute(strand1, strand2)
    raise ArgumentError, 'strand lengths not equal' unless strand1.length == strand2.length

    strand1.chars.zip(strand2.chars).count do |nucleotide1, nucleotide2|
      nucleotide1 != nucleotide2
    end
  end
end

module BookKeeping
  VERSION = 3
end
