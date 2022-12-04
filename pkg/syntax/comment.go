package syntax

func Comment(_ any) {

}

func If(b bool, f func()) {
	if b {
		f()
	}
}