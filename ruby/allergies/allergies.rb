# frozen_string_literal: true

class Allergies
  ALLERGENS = %w[
    eggs
    peanuts
    shellfish
    strawberries
    tomatoes
    chocolate
    pollen
    cats
  ].freeze

  def initialize(allergies)
    @allergies = allergies
  end

  def allergic_to?(allergen)
    raise unless bit = ALLERGENS.index(allergen)
    @allergies & (1 << bit) != 0
  end

  def list
    ALLERGENS.select { |a| allergic_to?(a) }
  end
end

module BookKeeping
  VERSION = 1
end
