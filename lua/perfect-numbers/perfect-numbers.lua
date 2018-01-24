local module = {}

local function factors(number)
  return coroutine.wrap(function ()
    for i = 1, math.floor(number / 2) do
      if number % i == 0 then coroutine.yield(i) end
    end
  end)
end

function module.aliquot_sum(number)
  local sum = 0

  for f in factors(number) do
    sum = sum + f
  end

  return sum
end

function module.classify(number)
  local sum = module.aliquot_sum(number)

  if sum < number then return 'deficient' end
  if sum > number then return 'abundant' end

  return 'perfect'
end

return module
