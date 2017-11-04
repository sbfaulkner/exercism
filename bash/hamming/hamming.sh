split() {
  echo $1 | fold -w 1
}

if [[ $# -ne 2 ]]; then
  echo "Usage: hamming.sh <string1> <string2>"
  exit 1
fi

STRAND1=($(split $1))
STRAND2=($(split $2))
LENGTH=${#STRAND1[@]}

if [[ ${LENGTH} -ne ${#STRAND2[@]} ]]; then
  echo "The two strands must have the same length."
  exit 1
fi

DISTANCE=0

for (( i=0; i < ${LENGTH}; i++ )); do
  [[ ${STRAND1[$i]} != ${STRAND2[$i]} ]] && let DISTANCE+=1
done

echo $DISTANCE
