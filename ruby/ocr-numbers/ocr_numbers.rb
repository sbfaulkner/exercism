class OcrNumbers
  NUMBERS = [
    ' _     _  _     _  _  _  _  _ ',
    '| |  | _| _||_||_ |_   ||_||_|',
    '|_|  ||_  _|  | _||_|  ||_| _|',
    '                              '
  ]

  def self.convert(text)
    new(text).convert
  end

  def initialize(text)
    @lines = text.lines.map(&:chomp).map do |line|
      raise ArgumentError unless (line.length % 3).zero?
      line.scan(/.{3}/)
    end

    raise ArgumentError unless (@lines.length % 4).zero?
  end

  def convert
    @lines.each_slice(4).map { |slice| convert_number(slice) }.join(',')
  end

  private

  def convert_number(lines)
    lines[0].zip(*lines[1..3]).map { |parts| patterns.dig(*parts) || '?' }.join
  end

  def patterns
    @patterns ||= 10.times.each_with_object({}) do |i, patterns|
      NUMBERS.map { |line| line[i * 3, 3] }.each_with_index do |pattern, p|
        if p < 3
          patterns = patterns[pattern] ||= {}
        else
          patterns[pattern] = i
        end
      end
    end
  end
end

module BookKeeping
  VERSION = 1
end
