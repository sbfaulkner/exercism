# RunLengthEncoding implements run-length encoding and decoding.
module RunLengthEncoding
  class << self
    def encode(decoded)
      decoded.gsub(/(([A-Za-z ])\2*)/) do |chunk|
        count = chunk.size
        count == 1 ? chunk : "#{count}#{chunk[0]}"
      end
    end

    def decode(encoded)
      encoded.gsub(/[0-9]*[A-Za-z ]/) do |chunk|
        count = chunk.to_i
        count.zero? ? chunk : (chunk[-1] * count)
      end
    end
  end
end

module BookKeeping
  VERSION = 3
end
