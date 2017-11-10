class Bst
  attr_reader :data, :left, :right

  def self.insert(data, bst)
    bst ? bst.insert(data) : new(data)
  end

  def initialize(data)
    @data = data
  end

  def insert(data)
    if data <= @data
      @left = self.class.insert(data, @left)
    else
      @right = self.class.insert(data, @right)
    end

    self
  end

  def each(&block)
    return enum_for(:each) unless block_given?

    @left&.each(&block)
    yield data
    @right&.each(&block)

    self
  end
end

module BookKeeping
  VERSION = 1
end
