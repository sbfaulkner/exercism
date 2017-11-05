module Accumulate
  def accumulate
    accumulator = []
    each { |item| accumulator << yield(item) }
    accumulator
  end
end

Array.include Accumulate

module BookKeeping
  VERSION = 1
end
