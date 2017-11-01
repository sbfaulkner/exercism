return function(array, target)
  local left = 1
  local right = #array

  while left <= right do
    local mid = math.floor((left+right)/2)
    local value = array[mid]

    if value < target then
      left = mid + 1
    elseif value > target then
      right = mid - 1
    else
      return mid
    end
  end

  return -1
end