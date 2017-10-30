local bob = {}

function bob.hey(text)
  if #text == 0 then return 'Fine, be that way.' end

  if text:sub(-1) == '?' then return 'Sure' end

  if text:upper() == text then return 'Whoa, chill out!' end

  return 'Whatever'
end

return bob
