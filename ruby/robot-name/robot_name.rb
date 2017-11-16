class Robot
  class NameGenerator
    ALL_NAMES = 'AA000'..'ZZ999'

    def initialize
      @prng = Random.new
      @names = ALL_NAMES.to_a
    end

    def next
      @names.delete_at(@prng.rand(@names.size))
    end
  end

  class << self
    attr_reader :name_generator

    def forget
      @name_generator = NameGenerator.new
    end
  end

  forget

  def initialize
    reset
  end

  attr_reader :name

  def reset
    @name = self.class.name_generator.next
  end
end

module BookKeeping
  VERSION = 3
end
