class Hexadecimal
  DIGITS = '0123456789abcdef'.freeze

  def initialize(hexadecimal)
    @hexadecimal = hexadecimal.downcase
  end

  def to_decimal
    @hexadecimal.chars.reduce(0) do |decimal, digit|
      value = DIGITS.index(digit)
      return 0 unless value
      decimal * 16 + value
    end
  end
end
