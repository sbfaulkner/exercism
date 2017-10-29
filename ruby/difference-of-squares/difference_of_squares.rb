# Squares defines methods to calculate the Difference Of Squares
class Squares
  def initialize(number)
    @range = 1..number
  end

  # square_of_sum calculates the square of the sum the terms in the series
  def square_of_sum
    @range.sum**2
  end

  # square_of_sum calculates the sum of the squares of the terms in the series
  def sum_of_squares
    @range.map { |n| n**2 }.sum
  end

  # difference calculates the difference between
  # the square of the sum of the terms and the sum of the squares of the terms
  def difference
    square_of_sum - sum_of_squares
  end
end

module BookKeeping
  VERSION = 4
end
