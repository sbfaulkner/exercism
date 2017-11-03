declare -A factors=(
  [3]=Pling
  [5]=Plang
  [7]=Plong
)

number=$1

string=''

for factor in ${!factors[@]}; do
  (( $number % $factor == 0 )) && string="${string}${factors[$factor]}"
done

echo ${string:-$number}
