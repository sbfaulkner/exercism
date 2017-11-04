# Bob simulates a lackadaisical teenager.
module Bob
  def self.hey(remark)
    remark.strip!

    if remark.empty?
      'Fine. Be that way!'
    elsif remark =~ /[a-z]/i && remark == remark.upcase
      'Whoa, chill out!'
    elsif remark.end_with?('?')
      'Sure.'
    else
      'Whatever.'
    end
  end
end

module BookKeeping
  VERSION = 1
end
