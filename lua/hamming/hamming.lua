local hamming = {}

-- compute() returns the Hamming distance between the two DNS strands
function hamming.compute(strand1, strand2)
  if #strand1 ~= #strand2 then return -1 end

  local distance = 0

  for i = 1, #strand1 do
    if strand1:sub(i, i) ~= strand2:sub(i, i) then distance = distance + 1 end
  end

  return distance
end

return hamming
