module Strain
  def keep
    kept = []
    each do |item|
      kept << item if yield(item)
    end
    kept
  end

  def discard
    kept = []
    each do |item|
      kept << item unless yield(item)
    end
    kept
  end
end

Array.include Strain
