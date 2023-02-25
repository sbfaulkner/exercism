module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz 1 = Just 0
collatz n
    | n <= 0 = Nothing
    | even n = fmap succ (collatz (div n 2))
    | otherwise = fmap succ (collatz (3*n+1))
