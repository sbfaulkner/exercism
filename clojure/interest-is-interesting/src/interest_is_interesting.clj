(ns interest-is-interesting)

(defn interest-rate
  "Returns the interest rate based on the specified balance."
  [balance]
  (cond
    (neg? balance) -3.213
    (< balance 1000.0) 0.5
    (< balance 5000.0) 1.621
    :else 2.475))

(defn- percent-of
  [balance rate]
  (* (abs balance) (bigdec rate) 0.01M))

(defn annual-balance-update
  "Returns the annual balance update, taking into account the interest rate."
  [balance]
  (+ balance (percent-of balance (interest-rate balance))))

(defn amount-to-donate
  "Returns how much money to donate based on the balance and the tax-free percentage."
  [balance tax-free-percentage]
  (if (pos? balance) (int (* 2 (percent-of balance tax-free-percentage))) 0))
