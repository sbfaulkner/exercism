class Trinary
  def initialize(trinary)
    @trinary = trinary
  end

  def to_decimal
    return 0 unless @trinary =~ /\A[0-2]+\z/
    @trinary.chars.reduce(0) { |decimal, digit| decimal * 3 + digit.to_i }
  end
end

module BookKeeping
  VERSION = 1
end
