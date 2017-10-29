local hamming = {}

-- compute() returns the Hamming distance between the two DNS strands
function hamming.compute(s1, s2)
  if s1:len() ~= s2:len() then return -1 end

  distance = 0

  for i = 1, #s1 do
    if s1:sub(i, i) ~= s2:sub(i, i) then distance = distance + 1 end
  end

  return distance
end

return hamming
