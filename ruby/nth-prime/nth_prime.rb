# Prime provides method to find the nth prime
module Prime
  @primes = [2, 3]

  def self.nth(n)
    raise ArgumentError if n < 1

    while n > @primes.size
      i = @primes.last + 2
      i += 2 while @primes.any? { |p| (i % p).zero? }
      @primes << i
    end

    @primes[n - 1]
  end
end

module BookKeeping
  VERSION = 1
end
