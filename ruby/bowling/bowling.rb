class Game
  BowlingError = Class.new(ArgumentError)

  def initialize
    @rolls = []
  end

  def roll(pins)
    raise BowlingError unless pins.between?(0, 10)
    @rolls << pins
  end

  def score
    total = 0
    r = 0

    10.times do
      if @rolls[r] == 10
        total += @rolls[r, 3].sum
        r += 1
      elsif @rolls[r, 2].sum == 10
        total += @rolls[r, 3].sum
        r += 2
      else
        total += @rolls[r, 2].sum
        r += 2
      end
    end

    total
  end
end
