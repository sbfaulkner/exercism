class Queens
  BOARD_SIZE = 8

  WHITE_POSITION = [0, 3]
  BLACK_POSITION = [7, 3]

  class Queen
    attr_reader :row, :column

    def initialize(row, column)
      raise ArgumentError unless row.between?(0, BOARD_SIZE - 1)
      raise ArgumentError unless column.between?(0, BOARD_SIZE - 1)

      @row = row
      @column = column
    end

    def intersecting_row?(other)
      row == other.row
    end

    def intersecting_column?(other)
      column == other.column
    end

    def intersecting_diagonal?(other)
      (column - other.column).abs == (row - other.row).abs
    end
  end

  def initialize(white: WHITE_POSITION, black: BLACK_POSITION)
    @white = Queen.new(*white)
    @black = Queen.new(*black)
  end

  def attack?
    @white.intersecting_row?(@black) || @white.intersecting_column?(@black) || @white.intersecting_diagonal?(@black)
  end
end

module BookKeeping
  VERSION = 2
end
