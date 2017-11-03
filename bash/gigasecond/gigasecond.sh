# argument is date of birth (or datetime)
DATE_OF_BIRTH="$(echo $1 | sed 's/\(-[0-9][0-9]\)Z/\1 00:00:00Z/')"

# input format
INPUT_FORMAT='%Y-%m-%d %H:%M:%SZ'

# required output format
OUTPUT_FORMAT='%a %b %e %H:%M:%S %Z %G'

# convert date
date -juR -v +1000000000S -f "$INPUT_FORMAT" "$DATE_OF_BIRTH" +"$OUTPUT_FORMAT" | sed 's/  / /'
