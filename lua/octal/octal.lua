function Octal(value)
  local to_decimal = function ()
                       decimal = 0
                       for i = 1, #value do
                         digit = tonumber(value:sub(i,i))
                         if digit == nil or digit > 7 then return 0 end
                         decimal = decimal * 8 + digit
                       end
                       return decimal
                     end

  return { to_decimal = to_decimal }
end

return Octal
