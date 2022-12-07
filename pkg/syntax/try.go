package syntax

// TryCatch is a try-catch statement.
func TryCatch(try func(), catch func(any)) (ok bool) {
	ok = true
	defer func() {
		if r := recover(); r != nil {
			ok = false
			catch(r)
		}
	}()
	try()
	return
}

// TryCatchFinally is a try-catch-finally statement.
func TryCatchFinally(try func(), catch func(any), finally func()) (ok bool) {
	ok = true
	defer func() {
		if r := recover(); r != nil {
			ok = false
			catch(r)
		}
		finally()
	}()
	try()
	return
}

// TryFinally is a try-finally statement.
func TryFinally(try func(), finally func()) {
	defer finally()
	try()
}
