class OcrNumbers
  PATTERNS = {
    ' _ ' => {
      '| |' => {
        '|_|' => {
          '   ' => 0,
        },
      },
      ' _|' => {
        '|_ ' => {
          '   ' => 2,
        },
        ' _|' => {
          '   ' => 3,
        },
      },
      '|_ ' => {
        ' _|' => {
          '   ' => 5,
        },
        '|_|' => {
          '   ' => 6,
        },
      },
      '  |' => {
        '  |' => {
          '   ' => 7,
        },
      },
      '|_|' => {
        '|_|' => {
          '   ' => 8,
        },
        ' _|' => {
          '   ' => 9,
        },
      }
    },
    '   ' => {
      '  |' => {
        '  |' => {
          '   ' => 1,
        },
      },
      '|_|' => {
        '  |' => {
          '   ' => 4,
        },
      }
    }
  }

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

  def convert_number(lines)
    lines[0].zip(*lines[1..3]).map { |parts| PATTERNS.dig(*parts) || '?' }.join
  end
end

module BookKeeping
  VERSION = 1
end
