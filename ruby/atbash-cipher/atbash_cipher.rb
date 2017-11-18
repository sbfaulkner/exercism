module Atbash
  ALPHABET = ('a'..'z').to_a.join.freeze
  REVERSED = ALPHABET.reverse.freeze
  NOISE = '^a-z0-9'.freeze
  GROUP = /.{1,5}/

  def self.encode(decoded)
    decoded.downcase.delete(NOISE).tr(ALPHABET, REVERSED).scan(GROUP).join(' ')
  end
end
