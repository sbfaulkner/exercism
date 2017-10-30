local diff = {}

function diff.square_of_sums(max)
  return (max * (max + 1) / 2)^2
end

function diff.sum_of_squares(max)
  return max * (max + 1) * (2 * max + 1) / 6
end

function diff.difference_of_squares(max)
  return diff.square_of_sums(max) - diff.sum_of_squares(max)
end

return diff
