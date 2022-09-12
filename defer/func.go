package deferbench

func dosth() {}

func useLoopDefer(n int) {
	for i := 0; i < n; i++ {
		defer dosth()
	}
}

func useFuncDefer(n int) {
	f := func() {
		defer dosth()
	}
	for i := 0; i < n; i++ {
		f()
	}
}

func useNormal(n int) {
	for i := 0; i < n; i++ {
		dosth()
	}
}
