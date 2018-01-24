declare -A NUCLEOTIDES=(
  [A]=0
  [C]=0
  [G]=0
  [T]=0
)

DNA=$1

for (( i=0; i < ${#DNA}; i++ ))
do
  NUCLEOTIDE=${DNA:i:1}

  if [[ -z ${NUCLEOTIDES[${NUCLEOTIDE}]} ]]
  then
    echo "Invalid nucleotide in strand"
    exit 1
  fi

  (( NUCLEOTIDES[${NUCLEOTIDE}]++ ))
done

for N in "${!NUCLEOTIDES[@]}"
do
  echo $N: ${NUCLEOTIDES[$N]}
done
