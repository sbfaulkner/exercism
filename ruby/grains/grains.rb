# Grains provides method to calculate the number of grains of wheat on the king's chessboard.
module Grains
  SQUARES = 1..64

  def self.square(number)
    raise ArgumentError unless SQUARES.include?(number)
    2**(number - 1)
  end

  def self.total
    SQUARES.sum { |i| square(i) }
  end
end

module BookKeeping
  VERSION = 1
end
