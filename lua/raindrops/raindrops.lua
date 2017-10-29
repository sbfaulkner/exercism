local factors = {
  {3, "Pling"},
  {5, "Plang"},
  {7, "Plong"},
}

function raindrops(number)
  text = ''

  for i = 1, #factors do
    if (number % factors[i][1]) == 0 then text = text .. factors[i][2] end
  end

  return #text > 0 and text or tostring(number)
end

return raindrops
