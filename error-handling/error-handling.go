// Package erratum implements functions to dealing with errors
package erratum

// Use takes a ResourceOpener, a input string and tries to open the resource
func Use(o ResourceOpener, input string) (err error) {
	var r Resource

	if r, err = o(); err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		return Use(o, input)
	}

	defer r.Close()
	defer func() {
		if rec := recover(); rec != nil {
			if frobError, ok := rec.(FrobError); ok {
				r.Defrob(frobError.defrobTag)
			}
			err = rec.(error)
		}
	}()

	r.Frob(input)
	return err
}
