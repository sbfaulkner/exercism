(ns elyses-destructured-enchantments)

(defn first-card
  "Returns the first card from deck."
  [deck]
  (let [[card1] deck] card1)
)

(defn second-card
  "Returns the second card from deck."
  [deck]
  (let [[_ card2] deck] card2)
)

(defn swap-top-two-cards
  "Returns the deck with first two items reversed."
  [deck]
  (let [[card1 card2 & rest] deck] (concat [card2 card1] rest))
)

(defn discard-top-card
  "Returns a sequence containing the first card and
   a sequence of the remaining cards in the deck."
  [deck]
  (let [[card1 & rest] deck] [card1 rest])
)

(def face-cards
  ["jack" "queen" "king"])

(defn insert-face-cards
  "Returns the deck with face cards between its head and tail."
  [deck]
  (let [[card1 & rest] deck] (concat (if card1 [card1]) face-cards rest))
)
