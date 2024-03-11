(ns squeaky-clean
   (:require [clojure.string :as str]))

 (defn- underscoreSpaces [s] (str/replace s " " "_"))

 (defn- replaceControlCharacters [s] (str/replace s #"[\u0000-\u001f\u007f-\u009f]" "CTRL"))

 (defn- camelize [s] (str/replace s #"-(.)" #(.toUpperCase (%1 1))))

 (defn- omitNonLetters [s] (str/replace s #"[^\p{IsAlphabetic}_]" ""))

 (defn- omitLowerGreek [s] (str/replace s #"[\p{InGreek}&&[^\p{Lu}]]" ""))

 (defn clean
   "Returns a clean identifier name."
   [s]
   (-> s underscoreSpaces replaceControlCharacters camelize omitNonLetters omitLowerGreek))
