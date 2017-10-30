local diff = {}

function diff.square_of_sums(max)
  return (max * (max + 1) / 2)^2
end

function diff.sum_of_squares(max)
  sum = 0

  for i = 1, max do
    sum = sum + i^2
  end

  return sum
end

function diff.difference_of_squares(max)
  return diff.square_of_sums(max) - diff.sum_of_squares(max)
end

return diff
