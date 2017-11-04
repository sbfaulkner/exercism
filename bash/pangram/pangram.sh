SENTENCE=${1^^*}

ALPHABET=(A B C D E F G H I J K L M N O P Q R S T U V W X Y Z)

for LETTER in ${ALPHABET[@]}; do
  [[ $SENTENCE =~ $LETTER ]] || exit 1
done
