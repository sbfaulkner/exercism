local Anagram = {}
Anagram.__index = Anagram

-- normalize word by converting to lowercase and sorting characters
local function normalize(word)
  f = { word:lower():match((word:gsub('.', '(.)'))) }
  table.sort(f)
  return table.concat(f)
end

-- match returns specified words that are anagrams of the root word
function Anagram:match(possibilities)
  matches = {}

  for _, p in pairs(possibilities) do
    if normalize(p) == self.root then matches[#matches+1] = p end
  end

  return matches
end

-- constructor for anagram detector
function Anagram:new(word)
  local anagram = {}
  setmetatable(anagram, Anagram)
  anagram.root = normalize(word)
  return anagram
end

return Anagram
