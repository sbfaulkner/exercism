return function (phrase)
  local letters = {}

  for letter in phrase:lower():gmatch('%a') do
    if letters[letter] then
      return false
    end

    letters[letter] = true
  end

  return true
end