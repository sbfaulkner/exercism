class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    @phrase.downcase.scan(/\b[\w']+\b/).each_with_object(Hash.new(0)) do |word, words|
      words[word] += 1
    end
  end
end

module BookKeeping
  VERSION = 1
end
