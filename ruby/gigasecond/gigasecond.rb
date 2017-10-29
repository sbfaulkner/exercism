# Gigasecond module to calculate a date 10^9 seconds into the future
module Gigasecond
  def self.from(dob)
    dob + 10**9
  end
end

module BookKeeping
  VERSION = 6
end
