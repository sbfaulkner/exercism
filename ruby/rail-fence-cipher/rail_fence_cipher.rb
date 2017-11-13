module RailFenceCipher
  VERSION = 1

  class << self
    def encode(decoded, count)
      return decoded if count == 1

      rails = Array.new(count) { '' }
      direction = 1
      rail = 1

      decoded.chars.each do |char|
        rails[rail - 1] << char
        rail += direction
        direction *= -1 if rail == 1 || rail == count
      end

      rails.join
    end

    def decode(encoded, count)
      return encoded if count == 1

      lengths = Array.new(count) { 0 }
      direction = 1
      rail = 1

      encoded.size.times do
        lengths[rail - 1] += 1
        rail += direction
        direction *= -1 if rail == 1 || rail == count
      end

      offset = 0
      rails = lengths.map do |length|
        part = encoded[offset, length].chars
        offset += length
        part
      end

      direction = 1
      rail = 1
      decoded = []

      encoded.size.times do
        decoded << rails[rail - 1].shift
        rail += direction
        direction *= -1 if rail == 1 || rail == count
      end

      decoded.join
    end
  end
end
