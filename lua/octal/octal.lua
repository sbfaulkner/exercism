return function (value)
  return {
    to_decimal = function ()
      local decimal = 0
      for i = 1, #value do
        local digit = tonumber(value:sub(i,i))
        if digit == nil or digit > 7 then return 0 end
        decimal = decimal * 8 + digit
      end
      return decimal
    end
  }
end
