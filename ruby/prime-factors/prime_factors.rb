module PrimeFactors
  def self.for(number)
    factors = []
    factor = 2

    while number > 1
      if (number % factor).zero?
        factors << factor
        number /= factor
      elsif factor > 2
        factor += 2
      else
        factor += 1
      end
    end

    factors
  end
end
