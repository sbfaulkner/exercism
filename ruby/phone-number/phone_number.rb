module PhoneNumber
  def self.clean(number)
    clean = number.gsub(/[^0-9]/, '').sub(/\A1/, '')
    return if clean.length != 10 || clean[3] < '2'
    clean
  end
end

module BookKeeping
  VERSION = 2
end
