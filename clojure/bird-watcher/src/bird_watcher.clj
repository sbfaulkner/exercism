(ns bird-watcher)

(def last-week [0, 2, 5, 3, 7, 8, 4])

(defn today [birds] (last birds))

(defn inc-bird [birds] (update birds (dec (count birds)) inc))

(defn day-without-birds? [birds] (not (every? pos? birds)))

(defn n-days-count [birds n] (apply + (take n birds)))

(defn busy-days [birds] (count (filter (partial < 4) birds)))

(defn odd-week? [birds] (= (take 7 (cycle [1 0])) birds))
