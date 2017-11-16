class Robot
  ALL_NAMES = ('AA000'..'ZZ999').to_a

  class << self
    attr_reader :names

    def forget
      @names = ALL_NAMES.shuffle(random: Random.new)
    end
  end

  forget

  def initialize
    reset
  end

  attr_reader :name

  def reset
    @name = self.class.names.shift
  end
end

module BookKeeping
  VERSION = 3
end
