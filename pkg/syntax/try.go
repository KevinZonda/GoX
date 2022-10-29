package syntax

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
