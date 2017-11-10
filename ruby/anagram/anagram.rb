class Anagram
  def initialize(word)
    @word = word.downcase
    @letters = @word.chars.sort
  end

  def match(words)
    words.select do |candidate|
      word = candidate.downcase
      word != @word && word.chars.sort == @letters
    end
  end
end

module BookKeeping
  VERSION = 2
end
