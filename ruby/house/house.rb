# frozen_string_literal: true

module House
  LEAD = 'This is '
  PHRASES = [
    "#{LEAD}the house that Jack built.\n",
    "the malt\nthat lay in ",
    "the rat\nthat ate ",
    "the cat\nthat killed ",
    "the dog\nthat worried ",
    "the cow with the crumpled horn\nthat tossed ",
    "the maiden all forlorn\nthat milked ",
    "the man all tattered and torn\nthat kissed ",
    "the priest all shaven and shorn\nthat married ",
    "the rooster that crowed in the morn\nthat woke ",
    "the farmer sowing his corn\nthat kept ",
    "the horse and the hound and the horn\nthat belonged to ",
  ]

  def self.recite
    Array.new(PHRASES.size) { |v| verse(v + 1) }.join("\n")
  end

  def self.verse(number)
    phrase(number - 1)
  end

  def self.phrase(index)
    return PHRASES[index].dup if index.zero?
    phrase(index - 1).insert(LEAD.length, PHRASES[index])
  end
end
