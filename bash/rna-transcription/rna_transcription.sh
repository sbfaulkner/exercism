declare -A complements=(
  [G]=C
  [C]=G
  [T]=A
  [A]=U
)

DNA=$1

for (( i=0; i<${#DNA}; i++ )); do
  NUCLEOTIDE=${complements[${DNA:$i:1}]}

  if [[ -z ${NUCLEOTIDE} ]]; then
    echo "Invalid nucleotide detected."
    exit 1
  fi

  RNA=${RNA}${NUCLEOTIDE}
done

echo ${RNA}
