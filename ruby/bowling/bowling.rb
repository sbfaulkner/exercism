require 'forwardable'

class Game
  BowlingError = Class.new(ArgumentError)

  class Frame
    extend Forwardable

    def_delegators :@rolls, :sum, :size

    def initialize
      @rolls = []
    end

    def inspect
      "<Frame @rolls=#{@rolls.inspect}>"
    end

    def <<(pins)
      @rolls << pins
    end

    def complete?
      strike? || @rolls[1]
    end

    def strike?
      @rolls[0] == 10
    end

    def spare?
      @rolls[0, 2].sum == 10 && !strike?
    end
  end

  def initialize
    @frames = [Frame.new]
  end

  def roll(pins)
    raise BowlingError unless pins.between?(0, 10)

    @frames[-1] << pins
    @frames[-2] << pins if @frames[-2] && (@frames[-2].strike? && @frames[-1].size < 3 || @frames[-2].spare? && @frames[-1].size == 1)
    @frames[-3] << pins if @frames[-3] && @frames[-3].strike? && @frames[-1].size == 1

    next_frame if @frames[-1].complete? && !complete?
  end

  def score
    @frames.sum(&:sum)
  end

  private

  def complete?
    @frames.size == 10
  end

  def next_frame
    @frames << Frame.new
  end
end
