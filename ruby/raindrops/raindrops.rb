# Raindrops converts a number to raindrop-speak.
module Raindrops
  def self.convert(number)
    raindrop = ''

    raindrop << 'Pling' if (number % 3).zero?
    raindrop << 'Plang' if (number % 5).zero?
    raindrop << 'Plong' if (number % 7).zero?

    raindrop = number.to_s if raindrop.empty?

    raindrop
  end
end

module BookKeeping
  VERSION = 3
end
