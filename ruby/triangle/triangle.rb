class Triangle
  def initialize(sides)
    @sides = sides
  end

  def equilateral?
    valid? && @sides.uniq.size == 1
  end

  def isosceles?
    valid? && @sides.uniq.size <= 2
  end

  def scalene?
    valid? && @sides.uniq.size == 3
  end

  private

  def valid?
    @sides.map.with_index { |side, index| side.positive? && triangle_inequality?(index) }.all?
  end

  def triangle_inequality?(index)
    @sides[index] < @sides.map.with_index { |s, i| i == index ? 0 : s }.sum
  end
end

module BookKeeping
  VERSION = 1
end
