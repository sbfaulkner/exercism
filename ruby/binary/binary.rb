module Binary
  def self.to_decimal(binary)
    raise ArgumentError unless binary =~ /\A[01]+\z/
    binary.chars.inject(0) { |value, char| value * 2 + char.to_i }
  end
end

module BookKeeping
  VERSION = 3
end
