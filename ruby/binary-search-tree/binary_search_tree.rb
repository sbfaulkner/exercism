class Bst
  attr_reader :data, :left, :right

  def initialize(data)
    @data = data
  end

  def insert(data)
    insert_or_new = lambda { |bst| bst ? bst.insert(data) : self.class.new(data) }

    if data <= @data
      @left = insert_or_new.call(@left)
    else
      @right = insert_or_new.call(@right)
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
