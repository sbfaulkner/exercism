class Matrix
  attr_reader :rows, :columns

  def initialize(text)
    @matrix = text.each_line.map { |row| row.split.map(&:to_i) }
    @rows = @matrix.dup
    @columns = @matrix.transpose
  end

  def saddle_points
    saddle_points = []

    rows.each_with_index do |row, r|
      row.each_with_index do |value, c|
        saddle_points << [r, c] if value == row.max && value == columns[c].min
      end
    end

    saddle_points
  end
end
