module Atbash
  ALPHABET = ('a'..'z').to_a.join.freeze
  REVERSED = ALPHABET.reverse.freeze
  NOISE = '^a-z0-9'.freeze
  GROUP_SIZE = 5

  def self.encode(decoded)
    decoded.downcase.tr(NOISE, '').tr(ALPHABET, REVERSED).chars.each_slice(GROUP_SIZE).map(&:join).join(' ')
  end
end