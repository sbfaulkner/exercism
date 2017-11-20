# frozen_string_literal: true

require 'strscan'

class WordProblem
  INTEGER_REGEX = /-?[0-9]+/
  WORD_REGEX = /[a-z]+/i
  OPERATOR_REGEX = Regexp.union(WORD_REGEX, /\?/)

  def initialize(question)
    @scanner = StringScanner.new(question)
  end

  def answer
    @answer ||= solve
  end

  private

  def solve
    answer = nil
    state = :start

    while state
      case state
      when :start
        expect_word('What')
        state = :what
      when :what
        expect_word('is')
        state = :is
      when :is
        answer = expect_integer
        state = :lvalue
      when :lvalue
        state = expect_operator
      when :+
        answer += expect_integer
        state = :lvalue
      when :-
        answer -= expect_integer
        state = :lvalue
      when :*
        answer *= expect_integer
        state = :lvalue
      when :/
        answer /= expect_integer
        state = :lvalue
      else
        raise NotImplementedError
      end
    end

    answer
  end

  def expect_word(word)
    token = get_token(WORD_REGEX)
    raise ArgumentError unless token == word
    token
  end

  def expect_integer
    token = get_token(INTEGER_REGEX)
    raise ArgumentError unless token
    token.to_i
  end

  def expect_operator
    case get_token(OPERATOR_REGEX)
    when 'plus'
      :+
    when 'minus'
      :-
    when 'multiplied'
      expect_word('by')
      :*
    when 'divided'
      expect_word('by')
      :/
    when '?'
      nil
    else
      raise ArgumentError
    end
  end

  def get_token(regex)
    token = @scanner.scan(regex)
    @scanner.skip(/\s+/)
    token
  end
end

module BookKeeping
  VERSION = 1
end
