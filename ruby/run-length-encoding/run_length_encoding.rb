# RunLengthEncoding implements run-length encoding and decoding.
module RunLengthEncoding
  class << self
    def encode(decoded)
      decoded.scan(/(([A-Za-z ])\2*)/).map do |chunk, char|
        count = chunk.size
        count == 1 ? char : "#{count}#{char}"
      end.join
    end

    def decode(encoded)
      encoded.scan(/[0-9]*[A-Za-z ]/).map do |chunk|
        count = chunk.to_i
        count.zero? ? chunk : (chunk[-1] * count)
      end.join
    end
  end
end

module BookKeeping
  VERSION = 3
end
