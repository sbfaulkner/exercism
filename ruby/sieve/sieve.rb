require 'set'

# Sieve class implements the Sieve of Eratosthenes.
class Sieve
  def initialize(limit)
    @primes = []

    candidates = (2..limit).to_set

    while c = candidates.first
      @primes << c

      i = c

      while i <= limit
        candidates.delete(i)

        i += c
      end
    end
  end

  attr_reader :primes
end

module BookKeeping
  VERSION = 1
end
