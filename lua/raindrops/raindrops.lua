local factors = {
  { factor = 3, sound = "Pling" },
  { factor = 5, sound = "Plang" },
  { factor = 7, sound = "Plong" },
}

return function (number)
  local text = ''

  for _, f in ipairs(factors) do
    if (number % f.factor) == 0 then text = text .. f.sound end
  end

  return #text > 0 and text or tostring(number)
end
