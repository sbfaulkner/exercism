# frozen_string_literal: true

module PigLatin
  REGEX = /\A(?<prefix>([^aeiouyx]|((?<=q)u)|((?<![^aeiou])y(?!t))|(x(?!r)))+)?(?<stem>.+)\z/

  def self.translate(english)
    english.split.map { |word| translate_word(word) }.join(' ')
  end

  def self.translate_word(word)
    m = REGEX.match(word)
    "#{m[:stem]}#{m[:prefix]}ay"
  end
end

module BookKeeping
  VERSION = 2
end
