# RomanNumerals module implementing conversion to roman numerals.
module RomanNumerals
  TENS = 'IXCM'.freeze
  FIVES = 'VLD'.freeze

  def to_roman
    decimal = self

    FIVES.length.downto(0).each_with_object('') do |i, roman|
      ten = 10**i

      case count = decimal / ten
      when 1, 2, 3
        roman << TENS[i] * count
      when 4
        roman << TENS[i] << FIVES[i]
      when 5
        roman << FIVES[i]
      when 6, 7, 8
        roman << FIVES[i] << TENS[i] * (count - 5)
      when 9
        roman << TENS[i, 2]
      end

      decimal %= ten
    end
  end
end

Integer.include RomanNumerals

module BookKeeping
  VERSION = 2
end
