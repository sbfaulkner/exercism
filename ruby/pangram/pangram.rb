module Pangram
  ALPHABET = 'a'..'z'

  def self.pangram?(sentence)
    ALPHABET.all? { |letter| sentence =~ /#{letter}/i }
  end
end

module BookKeeping
  VERSION = 6
end
