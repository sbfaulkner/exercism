(ns armstrong-numbers)

(defn- digits [num]
  (map #(Character/digit % 10) (str num)))

(defn- required-sum [num]
  (let [d (digits num) n (count d)]
    (reduce + (map #(Math/pow % n) d))))

(defn armstrong? [num]
  (== num (required-sum num)))
