module Acronym
  class << self
    def abbreviate(name)
      name.split(/\W/).map { |word| word[0] }.join.upcase
    end
  end
end

module BookKeeping
  VERSION = 4
end
