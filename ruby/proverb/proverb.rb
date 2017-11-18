# frozen_string_literal: true

class Proverb
  LINE = 'For want of a %<wanted>s the %<lost>s was lost.'
  FINAL_LINE = 'And all for the want of a %<wanted>s.'

  def initialize(*things, qualifier: nil)
    @things = things
    @wanted = qualifier ? "#{qualifier} #{things.first}" : things.first
  end

  def to_s
    lines = @things.each_cons(2).map { |wanted, lost| format(LINE, wanted: wanted, lost: lost) }
    lines << format(FINAL_LINE, wanted: @wanted)
    lines.join("\n")
  end
end
