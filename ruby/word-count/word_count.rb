class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    @phrase.downcase.scan(/\b[\w']+\b/).group_by(&:to_s).transform_values(&:count)
  end
end

module BookKeeping
  VERSION = 1
end
