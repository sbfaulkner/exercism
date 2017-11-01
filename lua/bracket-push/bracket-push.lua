return {
  valid = function (text)
    local brackets = { ['{'] = '}', ['('] = ')', ['['] = ']' }
    local closing = '})]'

    expecting = {}

    for i = 1, #text do
      local char = text:sub(i,i)
      local right = brackets[char]

      if right then
        table.insert(expecting, right)
      elseif closing:find(char) then
        if char ~= table.remove(expecting) then return false end
      end
    end

    return #expecting == 0
  end
}