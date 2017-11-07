module ETL
  def self.transform(source)
    source.each_with_object({}) do |(score, letters), transformed|
      letters.each { |letter| transformed[letter.downcase] = score }
    end
  end
end

module BookKeeping
  VERSION = 1
end
