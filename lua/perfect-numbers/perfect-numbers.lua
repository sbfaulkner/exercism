local module = {}

function module.aliquot_sum(number)
  local sum = 0
  local i = 1

  while i < number do
    if number % i == 0 then
      sum = sum + i
    end

    i = i + 1
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
