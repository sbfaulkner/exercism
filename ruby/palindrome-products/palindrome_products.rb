# frozen_string_literal: true

require 'forwardable'

class Palindromes
  class Product
    attr_reader :factors, :value

    def initialize(value, *factors)
      @factors = factors
      @value = value
    end

    def palindrome?
      @value.to_s == @value.to_s.reverse
    end

    def <=>(other)
      @value <=> other.value
    end
  end

  def initialize(min_factor: 1, max_factor:)
    @factors = (min_factor..max_factor).to_a.repeated_combination(2)
  end

  extend Forwardable
  def_delegator :@products, :max, :largest
  def_delegator :@products, :min, :smallest

  def generate
    @products ||= @factors.group_by { |a, b| a * b }.map { |c, factors| Product.new(c, *factors) }.select(&:palindrome?)
  end
end
