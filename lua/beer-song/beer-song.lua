local beer = {}

local function bottles(count)
  if count == 0 then return 'No more bottles' end
  if count == 1 then return '1 bottle' end
  return string.format('%d bottles', count)
end

local function beers(count)
  return bottles(count) .. ' of beer'
end

local function take(count)
  if count == 0 then return 'Go to the store and buy some more' end
  if count == 1 then return 'Take it down and pass it around' end
  return 'Take one down and pass it around'
end

function beer.verse(number)
  local b = beers(number)

  return b .. ' on the wall, ' .. b:lower() .. '.\n'
           .. take(number) .. ', ' .. beers((number-1)%100):lower() .. ' on the wall.\n'
end

function beer.sing(first, last)
  local song = {}

  for v = first, last or 0, -1 do
    song[#song+1] = beer.verse(v)
  end

  return table.concat(song, '\n')
end

return beer