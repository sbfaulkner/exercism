module Strain
  def keep
    each_with_object([]) do |item, kept|
      kept << item if yield(item)
    end
  end

  def discard
    each_with_object([]) do |item, kept|
      kept << item unless yield(item)
    end
  end
end

Array.include Strain
