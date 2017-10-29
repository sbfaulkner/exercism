# Raindrops converts a number to raindrop-speak.
module Raindrops
  FACTORS = {
    3 => 'Pling',
    5 => 'Plang',
    7 => 'Plong',
  }.freeze

  def self.convert(number)
    text = FACTORS.select { |f, _| (number % f).zero? }.values
    text.empty? ? number.to_s : text.join
  end
end

module BookKeeping
  VERSION = 3
end
