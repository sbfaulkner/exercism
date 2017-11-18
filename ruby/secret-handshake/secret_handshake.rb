# frozen_string_literal: true

class SecretHandshake
  COMMANDS = [
    'wink',
    'double blink',
    'close your eyes',
    'jump',
  ].freeze

  REVERSE = 0b10000

  def initialize(number)
    @number = number.to_i
  end

  def commands
    handshake = COMMANDS.reject.with_index { |_, i| (@number & 2**i).zero? }
    (@number & REVERSE).zero? ? handshake : handshake.reverse
  end
end
