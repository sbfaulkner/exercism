class Scrabble
  SCORES = {
    1 => 'AEIOULNRST',
    2 => 'DG',
    3 => 'BCMP',
    4 => 'FHVWY',
    5 => 'K',
    8 => 'JX',
    10 => 'QZ',
  }.freeze

  TILE_SCORES = SCORES.each_with_object({}) do |(score, tiles), scores|
    tiles.chars.each do |tile|
      scores[tile] = score
    end
  end

  def self.score(word)
    new(word).score
  end

  def initialize(word)
    @tiles = word.to_s.strip.upcase.chars
  end

  def score
    @tiles.map { |tile| TILE_SCORES[tile] }.sum
  end
end
