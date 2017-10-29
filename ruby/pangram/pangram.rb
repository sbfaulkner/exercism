require 'set'

# Pangram detects sentences that contain every letter in the alphabet at least once
module Pangram
  ALPHABET = ('a'..'z').to_set

  def self.pangram?(sentence)
    (ALPHABET - sentence.downcase.chars.to_set).empty?
  end
end

module BookKeeping
  VERSION = 6
end
