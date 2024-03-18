(ns card-games)

(defn rounds
  "Takes the current round number and returns
   a `list` with that round and the _next two_."
  [n]
  (list n (+ n 1) (+ n 2))
  )

(defn concat-rounds
  "Takes two lists and returns a single `list`
   consisting of all the rounds in the first `list`,
   followed by all the rounds in the second `list`"
  [l1 l2]
  (concat l1 l2)
  )

(defn contains-round?
  "Takes a list of rounds played and a round number.
   Returns `true` if the round is in the list, `false` if not."
  [l n]
  (boolean (some #{n} l))
  )

(defn- mean
  [l]
  (double (/ (reduce + l) (count l)))
  )

(defn card-average
  "Returns the average value of a hand"
  [hand]
  (mean hand)
  )

(defn- median
  [l]
  (double (nth l (int (Math/floor (/ (count l) 2)))))
  )

(defn approx-average?
  "Returns `true` if average is equal to either one of:
  - Take the average of the _first_ and _last_ number in the hand.
  - Using the median (middle card) of the hand."
  [hand]
  (let [avg (card-average hand)]
    (or (= avg (mean [(first hand) (last hand)]))
        (= avg (median hand))))
  )

(take-nth 2 [1 2 3 4 5 6 7])

(defn average-even-odd?
  "Returns true if the average of the cards at even indexes
   is the same as the average of the cards at odd indexes."
  [hand]
  (= (mean (take-nth 2 hand)) (mean (take-nth 2 (rest hand))))
  )

(defn maybe-double-last
  "If the last card is a Jack (11), doubles its value
   before returning the hand."
  [hand]
  (if (= 11 (last hand)) (concat (drop-last hand) [22]) hand)
  )
