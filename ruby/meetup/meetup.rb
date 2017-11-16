# frozen_string_literal: true

require 'date'

class Meetup
  OCCURENCES = %i[first second third fourth].freeze
  WEEKDAYS = %i[sunday monday tuesday wednesday thursday friday saturday].freeze

  def initialize(month, year)
    @month = month
    @year = year
  end

  def day(weekday, occurence)
    case occurence
    when :teenth
      date = Date.new(@year, @month, 13)
      date + days_until_weekday(date, weekday)
    when :last
      date = Date.new(@year + @month / 12, @month % 12 + 1, 1) - 7
      date + days_until_weekday(date, weekday)
    else
      date = Date.new(@year, @month, 1)
      date + days_until_weekday(date, weekday) + OCCURENCES.index(occurence) * 7
    end
  end

  def days_until_weekday(date, weekday)
    (7 + WEEKDAYS.index(weekday) - date.cwday) % 7
  end
end

module BookKeeping
  VERSION = 1
end
