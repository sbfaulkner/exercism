# frozen_string_literal: true

module Luhn
  def self.valid?(text)
    text = text.delete(' ')

    return false if text !~ /\A[0-9]{2,}\z/

    digits = text.reverse.chars.map.with_index do |ch, i|
      digit = ch.to_i

      if i.odd?
        digit *= 2
        digit -= 9 if digit > 9
      end

      digit
    end

    (digits.sum % 10).zero?
  end
end

module BookKeeping
  VERSION = 1
end
