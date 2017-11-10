require 'forwardable'

class Game
  BowlingError = Class.new(ArgumentError)

  class Frame
    extend Forwardable

    def_delegators :@rolls, :sum, :size

    def initialize
      @rolls = []
    end

    def <<(pins)
      raise BowlingError unless pins.between?(0, 10)
      @rolls << pins
      raise BowlingError unless valid?
    end

    def bonus_required?
      (strike? || spare?) && size < 3
    end

    def complete?
      strike? || @rolls[1]
    end

    def strike?
      @rolls[0] == 10
    end

    def spare?
      @rolls[0, 2].sum == 10
    end

    def valid?
      if strike?
        @rolls[1, 2].sum <= 10 || @rolls[1] == 10
      else
        @rolls[0, 2].sum <= 10
      end
    end
  end

  def initialize
    @frames = [Frame.new]
  end

  def roll(pins)
    raise BowlingError if complete?

    @frames[-1] << pins
    @frames[-2] << pins if frame > 1 && @frames[-2].bonus_required?
    @frames[-3] << pins if frame > 2 && @frames[-3].bonus_required?

    next_frame if @frames[-1].complete?
  end

  def score
    raise BowlingError unless complete?
    @frames.sum(&:sum)
  end

  private

  def complete?
    frame == 10 && @frames[-1].complete? && !@frames[-1].bonus_required?
  end

  def frame
    @frames.size
  end

  def next_frame
    return if frame == 10
    @frames << Frame.new
  end
end

module BookKeeping
  VERSION = 3
end
