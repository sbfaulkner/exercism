local vlq = {}

function vlq.decode(encoded)
  decoded = {}
  value = 0
  incomplete = false

  for i, b in ipairs(encoded) do
    incomplete = true
    value = bit32.lshift(value, 7) + bit32.band(b, 0x7F)

    if not bit32.btest(b, 0x80) then
      incomplete = false
      decoded[#decoded+1] = value
      value = 0
    end
  end

  if incomplete then error('incomplete byte sequence') end

  return decoded
end

function vlq.encode(decoded)
  encoded = {}
  index = 1

  for i, value in ipairs(decoded) do
    b = 0

    repeat
      b = bit32.bor(b, bit32.band(value, 0x7F))
      table.insert(encoded, index, b)
      value = bit32.rshift(value, 7)
      b = 0x80
    until value == 0

    index = #encoded + 1
  end

  return encoded
end

return vlq
