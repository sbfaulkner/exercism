class Crypto
  def initialize(plaintext)
    @normalized = plaintext.downcase.tr('^a-z0-9', '').chars
  end

  def ciphertext
    columns = Math.sqrt(@normalized.size).ceil
    rows = (@normalized.size.to_f / columns).ceil if columns.positive?

    square = Array.new(columns) do |c|
      rows.times.each_with_object('') do |i, row|
        row << @normalized.fetch(c + i * columns, ' ')
      end
    end

    square.join(' ')
  end
end

module BookKeeping
  VERSION = 1
end
