class Robot
  class NameGenerator
    ALL_NAMES = ('AA000'..'ZZ999').to_a

    def initialize
      @names = ALL_NAMES.shuffle(random: Random.new)
    end

    def next
      @names.shift
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
