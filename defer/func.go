package deferbench

func dosth() {}

func useDefer(n int) {
	for i := 0; i < n; i++ {
		defer dosth()
	}
}

func useNormal(n int) {
	for i := 0; i < n; i++ {
		dosth()
	}
}
