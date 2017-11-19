# frozen_string_literal: true

require 'securerandom'

class Cipher
  attr_reader :key

  ALPHABET = [*'a'..'z']
  NONALPHA = /[^a-z]/
  A = 'a'.ord

  def initialize(key = random_key)
    raise ArgumentError if key.empty? || key.match?(NONALPHA)
    @key = key
  end

  def encode(plaintext)
    shift_text(plaintext, 1)
  end

  def decode(ciphertext)
    shift_text(ciphertext, -1)
  end

  private

  def random_key
    Array.new(128) { ALPHABET[SecureRandom.random_number(26)] }.join
  end

  def shift_char(ch, shift)
    ((ch.ord - A + shift + 26) % 26 + A).chr
  end

  def shift_text(text, direction)
    text.each_char.zip(key.each_char.cycle).map { |ch, k| shift_char(ch, direction * (k.ord - A)) }.map(&:chr).join
  end
end
