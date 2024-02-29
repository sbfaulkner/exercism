(ns bob
  (:require [clojure.string :as str]))

(defn- question? [s] (re-find #"\?\s*\z" s))
(defn- yelling? [s] (and (re-find #"[A-Z]" s) (not (re-find #"[a-z]" s))))
(defn- silence? [s] (= (str/trim s) ""))

(defn response-for
  "Determine response when someone talks to Bob."
  [s]
  (cond
    (silence? s) "Fine. Be that way!"
    (and (yelling? s) (question? s)) "Calm down, I know what I'm doing!"
    (question? s) "Sure."
    (yelling? s) "Whoa, chill out!"
    :else "Whatever."))
