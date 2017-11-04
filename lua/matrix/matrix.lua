return function (text)
  local matrix = {}

  for line in text:gmatch('[^\n]+') do
    local row = {}

    for cell in line:gmatch('[^%s]+') do
      row[#row+1] = tonumber(cell)
    end

    matrix[#matrix+1] = row
  end

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
