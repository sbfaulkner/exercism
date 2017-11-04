# Year defines method to test for a leap year
module Year
  def self.leap?(year)
    (year % 4).zero? && !(year % 100).zero? || (year % 400).zero?
  end
end

module BookKeeping
  VERSION = 3
end
