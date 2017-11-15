class Say
  def initialize(number)
    @number = number
  end

  MAXIMUM_NUMBER = 999_999_999_999

  ZERO = 'zero'.freeze
  ONES = %w[zero one two three four five six seven eight nine].freeze
  TEENS = %w[ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen].freeze
  TENS = %w[zero ten twenty thirty forty fifty sixty seventy eighty ninety].freeze
  SCALES = %w[zero thousand million billion].freeze

  def in_english
    raise ArgumentError if @number.negative? || @number > MAXIMUM_NUMBER

    return ZERO if @number.zero?

    words_for(@number).join(' ')
  end

  def words_for(number, scale = 0)
    words = []

    thousands, ones = number.divmod(1000)

    words.unshift(SCALES[scale]) if scale.positive? && ones.positive?

    hundreds, ones = ones.divmod(100)
    tens, ones = ones.divmod(10)

    if ones.positive?
      if tens == 1
        words.unshift(TEENS[ones])
      elsif tens.positive?
        words.unshift("#{TENS[tens]}-#{ONES[ones]}")
      else
        words.unshift(ONES[ones])
      end
    elsif tens.positive?
      words.unshift(TENS[tens])
    end

    words.unshift("#{ONES[hundreds]} hundred") if hundreds.positive?

    words.unshift(*words_for(thousands, scale + 1)) if thousands.positive?

    words
  end
end

module BookKeeping
  VERSION = 1
end
