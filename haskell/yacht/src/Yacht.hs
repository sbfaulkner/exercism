module Yacht (yacht, Category(..)) where

import Data.List

data Category = Ones
              | Twos
              | Threes
              | Fours
              | Fives
              | Sixes
              | FullHouse
              | FourOfAKind
              | LittleStraight
              | BigStraight
              | Choice
              | Yacht

yacht :: Category -> [Int] -> Int
yacht Ones = scoreDie 1
yacht Twos = scoreDie 2
yacht Threes = scoreDie 3
yacht Fours = scoreDie 4
yacht Fives = scoreDie 5
yacht Sixes = scoreDie 6
yacht FullHouse = scoreFullHouse
yacht FourOfAKind = scoreFourOfAKind
yacht LittleStraight = scoreLittleStraight
yacht BigStraight = scoreBigStraight
yacht Choice = sum
yacht Yacht = scoreYacht

scoreDie :: Int -> [Int] -> Int
scoreDie n dice = sum (filter (== n) dice)

scoreFullHouse :: [Int] -> Int
scoreFullHouse dice
    | length grouped == 2 && ( length (head grouped) == 2 || length (head grouped) == 3 ) = sum dice
    | otherwise = 0
    where grouped = group (sort dice)

scoreFourOfAKind :: [Int] -> Int
scoreFourOfAKind dice
    | length grouped <= 2 && length (head grouped) >= 4 = 4 * head (head grouped)
    | length grouped == 2 && length (last grouped) == 4 = sum (last grouped)
    | otherwise = 0
    where grouped = group (sort dice)

scoreLittleStraight :: [Int] -> Int
scoreLittleStraight dice
    | sort dice == [1,2,3,4,5] = 30
    | otherwise = 0

scoreBigStraight :: [Int] -> Int
scoreBigStraight dice
    | sort dice == [2,3,4,5,6] = 30
    | otherwise = 0

scoreYacht :: [Int] -> Int
scoreYacht dice
    | length grouped == 1 = 50
    | otherwise = 0
    where grouped = group (sort dice)
