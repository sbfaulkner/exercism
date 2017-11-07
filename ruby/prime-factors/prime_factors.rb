require 'prime'

module PrimeFactors
  def self.for(number)
    Prime.prime_division(number).map { |p, e| [p] * e }.flatten
  end
end
