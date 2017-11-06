class Series
  def initialize(digits)
    @digits = digits.chars
  end

  def slices(length)
    raise ArgumentError if length > @digits.length
    @digits.each_cons(length).map(&:join)
  end
end
