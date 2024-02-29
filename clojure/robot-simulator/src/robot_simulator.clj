(ns robot-simulator)

(defn robot
  "Initialize a robot."
  [coordinates bearing]
  (assoc {} :bearing bearing :coordinates coordinates)
  )

(def dx {:north 0, :south 0, :east 1, :west -1})
(def dy {:north 1, :south -1, :east 0, :west 0})

(defn- advance
  "Move robot forward one position in current direction."
  [{bearing :bearing {x :x y :y} :coordinates}]
    (robot {:x (+ x (dx bearing)), :y (+ y (dy bearing))} bearing))

(def left {:north :west, :south :east, :east :north, :west :south})

(defn- turn-left
  "Turn robot to the left."
  [{bearing :bearing coordinates :coordinates}]
  (robot coordinates (left bearing)))

(def right {:north :east, :south :west, :east :south, :west :north})

(defn- turn-right
  "Turn robot to the right."
  [{bearing :bearing coordinates :coordinates}]
  (robot coordinates (right bearing)))

(def instruction {\A advance, \L turn-left, \R turn-right})

(defn- execute
  "Execute a single instruction."
  [robot i]
  ((instruction i) robot))

(defn simulate
  "Execute a series of instructions, returning the resulting robot."
  [instructions robot]
  (reduce execute robot instructions))
