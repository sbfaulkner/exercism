local function word_count(phrase)
  local words = {}

  for word in phrase:lower():gmatch("%w+") do
    words[word] = (words[word] or 0) + 1
  end

  return words
end

return { word_count = word_count }
