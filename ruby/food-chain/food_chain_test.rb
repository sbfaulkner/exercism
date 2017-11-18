require 'minitest/autorun'

require_relative 'food_chain'

class NoCheating < IOError
  def message
    "The use of File.open and IO.read is restricted.\n"                \
    'This exercise intends to help you improve your ability to work ' \
    'with data generated from your code. Your program must not read ' \
    'the song.txt file.'
  end
end

class FoodChainTest < Minitest::Test
  def setup
    @song_file = File.expand_path('../song.txt', __FILE__)
    @whole_song = IO.read(@song_file)
    @verses = @whole_song.gsub("\n\n", "\n|").split('|')
  end

  def test_verse_1
    assert_equal @verses[0], FoodChain.verse(1)
  end

  def test_verse_2
    assert_equal @verses[1], FoodChain.verse(2)
  end

  def test_verses
    8.times do |i|
      assert_equal @verses[i], FoodChain.verse(i + 1)
    end
  end

  # This test is an acceptance test.
  #
  # If you find it difficult to work the problem with so much
  # output, go ahead and add a `skip`, and write whatever
  # unit tests will help you. Then unskip it again
  # to make sure you got it right.
  # There's no need to submit the tests you write, unless you
  # specifically want feedback on them.
  def test_the_whole_song
    assert_equal @whole_song, FoodChain.song
  end

  # Tests that an error is effectively raised when IO.read or
  # File.open are used within FoodChain.
  def test_read_guard
    ["IO.read '#{@song_file}'", "File.open '#{@song_file}'"].each do |trigger|
      assert_raises(NoCheating) { FoodChain.send :class_eval, trigger }
    end
  end

  # Problems in exercism evolve over time,
  # as we find better ways to ask questions.
  # The version number refers to the version of the problem you solved,
  # not your solution.
  #
  # Define a constant named VERSION inside of BookKeeping.
  # If you are curious, read more about constants on RubyDoc:
  # http://ruby-doc.org/docs/ruby-doc-bundle/UsersGuide/rg/constants.html
  def test_version
    assert_equal 2, BookKeeping::VERSION
  end
end

module RestrictedClasses
  class File
    def self.open(*)
      fail NoCheating
    end

    def self.read(*)
      fail NoCheating
    end

    def open(*)
      fail NoCheating
    end

    def read(*)
      fail NoCheating
    end
  end

  class IO
    def self.read(*)
      fail NoCheating
    end

    def read(*)
      fail NoCheating
    end
  end
end

FoodChain.prepend RestrictedClasses
