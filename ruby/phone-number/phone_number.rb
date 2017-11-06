module PhoneNumber
  PHONE_NUMBER = /\A(?:\+?1 ?)?(?:\(([2-9][0-9]{2})\)|([2-9][0-9]{2}))[ .]?([2-9][0-9]{2})(?:[-.]| *)([0-9]{4}) *\z/

  def self.clean(number)
    PHONE_NUMBER.match(number)&.captures&.join
  end
end

module BookKeeping
  VERSION = 2
end
