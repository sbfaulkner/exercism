# frozen_string_literal: true

module House
  FIRST_LINE = 'This is the %<character>s'
  OTHER_LINE = 'that %<action>s the %<character>s'

  CHARACTERS = [
    'house that Jack built.',
    'malt',
    'rat',
    'cat',
    'dog',
    'cow with the crumpled horn',
    'maiden all forlorn',
    'man all tattered and torn',
    'priest all shaven and shorn',
    'rooster that crowed in the morn',
    'farmer sowing his corn',
    'horse and the hound and the horn',
  ]

  ACTIONS = [
    'lay in',
    'ate',
    'killed',
    'worried',
    'tossed',
    'milked',
    'kissed',
    'married',
    'woke',
    'kept',
    'belonged to',
  ]

  def self.recite
    Array.new(CHARACTERS.size) { |v| verse(v + 1) }.join("\n")
  end

  def self.verse(number)
    phrases = []

    phrases << format(FIRST_LINE, character: CHARACTERS[number - 1])

    (number - 2).downto(0) do |i|
      phrases << format(OTHER_LINE, action: ACTIONS[i], character: CHARACTERS[i])
    end

    phrases << ''

    phrases.join("\n")
  end
end
