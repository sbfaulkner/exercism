class Matrix
  attr_reader :rows, :columns

  def initialize(text)
    @rows = text.split("\n").map { |row| row.split.map(&:to_i) }.freeze
    @columns = Array.new(@rows[0].size) { |c| @rows.map { |row| row[c] } }.freeze
  end
end
