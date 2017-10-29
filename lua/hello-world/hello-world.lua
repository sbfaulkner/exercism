-- Table to be returned by the hello-world module.
local hello_world = {}

-- Add the hello() function to the table returned by this module.
function hello_world.hello()
  return 'Hello, World!'
end

-- Return the hello_world table to make it accessible as a module.
return hello_world
