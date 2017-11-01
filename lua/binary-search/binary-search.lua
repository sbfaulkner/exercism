return function(array, target)
  local l = 1
  local r = #array

  while l <= r do
    local m = math.floor((l+r)/2)
    local am = array[m]

    if am < target then
      l = m + 1
    elseif am > target then
      r = m - 1
    else
      return m
    end
  end

  return -1
end