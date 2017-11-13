module FlattenArray
  def self.flatten(array)
    array.each_with_object([]) do |item, accumulator|
      next if item.nil?

      if item.is_a?(Array)
        accumulator.concat(flatten(item))
      else
        accumulator << item
      end
    end
  end
end

module BookKeeping
  VERSION = 1
end
