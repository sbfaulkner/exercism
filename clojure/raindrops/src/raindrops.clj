(ns raindrops)

(defn- div? [n div] (zero? (mod n div)))

(defn convert [n]
  (if (some (partial div? n) [3 5 7])
    (str (if (div? n 3) "Pling" "") (if (div? n 5) "Plang" "") (if (div? n 7) "Plong" ""))
    (str n)))
