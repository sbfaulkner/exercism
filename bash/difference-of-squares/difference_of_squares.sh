square_of_the_sum() {
  local n=$1
  local sum=0

  for (( i = 1; i <= n; i++ )); do
    let sum=sum+i
  done

  let sum=sum**2

  echo $sum
}

sum_of_the_squares() {
  local n=$1
  local sum=0

  for (( i = 1; i <= n; i++ )); do
    let sum=sum+i**2
  done

  echo $sum
}

N=$1
FLAG=$2

case $FLAG in
-S)
  let SQUARE=$(square_of_the_sum $N)
  echo $SQUARE
  ;;
-s)
  let SUM=$(sum_of_the_squares $N)
  echo $SUM
  ;;
*)
  let DIFFERENCE=$(square_of_the_sum $N)-$(sum_of_the_squares $N)
  echo $DIFFERENCE
  ;;
esac
