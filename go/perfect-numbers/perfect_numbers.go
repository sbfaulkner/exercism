package perfect

const testVersion = 1

// Classification is used for Nicomachus' classification scheme for natural numbers.
type Classification int

// The number classifications.
const (
	ClassificationPerfect   Classification = 1 // aliquot sum = number
	ClassificationAbundant  Classification = 2 // aliquot sum > number
	ClassificationDeficient Classification = 3 // aliquot sum < number
)

// ClassificationError provides a custom type for Classify errors.
type ClassificationError string

// ErrOnlyPositive is returned if the number passed to Classify is zero.
const ErrOnlyPositive ClassificationError = "number must be positive"

// Classify returns the classification of the provided positive integer.
func Classify(number uint64) (classification Classification, err error) {
	if number == 0 {
		err = ErrOnlyPositive
		return
	}

	sum := aliquotSum(number)

	if sum == number {
		classification = ClassificationPerfect
	} else if sum > number {
		classification = ClassificationAbundant
	} else {
		classification = ClassificationDeficient
	}

	return
}

func (e ClassificationError) Error() string {
	return string(e)
}

func aliquotSum(number uint64) uint64 {
	sum := uint64(0)

	for _, f := range factors(number) {
		if f != number {
			sum += f
		}
	}

	return sum
}

func factors(number uint64) []uint64 {
	factors := []uint64{}

	for i := uint64(1); i*i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i, number/i)
		}
	}

	return factors
}
