return function(array)
  local function append(dest, src)
    for i, item in ipairs(src) do
      if type(item) == "table" then
        append(dest, item)
      else
        table.insert(dest, item)
      end
    end

    return dest
  end

  return append({}, array)
end
