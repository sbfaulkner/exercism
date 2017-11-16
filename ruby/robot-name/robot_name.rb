class Robot
  class << self
    ALL_NAMES = 'AA000'..'ZZ999'

    def forget
      @names = generate_names
    end

    def next_name
      @names ||= generate_names
      @names.delete_at(rand(@names.size))
    end

    private

    def generate_names
      ALL_NAMES.to_a
    end
  end

  def initialize
    reset
  end

  attr_reader :name

  def reset
    srand
    @name = self.class.next_name
  end
end

module BookKeeping
  VERSION = 3
end
