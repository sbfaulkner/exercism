(ns log-levels
  (:require [clojure.string :as str]))

(def log-format #"\[(.+)\]:\s+(.+?)\s*")

(defn message
  "Takes a string representing a log line
   and returns its message with whitespace trimmed."
  [s]
  (get (re-matches log-format s) 2))

(defn log-level
  "Takes a string representing a log line
   and returns its level in lower-case."
  [s]
  (str/lower-case (get (re-matches log-format s) 1)))

(defn reformat
  "Takes a string representing a log line and formats it
   with the message first and the log level in parentheses."
  [s]
  (format "%s (%s)" (message s) (log-level s)))
