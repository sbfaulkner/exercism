# Transcribe a DNA strand into it's complement.
module Complement
  DNA = 'GCTA'.freeze
  RNA = 'CGAU'.freeze

  def self.of_dna(strand)
    return '' unless strand.delete(DNA).empty?
    strand.tr(DNA, RNA)
  end
end

module BookKeeping
  VERSION = 4
end
