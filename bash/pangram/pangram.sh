SENTENCE=${1^^*}

for LETTER in {A..Z}; do
  [[ $SENTENCE =~ $LETTER ]] || exit 1
done
