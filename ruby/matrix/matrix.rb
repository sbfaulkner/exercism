class Matrix
  def initialize(text)
    @matrix = text.each_line.map { |row| row.split.map(&:to_i) }
  end

  def rows
    @matrix.dup
  end

  def columns
    @matrix.transpose
  end
end
