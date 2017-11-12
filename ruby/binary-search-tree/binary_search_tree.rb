class Bst
  attr_reader :data

  def initialize(data)
    @data = data
  end

  %w(left right).each do |subtree|
    class_eval <<-SUBTREE_METHODS
      def #{subtree}
        @#{subtree}.dup.freeze
      end

      def insert_#{subtree}(data)
        @#{subtree} = @#{subtree} ? @#{subtree}.insert(data) : self.class.new(data)
      end
    SUBTREE_METHODS
  end

  def insert(data)
    if data <= @data
      insert_left(data)
    else
      insert_right(data)
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
