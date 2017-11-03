STATEMENT="$(echo $1 | tr -d ' ')"

QUESTION_RESPONSE="Sure."
SHOUTING_RESPONSE="Whoa, chill out!"
SILENCE_RESPONSE="Fine. Be that way!"
DEFAULT_RESPONSE="Whatever."

if [[ -z ${STATEMENT} ]]; then
  echo $SILENCE_RESPONSE
elif [[ ${STATEMENT} =~ [A-Z] && ${STATEMENT} == ${STATEMENT^^*} ]]; then
  echo $SHOUTING_RESPONSE
elif [[ ${STATEMENT} =~ \?$ ]]; then
  echo $QUESTION_RESPONSE
else
  echo $DEFAULT_RESPONSE
fi
