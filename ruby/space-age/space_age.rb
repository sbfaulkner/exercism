class SpaceAge
  EARTH_YEAR = 31_557_600.0

  PLANETARY_YEARS = {
    earth: 1.0,
    mercury: 0.2408467,
    venus: 0.61519726,
    mars: 1.8808158,
    jupiter: 11.862615,
    saturn: 29.447498,
    uranus: 84.016846,
    neptune: 164.79132,
  }

  def initialize(seconds)
    @seconds = seconds
  end

  PLANETARY_YEARS.each do |planet, fraction|
    define_method("on_#{planet}") { @seconds / (EARTH_YEAR * fraction) }
  end
end

module BookKeeping
  VERSION = 1
end
