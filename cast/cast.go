package cast

//go:noinline
func CastFunc(i interface{}) int {
	return i.(int)
}

//go:noinline
func NonCastFunc(i int) int {
	return i
}
