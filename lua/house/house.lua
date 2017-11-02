local house = {}

local objects = {
  'house that Jack built.',
  'malt',
  'rat',
  'cat',
  'dog',
  'cow with the crumpled horn',
  'maiden all forlorn',
  'man all tattered and torn',
  'priest all shaven and shorn',
  'rooster that crowed in the morn',
  'farmer sowing his corn',
  'horse and the hound and the horn',
}

local actions = {
  '',
  'lay in',
  'ate',
  'killed',
  'worried',
  'tossed',
  'milked',
  'kissed',
  'married',
  'woke',
  'kept',
  'belonged to',
}

-- verse() returns the specified verse.
function house.verse(number)
  local text = 'This is the ' .. objects[number]

  while number > 1 do
    text = text .. '\nthat ' .. actions[number]
    number = number - 1
    text = text .. ' the ' .. objects[number]
  end

  return text
end

-- recite() returns the entire song.
function house.recite()
  local verses = {}

  for i = 1, #objects do
    verses[#verses+1] = house.verse(i)
  end

  return table.concat(verses, '\n')
end

return house
