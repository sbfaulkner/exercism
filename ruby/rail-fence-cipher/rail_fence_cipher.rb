module RailFenceCipher
  VERSION = 1

  class << self
    def encode(decoded, count)
      return decoded if count == 1

      rails = make_rails(decoded.length, count) { |i| decoded[i] }
      rails.join
    end

    def decode(encoded, count)
      return encoded if count == 1

      rails = make_rails(encoded.length, count) { |i| i }
      rails.reduce(&:concat).zip(encoded.chars).sort.map(&:last).join
    end

    private

    def make_rails(length, count)
      rails = Array.new(count) { [] }

      direction = 1
      rail = 1

      length.times do |i|
        rails[rail - 1] << yield(i)
        rail += direction
        direction *= -1 if rail == 1 || rail == count
      end

      rails
    end
  end
end
