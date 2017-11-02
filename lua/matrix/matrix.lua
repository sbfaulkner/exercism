return function (text)
  local matrix = {}

  text:gsub('([^\n]+)', function (line)
    local row = {}

    line:gsub('([^%s]+)', function (cell)
      row[#row+1] = tonumber(cell)
      return ''
    end)

    matrix[#matrix+1] = row
    return ''
  end)

  local function row(r)
    return matrix[r]
  end

  local function column(c)
    local column = {}

    for _, row in ipairs(matrix) do
      column[#column+1] = row[c]
    end

    return column
  end

  return { row = row, column = column }
end
