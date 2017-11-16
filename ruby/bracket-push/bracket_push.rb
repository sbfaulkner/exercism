# frozen_string_literal: true

require 'strscan'

module Brackets
  OPEN = '[({'
  CLOSE = '])}'

  def self.paired?(text)
    brackets = []

    text.chars.each do |char|
      if OPEN.include?(char)
        brackets.push(char)
      elsif CLOSE.include?(char)
        return false if brackets.pop != char.tr(CLOSE, OPEN)
      end
    end

    brackets.empty?
  end
end

module BookKeeping
  VERSION = 4
end
