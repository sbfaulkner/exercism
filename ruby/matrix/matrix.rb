class Matrix
  attr_reader :rows, :columns

  def initialize(text)
    @rows = text.lines.map { |row| row.split.map(&:to_i) }.freeze
    @columns = @rows.transpose.freeze
  end
end
