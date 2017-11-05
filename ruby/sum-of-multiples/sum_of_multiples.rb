class SumOfMultiples
  def initialize(*factors)
    @factors = factors
  end

  def to(limit)
    (1...limit).select { |i| @factors.any? { |f| (i % f).zero? } }.sum
  end
end

module BookKeeping
  VERSION = 2
end
