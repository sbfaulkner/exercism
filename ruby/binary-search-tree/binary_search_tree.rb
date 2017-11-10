class Bst
  attr_reader :data, :left, :right

  def initialize(data)
    @data = data
  end

  def insert(data)
    if data <= @data
      if @left
        @left.insert(data)
      else
        @left = self.class.new(data)
      end
    else
      if @right
        @right.insert(data)
      else
        @right = self.class.new(data)
      end
    end
  end

  def each(&block)
    return enum_for(:each) unless block_given?

    @left&.each(&block)
    yield data
    @right&.each(&block)
  end
end

module BookKeeping
  VERSION = 1
end
