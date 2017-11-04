# Bob simulates a lackadaisical teenager.
module Bob
  # Remark recognizes different types of remarks.
  class Remark
    require 'forwardable'
    extend Forwardable

    def_delegator :@text, :empty?

    def initialize(text)
      @text = text.strip
    end

    def shouting?
      @text =~ /[a-z]/i && @text == @text.upcase
    end

    def question?
      @text.end_with?('?')
    end
  end

  def self.hey(remark_text)
    remark = Remark.new(remark_text)

    if remark.empty?
      'Fine. Be that way!'
    elsif remark.shouting?
      'Whoa, chill out!'
    elsif remark.question?
      'Sure.'
    else
      'Whatever.'
    end
  end
end

module BookKeeping
  VERSION = 1
end
