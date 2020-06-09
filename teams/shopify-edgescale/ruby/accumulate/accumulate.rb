module Accumulate
  def accumulate
    return to_enum unless block_given?

    accumulator = []
    each { |item| accumulator << yield(item) }
    accumulator
  end
end

Array.include Accumulate
