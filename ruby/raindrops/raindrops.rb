# Raindrops converts a number to raindrop-speak.
module Raindrops
  FACTORS = [
    [3, 'Pling'],
    [5, 'Plang'],
    [7, 'Plong'],
  ].freeze

  def self.convert(number)
    text = FACTORS.select { |f, _| (number % f).zero? }.map(&:last).join
    return number.to_s if text.empty?
    text
  end
end

module BookKeeping
  VERSION = 3
end
