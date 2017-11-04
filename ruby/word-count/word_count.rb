class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    @phrase.split(/[^\w']+/).each_with_object(Hash.new(0)) do |word, words|
      words[word.sub(/\A'(.+)'\z/, '\1').downcase] += 1
    end
  end
end

module BookKeeping
  VERSION = 1
end
