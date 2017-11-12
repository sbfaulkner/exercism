class Clock
  def self.at(hour, minute)
    Clock.new(hour, minute)
  end

  attr_reader :hour, :minute

  def initialize(hour, minute)
    @hour = (hour + minute / 60) % 24
    @minute = minute % 60
  end

  def to_s
    format('%02d:%02d', hour, minute)
  end

  def +(minutes)
    Clock.new(@hour, @minute + minutes)
  end

  def ==(clock)
    clock.hour == hour && clock.minute == minute
  rescue
    false
  end
end

module BookKeeping
  VERSION = 2
end
