declare -A factors

factors[3]=Pling
factors[5]=Plang
factors[7]=Plong

number=$1

string=''

for factor in ${!factors[@]}; do
  (( $number % $factor == 0 )) && string="${string}${factors[$factor]}"
done

echo ${string:-$number}
