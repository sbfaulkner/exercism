# frozen_string_literal: true

class Element
  attr_reader :datum

  attr_accessor :next

  def initialize(datum)
    @datum = datum
  end
end

class SimpleLinkedList
  include Enumerable

  def initialize(array = [])
    array.each { |datum| push(Element.new(datum)) }
  end

  def push(element)
    element.next = @head
    @head = element
    self
  end

  def pop
    return unless @head

    element = @head
    @head = element.next
    element.next = nil
    element
  end

  def each
    return enum_for(:each) unless block_given?

    element = @head

    while element
      yield element.datum
      element = element.next
    end

    self
  end

  def reverse!
    next_element = @head
    @head = nil

    while element = next_element
      next_element = element.next
      push element
    end

    self
  end
end

module BookKeeping
  VERSION = 1
end
