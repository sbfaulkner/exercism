local etl = {}

function etl.transform(old)
  local new = {}

  for score, letters in pairs(old) do
    for _, letter in pairs(letters) do
      new[letter:lower()] = score
    end
  end

  return new
end

return etl
