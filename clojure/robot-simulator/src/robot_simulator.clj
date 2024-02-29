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
(def right {:north :east, :south :west, :east :south, :west :north})

(defn- turn
  "Turn robot left or right."
  [direction {bearing :bearing coordinates :coordinates}]
  (robot coordinates (direction bearing)))

(def instruction {\A advance, \L (partial turn left), \R (partial turn right)})

(defn- execute
  "Execute a single instruction."
  [robot i]
  ((instruction i) robot))

(defn simulate
  "Execute a series of instructions, returning the resulting robot."
  [instructions robot]
  (reduce execute robot instructions))
