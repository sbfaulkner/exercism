local diff = {}

function diff.square_of_sums(max)
  sum = 0

  for i = 1, max do
    sum = sum + i
  end

  return sum*sum
end

function diff.sum_of_squares(max)
  sum = 0

  for i = 1, max do
    sum = sum + i*i
  end

  return sum
end

function diff.difference_of_squares(max)
  return diff.square_of_sums(max) - diff.sum_of_squares(max)
end

return diff
