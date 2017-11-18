# frozen_string_literal: false

module FoodChain
  OLD_LADY = 'I know an old lady who swallowed a %{animal}.'
  REASON = 'She swallowed the %{animal} to catch the %{prey}%{details}.'

  ANIMALS = [
    'fly',
    'spider',
    'bird',
    'cat',
    'dog',
    'goat',
    'cow',
    'horse',
  ]

  REASONS = [
    'I don\'t know why she swallowed the fly. Perhaps she\'ll die.',
    'It wriggled and jiggled and tickled inside her.',
    'How absurd to swallow a bird!',
    'Imagine that, to swallow a cat!',
    'What a hog, to swallow a dog!',
    'Just opened her throat and swallowed a goat!',
    'I don\'t know how she swallowed a cow!',
    'She\'s dead, of course!',
  ]

  DETAILS = [
    '',
    ' that wriggled and jiggled and tickled inside her',
    '',
    '',
    '',
    '',
    '',
    '',
  ]

  def self.song
    1.upto(ANIMALS.size).map { |i| verse(i) }.join("\n")
  end

  def self.verse(number)
    lyrics = []

    lyrics << format(OLD_LADY, animal: ANIMALS[number - 1])
    lyrics << REASONS[number - 1] if number > 1

    if number < ANIMALS.size
      (number - 1).downto(1) do |meal|
        lyrics << format(REASON, animal: ANIMALS[meal], prey: ANIMALS[meal - 1], details: DETAILS[meal - 1])
      end

      lyrics << REASONS[0]
    end

    lyrics << ''

    lyrics.join("\n")
  end
end

module BookKeeping
  VERSION = 2
end
