package erratum

const testVersion = 2

// Use is an artificial test case that exercises various types of error handling
func Use(o ResourceOpener, input string) (err error) {
	var res Resource

	if res, err = o(); err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}

		return
	}

	defer func() {
		if r := recover(); r != nil {
			if f, ok := r.(FrobError); ok {
				res.Defrob(f.defrobTag)
			}
			err = r.(error)
		}

		res.Close()
	}()

	res.Frob(input)

	return
}
