# Hamming module to determine distance between two DNA strands
module Hamming
  def self.compute(s1, s2)
    raise ArgumentError, 'strand lengths not equal' unless s1.length == s2.length

    s1.chars.zip(s2.chars).count do |n1, n2|
      n1 != n2
    end
  end
end

module BookKeeping
  VERSION = 3
end
