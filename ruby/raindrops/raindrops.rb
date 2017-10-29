# Raindrops converts a number to raindrop-speak.
module Raindrops
  FACTORS = [
    [3, 'Pling'],
    [5, 'Plang'],
    [7, 'Plong'],
  ]

  def self.convert(number)
    text = FACTORS.map { |f, w| w if (number % f).zero? }.compact.join
    return number.to_s if text.empty?
    text
  end
end

module BookKeeping
  VERSION = 3
end
