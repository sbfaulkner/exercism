class Series
  def initialize(digits)
    @digits = digits
  end

  def slices(length)
    raise ArgumentError if length > @digits.length
    (0..(@digits.length - length)).map { |i| @digits[i,length] }
  end
end
