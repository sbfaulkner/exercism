module Alphametics
  class << self
    DIGITS = ('0'..'9').to_a

    EQUALS = ' == '
    PLUS = ' + '

    def solve(puzzle)
      letters = puzzle.tr('^A-Z', '').chars.uniq.sort.join

      left, right = puzzle.split(EQUALS)
      left = left.split(PLUS)

      DIGITS.permutation(letters.size).each do |digits|
        digits = digits.join

        actual = left.sum { |l| convert_letters_to_integer(l, letters, digits) || break }
        next unless actual

        expected = convert_letters_to_integer(right, letters, digits)
        next unless expected

        return Hash[letters.chars.zip(digits.chars.map(&:to_i))] if actual == expected
      end

      {}
    end

    private

    def convert_letters_to_integer(term, variables, digits)
      text = term.tr(variables, digits)
      return nil if text.size > 1 && text[0] == '0'
      text.to_i
    end
  end
end

module BookKeeping
  VERSION = 4
end
