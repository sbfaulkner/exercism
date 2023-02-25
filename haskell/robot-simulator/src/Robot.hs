module Robot
    ( Bearing(East,North,South,West)
    , bearing
    , coordinates
    , mkRobot
    , move
    ) where

data Bearing = North
             | East
             | South
             | West
             deriving (Eq, Show)

data Robot = Robot { bearing :: Bearing, coordinates :: (Integer, Integer) }

mkRobot :: Bearing -> (Integer, Integer) -> Robot
mkRobot d c = Robot d c

move :: Robot -> String -> Robot
move robot instructions = foldl oneMove robot instructions

oneMove :: Robot -> Char -> Robot
oneMove robot 'L' = turnLeft robot
oneMove robot 'R' = turnRight robot
oneMove robot 'A' = advance robot
oneMove _robot _   = error "Invalid instruction"

turnLeft :: Robot -> Robot
turnLeft robot = case (bearing robot) of
                    North -> robot { bearing = West }
                    East  -> robot { bearing = North }
                    South -> robot { bearing = East }
                    West  -> robot { bearing = South }

turnRight :: Robot -> Robot
turnRight robot = case (bearing robot) of
                    North -> robot { bearing = East }
                    East  -> robot { bearing = South }
                    South -> robot { bearing = West }
                    West  -> robot { bearing = North }

advance :: Robot -> Robot
advance robot = case (bearing robot) of
                    North -> robot { coordinates = (fst (coordinates robot), snd (coordinates robot) + 1) }
                    East  -> robot { coordinates = (fst (coordinates robot) + 1, snd (coordinates robot)) }
                    South -> robot { coordinates = (fst (coordinates robot), snd (coordinates robot) - 1) }
                    West  -> robot { coordinates = (fst (coordinates robot) - 1, snd (coordinates robot)) }
