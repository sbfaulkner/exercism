class Matrix
  attr_reader :rows, :columns

  def initialize(text)
    @matrix = text.each_line.map { |row| row.split.map(&:to_i) }
    @rows = @matrix.dup.freeze
    @columns = @matrix.transpose.freeze
  end

  def saddle_points
    cells.select { |r, c| saddle_point?(r, c) }
  end

  private

  def cells
    (0...rows.count).to_a.product((0...columns.count).to_a)
  end

  def max_in_row?(r, c)
    rows[r][c] == rows[r].max
  end

  def min_in_column?(r, c)
    columns[c][r] == columns[c].min
  end

  def saddle_point?(r, c)
    max_in_row?(r, c) && min_in_column?(r, c)
  end
end
