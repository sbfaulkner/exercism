return function (sentence)
  local letters = {}
  local count = 0

  for letter in sentence:lower():gmatch('%a') do
    if not letters[letter] then
      count = count + 1
      letters[letter] = true
    end
  end

  return count == 26
end