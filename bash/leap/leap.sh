usage() {
  echo "Usage: leap.sh <year>"
  exit 1
}

[[ $# -ne 1 ]] && usage

YEAR=$1

[[ ${YEAR} =~ [^0-9] ]] && usage

if (( ${YEAR} % 4 == 0 && ${YEAR} % 100 != 0 || ${YEAR} % 400 == 0 )); then
  echo "This is a leap year."
else
  echo "This is not a leap year."
fi
