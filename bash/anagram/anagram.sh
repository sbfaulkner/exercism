normalize() {
  local string=$1

  echo ${string} | fold -w 1 | sort | tr -d '\n'
}

WORD=$1
POSSIBILITIES=$2

ROOT=$(normalize ${WORD,,})
ANAGRAMS=()

for POSSIBILITY in $POSSIBILITIES; do
  NORMALIZED=$(normalize ${POSSIBILITY,,*})
  [[ $NORMALIZED == $ROOT && ${POSSIBILITY,,*} != ${WORD,,*} ]] && ANAGRAMS+=($POSSIBILITY)
done

echo ${ANAGRAMS[@]}
