package prelude

func Must(err error) {
	Assert(err == nil, err)
}

func Must1[T any](x T, err error) T {
	Assert(err == nil, err)
	return x
}

func Must2[T, U any](x T, y U, err error) (T, U) {
	Assert(err == nil, err)
	return x, y
}

func Assert(cond bool, msg ...any) {
	if !cond {
		// todo: allow Sprintf-style arguments here, with the first being a format string?
		panic(msg)
	}
}
